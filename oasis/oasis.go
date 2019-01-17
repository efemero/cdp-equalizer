// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oasis

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OasisABI is the input ABI used to generate the binding from.
const OasisABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"otc\",\"type\":\"address\"},{\"name\":\"payToken\",\"type\":\"address\"},{\"name\":\"payAmt\",\"type\":\"uint256\"},{\"name\":\"wethToken\",\"type\":\"address\"},{\"name\":\"minBuyAmt\",\"type\":\"uint256\"}],\"name\":\"sellAllAmountBuyEth\",\"outputs\":[{\"name\":\"wethAmt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"otc\",\"type\":\"address\"},{\"name\":\"payToken\",\"type\":\"address\"},{\"name\":\"payAmt\",\"type\":\"uint256\"},{\"name\":\"buyToken\",\"type\":\"address\"},{\"name\":\"minBuyAmt\",\"type\":\"uint256\"}],\"name\":\"sellAllAmount\",\"outputs\":[{\"name\":\"buyAmt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"otc\",\"type\":\"address\"},{\"name\":\"buyToken\",\"type\":\"address\"},{\"name\":\"buyAmt\",\"type\":\"uint256\"},{\"name\":\"payToken\",\"type\":\"address\"},{\"name\":\"maxPayAmt\",\"type\":\"uint256\"}],\"name\":\"buyAllAmount\",\"outputs\":[{\"name\":\"payAmt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"otc\",\"type\":\"address\"},{\"name\":\"buyToken\",\"type\":\"address\"},{\"name\":\"buyAmt\",\"type\":\"uint256\"},{\"name\":\"wethToken\",\"type\":\"address\"}],\"name\":\"buyAllAmountPayEth\",\"outputs\":[{\"name\":\"wethAmt\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"otc\",\"type\":\"address\"},{\"name\":\"wethToken\",\"type\":\"address\"},{\"name\":\"buyToken\",\"type\":\"address\"},{\"name\":\"minBuyAmt\",\"type\":\"uint256\"}],\"name\":\"sellAllAmountPayEth\",\"outputs\":[{\"name\":\"buyAmt\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"otc\",\"type\":\"address\"},{\"name\":\"wethToken\",\"type\":\"address\"},{\"name\":\"wethAmt\",\"type\":\"uint256\"},{\"name\":\"payToken\",\"type\":\"address\"},{\"name\":\"maxPayAmt\",\"type\":\"uint256\"}],\"name\":\"buyAllAmountBuyEth\",\"outputs\":[{\"name\":\"payAmt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// Oasis is an auto generated Go binding around an Ethereum contract.
type Oasis struct {
	OasisCaller     // Read-only binding to the contract
	OasisTransactor // Write-only binding to the contract
	OasisFilterer   // Log filterer for contract events
}

// OasisCaller is an auto generated read-only Go binding around an Ethereum contract.
type OasisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OasisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OasisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OasisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OasisSession struct {
	Contract     *Oasis            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OasisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OasisCallerSession struct {
	Contract *OasisCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OasisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OasisTransactorSession struct {
	Contract     *OasisTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OasisRaw is an auto generated low-level Go binding around an Ethereum contract.
type OasisRaw struct {
	Contract *Oasis // Generic contract binding to access the raw methods on
}

// OasisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OasisCallerRaw struct {
	Contract *OasisCaller // Generic read-only contract binding to access the raw methods on
}

// OasisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OasisTransactorRaw struct {
	Contract *OasisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOasis creates a new instance of Oasis, bound to a specific deployed contract.
func NewOasis(address common.Address, backend bind.ContractBackend) (*Oasis, error) {
	contract, err := bindOasis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oasis{OasisCaller: OasisCaller{contract: contract}, OasisTransactor: OasisTransactor{contract: contract}, OasisFilterer: OasisFilterer{contract: contract}}, nil
}

// NewOasisCaller creates a new read-only instance of Oasis, bound to a specific deployed contract.
func NewOasisCaller(address common.Address, caller bind.ContractCaller) (*OasisCaller, error) {
	contract, err := bindOasis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OasisCaller{contract: contract}, nil
}

// NewOasisTransactor creates a new write-only instance of Oasis, bound to a specific deployed contract.
func NewOasisTransactor(address common.Address, transactor bind.ContractTransactor) (*OasisTransactor, error) {
	contract, err := bindOasis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OasisTransactor{contract: contract}, nil
}

// NewOasisFilterer creates a new log filterer instance of Oasis, bound to a specific deployed contract.
func NewOasisFilterer(address common.Address, filterer bind.ContractFilterer) (*OasisFilterer, error) {
	contract, err := bindOasis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OasisFilterer{contract: contract}, nil
}

// bindOasis binds a generic wrapper to an already deployed contract.
func bindOasis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OasisABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oasis *OasisRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oasis.Contract.OasisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oasis *OasisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oasis.Contract.OasisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oasis *OasisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oasis.Contract.OasisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oasis *OasisCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oasis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oasis *OasisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oasis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oasis *OasisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oasis.Contract.contract.Transact(opts, method, params...)
}

// BuyAllAmount is a paid mutator transaction binding the contract method 0x3c0ab650.
//
// Solidity: function buyAllAmount(otc address, buyToken address, buyAmt uint256, payToken address, maxPayAmt uint256) returns(payAmt uint256)
func (_Oasis *OasisTransactor) BuyAllAmount(opts *bind.TransactOpts, otc common.Address, buyToken common.Address, buyAmt *big.Int, payToken common.Address, maxPayAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "buyAllAmount", otc, buyToken, buyAmt, payToken, maxPayAmt)
}

// BuyAllAmount is a paid mutator transaction binding the contract method 0x3c0ab650.
//
// Solidity: function buyAllAmount(otc address, buyToken address, buyAmt uint256, payToken address, maxPayAmt uint256) returns(payAmt uint256)
func (_Oasis *OasisSession) BuyAllAmount(otc common.Address, buyToken common.Address, buyAmt *big.Int, payToken common.Address, maxPayAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.BuyAllAmount(&_Oasis.TransactOpts, otc, buyToken, buyAmt, payToken, maxPayAmt)
}

// BuyAllAmount is a paid mutator transaction binding the contract method 0x3c0ab650.
//
// Solidity: function buyAllAmount(otc address, buyToken address, buyAmt uint256, payToken address, maxPayAmt uint256) returns(payAmt uint256)
func (_Oasis *OasisTransactorSession) BuyAllAmount(otc common.Address, buyToken common.Address, buyAmt *big.Int, payToken common.Address, maxPayAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.BuyAllAmount(&_Oasis.TransactOpts, otc, buyToken, buyAmt, payToken, maxPayAmt)
}

// BuyAllAmountBuyEth is a paid mutator transaction binding the contract method 0xf9a87d4f.
//
// Solidity: function buyAllAmountBuyEth(otc address, wethToken address, wethAmt uint256, payToken address, maxPayAmt uint256) returns(payAmt uint256)
func (_Oasis *OasisTransactor) BuyAllAmountBuyEth(opts *bind.TransactOpts, otc common.Address, wethToken common.Address, wethAmt *big.Int, payToken common.Address, maxPayAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "buyAllAmountBuyEth", otc, wethToken, wethAmt, payToken, maxPayAmt)
}

// BuyAllAmountBuyEth is a paid mutator transaction binding the contract method 0xf9a87d4f.
//
// Solidity: function buyAllAmountBuyEth(otc address, wethToken address, wethAmt uint256, payToken address, maxPayAmt uint256) returns(payAmt uint256)
func (_Oasis *OasisSession) BuyAllAmountBuyEth(otc common.Address, wethToken common.Address, wethAmt *big.Int, payToken common.Address, maxPayAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.BuyAllAmountBuyEth(&_Oasis.TransactOpts, otc, wethToken, wethAmt, payToken, maxPayAmt)
}

// BuyAllAmountBuyEth is a paid mutator transaction binding the contract method 0xf9a87d4f.
//
// Solidity: function buyAllAmountBuyEth(otc address, wethToken address, wethAmt uint256, payToken address, maxPayAmt uint256) returns(payAmt uint256)
func (_Oasis *OasisTransactorSession) BuyAllAmountBuyEth(otc common.Address, wethToken common.Address, wethAmt *big.Int, payToken common.Address, maxPayAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.BuyAllAmountBuyEth(&_Oasis.TransactOpts, otc, wethToken, wethAmt, payToken, maxPayAmt)
}

// BuyAllAmountPayEth is a paid mutator transaction binding the contract method 0x9a22dec5.
//
// Solidity: function buyAllAmountPayEth(otc address, buyToken address, buyAmt uint256, wethToken address) returns(wethAmt uint256)
func (_Oasis *OasisTransactor) BuyAllAmountPayEth(opts *bind.TransactOpts, otc common.Address, buyToken common.Address, buyAmt *big.Int, wethToken common.Address) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "buyAllAmountPayEth", otc, buyToken, buyAmt, wethToken)
}

// BuyAllAmountPayEth is a paid mutator transaction binding the contract method 0x9a22dec5.
//
// Solidity: function buyAllAmountPayEth(otc address, buyToken address, buyAmt uint256, wethToken address) returns(wethAmt uint256)
func (_Oasis *OasisSession) BuyAllAmountPayEth(otc common.Address, buyToken common.Address, buyAmt *big.Int, wethToken common.Address) (*types.Transaction, error) {
	return _Oasis.Contract.BuyAllAmountPayEth(&_Oasis.TransactOpts, otc, buyToken, buyAmt, wethToken)
}

// BuyAllAmountPayEth is a paid mutator transaction binding the contract method 0x9a22dec5.
//
// Solidity: function buyAllAmountPayEth(otc address, buyToken address, buyAmt uint256, wethToken address) returns(wethAmt uint256)
func (_Oasis *OasisTransactorSession) BuyAllAmountPayEth(otc common.Address, buyToken common.Address, buyAmt *big.Int, wethToken common.Address) (*types.Transaction, error) {
	return _Oasis.Contract.BuyAllAmountPayEth(&_Oasis.TransactOpts, otc, buyToken, buyAmt, wethToken)
}

// SellAllAmount is a paid mutator transaction binding the contract method 0x1ebf0d11.
//
// Solidity: function sellAllAmount(otc address, payToken address, payAmt uint256, buyToken address, minBuyAmt uint256) returns(buyAmt uint256)
func (_Oasis *OasisTransactor) SellAllAmount(opts *bind.TransactOpts, otc common.Address, payToken common.Address, payAmt *big.Int, buyToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "sellAllAmount", otc, payToken, payAmt, buyToken, minBuyAmt)
}

// SellAllAmount is a paid mutator transaction binding the contract method 0x1ebf0d11.
//
// Solidity: function sellAllAmount(otc address, payToken address, payAmt uint256, buyToken address, minBuyAmt uint256) returns(buyAmt uint256)
func (_Oasis *OasisSession) SellAllAmount(otc common.Address, payToken common.Address, payAmt *big.Int, buyToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SellAllAmount(&_Oasis.TransactOpts, otc, payToken, payAmt, buyToken, minBuyAmt)
}

// SellAllAmount is a paid mutator transaction binding the contract method 0x1ebf0d11.
//
// Solidity: function sellAllAmount(otc address, payToken address, payAmt uint256, buyToken address, minBuyAmt uint256) returns(buyAmt uint256)
func (_Oasis *OasisTransactorSession) SellAllAmount(otc common.Address, payToken common.Address, payAmt *big.Int, buyToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SellAllAmount(&_Oasis.TransactOpts, otc, payToken, payAmt, buyToken, minBuyAmt)
}

// SellAllAmountBuyEth is a paid mutator transaction binding the contract method 0x03e1b3c6.
//
// Solidity: function sellAllAmountBuyEth(otc address, payToken address, payAmt uint256, wethToken address, minBuyAmt uint256) returns(wethAmt uint256)
func (_Oasis *OasisTransactor) SellAllAmountBuyEth(opts *bind.TransactOpts, otc common.Address, payToken common.Address, payAmt *big.Int, wethToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "sellAllAmountBuyEth", otc, payToken, payAmt, wethToken, minBuyAmt)
}

// SellAllAmountBuyEth is a paid mutator transaction binding the contract method 0x03e1b3c6.
//
// Solidity: function sellAllAmountBuyEth(otc address, payToken address, payAmt uint256, wethToken address, minBuyAmt uint256) returns(wethAmt uint256)
func (_Oasis *OasisSession) SellAllAmountBuyEth(otc common.Address, payToken common.Address, payAmt *big.Int, wethToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SellAllAmountBuyEth(&_Oasis.TransactOpts, otc, payToken, payAmt, wethToken, minBuyAmt)
}

// SellAllAmountBuyEth is a paid mutator transaction binding the contract method 0x03e1b3c6.
//
// Solidity: function sellAllAmountBuyEth(otc address, payToken address, payAmt uint256, wethToken address, minBuyAmt uint256) returns(wethAmt uint256)
func (_Oasis *OasisTransactorSession) SellAllAmountBuyEth(otc common.Address, payToken common.Address, payAmt *big.Int, wethToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SellAllAmountBuyEth(&_Oasis.TransactOpts, otc, payToken, payAmt, wethToken, minBuyAmt)
}

// SellAllAmountPayEth is a paid mutator transaction binding the contract method 0xe50278a6.
//
// Solidity: function sellAllAmountPayEth(otc address, wethToken address, buyToken address, minBuyAmt uint256) returns(buyAmt uint256)
func (_Oasis *OasisTransactor) SellAllAmountPayEth(opts *bind.TransactOpts, otc common.Address, wethToken common.Address, buyToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.contract.Transact(opts, "sellAllAmountPayEth", otc, wethToken, buyToken, minBuyAmt)
}

// SellAllAmountPayEth is a paid mutator transaction binding the contract method 0xe50278a6.
//
// Solidity: function sellAllAmountPayEth(otc address, wethToken address, buyToken address, minBuyAmt uint256) returns(buyAmt uint256)
func (_Oasis *OasisSession) SellAllAmountPayEth(otc common.Address, wethToken common.Address, buyToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SellAllAmountPayEth(&_Oasis.TransactOpts, otc, wethToken, buyToken, minBuyAmt)
}

// SellAllAmountPayEth is a paid mutator transaction binding the contract method 0xe50278a6.
//
// Solidity: function sellAllAmountPayEth(otc address, wethToken address, buyToken address, minBuyAmt uint256) returns(buyAmt uint256)
func (_Oasis *OasisTransactorSession) SellAllAmountPayEth(otc common.Address, wethToken common.Address, buyToken common.Address, minBuyAmt *big.Int) (*types.Transaction, error) {
	return _Oasis.Contract.SellAllAmountPayEth(&_Oasis.TransactOpts, otc, wethToken, buyToken, minBuyAmt)
}
