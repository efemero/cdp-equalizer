// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package equalizer

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

// EqualizerABI is the input ABI used to generate the binding from.
const EqualizerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tub\",\"type\":\"address\"},{\"name\":\"otc\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"sai\",\"type\":\"address\"},{\"name\":\"freeAmt\",\"type\":\"uint256\"},{\"name\":\"weth\",\"type\":\"address\"},{\"name\":\"minWipeAmt\",\"type\":\"uint256\"}],\"name\":\"freeSellWipe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub\",\"type\":\"address\"},{\"name\":\"otc\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"sai\",\"type\":\"address\"},{\"name\":\"drawAmt\",\"type\":\"uint256\"},{\"name\":\"weth\",\"type\":\"address\"},{\"name\":\"minLockAmt\",\"type\":\"uint256\"}],\"name\":\"drawSellLock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Equalizer is an auto generated Go binding around an Ethereum contract.
type Equalizer struct {
	EqualizerCaller     // Read-only binding to the contract
	EqualizerTransactor // Write-only binding to the contract
	EqualizerFilterer   // Log filterer for contract events
}

// EqualizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EqualizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EqualizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EqualizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EqualizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EqualizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EqualizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EqualizerSession struct {
	Contract     *Equalizer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EqualizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EqualizerCallerSession struct {
	Contract *EqualizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EqualizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EqualizerTransactorSession struct {
	Contract     *EqualizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EqualizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EqualizerRaw struct {
	Contract *Equalizer // Generic contract binding to access the raw methods on
}

// EqualizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EqualizerCallerRaw struct {
	Contract *EqualizerCaller // Generic read-only contract binding to access the raw methods on
}

// EqualizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EqualizerTransactorRaw struct {
	Contract *EqualizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEqualizer creates a new instance of Equalizer, bound to a specific deployed contract.
func NewEqualizer(address common.Address, backend bind.ContractBackend) (*Equalizer, error) {
	contract, err := bindEqualizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Equalizer{EqualizerCaller: EqualizerCaller{contract: contract}, EqualizerTransactor: EqualizerTransactor{contract: contract}, EqualizerFilterer: EqualizerFilterer{contract: contract}}, nil
}

// NewEqualizerCaller creates a new read-only instance of Equalizer, bound to a specific deployed contract.
func NewEqualizerCaller(address common.Address, caller bind.ContractCaller) (*EqualizerCaller, error) {
	contract, err := bindEqualizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EqualizerCaller{contract: contract}, nil
}

// NewEqualizerTransactor creates a new write-only instance of Equalizer, bound to a specific deployed contract.
func NewEqualizerTransactor(address common.Address, transactor bind.ContractTransactor) (*EqualizerTransactor, error) {
	contract, err := bindEqualizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EqualizerTransactor{contract: contract}, nil
}

// NewEqualizerFilterer creates a new log filterer instance of Equalizer, bound to a specific deployed contract.
func NewEqualizerFilterer(address common.Address, filterer bind.ContractFilterer) (*EqualizerFilterer, error) {
	contract, err := bindEqualizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EqualizerFilterer{contract: contract}, nil
}

// bindEqualizer binds a generic wrapper to an already deployed contract.
func bindEqualizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EqualizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Equalizer *EqualizerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Equalizer.Contract.EqualizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Equalizer *EqualizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Equalizer.Contract.EqualizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Equalizer *EqualizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Equalizer.Contract.EqualizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Equalizer *EqualizerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Equalizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Equalizer *EqualizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Equalizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Equalizer *EqualizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Equalizer.Contract.contract.Transact(opts, method, params...)
}

// DrawSellLock is a paid mutator transaction binding the contract method 0xf7020cda.
//
// Solidity: function drawSellLock(tub address, otc address, cup bytes32, sai address, drawAmt uint256, weth address, minLockAmt uint256) returns()
func (_Equalizer *EqualizerTransactor) DrawSellLock(opts *bind.TransactOpts, tub common.Address, otc common.Address, cup [32]byte, sai common.Address, drawAmt *big.Int, weth common.Address, minLockAmt *big.Int) (*types.Transaction, error) {
	return _Equalizer.contract.Transact(opts, "drawSellLock", tub, otc, cup, sai, drawAmt, weth, minLockAmt)
}

// DrawSellLock is a paid mutator transaction binding the contract method 0xf7020cda.
//
// Solidity: function drawSellLock(tub address, otc address, cup bytes32, sai address, drawAmt uint256, weth address, minLockAmt uint256) returns()
func (_Equalizer *EqualizerSession) DrawSellLock(tub common.Address, otc common.Address, cup [32]byte, sai common.Address, drawAmt *big.Int, weth common.Address, minLockAmt *big.Int) (*types.Transaction, error) {
	return _Equalizer.Contract.DrawSellLock(&_Equalizer.TransactOpts, tub, otc, cup, sai, drawAmt, weth, minLockAmt)
}

// DrawSellLock is a paid mutator transaction binding the contract method 0xf7020cda.
//
// Solidity: function drawSellLock(tub address, otc address, cup bytes32, sai address, drawAmt uint256, weth address, minLockAmt uint256) returns()
func (_Equalizer *EqualizerTransactorSession) DrawSellLock(tub common.Address, otc common.Address, cup [32]byte, sai common.Address, drawAmt *big.Int, weth common.Address, minLockAmt *big.Int) (*types.Transaction, error) {
	return _Equalizer.Contract.DrawSellLock(&_Equalizer.TransactOpts, tub, otc, cup, sai, drawAmt, weth, minLockAmt)
}

// FreeSellWipe is a paid mutator transaction binding the contract method 0xe466979f.
//
// Solidity: function freeSellWipe(tub address, otc address, cup bytes32, sai address, freeAmt uint256, weth address, minWipeAmt uint256) returns()
func (_Equalizer *EqualizerTransactor) FreeSellWipe(opts *bind.TransactOpts, tub common.Address, otc common.Address, cup [32]byte, sai common.Address, freeAmt *big.Int, weth common.Address, minWipeAmt *big.Int) (*types.Transaction, error) {
	return _Equalizer.contract.Transact(opts, "freeSellWipe", tub, otc, cup, sai, freeAmt, weth, minWipeAmt)
}

// FreeSellWipe is a paid mutator transaction binding the contract method 0xe466979f.
//
// Solidity: function freeSellWipe(tub address, otc address, cup bytes32, sai address, freeAmt uint256, weth address, minWipeAmt uint256) returns()
func (_Equalizer *EqualizerSession) FreeSellWipe(tub common.Address, otc common.Address, cup [32]byte, sai common.Address, freeAmt *big.Int, weth common.Address, minWipeAmt *big.Int) (*types.Transaction, error) {
	return _Equalizer.Contract.FreeSellWipe(&_Equalizer.TransactOpts, tub, otc, cup, sai, freeAmt, weth, minWipeAmt)
}

// FreeSellWipe is a paid mutator transaction binding the contract method 0xe466979f.
//
// Solidity: function freeSellWipe(tub address, otc address, cup bytes32, sai address, freeAmt uint256, weth address, minWipeAmt uint256) returns()
func (_Equalizer *EqualizerTransactorSession) FreeSellWipe(tub common.Address, otc common.Address, cup [32]byte, sai common.Address, freeAmt *big.Int, weth common.Address, minWipeAmt *big.Int) (*types.Transaction, error) {
	return _Equalizer.Contract.FreeSellWipe(&_Equalizer.TransactOpts, tub, otc, cup, sai, freeAmt, weth, minWipeAmt)
}
