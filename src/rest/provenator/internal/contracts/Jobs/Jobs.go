// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Jobs

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

// ExtraInfo is an auto generated low-level Go binding around an user-defined struct.
type ExtraInfo struct {
	FileHash    string
	FileName    string
	FileUrl     string
	Description string
}

// Job is an auto generated low-level Go binding around an user-defined struct.
type Job struct {
	Status      uint8
	Location    string
	Description string
}

// Work is an auto generated low-level Go binding around an user-defined struct.
type Work struct {
	Status      uint8
	Subbie      common.Address
	StartDate   [32]byte
	EndDate     [32]byte
	StartTime   [32]byte
	EndTime     [32]byte
	SponsRef    [32]byte
	Quantity    *big.Int
	Cost        *big.Int
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

// IInfoABI is the input ABI used to generate the binding from.
const IInfoABI = "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IInfoFuncSigs maps the 4-byte function signature to its string representation.
var IInfoFuncSigs = map[string]string{
	"6d4ce63c": "get()",
}

// IInfo is an auto generated Go binding around an Ethereum contract.
type IInfo struct {
	IInfoCaller     // Read-only binding to the contract
	IInfoTransactor // Write-only binding to the contract
	IInfoFilterer   // Log filterer for contract events
}

// IInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInfoSession struct {
	Contract     *IInfo            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInfoCallerSession struct {
	Contract *IInfoCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInfoTransactorSession struct {
	Contract     *IInfoTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInfoRaw struct {
	Contract *IInfo // Generic contract binding to access the raw methods on
}

// IInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInfoCallerRaw struct {
	Contract *IInfoCaller // Generic read-only contract binding to access the raw methods on
}

// IInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInfoTransactorRaw struct {
	Contract *IInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInfo creates a new instance of IInfo, bound to a specific deployed contract.
func NewIInfo(address common.Address, backend bind.ContractBackend) (*IInfo, error) {
	contract, err := bindIInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInfo{IInfoCaller: IInfoCaller{contract: contract}, IInfoTransactor: IInfoTransactor{contract: contract}, IInfoFilterer: IInfoFilterer{contract: contract}}, nil
}

// NewIInfoCaller creates a new read-only instance of IInfo, bound to a specific deployed contract.
func NewIInfoCaller(address common.Address, caller bind.ContractCaller) (*IInfoCaller, error) {
	contract, err := bindIInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInfoCaller{contract: contract}, nil
}

// NewIInfoTransactor creates a new write-only instance of IInfo, bound to a specific deployed contract.
func NewIInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*IInfoTransactor, error) {
	contract, err := bindIInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInfoTransactor{contract: contract}, nil
}

// NewIInfoFilterer creates a new log filterer instance of IInfo, bound to a specific deployed contract.
func NewIInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*IInfoFilterer, error) {
	contract, err := bindIInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInfoFilterer{contract: contract}, nil
}

// bindIInfo binds a generic wrapper to an already deployed contract.
func bindIInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInfo *IInfoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IInfo.Contract.IInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInfo *IInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInfo.Contract.IInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInfo *IInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInfo.Contract.IInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInfo *IInfoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInfo *IInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInfo *IInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInfo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (string,string,string,string))
func (_IInfo *IInfoCaller) Get(opts *bind.CallOpts) ([32]byte, ExtraInfo, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(ExtraInfo)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IInfo.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (string,string,string,string))
func (_IInfo *IInfoSession) Get() ([32]byte, ExtraInfo, error) {
	return _IInfo.Contract.Get(&_IInfo.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (string,string,string,string))
func (_IInfo *IInfoCallerSession) Get() ([32]byte, ExtraInfo, error) {
	return _IInfo.Contract.Get(&_IInfo.CallOpts)
}

// IInfoFactoryABI is the input ABI used to generate the binding from.
const IInfoFactoryABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"addInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"}],\"name\":\"getInfoContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IInfoFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IInfoFactoryFuncSigs = map[string]string{
	"f538e358": "addInfo(bytes32,bytes32,(string,string,string,string))",
	"52c43d9f": "getInfoContract(bytes32,bytes32)",
}

// IInfoFactory is an auto generated Go binding around an Ethereum contract.
type IInfoFactory struct {
	IInfoFactoryCaller     // Read-only binding to the contract
	IInfoFactoryTransactor // Write-only binding to the contract
	IInfoFactoryFilterer   // Log filterer for contract events
}

// IInfoFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInfoFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInfoFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInfoFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInfoFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInfoFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInfoFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInfoFactorySession struct {
	Contract     *IInfoFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInfoFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInfoFactoryCallerSession struct {
	Contract *IInfoFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IInfoFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInfoFactoryTransactorSession struct {
	Contract     *IInfoFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IInfoFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInfoFactoryRaw struct {
	Contract *IInfoFactory // Generic contract binding to access the raw methods on
}

// IInfoFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInfoFactoryCallerRaw struct {
	Contract *IInfoFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IInfoFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInfoFactoryTransactorRaw struct {
	Contract *IInfoFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInfoFactory creates a new instance of IInfoFactory, bound to a specific deployed contract.
func NewIInfoFactory(address common.Address, backend bind.ContractBackend) (*IInfoFactory, error) {
	contract, err := bindIInfoFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInfoFactory{IInfoFactoryCaller: IInfoFactoryCaller{contract: contract}, IInfoFactoryTransactor: IInfoFactoryTransactor{contract: contract}, IInfoFactoryFilterer: IInfoFactoryFilterer{contract: contract}}, nil
}

// NewIInfoFactoryCaller creates a new read-only instance of IInfoFactory, bound to a specific deployed contract.
func NewIInfoFactoryCaller(address common.Address, caller bind.ContractCaller) (*IInfoFactoryCaller, error) {
	contract, err := bindIInfoFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInfoFactoryCaller{contract: contract}, nil
}

// NewIInfoFactoryTransactor creates a new write-only instance of IInfoFactory, bound to a specific deployed contract.
func NewIInfoFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IInfoFactoryTransactor, error) {
	contract, err := bindIInfoFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInfoFactoryTransactor{contract: contract}, nil
}

// NewIInfoFactoryFilterer creates a new log filterer instance of IInfoFactory, bound to a specific deployed contract.
func NewIInfoFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IInfoFactoryFilterer, error) {
	contract, err := bindIInfoFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInfoFactoryFilterer{contract: contract}, nil
}

// bindIInfoFactory binds a generic wrapper to an already deployed contract.
func bindIInfoFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInfoFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInfoFactory *IInfoFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IInfoFactory.Contract.IInfoFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInfoFactory *IInfoFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInfoFactory.Contract.IInfoFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInfoFactory *IInfoFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInfoFactory.Contract.IInfoFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInfoFactory *IInfoFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IInfoFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInfoFactory *IInfoFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInfoFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInfoFactory *IInfoFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInfoFactory.Contract.contract.Transact(opts, method, params...)
}

// GetInfoContract is a free data retrieval call binding the contract method 0x52c43d9f.
//
// Solidity: function getInfoContract(bytes32 _workRef, bytes32 _infoRef) view returns(address)
func (_IInfoFactory *IInfoFactoryCaller) GetInfoContract(opts *bind.CallOpts, _workRef [32]byte, _infoRef [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IInfoFactory.contract.Call(opts, out, "getInfoContract", _workRef, _infoRef)
	return *ret0, err
}

// GetInfoContract is a free data retrieval call binding the contract method 0x52c43d9f.
//
// Solidity: function getInfoContract(bytes32 _workRef, bytes32 _infoRef) view returns(address)
func (_IInfoFactory *IInfoFactorySession) GetInfoContract(_workRef [32]byte, _infoRef [32]byte) (common.Address, error) {
	return _IInfoFactory.Contract.GetInfoContract(&_IInfoFactory.CallOpts, _workRef, _infoRef)
}

// GetInfoContract is a free data retrieval call binding the contract method 0x52c43d9f.
//
// Solidity: function getInfoContract(bytes32 _workRef, bytes32 _infoRef) view returns(address)
func (_IInfoFactory *IInfoFactoryCallerSession) GetInfoContract(_workRef [32]byte, _infoRef [32]byte) (common.Address, error) {
	return _IInfoFactory.Contract.GetInfoContract(&_IInfoFactory.CallOpts, _workRef, _infoRef)
}

// AddInfo is a paid mutator transaction binding the contract method 0xf538e358.
//
// Solidity: function addInfo(bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IInfoFactory *IInfoFactoryTransactor) AddInfo(opts *bind.TransactOpts, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IInfoFactory.contract.Transact(opts, "addInfo", _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xf538e358.
//
// Solidity: function addInfo(bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IInfoFactory *IInfoFactorySession) AddInfo(_workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IInfoFactory.Contract.AddInfo(&_IInfoFactory.TransactOpts, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xf538e358.
//
// Solidity: function addInfo(bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IInfoFactory *IInfoFactoryTransactorSession) AddInfo(_workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IInfoFactory.Contract.AddInfo(&_IInfoFactory.TransactOpts, _workRef, _infoRef, _info)
}

// IJobABI is the input ABI used to generate the binding from.
const IJobABI = "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IJobFuncSigs maps the 4-byte function signature to its string representation.
var IJobFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"8de654ba": "setStatus(bytes32,uint8)",
}

// IJob is an auto generated Go binding around an Ethereum contract.
type IJob struct {
	IJobCaller     // Read-only binding to the contract
	IJobTransactor // Write-only binding to the contract
	IJobFilterer   // Log filterer for contract events
}

// IJobCaller is an auto generated read-only Go binding around an Ethereum contract.
type IJobCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IJobTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IJobFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IJobSession struct {
	Contract     *IJob             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJobCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IJobCallerSession struct {
	Contract *IJobCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IJobTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IJobTransactorSession struct {
	Contract     *IJobTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJobRaw is an auto generated low-level Go binding around an Ethereum contract.
type IJobRaw struct {
	Contract *IJob // Generic contract binding to access the raw methods on
}

// IJobCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IJobCallerRaw struct {
	Contract *IJobCaller // Generic read-only contract binding to access the raw methods on
}

// IJobTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IJobTransactorRaw struct {
	Contract *IJobTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIJob creates a new instance of IJob, bound to a specific deployed contract.
func NewIJob(address common.Address, backend bind.ContractBackend) (*IJob, error) {
	contract, err := bindIJob(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IJob{IJobCaller: IJobCaller{contract: contract}, IJobTransactor: IJobTransactor{contract: contract}, IJobFilterer: IJobFilterer{contract: contract}}, nil
}

// NewIJobCaller creates a new read-only instance of IJob, bound to a specific deployed contract.
func NewIJobCaller(address common.Address, caller bind.ContractCaller) (*IJobCaller, error) {
	contract, err := bindIJob(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IJobCaller{contract: contract}, nil
}

// NewIJobTransactor creates a new write-only instance of IJob, bound to a specific deployed contract.
func NewIJobTransactor(address common.Address, transactor bind.ContractTransactor) (*IJobTransactor, error) {
	contract, err := bindIJob(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IJobTransactor{contract: contract}, nil
}

// NewIJobFilterer creates a new log filterer instance of IJob, bound to a specific deployed contract.
func NewIJobFilterer(address common.Address, filterer bind.ContractFilterer) (*IJobFilterer, error) {
	contract, err := bindIJob(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IJobFilterer{contract: contract}, nil
}

// bindIJob binds a generic wrapper to an already deployed contract.
func bindIJob(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IJobABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJob *IJobRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IJob.Contract.IJobCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJob *IJobRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJob.Contract.IJobTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJob *IJobRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJob.Contract.IJobTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJob *IJobCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IJob.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJob *IJobTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJob.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJob *IJobTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJob.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,string,string))
func (_IJob *IJobCaller) Get(opts *bind.CallOpts) ([32]byte, Job, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(Job)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IJob.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,string,string))
func (_IJob *IJobSession) Get() ([32]byte, Job, error) {
	return _IJob.Contract.Get(&_IJob.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,string,string))
func (_IJob *IJobCallerSession) Get() ([32]byte, Job, error) {
	return _IJob.Contract.Get(&_IJob.CallOpts)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _jobRef, uint8 _status) returns()
func (_IJob *IJobTransactor) SetStatus(opts *bind.TransactOpts, _jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJob.contract.Transact(opts, "setStatus", _jobRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _jobRef, uint8 _status) returns()
func (_IJob *IJobSession) SetStatus(_jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJob.Contract.SetStatus(&_IJob.TransactOpts, _jobRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _jobRef, uint8 _status) returns()
func (_IJob *IJobTransactorSession) SetStatus(_jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJob.Contract.SetStatus(&_IJob.TransactOpts, _jobRef, _status)
}

// IJobFactoryABI is the input ABI used to generate the binding from.
const IJobFactoryABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"addInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"_job\",\"type\":\"tuple\"}],\"name\":\"addJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumWorkStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"subbie\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"startDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sponsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWork\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"}],\"name\":\"getJobContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_subbie\",\"type\":\"address\"}],\"name\":\"setSubbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setWorkStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IJobFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IJobFactoryFuncSigs = map[string]string{
	"dcb5ad27": "addInfo(bytes32,bytes32,bytes32,(string,string,string,string))",
	"8b7a9f98": "addJob(bytes32,(uint8,string,string))",
	"ddca1462": "addWork(bytes32,bytes32,(uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))",
	"c918e294": "getJobContract(bytes32)",
	"e1908676": "setJobStatus(bytes32,uint8)",
	"8896dca8": "setSubbie(bytes32,bytes32,address)",
	"e2aaa459": "setWorkStatus(bytes32,bytes32,uint8)",
}

// IJobFactory is an auto generated Go binding around an Ethereum contract.
type IJobFactory struct {
	IJobFactoryCaller     // Read-only binding to the contract
	IJobFactoryTransactor // Write-only binding to the contract
	IJobFactoryFilterer   // Log filterer for contract events
}

// IJobFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IJobFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IJobFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IJobFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IJobFactorySession struct {
	Contract     *IJobFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJobFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IJobFactoryCallerSession struct {
	Contract *IJobFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IJobFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IJobFactoryTransactorSession struct {
	Contract     *IJobFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IJobFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IJobFactoryRaw struct {
	Contract *IJobFactory // Generic contract binding to access the raw methods on
}

// IJobFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IJobFactoryCallerRaw struct {
	Contract *IJobFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IJobFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IJobFactoryTransactorRaw struct {
	Contract *IJobFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIJobFactory creates a new instance of IJobFactory, bound to a specific deployed contract.
func NewIJobFactory(address common.Address, backend bind.ContractBackend) (*IJobFactory, error) {
	contract, err := bindIJobFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IJobFactory{IJobFactoryCaller: IJobFactoryCaller{contract: contract}, IJobFactoryTransactor: IJobFactoryTransactor{contract: contract}, IJobFactoryFilterer: IJobFactoryFilterer{contract: contract}}, nil
}

// NewIJobFactoryCaller creates a new read-only instance of IJobFactory, bound to a specific deployed contract.
func NewIJobFactoryCaller(address common.Address, caller bind.ContractCaller) (*IJobFactoryCaller, error) {
	contract, err := bindIJobFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IJobFactoryCaller{contract: contract}, nil
}

// NewIJobFactoryTransactor creates a new write-only instance of IJobFactory, bound to a specific deployed contract.
func NewIJobFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IJobFactoryTransactor, error) {
	contract, err := bindIJobFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IJobFactoryTransactor{contract: contract}, nil
}

// NewIJobFactoryFilterer creates a new log filterer instance of IJobFactory, bound to a specific deployed contract.
func NewIJobFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IJobFactoryFilterer, error) {
	contract, err := bindIJobFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IJobFactoryFilterer{contract: contract}, nil
}

// bindIJobFactory binds a generic wrapper to an already deployed contract.
func bindIJobFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IJobFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJobFactory *IJobFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IJobFactory.Contract.IJobFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJobFactory *IJobFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJobFactory.Contract.IJobFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJobFactory *IJobFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJobFactory.Contract.IJobFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJobFactory *IJobFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IJobFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJobFactory *IJobFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJobFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJobFactory *IJobFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJobFactory.Contract.contract.Transact(opts, method, params...)
}

// GetJobContract is a free data retrieval call binding the contract method 0xc918e294.
//
// Solidity: function getJobContract(bytes32 _jobRef) view returns(address)
func (_IJobFactory *IJobFactoryCaller) GetJobContract(opts *bind.CallOpts, _jobRef [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IJobFactory.contract.Call(opts, out, "getJobContract", _jobRef)
	return *ret0, err
}

// GetJobContract is a free data retrieval call binding the contract method 0xc918e294.
//
// Solidity: function getJobContract(bytes32 _jobRef) view returns(address)
func (_IJobFactory *IJobFactorySession) GetJobContract(_jobRef [32]byte) (common.Address, error) {
	return _IJobFactory.Contract.GetJobContract(&_IJobFactory.CallOpts, _jobRef)
}

// GetJobContract is a free data retrieval call binding the contract method 0xc918e294.
//
// Solidity: function getJobContract(bytes32 _jobRef) view returns(address)
func (_IJobFactory *IJobFactoryCallerSession) GetJobContract(_jobRef [32]byte) (common.Address, error) {
	return _IJobFactory.Contract.GetJobContract(&_IJobFactory.CallOpts, _jobRef)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IJobFactory *IJobFactoryTransactor) AddInfo(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IJobFactory.contract.Transact(opts, "addInfo", _jobRef, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IJobFactory *IJobFactorySession) AddInfo(_jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IJobFactory.Contract.AddInfo(&_IJobFactory.TransactOpts, _jobRef, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IJobFactory *IJobFactoryTransactorSession) AddInfo(_jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IJobFactory.Contract.AddInfo(&_IJobFactory.TransactOpts, _jobRef, _workRef, _infoRef, _info)
}

// AddJob is a paid mutator transaction binding the contract method 0x8b7a9f98.
//
// Solidity: function addJob(bytes32 _jobRef, (uint8,string,string) _job) returns()
func (_IJobFactory *IJobFactoryTransactor) AddJob(opts *bind.TransactOpts, _jobRef [32]byte, _job Job) (*types.Transaction, error) {
	return _IJobFactory.contract.Transact(opts, "addJob", _jobRef, _job)
}

// AddJob is a paid mutator transaction binding the contract method 0x8b7a9f98.
//
// Solidity: function addJob(bytes32 _jobRef, (uint8,string,string) _job) returns()
func (_IJobFactory *IJobFactorySession) AddJob(_jobRef [32]byte, _job Job) (*types.Transaction, error) {
	return _IJobFactory.Contract.AddJob(&_IJobFactory.TransactOpts, _jobRef, _job)
}

// AddJob is a paid mutator transaction binding the contract method 0x8b7a9f98.
//
// Solidity: function addJob(bytes32 _jobRef, (uint8,string,string) _job) returns()
func (_IJobFactory *IJobFactoryTransactorSession) AddJob(_jobRef [32]byte, _job Job) (*types.Transaction, error) {
	return _IJobFactory.Contract.AddJob(&_IJobFactory.TransactOpts, _jobRef, _job)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_IJobFactory *IJobFactoryTransactor) AddWork(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _IJobFactory.contract.Transact(opts, "addWork", _jobRef, _workRef, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_IJobFactory *IJobFactorySession) AddWork(_jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _IJobFactory.Contract.AddWork(&_IJobFactory.TransactOpts, _jobRef, _workRef, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_IJobFactory *IJobFactoryTransactorSession) AddWork(_jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _IJobFactory.Contract.AddWork(&_IJobFactory.TransactOpts, _jobRef, _workRef, _work)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobRef, uint8 _status) returns()
func (_IJobFactory *IJobFactoryTransactor) SetJobStatus(opts *bind.TransactOpts, _jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJobFactory.contract.Transact(opts, "setJobStatus", _jobRef, _status)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobRef, uint8 _status) returns()
func (_IJobFactory *IJobFactorySession) SetJobStatus(_jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJobFactory.Contract.SetJobStatus(&_IJobFactory.TransactOpts, _jobRef, _status)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobRef, uint8 _status) returns()
func (_IJobFactory *IJobFactoryTransactorSession) SetJobStatus(_jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJobFactory.Contract.SetJobStatus(&_IJobFactory.TransactOpts, _jobRef, _status)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_IJobFactory *IJobFactoryTransactor) SetSubbie(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IJobFactory.contract.Transact(opts, "setSubbie", _jobRef, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_IJobFactory *IJobFactorySession) SetSubbie(_jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IJobFactory.Contract.SetSubbie(&_IJobFactory.TransactOpts, _jobRef, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_IJobFactory *IJobFactoryTransactorSession) SetSubbie(_jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IJobFactory.Contract.SetSubbie(&_IJobFactory.TransactOpts, _jobRef, _workRef, _subbie)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_IJobFactory *IJobFactoryTransactor) SetWorkStatus(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJobFactory.contract.Transact(opts, "setWorkStatus", _jobRef, _workRef, _status)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_IJobFactory *IJobFactorySession) SetWorkStatus(_jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJobFactory.Contract.SetWorkStatus(&_IJobFactory.TransactOpts, _jobRef, _workRef, _status)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_IJobFactory *IJobFactoryTransactorSession) SetWorkStatus(_jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IJobFactory.Contract.SetWorkStatus(&_IJobFactory.TransactOpts, _jobRef, _workRef, _status)
}

// IWorkABI is the input ABI used to generate the binding from.
const IWorkABI = "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumWorkStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"subbie\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"startDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sponsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_subbie\",\"type\":\"address\"}],\"name\":\"setSubbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWorkFuncSigs maps the 4-byte function signature to its string representation.
var IWorkFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"8de654ba": "setStatus(bytes32,uint8)",
	"d00c66af": "setSubbie(bytes32,address)",
}

// IWork is an auto generated Go binding around an Ethereum contract.
type IWork struct {
	IWorkCaller     // Read-only binding to the contract
	IWorkTransactor // Write-only binding to the contract
	IWorkFilterer   // Log filterer for contract events
}

// IWorkCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWorkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWorkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWorkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWorkSession struct {
	Contract     *IWork            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWorkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWorkCallerSession struct {
	Contract *IWorkCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWorkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWorkTransactorSession struct {
	Contract     *IWorkTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWorkRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWorkRaw struct {
	Contract *IWork // Generic contract binding to access the raw methods on
}

// IWorkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWorkCallerRaw struct {
	Contract *IWorkCaller // Generic read-only contract binding to access the raw methods on
}

// IWorkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWorkTransactorRaw struct {
	Contract *IWorkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWork creates a new instance of IWork, bound to a specific deployed contract.
func NewIWork(address common.Address, backend bind.ContractBackend) (*IWork, error) {
	contract, err := bindIWork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWork{IWorkCaller: IWorkCaller{contract: contract}, IWorkTransactor: IWorkTransactor{contract: contract}, IWorkFilterer: IWorkFilterer{contract: contract}}, nil
}

// NewIWorkCaller creates a new read-only instance of IWork, bound to a specific deployed contract.
func NewIWorkCaller(address common.Address, caller bind.ContractCaller) (*IWorkCaller, error) {
	contract, err := bindIWork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkCaller{contract: contract}, nil
}

// NewIWorkTransactor creates a new write-only instance of IWork, bound to a specific deployed contract.
func NewIWorkTransactor(address common.Address, transactor bind.ContractTransactor) (*IWorkTransactor, error) {
	contract, err := bindIWork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkTransactor{contract: contract}, nil
}

// NewIWorkFilterer creates a new log filterer instance of IWork, bound to a specific deployed contract.
func NewIWorkFilterer(address common.Address, filterer bind.ContractFilterer) (*IWorkFilterer, error) {
	contract, err := bindIWork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWorkFilterer{contract: contract}, nil
}

// bindIWork binds a generic wrapper to an already deployed contract.
func bindIWork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWorkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWork *IWorkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IWork.Contract.IWorkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWork *IWorkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWork.Contract.IWorkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWork *IWorkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWork.Contract.IWorkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWork *IWorkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IWork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWork *IWorkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWork *IWorkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWork.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))
func (_IWork *IWorkCaller) Get(opts *bind.CallOpts) ([32]byte, Work, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(Work)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IWork.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))
func (_IWork *IWorkSession) Get() ([32]byte, Work, error) {
	return _IWork.Contract.Get(&_IWork.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))
func (_IWork *IWorkCallerSession) Get() ([32]byte, Work, error) {
	return _IWork.Contract.Get(&_IWork.CallOpts)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _workRef, uint8 _status) returns()
func (_IWork *IWorkTransactor) SetStatus(opts *bind.TransactOpts, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IWork.contract.Transact(opts, "setStatus", _workRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _workRef, uint8 _status) returns()
func (_IWork *IWorkSession) SetStatus(_workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IWork.Contract.SetStatus(&_IWork.TransactOpts, _workRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _workRef, uint8 _status) returns()
func (_IWork *IWorkTransactorSession) SetStatus(_workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IWork.Contract.SetStatus(&_IWork.TransactOpts, _workRef, _status)
}

// SetSubbie is a paid mutator transaction binding the contract method 0xd00c66af.
//
// Solidity: function setSubbie(bytes32 _workRef, address _subbie) returns()
func (_IWork *IWorkTransactor) SetSubbie(opts *bind.TransactOpts, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IWork.contract.Transact(opts, "setSubbie", _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0xd00c66af.
//
// Solidity: function setSubbie(bytes32 _workRef, address _subbie) returns()
func (_IWork *IWorkSession) SetSubbie(_workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IWork.Contract.SetSubbie(&_IWork.TransactOpts, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0xd00c66af.
//
// Solidity: function setSubbie(bytes32 _workRef, address _subbie) returns()
func (_IWork *IWorkTransactorSession) SetSubbie(_workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IWork.Contract.SetSubbie(&_IWork.TransactOpts, _workRef, _subbie)
}

// IWorkFactoryABI is the input ABI used to generate the binding from.
const IWorkFactoryABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"addInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumWorkStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"subbie\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"startDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sponsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWork\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"}],\"name\":\"getWorkContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_subbie\",\"type\":\"address\"}],\"name\":\"setSubbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setWorkStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWorkFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IWorkFactoryFuncSigs = map[string]string{
	"dcb5ad27": "addInfo(bytes32,bytes32,bytes32,(string,string,string,string))",
	"ddca1462": "addWork(bytes32,bytes32,(uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))",
	"82f92482": "getWorkContract(bytes32,bytes32)",
	"8896dca8": "setSubbie(bytes32,bytes32,address)",
	"e2aaa459": "setWorkStatus(bytes32,bytes32,uint8)",
}

// IWorkFactory is an auto generated Go binding around an Ethereum contract.
type IWorkFactory struct {
	IWorkFactoryCaller     // Read-only binding to the contract
	IWorkFactoryTransactor // Write-only binding to the contract
	IWorkFactoryFilterer   // Log filterer for contract events
}

// IWorkFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWorkFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWorkFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWorkFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWorkFactorySession struct {
	Contract     *IWorkFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWorkFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWorkFactoryCallerSession struct {
	Contract *IWorkFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IWorkFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWorkFactoryTransactorSession struct {
	Contract     *IWorkFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IWorkFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWorkFactoryRaw struct {
	Contract *IWorkFactory // Generic contract binding to access the raw methods on
}

// IWorkFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWorkFactoryCallerRaw struct {
	Contract *IWorkFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IWorkFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWorkFactoryTransactorRaw struct {
	Contract *IWorkFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWorkFactory creates a new instance of IWorkFactory, bound to a specific deployed contract.
func NewIWorkFactory(address common.Address, backend bind.ContractBackend) (*IWorkFactory, error) {
	contract, err := bindIWorkFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWorkFactory{IWorkFactoryCaller: IWorkFactoryCaller{contract: contract}, IWorkFactoryTransactor: IWorkFactoryTransactor{contract: contract}, IWorkFactoryFilterer: IWorkFactoryFilterer{contract: contract}}, nil
}

// NewIWorkFactoryCaller creates a new read-only instance of IWorkFactory, bound to a specific deployed contract.
func NewIWorkFactoryCaller(address common.Address, caller bind.ContractCaller) (*IWorkFactoryCaller, error) {
	contract, err := bindIWorkFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkFactoryCaller{contract: contract}, nil
}

// NewIWorkFactoryTransactor creates a new write-only instance of IWorkFactory, bound to a specific deployed contract.
func NewIWorkFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IWorkFactoryTransactor, error) {
	contract, err := bindIWorkFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkFactoryTransactor{contract: contract}, nil
}

// NewIWorkFactoryFilterer creates a new log filterer instance of IWorkFactory, bound to a specific deployed contract.
func NewIWorkFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IWorkFactoryFilterer, error) {
	contract, err := bindIWorkFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWorkFactoryFilterer{contract: contract}, nil
}

// bindIWorkFactory binds a generic wrapper to an already deployed contract.
func bindIWorkFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWorkFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorkFactory *IWorkFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IWorkFactory.Contract.IWorkFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorkFactory *IWorkFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkFactory.Contract.IWorkFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorkFactory *IWorkFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorkFactory.Contract.IWorkFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorkFactory *IWorkFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IWorkFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorkFactory *IWorkFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorkFactory *IWorkFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorkFactory.Contract.contract.Transact(opts, method, params...)
}

// GetWorkContract is a free data retrieval call binding the contract method 0x82f92482.
//
// Solidity: function getWorkContract(bytes32 _jobRef, bytes32 _workRef) view returns(address)
func (_IWorkFactory *IWorkFactoryCaller) GetWorkContract(opts *bind.CallOpts, _jobRef [32]byte, _workRef [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IWorkFactory.contract.Call(opts, out, "getWorkContract", _jobRef, _workRef)
	return *ret0, err
}

// GetWorkContract is a free data retrieval call binding the contract method 0x82f92482.
//
// Solidity: function getWorkContract(bytes32 _jobRef, bytes32 _workRef) view returns(address)
func (_IWorkFactory *IWorkFactorySession) GetWorkContract(_jobRef [32]byte, _workRef [32]byte) (common.Address, error) {
	return _IWorkFactory.Contract.GetWorkContract(&_IWorkFactory.CallOpts, _jobRef, _workRef)
}

// GetWorkContract is a free data retrieval call binding the contract method 0x82f92482.
//
// Solidity: function getWorkContract(bytes32 _jobRef, bytes32 _workRef) view returns(address)
func (_IWorkFactory *IWorkFactoryCallerSession) GetWorkContract(_jobRef [32]byte, _workRef [32]byte) (common.Address, error) {
	return _IWorkFactory.Contract.GetWorkContract(&_IWorkFactory.CallOpts, _jobRef, _workRef)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IWorkFactory *IWorkFactoryTransactor) AddInfo(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IWorkFactory.contract.Transact(opts, "addInfo", _jobRef, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IWorkFactory *IWorkFactorySession) AddInfo(_jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IWorkFactory.Contract.AddInfo(&_IWorkFactory.TransactOpts, _jobRef, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_IWorkFactory *IWorkFactoryTransactorSession) AddInfo(_jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _IWorkFactory.Contract.AddInfo(&_IWorkFactory.TransactOpts, _jobRef, _workRef, _infoRef, _info)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_IWorkFactory *IWorkFactoryTransactor) AddWork(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _IWorkFactory.contract.Transact(opts, "addWork", _jobRef, _workRef, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_IWorkFactory *IWorkFactorySession) AddWork(_jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _IWorkFactory.Contract.AddWork(&_IWorkFactory.TransactOpts, _jobRef, _workRef, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_IWorkFactory *IWorkFactoryTransactorSession) AddWork(_jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _IWorkFactory.Contract.AddWork(&_IWorkFactory.TransactOpts, _jobRef, _workRef, _work)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_IWorkFactory *IWorkFactoryTransactor) SetSubbie(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IWorkFactory.contract.Transact(opts, "setSubbie", _jobRef, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_IWorkFactory *IWorkFactorySession) SetSubbie(_jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IWorkFactory.Contract.SetSubbie(&_IWorkFactory.TransactOpts, _jobRef, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_IWorkFactory *IWorkFactoryTransactorSession) SetSubbie(_jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _IWorkFactory.Contract.SetSubbie(&_IWorkFactory.TransactOpts, _jobRef, _workRef, _subbie)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_IWorkFactory *IWorkFactoryTransactor) SetWorkStatus(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IWorkFactory.contract.Transact(opts, "setWorkStatus", _jobRef, _workRef, _status)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_IWorkFactory *IWorkFactorySession) SetWorkStatus(_jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IWorkFactory.Contract.SetWorkStatus(&_IWorkFactory.TransactOpts, _jobRef, _workRef, _status)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_IWorkFactory *IWorkFactoryTransactorSession) SetWorkStatus(_jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _IWorkFactory.Contract.SetWorkStatus(&_IWorkFactory.TransactOpts, _jobRef, _workRef, _status)
}

// InfoNodeABI is the input ABI used to generate the binding from.
const InfoNodeABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// InfoNodeFuncSigs maps the 4-byte function signature to its string representation.
var InfoNodeFuncSigs = map[string]string{
	"6d4ce63c": "get()",
}

// InfoNodeBin is the compiled bytecode used for deploying new contracts.
var InfoNodeBin = "0x608060405234801561001057600080fd5b506040516106ff3803806106ff83398101604081905261002f916101f1565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001661006157600080fd5b60008290558051805161007c916001916020909101906100cf565b5060208082015180516100939260029201906100cf565b50604081015180516100ad916003916020909101906100cf565b50606081015180516100c7916004916020909101906100cf565b5050506102f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011057805160ff191683800117855561013d565b8280016001018555821561013d579182015b8281111561013d578251825591602001919060010190610122565b5061014992915061014d565b5090565b61016791905b808211156101495760008155600101610153565b90565b600082601f83011261017a578081fd5b81516001600160401b0381111561018f578182fd5b60206101a3601f8301601f191682016102d2565b925081835284818386010111156101b957600080fd5b60005b828110156101d75784810182015184820183015281016101bc565b828111156101e85760008284860101525b50505092915050565b60008060408385031215610203578182fd5b825160208401519092506001600160401b0380821115610221578283fd5b81850160808188031215610233578384fd5b61023d60806102d2565b925080518281111561024d578485fd5b6102598882840161016a565b84525060208101518281111561026d578485fd5b6102798882840161016a565b602085015250604081015182811115610290578485fd5b61029c8882840161016a565b6040850152506060810151828111156102b3578485fd5b6102bf8882840161016a565b6060850152505050809150509250929050565b6040518181016001600160401b03811182821017156102f057600080fd5b604052919050565b6103f8806103076000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80636d4ce63c14610030575b600080fd5b61003861004f565b60405161004692919061033e565b60405180910390f35b60006100596102cb565b600054604080516001805460a06020601f60026000196101008688161502019094169390930492830181900402840181019094526080830181815291938492849290918491908401828280156100f05780601f106100c5576101008083540402835291602001916100f0565b820191906000526020600020905b8154815290600101906020018083116100d357829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101925780601f1061016757610100808354040283529160200191610192565b820191906000526020600020905b81548152906001019060200180831161017557829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156102245780601f106101f957610100808354040283529160200191610224565b820191906000526020600020905b81548152906001019060200180831161020757829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156102b85780601f1061028d576101008083540402835291602001916102b8565b820191906000526020600020905b81548152906001019060200180831161029b57829003601f168201915b5050505050815250509050915091509091565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b60008151808452815b81811015610318576020818501810151868301820152016102fc565b818111156103295782602083870101525b50601f01601f19169290920160200192915050565b60008382526040602083015282516080604084015261036060c08401826102f3565b60208501519150603f198085830301606086015261037e82846102f3565b604087015193508186820301608087015261039981856102f3565b92505060608601519250808583030160a0860152506103b881836102f3565b969550505050505056fea264697066735822122051c640576777bbdc856893273c3ae7351824716c50548076450668cfda633c9364736f6c63430006080033"

// DeployInfoNode deploys a new Ethereum contract, binding an instance of InfoNode to it.
func DeployInfoNode(auth *bind.TransactOpts, backend bind.ContractBackend, _infoRef [32]byte, _info ExtraInfo) (common.Address, *types.Transaction, *InfoNode, error) {
	parsed, err := abi.JSON(strings.NewReader(InfoNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InfoNodeBin), backend, _infoRef, _info)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InfoNode{InfoNodeCaller: InfoNodeCaller{contract: contract}, InfoNodeTransactor: InfoNodeTransactor{contract: contract}, InfoNodeFilterer: InfoNodeFilterer{contract: contract}}, nil
}

// InfoNode is an auto generated Go binding around an Ethereum contract.
type InfoNode struct {
	InfoNodeCaller     // Read-only binding to the contract
	InfoNodeTransactor // Write-only binding to the contract
	InfoNodeFilterer   // Log filterer for contract events
}

// InfoNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type InfoNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InfoNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfoNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfoNodeSession struct {
	Contract     *InfoNode         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InfoNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfoNodeCallerSession struct {
	Contract *InfoNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// InfoNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfoNodeTransactorSession struct {
	Contract     *InfoNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InfoNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type InfoNodeRaw struct {
	Contract *InfoNode // Generic contract binding to access the raw methods on
}

// InfoNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfoNodeCallerRaw struct {
	Contract *InfoNodeCaller // Generic read-only contract binding to access the raw methods on
}

// InfoNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfoNodeTransactorRaw struct {
	Contract *InfoNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInfoNode creates a new instance of InfoNode, bound to a specific deployed contract.
func NewInfoNode(address common.Address, backend bind.ContractBackend) (*InfoNode, error) {
	contract, err := bindInfoNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfoNode{InfoNodeCaller: InfoNodeCaller{contract: contract}, InfoNodeTransactor: InfoNodeTransactor{contract: contract}, InfoNodeFilterer: InfoNodeFilterer{contract: contract}}, nil
}

// NewInfoNodeCaller creates a new read-only instance of InfoNode, bound to a specific deployed contract.
func NewInfoNodeCaller(address common.Address, caller bind.ContractCaller) (*InfoNodeCaller, error) {
	contract, err := bindInfoNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfoNodeCaller{contract: contract}, nil
}

// NewInfoNodeTransactor creates a new write-only instance of InfoNode, bound to a specific deployed contract.
func NewInfoNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*InfoNodeTransactor, error) {
	contract, err := bindInfoNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfoNodeTransactor{contract: contract}, nil
}

// NewInfoNodeFilterer creates a new log filterer instance of InfoNode, bound to a specific deployed contract.
func NewInfoNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*InfoNodeFilterer, error) {
	contract, err := bindInfoNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfoNodeFilterer{contract: contract}, nil
}

// bindInfoNode binds a generic wrapper to an already deployed contract.
func bindInfoNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InfoNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfoNode *InfoNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InfoNode.Contract.InfoNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfoNode *InfoNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfoNode.Contract.InfoNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfoNode *InfoNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfoNode.Contract.InfoNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfoNode *InfoNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InfoNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfoNode *InfoNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfoNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfoNode *InfoNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfoNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (string,string,string,string))
func (_InfoNode *InfoNodeCaller) Get(opts *bind.CallOpts) ([32]byte, ExtraInfo, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(ExtraInfo)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _InfoNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (string,string,string,string))
func (_InfoNode *InfoNodeSession) Get() ([32]byte, ExtraInfo, error) {
	return _InfoNode.Contract.Get(&_InfoNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (string,string,string,string))
func (_InfoNode *InfoNodeCallerSession) Get() ([32]byte, ExtraInfo, error) {
	return _InfoNode.Contract.Get(&_InfoNode.CallOpts)
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

// JobNodeABI is the input ABI used to generate the binding from.
const JobNodeABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"_job\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"addInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumWorkStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"subbie\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"startDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sponsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWork\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"}],\"name\":\"getWorkContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_subbie\",\"type\":\"address\"}],\"name\":\"setSubbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setWorkStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// JobNodeFuncSigs maps the 4-byte function signature to its string representation.
var JobNodeFuncSigs = map[string]string{
	"dcb5ad27": "addInfo(bytes32,bytes32,bytes32,(string,string,string,string))",
	"ddca1462": "addWork(bytes32,bytes32,(uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))",
	"6d4ce63c": "get()",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"82f92482": "getWorkContract(bytes32,bytes32)",
	"8de654ba": "setStatus(bytes32,uint8)",
	"8896dca8": "setSubbie(bytes32,bytes32,address)",
	"e2aaa459": "setWorkStatus(bytes32,bytes32,uint8)",
}

// JobNodeBin is the compiled bytecode used for deploying new contracts.
var JobNodeBin = "0x60806040523480156200001157600080fd5b506040516200257438038062002574833981016040819052620000349162000259565b8160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007957506000815160058111156200007757fe5b115b80156200009357506005815160058111156200009157fe5b105b8015620000a557506000816020015151115b8015620000b757506000816040015151115b620000c157600080fd5b600082905580516001805460ff191681836005811115620000de57fe5b02179055506020808201518051620000fb92600292019062000125565b5060408101518051620001179160039160209091019062000125565b50506000600655506200033d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200016857805160ff191683800117855562000198565b8280016001018555821562000198579182015b82811115620001985782518255916020019190600101906200017b565b50620001a6929150620001aa565b5090565b620001c791905b80821115620001a65760008155600101620001b1565b90565b600082601f830112620001db578081fd5b81516001600160401b03811115620001f1578182fd5b602062000207601f8301601f1916820162000316565b925081835284818386010111156200021e57600080fd5b60005b828110156200023e57848101820151848201830152810162000221565b82811115620002505760008284860101525b50505092915050565b600080604083850312156200026c578182fd5b825160208401519092506001600160401b03808211156200028b578283fd5b818501606081880312156200029e578384fd5b620002aa606062000316565b9250805160068110620002bb578485fd5b8352602081015182811115620002cf578485fd5b620002dd88828401620001ca565b602085015250604081015182811115620002f5578485fd5b6200030388828401620001ca565b6040850152505050809150509250929050565b6040518181016001600160401b03811182821017156200033557600080fd5b604052919050565b612227806200034d6000396000f3fe60806040523480156200001157600080fd5b5060043610620000a05760003560e01c80638de654ba116200006f5780638de654ba1462000120578063bc5993411462000137578063dcb5ad27146200014e578063ddca14621462000165578063e2aaa459146200017c57620000a0565b806367e0badb14620000a55780636d4ce63c14620000c757806382f9248214620000e15780638896dca81462000107575b600080fd5b620000af62000193565b604051620000be919062000b82565b60405180910390f35b620000d162000199565b604051620000be92919062000c35565b620000f8620000f236600462000819565b62000323565b604051620000be919062000b6e565b6200011e620001183660046200083b565b6200037c565b005b6200011e6200013136600462000ab5565b62000469565b620000af6200014836600462000aeb565b620004c6565b6200011e6200015f36600462000880565b62000500565b6200011e6200017636600462000984565b620005b8565b6200011e6200018d36600462000a81565b6200063d565b60065490565b6000620001a562000753565b6000546040805160608101909152600180549091908290829060ff166005811115620001cd57fe5b6005811115620001d957fe5b8152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156200027a5780601f106200024e576101008083540402835291602001916200027a565b820191906000526020600020905b8154815290600101906020018083116200025c57829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f81018390048302850183019091528084529381019390830182828015620003105780601f10620002e45761010080835404028352916020019162000310565b820191906000526020600020905b815481529060010190602001808311620002f257829003601f168201915b5050505050815250509050915091509091565b60008054831480156200034f57506000828152600460205260409020600101546001600160a01b031615155b6200035957600080fd5b506000818152600460205260409020600101546001600160a01b03165b92915050565b8260001a60f81b6001600160f81b031916158015906200039d575060005483145b8015620003b957508160001a60f81b6001600160f81b03191615155b8015620003df57506000828152600460205260409020600101546001600160a01b031615155b620003e957600080fd5b60008281526004602081905260409182902060010154915163d00c66af60e01b81526001600160a01b0390921691829163d00c66af916200042f91879187910162000b8b565b600060405180830381600087803b1580156200044a57600080fd5b505af11580156200045f573d6000803e3d6000fd5b5050505050505050565b600054821480156200047d575060ff811615155b80156200048d5750600560ff8216105b6200049757600080fd5b8060ff166005811115620004a757fe5b6001805460ff191681836005811115620004bd57fe5b02179055505050565b6006546000908210620004d857600080fd5b6005805483908110620004e757fe5b9060005260206000209060020201600001549050919050565b600054841480156200052b57506000838152600460205260409020600101546001600160a01b031615155b6200053557600080fd5b600083815260046020819052604091829020600101549151631ea71c6b60e31b81526001600160a01b0390921691829163f538e358916200057d918891889188910162000ba2565b600060405180830381600087803b1580156200059857600080fd5b505af1158015620005ad573d6000803e3d6000fd5b505050505050505050565b60005483148015620005d957508160001a60f81b6001600160f81b03191615155b620005e357600080fd5b60008282604051620005f59062000775565b6200060292919062000c96565b604051809103906000f0801580156200061f573d6000803e3d6000fd5b509050620006366004848363ffffffff620006b816565b5050505050565b600054831480156200066857506000828152600460205260409020600101546001600160a01b031615155b6200067257600080fd5b6000828152600460208190526040918290206001015491516346f32a5d60e11b81526001600160a01b03909216918291638de654ba916200042f91879187910162000d43565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b0385161790558015620006f85760019150506200074c565b50600180850180549182018082556000868152602088905260409020819055859190839081106200072557fe5b600091825260208220600291820201929092559086018054600101905591506200074c9050565b9392505050565b6040805160608101909152806000815260200160608152602001606081525090565b6114758062000d7d83390190565b80356001600160a01b03811681146200037657600080fd5b8035600981106200037657600080fd5b600082601f830112620007bc578081fd5b813567ffffffffffffffff811115620007d3578182fd5b620007e8601f8201601f191660200162000d54565b91508082528360208285010111156200080057600080fd5b8060208401602084013760009082016020015292915050565b600080604083850312156200082c578182fd5b50508035926020909101359150565b60008060006060848603121562000850578081fd5b833592506020840135915060408401356001600160a01b038116811462000875578182fd5b809150509250925092565b6000806000806080858703121562000896578081fd5b843593506020850135925060408501359150606085013567ffffffffffffffff80821115620008c3578283fd5b8187016080818a031215620008d6578384fd5b620008e2608062000d54565b9250803582811115620008f3578485fd5b620009018a828401620007ab565b84525060208101358281111562000916578485fd5b620009248a828401620007ab565b6020850152506040810135828111156200093c578485fd5b6200094a8a828401620007ab565b60408501525060608101358281111562000962578485fd5b620009708a828401620007ab565b606085015250959894975092955093505050565b60008060006060848603121562000999578283fd5b8335925060208401359150604084013567ffffffffffffffff80821115620009bf578283fd5b610140918601808803831315620009d4578384fd5b620009df8362000d54565b620009eb89836200079b565b8152620009fc896020840162000783565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c082013560c082015260e082013560e0820152610100935083820135848201526101209350838201358381111562000a60578586fd5b62000a6e8a828501620007ab565b8583015250809450505050509250925092565b60008060006060848603121562000a96578283fd5b8335925060208401359150604084013560ff8116811462000875578182fd5b6000806040838503121562000ac8578182fd5b82359150602083013560ff8116811462000ae0578182fd5b809150509250929050565b60006020828403121562000afd578081fd5b5035919050565b6001600160a01b03169052565b6009811062000b1c57fe5b9052565b60008151808452815b8181101562000b475760208185018101518683018201520162000b29565b8181111562000b595782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b90815260200190565b9182526001600160a01b0316602082015260400190565b60008482528360208301526060604083015282516080606084015262000bcc60e084018262000b20565b60208501519150605f198085830301608086015262000bec828462000b20565b60408701519350818682030160a087015262000c09818562000b20565b92505060608601519250808583030160c08601525062000c2a818362000b20565b979650505050505050565b60008382526040602083015282516006811062000c4e57fe5b6040830152602083015160608084015262000c6d60a084018262000b20565b6040850151848203603f19016080860152915062000c8c818362000b20565b9695505050505050565b60008382526040602083015262000cb260408301845162000b11565b602083015162000cc6606084018262000b04565b5060408301516080830152606083015160a0830152608083015160c083015260a083015160e083015260c0830151610100818185015260e0850151915061012082818601528186015192506101409150828286015280860151925050806101608501525062000d3a61018084018262000b20565b95945050505050565b91825260ff16602082015260400190565b60405181810167ffffffffffffffff8111828210171562000d7457600080fd5b60405291905056fe60806040523480156200001157600080fd5b5060405162001475380380620014758339810160408190526200003491620002c9565b8160001a60f81b6001600160f81b031916158015906200006157506000815160088111156200005f57fe5b115b80156200007b57506008815160088111156200007957fe5b105b80156200009b575060c081015160001a60f81b6001600160f81b03191615155b8015620000ab575060e081015115155b8015620000bc575061010081015115155b620000c657600080fd5b600082905580516001805460ff191681836008811115620000e357fe5b021790555060208181015160018054610100600160a81b0319166101006001600160a01b03909316830217905560c083015160065560408301516002556060830151600355608083015160045560a083015160055560e083015160075582015160085561012082015180516200015e92600992019062000167565b505050620003e4565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001aa57805160ff1916838001178555620001da565b82800160010185558215620001da579182015b82811115620001da578251825591602001919060010190620001bd565b50620001e8929150620001ec565b5090565b6200020991905b80821115620001e85760008155600101620001f3565b90565b80516001600160a01b03811681146200022457600080fd5b92915050565b8051600981106200022457600080fd5b600082601f8301126200024b578081fd5b81516001600160401b0381111562000261578182fd5b602062000277601f8301601f19168201620003bd565b925081835284818386010111156200028e57600080fd5b60005b82811015620002ae57848101820151848201830152810162000291565b82811115620002c05760008284860101525b50505092915050565b60008060408385031215620002dc578182fd5b825160208401519092506001600160401b0380821115620002fb578283fd5b61014091850180870383131562000310578384fd5b6200031b83620003bd565b6200032788836200022a565b81526200033888602084016200020c565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c082015160c082015260e082015160e082015261010093508382015184820152610120935083820151838111156200039c578586fd5b620003aa898285016200023a565b8583015250809450505050509250929050565b6040518181016001600160401b0381118282101715620003dc57600080fd5b604052919050565b61108180620003f46000396000f3fe60806040523480156200001157600080fd5b5060043610620000885760003560e01c80638de654ba11620000635780638de654ba14620000ef578063bc5993411462000108578063d00c66af146200011f578063f538e35814620001365762000088565b806352c43d9f146200008d57806367e0badb14620000bc5780636d4ce63c14620000d5575b600080fd5b620000a46200009e36600462000603565b6200014d565b604051620000b39190620007ce565b60405180910390f35b620000c6620001a4565b604051620000b39190620007e2565b620000df620001aa565b604051620000b392919062000877565b620001066200010036600462000720565b620002ef565b005b620000c6620001193660046200074b565b6200034c565b6200010662000130366004620005c7565b62000386565b620001066200014736600462000625565b620003d3565b60008054831480156200017957506000828152600a60205260409020600101546001600160a01b031615155b6200018357600080fd5b506000908152600a60205260409020600101546001600160a01b0316919050565b600c5490565b6000620001b6620004f3565b600054604080516101408101909152600180549091908290829060ff166008811115620001df57fe5b6008811115620001eb57fe5b81528154610100908190046001600160a01b0316602080840191909152600180850154604080860191909152600280870154606087015260038701546080870152600487015460a0870152600587015460c0870152600687015460e08701526007870154858701526008870180548351948116159096026000190190951604601f810184900484028301840190915280825261012090940193909291830182828015620002dc5780601f10620002b057610100808354040283529160200191620002dc565b820191906000526020600020905b815481529060010190602001808311620002be57829003601f168201915b5050505050815250509050915091509091565b6000548214801562000303575060ff811615155b8015620003135750600860ff8216105b6200031d57600080fd5b8060ff1660088111156200032d57fe5b6001805460ff1916818360088111156200034357fe5b02179055505050565b600c5460009082106200035e57600080fd5b600b8054839081106200036d57fe5b9060005260206000209060020201600001549050919050565b60005482148015620003a057506001600160a01b03811615155b620003aa57600080fd5b600180546001600160a01b0390921661010002610100600160a81b031990921691909117905550565b60005483148015620003f457508160001a60f81b6001600160f81b03191615155b620003fe57600080fd5b6000828260405162000410906200054b565b6200041d929190620007eb565b604051809103906000f0801580156200043a573d6000803e3d6000fd5b50905062000451600a848363ffffffff6200045816565b5050505050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000498576001915050620004ec565b5060018085018054918201808255600086815260208890526040902081905585919083908110620004c557fe5b60009182526020822060029182020192909255908601805460010190559150620004ec9050565b9392505050565b60408051610140810190915280600081526000602082018190526040820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301919091526101209091015290565b6106ff806200094d83390190565b600082601f8301126200056a578081fd5b813567ffffffffffffffff81111562000581578182fd5b62000596601f8201601f191660200162000924565b9150808252836020828501011115620005ae57600080fd5b8060208401602084013760009082016020015292915050565b60008060408385031215620005da578182fd5b8235915060208301356001600160a01b0381168114620005f8578182fd5b809150509250929050565b6000806040838503121562000616578182fd5b50508035926020909101359150565b6000806000606084860312156200063a578081fd5b8335925060208401359150604084013567ffffffffffffffff8082111562000660578283fd5b8186016080818903121562000673578384fd5b6200067f608062000924565b925080358281111562000690578485fd5b6200069e8982840162000559565b845250602081013582811115620006b3578485fd5b620006c18982840162000559565b602085015250604081013582811115620006d9578485fd5b620006e78982840162000559565b604085015250606081013582811115620006ff578485fd5b6200070d8982840162000559565b6060850152505050809150509250925092565b6000806040838503121562000733578182fd5b82359150602083013560ff81168114620005f8578182fd5b6000602082840312156200075d578081fd5b5035919050565b6001600160a01b03169052565b600981106200077c57fe5b9052565b60008151808452815b81811015620007a75760208185018101518683018201520162000789565b81811115620007b95782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b90815260200190565b6000838252604060208301528251608060408401526200080f60c084018262000780565b60208501519150603f19808583030160608601526200082f828462000780565b60408701519350818682030160808701526200084c818562000780565b92505060608601519250808583030160a0860152506200086d818362000780565b9695505050505050565b6000838252604060208301526200089360408301845162000771565b6020830151620008a7606084018262000764565b5060408301516080830152606083015160a0830152608083015160c083015260a083015160e083015260c0830151610100818185015260e085015191506101208281860152818601519250610140915082828601528086015192505080610160850152506200091b61018084018262000780565b95945050505050565b60405181810167ffffffffffffffff811182821017156200094457600080fd5b60405291905056fe608060405234801561001057600080fd5b506040516106ff3803806106ff83398101604081905261002f916101f1565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001661006157600080fd5b60008290558051805161007c916001916020909101906100cf565b5060208082015180516100939260029201906100cf565b50604081015180516100ad916003916020909101906100cf565b50606081015180516100c7916004916020909101906100cf565b5050506102f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011057805160ff191683800117855561013d565b8280016001018555821561013d579182015b8281111561013d578251825591602001919060010190610122565b5061014992915061014d565b5090565b61016791905b808211156101495760008155600101610153565b90565b600082601f83011261017a578081fd5b81516001600160401b0381111561018f578182fd5b60206101a3601f8301601f191682016102d2565b925081835284818386010111156101b957600080fd5b60005b828110156101d75784810182015184820183015281016101bc565b828111156101e85760008284860101525b50505092915050565b60008060408385031215610203578182fd5b825160208401519092506001600160401b0380821115610221578283fd5b81850160808188031215610233578384fd5b61023d60806102d2565b925080518281111561024d578485fd5b6102598882840161016a565b84525060208101518281111561026d578485fd5b6102798882840161016a565b602085015250604081015182811115610290578485fd5b61029c8882840161016a565b6040850152506060810151828111156102b3578485fd5b6102bf8882840161016a565b6060850152505050809150509250929050565b6040518181016001600160401b03811182821017156102f057600080fd5b604052919050565b6103f8806103076000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80636d4ce63c14610030575b600080fd5b61003861004f565b60405161004692919061033e565b60405180910390f35b60006100596102cb565b600054604080516001805460a06020601f60026000196101008688161502019094169390930492830181900402840181019094526080830181815291938492849290918491908401828280156100f05780601f106100c5576101008083540402835291602001916100f0565b820191906000526020600020905b8154815290600101906020018083116100d357829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101925780601f1061016757610100808354040283529160200191610192565b820191906000526020600020905b81548152906001019060200180831161017557829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156102245780601f106101f957610100808354040283529160200191610224565b820191906000526020600020905b81548152906001019060200180831161020757829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156102b85780601f1061028d576101008083540402835291602001916102b8565b820191906000526020600020905b81548152906001019060200180831161029b57829003601f168201915b5050505050815250509050915091509091565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b60008151808452815b81811015610318576020818501810151868301820152016102fc565b818111156103295782602083870101525b50601f01601f19169290920160200192915050565b60008382526040602083015282516080604084015261036060c08401826102f3565b60208501519150603f198085830301606086015261037e82846102f3565b604087015193508186820301608087015261039981856102f3565b92505060608601519250808583030160a0860152506103b881836102f3565b969550505050505056fea264697066735822122051c640576777bbdc856893273c3ae7351824716c50548076450668cfda633c9364736f6c63430006080033a2646970667358221220091b9b56f546f270a7ac47f850fc9812ce3c6b78ff3d38b959d6fdbce031155964736f6c63430006080033a26469706673582212206cf0adc33f806ba7b2b53ac27f512310dcfcdd752cc1fb117ef94f8866c2abd764736f6c63430006080033"

// DeployJobNode deploys a new Ethereum contract, binding an instance of JobNode to it.
func DeployJobNode(auth *bind.TransactOpts, backend bind.ContractBackend, _jobRef [32]byte, _job Job) (common.Address, *types.Transaction, *JobNode, error) {
	parsed, err := abi.JSON(strings.NewReader(JobNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JobNodeBin), backend, _jobRef, _job)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JobNode{JobNodeCaller: JobNodeCaller{contract: contract}, JobNodeTransactor: JobNodeTransactor{contract: contract}, JobNodeFilterer: JobNodeFilterer{contract: contract}}, nil
}

// JobNode is an auto generated Go binding around an Ethereum contract.
type JobNode struct {
	JobNodeCaller     // Read-only binding to the contract
	JobNodeTransactor // Write-only binding to the contract
	JobNodeFilterer   // Log filterer for contract events
}

// JobNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobNodeSession struct {
	Contract     *JobNode          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobNodeCallerSession struct {
	Contract *JobNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// JobNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobNodeTransactorSession struct {
	Contract     *JobNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// JobNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobNodeRaw struct {
	Contract *JobNode // Generic contract binding to access the raw methods on
}

// JobNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobNodeCallerRaw struct {
	Contract *JobNodeCaller // Generic read-only contract binding to access the raw methods on
}

// JobNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobNodeTransactorRaw struct {
	Contract *JobNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobNode creates a new instance of JobNode, bound to a specific deployed contract.
func NewJobNode(address common.Address, backend bind.ContractBackend) (*JobNode, error) {
	contract, err := bindJobNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JobNode{JobNodeCaller: JobNodeCaller{contract: contract}, JobNodeTransactor: JobNodeTransactor{contract: contract}, JobNodeFilterer: JobNodeFilterer{contract: contract}}, nil
}

// NewJobNodeCaller creates a new read-only instance of JobNode, bound to a specific deployed contract.
func NewJobNodeCaller(address common.Address, caller bind.ContractCaller) (*JobNodeCaller, error) {
	contract, err := bindJobNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobNodeCaller{contract: contract}, nil
}

// NewJobNodeTransactor creates a new write-only instance of JobNode, bound to a specific deployed contract.
func NewJobNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*JobNodeTransactor, error) {
	contract, err := bindJobNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobNodeTransactor{contract: contract}, nil
}

// NewJobNodeFilterer creates a new log filterer instance of JobNode, bound to a specific deployed contract.
func NewJobNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*JobNodeFilterer, error) {
	contract, err := bindJobNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobNodeFilterer{contract: contract}, nil
}

// bindJobNode binds a generic wrapper to an already deployed contract.
func bindJobNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JobNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobNode *JobNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JobNode.Contract.JobNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobNode *JobNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobNode.Contract.JobNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobNode *JobNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobNode.Contract.JobNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobNode *JobNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JobNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobNode *JobNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobNode *JobNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,string,string))
func (_JobNode *JobNodeCaller) Get(opts *bind.CallOpts) ([32]byte, Job, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(Job)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _JobNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,string,string))
func (_JobNode *JobNodeSession) Get() ([32]byte, Job, error) {
	return _JobNode.Contract.Get(&_JobNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,string,string))
func (_JobNode *JobNodeCallerSession) Get() ([32]byte, Job, error) {
	return _JobNode.Contract.Get(&_JobNode.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_JobNode *JobNodeCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobNode.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_JobNode *JobNodeSession) GetNum() (*big.Int, error) {
	return _JobNode.Contract.GetNum(&_JobNode.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_JobNode *JobNodeCallerSession) GetNum() (*big.Int, error) {
	return _JobNode.Contract.GetNum(&_JobNode.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_JobNode *JobNodeCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _JobNode.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_JobNode *JobNodeSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _JobNode.Contract.GetReference(&_JobNode.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_JobNode *JobNodeCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _JobNode.Contract.GetReference(&_JobNode.CallOpts, _index)
}

// GetWorkContract is a free data retrieval call binding the contract method 0x82f92482.
//
// Solidity: function getWorkContract(bytes32 _jobRef, bytes32 _workRef) view returns(address)
func (_JobNode *JobNodeCaller) GetWorkContract(opts *bind.CallOpts, _jobRef [32]byte, _workRef [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _JobNode.contract.Call(opts, out, "getWorkContract", _jobRef, _workRef)
	return *ret0, err
}

// GetWorkContract is a free data retrieval call binding the contract method 0x82f92482.
//
// Solidity: function getWorkContract(bytes32 _jobRef, bytes32 _workRef) view returns(address)
func (_JobNode *JobNodeSession) GetWorkContract(_jobRef [32]byte, _workRef [32]byte) (common.Address, error) {
	return _JobNode.Contract.GetWorkContract(&_JobNode.CallOpts, _jobRef, _workRef)
}

// GetWorkContract is a free data retrieval call binding the contract method 0x82f92482.
//
// Solidity: function getWorkContract(bytes32 _jobRef, bytes32 _workRef) view returns(address)
func (_JobNode *JobNodeCallerSession) GetWorkContract(_jobRef [32]byte, _workRef [32]byte) (common.Address, error) {
	return _JobNode.Contract.GetWorkContract(&_JobNode.CallOpts, _jobRef, _workRef)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_JobNode *JobNodeTransactor) AddInfo(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _JobNode.contract.Transact(opts, "addInfo", _jobRef, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_JobNode *JobNodeSession) AddInfo(_jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _JobNode.Contract.AddInfo(&_JobNode.TransactOpts, _jobRef, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_JobNode *JobNodeTransactorSession) AddInfo(_jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _JobNode.Contract.AddInfo(&_JobNode.TransactOpts, _jobRef, _workRef, _infoRef, _info)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_JobNode *JobNodeTransactor) AddWork(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _JobNode.contract.Transact(opts, "addWork", _jobRef, _workRef, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_JobNode *JobNodeSession) AddWork(_jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _JobNode.Contract.AddWork(&_JobNode.TransactOpts, _jobRef, _workRef, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_JobNode *JobNodeTransactorSession) AddWork(_jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _JobNode.Contract.AddWork(&_JobNode.TransactOpts, _jobRef, _workRef, _work)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _jobRef, uint8 _status) returns()
func (_JobNode *JobNodeTransactor) SetStatus(opts *bind.TransactOpts, _jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _JobNode.contract.Transact(opts, "setStatus", _jobRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _jobRef, uint8 _status) returns()
func (_JobNode *JobNodeSession) SetStatus(_jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _JobNode.Contract.SetStatus(&_JobNode.TransactOpts, _jobRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _jobRef, uint8 _status) returns()
func (_JobNode *JobNodeTransactorSession) SetStatus(_jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _JobNode.Contract.SetStatus(&_JobNode.TransactOpts, _jobRef, _status)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_JobNode *JobNodeTransactor) SetSubbie(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _JobNode.contract.Transact(opts, "setSubbie", _jobRef, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_JobNode *JobNodeSession) SetSubbie(_jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _JobNode.Contract.SetSubbie(&_JobNode.TransactOpts, _jobRef, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_JobNode *JobNodeTransactorSession) SetSubbie(_jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _JobNode.Contract.SetSubbie(&_JobNode.TransactOpts, _jobRef, _workRef, _subbie)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_JobNode *JobNodeTransactor) SetWorkStatus(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _JobNode.contract.Transact(opts, "setWorkStatus", _jobRef, _workRef, _status)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_JobNode *JobNodeSession) SetWorkStatus(_jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _JobNode.Contract.SetWorkStatus(&_JobNode.TransactOpts, _jobRef, _workRef, _status)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_JobNode *JobNodeTransactorSession) SetWorkStatus(_jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _JobNode.Contract.SetWorkStatus(&_JobNode.TransactOpts, _jobRef, _workRef, _status)
}

// JobsABI is the input ABI used to generate the binding from.
const JobsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"addInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"_job\",\"type\":\"tuple\"}],\"name\":\"addJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumWorkStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"subbie\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"startDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sponsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWork\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"}],\"name\":\"getJobContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_subbie\",\"type\":\"address\"}],\"name\":\"setSubbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setWorkStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// JobsFuncSigs maps the 4-byte function signature to its string representation.
var JobsFuncSigs = map[string]string{
	"dcb5ad27": "addInfo(bytes32,bytes32,bytes32,(string,string,string,string))",
	"8b7a9f98": "addJob(bytes32,(uint8,string,string))",
	"ddca1462": "addWork(bytes32,bytes32,(uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))",
	"c918e294": "getJobContract(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"e1908676": "setJobStatus(bytes32,uint8)",
	"8896dca8": "setSubbie(bytes32,bytes32,address)",
	"e2aaa459": "setWorkStatus(bytes32,bytes32,uint8)",
}

// JobsBin is the compiled bytecode used for deploying new contracts.
var JobsBin = "0x608060405234801561001057600080fd5b506000600255613269806100256000396000f3fe60806040523480156200001157600080fd5b5060043610620000a05760003560e01c8063c918e294116200006f578063c918e294146200010e578063dcb5ad271462000134578063ddca1462146200014b578063e19086761462000162578063e2aaa459146200017957620000a0565b806367e0badb14620000a55780638896dca814620000c75780638b7a9f9814620000e0578063bc59934114620000f7575b600080fd5b620000af62000190565b604051620000be919062000aa0565b60405180910390f35b620000de620000d8366004620006b5565b62000196565b005b620000de620000f13660046200092f565b6200023d565b620000af620001083660046200069c565b620002b1565b620001256200011f3660046200069c565b620002eb565b604051620000be919062000a8c565b620000de62000145366004620006fa565b6200032f565b620000de6200015c366004620007fe565b620003d9565b620000de62000173366004620009ec565b62000446565b620000de6200018a366004620008fb565b620004ea565b60025490565b6000838152602081905260409020600101546001600160a01b0316620001bb57600080fd5b60008381526020819052604090819020600101549051631112db9560e31b81526001600160a01b03909116908190638896dca890620002039087908790879060040162000aa9565b600060405180830381600087803b1580156200021e57600080fd5b505af115801562000233573d6000803e3d6000fd5b5050505050505050565b8160001a60f81b6001600160f81b0319166200025857600080fd5b600082826040516200026a90620005f2565b6200027792919062000c2f565b604051809103906000f08015801562000294573d6000803e3d6000fd5b509050620002ab6000848363ffffffff6200055716565b50505050565b6002546000908210620002c357600080fd5b6001805483908110620002d257fe5b9060005260206000209060020201600001549050919050565b6000818152602081905260408120600101546001600160a01b03166200031057600080fd5b506000908152602081905260409020600101546001600160a01b031690565b6000848152602081905260409020600101546001600160a01b03166200035457600080fd5b6000848152602081905260409081902060010154905163dcb5ad2760e01b81526001600160a01b0390911690819063dcb5ad27906200039e90889088908890889060040162000ac8565b600060405180830381600087803b158015620003b957600080fd5b505af1158015620003ce573d6000803e3d6000fd5b505050505050505050565b6000838152602081905260409020600101546001600160a01b0316620003fe57600080fd5b60008381526020819052604090819020600101549051636ee50a3160e11b81526001600160a01b0390911690819063ddca146290620002039087908790879060040162000b62565b6000828152602081905260409020600101546001600160a01b03166200046b57600080fd5b600082815260208190526040908190206001015490516346f32a5d60e11b81526001600160a01b03909116908190638de654ba90620004b1908690869060040162000c86565b600060405180830381600087803b158015620004cc57600080fd5b505af1158015620004e1573d6000803e3d6000fd5b50505050505050565b6000838152602081905260409020600101546001600160a01b03166200050f57600080fd5b6000838152602081905260409081902060010154905163e2aaa45960e01b81526001600160a01b0390911690819063e2aaa45990620002039087908790879060040162000c16565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000597576001915050620005eb565b5060018085018054918201808255600086815260208890526040902081905585919083908110620005c457fe5b60009182526020822060029182020192909255908601805460010190559150620005eb9050565b9392505050565b6125748062000cc083390190565b80356001600160a01b03811681146200061857600080fd5b92915050565b8035600981106200061857600080fd5b600082601f8301126200063f578081fd5b813567ffffffffffffffff81111562000656578182fd5b6200066b601f8201601f191660200162000c97565b91508082528360208285010111156200068357600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215620006ae578081fd5b5035919050565b600080600060608486031215620006ca578182fd5b833592506020840135915060408401356001600160a01b0381168114620006ef578182fd5b809150509250925092565b6000806000806080858703121562000710578081fd5b843593506020850135925060408501359150606085013567ffffffffffffffff808211156200073d578283fd5b8187016080818a03121562000750578384fd5b6200075c608062000c97565b92508035828111156200076d578485fd5b6200077b8a8284016200062e565b84525060208101358281111562000790578485fd5b6200079e8a8284016200062e565b602085015250604081013582811115620007b6578485fd5b620007c48a8284016200062e565b604085015250606081013582811115620007dc578485fd5b620007ea8a8284016200062e565b606085015250959894975092955093505050565b60008060006060848603121562000813578283fd5b8335925060208401359150604084013567ffffffffffffffff8082111562000839578283fd5b6101409186018088038313156200084e578384fd5b620008598362000c97565b6200086589836200061e565b815262000876896020840162000600565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c082013560c082015260e082013560e08201526101009350838201358482015261012093508382013583811115620008da578586fd5b620008e88a8285016200062e565b8583015250809450505050509250925092565b60008060006060848603121562000910578283fd5b8335925060208401359150604084013560ff81168114620006ef578182fd5b6000806040838503121562000942578182fd5b82359150602083013567ffffffffffffffff8082111562000961578283fd5b8185016060818803121562000974578384fd5b62000980606062000c97565b925080356006811062000991578485fd5b8352602081013582811115620009a5578485fd5b620009b3888284016200062e565b602085015250604081013582811115620009cb578485fd5b620009d9888284016200062e565b6040850152505050809150509250929050565b60008060408385031215620009ff578182fd5b82359150602083013560ff8116811462000a17578182fd5b809150509250929050565b6001600160a01b03169052565b6009811062000a3a57fe5b9052565b60008151808452815b8181101562000a655760208185018101518683018201520162000a47565b8181111562000a775782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b90815260200190565b92835260208301919091526001600160a01b0316604082015260600190565b600085825284602083015283604083015260806060830152825160808084015262000af861010084018262000a3e565b60208501519150607f19808583030160a086015262000b18828462000a3e565b60408701519350818682030160c087015262000b35818562000a3e565b92505060608601519250808583030160e08601525062000b56818362000a3e565b98975050505050505050565b60008482528360208301526060604083015262000b8460608301845162000a2f565b602083015162000b98608084018262000a22565b50604083015160a0830152606083015160c0830152608083015160e083015260a0830151610100818185015260c08501519150610120828186015260e086015192506101408381870152828701516101608701528187015193508061018087015250505062000c0c6101a084018262000a3e565b9695505050505050565b928352602083019190915260ff16604082015260600190565b60008382526040602083015282516006811062000c4857fe5b6040830152602083015160608084015262000c6760a084018262000a3e565b6040850151848203603f19016080860152915062000c0c818362000a3e565b91825260ff16602082015260400190565b60405181810167ffffffffffffffff8111828210171562000cb757600080fd5b60405291905056fe60806040523480156200001157600080fd5b506040516200257438038062002574833981016040819052620000349162000259565b8160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007957506000815160058111156200007757fe5b115b80156200009357506005815160058111156200009157fe5b105b8015620000a557506000816020015151115b8015620000b757506000816040015151115b620000c157600080fd5b600082905580516001805460ff191681836005811115620000de57fe5b02179055506020808201518051620000fb92600292019062000125565b5060408101518051620001179160039160209091019062000125565b50506000600655506200033d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200016857805160ff191683800117855562000198565b8280016001018555821562000198579182015b82811115620001985782518255916020019190600101906200017b565b50620001a6929150620001aa565b5090565b620001c791905b80821115620001a65760008155600101620001b1565b90565b600082601f830112620001db578081fd5b81516001600160401b03811115620001f1578182fd5b602062000207601f8301601f1916820162000316565b925081835284818386010111156200021e57600080fd5b60005b828110156200023e57848101820151848201830152810162000221565b82811115620002505760008284860101525b50505092915050565b600080604083850312156200026c578182fd5b825160208401519092506001600160401b03808211156200028b578283fd5b818501606081880312156200029e578384fd5b620002aa606062000316565b9250805160068110620002bb578485fd5b8352602081015182811115620002cf578485fd5b620002dd88828401620001ca565b602085015250604081015182811115620002f5578485fd5b6200030388828401620001ca565b6040850152505050809150509250929050565b6040518181016001600160401b03811182821017156200033557600080fd5b604052919050565b612227806200034d6000396000f3fe60806040523480156200001157600080fd5b5060043610620000a05760003560e01c80638de654ba116200006f5780638de654ba1462000120578063bc5993411462000137578063dcb5ad27146200014e578063ddca14621462000165578063e2aaa459146200017c57620000a0565b806367e0badb14620000a55780636d4ce63c14620000c757806382f9248214620000e15780638896dca81462000107575b600080fd5b620000af62000193565b604051620000be919062000b82565b60405180910390f35b620000d162000199565b604051620000be92919062000c35565b620000f8620000f236600462000819565b62000323565b604051620000be919062000b6e565b6200011e620001183660046200083b565b6200037c565b005b6200011e6200013136600462000ab5565b62000469565b620000af6200014836600462000aeb565b620004c6565b6200011e6200015f36600462000880565b62000500565b6200011e6200017636600462000984565b620005b8565b6200011e6200018d36600462000a81565b6200063d565b60065490565b6000620001a562000753565b6000546040805160608101909152600180549091908290829060ff166005811115620001cd57fe5b6005811115620001d957fe5b8152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156200027a5780601f106200024e576101008083540402835291602001916200027a565b820191906000526020600020905b8154815290600101906020018083116200025c57829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f81018390048302850183019091528084529381019390830182828015620003105780601f10620002e45761010080835404028352916020019162000310565b820191906000526020600020905b815481529060010190602001808311620002f257829003601f168201915b5050505050815250509050915091509091565b60008054831480156200034f57506000828152600460205260409020600101546001600160a01b031615155b6200035957600080fd5b506000818152600460205260409020600101546001600160a01b03165b92915050565b8260001a60f81b6001600160f81b031916158015906200039d575060005483145b8015620003b957508160001a60f81b6001600160f81b03191615155b8015620003df57506000828152600460205260409020600101546001600160a01b031615155b620003e957600080fd5b60008281526004602081905260409182902060010154915163d00c66af60e01b81526001600160a01b0390921691829163d00c66af916200042f91879187910162000b8b565b600060405180830381600087803b1580156200044a57600080fd5b505af11580156200045f573d6000803e3d6000fd5b5050505050505050565b600054821480156200047d575060ff811615155b80156200048d5750600560ff8216105b6200049757600080fd5b8060ff166005811115620004a757fe5b6001805460ff191681836005811115620004bd57fe5b02179055505050565b6006546000908210620004d857600080fd5b6005805483908110620004e757fe5b9060005260206000209060020201600001549050919050565b600054841480156200052b57506000838152600460205260409020600101546001600160a01b031615155b6200053557600080fd5b600083815260046020819052604091829020600101549151631ea71c6b60e31b81526001600160a01b0390921691829163f538e358916200057d918891889188910162000ba2565b600060405180830381600087803b1580156200059857600080fd5b505af1158015620005ad573d6000803e3d6000fd5b505050505050505050565b60005483148015620005d957508160001a60f81b6001600160f81b03191615155b620005e357600080fd5b60008282604051620005f59062000775565b6200060292919062000c96565b604051809103906000f0801580156200061f573d6000803e3d6000fd5b509050620006366004848363ffffffff620006b816565b5050505050565b600054831480156200066857506000828152600460205260409020600101546001600160a01b031615155b6200067257600080fd5b6000828152600460208190526040918290206001015491516346f32a5d60e11b81526001600160a01b03909216918291638de654ba916200042f91879187910162000d43565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b0385161790558015620006f85760019150506200074c565b50600180850180549182018082556000868152602088905260409020819055859190839081106200072557fe5b600091825260208220600291820201929092559086018054600101905591506200074c9050565b9392505050565b6040805160608101909152806000815260200160608152602001606081525090565b6114758062000d7d83390190565b80356001600160a01b03811681146200037657600080fd5b8035600981106200037657600080fd5b600082601f830112620007bc578081fd5b813567ffffffffffffffff811115620007d3578182fd5b620007e8601f8201601f191660200162000d54565b91508082528360208285010111156200080057600080fd5b8060208401602084013760009082016020015292915050565b600080604083850312156200082c578182fd5b50508035926020909101359150565b60008060006060848603121562000850578081fd5b833592506020840135915060408401356001600160a01b038116811462000875578182fd5b809150509250925092565b6000806000806080858703121562000896578081fd5b843593506020850135925060408501359150606085013567ffffffffffffffff80821115620008c3578283fd5b8187016080818a031215620008d6578384fd5b620008e2608062000d54565b9250803582811115620008f3578485fd5b620009018a828401620007ab565b84525060208101358281111562000916578485fd5b620009248a828401620007ab565b6020850152506040810135828111156200093c578485fd5b6200094a8a828401620007ab565b60408501525060608101358281111562000962578485fd5b620009708a828401620007ab565b606085015250959894975092955093505050565b60008060006060848603121562000999578283fd5b8335925060208401359150604084013567ffffffffffffffff80821115620009bf578283fd5b610140918601808803831315620009d4578384fd5b620009df8362000d54565b620009eb89836200079b565b8152620009fc896020840162000783565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c082013560c082015260e082013560e0820152610100935083820135848201526101209350838201358381111562000a60578586fd5b62000a6e8a828501620007ab565b8583015250809450505050509250925092565b60008060006060848603121562000a96578283fd5b8335925060208401359150604084013560ff8116811462000875578182fd5b6000806040838503121562000ac8578182fd5b82359150602083013560ff8116811462000ae0578182fd5b809150509250929050565b60006020828403121562000afd578081fd5b5035919050565b6001600160a01b03169052565b6009811062000b1c57fe5b9052565b60008151808452815b8181101562000b475760208185018101518683018201520162000b29565b8181111562000b595782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b90815260200190565b9182526001600160a01b0316602082015260400190565b60008482528360208301526060604083015282516080606084015262000bcc60e084018262000b20565b60208501519150605f198085830301608086015262000bec828462000b20565b60408701519350818682030160a087015262000c09818562000b20565b92505060608601519250808583030160c08601525062000c2a818362000b20565b979650505050505050565b60008382526040602083015282516006811062000c4e57fe5b6040830152602083015160608084015262000c6d60a084018262000b20565b6040850151848203603f19016080860152915062000c8c818362000b20565b9695505050505050565b60008382526040602083015262000cb260408301845162000b11565b602083015162000cc6606084018262000b04565b5060408301516080830152606083015160a0830152608083015160c083015260a083015160e083015260c0830151610100818185015260e0850151915061012082818601528186015192506101409150828286015280860151925050806101608501525062000d3a61018084018262000b20565b95945050505050565b91825260ff16602082015260400190565b60405181810167ffffffffffffffff8111828210171562000d7457600080fd5b60405291905056fe60806040523480156200001157600080fd5b5060405162001475380380620014758339810160408190526200003491620002c9565b8160001a60f81b6001600160f81b031916158015906200006157506000815160088111156200005f57fe5b115b80156200007b57506008815160088111156200007957fe5b105b80156200009b575060c081015160001a60f81b6001600160f81b03191615155b8015620000ab575060e081015115155b8015620000bc575061010081015115155b620000c657600080fd5b600082905580516001805460ff191681836008811115620000e357fe5b021790555060208181015160018054610100600160a81b0319166101006001600160a01b03909316830217905560c083015160065560408301516002556060830151600355608083015160045560a083015160055560e083015160075582015160085561012082015180516200015e92600992019062000167565b505050620003e4565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001aa57805160ff1916838001178555620001da565b82800160010185558215620001da579182015b82811115620001da578251825591602001919060010190620001bd565b50620001e8929150620001ec565b5090565b6200020991905b80821115620001e85760008155600101620001f3565b90565b80516001600160a01b03811681146200022457600080fd5b92915050565b8051600981106200022457600080fd5b600082601f8301126200024b578081fd5b81516001600160401b0381111562000261578182fd5b602062000277601f8301601f19168201620003bd565b925081835284818386010111156200028e57600080fd5b60005b82811015620002ae57848101820151848201830152810162000291565b82811115620002c05760008284860101525b50505092915050565b60008060408385031215620002dc578182fd5b825160208401519092506001600160401b0380821115620002fb578283fd5b61014091850180870383131562000310578384fd5b6200031b83620003bd565b6200032788836200022a565b81526200033888602084016200020c565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c082015160c082015260e082015160e082015261010093508382015184820152610120935083820151838111156200039c578586fd5b620003aa898285016200023a565b8583015250809450505050509250929050565b6040518181016001600160401b0381118282101715620003dc57600080fd5b604052919050565b61108180620003f46000396000f3fe60806040523480156200001157600080fd5b5060043610620000885760003560e01c80638de654ba11620000635780638de654ba14620000ef578063bc5993411462000108578063d00c66af146200011f578063f538e35814620001365762000088565b806352c43d9f146200008d57806367e0badb14620000bc5780636d4ce63c14620000d5575b600080fd5b620000a46200009e36600462000603565b6200014d565b604051620000b39190620007ce565b60405180910390f35b620000c6620001a4565b604051620000b39190620007e2565b620000df620001aa565b604051620000b392919062000877565b620001066200010036600462000720565b620002ef565b005b620000c6620001193660046200074b565b6200034c565b6200010662000130366004620005c7565b62000386565b620001066200014736600462000625565b620003d3565b60008054831480156200017957506000828152600a60205260409020600101546001600160a01b031615155b6200018357600080fd5b506000908152600a60205260409020600101546001600160a01b0316919050565b600c5490565b6000620001b6620004f3565b600054604080516101408101909152600180549091908290829060ff166008811115620001df57fe5b6008811115620001eb57fe5b81528154610100908190046001600160a01b0316602080840191909152600180850154604080860191909152600280870154606087015260038701546080870152600487015460a0870152600587015460c0870152600687015460e08701526007870154858701526008870180548351948116159096026000190190951604601f810184900484028301840190915280825261012090940193909291830182828015620002dc5780601f10620002b057610100808354040283529160200191620002dc565b820191906000526020600020905b815481529060010190602001808311620002be57829003601f168201915b5050505050815250509050915091509091565b6000548214801562000303575060ff811615155b8015620003135750600860ff8216105b6200031d57600080fd5b8060ff1660088111156200032d57fe5b6001805460ff1916818360088111156200034357fe5b02179055505050565b600c5460009082106200035e57600080fd5b600b8054839081106200036d57fe5b9060005260206000209060020201600001549050919050565b60005482148015620003a057506001600160a01b03811615155b620003aa57600080fd5b600180546001600160a01b0390921661010002610100600160a81b031990921691909117905550565b60005483148015620003f457508160001a60f81b6001600160f81b03191615155b620003fe57600080fd5b6000828260405162000410906200054b565b6200041d929190620007eb565b604051809103906000f0801580156200043a573d6000803e3d6000fd5b50905062000451600a848363ffffffff6200045816565b5050505050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000498576001915050620004ec565b5060018085018054918201808255600086815260208890526040902081905585919083908110620004c557fe5b60009182526020822060029182020192909255908601805460010190559150620004ec9050565b9392505050565b60408051610140810190915280600081526000602082018190526040820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301919091526101209091015290565b6106ff806200094d83390190565b600082601f8301126200056a578081fd5b813567ffffffffffffffff81111562000581578182fd5b62000596601f8201601f191660200162000924565b9150808252836020828501011115620005ae57600080fd5b8060208401602084013760009082016020015292915050565b60008060408385031215620005da578182fd5b8235915060208301356001600160a01b0381168114620005f8578182fd5b809150509250929050565b6000806040838503121562000616578182fd5b50508035926020909101359150565b6000806000606084860312156200063a578081fd5b8335925060208401359150604084013567ffffffffffffffff8082111562000660578283fd5b8186016080818903121562000673578384fd5b6200067f608062000924565b925080358281111562000690578485fd5b6200069e8982840162000559565b845250602081013582811115620006b3578485fd5b620006c18982840162000559565b602085015250604081013582811115620006d9578485fd5b620006e78982840162000559565b604085015250606081013582811115620006ff578485fd5b6200070d8982840162000559565b6060850152505050809150509250925092565b6000806040838503121562000733578182fd5b82359150602083013560ff81168114620005f8578182fd5b6000602082840312156200075d578081fd5b5035919050565b6001600160a01b03169052565b600981106200077c57fe5b9052565b60008151808452815b81811015620007a75760208185018101518683018201520162000789565b81811115620007b95782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b90815260200190565b6000838252604060208301528251608060408401526200080f60c084018262000780565b60208501519150603f19808583030160608601526200082f828462000780565b60408701519350818682030160808701526200084c818562000780565b92505060608601519250808583030160a0860152506200086d818362000780565b9695505050505050565b6000838252604060208301526200089360408301845162000771565b6020830151620008a7606084018262000764565b5060408301516080830152606083015160a0830152608083015160c083015260a083015160e083015260c0830151610100818185015260e085015191506101208281860152818601519250610140915082828601528086015192505080610160850152506200091b61018084018262000780565b95945050505050565b60405181810167ffffffffffffffff811182821017156200094457600080fd5b60405291905056fe608060405234801561001057600080fd5b506040516106ff3803806106ff83398101604081905261002f916101f1565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001661006157600080fd5b60008290558051805161007c916001916020909101906100cf565b5060208082015180516100939260029201906100cf565b50604081015180516100ad916003916020909101906100cf565b50606081015180516100c7916004916020909101906100cf565b5050506102f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011057805160ff191683800117855561013d565b8280016001018555821561013d579182015b8281111561013d578251825591602001919060010190610122565b5061014992915061014d565b5090565b61016791905b808211156101495760008155600101610153565b90565b600082601f83011261017a578081fd5b81516001600160401b0381111561018f578182fd5b60206101a3601f8301601f191682016102d2565b925081835284818386010111156101b957600080fd5b60005b828110156101d75784810182015184820183015281016101bc565b828111156101e85760008284860101525b50505092915050565b60008060408385031215610203578182fd5b825160208401519092506001600160401b0380821115610221578283fd5b81850160808188031215610233578384fd5b61023d60806102d2565b925080518281111561024d578485fd5b6102598882840161016a565b84525060208101518281111561026d578485fd5b6102798882840161016a565b602085015250604081015182811115610290578485fd5b61029c8882840161016a565b6040850152506060810151828111156102b3578485fd5b6102bf8882840161016a565b6060850152505050809150509250929050565b6040518181016001600160401b03811182821017156102f057600080fd5b604052919050565b6103f8806103076000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80636d4ce63c14610030575b600080fd5b61003861004f565b60405161004692919061033e565b60405180910390f35b60006100596102cb565b600054604080516001805460a06020601f60026000196101008688161502019094169390930492830181900402840181019094526080830181815291938492849290918491908401828280156100f05780601f106100c5576101008083540402835291602001916100f0565b820191906000526020600020905b8154815290600101906020018083116100d357829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101925780601f1061016757610100808354040283529160200191610192565b820191906000526020600020905b81548152906001019060200180831161017557829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156102245780601f106101f957610100808354040283529160200191610224565b820191906000526020600020905b81548152906001019060200180831161020757829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156102b85780601f1061028d576101008083540402835291602001916102b8565b820191906000526020600020905b81548152906001019060200180831161029b57829003601f168201915b5050505050815250509050915091509091565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b60008151808452815b81811015610318576020818501810151868301820152016102fc565b818111156103295782602083870101525b50601f01601f19169290920160200192915050565b60008382526040602083015282516080604084015261036060c08401826102f3565b60208501519150603f198085830301606086015261037e82846102f3565b604087015193508186820301608087015261039981856102f3565b92505060608601519250808583030160a0860152506103b881836102f3565b969550505050505056fea264697066735822122051c640576777bbdc856893273c3ae7351824716c50548076450668cfda633c9364736f6c63430006080033a2646970667358221220091b9b56f546f270a7ac47f850fc9812ce3c6b78ff3d38b959d6fdbce031155964736f6c63430006080033a26469706673582212206cf0adc33f806ba7b2b53ac27f512310dcfcdd752cc1fb117ef94f8866c2abd764736f6c63430006080033a26469706673582212205e9767c596aef312f883f5fb348e51947c8d8f11b30520a3068d57e0b196253d64736f6c63430006080033"

// DeployJobs deploys a new Ethereum contract, binding an instance of Jobs to it.
func DeployJobs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Jobs, error) {
	parsed, err := abi.JSON(strings.NewReader(JobsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JobsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jobs{JobsCaller: JobsCaller{contract: contract}, JobsTransactor: JobsTransactor{contract: contract}, JobsFilterer: JobsFilterer{contract: contract}}, nil
}

// Jobs is an auto generated Go binding around an Ethereum contract.
type Jobs struct {
	JobsCaller     // Read-only binding to the contract
	JobsTransactor // Write-only binding to the contract
	JobsFilterer   // Log filterer for contract events
}

// JobsCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobsSession struct {
	Contract     *Jobs             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobsCallerSession struct {
	Contract *JobsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JobsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobsTransactorSession struct {
	Contract     *JobsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobsRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobsRaw struct {
	Contract *Jobs // Generic contract binding to access the raw methods on
}

// JobsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobsCallerRaw struct {
	Contract *JobsCaller // Generic read-only contract binding to access the raw methods on
}

// JobsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobsTransactorRaw struct {
	Contract *JobsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobs creates a new instance of Jobs, bound to a specific deployed contract.
func NewJobs(address common.Address, backend bind.ContractBackend) (*Jobs, error) {
	contract, err := bindJobs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jobs{JobsCaller: JobsCaller{contract: contract}, JobsTransactor: JobsTransactor{contract: contract}, JobsFilterer: JobsFilterer{contract: contract}}, nil
}

// NewJobsCaller creates a new read-only instance of Jobs, bound to a specific deployed contract.
func NewJobsCaller(address common.Address, caller bind.ContractCaller) (*JobsCaller, error) {
	contract, err := bindJobs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobsCaller{contract: contract}, nil
}

// NewJobsTransactor creates a new write-only instance of Jobs, bound to a specific deployed contract.
func NewJobsTransactor(address common.Address, transactor bind.ContractTransactor) (*JobsTransactor, error) {
	contract, err := bindJobs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobsTransactor{contract: contract}, nil
}

// NewJobsFilterer creates a new log filterer instance of Jobs, bound to a specific deployed contract.
func NewJobsFilterer(address common.Address, filterer bind.ContractFilterer) (*JobsFilterer, error) {
	contract, err := bindJobs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobsFilterer{contract: contract}, nil
}

// bindJobs binds a generic wrapper to an already deployed contract.
func bindJobs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JobsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobs *JobsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Jobs.Contract.JobsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobs *JobsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobs.Contract.JobsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobs *JobsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobs.Contract.JobsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobs *JobsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Jobs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobs *JobsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobs *JobsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobs.Contract.contract.Transact(opts, method, params...)
}

// GetJobContract is a free data retrieval call binding the contract method 0xc918e294.
//
// Solidity: function getJobContract(bytes32 _jobRef) view returns(address)
func (_Jobs *JobsCaller) GetJobContract(opts *bind.CallOpts, _jobRef [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Jobs.contract.Call(opts, out, "getJobContract", _jobRef)
	return *ret0, err
}

// GetJobContract is a free data retrieval call binding the contract method 0xc918e294.
//
// Solidity: function getJobContract(bytes32 _jobRef) view returns(address)
func (_Jobs *JobsSession) GetJobContract(_jobRef [32]byte) (common.Address, error) {
	return _Jobs.Contract.GetJobContract(&_Jobs.CallOpts, _jobRef)
}

// GetJobContract is a free data retrieval call binding the contract method 0xc918e294.
//
// Solidity: function getJobContract(bytes32 _jobRef) view returns(address)
func (_Jobs *JobsCallerSession) GetJobContract(_jobRef [32]byte) (common.Address, error) {
	return _Jobs.Contract.GetJobContract(&_Jobs.CallOpts, _jobRef)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Jobs *JobsCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Jobs.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Jobs *JobsSession) GetNum() (*big.Int, error) {
	return _Jobs.Contract.GetNum(&_Jobs.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Jobs *JobsCallerSession) GetNum() (*big.Int, error) {
	return _Jobs.Contract.GetNum(&_Jobs.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Jobs *JobsCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Jobs.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Jobs *JobsSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Jobs.Contract.GetReference(&_Jobs.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Jobs *JobsCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Jobs.Contract.GetReference(&_Jobs.CallOpts, _index)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_Jobs *JobsTransactor) AddInfo(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "addInfo", _jobRef, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_Jobs *JobsSession) AddInfo(_jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _Jobs.Contract.AddInfo(&_Jobs.TransactOpts, _jobRef, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xdcb5ad27.
//
// Solidity: function addInfo(bytes32 _jobRef, bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_Jobs *JobsTransactorSession) AddInfo(_jobRef [32]byte, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _Jobs.Contract.AddInfo(&_Jobs.TransactOpts, _jobRef, _workRef, _infoRef, _info)
}

// AddJob is a paid mutator transaction binding the contract method 0x8b7a9f98.
//
// Solidity: function addJob(bytes32 _jobRef, (uint8,string,string) _job) returns()
func (_Jobs *JobsTransactor) AddJob(opts *bind.TransactOpts, _jobRef [32]byte, _job Job) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "addJob", _jobRef, _job)
}

// AddJob is a paid mutator transaction binding the contract method 0x8b7a9f98.
//
// Solidity: function addJob(bytes32 _jobRef, (uint8,string,string) _job) returns()
func (_Jobs *JobsSession) AddJob(_jobRef [32]byte, _job Job) (*types.Transaction, error) {
	return _Jobs.Contract.AddJob(&_Jobs.TransactOpts, _jobRef, _job)
}

// AddJob is a paid mutator transaction binding the contract method 0x8b7a9f98.
//
// Solidity: function addJob(bytes32 _jobRef, (uint8,string,string) _job) returns()
func (_Jobs *JobsTransactorSession) AddJob(_jobRef [32]byte, _job Job) (*types.Transaction, error) {
	return _Jobs.Contract.AddJob(&_Jobs.TransactOpts, _jobRef, _job)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_Jobs *JobsTransactor) AddWork(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "addWork", _jobRef, _workRef, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_Jobs *JobsSession) AddWork(_jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _Jobs.Contract.AddWork(&_Jobs.TransactOpts, _jobRef, _workRef, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0xddca1462.
//
// Solidity: function addWork(bytes32 _jobRef, bytes32 _workRef, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string) _work) returns()
func (_Jobs *JobsTransactorSession) AddWork(_jobRef [32]byte, _workRef [32]byte, _work Work) (*types.Transaction, error) {
	return _Jobs.Contract.AddWork(&_Jobs.TransactOpts, _jobRef, _workRef, _work)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobRef, uint8 _status) returns()
func (_Jobs *JobsTransactor) SetJobStatus(opts *bind.TransactOpts, _jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "setJobStatus", _jobRef, _status)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobRef, uint8 _status) returns()
func (_Jobs *JobsSession) SetJobStatus(_jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _Jobs.Contract.SetJobStatus(&_Jobs.TransactOpts, _jobRef, _status)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobRef, uint8 _status) returns()
func (_Jobs *JobsTransactorSession) SetJobStatus(_jobRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _Jobs.Contract.SetJobStatus(&_Jobs.TransactOpts, _jobRef, _status)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_Jobs *JobsTransactor) SetSubbie(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "setSubbie", _jobRef, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_Jobs *JobsSession) SetSubbie(_jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _Jobs.Contract.SetSubbie(&_Jobs.TransactOpts, _jobRef, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0x8896dca8.
//
// Solidity: function setSubbie(bytes32 _jobRef, bytes32 _workRef, address _subbie) returns()
func (_Jobs *JobsTransactorSession) SetSubbie(_jobRef [32]byte, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _Jobs.Contract.SetSubbie(&_Jobs.TransactOpts, _jobRef, _workRef, _subbie)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_Jobs *JobsTransactor) SetWorkStatus(opts *bind.TransactOpts, _jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "setWorkStatus", _jobRef, _workRef, _status)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_Jobs *JobsSession) SetWorkStatus(_jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _Jobs.Contract.SetWorkStatus(&_Jobs.TransactOpts, _jobRef, _workRef, _status)
}

// SetWorkStatus is a paid mutator transaction binding the contract method 0xe2aaa459.
//
// Solidity: function setWorkStatus(bytes32 _jobRef, bytes32 _workRef, uint8 _status) returns()
func (_Jobs *JobsTransactorSession) SetWorkStatus(_jobRef [32]byte, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _Jobs.Contract.SetWorkStatus(&_Jobs.TransactOpts, _jobRef, _workRef, _status)
}

// WorkNodeABI is the input ABI used to generate the binding from.
const WorkNodeABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumWorkStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"subbie\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"startDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sponsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWork\",\"name\":\"_work\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structExtraInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"addInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumWorkStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"subbie\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"startDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endDate\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"endTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sponsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_infoRef\",\"type\":\"bytes32\"}],\"name\":\"getInfoContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workRef\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_subbie\",\"type\":\"address\"}],\"name\":\"setSubbie\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WorkNodeFuncSigs maps the 4-byte function signature to its string representation.
var WorkNodeFuncSigs = map[string]string{
	"f538e358": "addInfo(bytes32,bytes32,(string,string,string,string))",
	"6d4ce63c": "get()",
	"52c43d9f": "getInfoContract(bytes32,bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"8de654ba": "setStatus(bytes32,uint8)",
	"d00c66af": "setSubbie(bytes32,address)",
}

// WorkNodeBin is the compiled bytecode used for deploying new contracts.
var WorkNodeBin = "0x60806040523480156200001157600080fd5b5060405162001475380380620014758339810160408190526200003491620002c9565b8160001a60f81b6001600160f81b031916158015906200006157506000815160088111156200005f57fe5b115b80156200007b57506008815160088111156200007957fe5b105b80156200009b575060c081015160001a60f81b6001600160f81b03191615155b8015620000ab575060e081015115155b8015620000bc575061010081015115155b620000c657600080fd5b600082905580516001805460ff191681836008811115620000e357fe5b021790555060208181015160018054610100600160a81b0319166101006001600160a01b03909316830217905560c083015160065560408301516002556060830151600355608083015160045560a083015160055560e083015160075582015160085561012082015180516200015e92600992019062000167565b505050620003e4565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001aa57805160ff1916838001178555620001da565b82800160010185558215620001da579182015b82811115620001da578251825591602001919060010190620001bd565b50620001e8929150620001ec565b5090565b6200020991905b80821115620001e85760008155600101620001f3565b90565b80516001600160a01b03811681146200022457600080fd5b92915050565b8051600981106200022457600080fd5b600082601f8301126200024b578081fd5b81516001600160401b0381111562000261578182fd5b602062000277601f8301601f19168201620003bd565b925081835284818386010111156200028e57600080fd5b60005b82811015620002ae57848101820151848201830152810162000291565b82811115620002c05760008284860101525b50505092915050565b60008060408385031215620002dc578182fd5b825160208401519092506001600160401b0380821115620002fb578283fd5b61014091850180870383131562000310578384fd5b6200031b83620003bd565b6200032788836200022a565b81526200033888602084016200020c565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c082015160c082015260e082015160e082015261010093508382015184820152610120935083820151838111156200039c578586fd5b620003aa898285016200023a565b8583015250809450505050509250929050565b6040518181016001600160401b0381118282101715620003dc57600080fd5b604052919050565b61108180620003f46000396000f3fe60806040523480156200001157600080fd5b5060043610620000885760003560e01c80638de654ba11620000635780638de654ba14620000ef578063bc5993411462000108578063d00c66af146200011f578063f538e35814620001365762000088565b806352c43d9f146200008d57806367e0badb14620000bc5780636d4ce63c14620000d5575b600080fd5b620000a46200009e36600462000603565b6200014d565b604051620000b39190620007ce565b60405180910390f35b620000c6620001a4565b604051620000b39190620007e2565b620000df620001aa565b604051620000b392919062000877565b620001066200010036600462000720565b620002ef565b005b620000c6620001193660046200074b565b6200034c565b6200010662000130366004620005c7565b62000386565b620001066200014736600462000625565b620003d3565b60008054831480156200017957506000828152600a60205260409020600101546001600160a01b031615155b6200018357600080fd5b506000908152600a60205260409020600101546001600160a01b0316919050565b600c5490565b6000620001b6620004f3565b600054604080516101408101909152600180549091908290829060ff166008811115620001df57fe5b6008811115620001eb57fe5b81528154610100908190046001600160a01b0316602080840191909152600180850154604080860191909152600280870154606087015260038701546080870152600487015460a0870152600587015460c0870152600687015460e08701526007870154858701526008870180548351948116159096026000190190951604601f810184900484028301840190915280825261012090940193909291830182828015620002dc5780601f10620002b057610100808354040283529160200191620002dc565b820191906000526020600020905b815481529060010190602001808311620002be57829003601f168201915b5050505050815250509050915091509091565b6000548214801562000303575060ff811615155b8015620003135750600860ff8216105b6200031d57600080fd5b8060ff1660088111156200032d57fe5b6001805460ff1916818360088111156200034357fe5b02179055505050565b600c5460009082106200035e57600080fd5b600b8054839081106200036d57fe5b9060005260206000209060020201600001549050919050565b60005482148015620003a057506001600160a01b03811615155b620003aa57600080fd5b600180546001600160a01b0390921661010002610100600160a81b031990921691909117905550565b60005483148015620003f457508160001a60f81b6001600160f81b03191615155b620003fe57600080fd5b6000828260405162000410906200054b565b6200041d929190620007eb565b604051809103906000f0801580156200043a573d6000803e3d6000fd5b50905062000451600a848363ffffffff6200045816565b5050505050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000498576001915050620004ec565b5060018085018054918201808255600086815260208890526040902081905585919083908110620004c557fe5b60009182526020822060029182020192909255908601805460010190559150620004ec9050565b9392505050565b60408051610140810190915280600081526000602082018190526040820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301919091526101209091015290565b6106ff806200094d83390190565b600082601f8301126200056a578081fd5b813567ffffffffffffffff81111562000581578182fd5b62000596601f8201601f191660200162000924565b9150808252836020828501011115620005ae57600080fd5b8060208401602084013760009082016020015292915050565b60008060408385031215620005da578182fd5b8235915060208301356001600160a01b0381168114620005f8578182fd5b809150509250929050565b6000806040838503121562000616578182fd5b50508035926020909101359150565b6000806000606084860312156200063a578081fd5b8335925060208401359150604084013567ffffffffffffffff8082111562000660578283fd5b8186016080818903121562000673578384fd5b6200067f608062000924565b925080358281111562000690578485fd5b6200069e8982840162000559565b845250602081013582811115620006b3578485fd5b620006c18982840162000559565b602085015250604081013582811115620006d9578485fd5b620006e78982840162000559565b604085015250606081013582811115620006ff578485fd5b6200070d8982840162000559565b6060850152505050809150509250925092565b6000806040838503121562000733578182fd5b82359150602083013560ff81168114620005f8578182fd5b6000602082840312156200075d578081fd5b5035919050565b6001600160a01b03169052565b600981106200077c57fe5b9052565b60008151808452815b81811015620007a75760208185018101518683018201520162000789565b81811115620007b95782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b90815260200190565b6000838252604060208301528251608060408401526200080f60c084018262000780565b60208501519150603f19808583030160608601526200082f828462000780565b60408701519350818682030160808701526200084c818562000780565b92505060608601519250808583030160a0860152506200086d818362000780565b9695505050505050565b6000838252604060208301526200089360408301845162000771565b6020830151620008a7606084018262000764565b5060408301516080830152606083015160a0830152608083015160c083015260a083015160e083015260c0830151610100818185015260e085015191506101208281860152818601519250610140915082828601528086015192505080610160850152506200091b61018084018262000780565b95945050505050565b60405181810167ffffffffffffffff811182821017156200094457600080fd5b60405291905056fe608060405234801561001057600080fd5b506040516106ff3803806106ff83398101604081905261002f916101f1565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001661006157600080fd5b60008290558051805161007c916001916020909101906100cf565b5060208082015180516100939260029201906100cf565b50604081015180516100ad916003916020909101906100cf565b50606081015180516100c7916004916020909101906100cf565b5050506102f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011057805160ff191683800117855561013d565b8280016001018555821561013d579182015b8281111561013d578251825591602001919060010190610122565b5061014992915061014d565b5090565b61016791905b808211156101495760008155600101610153565b90565b600082601f83011261017a578081fd5b81516001600160401b0381111561018f578182fd5b60206101a3601f8301601f191682016102d2565b925081835284818386010111156101b957600080fd5b60005b828110156101d75784810182015184820183015281016101bc565b828111156101e85760008284860101525b50505092915050565b60008060408385031215610203578182fd5b825160208401519092506001600160401b0380821115610221578283fd5b81850160808188031215610233578384fd5b61023d60806102d2565b925080518281111561024d578485fd5b6102598882840161016a565b84525060208101518281111561026d578485fd5b6102798882840161016a565b602085015250604081015182811115610290578485fd5b61029c8882840161016a565b6040850152506060810151828111156102b3578485fd5b6102bf8882840161016a565b6060850152505050809150509250929050565b6040518181016001600160401b03811182821017156102f057600080fd5b604052919050565b6103f8806103076000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80636d4ce63c14610030575b600080fd5b61003861004f565b60405161004692919061033e565b60405180910390f35b60006100596102cb565b600054604080516001805460a06020601f60026000196101008688161502019094169390930492830181900402840181019094526080830181815291938492849290918491908401828280156100f05780601f106100c5576101008083540402835291602001916100f0565b820191906000526020600020905b8154815290600101906020018083116100d357829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101925780601f1061016757610100808354040283529160200191610192565b820191906000526020600020905b81548152906001019060200180831161017557829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156102245780601f106101f957610100808354040283529160200191610224565b820191906000526020600020905b81548152906001019060200180831161020757829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156102b85780601f1061028d576101008083540402835291602001916102b8565b820191906000526020600020905b81548152906001019060200180831161029b57829003601f168201915b5050505050815250509050915091509091565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b60008151808452815b81811015610318576020818501810151868301820152016102fc565b818111156103295782602083870101525b50601f01601f19169290920160200192915050565b60008382526040602083015282516080604084015261036060c08401826102f3565b60208501519150603f198085830301606086015261037e82846102f3565b604087015193508186820301608087015261039981856102f3565b92505060608601519250808583030160a0860152506103b881836102f3565b969550505050505056fea264697066735822122051c640576777bbdc856893273c3ae7351824716c50548076450668cfda633c9364736f6c63430006080033a2646970667358221220091b9b56f546f270a7ac47f850fc9812ce3c6b78ff3d38b959d6fdbce031155964736f6c63430006080033"

// DeployWorkNode deploys a new Ethereum contract, binding an instance of WorkNode to it.
func DeployWorkNode(auth *bind.TransactOpts, backend bind.ContractBackend, _workRef [32]byte, _work Work) (common.Address, *types.Transaction, *WorkNode, error) {
	parsed, err := abi.JSON(strings.NewReader(WorkNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WorkNodeBin), backend, _workRef, _work)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WorkNode{WorkNodeCaller: WorkNodeCaller{contract: contract}, WorkNodeTransactor: WorkNodeTransactor{contract: contract}, WorkNodeFilterer: WorkNodeFilterer{contract: contract}}, nil
}

// WorkNode is an auto generated Go binding around an Ethereum contract.
type WorkNode struct {
	WorkNodeCaller     // Read-only binding to the contract
	WorkNodeTransactor // Write-only binding to the contract
	WorkNodeFilterer   // Log filterer for contract events
}

// WorkNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkNodeSession struct {
	Contract     *WorkNode         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkNodeCallerSession struct {
	Contract *WorkNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WorkNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkNodeTransactorSession struct {
	Contract     *WorkNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WorkNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkNodeRaw struct {
	Contract *WorkNode // Generic contract binding to access the raw methods on
}

// WorkNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkNodeCallerRaw struct {
	Contract *WorkNodeCaller // Generic read-only contract binding to access the raw methods on
}

// WorkNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkNodeTransactorRaw struct {
	Contract *WorkNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkNode creates a new instance of WorkNode, bound to a specific deployed contract.
func NewWorkNode(address common.Address, backend bind.ContractBackend) (*WorkNode, error) {
	contract, err := bindWorkNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkNode{WorkNodeCaller: WorkNodeCaller{contract: contract}, WorkNodeTransactor: WorkNodeTransactor{contract: contract}, WorkNodeFilterer: WorkNodeFilterer{contract: contract}}, nil
}

// NewWorkNodeCaller creates a new read-only instance of WorkNode, bound to a specific deployed contract.
func NewWorkNodeCaller(address common.Address, caller bind.ContractCaller) (*WorkNodeCaller, error) {
	contract, err := bindWorkNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkNodeCaller{contract: contract}, nil
}

// NewWorkNodeTransactor creates a new write-only instance of WorkNode, bound to a specific deployed contract.
func NewWorkNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkNodeTransactor, error) {
	contract, err := bindWorkNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkNodeTransactor{contract: contract}, nil
}

// NewWorkNodeFilterer creates a new log filterer instance of WorkNode, bound to a specific deployed contract.
func NewWorkNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkNodeFilterer, error) {
	contract, err := bindWorkNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkNodeFilterer{contract: contract}, nil
}

// bindWorkNode binds a generic wrapper to an already deployed contract.
func bindWorkNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WorkNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkNode *WorkNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WorkNode.Contract.WorkNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkNode *WorkNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkNode.Contract.WorkNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkNode *WorkNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkNode.Contract.WorkNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkNode *WorkNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WorkNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkNode *WorkNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkNode *WorkNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))
func (_WorkNode *WorkNodeCaller) Get(opts *bind.CallOpts) ([32]byte, Work, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(Work)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _WorkNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))
func (_WorkNode *WorkNodeSession) Get() ([32]byte, Work, error) {
	return _WorkNode.Contract.Get(&_WorkNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,address,bytes32,bytes32,bytes32,bytes32,bytes32,uint256,uint256,string))
func (_WorkNode *WorkNodeCallerSession) Get() ([32]byte, Work, error) {
	return _WorkNode.Contract.Get(&_WorkNode.CallOpts)
}

// GetInfoContract is a free data retrieval call binding the contract method 0x52c43d9f.
//
// Solidity: function getInfoContract(bytes32 _workRef, bytes32 _infoRef) view returns(address)
func (_WorkNode *WorkNodeCaller) GetInfoContract(opts *bind.CallOpts, _workRef [32]byte, _infoRef [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WorkNode.contract.Call(opts, out, "getInfoContract", _workRef, _infoRef)
	return *ret0, err
}

// GetInfoContract is a free data retrieval call binding the contract method 0x52c43d9f.
//
// Solidity: function getInfoContract(bytes32 _workRef, bytes32 _infoRef) view returns(address)
func (_WorkNode *WorkNodeSession) GetInfoContract(_workRef [32]byte, _infoRef [32]byte) (common.Address, error) {
	return _WorkNode.Contract.GetInfoContract(&_WorkNode.CallOpts, _workRef, _infoRef)
}

// GetInfoContract is a free data retrieval call binding the contract method 0x52c43d9f.
//
// Solidity: function getInfoContract(bytes32 _workRef, bytes32 _infoRef) view returns(address)
func (_WorkNode *WorkNodeCallerSession) GetInfoContract(_workRef [32]byte, _infoRef [32]byte) (common.Address, error) {
	return _WorkNode.Contract.GetInfoContract(&_WorkNode.CallOpts, _workRef, _infoRef)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_WorkNode *WorkNodeCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WorkNode.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_WorkNode *WorkNodeSession) GetNum() (*big.Int, error) {
	return _WorkNode.Contract.GetNum(&_WorkNode.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_WorkNode *WorkNodeCallerSession) GetNum() (*big.Int, error) {
	return _WorkNode.Contract.GetNum(&_WorkNode.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_WorkNode *WorkNodeCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WorkNode.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_WorkNode *WorkNodeSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _WorkNode.Contract.GetReference(&_WorkNode.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_WorkNode *WorkNodeCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _WorkNode.Contract.GetReference(&_WorkNode.CallOpts, _index)
}

// AddInfo is a paid mutator transaction binding the contract method 0xf538e358.
//
// Solidity: function addInfo(bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_WorkNode *WorkNodeTransactor) AddInfo(opts *bind.TransactOpts, _workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _WorkNode.contract.Transact(opts, "addInfo", _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xf538e358.
//
// Solidity: function addInfo(bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_WorkNode *WorkNodeSession) AddInfo(_workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _WorkNode.Contract.AddInfo(&_WorkNode.TransactOpts, _workRef, _infoRef, _info)
}

// AddInfo is a paid mutator transaction binding the contract method 0xf538e358.
//
// Solidity: function addInfo(bytes32 _workRef, bytes32 _infoRef, (string,string,string,string) _info) returns()
func (_WorkNode *WorkNodeTransactorSession) AddInfo(_workRef [32]byte, _infoRef [32]byte, _info ExtraInfo) (*types.Transaction, error) {
	return _WorkNode.Contract.AddInfo(&_WorkNode.TransactOpts, _workRef, _infoRef, _info)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _workRef, uint8 _status) returns()
func (_WorkNode *WorkNodeTransactor) SetStatus(opts *bind.TransactOpts, _workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _WorkNode.contract.Transact(opts, "setStatus", _workRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _workRef, uint8 _status) returns()
func (_WorkNode *WorkNodeSession) SetStatus(_workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _WorkNode.Contract.SetStatus(&_WorkNode.TransactOpts, _workRef, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x8de654ba.
//
// Solidity: function setStatus(bytes32 _workRef, uint8 _status) returns()
func (_WorkNode *WorkNodeTransactorSession) SetStatus(_workRef [32]byte, _status uint8) (*types.Transaction, error) {
	return _WorkNode.Contract.SetStatus(&_WorkNode.TransactOpts, _workRef, _status)
}

// SetSubbie is a paid mutator transaction binding the contract method 0xd00c66af.
//
// Solidity: function setSubbie(bytes32 _workRef, address _subbie) returns()
func (_WorkNode *WorkNodeTransactor) SetSubbie(opts *bind.TransactOpts, _workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _WorkNode.contract.Transact(opts, "setSubbie", _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0xd00c66af.
//
// Solidity: function setSubbie(bytes32 _workRef, address _subbie) returns()
func (_WorkNode *WorkNodeSession) SetSubbie(_workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _WorkNode.Contract.SetSubbie(&_WorkNode.TransactOpts, _workRef, _subbie)
}

// SetSubbie is a paid mutator transaction binding the contract method 0xd00c66af.
//
// Solidity: function setSubbie(bytes32 _workRef, address _subbie) returns()
func (_WorkNode *WorkNodeTransactorSession) SetSubbie(_workRef [32]byte, _subbie common.Address) (*types.Transaction, error) {
	return _WorkNode.Contract.SetSubbie(&_WorkNode.TransactOpts, _workRef, _subbie)
}
