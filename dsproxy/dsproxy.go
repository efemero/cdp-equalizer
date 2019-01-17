// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dsproxy

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

// DsproxyABI is the input ABI used to generate the binding from.
const DsproxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"response\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cache\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cacheAddr\",\"type\":\"address\"}],\"name\":\"setCache\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cacheAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// Dsproxy is an auto generated Go binding around an Ethereum contract.
type Dsproxy struct {
	DsproxyCaller     // Read-only binding to the contract
	DsproxyTransactor // Write-only binding to the contract
	DsproxyFilterer   // Log filterer for contract events
}

// DsproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DsproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DsproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DsproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DsproxySession struct {
	Contract     *Dsproxy          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DsproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DsproxyCallerSession struct {
	Contract *DsproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DsproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DsproxyTransactorSession struct {
	Contract     *DsproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DsproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DsproxyRaw struct {
	Contract *Dsproxy // Generic contract binding to access the raw methods on
}

// DsproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DsproxyCallerRaw struct {
	Contract *DsproxyCaller // Generic read-only contract binding to access the raw methods on
}

// DsproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DsproxyTransactorRaw struct {
	Contract *DsproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDsproxy creates a new instance of Dsproxy, bound to a specific deployed contract.
func NewDsproxy(address common.Address, backend bind.ContractBackend) (*Dsproxy, error) {
	contract, err := bindDsproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dsproxy{DsproxyCaller: DsproxyCaller{contract: contract}, DsproxyTransactor: DsproxyTransactor{contract: contract}, DsproxyFilterer: DsproxyFilterer{contract: contract}}, nil
}

// NewDsproxyCaller creates a new read-only instance of Dsproxy, bound to a specific deployed contract.
func NewDsproxyCaller(address common.Address, caller bind.ContractCaller) (*DsproxyCaller, error) {
	contract, err := bindDsproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DsproxyCaller{contract: contract}, nil
}

// NewDsproxyTransactor creates a new write-only instance of Dsproxy, bound to a specific deployed contract.
func NewDsproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DsproxyTransactor, error) {
	contract, err := bindDsproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DsproxyTransactor{contract: contract}, nil
}

// NewDsproxyFilterer creates a new log filterer instance of Dsproxy, bound to a specific deployed contract.
func NewDsproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DsproxyFilterer, error) {
	contract, err := bindDsproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DsproxyFilterer{contract: contract}, nil
}

// bindDsproxy binds a generic wrapper to an already deployed contract.
func bindDsproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DsproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dsproxy *DsproxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dsproxy.Contract.DsproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dsproxy *DsproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsproxy.Contract.DsproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dsproxy *DsproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dsproxy.Contract.DsproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dsproxy *DsproxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dsproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dsproxy *DsproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dsproxy *DsproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dsproxy.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Dsproxy *DsproxyCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsproxy.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Dsproxy *DsproxySession) Authority() (common.Address, error) {
	return _Dsproxy.Contract.Authority(&_Dsproxy.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Dsproxy *DsproxyCallerSession) Authority() (common.Address, error) {
	return _Dsproxy.Contract.Authority(&_Dsproxy.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_Dsproxy *DsproxyCaller) Cache(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsproxy.contract.Call(opts, out, "cache")
	return *ret0, err
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_Dsproxy *DsproxySession) Cache() (common.Address, error) {
	return _Dsproxy.Contract.Cache(&_Dsproxy.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_Dsproxy *DsproxyCallerSession) Cache() (common.Address, error) {
	return _Dsproxy.Contract.Cache(&_Dsproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dsproxy *DsproxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsproxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dsproxy *DsproxySession) Owner() (common.Address, error) {
	return _Dsproxy.Contract.Owner(&_Dsproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dsproxy *DsproxyCallerSession) Owner() (common.Address, error) {
	return _Dsproxy.Contract.Owner(&_Dsproxy.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(_target address, _data bytes) returns(response bytes32)
func (_Dsproxy *DsproxyTransactor) Execute(opts *bind.TransactOpts, _target common.Address, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "execute", _target, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(_target address, _data bytes) returns(response bytes32)
func (_Dsproxy *DsproxySession) Execute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.Contract.Execute(&_Dsproxy.TransactOpts, _target, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(_target address, _data bytes) returns(response bytes32)
func (_Dsproxy *DsproxyTransactorSession) Execute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _Dsproxy.Contract.Execute(&_Dsproxy.TransactOpts, _target, _data)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Dsproxy *DsproxyTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Dsproxy *DsproxySession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetAuthority(&_Dsproxy.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Dsproxy *DsproxyTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetAuthority(&_Dsproxy.TransactOpts, authority_)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(_cacheAddr address) returns(bool)
func (_Dsproxy *DsproxyTransactor) SetCache(opts *bind.TransactOpts, _cacheAddr common.Address) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "setCache", _cacheAddr)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(_cacheAddr address) returns(bool)
func (_Dsproxy *DsproxySession) SetCache(_cacheAddr common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetCache(&_Dsproxy.TransactOpts, _cacheAddr)
}

// SetCache is a paid mutator transaction binding the contract method 0x948f5076.
//
// Solidity: function setCache(_cacheAddr address) returns(bool)
func (_Dsproxy *DsproxyTransactorSession) SetCache(_cacheAddr common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetCache(&_Dsproxy.TransactOpts, _cacheAddr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Dsproxy *DsproxyTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Dsproxy *DsproxySession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetOwner(&_Dsproxy.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Dsproxy *DsproxyTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Dsproxy.Contract.SetOwner(&_Dsproxy.TransactOpts, owner_)
}

// DsproxyLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Dsproxy contract.
type DsproxyLogSetAuthorityIterator struct {
	Event *DsproxyLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *DsproxyLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DsproxyLogSetAuthority)
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
		it.Event = new(DsproxyLogSetAuthority)
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
func (it *DsproxyLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DsproxyLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DsproxyLogSetAuthority represents a LogSetAuthority event raised by the Dsproxy contract.
type DsproxyLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_Dsproxy *DsproxyFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*DsproxyLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Dsproxy.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &DsproxyLogSetAuthorityIterator{contract: _Dsproxy.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_Dsproxy *DsproxyFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *DsproxyLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Dsproxy.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DsproxyLogSetAuthority)
				if err := _Dsproxy.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// DsproxyLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Dsproxy contract.
type DsproxyLogSetOwnerIterator struct {
	Event *DsproxyLogSetOwner // Event containing the contract specifics and raw log

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
func (it *DsproxyLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DsproxyLogSetOwner)
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
		it.Event = new(DsproxyLogSetOwner)
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
func (it *DsproxyLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DsproxyLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DsproxyLogSetOwner represents a LogSetOwner event raised by the Dsproxy contract.
type DsproxyLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_Dsproxy *DsproxyFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*DsproxyLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Dsproxy.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DsproxyLogSetOwnerIterator{contract: _Dsproxy.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_Dsproxy *DsproxyFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *DsproxyLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Dsproxy.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DsproxyLogSetOwner)
				if err := _Dsproxy.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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
