package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"math/big"
	"net/http"
	"os"
	"sort"
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

const cdpTmpl = `<!doctype html>
<html>
<head>
<meta charset="utf-8" />
<title>CDP {{.Current.ID}} ({{.Current.Ratio}})</title>
</head>
<body class="activity-stream">
<h1>Status of CDP {{.Current.ID}}<h1>
<h2>Price: {{.Current.Price}} Net: {{.Current.EthNet}}Ξ ({{.Current.DaiNet}}DAI)</h2>
<h2>Ratio: {{.Current.Ratio}}</h2>
  <ul>
    <h2>Key prices</h2>
{{range .KeyPrices}}
<li>Price: {{.Price}} Net: {{.EthNet}}Ξ ({{.DaiNet}}DAI)</li>
{{end}}
  </ul>
  <ul style="float: left; border: 2px solid green;">
    <h2>UP</h2>
{{range .Up}}
<li>Price: {{.Price}} Net: {{.EthNet}}Ξ ({{.DaiNet}}DAI)</li>
{{end}}
  </ul>
  <ul style="float: left; border: 2px solid red;">
    <h2>DOWN</h2>
{{range .Down}}
<li>Price: {{.Price}} Net: {{.EthNet}}Ξ ({{.DaiNet}}DAI)</li>
{{end}}
  </ul>
  </body>
</html>`

var (
	maxLimit     *big.Float
	minLimit     *big.Float
	target       *big.Float
	templatesDir string

	client    *blockchain.Client
	mycdp     *cdp.CDP
	ethPrice  *big.Float
	pethRatio *big.Float
	cdpID     int64
	keyPrices = []float64{10, 50, 100, 150, 200, 300, 400, 500, 750, 1000, 1500, 2000}
)

type statuses struct {
	Current   *cdp.Status
	KeyPrices []*cdp.Status
	Up        []*cdp.Status
	Down      []*cdp.Status
}

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

	jsonMux := goji.SubMux()
	cdpMux := goji.SubMux()

	// HTML
	mux.Handle(pat.New("/cdp/*"), cdpMux)
	cdpMux.HandleFunc(pat.Get("/"), showCDP)
	cdpMux.HandleFunc(pat.Get("/:cdpID"), showCDP)

	// JSON
	cdpMux.Handle(pat.New("/json/*"), jsonMux)
	jsonMux.HandleFunc(pat.Get("/"), cdpJSON)
	jsonMux.HandleFunc(pat.Get("/:cdpID"), cdpJSON)
	http.ListenAndServe("localhost:8000", mux)
}

func showCDP(w http.ResponseWriter, r *http.Request) {
	cdps := getStatuses(w, r)

	t, err := template.New("cdp").Parse(cdpTmpl)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, cdps)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
}

func cdpJSON(w http.ResponseWriter, r *http.Request) {
	cdps := getStatuses(w, r)
	js, err := json.Marshal(cdps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getStatuses(w http.ResponseWriter, r *http.Request) *statuses {
	var (
		currentCDP               *cdp.CDP
		nextPrice, previousPrice *big.Float
	)
	value := ""
	if r != nil && r.Context().Value(pattern.Variable("cdpID")) != nil {
		value = pat.Param(r, "cdpID")
	}
	if value == "" || value == strconv.Itoa(int(mycdp.ID)) {
		currentCDP = mycdp
	} else {
		cdpID, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}
		currentCDP, err = client.GetCDP(cdpID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
	}

	cdps := statuses{}

	status, err := currentCDP.GetStatus(ethPrice, pethRatio, target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	cdps.Current = status

	// Up
	currentPrice := new(big.Float).Copy(ethPrice)
	nextCDP := &cdp.CDP{
		ID:      currentCDP.ID,
		BytesID: currentCDP.BytesID,
		DaiDebt: currentCDP.DaiDebt,
		PethCol: currentCDP.PethCol,
		EthCol:  currentCDP.EthCol,
	}
	cdps.Up = []*cdp.Status{}

	for currentPrice.Cmp(big.NewFloat(2000.0)) < 0 {
		previousPrice, currentPrice = nextCDP.GetChangePrices(currentPrice, minLimit, maxLimit, pethRatio)
		nextCDP, err = nextCDP.EqualizeCDP(currentPrice, target, pethRatio)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
		nextStatus, err := nextCDP.GetStatus(currentPrice, pethRatio, target)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
		cdps.Up = append(cdps.Up, nextStatus)
	LoopKPUp:
		for _, kp := range keyPrices {
			// check if a keyprice is in the nextStatus
			if currentPrice.Cmp(big.NewFloat(kp)) > 0 && previousPrice.Cmp(big.NewFloat(kp)) < 0 {
				for _, status := range cdps.KeyPrices {
					if big.NewFloat(kp).Cmp((*big.Float)(status.Price)) == 0 {
						break LoopKPUp
					}
				}
				keyStatus, err := nextCDP.GetStatus(big.NewFloat(kp), pethRatio, target)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return nil
				}
				cdps.KeyPrices = append(cdps.KeyPrices, keyStatus)
				break
			}
		}
	}

	// Down
	currentPrice = new(big.Float).Copy(ethPrice)
	nextCDP = &cdp.CDP{
		ID:      currentCDP.ID,
		BytesID: currentCDP.BytesID,
		DaiDebt: currentCDP.DaiDebt,
		PethCol: currentCDP.PethCol,
		EthCol:  currentCDP.EthCol,
	}
	cdps.Down = []*cdp.Status{}
	for currentPrice.Cmp(big.NewFloat(10.0)) > 0 {
		currentPrice, nextPrice = nextCDP.GetChangePrices(currentPrice, minLimit, maxLimit, pethRatio)
		nextCDP, err = nextCDP.EqualizeCDP(currentPrice, target, pethRatio)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
		nextStatus, err := nextCDP.GetStatus(currentPrice, pethRatio, target)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
		cdps.Down = append(cdps.Down, nextStatus)
	LoopKPDown:
		for _, kp := range keyPrices {
			// check if a keyprice is in the nextStatus
			if nextPrice.Cmp(big.NewFloat(kp)) > 0 && currentPrice.Cmp(big.NewFloat(kp)) < 0 {
				for _, status := range cdps.KeyPrices {
					if big.NewFloat(kp).Cmp((*big.Float)(status.Price)) == 0 {
						break LoopKPDown
					}
				}
				keyStatus, err := nextCDP.GetStatus(big.NewFloat(kp), pethRatio, target)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return nil
				}
				cdps.KeyPrices = append(cdps.KeyPrices, keyStatus)
				break
			}
		}
	}
	sort.Slice(cdps.KeyPrices, func(i, j int) bool {
		return (*big.Float)(cdps.KeyPrices[i].Price).Cmp((*big.Float)(cdps.KeyPrices[j].Price)) < 0
	})
	return &cdps

}

// watchCDP make some requests every 15 seconds to see if the CDP must be equalized
func watchCDP(client *blockchain.Client) {
	var (
		err error
		tx  *transaction.Tx
	)
	var i = 0
	mycdp.Log(ethPrice, pethRatio, target)

	ticker := time.NewTicker(time.Second * 15)
	for range ticker.C {
		i++
		tx, mycdp, pethRatio, ethPrice, err = getBase(client)
		if err != nil {
			log.Println(err)
			continue
		}

		if i%60 == 0 {
			mycdp.Log(ethPrice, pethRatio, target)
		}

		if tx == nil {

			eth, err := client.GetEthBalance()
			if err != nil {
				log.Println(err)
				continue
			}
			if eth.Cmp(big.NewFloat(0.25)) < 0 {
				tx, err := client.FreeEth(big.NewFloat(0.3), mycdp)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println("Starting free 0.3 ETH")
				transaction.SaveTx(&transaction.Tx{Hash: tx.Hash().Hex(), Kind: transaction.FreeEth})
				continue
			}
			ratio := mycdp.GetRatio(ethPrice, pethRatio)

			if ratio.Cmp(minLimit) < 0 {
				mycdp.Log(ethPrice, pethRatio, target)
				tx, err := client.DoFreeEth(mycdp, ethPrice, target)
				if err != nil {
					log.Println(err)
					continue
				}
				transaction.SaveTx(tx)
			} else if ratio.Cmp(maxLimit) > 0 {
				mycdp.Log(ethPrice, pethRatio, target)
				tx, err := client.DoDrawDai(mycdp, ethPrice, pethRatio, target)
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
					ratio := mycdp.GetRatio(ethPrice, pethRatio)
					if ratio.Cmp(minLimit) < 0 {
						tx, err := client.DoFreeEth(mycdp, ethPrice, target)
						if err != nil {
							log.Println(err)
							continue
						}
						transaction.SaveTx(tx)
						continue
					} else if ratio.Cmp(maxLimit) > 0 {
						tx, err := client.DoDrawDai(mycdp, ethPrice, pethRatio, target)
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
					mycdp.Log(ethPrice, pethRatio, target)
					continue
				}
			} else if tx.Kind == transaction.FreeSellWipe {
				if receipt.Status != 1 {
					ratio := mycdp.GetRatio(ethPrice, pethRatio)
					if ratio.Cmp(minLimit) < 0 {
						tx, err := client.DoFreeEth(mycdp, ethPrice, target)
						if err != nil {
							log.Println(err)
							continue
						}
						transaction.SaveTx(tx)
						continue
					} else if ratio.Cmp(maxLimit) > 0 {
						tx, err := client.DoDrawDai(mycdp, ethPrice, pethRatio, target)
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
					mycdp.Log(ethPrice, pethRatio, target)
					continue
				}
			}
		}

	}
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

func init() {
	max := flag.Float64("maxRatio", 2.15, "The maximum ratio of your CDP")
	min := flag.Float64("minRatio", 1.9, "The minimum ratio of your CDP")
	ratio := flag.Float64("targetRatio", 2.0, "The target ratio of your CDP")
	templates := flag.String("templatesDir", "/opt/cdp-equalizer/templates", "The base directory of the templates")
	flag.Parse()
	maxLimit = big.NewFloat(*max)
	minLimit = big.NewFloat(*min)
	target = big.NewFloat(*ratio)
	templatesDir = *templates
}
