package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/efemero/cdp-equalizer/blockchain"
	"github.com/efemero/cdp-equalizer/cdp"
	"github.com/efemero/cdp-equalizer/transaction"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	goji "goji.io"
	"goji.io/pat"
	"goji.io/pattern"
)

var (
	maxLimit *big.Float
	minLimit *big.Float
	target   *big.Float

	client    *blockchain.Client
	mycdp     *cdp.CDP
	ethPrice  *big.Float
	pethRatio *big.Float
	cdpID     int64
)

func main() {
	log.SetOutput(os.Stdout)

	var err error
	node := "wss://mainnet.infura.io/ws/v3/1f0c50a057c040a0879468a513501c6d"
	privateKey := os.Getenv("CDPPK")
	cdpID, err = strconv.ParseInt(os.Getenv("CDPID"), 10, 0)
	if err != nil {
		log.Fatal("define an environment variable `CDPID` with your CDP ID, err: ", err)
	}
	proxyAddress := os.Getenv("PROXY")
	if proxyAddress == "" {
		log.Fatal("define an environment variable `PROXY` with your dp-proxy address")
	}

	client, err = blockchain.NewClient(node, privateKey, cdpID, proxyAddress)
	if err != nil {
		log.Fatalf("error while creating client,  error: %v", err)
	}

	client.ApproveDai()
	client.ApproveMkr()

	_, mycdp, pethRatio, ethPrice, err = getBase(client)
	if err != nil {
		log.Fatalf("error while getting base informations, error: %v", err)
	}
	go watchCDP(client)

	launchServer()
}

func launchServer() {
	mux := goji.NewMux()

	mux.HandleFunc(pat.Get("/"), cdpJSON)
	mux.HandleFunc(pat.Get("/:cdpID"), cdpJSON)
	http.ListenAndServe("localhost:8000", mux)
}

func cdpJSON(w http.ResponseWriter, r *http.Request) {
	var currentCDP *cdp.CDP
	value := ""
	if r != nil && r.Context().Value(pattern.Variable("cdpID")) != nil {
		value = pat.Param(r, "cdpID")
		log.Println(value)
	}
	if value == "" {
		currentCDP = mycdp
	} else {
		cdpID, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		currentCDP, err = client.GetCDP(cdpID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	status, err := currentCDP.GetStatus(ethPrice, pethRatio, target)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// watchCDP make some requests every 15 seconds to see if the CDP must be equalized
func watchCDP(client *blockchain.Client) {
	var i = 0
	mycdp.Log(ethPrice, pethRatio, target)

	ticker := time.NewTicker(time.Second * 15)
	for range ticker.C {
		i++
		tx, cdp, pethRatio, ethPrice, err := getBase(client)
		mycdp = cdp
		if err != nil {
			log.Println(err)
			continue
		}

		if i%60 == 0 {
			cdp.Log(ethPrice, pethRatio, target)
		}

		if tx == nil {

			eth, err := client.GetEthBalance()
			if err != nil {
				log.Println(err)
				continue
			}
			if eth.Cmp(big.NewFloat(0.25)) < 0 {
				tx, err := client.FreeEth(big.NewFloat(0.3), cdp)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println("Starting free 0.3 ETH")
				transaction.SaveTx(&transaction.Tx{Hash: tx.Hash().Hex(), Kind: transaction.FreeEth})
				continue
			}
			ratio := cdp.GetRatio(ethPrice, pethRatio)

			if ratio.Cmp(minLimit) < 0 {
				cdp.Log(ethPrice, pethRatio, target)
				tx, err := client.DoFreeEth(cdp, ethPrice, target)
				if err != nil {
					log.Println(err)
					continue
				}
				transaction.SaveTx(tx)
			} else if ratio.Cmp(maxLimit) > 0 {
				cdp.Log(ethPrice, pethRatio, target)
				tx, err := client.DoDrawDai(cdp, ethPrice, pethRatio, target)
				if err != nil {
					log.Println(err)
					continue
				}
				transaction.SaveTx(tx)
				continue
			}

		} else {
			receipt, err := checkTransaction(client, tx, nil)
			if err != nil {
				log.Println(err)
				continue
			}
			if receipt == nil {
				// transaction not yet mined
				continue
			}
			if tx.Kind == transaction.FreeEth {
				if receipt.Status != 1 {
					log.Println("could not free ETH")
				} else {
					transaction.SaveTx(nil)
					log.Println("OK")
				}
				continue
			}
			if tx.Kind == transaction.DrawSellLock {
				if receipt.Status != 1 {
					ratio := cdp.GetRatio(ethPrice, pethRatio)
					if ratio.Cmp(minLimit) < 0 {
						tx, err := client.DoFreeEth(cdp, ethPrice, target)
						if err != nil {
							log.Println(err)
							continue
						}
						transaction.SaveTx(tx)
						continue
					} else if ratio.Cmp(maxLimit) > 0 {
						tx, err := client.DoDrawDai(cdp, ethPrice, pethRatio, target)
						if err != nil {
							log.Println(err)
							continue
						}
						transaction.SaveTx(tx)
						continue
					}
					continue
				} else {
					transaction.SaveTx(nil)
					log.Println("OK")
					cdp.Log(ethPrice, pethRatio, target)
					continue
				}
			} else if tx.Kind == transaction.FreeSellWipe {
				if receipt.Status != 1 {
					ratio := cdp.GetRatio(ethPrice, pethRatio)
					if ratio.Cmp(minLimit) < 0 {
						tx, err := client.DoFreeEth(cdp, ethPrice, target)
						if err != nil {
							log.Println(err)
							continue
						}
						transaction.SaveTx(tx)
						continue
					} else if ratio.Cmp(maxLimit) > 0 {
						tx, err := client.DoDrawDai(cdp, ethPrice, pethRatio, target)
						if err != nil {
							log.Println(err)
							continue
						}
						transaction.SaveTx(tx)
						continue
					}
					continue
				} else {
					transaction.SaveTx(nil)
					log.Println("OK")
					cdp.Log(ethPrice, pethRatio, target)
					continue
				}
			}
		}

	}
}

func init() {
	max := flag.Float64("maxRatio", 2.15, "The maximum ratio of your CDP")
	min := flag.Float64("minRatio", 1.9, "The minimum ratio of your CDP")
	ratio := flag.Float64("targetRatio", 2.0, "The target ratio of your CDP")
	flag.Parse()
	maxLimit = big.NewFloat(*max)
	minLimit = big.NewFloat(*min)
	target = big.NewFloat(*ratio)
}

func getBase(client *blockchain.Client) (tx *transaction.Tx, cdp *cdp.CDP, pethRatio, ethPrice *big.Float,
	err error) {

	tx, err = transaction.GetTx()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "cannot retrieve transaction")
	}
	cdp, err = client.GetCDP(cdpID)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "cannot retrieve CDP")
	}

	pethRatio, err = client.GetPethRatio()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "cannot retrieve PETH / ETH ratio")
	}

	ethPrice, err = client.GetEthPrice()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "cannot retrieve ETH price")
	}
	return
}

func checkTransaction(client *blockchain.Client, tx *transaction.Tx, header *types.Header) (receit *types.Receipt, err error) {
	receipt, err := client.GetReceipt(tx.Hash)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot retrieve transaction `%s`", tx.Hash)
	}
	/*
		var block *types.Block
		i := 0
		for i < 20 {
			block, err = client.Client.BlockByHash(context.Background(), header.Hash())
			if err == nil {
				break
			} else if err == ethereum.NotFound {
				i++
				continue
			} else {
				return nil, errors.Wrapf(err, "cannot retrieve block `%s`", header.Hash().Hex())
			}
		}
		if block.Transaction(common.HexToHash(tx.Hash)) != nil {
			return nil, errors.Wrapf(err, "Transaction `%s` found in current block, waiting 1 confirmation", tx.Hash)
		}*/
	return receipt, nil

}
