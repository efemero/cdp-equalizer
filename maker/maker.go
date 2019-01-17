// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maker

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

// MakerABI is the input ABI used to generate the binding from.
const MakerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"skr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"era\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"ink\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rho\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"air\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rhi\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"flow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"bite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"draw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cupi\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"axe\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tag\",\"outputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"off\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vox\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"rap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"wipe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tap_\",\"type\":\"address\"}],\"name\":\"turn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"per\",\"outputs\":[{\"name\":\"ray\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pip_\",\"type\":\"address\"}],\"name\":\"setPip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pie\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fit_\",\"type\":\"uint256\"},{\"name\":\"jam\",\"type\":\"uint256\"}],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sai\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"param\",\"type\":\"bytes32\"},{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"mold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tax\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mat\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pep\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"out\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"shut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"give\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"chi\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vox_\",\"type\":\"address\"}],\"name\":\"setVox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pip\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pep_\",\"type\":\"address\"}],\"name\":\"setPep\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"lad\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"din\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"ask\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"safe\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"tab\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"open\",\"outputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cups\",\"outputs\":[{\"name\":\"lad\",\"type\":\"address\"},{\"name\":\"ink\",\"type\":\"uint256\"},{\"name\":\"art\",\"type\":\"uint256\"},{\"name\":\"ire\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"sai_\",\"type\":\"address\"},{\"name\":\"sin_\",\"type\":\"address\"},{\"name\":\"skr_\",\"type\":\"address\"},{\"name\":\"gem_\",\"type\":\"address\"},{\"name\":\"gov_\",\"type\":\"address\"},{\"name\":\"pip_\",\"type\":\"address\"},{\"name\":\"pep_\",\"type\":\"address\"},{\"name\":\"vox_\",\"type\":\"address\"},{\"name\":\"pit_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"lad\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"LogNewCup\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// Maker is an auto generated Go binding around an Ethereum contract.
type Maker struct {
	MakerCaller     // Read-only binding to the contract
	MakerTransactor // Write-only binding to the contract
	MakerFilterer   // Log filterer for contract events
}

// MakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MakerSession struct {
	Contract     *Maker            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MakerCallerSession struct {
	Contract *MakerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MakerTransactorSession struct {
	Contract     *MakerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MakerRaw struct {
	Contract *Maker // Generic contract binding to access the raw methods on
}

// MakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MakerCallerRaw struct {
	Contract *MakerCaller // Generic read-only contract binding to access the raw methods on
}

// MakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MakerTransactorRaw struct {
	Contract *MakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMaker creates a new instance of Maker, bound to a specific deployed contract.
func NewMaker(address common.Address, backend bind.ContractBackend) (*Maker, error) {
	contract, err := bindMaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Maker{MakerCaller: MakerCaller{contract: contract}, MakerTransactor: MakerTransactor{contract: contract}, MakerFilterer: MakerFilterer{contract: contract}}, nil
}

// NewMakerCaller creates a new read-only instance of Maker, bound to a specific deployed contract.
func NewMakerCaller(address common.Address, caller bind.ContractCaller) (*MakerCaller, error) {
	contract, err := bindMaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MakerCaller{contract: contract}, nil
}

// NewMakerTransactor creates a new write-only instance of Maker, bound to a specific deployed contract.
func NewMakerTransactor(address common.Address, transactor bind.ContractTransactor) (*MakerTransactor, error) {
	contract, err := bindMaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MakerTransactor{contract: contract}, nil
}

// NewMakerFilterer creates a new log filterer instance of Maker, bound to a specific deployed contract.
func NewMakerFilterer(address common.Address, filterer bind.ContractFilterer) (*MakerFilterer, error) {
	contract, err := bindMaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MakerFilterer{contract: contract}, nil
}

// bindMaker binds a generic wrapper to an already deployed contract.
func bindMaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MakerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Maker *MakerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Maker.Contract.MakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Maker *MakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maker.Contract.MakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Maker *MakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Maker.Contract.MakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Maker *MakerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Maker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Maker *MakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Maker *MakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Maker.Contract.contract.Transact(opts, method, params...)
}

// Air is a free data retrieval call binding the contract method 0x27e7e21e.
//
// Solidity: function air() constant returns(uint256)
func (_Maker *MakerCaller) Air(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "air")
	return *ret0, err
}

// Air is a free data retrieval call binding the contract method 0x27e7e21e.
//
// Solidity: function air() constant returns(uint256)
func (_Maker *MakerSession) Air() (*big.Int, error) {
	return _Maker.Contract.Air(&_Maker.CallOpts)
}

// Air is a free data retrieval call binding the contract method 0x27e7e21e.
//
// Solidity: function air() constant returns(uint256)
func (_Maker *MakerCallerSession) Air() (*big.Int, error) {
	return _Maker.Contract.Air(&_Maker.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0xe47e7e66.
//
// Solidity: function ask(wad uint256) constant returns(uint256)
func (_Maker *MakerCaller) Ask(opts *bind.CallOpts, wad *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "ask", wad)
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0xe47e7e66.
//
// Solidity: function ask(wad uint256) constant returns(uint256)
func (_Maker *MakerSession) Ask(wad *big.Int) (*big.Int, error) {
	return _Maker.Contract.Ask(&_Maker.CallOpts, wad)
}

// Ask is a free data retrieval call binding the contract method 0xe47e7e66.
//
// Solidity: function ask(wad uint256) constant returns(uint256)
func (_Maker *MakerCallerSession) Ask(wad *big.Int) (*big.Int, error) {
	return _Maker.Contract.Ask(&_Maker.CallOpts, wad)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Maker *MakerCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Maker *MakerSession) Authority() (common.Address, error) {
	return _Maker.Contract.Authority(&_Maker.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Maker *MakerCallerSession) Authority() (common.Address, error) {
	return _Maker.Contract.Authority(&_Maker.CallOpts)
}

// Axe is a free data retrieval call binding the contract method 0x509bf2bf.
//
// Solidity: function axe() constant returns(uint256)
func (_Maker *MakerCaller) Axe(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "axe")
	return *ret0, err
}

// Axe is a free data retrieval call binding the contract method 0x509bf2bf.
//
// Solidity: function axe() constant returns(uint256)
func (_Maker *MakerSession) Axe() (*big.Int, error) {
	return _Maker.Contract.Axe(&_Maker.CallOpts)
}

// Axe is a free data retrieval call binding the contract method 0x509bf2bf.
//
// Solidity: function axe() constant returns(uint256)
func (_Maker *MakerCallerSession) Axe() (*big.Int, error) {
	return _Maker.Contract.Axe(&_Maker.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x454a2ab3.
//
// Solidity: function bid(wad uint256) constant returns(uint256)
func (_Maker *MakerCaller) Bid(opts *bind.CallOpts, wad *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "bid", wad)
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x454a2ab3.
//
// Solidity: function bid(wad uint256) constant returns(uint256)
func (_Maker *MakerSession) Bid(wad *big.Int) (*big.Int, error) {
	return _Maker.Contract.Bid(&_Maker.CallOpts, wad)
}

// Bid is a free data retrieval call binding the contract method 0x454a2ab3.
//
// Solidity: function bid(wad uint256) constant returns(uint256)
func (_Maker *MakerCallerSession) Bid(wad *big.Int) (*big.Int, error) {
	return _Maker.Contract.Bid(&_Maker.CallOpts, wad)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_Maker *MakerCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_Maker *MakerSession) Cap() (*big.Int, error) {
	return _Maker.Contract.Cap(&_Maker.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_Maker *MakerCallerSession) Cap() (*big.Int, error) {
	return _Maker.Contract.Cap(&_Maker.CallOpts)
}

// Cupi is a free data retrieval call binding the contract method 0x49955431.
//
// Solidity: function cupi() constant returns(uint256)
func (_Maker *MakerCaller) Cupi(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "cupi")
	return *ret0, err
}

// Cupi is a free data retrieval call binding the contract method 0x49955431.
//
// Solidity: function cupi() constant returns(uint256)
func (_Maker *MakerSession) Cupi() (*big.Int, error) {
	return _Maker.Contract.Cupi(&_Maker.CallOpts)
}

// Cupi is a free data retrieval call binding the contract method 0x49955431.
//
// Solidity: function cupi() constant returns(uint256)
func (_Maker *MakerCallerSession) Cupi() (*big.Int, error) {
	return _Maker.Contract.Cupi(&_Maker.CallOpts)
}

// Cups is a free data retrieval call binding the contract method 0xfdac0025.
//
// Solidity: function cups( bytes32) constant returns(lad address, ink uint256, art uint256, ire uint256)
func (_Maker *MakerCaller) Cups(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Lad common.Address
	Ink *big.Int
	Art *big.Int
	Ire *big.Int
}, error) {
	ret := new(struct {
		Lad common.Address
		Ink *big.Int
		Art *big.Int
		Ire *big.Int
	})
	out := ret
	err := _Maker.contract.Call(opts, out, "cups", arg0)
	return *ret, err
}

// Cups is a free data retrieval call binding the contract method 0xfdac0025.
//
// Solidity: function cups( bytes32) constant returns(lad address, ink uint256, art uint256, ire uint256)
func (_Maker *MakerSession) Cups(arg0 [32]byte) (struct {
	Lad common.Address
	Ink *big.Int
	Art *big.Int
	Ire *big.Int
}, error) {
	return _Maker.Contract.Cups(&_Maker.CallOpts, arg0)
}

// Cups is a free data retrieval call binding the contract method 0xfdac0025.
//
// Solidity: function cups( bytes32) constant returns(lad address, ink uint256, art uint256, ire uint256)
func (_Maker *MakerCallerSession) Cups(arg0 [32]byte) (struct {
	Lad common.Address
	Ink *big.Int
	Art *big.Int
	Ire *big.Int
}, error) {
	return _Maker.Contract.Cups(&_Maker.CallOpts, arg0)
}

// Era is a free data retrieval call binding the contract method 0x143e55e0.
//
// Solidity: function era() constant returns(uint256)
func (_Maker *MakerCaller) Era(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "era")
	return *ret0, err
}

// Era is a free data retrieval call binding the contract method 0x143e55e0.
//
// Solidity: function era() constant returns(uint256)
func (_Maker *MakerSession) Era() (*big.Int, error) {
	return _Maker.Contract.Era(&_Maker.CallOpts)
}

// Era is a free data retrieval call binding the contract method 0x143e55e0.
//
// Solidity: function era() constant returns(uint256)
func (_Maker *MakerCallerSession) Era() (*big.Int, error) {
	return _Maker.Contract.Era(&_Maker.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Maker *MakerCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Maker *MakerSession) Fee() (*big.Int, error) {
	return _Maker.Contract.Fee(&_Maker.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Maker *MakerCallerSession) Fee() (*big.Int, error) {
	return _Maker.Contract.Fee(&_Maker.CallOpts)
}

// Fit is a free data retrieval call binding the contract method 0xc8e13bb4.
//
// Solidity: function fit() constant returns(uint256)
func (_Maker *MakerCaller) Fit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "fit")
	return *ret0, err
}

// Fit is a free data retrieval call binding the contract method 0xc8e13bb4.
//
// Solidity: function fit() constant returns(uint256)
func (_Maker *MakerSession) Fit() (*big.Int, error) {
	return _Maker.Contract.Fit(&_Maker.CallOpts)
}

// Fit is a free data retrieval call binding the contract method 0xc8e13bb4.
//
// Solidity: function fit() constant returns(uint256)
func (_Maker *MakerCallerSession) Fit() (*big.Int, error) {
	return _Maker.Contract.Fit(&_Maker.CallOpts)
}

// Gap is a free data retrieval call binding the contract method 0x6c32c0a6.
//
// Solidity: function gap() constant returns(uint256)
func (_Maker *MakerCaller) Gap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "gap")
	return *ret0, err
}

// Gap is a free data retrieval call binding the contract method 0x6c32c0a6.
//
// Solidity: function gap() constant returns(uint256)
func (_Maker *MakerSession) Gap() (*big.Int, error) {
	return _Maker.Contract.Gap(&_Maker.CallOpts)
}

// Gap is a free data retrieval call binding the contract method 0x6c32c0a6.
//
// Solidity: function gap() constant returns(uint256)
func (_Maker *MakerCallerSession) Gap() (*big.Int, error) {
	return _Maker.Contract.Gap(&_Maker.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Maker *MakerCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "gem")
	return *ret0, err
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Maker *MakerSession) Gem() (common.Address, error) {
	return _Maker.Contract.Gem(&_Maker.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() constant returns(address)
func (_Maker *MakerCallerSession) Gem() (common.Address, error) {
	return _Maker.Contract.Gem(&_Maker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() constant returns(address)
func (_Maker *MakerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "gov")
	return *ret0, err
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() constant returns(address)
func (_Maker *MakerSession) Gov() (common.Address, error) {
	return _Maker.Contract.Gov(&_Maker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() constant returns(address)
func (_Maker *MakerCallerSession) Gov() (common.Address, error) {
	return _Maker.Contract.Gov(&_Maker.CallOpts)
}

// Ink is a free data retrieval call binding the contract method 0x1f3634ed.
//
// Solidity: function ink(cup bytes32) constant returns(uint256)
func (_Maker *MakerCaller) Ink(opts *bind.CallOpts, cup [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "ink", cup)
	return *ret0, err
}

// Ink is a free data retrieval call binding the contract method 0x1f3634ed.
//
// Solidity: function ink(cup bytes32) constant returns(uint256)
func (_Maker *MakerSession) Ink(cup [32]byte) (*big.Int, error) {
	return _Maker.Contract.Ink(&_Maker.CallOpts, cup)
}

// Ink is a free data retrieval call binding the contract method 0x1f3634ed.
//
// Solidity: function ink(cup bytes32) constant returns(uint256)
func (_Maker *MakerCallerSession) Ink(cup [32]byte) (*big.Int, error) {
	return _Maker.Contract.Ink(&_Maker.CallOpts, cup)
}

// Lad is a free data retrieval call binding the contract method 0xde5f5517.
//
// Solidity: function lad(cup bytes32) constant returns(address)
func (_Maker *MakerCaller) Lad(opts *bind.CallOpts, cup [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "lad", cup)
	return *ret0, err
}

// Lad is a free data retrieval call binding the contract method 0xde5f5517.
//
// Solidity: function lad(cup bytes32) constant returns(address)
func (_Maker *MakerSession) Lad(cup [32]byte) (common.Address, error) {
	return _Maker.Contract.Lad(&_Maker.CallOpts, cup)
}

// Lad is a free data retrieval call binding the contract method 0xde5f5517.
//
// Solidity: function lad(cup bytes32) constant returns(address)
func (_Maker *MakerCallerSession) Lad(cup [32]byte) (common.Address, error) {
	return _Maker.Contract.Lad(&_Maker.CallOpts, cup)
}

// Mat is a free data retrieval call binding the contract method 0xab0783da.
//
// Solidity: function mat() constant returns(uint256)
func (_Maker *MakerCaller) Mat(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "mat")
	return *ret0, err
}

// Mat is a free data retrieval call binding the contract method 0xab0783da.
//
// Solidity: function mat() constant returns(uint256)
func (_Maker *MakerSession) Mat() (*big.Int, error) {
	return _Maker.Contract.Mat(&_Maker.CallOpts)
}

// Mat is a free data retrieval call binding the contract method 0xab0783da.
//
// Solidity: function mat() constant returns(uint256)
func (_Maker *MakerCallerSession) Mat() (*big.Int, error) {
	return _Maker.Contract.Mat(&_Maker.CallOpts)
}

// Off is a free data retrieval call binding the contract method 0x6626b26d.
//
// Solidity: function off() constant returns(bool)
func (_Maker *MakerCaller) Off(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "off")
	return *ret0, err
}

// Off is a free data retrieval call binding the contract method 0x6626b26d.
//
// Solidity: function off() constant returns(bool)
func (_Maker *MakerSession) Off() (bool, error) {
	return _Maker.Contract.Off(&_Maker.CallOpts)
}

// Off is a free data retrieval call binding the contract method 0x6626b26d.
//
// Solidity: function off() constant returns(bool)
func (_Maker *MakerCallerSession) Off() (bool, error) {
	return _Maker.Contract.Off(&_Maker.CallOpts)
}

// Out is a free data retrieval call binding the contract method 0xb2a1449b.
//
// Solidity: function out() constant returns(bool)
func (_Maker *MakerCaller) Out(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "out")
	return *ret0, err
}

// Out is a free data retrieval call binding the contract method 0xb2a1449b.
//
// Solidity: function out() constant returns(bool)
func (_Maker *MakerSession) Out() (bool, error) {
	return _Maker.Contract.Out(&_Maker.CallOpts)
}

// Out is a free data retrieval call binding the contract method 0xb2a1449b.
//
// Solidity: function out() constant returns(bool)
func (_Maker *MakerCallerSession) Out() (bool, error) {
	return _Maker.Contract.Out(&_Maker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Maker *MakerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Maker *MakerSession) Owner() (common.Address, error) {
	return _Maker.Contract.Owner(&_Maker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Maker *MakerCallerSession) Owner() (common.Address, error) {
	return _Maker.Contract.Owner(&_Maker.CallOpts)
}

// Pep is a free data retrieval call binding the contract method 0xace237f5.
//
// Solidity: function pep() constant returns(address)
func (_Maker *MakerCaller) Pep(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "pep")
	return *ret0, err
}

// Pep is a free data retrieval call binding the contract method 0xace237f5.
//
// Solidity: function pep() constant returns(address)
func (_Maker *MakerSession) Pep() (common.Address, error) {
	return _Maker.Contract.Pep(&_Maker.CallOpts)
}

// Pep is a free data retrieval call binding the contract method 0xace237f5.
//
// Solidity: function pep() constant returns(address)
func (_Maker *MakerCallerSession) Pep() (common.Address, error) {
	return _Maker.Contract.Pep(&_Maker.CallOpts)
}

// Per is a free data retrieval call binding the contract method 0x7ec9c3b8.
//
// Solidity: function per() constant returns(ray uint256)
func (_Maker *MakerCaller) Per(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "per")
	return *ret0, err
}

// Per is a free data retrieval call binding the contract method 0x7ec9c3b8.
//
// Solidity: function per() constant returns(ray uint256)
func (_Maker *MakerSession) Per() (*big.Int, error) {
	return _Maker.Contract.Per(&_Maker.CallOpts)
}

// Per is a free data retrieval call binding the contract method 0x7ec9c3b8.
//
// Solidity: function per() constant returns(ray uint256)
func (_Maker *MakerCallerSession) Per() (*big.Int, error) {
	return _Maker.Contract.Per(&_Maker.CallOpts)
}

// Pie is a free data retrieval call binding the contract method 0x8a95a746.
//
// Solidity: function pie() constant returns(uint256)
func (_Maker *MakerCaller) Pie(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "pie")
	return *ret0, err
}

// Pie is a free data retrieval call binding the contract method 0x8a95a746.
//
// Solidity: function pie() constant returns(uint256)
func (_Maker *MakerSession) Pie() (*big.Int, error) {
	return _Maker.Contract.Pie(&_Maker.CallOpts)
}

// Pie is a free data retrieval call binding the contract method 0x8a95a746.
//
// Solidity: function pie() constant returns(uint256)
func (_Maker *MakerCallerSession) Pie() (*big.Int, error) {
	return _Maker.Contract.Pie(&_Maker.CallOpts)
}

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() constant returns(address)
func (_Maker *MakerCaller) Pip(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "pip")
	return *ret0, err
}

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() constant returns(address)
func (_Maker *MakerSession) Pip() (common.Address, error) {
	return _Maker.Contract.Pip(&_Maker.CallOpts)
}

// Pip is a free data retrieval call binding the contract method 0xd741e2f9.
//
// Solidity: function pip() constant returns(address)
func (_Maker *MakerCallerSession) Pip() (common.Address, error) {
	return _Maker.Contract.Pip(&_Maker.CallOpts)
}

// Pit is a free data retrieval call binding the contract method 0xf03c7c6e.
//
// Solidity: function pit() constant returns(address)
func (_Maker *MakerCaller) Pit(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "pit")
	return *ret0, err
}

// Pit is a free data retrieval call binding the contract method 0xf03c7c6e.
//
// Solidity: function pit() constant returns(address)
func (_Maker *MakerSession) Pit() (common.Address, error) {
	return _Maker.Contract.Pit(&_Maker.CallOpts)
}

// Pit is a free data retrieval call binding the contract method 0xf03c7c6e.
//
// Solidity: function pit() constant returns(address)
func (_Maker *MakerCallerSession) Pit() (common.Address, error) {
	return _Maker.Contract.Pit(&_Maker.CallOpts)
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() constant returns(uint256)
func (_Maker *MakerCaller) Rho(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "rho")
	return *ret0, err
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() constant returns(uint256)
func (_Maker *MakerSession) Rho() (*big.Int, error) {
	return _Maker.Contract.Rho(&_Maker.CallOpts)
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() constant returns(uint256)
func (_Maker *MakerCallerSession) Rho() (*big.Int, error) {
	return _Maker.Contract.Rho(&_Maker.CallOpts)
}

// Rum is a free data retrieval call binding the contract method 0x8cf0c191.
//
// Solidity: function rum() constant returns(uint256)
func (_Maker *MakerCaller) Rum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "rum")
	return *ret0, err
}

// Rum is a free data retrieval call binding the contract method 0x8cf0c191.
//
// Solidity: function rum() constant returns(uint256)
func (_Maker *MakerSession) Rum() (*big.Int, error) {
	return _Maker.Contract.Rum(&_Maker.CallOpts)
}

// Rum is a free data retrieval call binding the contract method 0x8cf0c191.
//
// Solidity: function rum() constant returns(uint256)
func (_Maker *MakerCallerSession) Rum() (*big.Int, error) {
	return _Maker.Contract.Rum(&_Maker.CallOpts)
}

// Sai is a free data retrieval call binding the contract method 0x9166cba4.
//
// Solidity: function sai() constant returns(address)
func (_Maker *MakerCaller) Sai(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "sai")
	return *ret0, err
}

// Sai is a free data retrieval call binding the contract method 0x9166cba4.
//
// Solidity: function sai() constant returns(address)
func (_Maker *MakerSession) Sai() (common.Address, error) {
	return _Maker.Contract.Sai(&_Maker.CallOpts)
}

// Sai is a free data retrieval call binding the contract method 0x9166cba4.
//
// Solidity: function sai() constant returns(address)
func (_Maker *MakerCallerSession) Sai() (common.Address, error) {
	return _Maker.Contract.Sai(&_Maker.CallOpts)
}

// Sin is a free data retrieval call binding the contract method 0x071bafb5.
//
// Solidity: function sin() constant returns(address)
func (_Maker *MakerCaller) Sin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "sin")
	return *ret0, err
}

// Sin is a free data retrieval call binding the contract method 0x071bafb5.
//
// Solidity: function sin() constant returns(address)
func (_Maker *MakerSession) Sin() (common.Address, error) {
	return _Maker.Contract.Sin(&_Maker.CallOpts)
}

// Sin is a free data retrieval call binding the contract method 0x071bafb5.
//
// Solidity: function sin() constant returns(address)
func (_Maker *MakerCallerSession) Sin() (common.Address, error) {
	return _Maker.Contract.Sin(&_Maker.CallOpts)
}

// Skr is a free data retrieval call binding the contract method 0x0f8a771e.
//
// Solidity: function skr() constant returns(address)
func (_Maker *MakerCaller) Skr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "skr")
	return *ret0, err
}

// Skr is a free data retrieval call binding the contract method 0x0f8a771e.
//
// Solidity: function skr() constant returns(address)
func (_Maker *MakerSession) Skr() (common.Address, error) {
	return _Maker.Contract.Skr(&_Maker.CallOpts)
}

// Skr is a free data retrieval call binding the contract method 0x0f8a771e.
//
// Solidity: function skr() constant returns(address)
func (_Maker *MakerCallerSession) Skr() (common.Address, error) {
	return _Maker.Contract.Skr(&_Maker.CallOpts)
}

// Tag is a free data retrieval call binding the contract method 0x51f91066.
//
// Solidity: function tag() constant returns(wad uint256)
func (_Maker *MakerCaller) Tag(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "tag")
	return *ret0, err
}

// Tag is a free data retrieval call binding the contract method 0x51f91066.
//
// Solidity: function tag() constant returns(wad uint256)
func (_Maker *MakerSession) Tag() (*big.Int, error) {
	return _Maker.Contract.Tag(&_Maker.CallOpts)
}

// Tag is a free data retrieval call binding the contract method 0x51f91066.
//
// Solidity: function tag() constant returns(wad uint256)
func (_Maker *MakerCallerSession) Tag() (*big.Int, error) {
	return _Maker.Contract.Tag(&_Maker.CallOpts)
}

// Tap is a free data retrieval call binding the contract method 0xfd221031.
//
// Solidity: function tap() constant returns(address)
func (_Maker *MakerCaller) Tap(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "tap")
	return *ret0, err
}

// Tap is a free data retrieval call binding the contract method 0xfd221031.
//
// Solidity: function tap() constant returns(address)
func (_Maker *MakerSession) Tap() (common.Address, error) {
	return _Maker.Contract.Tap(&_Maker.CallOpts)
}

// Tap is a free data retrieval call binding the contract method 0xfd221031.
//
// Solidity: function tap() constant returns(address)
func (_Maker *MakerCallerSession) Tap() (common.Address, error) {
	return _Maker.Contract.Tap(&_Maker.CallOpts)
}

// Tax is a free data retrieval call binding the contract method 0x99c8d556.
//
// Solidity: function tax() constant returns(uint256)
func (_Maker *MakerCaller) Tax(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "tax")
	return *ret0, err
}

// Tax is a free data retrieval call binding the contract method 0x99c8d556.
//
// Solidity: function tax() constant returns(uint256)
func (_Maker *MakerSession) Tax() (*big.Int, error) {
	return _Maker.Contract.Tax(&_Maker.CallOpts)
}

// Tax is a free data retrieval call binding the contract method 0x99c8d556.
//
// Solidity: function tax() constant returns(uint256)
func (_Maker *MakerCallerSession) Tax() (*big.Int, error) {
	return _Maker.Contract.Tax(&_Maker.CallOpts)
}

// Vox is a free data retrieval call binding the contract method 0x67550a35.
//
// Solidity: function vox() constant returns(address)
func (_Maker *MakerCaller) Vox(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Maker.contract.Call(opts, out, "vox")
	return *ret0, err
}

// Vox is a free data retrieval call binding the contract method 0x67550a35.
//
// Solidity: function vox() constant returns(address)
func (_Maker *MakerSession) Vox() (common.Address, error) {
	return _Maker.Contract.Vox(&_Maker.CallOpts)
}

// Vox is a free data retrieval call binding the contract method 0x67550a35.
//
// Solidity: function vox() constant returns(address)
func (_Maker *MakerCallerSession) Vox() (common.Address, error) {
	return _Maker.Contract.Vox(&_Maker.CallOpts)
}

// Bite is a paid mutator transaction binding the contract method 0x40cc8854.
//
// Solidity: function bite(cup bytes32) returns()
func (_Maker *MakerTransactor) Bite(opts *bind.TransactOpts, cup [32]byte) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "bite", cup)
}

// Bite is a paid mutator transaction binding the contract method 0x40cc8854.
//
// Solidity: function bite(cup bytes32) returns()
func (_Maker *MakerSession) Bite(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Bite(&_Maker.TransactOpts, cup)
}

// Bite is a paid mutator transaction binding the contract method 0x40cc8854.
//
// Solidity: function bite(cup bytes32) returns()
func (_Maker *MakerTransactorSession) Bite(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Bite(&_Maker.TransactOpts, cup)
}

// Cage is a paid mutator transaction binding the contract method 0x8ceedb47.
//
// Solidity: function cage(fit_ uint256, jam uint256) returns()
func (_Maker *MakerTransactor) Cage(opts *bind.TransactOpts, fit_ *big.Int, jam *big.Int) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "cage", fit_, jam)
}

// Cage is a paid mutator transaction binding the contract method 0x8ceedb47.
//
// Solidity: function cage(fit_ uint256, jam uint256) returns()
func (_Maker *MakerSession) Cage(fit_ *big.Int, jam *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Cage(&_Maker.TransactOpts, fit_, jam)
}

// Cage is a paid mutator transaction binding the contract method 0x8ceedb47.
//
// Solidity: function cage(fit_ uint256, jam uint256) returns()
func (_Maker *MakerTransactorSession) Cage(fit_ *big.Int, jam *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Cage(&_Maker.TransactOpts, fit_, jam)
}

// Chi is a paid mutator transaction binding the contract method 0xc92aecc4.
//
// Solidity: function chi() returns(uint256)
func (_Maker *MakerTransactor) Chi(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "chi")
}

// Chi is a paid mutator transaction binding the contract method 0xc92aecc4.
//
// Solidity: function chi() returns(uint256)
func (_Maker *MakerSession) Chi() (*types.Transaction, error) {
	return _Maker.Contract.Chi(&_Maker.TransactOpts)
}

// Chi is a paid mutator transaction binding the contract method 0xc92aecc4.
//
// Solidity: function chi() returns(uint256)
func (_Maker *MakerTransactorSession) Chi() (*types.Transaction, error) {
	return _Maker.Contract.Chi(&_Maker.TransactOpts)
}

// Din is a paid mutator transaction binding the contract method 0xe0ae96e9.
//
// Solidity: function din() returns(uint256)
func (_Maker *MakerTransactor) Din(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "din")
}

// Din is a paid mutator transaction binding the contract method 0xe0ae96e9.
//
// Solidity: function din() returns(uint256)
func (_Maker *MakerSession) Din() (*types.Transaction, error) {
	return _Maker.Contract.Din(&_Maker.TransactOpts)
}

// Din is a paid mutator transaction binding the contract method 0xe0ae96e9.
//
// Solidity: function din() returns(uint256)
func (_Maker *MakerTransactorSession) Din() (*types.Transaction, error) {
	return _Maker.Contract.Din(&_Maker.TransactOpts)
}

// Draw is a paid mutator transaction binding the contract method 0x440f19ba.
//
// Solidity: function draw(cup bytes32, wad uint256) returns()
func (_Maker *MakerTransactor) Draw(opts *bind.TransactOpts, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "draw", cup, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x440f19ba.
//
// Solidity: function draw(cup bytes32, wad uint256) returns()
func (_Maker *MakerSession) Draw(cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Draw(&_Maker.TransactOpts, cup, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x440f19ba.
//
// Solidity: function draw(cup bytes32, wad uint256) returns()
func (_Maker *MakerTransactorSession) Draw(cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Draw(&_Maker.TransactOpts, cup, wad)
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns()
func (_Maker *MakerTransactor) Drip(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "drip")
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns()
func (_Maker *MakerSession) Drip() (*types.Transaction, error) {
	return _Maker.Contract.Drip(&_Maker.TransactOpts)
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns()
func (_Maker *MakerTransactorSession) Drip() (*types.Transaction, error) {
	return _Maker.Contract.Drip(&_Maker.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(wad uint256) returns()
func (_Maker *MakerTransactor) Exit(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "exit", wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(wad uint256) returns()
func (_Maker *MakerSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Exit(&_Maker.TransactOpts, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(wad uint256) returns()
func (_Maker *MakerTransactorSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Exit(&_Maker.TransactOpts, wad)
}

// Flow is a paid mutator transaction binding the contract method 0x343aad82.
//
// Solidity: function flow() returns()
func (_Maker *MakerTransactor) Flow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "flow")
}

// Flow is a paid mutator transaction binding the contract method 0x343aad82.
//
// Solidity: function flow() returns()
func (_Maker *MakerSession) Flow() (*types.Transaction, error) {
	return _Maker.Contract.Flow(&_Maker.TransactOpts)
}

// Flow is a paid mutator transaction binding the contract method 0x343aad82.
//
// Solidity: function flow() returns()
func (_Maker *MakerTransactorSession) Flow() (*types.Transaction, error) {
	return _Maker.Contract.Flow(&_Maker.TransactOpts)
}

// Free is a paid mutator transaction binding the contract method 0xa5cd184e.
//
// Solidity: function free(cup bytes32, wad uint256) returns()
func (_Maker *MakerTransactor) Free(opts *bind.TransactOpts, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "free", cup, wad)
}

// Free is a paid mutator transaction binding the contract method 0xa5cd184e.
//
// Solidity: function free(cup bytes32, wad uint256) returns()
func (_Maker *MakerSession) Free(cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Free(&_Maker.TransactOpts, cup, wad)
}

// Free is a paid mutator transaction binding the contract method 0xa5cd184e.
//
// Solidity: function free(cup bytes32, wad uint256) returns()
func (_Maker *MakerTransactorSession) Free(cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Free(&_Maker.TransactOpts, cup, wad)
}

// Give is a paid mutator transaction binding the contract method 0xbaa8529c.
//
// Solidity: function give(cup bytes32, guy address) returns()
func (_Maker *MakerTransactor) Give(opts *bind.TransactOpts, cup [32]byte, guy common.Address) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "give", cup, guy)
}

// Give is a paid mutator transaction binding the contract method 0xbaa8529c.
//
// Solidity: function give(cup bytes32, guy address) returns()
func (_Maker *MakerSession) Give(cup [32]byte, guy common.Address) (*types.Transaction, error) {
	return _Maker.Contract.Give(&_Maker.TransactOpts, cup, guy)
}

// Give is a paid mutator transaction binding the contract method 0xbaa8529c.
//
// Solidity: function give(cup bytes32, guy address) returns()
func (_Maker *MakerTransactorSession) Give(cup [32]byte, guy common.Address) (*types.Transaction, error) {
	return _Maker.Contract.Give(&_Maker.TransactOpts, cup, guy)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(wad uint256) returns()
func (_Maker *MakerTransactor) Join(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "join", wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(wad uint256) returns()
func (_Maker *MakerSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Join(&_Maker.TransactOpts, wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(wad uint256) returns()
func (_Maker *MakerTransactorSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Join(&_Maker.TransactOpts, wad)
}

// Lock is a paid mutator transaction binding the contract method 0xb3b77a51.
//
// Solidity: function lock(cup bytes32, wad uint256) returns()
func (_Maker *MakerTransactor) Lock(opts *bind.TransactOpts, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "lock", cup, wad)
}

// Lock is a paid mutator transaction binding the contract method 0xb3b77a51.
//
// Solidity: function lock(cup bytes32, wad uint256) returns()
func (_Maker *MakerSession) Lock(cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Lock(&_Maker.TransactOpts, cup, wad)
}

// Lock is a paid mutator transaction binding the contract method 0xb3b77a51.
//
// Solidity: function lock(cup bytes32, wad uint256) returns()
func (_Maker *MakerTransactorSession) Lock(cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Lock(&_Maker.TransactOpts, cup, wad)
}

// Mold is a paid mutator transaction binding the contract method 0x92b0d721.
//
// Solidity: function mold(param bytes32, val uint256) returns()
func (_Maker *MakerTransactor) Mold(opts *bind.TransactOpts, param [32]byte, val *big.Int) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "mold", param, val)
}

// Mold is a paid mutator transaction binding the contract method 0x92b0d721.
//
// Solidity: function mold(param bytes32, val uint256) returns()
func (_Maker *MakerSession) Mold(param [32]byte, val *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Mold(&_Maker.TransactOpts, param, val)
}

// Mold is a paid mutator transaction binding the contract method 0x92b0d721.
//
// Solidity: function mold(param bytes32, val uint256) returns()
func (_Maker *MakerTransactorSession) Mold(param [32]byte, val *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Mold(&_Maker.TransactOpts, param, val)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns(cup bytes32)
func (_Maker *MakerTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns(cup bytes32)
func (_Maker *MakerSession) Open() (*types.Transaction, error) {
	return _Maker.Contract.Open(&_Maker.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns(cup bytes32)
func (_Maker *MakerTransactorSession) Open() (*types.Transaction, error) {
	return _Maker.Contract.Open(&_Maker.TransactOpts)
}

// Rap is a paid mutator transaction binding the contract method 0x6f78ee0d.
//
// Solidity: function rap(cup bytes32) returns(uint256)
func (_Maker *MakerTransactor) Rap(opts *bind.TransactOpts, cup [32]byte) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "rap", cup)
}

// Rap is a paid mutator transaction binding the contract method 0x6f78ee0d.
//
// Solidity: function rap(cup bytes32) returns(uint256)
func (_Maker *MakerSession) Rap(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Rap(&_Maker.TransactOpts, cup)
}

// Rap is a paid mutator transaction binding the contract method 0x6f78ee0d.
//
// Solidity: function rap(cup bytes32) returns(uint256)
func (_Maker *MakerTransactorSession) Rap(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Rap(&_Maker.TransactOpts, cup)
}

// Rhi is a paid mutator transaction binding the contract method 0x338a0261.
//
// Solidity: function rhi() returns(uint256)
func (_Maker *MakerTransactor) Rhi(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "rhi")
}

// Rhi is a paid mutator transaction binding the contract method 0x338a0261.
//
// Solidity: function rhi() returns(uint256)
func (_Maker *MakerSession) Rhi() (*types.Transaction, error) {
	return _Maker.Contract.Rhi(&_Maker.TransactOpts)
}

// Rhi is a paid mutator transaction binding the contract method 0x338a0261.
//
// Solidity: function rhi() returns(uint256)
func (_Maker *MakerTransactorSession) Rhi() (*types.Transaction, error) {
	return _Maker.Contract.Rhi(&_Maker.TransactOpts)
}

// Safe is a paid mutator transaction binding the contract method 0xe95823ad.
//
// Solidity: function safe(cup bytes32) returns(bool)
func (_Maker *MakerTransactor) Safe(opts *bind.TransactOpts, cup [32]byte) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "safe", cup)
}

// Safe is a paid mutator transaction binding the contract method 0xe95823ad.
//
// Solidity: function safe(cup bytes32) returns(bool)
func (_Maker *MakerSession) Safe(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Safe(&_Maker.TransactOpts, cup)
}

// Safe is a paid mutator transaction binding the contract method 0xe95823ad.
//
// Solidity: function safe(cup bytes32) returns(bool)
func (_Maker *MakerTransactorSession) Safe(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Safe(&_Maker.TransactOpts, cup)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Maker *MakerTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Maker *MakerSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetAuthority(&_Maker.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Maker *MakerTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetAuthority(&_Maker.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Maker *MakerTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Maker *MakerSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetOwner(&_Maker.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Maker *MakerTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetOwner(&_Maker.TransactOpts, owner_)
}

// SetPep is a paid mutator transaction binding the contract method 0xd9c27cc6.
//
// Solidity: function setPep(pep_ address) returns()
func (_Maker *MakerTransactor) SetPep(opts *bind.TransactOpts, pep_ common.Address) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "setPep", pep_)
}

// SetPep is a paid mutator transaction binding the contract method 0xd9c27cc6.
//
// Solidity: function setPep(pep_ address) returns()
func (_Maker *MakerSession) SetPep(pep_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetPep(&_Maker.TransactOpts, pep_)
}

// SetPep is a paid mutator transaction binding the contract method 0xd9c27cc6.
//
// Solidity: function setPep(pep_ address) returns()
func (_Maker *MakerTransactorSession) SetPep(pep_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetPep(&_Maker.TransactOpts, pep_)
}

// SetPip is a paid mutator transaction binding the contract method 0x82bf9a75.
//
// Solidity: function setPip(pip_ address) returns()
func (_Maker *MakerTransactor) SetPip(opts *bind.TransactOpts, pip_ common.Address) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "setPip", pip_)
}

// SetPip is a paid mutator transaction binding the contract method 0x82bf9a75.
//
// Solidity: function setPip(pip_ address) returns()
func (_Maker *MakerSession) SetPip(pip_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetPip(&_Maker.TransactOpts, pip_)
}

// SetPip is a paid mutator transaction binding the contract method 0x82bf9a75.
//
// Solidity: function setPip(pip_ address) returns()
func (_Maker *MakerTransactorSession) SetPip(pip_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetPip(&_Maker.TransactOpts, pip_)
}

// SetVox is a paid mutator transaction binding the contract method 0xcf48d1a6.
//
// Solidity: function setVox(vox_ address) returns()
func (_Maker *MakerTransactor) SetVox(opts *bind.TransactOpts, vox_ common.Address) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "setVox", vox_)
}

// SetVox is a paid mutator transaction binding the contract method 0xcf48d1a6.
//
// Solidity: function setVox(vox_ address) returns()
func (_Maker *MakerSession) SetVox(vox_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetVox(&_Maker.TransactOpts, vox_)
}

// SetVox is a paid mutator transaction binding the contract method 0xcf48d1a6.
//
// Solidity: function setVox(vox_ address) returns()
func (_Maker *MakerTransactorSession) SetVox(vox_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.SetVox(&_Maker.TransactOpts, vox_)
}

// Shut is a paid mutator transaction binding the contract method 0xb84d2106.
//
// Solidity: function shut(cup bytes32) returns()
func (_Maker *MakerTransactor) Shut(opts *bind.TransactOpts, cup [32]byte) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "shut", cup)
}

// Shut is a paid mutator transaction binding the contract method 0xb84d2106.
//
// Solidity: function shut(cup bytes32) returns()
func (_Maker *MakerSession) Shut(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Shut(&_Maker.TransactOpts, cup)
}

// Shut is a paid mutator transaction binding the contract method 0xb84d2106.
//
// Solidity: function shut(cup bytes32) returns()
func (_Maker *MakerTransactorSession) Shut(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Shut(&_Maker.TransactOpts, cup)
}

// Tab is a paid mutator transaction binding the contract method 0xf7c8d634.
//
// Solidity: function tab(cup bytes32) returns(uint256)
func (_Maker *MakerTransactor) Tab(opts *bind.TransactOpts, cup [32]byte) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "tab", cup)
}

// Tab is a paid mutator transaction binding the contract method 0xf7c8d634.
//
// Solidity: function tab(cup bytes32) returns(uint256)
func (_Maker *MakerSession) Tab(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Tab(&_Maker.TransactOpts, cup)
}

// Tab is a paid mutator transaction binding the contract method 0xf7c8d634.
//
// Solidity: function tab(cup bytes32) returns(uint256)
func (_Maker *MakerTransactorSession) Tab(cup [32]byte) (*types.Transaction, error) {
	return _Maker.Contract.Tab(&_Maker.TransactOpts, cup)
}

// Turn is a paid mutator transaction binding the contract method 0x7e74325f.
//
// Solidity: function turn(tap_ address) returns()
func (_Maker *MakerTransactor) Turn(opts *bind.TransactOpts, tap_ common.Address) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "turn", tap_)
}

// Turn is a paid mutator transaction binding the contract method 0x7e74325f.
//
// Solidity: function turn(tap_ address) returns()
func (_Maker *MakerSession) Turn(tap_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.Turn(&_Maker.TransactOpts, tap_)
}

// Turn is a paid mutator transaction binding the contract method 0x7e74325f.
//
// Solidity: function turn(tap_ address) returns()
func (_Maker *MakerTransactorSession) Turn(tap_ common.Address) (*types.Transaction, error) {
	return _Maker.Contract.Turn(&_Maker.TransactOpts, tap_)
}

// Wipe is a paid mutator transaction binding the contract method 0x73b38101.
//
// Solidity: function wipe(cup bytes32, wad uint256) returns()
func (_Maker *MakerTransactor) Wipe(opts *bind.TransactOpts, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.contract.Transact(opts, "wipe", cup, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0x73b38101.
//
// Solidity: function wipe(cup bytes32, wad uint256) returns()
func (_Maker *MakerSession) Wipe(cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Wipe(&_Maker.TransactOpts, cup, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0x73b38101.
//
// Solidity: function wipe(cup bytes32, wad uint256) returns()
func (_Maker *MakerTransactorSession) Wipe(cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Maker.Contract.Wipe(&_Maker.TransactOpts, cup, wad)
}

// MakerLogNewCupIterator is returned from FilterLogNewCup and is used to iterate over the raw logs and unpacked data for LogNewCup events raised by the Maker contract.
type MakerLogNewCupIterator struct {
	Event *MakerLogNewCup // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MakerLogNewCupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerLogNewCup)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MakerLogNewCup)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MakerLogNewCupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerLogNewCupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerLogNewCup represents a LogNewCup event raised by the Maker contract.
type MakerLogNewCup struct {
	Lad common.Address
	Cup [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNewCup is a free log retrieval operation binding the contract event 0x89b8893b806db50897c8e2362c71571cfaeb9761ee40727f683f1793cda9df16.
//
// Solidity: e LogNewCup(lad indexed address, cup bytes32)
func (_Maker *MakerFilterer) FilterLogNewCup(opts *bind.FilterOpts, lad []common.Address) (*MakerLogNewCupIterator, error) {

	var ladRule []interface{}
	for _, ladItem := range lad {
		ladRule = append(ladRule, ladItem)
	}

	logs, sub, err := _Maker.contract.FilterLogs(opts, "LogNewCup", ladRule)
	if err != nil {
		return nil, err
	}
	return &MakerLogNewCupIterator{contract: _Maker.contract, event: "LogNewCup", logs: logs, sub: sub}, nil
}

// WatchLogNewCup is a free log subscription operation binding the contract event 0x89b8893b806db50897c8e2362c71571cfaeb9761ee40727f683f1793cda9df16.
//
// Solidity: e LogNewCup(lad indexed address, cup bytes32)
func (_Maker *MakerFilterer) WatchLogNewCup(opts *bind.WatchOpts, sink chan<- *MakerLogNewCup, lad []common.Address) (event.Subscription, error) {

	var ladRule []interface{}
	for _, ladItem := range lad {
		ladRule = append(ladRule, ladItem)
	}

	logs, sub, err := _Maker.contract.WatchLogs(opts, "LogNewCup", ladRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerLogNewCup)
				if err := _Maker.contract.UnpackLog(event, "LogNewCup", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MakerLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Maker contract.
type MakerLogSetAuthorityIterator struct {
	Event *MakerLogSetAuthority // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MakerLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerLogSetAuthority)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MakerLogSetAuthority)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MakerLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerLogSetAuthority represents a LogSetAuthority event raised by the Maker contract.
type MakerLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_Maker *MakerFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*MakerLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Maker.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &MakerLogSetAuthorityIterator{contract: _Maker.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_Maker *MakerFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *MakerLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Maker.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerLogSetAuthority)
				if err := _Maker.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MakerLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Maker contract.
type MakerLogSetOwnerIterator struct {
	Event *MakerLogSetOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MakerLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerLogSetOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MakerLogSetOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MakerLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerLogSetOwner represents a LogSetOwner event raised by the Maker contract.
type MakerLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_Maker *MakerFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*MakerLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Maker.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MakerLogSetOwnerIterator{contract: _Maker.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_Maker *MakerFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *MakerLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Maker.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerLogSetOwner)
				if err := _Maker.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
