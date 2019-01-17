// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package portal

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

// PortalABI is the input ABI used to generate the binding from.
const PortalABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"draw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"jam\",\"type\":\"uint256\"},{\"name\":\"wad\",\"type\":\"uint256\"},{\"name\":\"otc_\",\"type\":\"address\"}],\"name\":\"wipeAndFree\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"lockAndDraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"lockAndDraw\",\"outputs\":[{\"name\":\"cup\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registry_\",\"type\":\"address\"},{\"name\":\"tub_\",\"type\":\"address\"}],\"name\":\"createAndOpen\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"otc_\",\"type\":\"address\"}],\"name\":\"shut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"wad\",\"type\":\"uint256\"},{\"name\":\"otc_\",\"type\":\"address\"}],\"name\":\"wipe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"wipe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"}],\"name\":\"open\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"shut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registry_\",\"type\":\"address\"},{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"createOpenLockAndDraw\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"lad\",\"type\":\"address\"}],\"name\":\"give\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registry_\",\"type\":\"address\"},{\"name\":\"tub_\",\"type\":\"address\"}],\"name\":\"createOpenAndLock\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"jam\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tub_\",\"type\":\"address\"},{\"name\":\"cup\",\"type\":\"bytes32\"},{\"name\":\"jam\",\"type\":\"uint256\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"wipeAndFree\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Portal is an auto generated Go binding around an Ethereum contract.
type Portal struct {
	PortalCaller     // Read-only binding to the contract
	PortalTransactor // Write-only binding to the contract
	PortalFilterer   // Log filterer for contract events
}

// PortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type PortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PortalSession struct {
	Contract     *Portal           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PortalCallerSession struct {
	Contract *PortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PortalTransactorSession struct {
	Contract     *PortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type PortalRaw struct {
	Contract *Portal // Generic contract binding to access the raw methods on
}

// PortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PortalCallerRaw struct {
	Contract *PortalCaller // Generic read-only contract binding to access the raw methods on
}

// PortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PortalTransactorRaw struct {
	Contract *PortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPortal creates a new instance of Portal, bound to a specific deployed contract.
func NewPortal(address common.Address, backend bind.ContractBackend) (*Portal, error) {
	contract, err := bindPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Portal{PortalCaller: PortalCaller{contract: contract}, PortalTransactor: PortalTransactor{contract: contract}, PortalFilterer: PortalFilterer{contract: contract}}, nil
}

// NewPortalCaller creates a new read-only instance of Portal, bound to a specific deployed contract.
func NewPortalCaller(address common.Address, caller bind.ContractCaller) (*PortalCaller, error) {
	contract, err := bindPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PortalCaller{contract: contract}, nil
}

// NewPortalTransactor creates a new write-only instance of Portal, bound to a specific deployed contract.
func NewPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*PortalTransactor, error) {
	contract, err := bindPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PortalTransactor{contract: contract}, nil
}

// NewPortalFilterer creates a new log filterer instance of Portal, bound to a specific deployed contract.
func NewPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*PortalFilterer, error) {
	contract, err := bindPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PortalFilterer{contract: contract}, nil
}

// bindPortal binds a generic wrapper to an already deployed contract.
func bindPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Portal *PortalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Portal.Contract.PortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Portal *PortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.Contract.PortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Portal *PortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Portal.Contract.PortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Portal *PortalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Portal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Portal *PortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Portal *PortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Portal.Contract.contract.Transact(opts, method, params...)
}

// CreateAndOpen is a paid mutator transaction binding the contract method 0x581f3c50.
//
// Solidity: function createAndOpen(registry_ address, tub_ address) returns(proxy address, cup bytes32)
func (_Portal *PortalTransactor) CreateAndOpen(opts *bind.TransactOpts, registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "createAndOpen", registry_, tub_)
}

// CreateAndOpen is a paid mutator transaction binding the contract method 0x581f3c50.
//
// Solidity: function createAndOpen(registry_ address, tub_ address) returns(proxy address, cup bytes32)
func (_Portal *PortalSession) CreateAndOpen(registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _Portal.Contract.CreateAndOpen(&_Portal.TransactOpts, registry_, tub_)
}

// CreateAndOpen is a paid mutator transaction binding the contract method 0x581f3c50.
//
// Solidity: function createAndOpen(registry_ address, tub_ address) returns(proxy address, cup bytes32)
func (_Portal *PortalTransactorSession) CreateAndOpen(registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _Portal.Contract.CreateAndOpen(&_Portal.TransactOpts, registry_, tub_)
}

// CreateOpenAndLock is a paid mutator transaction binding the contract method 0xeefe3818.
//
// Solidity: function createOpenAndLock(registry_ address, tub_ address) returns(proxy address, cup bytes32)
func (_Portal *PortalTransactor) CreateOpenAndLock(opts *bind.TransactOpts, registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "createOpenAndLock", registry_, tub_)
}

// CreateOpenAndLock is a paid mutator transaction binding the contract method 0xeefe3818.
//
// Solidity: function createOpenAndLock(registry_ address, tub_ address) returns(proxy address, cup bytes32)
func (_Portal *PortalSession) CreateOpenAndLock(registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _Portal.Contract.CreateOpenAndLock(&_Portal.TransactOpts, registry_, tub_)
}

// CreateOpenAndLock is a paid mutator transaction binding the contract method 0xeefe3818.
//
// Solidity: function createOpenAndLock(registry_ address, tub_ address) returns(proxy address, cup bytes32)
func (_Portal *PortalTransactorSession) CreateOpenAndLock(registry_ common.Address, tub_ common.Address) (*types.Transaction, error) {
	return _Portal.Contract.CreateOpenAndLock(&_Portal.TransactOpts, registry_, tub_)
}

// CreateOpenLockAndDraw is a paid mutator transaction binding the contract method 0xd3140a65.
//
// Solidity: function createOpenLockAndDraw(registry_ address, tub_ address, wad uint256) returns(proxy address, cup bytes32)
func (_Portal *PortalTransactor) CreateOpenLockAndDraw(opts *bind.TransactOpts, registry_ common.Address, tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "createOpenLockAndDraw", registry_, tub_, wad)
}

// CreateOpenLockAndDraw is a paid mutator transaction binding the contract method 0xd3140a65.
//
// Solidity: function createOpenLockAndDraw(registry_ address, tub_ address, wad uint256) returns(proxy address, cup bytes32)
func (_Portal *PortalSession) CreateOpenLockAndDraw(registry_ common.Address, tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.CreateOpenLockAndDraw(&_Portal.TransactOpts, registry_, tub_, wad)
}

// CreateOpenLockAndDraw is a paid mutator transaction binding the contract method 0xd3140a65.
//
// Solidity: function createOpenLockAndDraw(registry_ address, tub_ address, wad uint256) returns(proxy address, cup bytes32)
func (_Portal *PortalTransactorSession) CreateOpenLockAndDraw(registry_ common.Address, tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.CreateOpenLockAndDraw(&_Portal.TransactOpts, registry_, tub_, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x0344a36f.
//
// Solidity: function draw(tub_ address, cup bytes32, wad uint256) returns()
func (_Portal *PortalTransactor) Draw(opts *bind.TransactOpts, tub_ common.Address, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "draw", tub_, cup, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x0344a36f.
//
// Solidity: function draw(tub_ address, cup bytes32, wad uint256) returns()
func (_Portal *PortalSession) Draw(tub_ common.Address, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.Draw(&_Portal.TransactOpts, tub_, cup, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x0344a36f.
//
// Solidity: function draw(tub_ address, cup bytes32, wad uint256) returns()
func (_Portal *PortalTransactorSession) Draw(tub_ common.Address, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.Draw(&_Portal.TransactOpts, tub_, cup, wad)
}

// Free is a paid mutator transaction binding the contract method 0xf9ef04be.
//
// Solidity: function free(tub_ address, cup bytes32, jam uint256) returns()
func (_Portal *PortalTransactor) Free(opts *bind.TransactOpts, tub_ common.Address, cup [32]byte, jam *big.Int) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "free", tub_, cup, jam)
}

// Free is a paid mutator transaction binding the contract method 0xf9ef04be.
//
// Solidity: function free(tub_ address, cup bytes32, jam uint256) returns()
func (_Portal *PortalSession) Free(tub_ common.Address, cup [32]byte, jam *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.Free(&_Portal.TransactOpts, tub_, cup, jam)
}

// Free is a paid mutator transaction binding the contract method 0xf9ef04be.
//
// Solidity: function free(tub_ address, cup bytes32, jam uint256) returns()
func (_Portal *PortalTransactorSession) Free(tub_ common.Address, cup [32]byte, jam *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.Free(&_Portal.TransactOpts, tub_, cup, jam)
}

// Give is a paid mutator transaction binding the contract method 0xda93dfcf.
//
// Solidity: function give(tub_ address, cup bytes32, lad address) returns()
func (_Portal *PortalTransactor) Give(opts *bind.TransactOpts, tub_ common.Address, cup [32]byte, lad common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "give", tub_, cup, lad)
}

// Give is a paid mutator transaction binding the contract method 0xda93dfcf.
//
// Solidity: function give(tub_ address, cup bytes32, lad address) returns()
func (_Portal *PortalSession) Give(tub_ common.Address, cup [32]byte, lad common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Give(&_Portal.TransactOpts, tub_, cup, lad)
}

// Give is a paid mutator transaction binding the contract method 0xda93dfcf.
//
// Solidity: function give(tub_ address, cup bytes32, lad address) returns()
func (_Portal *PortalTransactorSession) Give(tub_ common.Address, cup [32]byte, lad common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Give(&_Portal.TransactOpts, tub_, cup, lad)
}

// Lock is a paid mutator transaction binding the contract method 0xbc25a810.
//
// Solidity: function lock(tub_ address, cup bytes32) returns()
func (_Portal *PortalTransactor) Lock(opts *bind.TransactOpts, tub_ common.Address, cup [32]byte) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "lock", tub_, cup)
}

// Lock is a paid mutator transaction binding the contract method 0xbc25a810.
//
// Solidity: function lock(tub_ address, cup bytes32) returns()
func (_Portal *PortalSession) Lock(tub_ common.Address, cup [32]byte) (*types.Transaction, error) {
	return _Portal.Contract.Lock(&_Portal.TransactOpts, tub_, cup)
}

// Lock is a paid mutator transaction binding the contract method 0xbc25a810.
//
// Solidity: function lock(tub_ address, cup bytes32) returns()
func (_Portal *PortalTransactorSession) Lock(tub_ common.Address, cup [32]byte) (*types.Transaction, error) {
	return _Portal.Contract.Lock(&_Portal.TransactOpts, tub_, cup)
}

// LockAndDraw is a paid mutator transaction binding the contract method 0x516e9aec.
//
// Solidity: function lockAndDraw(tub_ address, wad uint256) returns(cup bytes32)
func (_Portal *PortalTransactor) LockAndDraw(opts *bind.TransactOpts, tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "lockAndDraw", tub_, wad)
}

// LockAndDraw is a paid mutator transaction binding the contract method 0x516e9aec.
//
// Solidity: function lockAndDraw(tub_ address, wad uint256) returns(cup bytes32)
func (_Portal *PortalSession) LockAndDraw(tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.LockAndDraw(&_Portal.TransactOpts, tub_, wad)
}

// LockAndDraw is a paid mutator transaction binding the contract method 0x516e9aec.
//
// Solidity: function lockAndDraw(tub_ address, wad uint256) returns(cup bytes32)
func (_Portal *PortalTransactorSession) LockAndDraw(tub_ common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.LockAndDraw(&_Portal.TransactOpts, tub_, wad)
}

// Open is a paid mutator transaction binding the contract method 0xb95460f8.
//
// Solidity: function open(tub_ address) returns(bytes32)
func (_Portal *PortalTransactor) Open(opts *bind.TransactOpts, tub_ common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "open", tub_)
}

// Open is a paid mutator transaction binding the contract method 0xb95460f8.
//
// Solidity: function open(tub_ address) returns(bytes32)
func (_Portal *PortalSession) Open(tub_ common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Open(&_Portal.TransactOpts, tub_)
}

// Open is a paid mutator transaction binding the contract method 0xb95460f8.
//
// Solidity: function open(tub_ address) returns(bytes32)
func (_Portal *PortalTransactorSession) Open(tub_ common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Open(&_Portal.TransactOpts, tub_)
}

// Shut is a paid mutator transaction binding the contract method 0xbc244c11.
//
// Solidity: function shut(tub_ address, cup bytes32) returns()
func (_Portal *PortalTransactor) Shut(opts *bind.TransactOpts, tub_ common.Address, cup [32]byte) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "shut", tub_, cup)
}

// Shut is a paid mutator transaction binding the contract method 0xbc244c11.
//
// Solidity: function shut(tub_ address, cup bytes32) returns()
func (_Portal *PortalSession) Shut(tub_ common.Address, cup [32]byte) (*types.Transaction, error) {
	return _Portal.Contract.Shut(&_Portal.TransactOpts, tub_, cup)
}

// Shut is a paid mutator transaction binding the contract method 0xbc244c11.
//
// Solidity: function shut(tub_ address, cup bytes32) returns()
func (_Portal *PortalTransactorSession) Shut(tub_ common.Address, cup [32]byte) (*types.Transaction, error) {
	return _Portal.Contract.Shut(&_Portal.TransactOpts, tub_, cup)
}

// Wipe is a paid mutator transaction binding the contract method 0xa3dc65a7.
//
// Solidity: function wipe(tub_ address, cup bytes32, wad uint256) returns()
func (_Portal *PortalTransactor) Wipe(opts *bind.TransactOpts, tub_ common.Address, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "wipe", tub_, cup, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xa3dc65a7.
//
// Solidity: function wipe(tub_ address, cup bytes32, wad uint256) returns()
func (_Portal *PortalSession) Wipe(tub_ common.Address, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.Wipe(&_Portal.TransactOpts, tub_, cup, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xa3dc65a7.
//
// Solidity: function wipe(tub_ address, cup bytes32, wad uint256) returns()
func (_Portal *PortalTransactorSession) Wipe(tub_ common.Address, cup [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.Wipe(&_Portal.TransactOpts, tub_, cup, wad)
}

// WipeAndFree is a paid mutator transaction binding the contract method 0xfaed77ab.
//
// Solidity: function wipeAndFree(tub_ address, cup bytes32, jam uint256, wad uint256) returns()
func (_Portal *PortalTransactor) WipeAndFree(opts *bind.TransactOpts, tub_ common.Address, cup [32]byte, jam *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "wipeAndFree", tub_, cup, jam, wad)
}

// WipeAndFree is a paid mutator transaction binding the contract method 0xfaed77ab.
//
// Solidity: function wipeAndFree(tub_ address, cup bytes32, jam uint256, wad uint256) returns()
func (_Portal *PortalSession) WipeAndFree(tub_ common.Address, cup [32]byte, jam *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.WipeAndFree(&_Portal.TransactOpts, tub_, cup, jam, wad)
}

// WipeAndFree is a paid mutator transaction binding the contract method 0xfaed77ab.
//
// Solidity: function wipeAndFree(tub_ address, cup bytes32, jam uint256, wad uint256) returns()
func (_Portal *PortalTransactorSession) WipeAndFree(tub_ common.Address, cup [32]byte, jam *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Portal.Contract.WipeAndFree(&_Portal.TransactOpts, tub_, cup, jam, wad)
}
