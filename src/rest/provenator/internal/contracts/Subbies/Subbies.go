// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Subbies

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Subbie is an auto generated low-level Go binding around an user-defined struct.
type Subbie struct {
	MainType    uint8
	Status      uint8
	Name        string
	Location    string
	Description string
}

// IFactoryABI is the input ABI used to generate the binding from.
const IFactoryABI = "[{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IFactoryFuncSigs = map[string]string{
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
}

// IFactory is an auto generated Go binding around an Ethereum contract.
type IFactory struct {
	IFactoryCaller     // Read-only binding to the contract
	IFactoryTransactor // Write-only binding to the contract
	IFactoryFilterer   // Log filterer for contract events
}

// IFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFactorySession struct {
	Contract     *IFactory         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFactoryCallerSession struct {
	Contract *IFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFactoryTransactorSession struct {
	Contract     *IFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFactoryRaw struct {
	Contract *IFactory // Generic contract binding to access the raw methods on
}

// IFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFactoryCallerRaw struct {
	Contract *IFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFactoryTransactorRaw struct {
	Contract *IFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFactory creates a new instance of IFactory, bound to a specific deployed contract.
func NewIFactory(address common.Address, backend bind.ContractBackend) (*IFactory, error) {
	contract, err := bindIFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFactory{IFactoryCaller: IFactoryCaller{contract: contract}, IFactoryTransactor: IFactoryTransactor{contract: contract}, IFactoryFilterer: IFactoryFilterer{contract: contract}}, nil
}

// NewIFactoryCaller creates a new read-only instance of IFactory, bound to a specific deployed contract.
func NewIFactoryCaller(address common.Address, caller bind.ContractCaller) (*IFactoryCaller, error) {
	contract, err := bindIFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFactoryCaller{contract: contract}, nil
}

// NewIFactoryTransactor creates a new write-only instance of IFactory, bound to a specific deployed contract.
func NewIFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IFactoryTransactor, error) {
	contract, err := bindIFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFactoryTransactor{contract: contract}, nil
}

// NewIFactoryFilterer creates a new log filterer instance of IFactory, bound to a specific deployed contract.
func NewIFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IFactoryFilterer, error) {
	contract, err := bindIFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFactoryFilterer{contract: contract}, nil
}

// bindIFactory binds a generic wrapper to an already deployed contract.
func bindIFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFactory *IFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IFactory.Contract.IFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFactory *IFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFactory.Contract.IFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFactory *IFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFactory.Contract.IFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFactory *IFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFactory *IFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFactory *IFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFactory.Contract.contract.Transact(opts, method, params...)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IFactory *IFactoryCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFactory.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IFactory *IFactorySession) GetNum() (*big.Int, error) {
	return _IFactory.Contract.GetNum(&_IFactory.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IFactory *IFactoryCallerSession) GetNum() (*big.Int, error) {
	return _IFactory.Contract.GetNum(&_IFactory.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IFactory *IFactoryCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IFactory.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IFactory *IFactorySession) GetReference(_index *big.Int) ([32]byte, error) {
	return _IFactory.Contract.GetReference(&_IFactory.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IFactory *IFactoryCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _IFactory.Contract.GetReference(&_IFactory.CallOpts, _index)
}

// ISubbieABI is the input ABI used to generate the binding from.
const ISubbieABI = "[{\"inputs\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"mainType\",\"type\":\"uint8\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structSubbie\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSubbieStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISubbieFuncSigs maps the 4-byte function signature to its string representation.
var ISubbieFuncSigs = map[string]string{
	"e188130b": "addType(uint8)",
	"6d4ce63c": "get()",
	"222f74e7": "isType(uint8)",
	"2e49d78b": "setStatus(uint8)",
}

// ISubbie is an auto generated Go binding around an Ethereum contract.
type ISubbie struct {
	ISubbieCaller     // Read-only binding to the contract
	ISubbieTransactor // Write-only binding to the contract
	ISubbieFilterer   // Log filterer for contract events
}

// ISubbieCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISubbieCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubbieTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISubbieTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubbieFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISubbieFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubbieSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISubbieSession struct {
	Contract     *ISubbie          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISubbieCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISubbieCallerSession struct {
	Contract *ISubbieCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ISubbieTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISubbieTransactorSession struct {
	Contract     *ISubbieTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ISubbieRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISubbieRaw struct {
	Contract *ISubbie // Generic contract binding to access the raw methods on
}

// ISubbieCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISubbieCallerRaw struct {
	Contract *ISubbieCaller // Generic read-only contract binding to access the raw methods on
}

// ISubbieTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISubbieTransactorRaw struct {
	Contract *ISubbieTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISubbie creates a new instance of ISubbie, bound to a specific deployed contract.
func NewISubbie(address common.Address, backend bind.ContractBackend) (*ISubbie, error) {
	contract, err := bindISubbie(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISubbie{ISubbieCaller: ISubbieCaller{contract: contract}, ISubbieTransactor: ISubbieTransactor{contract: contract}, ISubbieFilterer: ISubbieFilterer{contract: contract}}, nil
}

// NewISubbieCaller creates a new read-only instance of ISubbie, bound to a specific deployed contract.
func NewISubbieCaller(address common.Address, caller bind.ContractCaller) (*ISubbieCaller, error) {
	contract, err := bindISubbie(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISubbieCaller{contract: contract}, nil
}

// NewISubbieTransactor creates a new write-only instance of ISubbie, bound to a specific deployed contract.
func NewISubbieTransactor(address common.Address, transactor bind.ContractTransactor) (*ISubbieTransactor, error) {
	contract, err := bindISubbie(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISubbieTransactor{contract: contract}, nil
}

// NewISubbieFilterer creates a new log filterer instance of ISubbie, bound to a specific deployed contract.
func NewISubbieFilterer(address common.Address, filterer bind.ContractFilterer) (*ISubbieFilterer, error) {
	contract, err := bindISubbie(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISubbieFilterer{contract: contract}, nil
}

// bindISubbie binds a generic wrapper to an already deployed contract.
func bindISubbie(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISubbieABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubbie *ISubbieRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubbie.Contract.ISubbieCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubbie *ISubbieRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubbie.Contract.ISubbieTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubbie *ISubbieRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubbie.Contract.ISubbieTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubbie *ISubbieCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubbie.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubbie *ISubbieTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubbie.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubbie *ISubbieTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubbie.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(address, (uint8,uint8,string,string,string))
func (_ISubbie *ISubbieCaller) Get(opts *bind.CallOpts) (common.Address, Subbie, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(Subbie)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ISubbie.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(address, (uint8,uint8,string,string,string))
func (_ISubbie *ISubbieSession) Get() (common.Address, Subbie, error) {
	return _ISubbie.Contract.Get(&_ISubbie.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(address, (uint8,uint8,string,string,string))
func (_ISubbie *ISubbieCallerSession) Get() (common.Address, Subbie, error) {
	return _ISubbie.Contract.Get(&_ISubbie.CallOpts)
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_ISubbie *ISubbieCaller) IsType(opts *bind.CallOpts, _type uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ISubbie.contract.Call(opts, out, "isType", _type)
	return *ret0, err
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_ISubbie *ISubbieSession) IsType(_type uint8) (bool, error) {
	return _ISubbie.Contract.IsType(&_ISubbie.CallOpts, _type)
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_ISubbie *ISubbieCallerSession) IsType(_type uint8) (bool, error) {
	return _ISubbie.Contract.IsType(&_ISubbie.CallOpts, _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_ISubbie *ISubbieTransactor) AddType(opts *bind.TransactOpts, _type uint8) (*types.Transaction, error) {
	return _ISubbie.contract.Transact(opts, "addType", _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_ISubbie *ISubbieSession) AddType(_type uint8) (*types.Transaction, error) {
	return _ISubbie.Contract.AddType(&_ISubbie.TransactOpts, _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_ISubbie *ISubbieTransactorSession) AddType(_type uint8) (*types.Transaction, error) {
	return _ISubbie.Contract.AddType(&_ISubbie.TransactOpts, _type)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_ISubbie *ISubbieTransactor) SetStatus(opts *bind.TransactOpts, _status uint8) (*types.Transaction, error) {
	return _ISubbie.contract.Transact(opts, "setStatus", _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_ISubbie *ISubbieSession) SetStatus(_status uint8) (*types.Transaction, error) {
	return _ISubbie.Contract.SetStatus(&_ISubbie.TransactOpts, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_ISubbie *ISubbieTransactorSession) SetStatus(_status uint8) (*types.Transaction, error) {
	return _ISubbie.Contract.SetStatus(&_ISubbie.TransactOpts, _status)
}

// ISubbieFactoryABI is the input ABI used to generate the binding from.
const ISubbieFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"mainType\",\"type\":\"uint8\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structSubbie\",\"name\":\"_subbie\",\"type\":\"tuple\"}],\"name\":\"addSubbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"},{\"internalType\":\"enumContractorTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"}],\"name\":\"getSubbie\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"mainType\",\"type\":\"uint8\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structSubbie\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"},{\"internalType\":\"enumContractorTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISubbieFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ISubbieFactoryFuncSigs = map[string]string{
	"e28afe96": "addSubbie(address,(uint8,uint8,string,string,string))",
	"90991c91": "addType(address,uint8)",
	"9790ea30": "getSubbie(address)",
	"6359e6a3": "isType(address,uint8)",
	"278e07ce": "setStatus(address,uint8)",
}

// ISubbieFactory is an auto generated Go binding around an Ethereum contract.
type ISubbieFactory struct {
	ISubbieFactoryCaller     // Read-only binding to the contract
	ISubbieFactoryTransactor // Write-only binding to the contract
	ISubbieFactoryFilterer   // Log filterer for contract events
}

// ISubbieFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISubbieFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubbieFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISubbieFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubbieFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISubbieFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubbieFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISubbieFactorySession struct {
	Contract     *ISubbieFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISubbieFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISubbieFactoryCallerSession struct {
	Contract *ISubbieFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ISubbieFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISubbieFactoryTransactorSession struct {
	Contract     *ISubbieFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ISubbieFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISubbieFactoryRaw struct {
	Contract *ISubbieFactory // Generic contract binding to access the raw methods on
}

// ISubbieFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISubbieFactoryCallerRaw struct {
	Contract *ISubbieFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ISubbieFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISubbieFactoryTransactorRaw struct {
	Contract *ISubbieFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISubbieFactory creates a new instance of ISubbieFactory, bound to a specific deployed contract.
func NewISubbieFactory(address common.Address, backend bind.ContractBackend) (*ISubbieFactory, error) {
	contract, err := bindISubbieFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISubbieFactory{ISubbieFactoryCaller: ISubbieFactoryCaller{contract: contract}, ISubbieFactoryTransactor: ISubbieFactoryTransactor{contract: contract}, ISubbieFactoryFilterer: ISubbieFactoryFilterer{contract: contract}}, nil
}

// NewISubbieFactoryCaller creates a new read-only instance of ISubbieFactory, bound to a specific deployed contract.
func NewISubbieFactoryCaller(address common.Address, caller bind.ContractCaller) (*ISubbieFactoryCaller, error) {
	contract, err := bindISubbieFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISubbieFactoryCaller{contract: contract}, nil
}

// NewISubbieFactoryTransactor creates a new write-only instance of ISubbieFactory, bound to a specific deployed contract.
func NewISubbieFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ISubbieFactoryTransactor, error) {
	contract, err := bindISubbieFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISubbieFactoryTransactor{contract: contract}, nil
}

// NewISubbieFactoryFilterer creates a new log filterer instance of ISubbieFactory, bound to a specific deployed contract.
func NewISubbieFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ISubbieFactoryFilterer, error) {
	contract, err := bindISubbieFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISubbieFactoryFilterer{contract: contract}, nil
}

// bindISubbieFactory binds a generic wrapper to an already deployed contract.
func bindISubbieFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISubbieFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubbieFactory *ISubbieFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubbieFactory.Contract.ISubbieFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubbieFactory *ISubbieFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.ISubbieFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubbieFactory *ISubbieFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.ISubbieFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubbieFactory *ISubbieFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISubbieFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubbieFactory *ISubbieFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubbieFactory *ISubbieFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.contract.Transact(opts, method, params...)
}

// GetSubbie is a free data retrieval call binding the contract method 0x9790ea30.
//
// Solidity: function getSubbie(address _subbieRef) view returns(address, (uint8,uint8,string,string,string))
func (_ISubbieFactory *ISubbieFactoryCaller) GetSubbie(opts *bind.CallOpts, _subbieRef common.Address) (common.Address, Subbie, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(Subbie)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ISubbieFactory.contract.Call(opts, out, "getSubbie", _subbieRef)
	return *ret0, *ret1, err
}

// GetSubbie is a free data retrieval call binding the contract method 0x9790ea30.
//
// Solidity: function getSubbie(address _subbieRef) view returns(address, (uint8,uint8,string,string,string))
func (_ISubbieFactory *ISubbieFactorySession) GetSubbie(_subbieRef common.Address) (common.Address, Subbie, error) {
	return _ISubbieFactory.Contract.GetSubbie(&_ISubbieFactory.CallOpts, _subbieRef)
}

// GetSubbie is a free data retrieval call binding the contract method 0x9790ea30.
//
// Solidity: function getSubbie(address _subbieRef) view returns(address, (uint8,uint8,string,string,string))
func (_ISubbieFactory *ISubbieFactoryCallerSession) GetSubbie(_subbieRef common.Address) (common.Address, Subbie, error) {
	return _ISubbieFactory.Contract.GetSubbie(&_ISubbieFactory.CallOpts, _subbieRef)
}

// IsType is a free data retrieval call binding the contract method 0x6359e6a3.
//
// Solidity: function isType(address _subbieRef, uint8 _type) view returns(bool)
func (_ISubbieFactory *ISubbieFactoryCaller) IsType(opts *bind.CallOpts, _subbieRef common.Address, _type uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ISubbieFactory.contract.Call(opts, out, "isType", _subbieRef, _type)
	return *ret0, err
}

// IsType is a free data retrieval call binding the contract method 0x6359e6a3.
//
// Solidity: function isType(address _subbieRef, uint8 _type) view returns(bool)
func (_ISubbieFactory *ISubbieFactorySession) IsType(_subbieRef common.Address, _type uint8) (bool, error) {
	return _ISubbieFactory.Contract.IsType(&_ISubbieFactory.CallOpts, _subbieRef, _type)
}

// IsType is a free data retrieval call binding the contract method 0x6359e6a3.
//
// Solidity: function isType(address _subbieRef, uint8 _type) view returns(bool)
func (_ISubbieFactory *ISubbieFactoryCallerSession) IsType(_subbieRef common.Address, _type uint8) (bool, error) {
	return _ISubbieFactory.Contract.IsType(&_ISubbieFactory.CallOpts, _subbieRef, _type)
}

// AddSubbie is a paid mutator transaction binding the contract method 0xe28afe96.
//
// Solidity: function addSubbie(address _subbieRef, (uint8,uint8,string,string,string) _subbie) returns()
func (_ISubbieFactory *ISubbieFactoryTransactor) AddSubbie(opts *bind.TransactOpts, _subbieRef common.Address, _subbie Subbie) (*types.Transaction, error) {
	return _ISubbieFactory.contract.Transact(opts, "addSubbie", _subbieRef, _subbie)
}

// AddSubbie is a paid mutator transaction binding the contract method 0xe28afe96.
//
// Solidity: function addSubbie(address _subbieRef, (uint8,uint8,string,string,string) _subbie) returns()
func (_ISubbieFactory *ISubbieFactorySession) AddSubbie(_subbieRef common.Address, _subbie Subbie) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.AddSubbie(&_ISubbieFactory.TransactOpts, _subbieRef, _subbie)
}

// AddSubbie is a paid mutator transaction binding the contract method 0xe28afe96.
//
// Solidity: function addSubbie(address _subbieRef, (uint8,uint8,string,string,string) _subbie) returns()
func (_ISubbieFactory *ISubbieFactoryTransactorSession) AddSubbie(_subbieRef common.Address, _subbie Subbie) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.AddSubbie(&_ISubbieFactory.TransactOpts, _subbieRef, _subbie)
}

// AddType is a paid mutator transaction binding the contract method 0x90991c91.
//
// Solidity: function addType(address _subbieRef, uint8 _type) returns()
func (_ISubbieFactory *ISubbieFactoryTransactor) AddType(opts *bind.TransactOpts, _subbieRef common.Address, _type uint8) (*types.Transaction, error) {
	return _ISubbieFactory.contract.Transact(opts, "addType", _subbieRef, _type)
}

// AddType is a paid mutator transaction binding the contract method 0x90991c91.
//
// Solidity: function addType(address _subbieRef, uint8 _type) returns()
func (_ISubbieFactory *ISubbieFactorySession) AddType(_subbieRef common.Address, _type uint8) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.AddType(&_ISubbieFactory.TransactOpts, _subbieRef, _type)
}

// AddType is a paid mutator transaction binding the contract method 0x90991c91.
//
// Solidity: function addType(address _subbieRef, uint8 _type) returns()
func (_ISubbieFactory *ISubbieFactoryTransactorSession) AddType(_subbieRef common.Address, _type uint8) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.AddType(&_ISubbieFactory.TransactOpts, _subbieRef, _type)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _subbieRef, uint8 _status) returns()
func (_ISubbieFactory *ISubbieFactoryTransactor) SetStatus(opts *bind.TransactOpts, _subbieRef common.Address, _status uint8) (*types.Transaction, error) {
	return _ISubbieFactory.contract.Transact(opts, "setStatus", _subbieRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _subbieRef, uint8 _status) returns()
func (_ISubbieFactory *ISubbieFactorySession) SetStatus(_subbieRef common.Address, _status uint8) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.SetStatus(&_ISubbieFactory.TransactOpts, _subbieRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _subbieRef, uint8 _status) returns()
func (_ISubbieFactory *ISubbieFactoryTransactorSession) SetStatus(_subbieRef common.Address, _status uint8) (*types.Transaction, error) {
	return _ISubbieFactory.Contract.SetStatus(&_ISubbieFactory.TransactOpts, _subbieRef, _status)
}

// IterableDataABI is the input ABI used to generate the binding from.
const IterableDataABI = "[]"

// IterableDataBin is the compiled bytecode used for deploying new contracts.
var IterableDataBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205e295e30bfdaef3f7f2dfc0a760aba8bc4713a34baf1a8da337d12da990da76664736f6c63430006080033"

// DeployIterableData deploys a new Ethereum contract, binding an instance of IterableData to it.
func DeployIterableData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IterableData, error) {
	parsed, err := abi.JSON(strings.NewReader(IterableDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IterableDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IterableData{IterableDataCaller: IterableDataCaller{contract: contract}, IterableDataTransactor: IterableDataTransactor{contract: contract}, IterableDataFilterer: IterableDataFilterer{contract: contract}}, nil
}

// IterableData is an auto generated Go binding around an Ethereum contract.
type IterableData struct {
	IterableDataCaller     // Read-only binding to the contract
	IterableDataTransactor // Write-only binding to the contract
	IterableDataFilterer   // Log filterer for contract events
}

// IterableDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IterableDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IterableDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IterableDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IterableDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IterableDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IterableDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IterableDataSession struct {
	Contract     *IterableData     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IterableDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IterableDataCallerSession struct {
	Contract *IterableDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IterableDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IterableDataTransactorSession struct {
	Contract     *IterableDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IterableDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IterableDataRaw struct {
	Contract *IterableData // Generic contract binding to access the raw methods on
}

// IterableDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IterableDataCallerRaw struct {
	Contract *IterableDataCaller // Generic read-only contract binding to access the raw methods on
}

// IterableDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IterableDataTransactorRaw struct {
	Contract *IterableDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIterableData creates a new instance of IterableData, bound to a specific deployed contract.
func NewIterableData(address common.Address, backend bind.ContractBackend) (*IterableData, error) {
	contract, err := bindIterableData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IterableData{IterableDataCaller: IterableDataCaller{contract: contract}, IterableDataTransactor: IterableDataTransactor{contract: contract}, IterableDataFilterer: IterableDataFilterer{contract: contract}}, nil
}

// NewIterableDataCaller creates a new read-only instance of IterableData, bound to a specific deployed contract.
func NewIterableDataCaller(address common.Address, caller bind.ContractCaller) (*IterableDataCaller, error) {
	contract, err := bindIterableData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IterableDataCaller{contract: contract}, nil
}

// NewIterableDataTransactor creates a new write-only instance of IterableData, bound to a specific deployed contract.
func NewIterableDataTransactor(address common.Address, transactor bind.ContractTransactor) (*IterableDataTransactor, error) {
	contract, err := bindIterableData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IterableDataTransactor{contract: contract}, nil
}

// NewIterableDataFilterer creates a new log filterer instance of IterableData, bound to a specific deployed contract.
func NewIterableDataFilterer(address common.Address, filterer bind.ContractFilterer) (*IterableDataFilterer, error) {
	contract, err := bindIterableData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IterableDataFilterer{contract: contract}, nil
}

// bindIterableData binds a generic wrapper to an already deployed contract.
func bindIterableData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IterableDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IterableData *IterableDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IterableData.Contract.IterableDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IterableData *IterableDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IterableData.Contract.IterableDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IterableData *IterableDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IterableData.Contract.IterableDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IterableData *IterableDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IterableData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IterableData *IterableDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IterableData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IterableData *IterableDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IterableData.Contract.contract.Transact(opts, method, params...)
}

// SubbieNodeABI is the input ABI used to generate the binding from.
const SubbieNodeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeRef\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"mainType\",\"type\":\"uint8\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structSubbie\",\"name\":\"_node\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"mainType\",\"type\":\"uint8\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structSubbie\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSubbieStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SubbieNodeFuncSigs maps the 4-byte function signature to its string representation.
var SubbieNodeFuncSigs = map[string]string{
	"e188130b": "addType(uint8)",
	"6d4ce63c": "get()",
	"222f74e7": "isType(uint8)",
	"2e49d78b": "setStatus(uint8)",
}

// SubbieNodeBin is the compiled bytecode used for deploying new contracts.
var SubbieNodeBin = "0x60806040523480156200001157600080fd5b5060405162000a7638038062000a768339810160408190526200003491620002ed565b6001600160a01b038216158015906200005a575060008151604a8111156200005857fe5b115b8015620000745750604a8151604a8111156200007257fe5b105b80156200009157506000816020015160048111156200008f57fe5b115b8015620000ae5750600481602001516004811115620000ac57fe5b105b8015620000c057506000816040015151115b8015620000d257506000816060015151115b620000dc57600080fd5b600080546001600160a01b0319166001600160a01b03841617905580516001805460ff19168183604a8111156200010f57fe5b021790555060208101516001805461ff0019166101008360048111156200013257fe5b021790555060408101518051620001529160029160209091019062000193565b50606081015180516200016e9160039160209091019062000193565b50608081015180516200018a9160049160209091019062000193565b5050506200041d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001d657805160ff191683800117855562000206565b8280016001018555821562000206579182015b8281111562000206578251825591602001919060010190620001e9565b506200021492915062000218565b5090565b6200023591905b808211156200021457600081556001016200021f565b90565b8051604b81106200024857600080fd5b92915050565b8051600581106200024857600080fd5b600082601f8301126200026f578081fd5b81516001600160401b0381111562000285578182fd5b60206200029b601f8301601f19168201620003f6565b92508183528481838601011115620002b257600080fd5b60005b82811015620002d2578481018201518482018301528101620002b5565b82811115620002e45760008284860101525b50505092915050565b6000806040838503121562000300578182fd5b82516001600160a01b038116811462000317578283fd5b60208401519092506001600160401b038082111562000334578283fd5b81850160a0818803121562000347578384fd5b6200035360a0620003f6565b925062000361878262000238565b83526200037287602083016200024e565b602084015260408101518281111562000389578485fd5b62000397888284016200025e565b604085015250606081015182811115620003af578485fd5b620003bd888284016200025e565b606085015250608081015182811115620003d5578485fd5b620003e3888284016200025e565b6080850152505050809150509250929050565b6040518181016001600160401b03811182821017156200041557600080fd5b604052919050565b610649806200042d6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063222f74e7146100515780632e49d78b1461007a5780636d4ce63c1461008f578063e188130b146100a5575b600080fd5b61006461005f3660046104d2565b6100b8565b6040516100719190610608565b60405180910390f35b61008d6100883660046104f8565b610185565b005b6100976101d7565b604051610071929190610570565b61008d6100b33660046104d2565b61041d565b60008082604a8111156100c757fe5b1180156100df5750604a82604a8111156100dd57fe5b105b6100e857600080fd5b600082604a8111156100f657fe5b60015460ff16604a81111561010757fe5b14156101155750600161017f565b60005b60055481101561017d5783604a81111561012e57fe5b6005828154811061013b57fe5b90600052602060002090602091828204019190069054906101000a900460ff16604a81111561016657fe5b1415610175576001915061017d565b600101610118565b505b92915050565b600081600481111561019357fe5b1180156101ab575060048160048111156101a957fe5b105b6101b457600080fd5b6001805482919061ff0019166101008360048111156101cf57fe5b021790555050565b60006101e16104a2565b6000546040805160a08101909152600180546001600160a01b03909316929091908290829060ff16604a81111561021457fe5b604a81111561021f57fe5b81528154602090910190610100900460ff16600481111561023c57fe5b600481111561024757fe5b8152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102e45780601f106102b9576101008083540402835291602001916102e4565b820191906000526020600020905b8154815290600101906020018083116102c757829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156103765780601f1061034b57610100808354040283529160200191610376565b820191906000526020600020905b81548152906001019060200180831161035957829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561040a5780601f106103df5761010080835404028352916020019161040a565b820191906000526020600020905b8154815290600101906020018083116103ed57829003601f168201915b5050505050815250509050915091509091565b600081604a81111561042b57fe5b1180156104435750604a81604a81111561044157fe5b105b61044c57600080fd5b600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0602082040180548392601f166101000a60ff8102199091169083604a8111156101cf57fe5b6040805160a081019091528060008152602001600081526020016060815260200160608152602001606081525090565b6000602082840312156104e3578081fd5b8135604b81106104f1578182fd5b9392505050565b600060208284031215610509578081fd5b8135600581106104f1578182fd5b6005811061052157fe5b9052565b60008151808452815b8181101561054a5760208185018101518683018201520161052e565b8181111561055b5782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0383168152604060208201528151600090604b811061059257fe5b604083015260208301516105a96060840182610517565b50604083015160a060808401526105c360e0840182610525565b60608501519150603f19808583030160a08601526105e18284610525565b60808701519350818682030160c08701526105fc8185610525565b98975050505050505050565b90151581526020019056fea264697066735822122047ff0b867a2e4b5ec8015a005463bff69185b5e39b0e013b75834da9099925c064736f6c63430006080033"

// DeploySubbieNode deploys a new Ethereum contract, binding an instance of SubbieNode to it.
func DeploySubbieNode(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeRef common.Address, _node Subbie) (common.Address, *types.Transaction, *SubbieNode, error) {
	parsed, err := abi.JSON(strings.NewReader(SubbieNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SubbieNodeBin), backend, _nodeRef, _node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubbieNode{SubbieNodeCaller: SubbieNodeCaller{contract: contract}, SubbieNodeTransactor: SubbieNodeTransactor{contract: contract}, SubbieNodeFilterer: SubbieNodeFilterer{contract: contract}}, nil
}

// SubbieNode is an auto generated Go binding around an Ethereum contract.
type SubbieNode struct {
	SubbieNodeCaller     // Read-only binding to the contract
	SubbieNodeTransactor // Write-only binding to the contract
	SubbieNodeFilterer   // Log filterer for contract events
}

// SubbieNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubbieNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubbieNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubbieNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubbieNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubbieNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubbieNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubbieNodeSession struct {
	Contract     *SubbieNode       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubbieNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubbieNodeCallerSession struct {
	Contract *SubbieNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SubbieNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubbieNodeTransactorSession struct {
	Contract     *SubbieNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SubbieNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubbieNodeRaw struct {
	Contract *SubbieNode // Generic contract binding to access the raw methods on
}

// SubbieNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubbieNodeCallerRaw struct {
	Contract *SubbieNodeCaller // Generic read-only contract binding to access the raw methods on
}

// SubbieNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubbieNodeTransactorRaw struct {
	Contract *SubbieNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubbieNode creates a new instance of SubbieNode, bound to a specific deployed contract.
func NewSubbieNode(address common.Address, backend bind.ContractBackend) (*SubbieNode, error) {
	contract, err := bindSubbieNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubbieNode{SubbieNodeCaller: SubbieNodeCaller{contract: contract}, SubbieNodeTransactor: SubbieNodeTransactor{contract: contract}, SubbieNodeFilterer: SubbieNodeFilterer{contract: contract}}, nil
}

// NewSubbieNodeCaller creates a new read-only instance of SubbieNode, bound to a specific deployed contract.
func NewSubbieNodeCaller(address common.Address, caller bind.ContractCaller) (*SubbieNodeCaller, error) {
	contract, err := bindSubbieNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubbieNodeCaller{contract: contract}, nil
}

// NewSubbieNodeTransactor creates a new write-only instance of SubbieNode, bound to a specific deployed contract.
func NewSubbieNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*SubbieNodeTransactor, error) {
	contract, err := bindSubbieNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubbieNodeTransactor{contract: contract}, nil
}

// NewSubbieNodeFilterer creates a new log filterer instance of SubbieNode, bound to a specific deployed contract.
func NewSubbieNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*SubbieNodeFilterer, error) {
	contract, err := bindSubbieNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubbieNodeFilterer{contract: contract}, nil
}

// bindSubbieNode binds a generic wrapper to an already deployed contract.
func bindSubbieNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubbieNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubbieNode *SubbieNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubbieNode.Contract.SubbieNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubbieNode *SubbieNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubbieNode.Contract.SubbieNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubbieNode *SubbieNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubbieNode.Contract.SubbieNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubbieNode *SubbieNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubbieNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubbieNode *SubbieNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubbieNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubbieNode *SubbieNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubbieNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(address, (uint8,uint8,string,string,string))
func (_SubbieNode *SubbieNodeCaller) Get(opts *bind.CallOpts) (common.Address, Subbie, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(Subbie)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SubbieNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(address, (uint8,uint8,string,string,string))
func (_SubbieNode *SubbieNodeSession) Get() (common.Address, Subbie, error) {
	return _SubbieNode.Contract.Get(&_SubbieNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(address, (uint8,uint8,string,string,string))
func (_SubbieNode *SubbieNodeCallerSession) Get() (common.Address, Subbie, error) {
	return _SubbieNode.Contract.Get(&_SubbieNode.CallOpts)
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_SubbieNode *SubbieNodeCaller) IsType(opts *bind.CallOpts, _type uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SubbieNode.contract.Call(opts, out, "isType", _type)
	return *ret0, err
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_SubbieNode *SubbieNodeSession) IsType(_type uint8) (bool, error) {
	return _SubbieNode.Contract.IsType(&_SubbieNode.CallOpts, _type)
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_SubbieNode *SubbieNodeCallerSession) IsType(_type uint8) (bool, error) {
	return _SubbieNode.Contract.IsType(&_SubbieNode.CallOpts, _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_SubbieNode *SubbieNodeTransactor) AddType(opts *bind.TransactOpts, _type uint8) (*types.Transaction, error) {
	return _SubbieNode.contract.Transact(opts, "addType", _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_SubbieNode *SubbieNodeSession) AddType(_type uint8) (*types.Transaction, error) {
	return _SubbieNode.Contract.AddType(&_SubbieNode.TransactOpts, _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_SubbieNode *SubbieNodeTransactorSession) AddType(_type uint8) (*types.Transaction, error) {
	return _SubbieNode.Contract.AddType(&_SubbieNode.TransactOpts, _type)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_SubbieNode *SubbieNodeTransactor) SetStatus(opts *bind.TransactOpts, _status uint8) (*types.Transaction, error) {
	return _SubbieNode.contract.Transact(opts, "setStatus", _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_SubbieNode *SubbieNodeSession) SetStatus(_status uint8) (*types.Transaction, error) {
	return _SubbieNode.Contract.SetStatus(&_SubbieNode.TransactOpts, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_SubbieNode *SubbieNodeTransactorSession) SetStatus(_status uint8) (*types.Transaction, error) {
	return _SubbieNode.Contract.SetStatus(&_SubbieNode.TransactOpts, _status)
}

// SubbiesABI is the input ABI used to generate the binding from.
const SubbiesABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"mainType\",\"type\":\"uint8\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structSubbie\",\"name\":\"_subbie\",\"type\":\"tuple\"}],\"name\":\"addSubbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"},{\"internalType\":\"enumContractorTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"}],\"name\":\"getSubbie\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumContractorTypes\",\"name\":\"mainType\",\"type\":\"uint8\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structSubbie\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"},{\"internalType\":\"enumContractorTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subbieRef\",\"type\":\"address\"},{\"internalType\":\"enumSubbieStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SubbiesFuncSigs maps the 4-byte function signature to its string representation.
var SubbiesFuncSigs = map[string]string{
	"e28afe96": "addSubbie(address,(uint8,uint8,string,string,string))",
	"90991c91": "addType(address,uint8)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"9790ea30": "getSubbie(address)",
	"6359e6a3": "isType(address,uint8)",
	"278e07ce": "setStatus(address,uint8)",
}

// SubbiesBin is the compiled bytecode used for deploying new contracts.
var SubbiesBin = "0x608060405234801561001057600080fd5b506000600255611606806100256000396000f3fe60806040523480156200001157600080fd5b5060043610620000885760003560e01c806390991c91116200006357806390991c9114620000ee5780639790ea301462000105578063bc599341146200012c578063e28afe9614620001435762000088565b8063278e07ce146200008d5780636359e6a314620000a657806367e0badb14620000d5575b600080fd5b620000a46200009e36600462000739565b6200015a565b005b620000bd620000b7366004620006fc565b62000218565b604051620000cc919062000a70565b60405180910390f35b620000df620002fb565b604051620000cc919062000a7b565b620000a4620000ff366004620006fc565b62000301565b6200011c62000116366004620006dd565b62000386565b604051620000cc929190620009cf565b620000df6200013d36600462000979565b6200046d565b620000a4620001543660046200076b565b620004a7565b6001600160a01b038216158015906200018f57506001600160a01b038281166000908152602081905260409020600101541615155b6200019957600080fd5b6001600160a01b0380831660009081526020819052604090819020600101549051632e49d78b60e01b81529116908190632e49d78b90620001df90859060040162000a99565b600060405180830381600087803b158015620001fa57600080fd5b505af11580156200020f573d6000803e3d6000fd5b50505050505050565b60006001600160a01b038316158015906200024f57506001600160a01b038381166000908152602081905260409020600101541615155b6200025957600080fd5b6001600160a01b038084166000908152602081905260409081902060010154905163222f74e760e01b8152911690819063222f74e7906200029f90869060040162000a84565b60206040518083038186803b158015620002b857600080fd5b505afa158015620002cd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002f3919062000957565b505092915050565b60025490565b6001600160a01b038216158015906200033657506001600160a01b038281166000908152602081905260409020600101541615155b6200034057600080fd5b6001600160a01b038083166000908152602081905260409081902060010154905163e188130b60e01b8152911690819063e188130b90620001df90859060040162000a84565b600062000392620005b8565b6001600160a01b03831615801590620003c757506001600160a01b038381166000908152602081905260409020600101541615155b620003d157600080fd5b6001600160a01b03808416600090815260208190526040808220600101548151631b53398f60e21b815291519316928392636d4ce63c9260048082019391829003018186803b1580156200042457600080fd5b505afa15801562000439573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526200046391908101906200086a565b9250925050915091565b60025460009082106200047f57600080fd5b60018054839081106200048e57fe5b9060005260206000209060020201600001549050919050565b6001600160a01b038216620004bb57600080fd5b60008282604051620004cd90620005e8565b620004da929190620009cf565b604051809103906000f080158015620004f7573d6000803e3d6000fd5b5090506200051760006001600160a01b0385168363ffffffff6200051d16565b50505050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b03851617905580156200055d576001915050620005b1565b50600180850180549182018082556000868152602088905260409020819055859190839081106200058a57fe5b60009182526020822060029182020192909255908601805460010190559150620005b19050565b9392505050565b6040805160a081019091528060008152602001600081526020016060815260200160608152602001606081525090565b610a768062000b5b83390190565b8035620006038162000b3e565b92915050565b8051620006038162000b3e565b8035620006038162000b4c565b8051620006038162000b4c565b600082601f83011262000641578081fd5b813562000658620006528262000ad1565b62000aa9565b91508082528360208285010111156200067057600080fd5b8060208401602084013760009082016020015292915050565b600082601f8301126200069a578081fd5b8151620006ab620006528262000ad1565b9150808252836020828501011115620006c357600080fd5b620006d681602084016020860162000af6565b5092915050565b600060208284031215620006ef578081fd5b8135620005b18162000b25565b600080604083850312156200070f578081fd5b82356200071c8162000b25565b915060208301356200072e8162000b3e565b809150509250929050565b600080604083850312156200074c578182fd5b8235620007598162000b25565b915060208301356200072e8162000b4c565b600080604083850312156200077e578182fd5b82356200078b8162000b25565b9150602083013567ffffffffffffffff80821115620007a8578283fd5b81850160a08188031215620007bb578384fd5b620007c760a062000aa9565b9250620007d58782620005f6565b8352620007e6876020830162000616565b6020840152604081013582811115620007fd578485fd5b6200080b8882840162000630565b60408501525060608101358281111562000823578485fd5b620008318882840162000630565b60608501525060808101358281111562000849578485fd5b620008578882840162000630565b6080850152505050809150509250929050565b600080604083850312156200087d578182fd5b82516200088a8162000b25565b602084015190925067ffffffffffffffff80821115620008a8578283fd5b81850160a08188031215620008bb578384fd5b620008c760a062000aa9565b9250620008d5878262000609565b8352620008e6876020830162000623565b6020840152604081015182811115620008fd578485fd5b6200090b8882840162000689565b60408501525060608101518281111562000923578485fd5b620009318882840162000689565b60608501525060808101518281111562000949578485fd5b620008578882840162000689565b60006020828403121562000969578081fd5b81518015158114620005b1578182fd5b6000602082840312156200098b578081fd5b5035919050565b600581106200099d57fe5b9052565b60008151808452620009bb81602086016020860162000af6565b601f01601f19169290920160200192915050565b6001600160a01b0383168152604060208201528151600090604b8110620009f257fe5b6040830152602083015162000a0b606084018262000992565b50604083015160a0608084015262000a2760e0840182620009a1565b60608501519150603f19808583030160a086015262000a478284620009a1565b60808701519350818682030160c087015262000a648185620009a1565b98975050505050505050565b901515815260200190565b90815260200190565b60208101604b831062000a9357fe5b91905290565b6020810162000603828462000992565b60405181810167ffffffffffffffff8111828210171562000ac957600080fd5b604052919050565b600067ffffffffffffffff82111562000ae8578081fd5b50601f01601f191660200190565b60005b8381101562000b1357818101518382015260200162000af9565b83811115620005175750506000910152565b6001600160a01b038116811462000b3b57600080fd5b50565b604b811062000b3b57600080fd5b6005811062000b3b57600080fdfe60806040523480156200001157600080fd5b5060405162000a7638038062000a768339810160408190526200003491620002ed565b6001600160a01b038216158015906200005a575060008151604a8111156200005857fe5b115b8015620000745750604a8151604a8111156200007257fe5b105b80156200009157506000816020015160048111156200008f57fe5b115b8015620000ae5750600481602001516004811115620000ac57fe5b105b8015620000c057506000816040015151115b8015620000d257506000816060015151115b620000dc57600080fd5b600080546001600160a01b0319166001600160a01b03841617905580516001805460ff19168183604a8111156200010f57fe5b021790555060208101516001805461ff0019166101008360048111156200013257fe5b021790555060408101518051620001529160029160209091019062000193565b50606081015180516200016e9160039160209091019062000193565b50608081015180516200018a9160049160209091019062000193565b5050506200041d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001d657805160ff191683800117855562000206565b8280016001018555821562000206579182015b8281111562000206578251825591602001919060010190620001e9565b506200021492915062000218565b5090565b6200023591905b808211156200021457600081556001016200021f565b90565b8051604b81106200024857600080fd5b92915050565b8051600581106200024857600080fd5b600082601f8301126200026f578081fd5b81516001600160401b0381111562000285578182fd5b60206200029b601f8301601f19168201620003f6565b92508183528481838601011115620002b257600080fd5b60005b82811015620002d2578481018201518482018301528101620002b5565b82811115620002e45760008284860101525b50505092915050565b6000806040838503121562000300578182fd5b82516001600160a01b038116811462000317578283fd5b60208401519092506001600160401b038082111562000334578283fd5b81850160a0818803121562000347578384fd5b6200035360a0620003f6565b925062000361878262000238565b83526200037287602083016200024e565b602084015260408101518281111562000389578485fd5b62000397888284016200025e565b604085015250606081015182811115620003af578485fd5b620003bd888284016200025e565b606085015250608081015182811115620003d5578485fd5b620003e3888284016200025e565b6080850152505050809150509250929050565b6040518181016001600160401b03811182821017156200041557600080fd5b604052919050565b610649806200042d6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063222f74e7146100515780632e49d78b1461007a5780636d4ce63c1461008f578063e188130b146100a5575b600080fd5b61006461005f3660046104d2565b6100b8565b6040516100719190610608565b60405180910390f35b61008d6100883660046104f8565b610185565b005b6100976101d7565b604051610071929190610570565b61008d6100b33660046104d2565b61041d565b60008082604a8111156100c757fe5b1180156100df5750604a82604a8111156100dd57fe5b105b6100e857600080fd5b600082604a8111156100f657fe5b60015460ff16604a81111561010757fe5b14156101155750600161017f565b60005b60055481101561017d5783604a81111561012e57fe5b6005828154811061013b57fe5b90600052602060002090602091828204019190069054906101000a900460ff16604a81111561016657fe5b1415610175576001915061017d565b600101610118565b505b92915050565b600081600481111561019357fe5b1180156101ab575060048160048111156101a957fe5b105b6101b457600080fd5b6001805482919061ff0019166101008360048111156101cf57fe5b021790555050565b60006101e16104a2565b6000546040805160a08101909152600180546001600160a01b03909316929091908290829060ff16604a81111561021457fe5b604a81111561021f57fe5b81528154602090910190610100900460ff16600481111561023c57fe5b600481111561024757fe5b8152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102e45780601f106102b9576101008083540402835291602001916102e4565b820191906000526020600020905b8154815290600101906020018083116102c757829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156103765780601f1061034b57610100808354040283529160200191610376565b820191906000526020600020905b81548152906001019060200180831161035957829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561040a5780601f106103df5761010080835404028352916020019161040a565b820191906000526020600020905b8154815290600101906020018083116103ed57829003601f168201915b5050505050815250509050915091509091565b600081604a81111561042b57fe5b1180156104435750604a81604a81111561044157fe5b105b61044c57600080fd5b600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0602082040180548392601f166101000a60ff8102199091169083604a8111156101cf57fe5b6040805160a081019091528060008152602001600081526020016060815260200160608152602001606081525090565b6000602082840312156104e3578081fd5b8135604b81106104f1578182fd5b9392505050565b600060208284031215610509578081fd5b8135600581106104f1578182fd5b6005811061052157fe5b9052565b60008151808452815b8181101561054a5760208185018101518683018201520161052e565b8181111561055b5782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0383168152604060208201528151600090604b811061059257fe5b604083015260208301516105a96060840182610517565b50604083015160a060808401526105c360e0840182610525565b60608501519150603f19808583030160a08601526105e18284610525565b60808701519350818682030160c08701526105fc8185610525565b98975050505050505050565b90151581526020019056fea264697066735822122047ff0b867a2e4b5ec8015a005463bff69185b5e39b0e013b75834da9099925c064736f6c63430006080033a2646970667358221220a991df4da6e54bcbdf2ae960a09d66ddfb67410013691278f01761d31316cfcf64736f6c63430006080033"

// DeploySubbies deploys a new Ethereum contract, binding an instance of Subbies to it.
func DeploySubbies(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Subbies, error) {
	parsed, err := abi.JSON(strings.NewReader(SubbiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SubbiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Subbies{SubbiesCaller: SubbiesCaller{contract: contract}, SubbiesTransactor: SubbiesTransactor{contract: contract}, SubbiesFilterer: SubbiesFilterer{contract: contract}}, nil
}

// Subbies is an auto generated Go binding around an Ethereum contract.
type Subbies struct {
	SubbiesCaller     // Read-only binding to the contract
	SubbiesTransactor // Write-only binding to the contract
	SubbiesFilterer   // Log filterer for contract events
}

// SubbiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubbiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubbiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubbiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubbiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubbiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubbiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubbiesSession struct {
	Contract     *Subbies          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubbiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubbiesCallerSession struct {
	Contract *SubbiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SubbiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubbiesTransactorSession struct {
	Contract     *SubbiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SubbiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubbiesRaw struct {
	Contract *Subbies // Generic contract binding to access the raw methods on
}

// SubbiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubbiesCallerRaw struct {
	Contract *SubbiesCaller // Generic read-only contract binding to access the raw methods on
}

// SubbiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubbiesTransactorRaw struct {
	Contract *SubbiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubbies creates a new instance of Subbies, bound to a specific deployed contract.
func NewSubbies(address common.Address, backend bind.ContractBackend) (*Subbies, error) {
	contract, err := bindSubbies(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subbies{SubbiesCaller: SubbiesCaller{contract: contract}, SubbiesTransactor: SubbiesTransactor{contract: contract}, SubbiesFilterer: SubbiesFilterer{contract: contract}}, nil
}

// NewSubbiesCaller creates a new read-only instance of Subbies, bound to a specific deployed contract.
func NewSubbiesCaller(address common.Address, caller bind.ContractCaller) (*SubbiesCaller, error) {
	contract, err := bindSubbies(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubbiesCaller{contract: contract}, nil
}

// NewSubbiesTransactor creates a new write-only instance of Subbies, bound to a specific deployed contract.
func NewSubbiesTransactor(address common.Address, transactor bind.ContractTransactor) (*SubbiesTransactor, error) {
	contract, err := bindSubbies(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubbiesTransactor{contract: contract}, nil
}

// NewSubbiesFilterer creates a new log filterer instance of Subbies, bound to a specific deployed contract.
func NewSubbiesFilterer(address common.Address, filterer bind.ContractFilterer) (*SubbiesFilterer, error) {
	contract, err := bindSubbies(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubbiesFilterer{contract: contract}, nil
}

// bindSubbies binds a generic wrapper to an already deployed contract.
func bindSubbies(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubbiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subbies *SubbiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Subbies.Contract.SubbiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subbies *SubbiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subbies.Contract.SubbiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subbies *SubbiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subbies.Contract.SubbiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subbies *SubbiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Subbies.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subbies *SubbiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subbies.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subbies *SubbiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subbies.Contract.contract.Transact(opts, method, params...)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Subbies *SubbiesCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subbies.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Subbies *SubbiesSession) GetNum() (*big.Int, error) {
	return _Subbies.Contract.GetNum(&_Subbies.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Subbies *SubbiesCallerSession) GetNum() (*big.Int, error) {
	return _Subbies.Contract.GetNum(&_Subbies.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Subbies *SubbiesCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Subbies.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Subbies *SubbiesSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Subbies.Contract.GetReference(&_Subbies.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Subbies *SubbiesCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Subbies.Contract.GetReference(&_Subbies.CallOpts, _index)
}

// GetSubbie is a free data retrieval call binding the contract method 0x9790ea30.
//
// Solidity: function getSubbie(address _subbieRef) view returns(address, (uint8,uint8,string,string,string))
func (_Subbies *SubbiesCaller) GetSubbie(opts *bind.CallOpts, _subbieRef common.Address) (common.Address, Subbie, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(Subbie)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Subbies.contract.Call(opts, out, "getSubbie", _subbieRef)
	return *ret0, *ret1, err
}

// GetSubbie is a free data retrieval call binding the contract method 0x9790ea30.
//
// Solidity: function getSubbie(address _subbieRef) view returns(address, (uint8,uint8,string,string,string))
func (_Subbies *SubbiesSession) GetSubbie(_subbieRef common.Address) (common.Address, Subbie, error) {
	return _Subbies.Contract.GetSubbie(&_Subbies.CallOpts, _subbieRef)
}

// GetSubbie is a free data retrieval call binding the contract method 0x9790ea30.
//
// Solidity: function getSubbie(address _subbieRef) view returns(address, (uint8,uint8,string,string,string))
func (_Subbies *SubbiesCallerSession) GetSubbie(_subbieRef common.Address) (common.Address, Subbie, error) {
	return _Subbies.Contract.GetSubbie(&_Subbies.CallOpts, _subbieRef)
}

// IsType is a free data retrieval call binding the contract method 0x6359e6a3.
//
// Solidity: function isType(address _subbieRef, uint8 _type) view returns(bool)
func (_Subbies *SubbiesCaller) IsType(opts *bind.CallOpts, _subbieRef common.Address, _type uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Subbies.contract.Call(opts, out, "isType", _subbieRef, _type)
	return *ret0, err
}

// IsType is a free data retrieval call binding the contract method 0x6359e6a3.
//
// Solidity: function isType(address _subbieRef, uint8 _type) view returns(bool)
func (_Subbies *SubbiesSession) IsType(_subbieRef common.Address, _type uint8) (bool, error) {
	return _Subbies.Contract.IsType(&_Subbies.CallOpts, _subbieRef, _type)
}

// IsType is a free data retrieval call binding the contract method 0x6359e6a3.
//
// Solidity: function isType(address _subbieRef, uint8 _type) view returns(bool)
func (_Subbies *SubbiesCallerSession) IsType(_subbieRef common.Address, _type uint8) (bool, error) {
	return _Subbies.Contract.IsType(&_Subbies.CallOpts, _subbieRef, _type)
}

// AddSubbie is a paid mutator transaction binding the contract method 0xe28afe96.
//
// Solidity: function addSubbie(address _subbieRef, (uint8,uint8,string,string,string) _subbie) returns()
func (_Subbies *SubbiesTransactor) AddSubbie(opts *bind.TransactOpts, _subbieRef common.Address, _subbie Subbie) (*types.Transaction, error) {
	return _Subbies.contract.Transact(opts, "addSubbie", _subbieRef, _subbie)
}

// AddSubbie is a paid mutator transaction binding the contract method 0xe28afe96.
//
// Solidity: function addSubbie(address _subbieRef, (uint8,uint8,string,string,string) _subbie) returns()
func (_Subbies *SubbiesSession) AddSubbie(_subbieRef common.Address, _subbie Subbie) (*types.Transaction, error) {
	return _Subbies.Contract.AddSubbie(&_Subbies.TransactOpts, _subbieRef, _subbie)
}

// AddSubbie is a paid mutator transaction binding the contract method 0xe28afe96.
//
// Solidity: function addSubbie(address _subbieRef, (uint8,uint8,string,string,string) _subbie) returns()
func (_Subbies *SubbiesTransactorSession) AddSubbie(_subbieRef common.Address, _subbie Subbie) (*types.Transaction, error) {
	return _Subbies.Contract.AddSubbie(&_Subbies.TransactOpts, _subbieRef, _subbie)
}

// AddType is a paid mutator transaction binding the contract method 0x90991c91.
//
// Solidity: function addType(address _subbieRef, uint8 _type) returns()
func (_Subbies *SubbiesTransactor) AddType(opts *bind.TransactOpts, _subbieRef common.Address, _type uint8) (*types.Transaction, error) {
	return _Subbies.contract.Transact(opts, "addType", _subbieRef, _type)
}

// AddType is a paid mutator transaction binding the contract method 0x90991c91.
//
// Solidity: function addType(address _subbieRef, uint8 _type) returns()
func (_Subbies *SubbiesSession) AddType(_subbieRef common.Address, _type uint8) (*types.Transaction, error) {
	return _Subbies.Contract.AddType(&_Subbies.TransactOpts, _subbieRef, _type)
}

// AddType is a paid mutator transaction binding the contract method 0x90991c91.
//
// Solidity: function addType(address _subbieRef, uint8 _type) returns()
func (_Subbies *SubbiesTransactorSession) AddType(_subbieRef common.Address, _type uint8) (*types.Transaction, error) {
	return _Subbies.Contract.AddType(&_Subbies.TransactOpts, _subbieRef, _type)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _subbieRef, uint8 _status) returns()
func (_Subbies *SubbiesTransactor) SetStatus(opts *bind.TransactOpts, _subbieRef common.Address, _status uint8) (*types.Transaction, error) {
	return _Subbies.contract.Transact(opts, "setStatus", _subbieRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _subbieRef, uint8 _status) returns()
func (_Subbies *SubbiesSession) SetStatus(_subbieRef common.Address, _status uint8) (*types.Transaction, error) {
	return _Subbies.Contract.SetStatus(&_Subbies.TransactOpts, _subbieRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _subbieRef, uint8 _status) returns()
func (_Subbies *SubbiesTransactorSession) SetStatus(_subbieRef common.Address, _status uint8) (*types.Transaction, error) {
	return _Subbies.Contract.SetStatus(&_Subbies.TransactOpts, _subbieRef, _status)
}
