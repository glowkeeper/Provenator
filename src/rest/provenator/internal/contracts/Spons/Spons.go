// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Spons

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

// CostModel is an auto generated low-level Go binding around an user-defined struct.
type CostModel struct {
	Unit        uint8
	Rate        *big.Int
	Description string
}

// PriceForMeasuredWorks is an auto generated low-level Go binding around an user-defined struct.
type PriceForMeasuredWorks struct {
	Unit          uint8
	PrimeCost     *big.Int
	LabourMinutes *big.Int
	LaboutCost    *big.Int
	Plant         *big.Int
	Material      *big.Int
	Description   string
}

// PricePerFunctionalUnit is an auto generated low-level Go binding around an user-defined struct.
type PricePerFunctionalUnit struct {
	Unit        uint8
	LowerBound  *big.Int
	UpperBound  *big.Int
	Description string
}

// PricePerSM is an auto generated low-level Go binding around an user-defined struct.
type PricePerSM struct {
	LowerBound  *big.Int
	UpperBound  *big.Int
	Description string
}

// CostNodeABI is the input ABI used to generate the binding from.
const CostNodeABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nodeRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCostModel\",\"name\":\"_node\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCostModel\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumPriceModels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CostNodeFuncSigs maps the 4-byte function signature to its string representation.
var CostNodeFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"15dae03e": "getType()",
}

// CostNodeBin is the compiled bytecode used for deploying new contracts.
var CostNodeBin = "0x608060405234801561001057600080fd5b5060405161054438038061054483398101604081905261002f91610184565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615801590610072575060008151601381111561007057fe5b115b801561008a575060138151601381111561008857fe5b105b80156100995750602081015115155b6100a257600080fd5b600082905580516001805460ff1916818360138111156100be57fe5b0217905550602080820151600255604082015180516100e19260039201906100e9565b5050506102db565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061012a57805160ff1916838001178555610157565b82800160010185558215610157579182015b8281111561015757825182559160200191906001019061013c565b50610163929150610167565b5090565b61018191905b80821115610163576000815560010161016d565b90565b60008060408385031215610196578182fd5b8251602080850151919350906001600160401b03808211156101b6578384fd5b818601606081890312156101c8578485fd5b6101d26060610252565b925080516101df816102cb565b835280840151848401526040810151828111156101fa578586fd5b019050601f8101871361020b578384fd5b805161021e61021982610278565b610252565b8181528885838501011115610231578586fd5b6102408286830187860161029b565b60408401525093969095509350505050565b6040518181016001600160401b038111828210171561027057600080fd5b604052919050565b60006001600160401b0382111561028d578081fd5b50601f01601f191660200190565b60005b838110156102b657818101518382015260200161029e565b838111156102c5576000848401525b50505050565b601481106102d857600080fd5b50565b61025a806102ea6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315dae03e1461003b5780636d4ce63c14610059575b600080fd5b61004361006f565b6040516100509190610210565b60405180910390f35b610061610074565b60405161005092919061018c565b600190565b600061007e61016a565b6000546040805160608101909152600180549091908290829060ff1660138111156100a557fe5b60138111156100b057fe5b815260200160018201548152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101575780601f1061012c57610100808354040283529160200191610157565b820191906000526020600020905b81548152906001019060200180831161013a57829003601f168201915b5050505050815250509050915091509091565b6040805160608101909152806000815260200160008152602001606081525090565b600083825260206040818401528351601481106101a557fe5b8060408501525080840151606084015260408401516060608085015280518060a0860152835b818110156101e75782810184015186820160c0015283016101cb565b818111156101f8578460c083880101525b50601f01601f19169390930160c00195945050505050565b602081016006831061021e57fe5b9190529056fea264697066735822122039f39f0b100e1618cfb04ef136841d6cfd7aa3c7c7ffd842bc9a807f34962aa164736f6c63430006080033"

// DeployCostNode deploys a new Ethereum contract, binding an instance of CostNode to it.
func DeployCostNode(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeRef [32]byte, _node CostModel) (common.Address, *types.Transaction, *CostNode, error) {
	parsed, err := abi.JSON(strings.NewReader(CostNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CostNodeBin), backend, _nodeRef, _node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CostNode{CostNodeCaller: CostNodeCaller{contract: contract}, CostNodeTransactor: CostNodeTransactor{contract: contract}, CostNodeFilterer: CostNodeFilterer{contract: contract}}, nil
}

// CostNode is an auto generated Go binding around an Ethereum contract.
type CostNode struct {
	CostNodeCaller     // Read-only binding to the contract
	CostNodeTransactor // Write-only binding to the contract
	CostNodeFilterer   // Log filterer for contract events
}

// CostNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CostNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CostNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CostNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CostNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CostNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CostNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CostNodeSession struct {
	Contract     *CostNode         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CostNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CostNodeCallerSession struct {
	Contract *CostNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CostNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CostNodeTransactorSession struct {
	Contract     *CostNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CostNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CostNodeRaw struct {
	Contract *CostNode // Generic contract binding to access the raw methods on
}

// CostNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CostNodeCallerRaw struct {
	Contract *CostNodeCaller // Generic read-only contract binding to access the raw methods on
}

// CostNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CostNodeTransactorRaw struct {
	Contract *CostNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCostNode creates a new instance of CostNode, bound to a specific deployed contract.
func NewCostNode(address common.Address, backend bind.ContractBackend) (*CostNode, error) {
	contract, err := bindCostNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CostNode{CostNodeCaller: CostNodeCaller{contract: contract}, CostNodeTransactor: CostNodeTransactor{contract: contract}, CostNodeFilterer: CostNodeFilterer{contract: contract}}, nil
}

// NewCostNodeCaller creates a new read-only instance of CostNode, bound to a specific deployed contract.
func NewCostNodeCaller(address common.Address, caller bind.ContractCaller) (*CostNodeCaller, error) {
	contract, err := bindCostNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CostNodeCaller{contract: contract}, nil
}

// NewCostNodeTransactor creates a new write-only instance of CostNode, bound to a specific deployed contract.
func NewCostNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*CostNodeTransactor, error) {
	contract, err := bindCostNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CostNodeTransactor{contract: contract}, nil
}

// NewCostNodeFilterer creates a new log filterer instance of CostNode, bound to a specific deployed contract.
func NewCostNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*CostNodeFilterer, error) {
	contract, err := bindCostNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CostNodeFilterer{contract: contract}, nil
}

// bindCostNode binds a generic wrapper to an already deployed contract.
func bindCostNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CostNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CostNode *CostNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CostNode.Contract.CostNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CostNode *CostNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CostNode.Contract.CostNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CostNode *CostNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CostNode.Contract.CostNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CostNode *CostNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CostNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CostNode *CostNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CostNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CostNode *CostNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CostNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,string))
func (_CostNode *CostNodeCaller) Get(opts *bind.CallOpts) ([32]byte, CostModel, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(CostModel)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _CostNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,string))
func (_CostNode *CostNodeSession) Get() ([32]byte, CostModel, error) {
	return _CostNode.Contract.Get(&_CostNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,string))
func (_CostNode *CostNodeCallerSession) Get() ([32]byte, CostModel, error) {
	return _CostNode.Contract.Get(&_CostNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_CostNode *CostNodeCaller) GetType(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _CostNode.contract.Call(opts, out, "getType")
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_CostNode *CostNodeSession) GetType() (uint8, error) {
	return _CostNode.Contract.GetType(&_CostNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_CostNode *CostNodeCallerSession) GetType() (uint8, error) {
	return _CostNode.Contract.GetType(&_CostNode.CallOpts)
}

// FunctionalUnitNodeABI is the input ABI used to generate the binding from.
const FunctionalUnitNodeABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nodeRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumFunctionalUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerFunctionalUnit\",\"name\":\"_node\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumFunctionalUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerFunctionalUnit\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumPriceModels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FunctionalUnitNodeFuncSigs maps the 4-byte function signature to its string representation.
var FunctionalUnitNodeFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"15dae03e": "getType()",
}

// FunctionalUnitNodeBin is the compiled bytecode used for deploying new contracts.
var FunctionalUnitNodeBin = "0x608060405234801561001057600080fd5b5060405161058138038061058183398101604081905261002f91610227565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615801590610072575060008151600b81111561007057fe5b115b801561008a5750600b8151600b81111561008857fe5b105b80156100995750602081015115155b80156100ad57508060200151816040015110155b6100b657600080fd5b600082905580516001805460ff19168183600b8111156100d257fe5b02179055506020808201516002556040820151600355606082015180516100fd926004920190610105565b5050506102f1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014657805160ff1916838001178555610173565b82800160010185558215610173579182015b82811115610173578251825591602001919060010190610158565b5061017f929150610183565b5090565b61019d91905b8082111561017f5760008155600101610189565b90565b600082601f8301126101b0578081fd5b81516001600160401b038111156101c5578182fd5b60206101d9601f8301601f191682016102cb565b925081835284818386010111156101ef57600080fd5b60005b8281101561020d5784810182015184820183015281016101f2565b8281111561021e5760008284860101525b50505092915050565b60008060408385031215610239578182fd5b825160208401519092506001600160401b0380821115610257578283fd5b81850160808188031215610269578384fd5b61027360806102cb565b92508051600c8110610283578485fd5b8084525060208101516020840152604081015160408401526060810151828111156102ac578485fd5b6102b8888284016101a0565b6060850152505050809150509250929050565b6040518181016001600160401b03811182821017156102e957600080fd5b604052919050565b610281806103006000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315dae03e1461003b5780636d4ce63c14610059575b600080fd5b61004361006f565b604051610050919061022a565b60405180910390f35b610061610074565b60405161005092919061019d565b600290565b600061007e610174565b6000546040805160808101909152600180549091908290829060ff16600b8111156100a557fe5b600b8111156100b057fe5b81526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101615780601f1061013657610100808354040283529160200191610161565b820191906000526020600020905b81548152906001019060200180831161014457829003601f168201915b5050505050815250509050915091509091565b604080516080810190915280600081526020016000815260200160008152602001606081525090565b6000838252602060408184015283516101b58161023e565b80604085015250808401516060840152604084015160808401526060840151608060a085015280518060c0860152835b818110156102015782810184015186820160e0015283016101e5565b81811115610212578460e083880101525b50601f01601f19169390930160e00195945050505050565b602081016006831061023857fe5b91905290565b600c811061024857fe5b5056fea2646970667358221220e60f4aca6429187c9a8a96d9c960238c26149e6cbd81f66344090002e14a3df164736f6c63430006080033"

// DeployFunctionalUnitNode deploys a new Ethereum contract, binding an instance of FunctionalUnitNode to it.
func DeployFunctionalUnitNode(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeRef [32]byte, _node PricePerFunctionalUnit) (common.Address, *types.Transaction, *FunctionalUnitNode, error) {
	parsed, err := abi.JSON(strings.NewReader(FunctionalUnitNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FunctionalUnitNodeBin), backend, _nodeRef, _node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FunctionalUnitNode{FunctionalUnitNodeCaller: FunctionalUnitNodeCaller{contract: contract}, FunctionalUnitNodeTransactor: FunctionalUnitNodeTransactor{contract: contract}, FunctionalUnitNodeFilterer: FunctionalUnitNodeFilterer{contract: contract}}, nil
}

// FunctionalUnitNode is an auto generated Go binding around an Ethereum contract.
type FunctionalUnitNode struct {
	FunctionalUnitNodeCaller     // Read-only binding to the contract
	FunctionalUnitNodeTransactor // Write-only binding to the contract
	FunctionalUnitNodeFilterer   // Log filterer for contract events
}

// FunctionalUnitNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FunctionalUnitNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionalUnitNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FunctionalUnitNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionalUnitNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FunctionalUnitNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionalUnitNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FunctionalUnitNodeSession struct {
	Contract     *FunctionalUnitNode // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FunctionalUnitNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FunctionalUnitNodeCallerSession struct {
	Contract *FunctionalUnitNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// FunctionalUnitNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FunctionalUnitNodeTransactorSession struct {
	Contract     *FunctionalUnitNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FunctionalUnitNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FunctionalUnitNodeRaw struct {
	Contract *FunctionalUnitNode // Generic contract binding to access the raw methods on
}

// FunctionalUnitNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FunctionalUnitNodeCallerRaw struct {
	Contract *FunctionalUnitNodeCaller // Generic read-only contract binding to access the raw methods on
}

// FunctionalUnitNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FunctionalUnitNodeTransactorRaw struct {
	Contract *FunctionalUnitNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFunctionalUnitNode creates a new instance of FunctionalUnitNode, bound to a specific deployed contract.
func NewFunctionalUnitNode(address common.Address, backend bind.ContractBackend) (*FunctionalUnitNode, error) {
	contract, err := bindFunctionalUnitNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FunctionalUnitNode{FunctionalUnitNodeCaller: FunctionalUnitNodeCaller{contract: contract}, FunctionalUnitNodeTransactor: FunctionalUnitNodeTransactor{contract: contract}, FunctionalUnitNodeFilterer: FunctionalUnitNodeFilterer{contract: contract}}, nil
}

// NewFunctionalUnitNodeCaller creates a new read-only instance of FunctionalUnitNode, bound to a specific deployed contract.
func NewFunctionalUnitNodeCaller(address common.Address, caller bind.ContractCaller) (*FunctionalUnitNodeCaller, error) {
	contract, err := bindFunctionalUnitNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FunctionalUnitNodeCaller{contract: contract}, nil
}

// NewFunctionalUnitNodeTransactor creates a new write-only instance of FunctionalUnitNode, bound to a specific deployed contract.
func NewFunctionalUnitNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*FunctionalUnitNodeTransactor, error) {
	contract, err := bindFunctionalUnitNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FunctionalUnitNodeTransactor{contract: contract}, nil
}

// NewFunctionalUnitNodeFilterer creates a new log filterer instance of FunctionalUnitNode, bound to a specific deployed contract.
func NewFunctionalUnitNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*FunctionalUnitNodeFilterer, error) {
	contract, err := bindFunctionalUnitNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FunctionalUnitNodeFilterer{contract: contract}, nil
}

// bindFunctionalUnitNode binds a generic wrapper to an already deployed contract.
func bindFunctionalUnitNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FunctionalUnitNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FunctionalUnitNode *FunctionalUnitNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FunctionalUnitNode.Contract.FunctionalUnitNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FunctionalUnitNode *FunctionalUnitNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionalUnitNode.Contract.FunctionalUnitNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FunctionalUnitNode *FunctionalUnitNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FunctionalUnitNode.Contract.FunctionalUnitNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FunctionalUnitNode *FunctionalUnitNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FunctionalUnitNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FunctionalUnitNode *FunctionalUnitNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionalUnitNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FunctionalUnitNode *FunctionalUnitNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FunctionalUnitNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,uint256,string))
func (_FunctionalUnitNode *FunctionalUnitNodeCaller) Get(opts *bind.CallOpts) ([32]byte, PricePerFunctionalUnit, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PricePerFunctionalUnit)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _FunctionalUnitNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,uint256,string))
func (_FunctionalUnitNode *FunctionalUnitNodeSession) Get() ([32]byte, PricePerFunctionalUnit, error) {
	return _FunctionalUnitNode.Contract.Get(&_FunctionalUnitNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,uint256,string))
func (_FunctionalUnitNode *FunctionalUnitNodeCallerSession) Get() ([32]byte, PricePerFunctionalUnit, error) {
	return _FunctionalUnitNode.Contract.Get(&_FunctionalUnitNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_FunctionalUnitNode *FunctionalUnitNodeCaller) GetType(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _FunctionalUnitNode.contract.Call(opts, out, "getType")
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_FunctionalUnitNode *FunctionalUnitNodeSession) GetType() (uint8, error) {
	return _FunctionalUnitNode.Contract.GetType(&_FunctionalUnitNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_FunctionalUnitNode *FunctionalUnitNodeCallerSession) GetType() (uint8, error) {
	return _FunctionalUnitNode.Contract.GetType(&_FunctionalUnitNode.CallOpts)
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

// ISponsABI is the input ABI used to generate the binding from.
const ISponsABI = "[{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumPriceModels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ISponsFuncSigs maps the 4-byte function signature to its string representation.
var ISponsFuncSigs = map[string]string{
	"15dae03e": "getType()",
}

// ISpons is an auto generated Go binding around an Ethereum contract.
type ISpons struct {
	ISponsCaller     // Read-only binding to the contract
	ISponsTransactor // Write-only binding to the contract
	ISponsFilterer   // Log filterer for contract events
}

// ISponsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISponsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISponsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISponsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISponsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISponsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISponsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISponsSession struct {
	Contract     *ISpons           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISponsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISponsCallerSession struct {
	Contract *ISponsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISponsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISponsTransactorSession struct {
	Contract     *ISponsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISponsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISponsRaw struct {
	Contract *ISpons // Generic contract binding to access the raw methods on
}

// ISponsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISponsCallerRaw struct {
	Contract *ISponsCaller // Generic read-only contract binding to access the raw methods on
}

// ISponsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISponsTransactorRaw struct {
	Contract *ISponsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISpons creates a new instance of ISpons, bound to a specific deployed contract.
func NewISpons(address common.Address, backend bind.ContractBackend) (*ISpons, error) {
	contract, err := bindISpons(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISpons{ISponsCaller: ISponsCaller{contract: contract}, ISponsTransactor: ISponsTransactor{contract: contract}, ISponsFilterer: ISponsFilterer{contract: contract}}, nil
}

// NewISponsCaller creates a new read-only instance of ISpons, bound to a specific deployed contract.
func NewISponsCaller(address common.Address, caller bind.ContractCaller) (*ISponsCaller, error) {
	contract, err := bindISpons(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISponsCaller{contract: contract}, nil
}

// NewISponsTransactor creates a new write-only instance of ISpons, bound to a specific deployed contract.
func NewISponsTransactor(address common.Address, transactor bind.ContractTransactor) (*ISponsTransactor, error) {
	contract, err := bindISpons(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISponsTransactor{contract: contract}, nil
}

// NewISponsFilterer creates a new log filterer instance of ISpons, bound to a specific deployed contract.
func NewISponsFilterer(address common.Address, filterer bind.ContractFilterer) (*ISponsFilterer, error) {
	contract, err := bindISpons(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISponsFilterer{contract: contract}, nil
}

// bindISpons binds a generic wrapper to an already deployed contract.
func bindISpons(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISponsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISpons *ISponsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISpons.Contract.ISponsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISpons *ISponsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISpons.Contract.ISponsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISpons *ISponsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISpons.Contract.ISponsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISpons *ISponsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISpons.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISpons *ISponsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISpons.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISpons *ISponsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISpons.Contract.contract.Transact(opts, method, params...)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_ISpons *ISponsCaller) GetType(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ISpons.contract.Call(opts, out, "getType")
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_ISpons *ISponsSession) GetType() (uint8, error) {
	return _ISpons.Contract.GetType(&_ISpons.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_ISpons *ISponsCallerSession) GetType() (uint8, error) {
	return _ISpons.Contract.GetType(&_ISpons.CallOpts)
}

// ISponsFactoryABI is the input ABI used to generate the binding from.
const ISponsFactoryABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCostModel\",\"name\":\"_node\",\"type\":\"tuple\"}],\"name\":\"addCostModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumFunctionalUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerFunctionalUnit\",\"name\":\"_node\",\"type\":\"tuple\"}],\"name\":\"addFunctionalUnit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"primeCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"labourMinutes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"laboutCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"plant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"material\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPriceForMeasuredWorks\",\"name\":\"_node\",\"type\":\"tuple\"}],\"name\":\"addMeasuredWorks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerSM\",\"name\":\"_node\",\"type\":\"tuple\"}],\"name\":\"addSM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getCostModel\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCostModel\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getFunctionalUnit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumFunctionalUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerFunctionalUnit\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getMeasuredWorks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"primeCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"labourMinutes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"laboutCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"plant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"material\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPriceForMeasuredWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getSM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerSM\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumPriceModels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ISponsFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ISponsFactoryFuncSigs = map[string]string{
	"eafd03bb": "addCostModel(bytes32,(uint8,uint256,string))",
	"71126457": "addFunctionalUnit(bytes32,(uint8,uint256,uint256,string))",
	"56b8ecb1": "addMeasuredWorks(bytes32,(uint8,uint256,uint256,uint256,uint256,uint256,string))",
	"62b119f9": "addSM(bytes32,(uint256,uint256,string))",
	"a42425e6": "getCostModel(bytes32)",
	"e9eaa8bb": "getFunctionalUnit(bytes32)",
	"429ae1ca": "getMeasuredWorks(bytes32)",
	"256ec300": "getSM(bytes32)",
	"7f2aeea4": "getType(bytes32)",
}

// ISponsFactory is an auto generated Go binding around an Ethereum contract.
type ISponsFactory struct {
	ISponsFactoryCaller     // Read-only binding to the contract
	ISponsFactoryTransactor // Write-only binding to the contract
	ISponsFactoryFilterer   // Log filterer for contract events
}

// ISponsFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISponsFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISponsFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISponsFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISponsFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISponsFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISponsFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISponsFactorySession struct {
	Contract     *ISponsFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISponsFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISponsFactoryCallerSession struct {
	Contract *ISponsFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISponsFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISponsFactoryTransactorSession struct {
	Contract     *ISponsFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISponsFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISponsFactoryRaw struct {
	Contract *ISponsFactory // Generic contract binding to access the raw methods on
}

// ISponsFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISponsFactoryCallerRaw struct {
	Contract *ISponsFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ISponsFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISponsFactoryTransactorRaw struct {
	Contract *ISponsFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISponsFactory creates a new instance of ISponsFactory, bound to a specific deployed contract.
func NewISponsFactory(address common.Address, backend bind.ContractBackend) (*ISponsFactory, error) {
	contract, err := bindISponsFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISponsFactory{ISponsFactoryCaller: ISponsFactoryCaller{contract: contract}, ISponsFactoryTransactor: ISponsFactoryTransactor{contract: contract}, ISponsFactoryFilterer: ISponsFactoryFilterer{contract: contract}}, nil
}

// NewISponsFactoryCaller creates a new read-only instance of ISponsFactory, bound to a specific deployed contract.
func NewISponsFactoryCaller(address common.Address, caller bind.ContractCaller) (*ISponsFactoryCaller, error) {
	contract, err := bindISponsFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISponsFactoryCaller{contract: contract}, nil
}

// NewISponsFactoryTransactor creates a new write-only instance of ISponsFactory, bound to a specific deployed contract.
func NewISponsFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ISponsFactoryTransactor, error) {
	contract, err := bindISponsFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISponsFactoryTransactor{contract: contract}, nil
}

// NewISponsFactoryFilterer creates a new log filterer instance of ISponsFactory, bound to a specific deployed contract.
func NewISponsFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ISponsFactoryFilterer, error) {
	contract, err := bindISponsFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISponsFactoryFilterer{contract: contract}, nil
}

// bindISponsFactory binds a generic wrapper to an already deployed contract.
func bindISponsFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISponsFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISponsFactory *ISponsFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISponsFactory.Contract.ISponsFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISponsFactory *ISponsFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISponsFactory.Contract.ISponsFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISponsFactory *ISponsFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISponsFactory.Contract.ISponsFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISponsFactory *ISponsFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ISponsFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISponsFactory *ISponsFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISponsFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISponsFactory *ISponsFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISponsFactory.Contract.contract.Transact(opts, method, params...)
}

// GetCostModel is a free data retrieval call binding the contract method 0xa42425e6.
//
// Solidity: function getCostModel(bytes32 _ref) view returns(bytes32, (uint8,uint256,string))
func (_ISponsFactory *ISponsFactoryCaller) GetCostModel(opts *bind.CallOpts, _ref [32]byte) ([32]byte, CostModel, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(CostModel)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ISponsFactory.contract.Call(opts, out, "getCostModel", _ref)
	return *ret0, *ret1, err
}

// GetCostModel is a free data retrieval call binding the contract method 0xa42425e6.
//
// Solidity: function getCostModel(bytes32 _ref) view returns(bytes32, (uint8,uint256,string))
func (_ISponsFactory *ISponsFactorySession) GetCostModel(_ref [32]byte) ([32]byte, CostModel, error) {
	return _ISponsFactory.Contract.GetCostModel(&_ISponsFactory.CallOpts, _ref)
}

// GetCostModel is a free data retrieval call binding the contract method 0xa42425e6.
//
// Solidity: function getCostModel(bytes32 _ref) view returns(bytes32, (uint8,uint256,string))
func (_ISponsFactory *ISponsFactoryCallerSession) GetCostModel(_ref [32]byte) ([32]byte, CostModel, error) {
	return _ISponsFactory.Contract.GetCostModel(&_ISponsFactory.CallOpts, _ref)
}

// GetFunctionalUnit is a free data retrieval call binding the contract method 0xe9eaa8bb.
//
// Solidity: function getFunctionalUnit(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,string))
func (_ISponsFactory *ISponsFactoryCaller) GetFunctionalUnit(opts *bind.CallOpts, _ref [32]byte) ([32]byte, PricePerFunctionalUnit, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PricePerFunctionalUnit)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ISponsFactory.contract.Call(opts, out, "getFunctionalUnit", _ref)
	return *ret0, *ret1, err
}

// GetFunctionalUnit is a free data retrieval call binding the contract method 0xe9eaa8bb.
//
// Solidity: function getFunctionalUnit(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,string))
func (_ISponsFactory *ISponsFactorySession) GetFunctionalUnit(_ref [32]byte) ([32]byte, PricePerFunctionalUnit, error) {
	return _ISponsFactory.Contract.GetFunctionalUnit(&_ISponsFactory.CallOpts, _ref)
}

// GetFunctionalUnit is a free data retrieval call binding the contract method 0xe9eaa8bb.
//
// Solidity: function getFunctionalUnit(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,string))
func (_ISponsFactory *ISponsFactoryCallerSession) GetFunctionalUnit(_ref [32]byte) ([32]byte, PricePerFunctionalUnit, error) {
	return _ISponsFactory.Contract.GetFunctionalUnit(&_ISponsFactory.CallOpts, _ref)
}

// GetMeasuredWorks is a free data retrieval call binding the contract method 0x429ae1ca.
//
// Solidity: function getMeasuredWorks(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_ISponsFactory *ISponsFactoryCaller) GetMeasuredWorks(opts *bind.CallOpts, _ref [32]byte) ([32]byte, PriceForMeasuredWorks, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PriceForMeasuredWorks)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ISponsFactory.contract.Call(opts, out, "getMeasuredWorks", _ref)
	return *ret0, *ret1, err
}

// GetMeasuredWorks is a free data retrieval call binding the contract method 0x429ae1ca.
//
// Solidity: function getMeasuredWorks(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_ISponsFactory *ISponsFactorySession) GetMeasuredWorks(_ref [32]byte) ([32]byte, PriceForMeasuredWorks, error) {
	return _ISponsFactory.Contract.GetMeasuredWorks(&_ISponsFactory.CallOpts, _ref)
}

// GetMeasuredWorks is a free data retrieval call binding the contract method 0x429ae1ca.
//
// Solidity: function getMeasuredWorks(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_ISponsFactory *ISponsFactoryCallerSession) GetMeasuredWorks(_ref [32]byte) ([32]byte, PriceForMeasuredWorks, error) {
	return _ISponsFactory.Contract.GetMeasuredWorks(&_ISponsFactory.CallOpts, _ref)
}

// GetSM is a free data retrieval call binding the contract method 0x256ec300.
//
// Solidity: function getSM(bytes32 _ref) view returns(bytes32, (uint256,uint256,string))
func (_ISponsFactory *ISponsFactoryCaller) GetSM(opts *bind.CallOpts, _ref [32]byte) ([32]byte, PricePerSM, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PricePerSM)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ISponsFactory.contract.Call(opts, out, "getSM", _ref)
	return *ret0, *ret1, err
}

// GetSM is a free data retrieval call binding the contract method 0x256ec300.
//
// Solidity: function getSM(bytes32 _ref) view returns(bytes32, (uint256,uint256,string))
func (_ISponsFactory *ISponsFactorySession) GetSM(_ref [32]byte) ([32]byte, PricePerSM, error) {
	return _ISponsFactory.Contract.GetSM(&_ISponsFactory.CallOpts, _ref)
}

// GetSM is a free data retrieval call binding the contract method 0x256ec300.
//
// Solidity: function getSM(bytes32 _ref) view returns(bytes32, (uint256,uint256,string))
func (_ISponsFactory *ISponsFactoryCallerSession) GetSM(_ref [32]byte) ([32]byte, PricePerSM, error) {
	return _ISponsFactory.Contract.GetSM(&_ISponsFactory.CallOpts, _ref)
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _ref) view returns(uint8)
func (_ISponsFactory *ISponsFactoryCaller) GetType(opts *bind.CallOpts, _ref [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ISponsFactory.contract.Call(opts, out, "getType", _ref)
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _ref) view returns(uint8)
func (_ISponsFactory *ISponsFactorySession) GetType(_ref [32]byte) (uint8, error) {
	return _ISponsFactory.Contract.GetType(&_ISponsFactory.CallOpts, _ref)
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _ref) view returns(uint8)
func (_ISponsFactory *ISponsFactoryCallerSession) GetType(_ref [32]byte) (uint8, error) {
	return _ISponsFactory.Contract.GetType(&_ISponsFactory.CallOpts, _ref)
}

// AddCostModel is a paid mutator transaction binding the contract method 0xeafd03bb.
//
// Solidity: function addCostModel(bytes32 _ref, (uint8,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactoryTransactor) AddCostModel(opts *bind.TransactOpts, _ref [32]byte, _node CostModel) (*types.Transaction, error) {
	return _ISponsFactory.contract.Transact(opts, "addCostModel", _ref, _node)
}

// AddCostModel is a paid mutator transaction binding the contract method 0xeafd03bb.
//
// Solidity: function addCostModel(bytes32 _ref, (uint8,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactorySession) AddCostModel(_ref [32]byte, _node CostModel) (*types.Transaction, error) {
	return _ISponsFactory.Contract.AddCostModel(&_ISponsFactory.TransactOpts, _ref, _node)
}

// AddCostModel is a paid mutator transaction binding the contract method 0xeafd03bb.
//
// Solidity: function addCostModel(bytes32 _ref, (uint8,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactoryTransactorSession) AddCostModel(_ref [32]byte, _node CostModel) (*types.Transaction, error) {
	return _ISponsFactory.Contract.AddCostModel(&_ISponsFactory.TransactOpts, _ref, _node)
}

// AddFunctionalUnit is a paid mutator transaction binding the contract method 0x71126457.
//
// Solidity: function addFunctionalUnit(bytes32 _ref, (uint8,uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactoryTransactor) AddFunctionalUnit(opts *bind.TransactOpts, _ref [32]byte, _node PricePerFunctionalUnit) (*types.Transaction, error) {
	return _ISponsFactory.contract.Transact(opts, "addFunctionalUnit", _ref, _node)
}

// AddFunctionalUnit is a paid mutator transaction binding the contract method 0x71126457.
//
// Solidity: function addFunctionalUnit(bytes32 _ref, (uint8,uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactorySession) AddFunctionalUnit(_ref [32]byte, _node PricePerFunctionalUnit) (*types.Transaction, error) {
	return _ISponsFactory.Contract.AddFunctionalUnit(&_ISponsFactory.TransactOpts, _ref, _node)
}

// AddFunctionalUnit is a paid mutator transaction binding the contract method 0x71126457.
//
// Solidity: function addFunctionalUnit(bytes32 _ref, (uint8,uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactoryTransactorSession) AddFunctionalUnit(_ref [32]byte, _node PricePerFunctionalUnit) (*types.Transaction, error) {
	return _ISponsFactory.Contract.AddFunctionalUnit(&_ISponsFactory.TransactOpts, _ref, _node)
}

// AddMeasuredWorks is a paid mutator transaction binding the contract method 0x56b8ecb1.
//
// Solidity: function addMeasuredWorks(bytes32 _ref, (uint8,uint256,uint256,uint256,uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactoryTransactor) AddMeasuredWorks(opts *bind.TransactOpts, _ref [32]byte, _node PriceForMeasuredWorks) (*types.Transaction, error) {
	return _ISponsFactory.contract.Transact(opts, "addMeasuredWorks", _ref, _node)
}

// AddMeasuredWorks is a paid mutator transaction binding the contract method 0x56b8ecb1.
//
// Solidity: function addMeasuredWorks(bytes32 _ref, (uint8,uint256,uint256,uint256,uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactorySession) AddMeasuredWorks(_ref [32]byte, _node PriceForMeasuredWorks) (*types.Transaction, error) {
	return _ISponsFactory.Contract.AddMeasuredWorks(&_ISponsFactory.TransactOpts, _ref, _node)
}

// AddMeasuredWorks is a paid mutator transaction binding the contract method 0x56b8ecb1.
//
// Solidity: function addMeasuredWorks(bytes32 _ref, (uint8,uint256,uint256,uint256,uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactoryTransactorSession) AddMeasuredWorks(_ref [32]byte, _node PriceForMeasuredWorks) (*types.Transaction, error) {
	return _ISponsFactory.Contract.AddMeasuredWorks(&_ISponsFactory.TransactOpts, _ref, _node)
}

// AddSM is a paid mutator transaction binding the contract method 0x62b119f9.
//
// Solidity: function addSM(bytes32 _ref, (uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactoryTransactor) AddSM(opts *bind.TransactOpts, _ref [32]byte, _node PricePerSM) (*types.Transaction, error) {
	return _ISponsFactory.contract.Transact(opts, "addSM", _ref, _node)
}

// AddSM is a paid mutator transaction binding the contract method 0x62b119f9.
//
// Solidity: function addSM(bytes32 _ref, (uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactorySession) AddSM(_ref [32]byte, _node PricePerSM) (*types.Transaction, error) {
	return _ISponsFactory.Contract.AddSM(&_ISponsFactory.TransactOpts, _ref, _node)
}

// AddSM is a paid mutator transaction binding the contract method 0x62b119f9.
//
// Solidity: function addSM(bytes32 _ref, (uint256,uint256,string) _node) returns()
func (_ISponsFactory *ISponsFactoryTransactorSession) AddSM(_ref [32]byte, _node PricePerSM) (*types.Transaction, error) {
	return _ISponsFactory.Contract.AddSM(&_ISponsFactory.TransactOpts, _ref, _node)
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

// MeasuredWorksNodeABI is the input ABI used to generate the binding from.
const MeasuredWorksNodeABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nodeRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"primeCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"labourMinutes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"laboutCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"plant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"material\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPriceForMeasuredWorks\",\"name\":\"_node\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"primeCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"labourMinutes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"laboutCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"plant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"material\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPriceForMeasuredWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumPriceModels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MeasuredWorksNodeFuncSigs maps the 4-byte function signature to its string representation.
var MeasuredWorksNodeFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"15dae03e": "getType()",
}

// MeasuredWorksNodeBin is the compiled bytecode used for deploying new contracts.
var MeasuredWorksNodeBin = "0x608060405234801561001057600080fd5b506040516105fc3803806105fc83398101604081905261002f91610231565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615801590610072575060008151601381111561007057fe5b115b801561008a575060138151601381111561008857fe5b105b61009357600080fd5b600082905580516001805460ff1916818360138111156100af57fe5b021790555060208082015160025560408201516003556060820151600455608082015160055560a082015160065560c082015180516100f29260079201906100fa565b505050610313565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013b57805160ff1916838001178555610168565b82800160010185558215610168579182015b8281111561016857825182559160200191906001019061014d565b50610174929150610178565b5090565b61019291905b80821115610174576000815560010161017e565b90565b8051601481106101a457600080fd5b92915050565b600082601f8301126101ba578081fd5b81516001600160401b038111156101cf578182fd5b60206101e3601f8301601f191682016102ed565b925081835284818386010111156101f957600080fd5b60005b828110156102175784810182015184820183015281016101fc565b828111156102285760008284860101525b50505092915050565b60008060408385031215610243578182fd5b825160208401519092506001600160401b0380821115610261578283fd5b81850160e08188031215610273578384fd5b61027d60e06102ed565b92506102898782610195565b83526020810151602084015260408101516040840152606081015160608401526080810151608084015260a081015160a084015260c0810151828111156102ce578485fd5b6102da888284016101aa565b60c0850152505050809150509250929050565b6040518181016001600160401b038111828210171561030b57600080fd5b604052919050565b6102da806103226000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315dae03e1461003b5780636d4ce63c14610059575b600080fd5b61004361006f565b6040516100509190610290565b60405180910390f35b610061610074565b60405161005092919061021b565b600490565b600061007e610192565b6000546040805160e08101909152600180549091908290829060ff1660138111156100a557fe5b60138111156100b057fe5b81526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561017f5780601f106101545761010080835404028352916020019161017f565b820191906000526020600020905b81548152906001019060200180831161016257829003601f168201915b5050505050815250509050915091509091565b6040805160e0810190915280600081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b60008151808452815b818110156101f5576020818501810151868301820152016101d9565b818111156102065782602083870101525b50601f01601f19169290920160200192915050565b60008382526040602083015282516014811061023357fe5b806040840152506020830151606083015260408301516080830152606083015160a0830152608083015160c083015260a083015160e083015260c083015160e06101008401526102876101208401826101d0565b95945050505050565b602081016006831061029e57fe5b9190529056fea2646970667358221220559be30cfcd01feb46c4f249788feb20190c359b0a71f2cca2d8f02d2676903164736f6c63430006080033"

// DeployMeasuredWorksNode deploys a new Ethereum contract, binding an instance of MeasuredWorksNode to it.
func DeployMeasuredWorksNode(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeRef [32]byte, _node PriceForMeasuredWorks) (common.Address, *types.Transaction, *MeasuredWorksNode, error) {
	parsed, err := abi.JSON(strings.NewReader(MeasuredWorksNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MeasuredWorksNodeBin), backend, _nodeRef, _node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MeasuredWorksNode{MeasuredWorksNodeCaller: MeasuredWorksNodeCaller{contract: contract}, MeasuredWorksNodeTransactor: MeasuredWorksNodeTransactor{contract: contract}, MeasuredWorksNodeFilterer: MeasuredWorksNodeFilterer{contract: contract}}, nil
}

// MeasuredWorksNode is an auto generated Go binding around an Ethereum contract.
type MeasuredWorksNode struct {
	MeasuredWorksNodeCaller     // Read-only binding to the contract
	MeasuredWorksNodeTransactor // Write-only binding to the contract
	MeasuredWorksNodeFilterer   // Log filterer for contract events
}

// MeasuredWorksNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MeasuredWorksNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeasuredWorksNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MeasuredWorksNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeasuredWorksNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MeasuredWorksNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeasuredWorksNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MeasuredWorksNodeSession struct {
	Contract     *MeasuredWorksNode // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MeasuredWorksNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MeasuredWorksNodeCallerSession struct {
	Contract *MeasuredWorksNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MeasuredWorksNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MeasuredWorksNodeTransactorSession struct {
	Contract     *MeasuredWorksNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MeasuredWorksNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MeasuredWorksNodeRaw struct {
	Contract *MeasuredWorksNode // Generic contract binding to access the raw methods on
}

// MeasuredWorksNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MeasuredWorksNodeCallerRaw struct {
	Contract *MeasuredWorksNodeCaller // Generic read-only contract binding to access the raw methods on
}

// MeasuredWorksNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MeasuredWorksNodeTransactorRaw struct {
	Contract *MeasuredWorksNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMeasuredWorksNode creates a new instance of MeasuredWorksNode, bound to a specific deployed contract.
func NewMeasuredWorksNode(address common.Address, backend bind.ContractBackend) (*MeasuredWorksNode, error) {
	contract, err := bindMeasuredWorksNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MeasuredWorksNode{MeasuredWorksNodeCaller: MeasuredWorksNodeCaller{contract: contract}, MeasuredWorksNodeTransactor: MeasuredWorksNodeTransactor{contract: contract}, MeasuredWorksNodeFilterer: MeasuredWorksNodeFilterer{contract: contract}}, nil
}

// NewMeasuredWorksNodeCaller creates a new read-only instance of MeasuredWorksNode, bound to a specific deployed contract.
func NewMeasuredWorksNodeCaller(address common.Address, caller bind.ContractCaller) (*MeasuredWorksNodeCaller, error) {
	contract, err := bindMeasuredWorksNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MeasuredWorksNodeCaller{contract: contract}, nil
}

// NewMeasuredWorksNodeTransactor creates a new write-only instance of MeasuredWorksNode, bound to a specific deployed contract.
func NewMeasuredWorksNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*MeasuredWorksNodeTransactor, error) {
	contract, err := bindMeasuredWorksNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MeasuredWorksNodeTransactor{contract: contract}, nil
}

// NewMeasuredWorksNodeFilterer creates a new log filterer instance of MeasuredWorksNode, bound to a specific deployed contract.
func NewMeasuredWorksNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*MeasuredWorksNodeFilterer, error) {
	contract, err := bindMeasuredWorksNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MeasuredWorksNodeFilterer{contract: contract}, nil
}

// bindMeasuredWorksNode binds a generic wrapper to an already deployed contract.
func bindMeasuredWorksNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MeasuredWorksNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MeasuredWorksNode *MeasuredWorksNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MeasuredWorksNode.Contract.MeasuredWorksNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MeasuredWorksNode *MeasuredWorksNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MeasuredWorksNode.Contract.MeasuredWorksNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MeasuredWorksNode *MeasuredWorksNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MeasuredWorksNode.Contract.MeasuredWorksNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MeasuredWorksNode *MeasuredWorksNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MeasuredWorksNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MeasuredWorksNode *MeasuredWorksNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MeasuredWorksNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MeasuredWorksNode *MeasuredWorksNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MeasuredWorksNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_MeasuredWorksNode *MeasuredWorksNodeCaller) Get(opts *bind.CallOpts) ([32]byte, PriceForMeasuredWorks, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PriceForMeasuredWorks)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MeasuredWorksNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_MeasuredWorksNode *MeasuredWorksNodeSession) Get() ([32]byte, PriceForMeasuredWorks, error) {
	return _MeasuredWorksNode.Contract.Get(&_MeasuredWorksNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_MeasuredWorksNode *MeasuredWorksNodeCallerSession) Get() ([32]byte, PriceForMeasuredWorks, error) {
	return _MeasuredWorksNode.Contract.Get(&_MeasuredWorksNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_MeasuredWorksNode *MeasuredWorksNodeCaller) GetType(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MeasuredWorksNode.contract.Call(opts, out, "getType")
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_MeasuredWorksNode *MeasuredWorksNodeSession) GetType() (uint8, error) {
	return _MeasuredWorksNode.Contract.GetType(&_MeasuredWorksNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_MeasuredWorksNode *MeasuredWorksNodeCallerSession) GetType() (uint8, error) {
	return _MeasuredWorksNode.Contract.GetType(&_MeasuredWorksNode.CallOpts)
}

// SMNodeABI is the input ABI used to generate the binding from.
const SMNodeABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nodeRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerSM\",\"name\":\"_node\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerSM\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumPriceModels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SMNodeFuncSigs maps the 4-byte function signature to its string representation.
var SMNodeFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"15dae03e": "getType()",
}

// SMNodeBin is the compiled bytecode used for deploying new contracts.
var SMNodeBin = "0x608060405234801561001057600080fd5b5060405161049338038061049383398101604081905261002f9161014b565b8160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906100665750805115155b801561007757508051602082015110155b61008057600080fd5b60008290558051600155602080820151600255604082015180516100a89260039201906100b0565b50505061026b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f157805160ff191683800117855561011e565b8280016001018555821561011e579182015b8281111561011e578251825591602001919060010190610103565b5061012a92915061012e565b5090565b61014891905b8082111561012a5760008155600101610134565b90565b6000806040838503121561015d578182fd5b8251602080850151919350906001600160401b038082111561017d578384fd5b8186016060818903121561018f578485fd5b6101996060610245565b92508051835283810151848401526040810151828111156101b8578586fd5b80820189601f8201126101c9578687fd5b80519250838311156101d9578687fd5b6101eb601f8401601f19168701610245565b93508284528986848301011115610200578687fd5b8691505b828210156102215780820186015184830187015290850190610204565b508181111561023257858583850101525b5050604082015292959294509192505050565b6040518181016001600160401b038111828210171561026357600080fd5b604052919050565b6102198061027a6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315dae03e1461003b5780636d4ce63c14610059575b600080fd5b61004361006f565b60405161005091906101cf565b60405180910390f35b610061610074565b604051610050929190610157565b600390565b600061007e610136565b600054604080516060810182526001805482526002805460208085019190915260038054865181861615610100026000190190911693909304601f810183900483028401830187528084529395869490860193928301828280156101235780601f106100f857610100808354040283529160200191610123565b820191906000526020600020905b81548152906001019060200180831161010657829003601f168201915b5050505050815250509050915091509091565b60405180606001604052806000815260200160008152602001606081525090565b600083825260206040818401528351604084015280840151606084015260408401516060608085015280518060a0860152835b818110156101a65782810184015186820160c00152830161018a565b818111156101b7578460c083880101525b50601f01601f19169390930160c00195945050505050565b60208101600683106101dd57fe5b9190529056fea264697066735822122083ab392ed34f72458c61f6ad48d7c7e7d98d69f6116a4ea9044c4909694486ca64736f6c63430006080033"

// DeploySMNode deploys a new Ethereum contract, binding an instance of SMNode to it.
func DeploySMNode(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeRef [32]byte, _node PricePerSM) (common.Address, *types.Transaction, *SMNode, error) {
	parsed, err := abi.JSON(strings.NewReader(SMNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SMNodeBin), backend, _nodeRef, _node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SMNode{SMNodeCaller: SMNodeCaller{contract: contract}, SMNodeTransactor: SMNodeTransactor{contract: contract}, SMNodeFilterer: SMNodeFilterer{contract: contract}}, nil
}

// SMNode is an auto generated Go binding around an Ethereum contract.
type SMNode struct {
	SMNodeCaller     // Read-only binding to the contract
	SMNodeTransactor // Write-only binding to the contract
	SMNodeFilterer   // Log filterer for contract events
}

// SMNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SMNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SMNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SMNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SMNodeSession struct {
	Contract     *SMNode           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SMNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SMNodeCallerSession struct {
	Contract *SMNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SMNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SMNodeTransactorSession struct {
	Contract     *SMNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SMNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SMNodeRaw struct {
	Contract *SMNode // Generic contract binding to access the raw methods on
}

// SMNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SMNodeCallerRaw struct {
	Contract *SMNodeCaller // Generic read-only contract binding to access the raw methods on
}

// SMNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SMNodeTransactorRaw struct {
	Contract *SMNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSMNode creates a new instance of SMNode, bound to a specific deployed contract.
func NewSMNode(address common.Address, backend bind.ContractBackend) (*SMNode, error) {
	contract, err := bindSMNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SMNode{SMNodeCaller: SMNodeCaller{contract: contract}, SMNodeTransactor: SMNodeTransactor{contract: contract}, SMNodeFilterer: SMNodeFilterer{contract: contract}}, nil
}

// NewSMNodeCaller creates a new read-only instance of SMNode, bound to a specific deployed contract.
func NewSMNodeCaller(address common.Address, caller bind.ContractCaller) (*SMNodeCaller, error) {
	contract, err := bindSMNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SMNodeCaller{contract: contract}, nil
}

// NewSMNodeTransactor creates a new write-only instance of SMNode, bound to a specific deployed contract.
func NewSMNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*SMNodeTransactor, error) {
	contract, err := bindSMNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SMNodeTransactor{contract: contract}, nil
}

// NewSMNodeFilterer creates a new log filterer instance of SMNode, bound to a specific deployed contract.
func NewSMNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*SMNodeFilterer, error) {
	contract, err := bindSMNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SMNodeFilterer{contract: contract}, nil
}

// bindSMNode binds a generic wrapper to an already deployed contract.
func bindSMNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SMNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SMNode *SMNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SMNode.Contract.SMNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SMNode *SMNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMNode.Contract.SMNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SMNode *SMNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SMNode.Contract.SMNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SMNode *SMNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SMNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SMNode *SMNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SMNode *SMNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SMNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint256,uint256,string))
func (_SMNode *SMNodeCaller) Get(opts *bind.CallOpts) ([32]byte, PricePerSM, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PricePerSM)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SMNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint256,uint256,string))
func (_SMNode *SMNodeSession) Get() ([32]byte, PricePerSM, error) {
	return _SMNode.Contract.Get(&_SMNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(bytes32, (uint256,uint256,string))
func (_SMNode *SMNodeCallerSession) Get() ([32]byte, PricePerSM, error) {
	return _SMNode.Contract.Get(&_SMNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_SMNode *SMNodeCaller) GetType(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SMNode.contract.Call(opts, out, "getType")
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_SMNode *SMNodeSession) GetType() (uint8, error) {
	return _SMNode.Contract.GetType(&_SMNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_SMNode *SMNodeCallerSession) GetType() (uint8, error) {
	return _SMNode.Contract.GetType(&_SMNode.CallOpts)
}

// SponsABI is the input ABI used to generate the binding from.
const SponsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCostModel\",\"name\":\"_node\",\"type\":\"tuple\"}],\"name\":\"addCostModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumFunctionalUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerFunctionalUnit\",\"name\":\"_node\",\"type\":\"tuple\"}],\"name\":\"addFunctionalUnit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"primeCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"labourMinutes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"laboutCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"plant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"material\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPriceForMeasuredWorks\",\"name\":\"_node\",\"type\":\"tuple\"}],\"name\":\"addMeasuredWorks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerSM\",\"name\":\"_node\",\"type\":\"tuple\"}],\"name\":\"addSM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getCostModel\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCostModel\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getFunctionalUnit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumFunctionalUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerFunctionalUnit\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getMeasuredWorks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumUnits\",\"name\":\"unit\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"primeCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"labourMinutes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"laboutCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"plant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"material\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPriceForMeasuredWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getSM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structPricePerSM\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumPriceModels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SponsFuncSigs maps the 4-byte function signature to its string representation.
var SponsFuncSigs = map[string]string{
	"eafd03bb": "addCostModel(bytes32,(uint8,uint256,string))",
	"71126457": "addFunctionalUnit(bytes32,(uint8,uint256,uint256,string))",
	"56b8ecb1": "addMeasuredWorks(bytes32,(uint8,uint256,uint256,uint256,uint256,uint256,string))",
	"62b119f9": "addSM(bytes32,(uint256,uint256,string))",
	"a42425e6": "getCostModel(bytes32)",
	"e9eaa8bb": "getFunctionalUnit(bytes32)",
	"429ae1ca": "getMeasuredWorks(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"256ec300": "getSM(bytes32)",
	"7f2aeea4": "getType(bytes32)",
}

// SponsBin is the compiled bytecode used for deploying new contracts.
var SponsBin = "0x608060405234801561001057600080fd5b50600060025561291d806100256000396000f3fe60806040523480156200001157600080fd5b5060043610620000b85760003560e01c806371126457116200007b57806371126457146200015d5780637f2aeea41462000174578063a42425e6146200019a578063bc59934114620001c1578063e9eaa8bb14620001d8578063eafd03bb14620001ff57620000b8565b8063256ec30014620000bd578063429ae1ca14620000ed57806356b8ecb1146200011457806362b119f9146200012d57806367e0badb1462000144575b600080fd5b620000d4620000ce36600462000c61565b62000216565b604051620000e4929190620012ab565b60405180910390f35b62000104620000fe36600462000c61565b62000391565b604051620000e4929190620011ed565b6200012b6200012536600462000da5565b62000502565b005b6200012b6200013e36600462001061565b62000576565b6200014e620005b0565b604051620000e4919062001196565b6200012b6200016e36600462000f1e565b620005b6565b6200018b6200018536600462000c61565b620005f0565b604051620000e49190620012e3565b620001b1620001ab36600462000c61565b620006c7565b604051620000e49291906200119f565b6200014e620001d236600462000c61565b62000839565b620001ef620001e936600462000c61565b62000873565b604051620000e49291906200125c565b6200012b6200021036600462000c7a565b620009e4565b60006200022262000ab2565b8260001a60f81b6001600160f81b031916158015906200025b57506000838152602081905260409020600101546001600160a01b031615155b6200026557600080fd5b6000838152602081905260409020600101546001600160a01b03166003816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b158015620002bc57600080fd5b505afa158015620002d1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002f7919062001147565b60058111156200030357fe5b146200030e57600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b1580156200034857600080fd5b505afa1580156200035d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620003879190810190620010d4565b9250925050915091565b60006200039d62000ad3565b8260001a60f81b6001600160f81b03191615801590620003d657506000838152602081905260409020600101546001600160a01b031615155b620003e057600080fd5b6000838152602081905260409020600101546001600160a01b03166004816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200043757600080fd5b505afa1580156200044c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000472919062001147565b60058111156200047e57fe5b146200048957600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b158015620004c357600080fd5b505afa158015620004d8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000387919081019062000e6b565b8160001a60f81b6001600160f81b0319166200051d57600080fd5b600082826040516200052f9062000b11565b6200053c929190620011ed565b604051809103906000f08015801562000559573d6000803e3d6000fd5b509050620005706000848363ffffffff62000a1e16565b50505050565b8160001a60f81b6001600160f81b0319166200059157600080fd5b60008282604051620005a39062000b1f565b6200053c929190620012ab565b60025490565b8160001a60f81b6001600160f81b031916620005d157600080fd5b60008282604051620005e39062000b2d565b6200053c9291906200125c565b600081811a60f81b6001600160f81b031916158015906200062a57506000828152602081905260409020600101546001600160a01b031615155b6200063457600080fd5b60008281526020818152604091829020600101548251630aed701f60e11b815292516001600160a01b039091169283926315dae03e92600480840193829003018186803b1580156200068557600080fd5b505afa1580156200069a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620006c0919062001147565b9392505050565b6000620006d362000b3b565b8260001a60f81b6001600160f81b031916158015906200070c57506000838152602081905260409020600101546001600160a01b031615155b6200071657600080fd5b60008381526020819052604090206001908101546001600160a01b031690816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200076e57600080fd5b505afa15801562000783573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620007a9919062001147565b6005811115620007b557fe5b14620007c057600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b158015620007fa57600080fd5b505afa1580156200080f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000387919081019062000d19565b60025460009082106200084b57600080fd5b60018054839081106200085a57fe5b9060005260206000209060020201600001549050919050565b60006200087f62000b5d565b8260001a60f81b6001600160f81b03191615801590620008b857506000838152602081905260409020600101546001600160a01b031615155b620008c257600080fd5b6000838152602081905260409020600101546001600160a01b03166002816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156200091957600080fd5b505afa1580156200092e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000954919062001147565b60058111156200096057fe5b146200096b57600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b158015620009a557600080fd5b505afa158015620009ba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000387919081019062000fc9565b8160001a60f81b6001600160f81b031916620009ff57600080fd5b6000828260405162000a119062000b86565b6200053c9291906200119f565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000a5e576001915050620006c0565b506001808501805491820180825560008681526020889052604090208190558591908390811062000a8b57fe5b60009182526020822060029182020192909255908601805460010190559150620006c09050565b60405180606001604052806000815260200160008152602001606081525090565b6040805160e0810190915280600081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b6105fc806200139483390190565b610493806200199083390190565b6105818062001e2383390190565b6040805160608101909152806000815260200160008152602001606081525090565b604080516080810190915280600081526020016000815260200160008152602001606081525090565b61054480620023a483390190565b803562000ba18162001385565b92915050565b805162000ba18162001385565b600082601f83011262000bc5578081fd5b813562000bdc62000bd68262001320565b620012f8565b915080825283602082850101111562000bf457600080fd5b8060208401602084013760009082016020015292915050565b600082601f83011262000c1e578081fd5b815162000c2f62000bd68262001320565b915080825283602082850101111562000c4757600080fd5b62000c5a81602084016020860162001345565b5092915050565b60006020828403121562000c73578081fd5b5035919050565b6000806040838503121562000c8d578081fd5b82359150602083013567ffffffffffffffff8082111562000cac578283fd5b8185016060818803121562000cbf578384fd5b62000ccb6060620012f8565b9250803562000cda8162001385565b83526020818101359084015260408101358281111562000cf8578485fd5b62000d068882840162000bb4565b6040850152505050809150509250929050565b6000806040838503121562000d2c578081fd5b82519150602083015167ffffffffffffffff8082111562000d4b578283fd5b8185016060818803121562000d5e578384fd5b62000d6a6060620012f8565b9250805162000d798162001385565b83526020818101519084015260408101518281111562000d97578485fd5b62000d068882840162000c0d565b6000806040838503121562000db8578182fd5b82359150602083013567ffffffffffffffff8082111562000dd7578283fd5b81850160e0818803121562000dea578384fd5b62000df660e0620012f8565b925062000e04878262000b94565b83526020810135602084015260408101356040840152606081013560608401526080810135608084015260a081013560a084015260c08101358281111562000e4a578485fd5b62000e588882840162000bb4565b60c0850152505050809150509250929050565b6000806040838503121562000e7e578182fd5b82519150602083015167ffffffffffffffff8082111562000e9d578283fd5b81850160e0818803121562000eb0578384fd5b62000ebc60e0620012f8565b925062000eca878262000ba7565b83526020810151602084015260408101516040840152606081015160608401526080810151608084015260a081015160a084015260c08101518281111562000f10578485fd5b62000e588882840162000c0d565b6000806040838503121562000f31578182fd5b82359150602083013567ffffffffffffffff8082111562000f50578283fd5b8185016080818803121562000f63578384fd5b62000f6f6080620012f8565b9250803562000f7e8162001374565b80845250602081013560208401526040810135604084015260608101358281111562000fa8578485fd5b62000fb68882840162000bb4565b6060850152505050809150509250929050565b6000806040838503121562000fdc578182fd5b82519150602083015167ffffffffffffffff8082111562000ffb578283fd5b818501608081880312156200100e578384fd5b6200101a6080620012f8565b92508051620010298162001374565b80845250602081015160208401526040810151604084015260608101518281111562001053578485fd5b62000fb68882840162000c0d565b6000806040838503121562001074578182fd5b82359150602083013567ffffffffffffffff8082111562001093578283fd5b81850160608188031215620010a6578384fd5b620010b26060620012f8565b9250803583526020810135602084015260408101358281111562000cf8578485fd5b60008060408385031215620010e7578182fd5b82519150602083015167ffffffffffffffff8082111562001106578283fd5b8185016060818803121562001119578384fd5b620011256060620012f8565b9250805183526020810151602084015260408101518281111562000d97578485fd5b60006020828403121562001159578081fd5b815160068110620006c0578182fd5b600081518084526200118281602086016020860162001345565b601f01601f19169290920160200192915050565b90815260200190565b600083825260406020830152825160148110620011b857fe5b8060408401525060208301516060830152604083015160606080840152620011e460a084018262001168565b95945050505050565b6000838252604060208301528251601481106200120657fe5b806040840152506020830151606083015260408301516080830152606083015160a0830152608083015160c083015260a083015160e083015260c083015160e0610100840152620011e461012084018262001168565b6000838252604060208301528251600c81106200127557fe5b8060408401525060208301516060830152604083015160808301526060830151608060a0840152620011e460c084018262001168565b6000838252604060208301528251604083015260208301516060830152604083015160606080840152620011e460a084018262001168565b6020810160068310620012f257fe5b91905290565b60405181810167ffffffffffffffff811182821017156200131857600080fd5b604052919050565b600067ffffffffffffffff82111562001337578081fd5b50601f01601f191660200190565b60005b838110156200136257818101518382015260200162001348565b83811115620005705750506000910152565b600c81106200138257600080fd5b50565b601481106200138257600080fdfe608060405234801561001057600080fd5b506040516105fc3803806105fc83398101604081905261002f91610231565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615801590610072575060008151601381111561007057fe5b115b801561008a575060138151601381111561008857fe5b105b61009357600080fd5b600082905580516001805460ff1916818360138111156100af57fe5b021790555060208082015160025560408201516003556060820151600455608082015160055560a082015160065560c082015180516100f29260079201906100fa565b505050610313565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013b57805160ff1916838001178555610168565b82800160010185558215610168579182015b8281111561016857825182559160200191906001019061014d565b50610174929150610178565b5090565b61019291905b80821115610174576000815560010161017e565b90565b8051601481106101a457600080fd5b92915050565b600082601f8301126101ba578081fd5b81516001600160401b038111156101cf578182fd5b60206101e3601f8301601f191682016102ed565b925081835284818386010111156101f957600080fd5b60005b828110156102175784810182015184820183015281016101fc565b828111156102285760008284860101525b50505092915050565b60008060408385031215610243578182fd5b825160208401519092506001600160401b0380821115610261578283fd5b81850160e08188031215610273578384fd5b61027d60e06102ed565b92506102898782610195565b83526020810151602084015260408101516040840152606081015160608401526080810151608084015260a081015160a084015260c0810151828111156102ce578485fd5b6102da888284016101aa565b60c0850152505050809150509250929050565b6040518181016001600160401b038111828210171561030b57600080fd5b604052919050565b6102da806103226000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315dae03e1461003b5780636d4ce63c14610059575b600080fd5b61004361006f565b6040516100509190610290565b60405180910390f35b610061610074565b60405161005092919061021b565b600490565b600061007e610192565b6000546040805160e08101909152600180549091908290829060ff1660138111156100a557fe5b60138111156100b057fe5b81526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561017f5780601f106101545761010080835404028352916020019161017f565b820191906000526020600020905b81548152906001019060200180831161016257829003601f168201915b5050505050815250509050915091509091565b6040805160e0810190915280600081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b60008151808452815b818110156101f5576020818501810151868301820152016101d9565b818111156102065782602083870101525b50601f01601f19169290920160200192915050565b60008382526040602083015282516014811061023357fe5b806040840152506020830151606083015260408301516080830152606083015160a0830152608083015160c083015260a083015160e083015260c083015160e06101008401526102876101208401826101d0565b95945050505050565b602081016006831061029e57fe5b9190529056fea2646970667358221220559be30cfcd01feb46c4f249788feb20190c359b0a71f2cca2d8f02d2676903164736f6c63430006080033608060405234801561001057600080fd5b5060405161049338038061049383398101604081905261002f9161014b565b8160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906100665750805115155b801561007757508051602082015110155b61008057600080fd5b60008290558051600155602080820151600255604082015180516100a89260039201906100b0565b50505061026b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f157805160ff191683800117855561011e565b8280016001018555821561011e579182015b8281111561011e578251825591602001919060010190610103565b5061012a92915061012e565b5090565b61014891905b8082111561012a5760008155600101610134565b90565b6000806040838503121561015d578182fd5b8251602080850151919350906001600160401b038082111561017d578384fd5b8186016060818903121561018f578485fd5b6101996060610245565b92508051835283810151848401526040810151828111156101b8578586fd5b80820189601f8201126101c9578687fd5b80519250838311156101d9578687fd5b6101eb601f8401601f19168701610245565b93508284528986848301011115610200578687fd5b8691505b828210156102215780820186015184830187015290850190610204565b508181111561023257858583850101525b5050604082015292959294509192505050565b6040518181016001600160401b038111828210171561026357600080fd5b604052919050565b6102198061027a6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315dae03e1461003b5780636d4ce63c14610059575b600080fd5b61004361006f565b60405161005091906101cf565b60405180910390f35b610061610074565b604051610050929190610157565b600390565b600061007e610136565b600054604080516060810182526001805482526002805460208085019190915260038054865181861615610100026000190190911693909304601f810183900483028401830187528084529395869490860193928301828280156101235780601f106100f857610100808354040283529160200191610123565b820191906000526020600020905b81548152906001019060200180831161010657829003601f168201915b5050505050815250509050915091509091565b60405180606001604052806000815260200160008152602001606081525090565b600083825260206040818401528351604084015280840151606084015260408401516060608085015280518060a0860152835b818110156101a65782810184015186820160c00152830161018a565b818111156101b7578460c083880101525b50601f01601f19169390930160c00195945050505050565b60208101600683106101dd57fe5b9190529056fea264697066735822122083ab392ed34f72458c61f6ad48d7c7e7d98d69f6116a4ea9044c4909694486ca64736f6c63430006080033608060405234801561001057600080fd5b5060405161058138038061058183398101604081905261002f91610227565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615801590610072575060008151600b81111561007057fe5b115b801561008a5750600b8151600b81111561008857fe5b105b80156100995750602081015115155b80156100ad57508060200151816040015110155b6100b657600080fd5b600082905580516001805460ff19168183600b8111156100d257fe5b02179055506020808201516002556040820151600355606082015180516100fd926004920190610105565b5050506102f1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014657805160ff1916838001178555610173565b82800160010185558215610173579182015b82811115610173578251825591602001919060010190610158565b5061017f929150610183565b5090565b61019d91905b8082111561017f5760008155600101610189565b90565b600082601f8301126101b0578081fd5b81516001600160401b038111156101c5578182fd5b60206101d9601f8301601f191682016102cb565b925081835284818386010111156101ef57600080fd5b60005b8281101561020d5784810182015184820183015281016101f2565b8281111561021e5760008284860101525b50505092915050565b60008060408385031215610239578182fd5b825160208401519092506001600160401b0380821115610257578283fd5b81850160808188031215610269578384fd5b61027360806102cb565b92508051600c8110610283578485fd5b8084525060208101516020840152604081015160408401526060810151828111156102ac578485fd5b6102b8888284016101a0565b6060850152505050809150509250929050565b6040518181016001600160401b03811182821017156102e957600080fd5b604052919050565b610281806103006000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315dae03e1461003b5780636d4ce63c14610059575b600080fd5b61004361006f565b604051610050919061022a565b60405180910390f35b610061610074565b60405161005092919061019d565b600290565b600061007e610174565b6000546040805160808101909152600180549091908290829060ff16600b8111156100a557fe5b600b8111156100b057fe5b81526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101615780601f1061013657610100808354040283529160200191610161565b820191906000526020600020905b81548152906001019060200180831161014457829003601f168201915b5050505050815250509050915091509091565b604080516080810190915280600081526020016000815260200160008152602001606081525090565b6000838252602060408184015283516101b58161023e565b80604085015250808401516060840152604084015160808401526060840151608060a085015280518060c0860152835b818110156102015782810184015186820160e0015283016101e5565b81811115610212578460e083880101525b50601f01601f19169390930160e00195945050505050565b602081016006831061023857fe5b91905290565b600c811061024857fe5b5056fea2646970667358221220e60f4aca6429187c9a8a96d9c960238c26149e6cbd81f66344090002e14a3df164736f6c63430006080033608060405234801561001057600080fd5b5060405161054438038061054483398101604081905261002f91610184565b8160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615801590610072575060008151601381111561007057fe5b115b801561008a575060138151601381111561008857fe5b105b80156100995750602081015115155b6100a257600080fd5b600082905580516001805460ff1916818360138111156100be57fe5b0217905550602080820151600255604082015180516100e19260039201906100e9565b5050506102db565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061012a57805160ff1916838001178555610157565b82800160010185558215610157579182015b8281111561015757825182559160200191906001019061013c565b50610163929150610167565b5090565b61018191905b80821115610163576000815560010161016d565b90565b60008060408385031215610196578182fd5b8251602080850151919350906001600160401b03808211156101b6578384fd5b818601606081890312156101c8578485fd5b6101d26060610252565b925080516101df816102cb565b835280840151848401526040810151828111156101fa578586fd5b019050601f8101871361020b578384fd5b805161021e61021982610278565b610252565b8181528885838501011115610231578586fd5b6102408286830187860161029b565b60408401525093969095509350505050565b6040518181016001600160401b038111828210171561027057600080fd5b604052919050565b60006001600160401b0382111561028d578081fd5b50601f01601f191660200190565b60005b838110156102b657818101518382015260200161029e565b838111156102c5576000848401525b50505050565b601481106102d857600080fd5b50565b61025a806102ea6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806315dae03e1461003b5780636d4ce63c14610059575b600080fd5b61004361006f565b6040516100509190610210565b60405180910390f35b610061610074565b60405161005092919061018c565b600190565b600061007e61016a565b6000546040805160608101909152600180549091908290829060ff1660138111156100a557fe5b60138111156100b057fe5b815260200160018201548152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101575780601f1061012c57610100808354040283529160200191610157565b820191906000526020600020905b81548152906001019060200180831161013a57829003601f168201915b5050505050815250509050915091509091565b6040805160608101909152806000815260200160008152602001606081525090565b600083825260206040818401528351601481106101a557fe5b8060408501525080840151606084015260408401516060608085015280518060a0860152835b818110156101e75782810184015186820160c0015283016101cb565b818111156101f8578460c083880101525b50601f01601f19169390930160c00195945050505050565b602081016006831061021e57fe5b9190529056fea264697066735822122039f39f0b100e1618cfb04ef136841d6cfd7aa3c7c7ffd842bc9a807f34962aa164736f6c63430006080033a2646970667358221220c201219b66f8ecc3a56c7937ea6910df0345480e6ce2cc7f96a9c01485f43e4764736f6c63430006080033"

// DeploySpons deploys a new Ethereum contract, binding an instance of Spons to it.
func DeploySpons(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Spons, error) {
	parsed, err := abi.JSON(strings.NewReader(SponsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SponsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Spons{SponsCaller: SponsCaller{contract: contract}, SponsTransactor: SponsTransactor{contract: contract}, SponsFilterer: SponsFilterer{contract: contract}}, nil
}

// Spons is an auto generated Go binding around an Ethereum contract.
type Spons struct {
	SponsCaller     // Read-only binding to the contract
	SponsTransactor // Write-only binding to the contract
	SponsFilterer   // Log filterer for contract events
}

// SponsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SponsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SponsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SponsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SponsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SponsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SponsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SponsSession struct {
	Contract     *Spons            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SponsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SponsCallerSession struct {
	Contract *SponsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SponsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SponsTransactorSession struct {
	Contract     *SponsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SponsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SponsRaw struct {
	Contract *Spons // Generic contract binding to access the raw methods on
}

// SponsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SponsCallerRaw struct {
	Contract *SponsCaller // Generic read-only contract binding to access the raw methods on
}

// SponsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SponsTransactorRaw struct {
	Contract *SponsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpons creates a new instance of Spons, bound to a specific deployed contract.
func NewSpons(address common.Address, backend bind.ContractBackend) (*Spons, error) {
	contract, err := bindSpons(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spons{SponsCaller: SponsCaller{contract: contract}, SponsTransactor: SponsTransactor{contract: contract}, SponsFilterer: SponsFilterer{contract: contract}}, nil
}

// NewSponsCaller creates a new read-only instance of Spons, bound to a specific deployed contract.
func NewSponsCaller(address common.Address, caller bind.ContractCaller) (*SponsCaller, error) {
	contract, err := bindSpons(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SponsCaller{contract: contract}, nil
}

// NewSponsTransactor creates a new write-only instance of Spons, bound to a specific deployed contract.
func NewSponsTransactor(address common.Address, transactor bind.ContractTransactor) (*SponsTransactor, error) {
	contract, err := bindSpons(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SponsTransactor{contract: contract}, nil
}

// NewSponsFilterer creates a new log filterer instance of Spons, bound to a specific deployed contract.
func NewSponsFilterer(address common.Address, filterer bind.ContractFilterer) (*SponsFilterer, error) {
	contract, err := bindSpons(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SponsFilterer{contract: contract}, nil
}

// bindSpons binds a generic wrapper to an already deployed contract.
func bindSpons(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SponsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spons *SponsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spons.Contract.SponsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spons *SponsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spons.Contract.SponsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spons *SponsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spons.Contract.SponsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spons *SponsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Spons.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spons *SponsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spons.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spons *SponsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spons.Contract.contract.Transact(opts, method, params...)
}

// GetCostModel is a free data retrieval call binding the contract method 0xa42425e6.
//
// Solidity: function getCostModel(bytes32 _ref) view returns(bytes32, (uint8,uint256,string))
func (_Spons *SponsCaller) GetCostModel(opts *bind.CallOpts, _ref [32]byte) ([32]byte, CostModel, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(CostModel)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Spons.contract.Call(opts, out, "getCostModel", _ref)
	return *ret0, *ret1, err
}

// GetCostModel is a free data retrieval call binding the contract method 0xa42425e6.
//
// Solidity: function getCostModel(bytes32 _ref) view returns(bytes32, (uint8,uint256,string))
func (_Spons *SponsSession) GetCostModel(_ref [32]byte) ([32]byte, CostModel, error) {
	return _Spons.Contract.GetCostModel(&_Spons.CallOpts, _ref)
}

// GetCostModel is a free data retrieval call binding the contract method 0xa42425e6.
//
// Solidity: function getCostModel(bytes32 _ref) view returns(bytes32, (uint8,uint256,string))
func (_Spons *SponsCallerSession) GetCostModel(_ref [32]byte) ([32]byte, CostModel, error) {
	return _Spons.Contract.GetCostModel(&_Spons.CallOpts, _ref)
}

// GetFunctionalUnit is a free data retrieval call binding the contract method 0xe9eaa8bb.
//
// Solidity: function getFunctionalUnit(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,string))
func (_Spons *SponsCaller) GetFunctionalUnit(opts *bind.CallOpts, _ref [32]byte) ([32]byte, PricePerFunctionalUnit, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PricePerFunctionalUnit)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Spons.contract.Call(opts, out, "getFunctionalUnit", _ref)
	return *ret0, *ret1, err
}

// GetFunctionalUnit is a free data retrieval call binding the contract method 0xe9eaa8bb.
//
// Solidity: function getFunctionalUnit(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,string))
func (_Spons *SponsSession) GetFunctionalUnit(_ref [32]byte) ([32]byte, PricePerFunctionalUnit, error) {
	return _Spons.Contract.GetFunctionalUnit(&_Spons.CallOpts, _ref)
}

// GetFunctionalUnit is a free data retrieval call binding the contract method 0xe9eaa8bb.
//
// Solidity: function getFunctionalUnit(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,string))
func (_Spons *SponsCallerSession) GetFunctionalUnit(_ref [32]byte) ([32]byte, PricePerFunctionalUnit, error) {
	return _Spons.Contract.GetFunctionalUnit(&_Spons.CallOpts, _ref)
}

// GetMeasuredWorks is a free data retrieval call binding the contract method 0x429ae1ca.
//
// Solidity: function getMeasuredWorks(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_Spons *SponsCaller) GetMeasuredWorks(opts *bind.CallOpts, _ref [32]byte) ([32]byte, PriceForMeasuredWorks, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PriceForMeasuredWorks)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Spons.contract.Call(opts, out, "getMeasuredWorks", _ref)
	return *ret0, *ret1, err
}

// GetMeasuredWorks is a free data retrieval call binding the contract method 0x429ae1ca.
//
// Solidity: function getMeasuredWorks(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_Spons *SponsSession) GetMeasuredWorks(_ref [32]byte) ([32]byte, PriceForMeasuredWorks, error) {
	return _Spons.Contract.GetMeasuredWorks(&_Spons.CallOpts, _ref)
}

// GetMeasuredWorks is a free data retrieval call binding the contract method 0x429ae1ca.
//
// Solidity: function getMeasuredWorks(bytes32 _ref) view returns(bytes32, (uint8,uint256,uint256,uint256,uint256,uint256,string))
func (_Spons *SponsCallerSession) GetMeasuredWorks(_ref [32]byte) ([32]byte, PriceForMeasuredWorks, error) {
	return _Spons.Contract.GetMeasuredWorks(&_Spons.CallOpts, _ref)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Spons *SponsCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Spons.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Spons *SponsSession) GetNum() (*big.Int, error) {
	return _Spons.Contract.GetNum(&_Spons.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Spons *SponsCallerSession) GetNum() (*big.Int, error) {
	return _Spons.Contract.GetNum(&_Spons.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Spons *SponsCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Spons.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Spons *SponsSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Spons.Contract.GetReference(&_Spons.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Spons *SponsCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Spons.Contract.GetReference(&_Spons.CallOpts, _index)
}

// GetSM is a free data retrieval call binding the contract method 0x256ec300.
//
// Solidity: function getSM(bytes32 _ref) view returns(bytes32, (uint256,uint256,string))
func (_Spons *SponsCaller) GetSM(opts *bind.CallOpts, _ref [32]byte) ([32]byte, PricePerSM, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(PricePerSM)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Spons.contract.Call(opts, out, "getSM", _ref)
	return *ret0, *ret1, err
}

// GetSM is a free data retrieval call binding the contract method 0x256ec300.
//
// Solidity: function getSM(bytes32 _ref) view returns(bytes32, (uint256,uint256,string))
func (_Spons *SponsSession) GetSM(_ref [32]byte) ([32]byte, PricePerSM, error) {
	return _Spons.Contract.GetSM(&_Spons.CallOpts, _ref)
}

// GetSM is a free data retrieval call binding the contract method 0x256ec300.
//
// Solidity: function getSM(bytes32 _ref) view returns(bytes32, (uint256,uint256,string))
func (_Spons *SponsCallerSession) GetSM(_ref [32]byte) ([32]byte, PricePerSM, error) {
	return _Spons.Contract.GetSM(&_Spons.CallOpts, _ref)
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _ref) view returns(uint8)
func (_Spons *SponsCaller) GetType(opts *bind.CallOpts, _ref [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Spons.contract.Call(opts, out, "getType", _ref)
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _ref) view returns(uint8)
func (_Spons *SponsSession) GetType(_ref [32]byte) (uint8, error) {
	return _Spons.Contract.GetType(&_Spons.CallOpts, _ref)
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _ref) view returns(uint8)
func (_Spons *SponsCallerSession) GetType(_ref [32]byte) (uint8, error) {
	return _Spons.Contract.GetType(&_Spons.CallOpts, _ref)
}

// AddCostModel is a paid mutator transaction binding the contract method 0xeafd03bb.
//
// Solidity: function addCostModel(bytes32 _ref, (uint8,uint256,string) _node) returns()
func (_Spons *SponsTransactor) AddCostModel(opts *bind.TransactOpts, _ref [32]byte, _node CostModel) (*types.Transaction, error) {
	return _Spons.contract.Transact(opts, "addCostModel", _ref, _node)
}

// AddCostModel is a paid mutator transaction binding the contract method 0xeafd03bb.
//
// Solidity: function addCostModel(bytes32 _ref, (uint8,uint256,string) _node) returns()
func (_Spons *SponsSession) AddCostModel(_ref [32]byte, _node CostModel) (*types.Transaction, error) {
	return _Spons.Contract.AddCostModel(&_Spons.TransactOpts, _ref, _node)
}

// AddCostModel is a paid mutator transaction binding the contract method 0xeafd03bb.
//
// Solidity: function addCostModel(bytes32 _ref, (uint8,uint256,string) _node) returns()
func (_Spons *SponsTransactorSession) AddCostModel(_ref [32]byte, _node CostModel) (*types.Transaction, error) {
	return _Spons.Contract.AddCostModel(&_Spons.TransactOpts, _ref, _node)
}

// AddFunctionalUnit is a paid mutator transaction binding the contract method 0x71126457.
//
// Solidity: function addFunctionalUnit(bytes32 _ref, (uint8,uint256,uint256,string) _node) returns()
func (_Spons *SponsTransactor) AddFunctionalUnit(opts *bind.TransactOpts, _ref [32]byte, _node PricePerFunctionalUnit) (*types.Transaction, error) {
	return _Spons.contract.Transact(opts, "addFunctionalUnit", _ref, _node)
}

// AddFunctionalUnit is a paid mutator transaction binding the contract method 0x71126457.
//
// Solidity: function addFunctionalUnit(bytes32 _ref, (uint8,uint256,uint256,string) _node) returns()
func (_Spons *SponsSession) AddFunctionalUnit(_ref [32]byte, _node PricePerFunctionalUnit) (*types.Transaction, error) {
	return _Spons.Contract.AddFunctionalUnit(&_Spons.TransactOpts, _ref, _node)
}

// AddFunctionalUnit is a paid mutator transaction binding the contract method 0x71126457.
//
// Solidity: function addFunctionalUnit(bytes32 _ref, (uint8,uint256,uint256,string) _node) returns()
func (_Spons *SponsTransactorSession) AddFunctionalUnit(_ref [32]byte, _node PricePerFunctionalUnit) (*types.Transaction, error) {
	return _Spons.Contract.AddFunctionalUnit(&_Spons.TransactOpts, _ref, _node)
}

// AddMeasuredWorks is a paid mutator transaction binding the contract method 0x56b8ecb1.
//
// Solidity: function addMeasuredWorks(bytes32 _ref, (uint8,uint256,uint256,uint256,uint256,uint256,string) _node) returns()
func (_Spons *SponsTransactor) AddMeasuredWorks(opts *bind.TransactOpts, _ref [32]byte, _node PriceForMeasuredWorks) (*types.Transaction, error) {
	return _Spons.contract.Transact(opts, "addMeasuredWorks", _ref, _node)
}

// AddMeasuredWorks is a paid mutator transaction binding the contract method 0x56b8ecb1.
//
// Solidity: function addMeasuredWorks(bytes32 _ref, (uint8,uint256,uint256,uint256,uint256,uint256,string) _node) returns()
func (_Spons *SponsSession) AddMeasuredWorks(_ref [32]byte, _node PriceForMeasuredWorks) (*types.Transaction, error) {
	return _Spons.Contract.AddMeasuredWorks(&_Spons.TransactOpts, _ref, _node)
}

// AddMeasuredWorks is a paid mutator transaction binding the contract method 0x56b8ecb1.
//
// Solidity: function addMeasuredWorks(bytes32 _ref, (uint8,uint256,uint256,uint256,uint256,uint256,string) _node) returns()
func (_Spons *SponsTransactorSession) AddMeasuredWorks(_ref [32]byte, _node PriceForMeasuredWorks) (*types.Transaction, error) {
	return _Spons.Contract.AddMeasuredWorks(&_Spons.TransactOpts, _ref, _node)
}

// AddSM is a paid mutator transaction binding the contract method 0x62b119f9.
//
// Solidity: function addSM(bytes32 _ref, (uint256,uint256,string) _node) returns()
func (_Spons *SponsTransactor) AddSM(opts *bind.TransactOpts, _ref [32]byte, _node PricePerSM) (*types.Transaction, error) {
	return _Spons.contract.Transact(opts, "addSM", _ref, _node)
}

// AddSM is a paid mutator transaction binding the contract method 0x62b119f9.
//
// Solidity: function addSM(bytes32 _ref, (uint256,uint256,string) _node) returns()
func (_Spons *SponsSession) AddSM(_ref [32]byte, _node PricePerSM) (*types.Transaction, error) {
	return _Spons.Contract.AddSM(&_Spons.TransactOpts, _ref, _node)
}

// AddSM is a paid mutator transaction binding the contract method 0x62b119f9.
//
// Solidity: function addSM(bytes32 _ref, (uint256,uint256,string) _node) returns()
func (_Spons *SponsTransactorSession) AddSM(_ref [32]byte, _node PricePerSM) (*types.Transaction, error) {
	return _Spons.Contract.AddSM(&_Spons.TransactOpts, _ref, _node)
}
