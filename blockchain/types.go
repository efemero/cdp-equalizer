package blockchain

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"

	"github.com/Efemero/cdp-equalizer/cdp"
	"github.com/Efemero/cdp-equalizer/dsproxy"
	"github.com/Efemero/cdp-equalizer/equalizer"
	"github.com/Efemero/cdp-equalizer/erc20"
	"github.com/Efemero/cdp-equalizer/ethutils"
	"github.com/Efemero/cdp-equalizer/maker"
	"github.com/Efemero/cdp-equalizer/oasis"
	"github.com/Efemero/cdp-equalizer/oracle"
	"github.com/Efemero/cdp-equalizer/portal"
	"github.com/Efemero/cdp-equalizer/transaction"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

var (
	makerAddress       = common.HexToAddress("0x448a5065aeBB8E423F0896E6c5D525C040f59af3")
	oracleAddress      = common.HexToAddress("0x729D19f657BD0614b4985Cf1D82531c67569197B")
	stationAddress     = common.HexToAddress("0xdc919494349E803fbd2D4064106944418381EDb3")
	daiAddress         = common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359")
	pethAddress        = common.HexToAddress("0xf53ad2c6851052a81b42133467480961b2321c09")
	wethAddress        = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	mkrAddress         = common.HexToAddress("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2")
	oasisDirectAddress = common.HexToAddress("0x279594b6843014376a422ebb26a6eab7a30e36f0")
	otcAddress         = common.HexToAddress("0xb7ac09c2c0217b07d7c103029b4918a2c401eecb")
	portalAddress      = common.HexToAddress("0xd64979357160E8146F6e1d805cf20437397bF1ba")
	equalizerAddress   = common.HexToAddress("0xF1E1d750137Ae5c1bD91fe7Bd0da692a3Ed1d553")
)

// Client is a struct that can make some calls to some contracts
type Client struct {
	Client         *ethclient.Client
	dsproxy        *dsproxy.Dsproxy
	dsproxyAddress common.Address
	maker          *maker.Maker
	oracle         *oracle.Oracle
	portal         *portal.Portal
	equalizer      *equalizer.Equalizer
	dai            *erc20.Erc20
	peth           *erc20.Erc20
	mkr            *erc20.Erc20
	cdpID          int64
	privateKey     *ecdsa.PrivateKey
	publicKey      *ecdsa.PublicKey
	address        common.Address
}

// NewClient initialize a new Client, using an ethClient
func NewClient(node string, privateKey string, cdpID int64, proxy string) (client *Client, err error) {
	client = new(Client)
	client.cdpID = cdpID
	client.Client, err = ethclient.Dial(node)
	if err != nil {
		log.Fatal("be sure to have an ethereum node running", err)
	}
	client.dsproxyAddress = common.HexToAddress(proxy)
	client.dsproxy, err = dsproxy.NewDsproxy(client.dsproxyAddress, client.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "could not instantiate dsproxy at address '%s'", client.dsproxyAddress.Hex())
	}
	client.maker, err = maker.NewMaker(makerAddress, client.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "could not instantiate maker at address '%s'", makerAddress.Hex())
	}
	client.oracle, err = oracle.NewOracle(oracleAddress, client.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "could not instantiate oracle at address '%s'", oracleAddress.Hex())
	}
	client.equalizer, err = equalizer.NewEqualizer(equalizerAddress, client.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "could not instantiate equalizer at address '%s'", equalizerAddress.Hex())
	}
	client.portal, err = portal.NewPortal(portalAddress, client.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "could not instantiate portal at address '%s'", portalAddress.Hex())
	}
	client.dai, err = erc20.NewErc20(daiAddress, client.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "could not instantiate DAI token contract at address '%s'", daiAddress.Hex())
	}
	client.peth, err = erc20.NewErc20(pethAddress, client.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "could not instantiate PETH token contract at address '%s'", pethAddress.Hex())
	}
	client.mkr, err = erc20.NewErc20(mkrAddress, client.Client)
	if err != nil {
		return nil, errors.Wrapf(err, "could not instantiate MKR token contract at address '%s'", mkrAddress.Hex())
	}
	client.privateKey, err = crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrapf(err, "`%s` is not a valid private key", privateKey)
	}

	publicKey := client.privateKey.Public()
	var ok bool
	client.publicKey, ok = publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.Wrap(err, "error casting public key to ECDSA")
	}
	client.address = crypto.PubkeyToAddress(*client.publicKey)

	return client, nil
}

// GetHeaderChan returns a channel that is filled with the headers of new blocks
func (client *Client) GetHeaderChan() (chan *types.Header, error) {
	headers := make(chan *types.Header)
	_, err := client.Client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		close(headers)
		return nil, errors.Wrap(err, "cannot initialize header chan")
	}
	return headers, nil
}

// GetAuth returns a *bind.TransactOpts to go to signed transactions
func (client *Client) GetAuth() (*bind.TransactOpts, error) {

	nonce, err := client.Client.PendingNonceAt(context.Background(), client.address)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get nonce")
	}
	gasPrice, err := client.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "cannot get gas price suggestion")
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2))

	auth := bind.NewKeyedTransactor(client.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth, nil
}

// GetCDP returns the given CDP
func (client *Client) GetCDP() (*cdp.CDP, error) {
	var cdp = cdp.CDP{ID: client.cdpID}
	copy(cdp.BytesID[:], abi.U256(big.NewInt(client.cdpID)))
	result, err := client.maker.Cups(&bind.CallOpts{}, cdp.BytesID)
	if err != nil {
		return nil, errors.Wrapf(err, "could not retrieve informations about cdp.CDP `%d`", client.cdpID)
	}
	cdp.DaiDebt = ethutils.FromWei(result.Art, 18)
	cdp.PethCol = ethutils.FromWei(result.Ink, 18)
	pethRatio, err := client.GetPethRatio()
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve peth ratio")
	}
	cdp.EthCol = new(big.Float).Mul(cdp.PethCol, pethRatio)

	return &cdp, nil
}

// GetPethRatio retrieve the Peth / Eth ratio on the blockchain
func (client *Client) GetPethRatio() (ratio *big.Float, err error) {
	ptoeWei, err := client.maker.Per(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve the Peth / Eth ratio on the blockchain")
	}
	return ethutils.FromWei(ptoeWei, 27), nil
}

// GetEthPrice returns the price from the makerDAO oracle
func (client *Client) GetEthPrice() (eth *big.Float, err error) {
	result, err := client.oracle.Read(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "could not read the price from oracle")
	}

	wei := new(big.Int).SetBytes(result[:])
	eth = ethutils.FromWei(wei, 18)
	return eth, nil
}

// IsMkrAllowed returns true if the CDP Station contract is allowed to use our PETH
func (client *Client) IsMkrAllowed() (allowed bool, err error) {
	result, err := client.mkr.Allowance(&bind.CallOpts{}, client.address, client.dsproxyAddress)
	if err != nil {
		return false, errors.Wrap(err, "could not read the MKR status")
	}
	allowed = result.Cmp(ethutils.ToWei(big.NewFloat(1000000), 18)) > 0
	return allowed, nil
}

// IsPethAllowed returns true if the CDP Station contract is allowed to use our PETH
func (client *Client) IsPethAllowed() (allowed bool, err error) {
	result, err := client.peth.Allowance(&bind.CallOpts{}, client.address, stationAddress)
	if err != nil {
		return false, errors.Wrap(err, "could not read the PETH status")
	}
	allowed = result.Cmp(ethutils.ToWei(big.NewFloat(1000000), 18)) > 0
	return allowed, nil
}

// DrawSellLock retrieve DAI from CDP, sells it on oasis.direct and locks ETH in the CDP
func (client Client) DrawSellLock(dai *big.Float) (*types.Transaction, error) {
	auth, err := client.GetAuth()
	if err != nil {
		return nil, errors.Wrap(err, "could not get auth for transaction DrawSellLock")
	}

	a, err := abi.JSON(strings.NewReader(equalizer.EqualizerABI))
	if err != nil {
		return nil, errors.Wrap(err, "could not read equalizer ABI")
	}

	cdpIDb, err := ethutils.SliceToBytes32(big.NewInt(client.cdpID).Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "could not get [32]byte")
	}

	minLockDai := new(big.Float).Mul(dai, big.NewFloat(0.9))
	price, err := client.GetEthPrice()
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve ETH price")
	}
	minLock := new(big.Float).Quo(minLockDai, price)
	b, err := a.Pack("drawSellLock", makerAddress, otcAddress, cdpIDb, daiAddress, ethutils.ToWei(dai, 18), wethAddress, ethutils.ToWei(minLock, 18))
	if err != nil {
		return nil, errors.Wrapf(err, "could not pack data")
	}
	tx, err := client.dsproxy.Execute(auth, equalizerAddress, b)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrapf(err, "could not DrawSellLock %.2f DAI in CDP #%d", dai, client.cdpID)
	}
	return tx, nil
}

// FreeSellWipe retrieve ETH from CDP, sells it on oasis.direct and wipe DAI in the CDP
func (client Client) FreeSellWipe(eth *big.Float) (*types.Transaction, error) {
	auth, err := client.GetAuth()
	if err != nil {
		return nil, errors.Wrap(err, "could not get auth for transaction FreeSellWipe")
	}

	a, err := abi.JSON(strings.NewReader(equalizer.EqualizerABI))
	if err != nil {
		return nil, errors.Wrapf(err, "could not read equalizer ABI")
	}

	cdpIDb, err := ethutils.SliceToBytes32(big.NewInt(client.cdpID).Bytes())
	if err != nil {
		return nil, errors.Wrapf(err, "could not get [32]byte")
	}

	minWipe := new(big.Float).Mul(eth, big.NewFloat(0.9))
	b, err := a.Pack("freeSellWipe", makerAddress, otcAddress, cdpIDb, daiAddress, ethutils.ToWei(eth, 18), wethAddress, ethutils.ToWei(minWipe, 18))
	if err != nil {
		return nil, errors.Wrapf(err, "could not pack data")
	}
	tx, err := client.dsproxy.Execute(auth, equalizerAddress, b)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrapf(err, "could not freeSellWipe %.5f ETH in CDP #%d", eth, client.cdpID)
	}
	return tx, nil
}

// LockEth send eth amount into the CDP
func (client Client) LockEth(eth *big.Float) (*types.Transaction, error) {
	auth, err := client.GetAuth()
	if err != nil {
		return nil, errors.Wrap(err, "could not get auth for transaction LockEth")
	}
	auth.Value = ethutils.ToWei(eth, 18)

	a, err := abi.JSON(strings.NewReader(portal.PortalABI))
	if err != nil {
		return nil, errors.Wrapf(err, "could not read portal ABI")
	}

	cdpIDb, err := ethutils.SliceToBytes32(big.NewInt(client.cdpID).Bytes())
	if err != nil {
		return nil, errors.Wrapf(err, "could not get [32]byte")
	}

	b, err := a.Pack("lock", makerAddress, cdpIDb)
	if err != nil {
		return nil, errors.Wrapf(err, "could not pack data")
	}
	tx, err := client.dsproxy.Execute(auth, portalAddress, b)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrapf(err, "could not lock %.5f ETH to CDP #%d", eth, client.cdpID)
	}
	return tx, nil
}

// GetEthBalance returns the ETH balance of the account
func (client *Client) GetEthBalance() (eth *big.Float, err error) {
	wei, err := client.Client.BalanceAt(context.Background(), client.address, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read the ETH balance of `%s`", client.address.Hex())
	}

	return ethutils.FromWei(wei, 18), nil
}

// Draw creates DAI from the CDP
func (client *Client) Draw(dai *big.Float, cdp *cdp.CDP) (*types.Transaction, error) {
	auth, err := client.GetAuth()
	if err != nil {
		return nil, errors.Wrap(err, "could not generate the Auth")
	}

	a, err := abi.JSON(strings.NewReader(portal.PortalABI))
	if err != nil {
		return nil, errors.Wrapf(err, "could not read portal ABI")
	}

	cdpIDb, err := ethutils.SliceToBytes32(big.NewInt(client.cdpID).Bytes())
	if err != nil {
		return nil, errors.Wrapf(err, "could not get [32]byte")
	}

	b, err := a.Pack("draw", makerAddress, cdpIDb, ethutils.ToWei(dai, 18))
	if err != nil {
		return nil, errors.Wrapf(err, "could not pack data")
	}
	tx, err := client.dsproxy.Execute(auth, portalAddress, b)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrapf(err, "could not draw %.5f DAI from CDP #%d", dai, client.cdpID)
	}

	return tx, nil
}

// ChangeEthToDai changes ETH to DAI via oasis direct contract
func (client *Client) ChangeEthToDai(eth *big.Float) (*types.Transaction, error) {
	auth, err := client.GetAuth()
	if err != nil {
		return nil, errors.Wrapf(err, "could not read the ETH balance of `%s`", client.address.Hex())
	}

	a, err := abi.JSON(strings.NewReader(oasis.OasisABI))
	if err != nil {
		return nil, errors.Wrapf(err, "could not read oasis ABI")
	}
	ethPrice, err := client.GetEthPrice()
	if err != nil {
		return nil, errors.Wrapf(err, "could not read ETH price")
	}
	expectedDai := big.NewFloat(0).Mul(eth, ethPrice)
	minDai := big.NewFloat(0).Mul(expectedDai, big.NewFloat(0.95))

	b, err := a.Pack("sellAllAmountPayEth", otcAddress, wethAddress, daiAddress, ethutils.ToWei(minDai, 18))
	if err != nil {
		return nil, errors.Wrapf(err, "could not pack data")
	}

	auth.Value = ethutils.ToWei(eth, 18)
	tx, err := client.dsproxy.Execute(auth, oasisDirectAddress, b)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrap(err, "could not change ETH to DAI")
	}
	return tx, nil

}

// ChangeDaiToEth changes DAI to ETH via oasis direct contract
func (client *Client) ChangeDaiToEth(dai *big.Float) (*types.Transaction, error) {
	auth, err := client.GetAuth()
	if err != nil {
		return nil, errors.Wrapf(err, "could not read the ETH balance of `%s`", client.address.Hex())
	}

	a, err := abi.JSON(strings.NewReader(oasis.OasisABI))
	if err != nil {
		return nil, errors.Wrapf(err, "could not read oasis ABI")
	}
	ethPrice, err := client.GetEthPrice()
	if err != nil {
		return nil, errors.Wrapf(err, "could not read ETH price")
	}
	expectedEth := big.NewFloat(0).Quo(dai, ethPrice)
	minEth := big.NewFloat(0).Mul(expectedEth, big.NewFloat(0.95))

	b, err := a.Pack("sellAllAmountBuyEth", otcAddress, daiAddress, ethutils.ToWei(dai, 18), wethAddress, ethutils.ToWei(minEth, 18))
	if err != nil {
		return nil, errors.Wrapf(err, "could not pack data")
	}

	tx, err := client.dsproxy.Execute(auth, oasisDirectAddress, b)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrap(err, "could not change DAI to ETH")
	}
	return tx, nil

}

// Wipe refund DAI to the CDP
func (client *Client) Wipe(dai *big.Float, cdp *cdp.CDP) (*types.Transaction, error) {
	auth, err := client.GetAuth()
	if err != nil {
		return nil, errors.Wrapf(err, "could not read the ETH balance of `%s`", client.address.Hex())
	}

	tx, err := client.maker.Wipe(auth, cdp.BytesID, ethutils.ToWei(dai, 18))
	if err != nil {
		return nil, errors.Wrap(err, "could not wipe DAI")
	}

	return tx, nil
}

// FreeEth unlock some ETH from the CDP
func (client *Client) FreeEth(eth *big.Float, cdp *cdp.CDP) (*types.Transaction, error) {
	auth, err := client.GetAuth()
	if err != nil {
		return nil, errors.Wrap(err, "could not generate the Auth")
	}

	a, err := abi.JSON(strings.NewReader(portal.PortalABI))
	if err != nil {
		return nil, errors.Wrapf(err, "could not read portal ABI")
	}

	cdpIDb, err := ethutils.SliceToBytes32(big.NewInt(client.cdpID).Bytes())
	if err != nil {
		return nil, errors.Wrapf(err, "could not get [32]byte")
	}

	b, err := a.Pack("free", makerAddress, cdpIDb, ethutils.ToWei(eth, 18))
	if err != nil {
		return nil, errors.Wrapf(err, "could not pack data")
	}
	tx, err := client.dsproxy.Execute(auth, portalAddress, b)
	if err != nil {
		log.Fatal(err)
		return nil, errors.Wrapf(err, "could not lock %.5f ETH to CDP #%d", eth, client.cdpID)
	}

	return tx, nil
}

// GetPethBalance returns the PETH balance of the account
func (client *Client) GetPethBalance() (peth *big.Float, err error) {
	pethWei, err := client.peth.BalanceOf(&bind.CallOpts{}, client.address)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get PETH balance of `%s`", client.address)
	}

	peth = ethutils.FromWei(pethWei, 18)
	return peth, nil
}

// GetDaiBalance returns the DAI balance of the account
func (client *Client) GetDaiBalance() (dai *big.Float, err error) {
	daiWei, err := client.dai.BalanceOf(&bind.CallOpts{}, client.address)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get DAI balance of `%s`", client.address)
	}

	dai = ethutils.FromWei(daiWei, 18)
	return dai, nil
}

// ApproveMkr approves that the station contract will play with our PETH
func (client *Client) ApproveMkr() (*types.Transaction, error) {
	approved, err := client.IsMkrAllowed()
	if err != nil {
		return nil, errors.Wrapf(err, "could not get Mkr approval status")
	}
	if !approved {
		auth, err := client.GetAuth()
		if err != nil {
			return nil, errors.Wrapf(err, "could not get Auth")
		}

		tx, err := client.peth.Approve(auth, client.dsproxyAddress, ethutils.ToWei(big.NewFloat(1000000000), 18))
		if err != nil {
			return nil, errors.Wrapf(err, "could not approve `%s` to use MKR", client.dsproxyAddress.Hex())
		}

		return tx, nil
	}
	return nil, nil
}

// ApprovePeth approves that the station contract will play with our PETH
func (client *Client) ApprovePeth() (*types.Transaction, error) {
	approved, err := client.IsPethAllowed()
	if err != nil {
		return nil, errors.Wrapf(err, "could not get PETH approval status")
	}
	if !approved {
		auth, err := client.GetAuth()
		if err != nil {
			return nil, errors.Wrapf(err, "could not get Auth")
		}

		tx, err := client.peth.Approve(auth, stationAddress, ethutils.ToWei(big.NewFloat(1000000000), 18))
		if err != nil {
			return nil, errors.Wrapf(err, "could not approve `%s` to use PETH", stationAddress.Hex())
		}

		return tx, nil
	}
	return nil, nil
}

// GetReceipt returns the transaction receipt defined by hash. return nil if not yet mined
func (client *Client) GetReceipt(hash string) (*types.Receipt, error) {
	tx, pending, err := client.Client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return nil, errors.Wrapf(err, "could not retrieve transaction `%s`", hash)
	}
	if pending {
		return nil, nil
	}
	receipt, err := client.Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, errors.Wrapf(err, "could not retrieve receipt of transaction `%s`", hash)
	}
	return receipt, nil
}

// DoFreeEth is the high level function to free ETH from the CDP
func (client *Client) DoFreeEth(cdp *cdp.CDP, ethPrice, target *big.Float) (*transaction.Tx, error) {
	eth := cdp.GetEthToFree(ethPrice, target)
	if eth == nil {
		return nil, errors.New("Cannot free any ETH, there is nothing in the CDP")
	}
	tx, err := client.FreeSellWipe(eth)
	if err != nil {
		return nil, errors.Wrapf(err, "error while trying to free %.5f ETH", eth)
	}
	log.Printf("Starting free %.5f ETH, tx: %s", eth, tx.Hash().Hex())
	return &transaction.Tx{Hash: tx.Hash().Hex(), Kind: transaction.FreeSellWipe}, nil
}

// DoDrawDai is the high level function to draw DAI from the CDP
func (client *Client) DoDrawDai(cdp *cdp.CDP, ethPrice, pethRatio, target *big.Float) (*transaction.Tx, error) {
	dai := cdp.GetDaiToDraw(ethPrice, pethRatio, target)
	tx, err := client.DrawSellLock(dai)
	if err != nil {
		return nil, errors.Wrapf(err, "error while trying to draw %.2f DAI", dai)
	}
	log.Printf("Starting drawing %.2f DAI, tx: %s", dai, tx.Hash().Hex())
	return &transaction.Tx{Hash: tx.Hash().Hex(), Kind: transaction.DrawSellLock}, nil
}
