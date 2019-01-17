// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compute\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"unset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"name\":\"values\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min_\",\"type\":\"uint96\"}],\"name\":\"setMin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"bytes12\"},{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"bytes12\"}],\"name\":\"unset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next_\",\"type\":\"bytes12\"}],\"name\":\"setNext\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"min\",\"outputs\":[{\"name\":\"\",\"type\":\"uint96\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Oracle *OracleCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Oracle *OracleSession) Authority() (common.Address, error) {
	return _Oracle.Contract.Authority(&_Oracle.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Oracle *OracleCallerSession) Authority() (common.Address, error) {
	return _Oracle.Contract.Authority(&_Oracle.CallOpts)
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_Oracle *OracleCaller) Compute(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Oracle.contract.Call(opts, out, "compute")
	return *ret0, *ret1, err
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_Oracle *OracleSession) Compute() ([32]byte, bool, error) {
	return _Oracle.Contract.Compute(&_Oracle.CallOpts)
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_Oracle *OracleCallerSession) Compute() ([32]byte, bool, error) {
	return _Oracle.Contract.Compute(&_Oracle.CallOpts)
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes( address) constant returns(bytes12)
func (_Oracle *OracleCaller) Indexes(opts *bind.CallOpts, arg0 common.Address) ([12]byte, error) {
	var (
		ret0 = new([12]byte)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "indexes", arg0)
	return *ret0, err
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes( address) constant returns(bytes12)
func (_Oracle *OracleSession) Indexes(arg0 common.Address) ([12]byte, error) {
	return _Oracle.Contract.Indexes(&_Oracle.CallOpts, arg0)
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes( address) constant returns(bytes12)
func (_Oracle *OracleCallerSession) Indexes(arg0 common.Address) ([12]byte, error) {
	return _Oracle.Contract.Indexes(&_Oracle.CallOpts, arg0)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_Oracle *OracleCaller) Min(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "min")
	return *ret0, err
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_Oracle *OracleSession) Min() (*big.Int, error) {
	return _Oracle.Contract.Min(&_Oracle.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_Oracle *OracleCallerSession) Min() (*big.Int, error) {
	return _Oracle.Contract.Min(&_Oracle.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_Oracle *OracleCaller) Next(opts *bind.CallOpts) ([12]byte, error) {
	var (
		ret0 = new([12]byte)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "next")
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_Oracle *OracleSession) Next() ([12]byte, error) {
	return _Oracle.Contract.Next(&_Oracle.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_Oracle *OracleCallerSession) Next() ([12]byte, error) {
	return _Oracle.Contract.Next(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Oracle *OracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Oracle *OracleSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Oracle *OracleCallerSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_Oracle *OracleCaller) Peek(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Oracle.contract.Call(opts, out, "peek")
	return *ret0, *ret1, err
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_Oracle *OracleSession) Peek() ([32]byte, bool, error) {
	return _Oracle.Contract.Peek(&_Oracle.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_Oracle *OracleCallerSession) Peek() ([32]byte, bool, error) {
	return _Oracle.Contract.Peek(&_Oracle.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_Oracle *OracleCaller) Read(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "read")
	return *ret0, err
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_Oracle *OracleSession) Read() ([32]byte, error) {
	return _Oracle.Contract.Read(&_Oracle.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_Oracle *OracleCallerSession) Read() ([32]byte, error) {
	return _Oracle.Contract.Read(&_Oracle.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values( bytes12) constant returns(address)
func (_Oracle *OracleCaller) Values(opts *bind.CallOpts, arg0 [12]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "values", arg0)
	return *ret0, err
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values( bytes12) constant returns(address)
func (_Oracle *OracleSession) Values(arg0 [12]byte) (common.Address, error) {
	return _Oracle.Contract.Values(&_Oracle.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values( bytes12) constant returns(address)
func (_Oracle *OracleCallerSession) Values(arg0 [12]byte) (common.Address, error) {
	return _Oracle.Contract.Values(&_Oracle.CallOpts, arg0)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Oracle *OracleTransactor) Poke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "poke")
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Oracle *OracleSession) Poke() (*types.Transaction, error) {
	return _Oracle.Contract.Poke(&_Oracle.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Oracle *OracleTransactorSession) Poke() (*types.Transaction, error) {
	return _Oracle.Contract.Poke(&_Oracle.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(pos bytes12, wat address) returns()
func (_Oracle *OracleTransactor) Set(opts *bind.TransactOpts, pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "set", pos, wat)
}

// Set is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(pos bytes12, wat address) returns()
func (_Oracle *OracleSession) Set(pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Set(&_Oracle.TransactOpts, pos, wat)
}

// Set is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(pos bytes12, wat address) returns()
func (_Oracle *OracleTransactorSession) Set(pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Set(&_Oracle.TransactOpts, pos, wat)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Oracle *OracleTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Oracle *OracleSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAuthority(&_Oracle.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Oracle *OracleTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAuthority(&_Oracle.TransactOpts, authority_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(min_ uint96) returns()
func (_Oracle *OracleTransactor) SetMin(opts *bind.TransactOpts, min_ *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setMin", min_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(min_ uint96) returns()
func (_Oracle *OracleSession) SetMin(min_ *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetMin(&_Oracle.TransactOpts, min_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(min_ uint96) returns()
func (_Oracle *OracleTransactorSession) SetMin(min_ *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetMin(&_Oracle.TransactOpts, min_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(next_ bytes12) returns()
func (_Oracle *OracleTransactor) SetNext(opts *bind.TransactOpts, next_ [12]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setNext", next_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(next_ bytes12) returns()
func (_Oracle *OracleSession) SetNext(next_ [12]byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetNext(&_Oracle.TransactOpts, next_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(next_ bytes12) returns()
func (_Oracle *OracleTransactorSession) SetNext(next_ [12]byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetNext(&_Oracle.TransactOpts, next_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Oracle *OracleTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Oracle *OracleSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetOwner(&_Oracle.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Oracle *OracleTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetOwner(&_Oracle.TransactOpts, owner_)
}

// Unset is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(pos bytes12) returns()
func (_Oracle *OracleTransactor) Unset(opts *bind.TransactOpts, pos [12]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "unset", pos)
}

// Unset is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(pos bytes12) returns()
func (_Oracle *OracleSession) Unset(pos [12]byte) (*types.Transaction, error) {
	return _Oracle.Contract.Unset(&_Oracle.TransactOpts, pos)
}

// Unset is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(pos bytes12) returns()
func (_Oracle *OracleTransactorSession) Unset(pos [12]byte) (*types.Transaction, error) {
	return _Oracle.Contract.Unset(&_Oracle.TransactOpts, pos)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Oracle *OracleTransactor) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "void")
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Oracle *OracleSession) Void() (*types.Transaction, error) {
	return _Oracle.Contract.Void(&_Oracle.TransactOpts)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Oracle *OracleTransactorSession) Void() (*types.Transaction, error) {
	return _Oracle.Contract.Void(&_Oracle.TransactOpts)
}

// OracleLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Oracle contract.
type OracleLogSetAuthorityIterator struct {
	Event *OracleLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *OracleLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleLogSetAuthority)
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
		it.Event = new(OracleLogSetAuthority)
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
func (it *OracleLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleLogSetAuthority represents a LogSetAuthority event raised by the Oracle contract.
type OracleLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_Oracle *OracleFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*OracleLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &OracleLogSetAuthorityIterator{contract: _Oracle.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_Oracle *OracleFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *OracleLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleLogSetAuthority)
				if err := _Oracle.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// OracleLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Oracle contract.
type OracleLogSetOwnerIterator struct {
	Event *OracleLogSetOwner // Event containing the contract specifics and raw log

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
func (it *OracleLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleLogSetOwner)
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
		it.Event = new(OracleLogSetOwner)
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
func (it *OracleLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleLogSetOwner represents a LogSetOwner event raised by the Oracle contract.
type OracleLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_Oracle *OracleFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*OracleLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &OracleLogSetOwnerIterator{contract: _Oracle.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_Oracle *OracleFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *OracleLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleLogSetOwner)
				if err := _Oracle.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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
