// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"Added\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"a\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f553480156012575f5ffd5b506101de806100205f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80630dbe671f146100385780634f2be91f14610056575b5f5ffd5b610040610074565b60405161004d91906100e8565b60405180910390f35b61005e610079565b60405161006b91906100e8565b60405180910390f35b5f5481565b5f60015f54610088919061012e565b5f547fd822787006f6794403d16891b4a593a2085bdc9fc3d768f12f0b13843fdaa59b60405160405180910390a35f5f8154809291906100c790610161565b91905055905090565b5f819050919050565b6100e2816100d0565b82525050565b5f6020820190506100fb5f8301846100d9565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610138826100d0565b9150610143836100d0565b925082820190508082111561015b5761015a610101565b5b92915050565b5f61016b826100d0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361019d5761019c610101565b5b60018201905091905056fea2646970667358221220415272adca5b362202750d97b244b38d5ba33df685d11cf55d41ebe9607b032c64736f6c634300081e0033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Store *StoreCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Store *StoreSession) A() (*big.Int, error) {
	return _Store.Contract.A(&_Store.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Store *StoreCallerSession) A() (*big.Int, error) {
	return _Store.Contract.A(&_Store.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns(uint256)
func (_Store *StoreTransactor) Add(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "add")
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns(uint256)
func (_Store *StoreSession) Add() (*types.Transaction, error) {
	return _Store.Contract.Add(&_Store.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns(uint256)
func (_Store *StoreTransactorSession) Add() (*types.Transaction, error) {
	return _Store.Contract.Add(&_Store.TransactOpts)
}

// StoreAddedIterator is returned from FilterAdded and is used to iterate over the raw logs and unpacked data for Added events raised by the Store contract.
type StoreAddedIterator struct {
	Event *StoreAdded // Event containing the contract specifics and raw log

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
func (it *StoreAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreAdded)
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
		it.Event = new(StoreAdded)
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
func (it *StoreAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreAdded represents a Added event raised by the Store contract.
type StoreAdded struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdded is a free log retrieval operation binding the contract event 0xd822787006f6794403d16891b4a593a2085bdc9fc3d768f12f0b13843fdaa59b.
//
// Solidity: event Added(uint256 indexed oldValue, uint256 indexed newValue)
func (_Store *StoreFilterer) FilterAdded(opts *bind.FilterOpts, oldValue []*big.Int, newValue []*big.Int) (*StoreAddedIterator, error) {

	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Added", oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return &StoreAddedIterator{contract: _Store.contract, event: "Added", logs: logs, sub: sub}, nil
}

// WatchAdded is a free log subscription operation binding the contract event 0xd822787006f6794403d16891b4a593a2085bdc9fc3d768f12f0b13843fdaa59b.
//
// Solidity: event Added(uint256 indexed oldValue, uint256 indexed newValue)
func (_Store *StoreFilterer) WatchAdded(opts *bind.WatchOpts, sink chan<- *StoreAdded, oldValue []*big.Int, newValue []*big.Int) (event.Subscription, error) {

	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Added", oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreAdded)
				if err := _Store.contract.UnpackLog(event, "Added", log); err != nil {
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

// ParseAdded is a log parse operation binding the contract event 0xd822787006f6794403d16891b4a593a2085bdc9fc3d768f12f0b13843fdaa59b.
//
// Solidity: event Added(uint256 indexed oldValue, uint256 indexed newValue)
func (_Store *StoreFilterer) ParseAdded(log types.Log) (*StoreAdded, error) {
	event := new(StoreAdded)
	if err := _Store.contract.UnpackLog(event, "Added", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
