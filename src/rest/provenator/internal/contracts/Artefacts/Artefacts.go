// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Artefacts

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

// CreativeEntities is an auto generated low-level Go binding around an user-defined struct.
type CreativeEntities struct {
	Id    [32]byte
	Name  string
	Email string
	Url   string
}

// CreativeWorks is an auto generated low-level Go binding around an user-defined struct.
type CreativeWorks struct {
	WorkType     uint8
	License      uint8
	Id           [32]byte
	AuthorId     [32]byte
	DateCreated  [32]byte
	DateModified [32]byte
	Url          string
	Name         string
	Description  string
}

// ArtefactNodeABI is the input ABI used to generate the binding from.
const ArtefactNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"addAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"addCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"addPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthors\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCopyrightHolders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublishers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isCopyrightHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isPublisher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"removeAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"removeCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"removePublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArtefactNodeFuncSigs maps the 4-byte function signature to its string representation.
var ArtefactNodeFuncSigs = map[string]string{
	"a9c56065": "addAuthor(bytes32)",
	"c3a30a6f": "addCopyrightHolder(bytes32)",
	"c28b16dc": "addPublisher(bytes32)",
	"9673484e": "amend((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))",
	"6d4ce63c": "get()",
	"7fc622f4": "getAuthors()",
	"d6d54164": "getCopyrightHolders()",
	"6c6071aa": "getPublishers()",
	"59c0161c": "isAuthor(bytes32)",
	"476caa00": "isCopyrightHolder(bytes32)",
	"240b5600": "isPublisher(bytes32)",
	"33966d03": "removeAuthor(bytes32)",
	"2424dccc": "removeCopyrightHolder(bytes32)",
	"d13c8f23": "removePublisher(bytes32)",
}

// ArtefactNodeBin is the compiled bytecode used for deploying new contracts.
var ArtefactNodeBin = "0x60806040523480156200001157600080fd5b50604051620013ce380380620013ce833981016040819052620000349162000338565b604081015160001a60f81b6001600160f81b031916158015906200006b5750606081015160001a60f81b6001600160f81b03191615155b80156200008557506000815160158111156200008357fe5b115b80156200009f57506015815160158111156200009d57fe5b105b8015620000bc5750601481602001516014811115620000ba57fe5b105b8015620000dc5750608081015160001a60f81b6001600160f81b03191615155b8015620000ee575060008160e0015151115b8015620001015750600081610100015151115b6200010b57600080fd5b805160008054839290829060ff191660018360158111156200012957fe5b021790555060208201518154829061ff0019166101008360148111156200014c57fe5b021790555060408201516001820155606082015160028201556080820151600382015560a0820151600482015560c0820151805162000196916005840191602090910190620001de565b5060e08201518051620001b4916006840191602090910190620001de565b506101008201518051620001d3916007840191602090910190620001de565b509050505062000475565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022157805160ff191683800117855562000251565b8280016001018555821562000251579182015b828111156200025157825182559160200191906001019062000234565b506200025f92915062000263565b5090565b6200028091905b808211156200025f57600081556001016200026a565b90565b8051601581106200029357600080fd5b92915050565b8051601681106200029357600080fd5b600082601f830112620002ba578081fd5b81516001600160401b03811115620002d0578182fd5b6020620002e6601f8301601f191682016200044e565b92508183528481838601011115620002fd57600080fd5b60005b828110156200031d57848101820151848201830152810162000300565b828111156200032f5760008284860101525b50505092915050565b6000602082840312156200034a578081fd5b81516001600160401b038082111562000361578283fd5b61012091840180860383131562000376578384fd5b62000381836200044e565b6200038d878362000299565b81526200039e876020840162000283565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c0820151935082841115620003df578485fd5b620003ed87858401620002a9565b60c082015260e082015193508284111562000406578485fd5b6200041487858401620002a9565b60e0820152610100935083820151838111156200042f578586fd5b6200043d88828501620002a9565b948201949094529695505050505050565b6040518181016001600160401b03811182821017156200046d57600080fd5b604052919050565b610f4980620004856000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637fc622f41161008c578063c28b16dc11610066578063c28b16dc146101be578063c3a30a6f146101d1578063d13c8f23146101e4578063d6d54164146101f7576100ea565b80637fc622f4146101905780639673484e14610198578063a9c56065146101ab576100ea565b8063476caa00116100c8578063476caa001461014057806359c0161c146101535780636c6071aa146101665780636d4ce63c1461017b576100ea565b8063240b5600146100ef5780632424dccc1461011857806333966d031461012d575b600080fd5b6101026100fd366004610c6c565b6101ff565b60405161010f9190610e30565b60405180910390f35b61012b610126366004610c6c565b61025f565b005b61012b61013b366004610c6c565b610333565b61010261014e366004610c6c565b6103e8565b610102610161366004610c6c565b610441565b61016e61049a565b60405161010f9190610dec565b6101836104f3565b60405161010f9190610e3b565b61016e610749565b61012b6101a6366004610c84565b61079f565b61012b6101b9366004610c6c565b610923565b61012b6101cc366004610c6c565b610969565b61012b6101df366004610c6c565b6109ab565b61012b6101f2366004610c6c565b6109ed565b61016e610aa2565b600081811a60f81b6001600160f81b03191661021a57600080fd5b6000805b600a548110156102585783600a828154811061023657fe5b906000526020600020015414156102505760019150610258565b60010161021e565b5092915050565b8060001a60f81b6001600160f81b03191661027957600080fd5b60005b60095481101561032f57816009828154811061029457fe5b90600052602060002001541415610327576009546000190181101561030157805b600954600019018110156102ff57600981600101815481106102d357fe5b9060005260206000200154600982815481106102eb57fe5b6000918252602090912001556001016102b5565b505b600980548061030c57fe5b6001900381819060005260206000200160009055905561032f565b60010161027c565b5050565b8060001a60f81b6001600160f81b03191661034d57600080fd5b60005b60085481101561032f57816008828154811061036857fe5b906000526020600020015414156103e057600854600019018110156103d557805b600854600019018110156103d357600881600101815481106103a757fe5b9060005260206000200154600882815481106103bf57fe5b600091825260209091200155600101610389565b505b600880548061030c57fe5b600101610350565b600081811a60f81b6001600160f81b03191661040357600080fd5b6000805b60095481101561025857836009828154811061041f57fe5b906000526020600020015414156104395760019150610258565b600101610407565b600081811a60f81b6001600160f81b03191661045c57600080fd5b6000805b60085481101561025857836008828154811061047857fe5b906000526020600020015414156104925760019150610258565b600101610460565b6060600a8054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020905b8154815260200190600101908083116104d4575b505050505090505b90565b6104fb610af8565b60408051610120810190915260008054829060ff16601581111561051b57fe5b601581111561052657fe5b81528154602090910190610100900460ff16601481111561054357fe5b601481111561054e57fe5b815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106135780601f106105e857610100808354040283529160200191610613565b820191906000526020600020905b8154815290600101906020018083116105f657829003601f168201915b505050918352505060068201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b505050918352505060078201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561073b5780601f106107105761010080835404028352916020019161073b565b820191906000526020600020905b81548152906001019060200180831161071e57829003601f168201915b505050505081525050905090565b606060088054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020908154815260200190600101908083116104d4575050505050905090565b60408101516001541480156107c75750606081015160001a60f81b6001600160f81b03191615155b80156107df57506000815160158111156107dd57fe5b115b80156107f757506015815160158111156107f557fe5b105b8015610812575060148160200151601481111561081057fe5b105b80156108315750608081015160001a60f81b6001600160f81b03191615155b8015610842575060008160e0015151115b80156108545750600081610100015151115b61085d57600080fd5b805160008054839290829060ff1916600183601581111561087a57fe5b021790555060208201518154829061ff00191661010083601481111561089c57fe5b021790555060408201516001820155606082015160028201556080820151600382015560a0820151600482015560c082015180516108e4916005840191602090910190610b47565b5060e08201518051610900916006840191602090910190610b47565b50610100820151805161091d916007840191602090910190610b47565b50505050565b61092c81610441565b61096657600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018190555b50565b610972816101ff565b61096657600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80155565b6109b4816103e8565b61096657600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0155565b8060001a60f81b6001600160f81b031916610a0757600080fd5b60005b600a5481101561032f5781600a8281548110610a2257fe5b90600052602060002001541415610a9a57600a5460001901811015610a8f57805b600a5460001901811015610a8d57600a8160010181548110610a6157fe5b9060005260206000200154600a8281548110610a7957fe5b600091825260209091200155600101610a43565b505b600a80548061030c57fe5b600101610a0a565b606060098054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020908154815260200190600101908083116104d4575050505050905090565b604080516101208101909152806000815260200160008152600060208201819052604082018190526060808301829052608083019190915260a0820181905260c0820181905260e09091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b8857805160ff1916838001178555610bb5565b82800160010185558215610bb5579182015b82811115610bb5578251825591602001919060010190610b9a565b50610bc1929150610bc5565b5090565b6104f091905b80821115610bc15760008155600101610bcb565b803560158110610bee57600080fd5b92915050565b803560168110610bee57600080fd5b600082601f830112610c13578081fd5b813567ffffffffffffffff811115610c29578182fd5b610c3c601f8201601f1916602001610eec565b9150808252836020828501011115610c5357600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610c7d578081fd5b5035919050565b600060208284031215610c95578081fd5b813567ffffffffffffffff80821115610cac578283fd5b610120918401808603831315610cc0578384fd5b610cc983610eec565b610cd38783610bf4565b8152610ce28760208401610bdf565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c0820135935082841115610d22578485fd5b610d2e87858401610c03565b60c082015260e0820135935082841115610d46578485fd5b610d5287858401610c03565b60e082015261010093508382013583811115610d6c578586fd5b610d7888828501610c03565b948201949094529695505050505050565b60158110610d9357fe5b9052565b60168110610d9357fe5b60008151808452815b81811015610dc657602081850181015186830182015201610daa565b81811115610dd75782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610e2457835183529284019291840191600101610e08565b50909695505050505050565b901515815260200190565b600060208252610e4f602083018451610d97565b6020830151610e616040840182610d89565b506040830151606083015260608301516080830152608083015160a083015260a083015160c083015260c08301516101208060e0850152610ea6610140850183610da1565b60e08601519250601f19610100818784030181880152610ec68386610da1565b818901519550828882030185890152610edf8187610da1565b9998505050505050505050565b60405181810167ffffffffffffffff81118282101715610f0b57600080fd5b60405291905056fea26469706673582212200add9a2229093fd51d68df0d12b54d1094b456591164167352eb1893863d88f864736f6c63430006080033"

// DeployArtefactNode deploys a new Ethereum contract, binding an instance of ArtefactNode to it.
func DeployArtefactNode(auth *bind.TransactOpts, backend bind.ContractBackend, _work CreativeWorks) (common.Address, *types.Transaction, *ArtefactNode, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtefactNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArtefactNodeBin), backend, _work)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArtefactNode{ArtefactNodeCaller: ArtefactNodeCaller{contract: contract}, ArtefactNodeTransactor: ArtefactNodeTransactor{contract: contract}, ArtefactNodeFilterer: ArtefactNodeFilterer{contract: contract}}, nil
}

// ArtefactNode is an auto generated Go binding around an Ethereum contract.
type ArtefactNode struct {
	ArtefactNodeCaller     // Read-only binding to the contract
	ArtefactNodeTransactor // Write-only binding to the contract
	ArtefactNodeFilterer   // Log filterer for contract events
}

// ArtefactNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArtefactNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtefactNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtefactNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtefactNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtefactNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtefactNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtefactNodeSession struct {
	Contract     *ArtefactNode     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtefactNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtefactNodeCallerSession struct {
	Contract *ArtefactNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ArtefactNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtefactNodeTransactorSession struct {
	Contract     *ArtefactNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ArtefactNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArtefactNodeRaw struct {
	Contract *ArtefactNode // Generic contract binding to access the raw methods on
}

// ArtefactNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtefactNodeCallerRaw struct {
	Contract *ArtefactNodeCaller // Generic read-only contract binding to access the raw methods on
}

// ArtefactNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtefactNodeTransactorRaw struct {
	Contract *ArtefactNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArtefactNode creates a new instance of ArtefactNode, bound to a specific deployed contract.
func NewArtefactNode(address common.Address, backend bind.ContractBackend) (*ArtefactNode, error) {
	contract, err := bindArtefactNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArtefactNode{ArtefactNodeCaller: ArtefactNodeCaller{contract: contract}, ArtefactNodeTransactor: ArtefactNodeTransactor{contract: contract}, ArtefactNodeFilterer: ArtefactNodeFilterer{contract: contract}}, nil
}

// NewArtefactNodeCaller creates a new read-only instance of ArtefactNode, bound to a specific deployed contract.
func NewArtefactNodeCaller(address common.Address, caller bind.ContractCaller) (*ArtefactNodeCaller, error) {
	contract, err := bindArtefactNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtefactNodeCaller{contract: contract}, nil
}

// NewArtefactNodeTransactor creates a new write-only instance of ArtefactNode, bound to a specific deployed contract.
func NewArtefactNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*ArtefactNodeTransactor, error) {
	contract, err := bindArtefactNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtefactNodeTransactor{contract: contract}, nil
}

// NewArtefactNodeFilterer creates a new log filterer instance of ArtefactNode, bound to a specific deployed contract.
func NewArtefactNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*ArtefactNodeFilterer, error) {
	contract, err := bindArtefactNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtefactNodeFilterer{contract: contract}, nil
}

// bindArtefactNode binds a generic wrapper to an already deployed contract.
func bindArtefactNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtefactNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtefactNode *ArtefactNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArtefactNode.Contract.ArtefactNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtefactNode *ArtefactNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtefactNode.Contract.ArtefactNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtefactNode *ArtefactNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtefactNode.Contract.ArtefactNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArtefactNode *ArtefactNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArtefactNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArtefactNode *ArtefactNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArtefactNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArtefactNode *ArtefactNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArtefactNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_ArtefactNode *ArtefactNodeCaller) Get(opts *bind.CallOpts) (CreativeWorks, error) {
	var (
		ret0 = new(CreativeWorks)
	)
	out := ret0
	err := _ArtefactNode.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_ArtefactNode *ArtefactNodeSession) Get() (CreativeWorks, error) {
	return _ArtefactNode.Contract.Get(&_ArtefactNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_ArtefactNode *ArtefactNodeCallerSession) Get() (CreativeWorks, error) {
	return _ArtefactNode.Contract.Get(&_ArtefactNode.CallOpts)
}

// GetAuthors is a free data retrieval call binding the contract method 0x7fc622f4.
//
// Solidity: function getAuthors() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeCaller) GetAuthors(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ArtefactNode.contract.Call(opts, out, "getAuthors")
	return *ret0, err
}

// GetAuthors is a free data retrieval call binding the contract method 0x7fc622f4.
//
// Solidity: function getAuthors() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeSession) GetAuthors() ([][32]byte, error) {
	return _ArtefactNode.Contract.GetAuthors(&_ArtefactNode.CallOpts)
}

// GetAuthors is a free data retrieval call binding the contract method 0x7fc622f4.
//
// Solidity: function getAuthors() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeCallerSession) GetAuthors() ([][32]byte, error) {
	return _ArtefactNode.Contract.GetAuthors(&_ArtefactNode.CallOpts)
}

// GetCopyrightHolders is a free data retrieval call binding the contract method 0xd6d54164.
//
// Solidity: function getCopyrightHolders() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeCaller) GetCopyrightHolders(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ArtefactNode.contract.Call(opts, out, "getCopyrightHolders")
	return *ret0, err
}

// GetCopyrightHolders is a free data retrieval call binding the contract method 0xd6d54164.
//
// Solidity: function getCopyrightHolders() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeSession) GetCopyrightHolders() ([][32]byte, error) {
	return _ArtefactNode.Contract.GetCopyrightHolders(&_ArtefactNode.CallOpts)
}

// GetCopyrightHolders is a free data retrieval call binding the contract method 0xd6d54164.
//
// Solidity: function getCopyrightHolders() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeCallerSession) GetCopyrightHolders() ([][32]byte, error) {
	return _ArtefactNode.Contract.GetCopyrightHolders(&_ArtefactNode.CallOpts)
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeCaller) GetPublishers(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ArtefactNode.contract.Call(opts, out, "getPublishers")
	return *ret0, err
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeSession) GetPublishers() ([][32]byte, error) {
	return _ArtefactNode.Contract.GetPublishers(&_ArtefactNode.CallOpts)
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(bytes32[])
func (_ArtefactNode *ArtefactNodeCallerSession) GetPublishers() ([][32]byte, error) {
	return _ArtefactNode.Contract.GetPublishers(&_ArtefactNode.CallOpts)
}

// IsAuthor is a free data retrieval call binding the contract method 0x59c0161c.
//
// Solidity: function isAuthor(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeCaller) IsAuthor(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArtefactNode.contract.Call(opts, out, "isAuthor", _id)
	return *ret0, err
}

// IsAuthor is a free data retrieval call binding the contract method 0x59c0161c.
//
// Solidity: function isAuthor(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeSession) IsAuthor(_id [32]byte) (bool, error) {
	return _ArtefactNode.Contract.IsAuthor(&_ArtefactNode.CallOpts, _id)
}

// IsAuthor is a free data retrieval call binding the contract method 0x59c0161c.
//
// Solidity: function isAuthor(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeCallerSession) IsAuthor(_id [32]byte) (bool, error) {
	return _ArtefactNode.Contract.IsAuthor(&_ArtefactNode.CallOpts, _id)
}

// IsCopyrightHolder is a free data retrieval call binding the contract method 0x476caa00.
//
// Solidity: function isCopyrightHolder(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeCaller) IsCopyrightHolder(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArtefactNode.contract.Call(opts, out, "isCopyrightHolder", _id)
	return *ret0, err
}

// IsCopyrightHolder is a free data retrieval call binding the contract method 0x476caa00.
//
// Solidity: function isCopyrightHolder(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeSession) IsCopyrightHolder(_id [32]byte) (bool, error) {
	return _ArtefactNode.Contract.IsCopyrightHolder(&_ArtefactNode.CallOpts, _id)
}

// IsCopyrightHolder is a free data retrieval call binding the contract method 0x476caa00.
//
// Solidity: function isCopyrightHolder(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeCallerSession) IsCopyrightHolder(_id [32]byte) (bool, error) {
	return _ArtefactNode.Contract.IsCopyrightHolder(&_ArtefactNode.CallOpts, _id)
}

// IsPublisher is a free data retrieval call binding the contract method 0x240b5600.
//
// Solidity: function isPublisher(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeCaller) IsPublisher(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArtefactNode.contract.Call(opts, out, "isPublisher", _id)
	return *ret0, err
}

// IsPublisher is a free data retrieval call binding the contract method 0x240b5600.
//
// Solidity: function isPublisher(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeSession) IsPublisher(_id [32]byte) (bool, error) {
	return _ArtefactNode.Contract.IsPublisher(&_ArtefactNode.CallOpts, _id)
}

// IsPublisher is a free data retrieval call binding the contract method 0x240b5600.
//
// Solidity: function isPublisher(bytes32 _id) view returns(bool)
func (_ArtefactNode *ArtefactNodeCallerSession) IsPublisher(_id [32]byte) (bool, error) {
	return _ArtefactNode.Contract.IsPublisher(&_ArtefactNode.CallOpts, _id)
}

// AddAuthor is a paid mutator transaction binding the contract method 0xa9c56065.
//
// Solidity: function addAuthor(bytes32 _authorId) returns()
func (_ArtefactNode *ArtefactNodeTransactor) AddAuthor(opts *bind.TransactOpts, _authorId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "addAuthor", _authorId)
}

// AddAuthor is a paid mutator transaction binding the contract method 0xa9c56065.
//
// Solidity: function addAuthor(bytes32 _authorId) returns()
func (_ArtefactNode *ArtefactNodeSession) AddAuthor(_authorId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddAuthor(&_ArtefactNode.TransactOpts, _authorId)
}

// AddAuthor is a paid mutator transaction binding the contract method 0xa9c56065.
//
// Solidity: function addAuthor(bytes32 _authorId) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) AddAuthor(_authorId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddAuthor(&_ArtefactNode.TransactOpts, _authorId)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0xc3a30a6f.
//
// Solidity: function addCopyrightHolder(bytes32 _copyrightHolderId) returns()
func (_ArtefactNode *ArtefactNodeTransactor) AddCopyrightHolder(opts *bind.TransactOpts, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "addCopyrightHolder", _copyrightHolderId)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0xc3a30a6f.
//
// Solidity: function addCopyrightHolder(bytes32 _copyrightHolderId) returns()
func (_ArtefactNode *ArtefactNodeSession) AddCopyrightHolder(_copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddCopyrightHolder(&_ArtefactNode.TransactOpts, _copyrightHolderId)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0xc3a30a6f.
//
// Solidity: function addCopyrightHolder(bytes32 _copyrightHolderId) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) AddCopyrightHolder(_copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddCopyrightHolder(&_ArtefactNode.TransactOpts, _copyrightHolderId)
}

// AddPublisher is a paid mutator transaction binding the contract method 0xc28b16dc.
//
// Solidity: function addPublisher(bytes32 _publisherId) returns()
func (_ArtefactNode *ArtefactNodeTransactor) AddPublisher(opts *bind.TransactOpts, _publisherId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "addPublisher", _publisherId)
}

// AddPublisher is a paid mutator transaction binding the contract method 0xc28b16dc.
//
// Solidity: function addPublisher(bytes32 _publisherId) returns()
func (_ArtefactNode *ArtefactNodeSession) AddPublisher(_publisherId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddPublisher(&_ArtefactNode.TransactOpts, _publisherId)
}

// AddPublisher is a paid mutator transaction binding the contract method 0xc28b16dc.
//
// Solidity: function addPublisher(bytes32 _publisherId) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) AddPublisher(_publisherId [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddPublisher(&_ArtefactNode.TransactOpts, _publisherId)
}

// Amend is a paid mutator transaction binding the contract method 0x9673484e.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_ArtefactNode *ArtefactNodeTransactor) Amend(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "amend", _work)
}

// Amend is a paid mutator transaction binding the contract method 0x9673484e.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_ArtefactNode *ArtefactNodeSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.Contract.Amend(&_ArtefactNode.TransactOpts, _work)
}

// Amend is a paid mutator transaction binding the contract method 0x9673484e.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.Contract.Amend(&_ArtefactNode.TransactOpts, _work)
}

// RemoveAuthor is a paid mutator transaction binding the contract method 0x33966d03.
//
// Solidity: function removeAuthor(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeTransactor) RemoveAuthor(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "removeAuthor", _id)
}

// RemoveAuthor is a paid mutator transaction binding the contract method 0x33966d03.
//
// Solidity: function removeAuthor(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeSession) RemoveAuthor(_id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.RemoveAuthor(&_ArtefactNode.TransactOpts, _id)
}

// RemoveAuthor is a paid mutator transaction binding the contract method 0x33966d03.
//
// Solidity: function removeAuthor(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) RemoveAuthor(_id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.RemoveAuthor(&_ArtefactNode.TransactOpts, _id)
}

// RemoveCopyrightHolder is a paid mutator transaction binding the contract method 0x2424dccc.
//
// Solidity: function removeCopyrightHolder(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeTransactor) RemoveCopyrightHolder(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "removeCopyrightHolder", _id)
}

// RemoveCopyrightHolder is a paid mutator transaction binding the contract method 0x2424dccc.
//
// Solidity: function removeCopyrightHolder(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeSession) RemoveCopyrightHolder(_id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.RemoveCopyrightHolder(&_ArtefactNode.TransactOpts, _id)
}

// RemoveCopyrightHolder is a paid mutator transaction binding the contract method 0x2424dccc.
//
// Solidity: function removeCopyrightHolder(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) RemoveCopyrightHolder(_id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.RemoveCopyrightHolder(&_ArtefactNode.TransactOpts, _id)
}

// RemovePublisher is a paid mutator transaction binding the contract method 0xd13c8f23.
//
// Solidity: function removePublisher(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeTransactor) RemovePublisher(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "removePublisher", _id)
}

// RemovePublisher is a paid mutator transaction binding the contract method 0xd13c8f23.
//
// Solidity: function removePublisher(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeSession) RemovePublisher(_id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.RemovePublisher(&_ArtefactNode.TransactOpts, _id)
}

// RemovePublisher is a paid mutator transaction binding the contract method 0xd13c8f23.
//
// Solidity: function removePublisher(bytes32 _id) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) RemovePublisher(_id [32]byte) (*types.Transaction, error) {
	return _ArtefactNode.Contract.RemovePublisher(&_ArtefactNode.TransactOpts, _id)
}

// ArtefactsABI is the input ABI used to generate the binding from.
const ArtefactsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"addWorkAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"addWorkCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"addWorkPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amendWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReferences\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWork\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkAuthors\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkCopyrightHolders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkPublishers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"isWorkAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"isWorkCopyrightHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"isWorkPublisher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"removeWorkAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"removeWorkCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"removeWorkPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArtefactsFuncSigs maps the 4-byte function signature to its string representation.
var ArtefactsFuncSigs = map[string]string{
	"1d339e24": "addWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))",
	"55884f0f": "addWorkAuthor(bytes32,bytes32)",
	"0f1ee4a2": "addWorkCopyrightHolder(bytes32,bytes32)",
	"e17989ea": "addWorkPublisher(bytes32,bytes32)",
	"4e542a6e": "amendWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"7a6337fa": "getReferences()",
	"30266537": "getWork(bytes32)",
	"426aa803": "getWorkAuthors(bytes32)",
	"4d2522a2": "getWorkContract(bytes32)",
	"86105c65": "getWorkCopyrightHolders(bytes32)",
	"5ffd09d8": "getWorkPublishers(bytes32)",
	"b2e459cf": "isWorkAuthor(bytes32,bytes32)",
	"f6a50965": "isWorkCopyrightHolder(bytes32,bytes32)",
	"ef90d9c9": "isWorkPublisher(bytes32,bytes32)",
	"e229e173": "removeWorkAuthor(bytes32,bytes32)",
	"521ffdce": "removeWorkCopyrightHolder(bytes32,bytes32)",
	"0c99420f": "removeWorkPublisher(bytes32,bytes32)",
}

// ArtefactsBin is the compiled bytecode used for deploying new contracts.
var ArtefactsBin = "0x608060405234801561001057600080fd5b5060006002556126f4806100256000396000f3fe60806040523480156200001157600080fd5b50600436106200013c5760003560e01c80635ffd09d811620000bd578063bc599341116200007b578063bc59934114620002a8578063e17989ea14620002bf578063e229e17314620002d6578063ef90d9c914620002ed578063f6a509651462000304576200013c565b80635ffd09d8146200023157806367e0badb14620002485780637a6337fa146200026157806386105c65146200026b578063b2e459cf1462000282576200013c565b8063426aa803116200010b578063426aa80314620001b75780634d2522a214620001dd5780634e542a6e1462000203578063521ffdce146200014157806355884f0f146200021a576200013c565b80630c99420f14620001415780630f1ee4a2146200015a5780631d339e241462000171578063302665371462000188575b600080fd5b620001586200015236600462000e83565b6200031b565b005b620001586200016b36600462000e83565b620003bd565b620001586200018236600462000ea5565b62000426565b6200019f6200019936600462000e6a565b620004b4565b604051620001ae919062001178565b60405180910390f35b620001ce620001c836600462000e6a565b6200057a565b604051620001ae91906200111e565b620001f4620001ee36600462000e6a565b62000632565b604051620001ae91906200110a565b620001586200021436600462000ea5565b62000676565b620001586200022b36600462000e83565b6200071d565b620001ce6200024236600462000e6a565b62000786565b62000252620007ff565b604051620001ae91906200116f565b620001ce62000805565b620001ce6200027c36600462000e6a565b620008a9565b620002996200029336600462000e83565b62000922565b604051620001ae919062001164565b62000252620002b936600462000e6a565b620009e9565b62000158620002d036600462000e83565b62000a23565b62000158620002e736600462000e83565b62000a8c565b62000299620002fe36600462000e83565b62000af5565b620002996200031536600462000e83565b62000b5d565b6000828152602081905260409020600101546001600160a01b03166200034057600080fd5b60008281526020819052604090819020600101549051630909373360e21b81526001600160a01b03909116908190632424dccc90620003849085906004016200116f565b600060405180830381600087803b1580156200039f57600080fd5b505af1158015620003b4573d6000803e3d6000fd5b50505050505050565b6000828152602081905260409020600101546001600160a01b0316620003e257600080fd5b6000828152602081905260409081902060010154905163c3a30a6f60e01b81526001600160a01b0390911690819063c3a30a6f90620003849085906004016200116f565b60408101516200043f9060009063ffffffff62000bc516565b156200045657620004508162000676565b620004b1565b600081604051620004679062000c6e565b62000473919062001178565b604051809103906000f08015801562000490573d6000803e3d6000fd5b506040830151909150620004ae906000908363ffffffff62000bda16565b50505b50565b620004be62000c7c565b6000828152602081905260409020600101546001600160a01b0316620004e357600080fd5b600082815260208190526040808220600101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b1580156200053457600080fd5b505afa15801562000549573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000573919081019062000fbc565b9392505050565b6000818152602081905260409020600101546060906001600160a01b0316620005a257600080fd5b600082815260208190526040808220600101548151631ff188bd60e21b815291516001600160a01b03909116928392637fc622f49260048083019392829003018186803b158015620005f357600080fd5b505afa15801562000608573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000573919081019062000dac565b6000818152602081905260408120600101546001600160a01b03166200065757600080fd5b506000908152602081905260409020600101546001600160a01b031690565b6040808201516000908152602081905220600101546001600160a01b03166200069e57600080fd5b60408082015160009081526020819052819020600101549051634b39a42760e11b81526001600160a01b03909116908190639673484e90620006e590859060040162001178565b600060405180830381600087803b1580156200070057600080fd5b505af115801562000715573d6000803e3d6000fd5b505050505050565b6000828152602081905260409020600101546001600160a01b03166200074257600080fd5b6000828152602081905260409081902060010154905163a9c5606560e01b81526001600160a01b0390911690819063a9c5606590620003849085906004016200116f565b6000818152602081905260409020600101546060906001600160a01b0316620007ae57600080fd5b60008281526020819052604080822060010154815163363038d560e11b815291516001600160a01b03909116928392636c6071aa9260048083019392829003018186803b158015620005f357600080fd5b60025490565b60608060006002015467ffffffffffffffff811180156200082557600080fd5b5060405190808252806020026020018201604052801562000850578160200160208202803683370190505b50905060005b600254811015620008a35760018054829081106200087057fe5b9060005260206000209060020201600001548282815181106200088f57fe5b602090810291909101015260010162000856565b50905090565b6000818152602081905260409020600101546060906001600160a01b0316620008d157600080fd5b6000828152602081905260408082206001015481516335b5505960e21b815291516001600160a01b0390911692839263d6d541649260048083019392829003018186803b158015620005f357600080fd5b6000828152602081905260408120600101546001600160a01b03166200094757600080fd5b60008381526020819052604090819020600101549051631670058760e21b81526001600160a01b039091169081906359c0161c906200098b9086906004016200116f565b60206040518083038186803b158015620009a457600080fd5b505afa158015620009b9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620009df919062000e48565b9150505b92915050565b6002546000908210620009fb57600080fd5b600180548390811062000a0a57fe5b9060005260206000209060020201600001549050919050565b6000828152602081905260409020600101546001600160a01b031662000a4857600080fd5b600082815260208190526040908190206001015490516330a2c5b760e21b81526001600160a01b0390911690819063c28b16dc90620003849085906004016200116f565b6000828152602081905260409020600101546001600160a01b031662000ab157600080fd5b600082815260208190526040908190206001015490516333966d0360e01b81526001600160a01b039091169081906333966d0390620003849085906004016200116f565b6000828152602081905260408120600101546001600160a01b031662000b1a57600080fd5b60008381526020819052604090819020600101549051621205ab60e91b81526001600160a01b0390911690819063240b5600906200098b9086906004016200116f565b6000828152602081905260408120600101546001600160a01b031662000b8257600080fd5b600083815260208190526040908190206001015490516223b65560e91b81526001600160a01b0390911690819063476caa00906200098b9086906004016200116f565b60009081526020919091526040902054151590565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000c1a57600191505062000573565b506001808501805491820180825560008681526020889052604090208190558591908390811062000c4757fe5b60009182526020822060029182020192909255908601805460010190559150620005739050565b6113ce80620012f183390190565b604080516101208101909152806000815260200160008152600060208201819052604082018190526060808301829052608083019190915260a0820181905260c0820181905260e09091015290565b8035620009e381620012d4565b8051620009e381620012d4565b8035620009e381620012e2565b8051620009e381620012e2565b600082601f83011262000d10578081fd5b813562000d2762000d21826200127c565b62001233565b915080825283602082850101111562000d3f57600080fd5b8060208401602084013760009082016020015292915050565b600082601f83011262000d69578081fd5b815162000d7a62000d21826200127c565b915080825283602082850101111562000d9257600080fd5b62000da5816020840160208601620012a1565b5092915050565b6000602080838503121562000dbf578182fd5b825167ffffffffffffffff81111562000dd6578283fd5b80840185601f82011262000de8578384fd5b8051915062000dfb62000d21836200125b565b828152838101908285018585028401860189101562000e18578687fd5b8693505b8484101562000e3c57805183526001939093019291850191850162000e1c565b50979650505050505050565b60006020828403121562000e5a578081fd5b8151801515811462000573578182fd5b60006020828403121562000e7c578081fd5b5035919050565b6000806040838503121562000e96578081fd5b50508035926020909101359150565b60006020828403121562000eb7578081fd5b813567ffffffffffffffff8082111562000ecf578283fd5b61012091840180860383131562000ee4578384fd5b62000eef8362001233565b62000efb878362000ce5565b815262000f0c876020840162000ccb565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c082013593508284111562000f4d578485fd5b62000f5b8785840162000cff565b60c082015260e082013593508284111562000f74578485fd5b62000f828785840162000cff565b60e08201526101009350838201358381111562000f9d578586fd5b62000fab8882850162000cff565b948201949094529695505050505050565b60006020828403121562000fce578081fd5b815167ffffffffffffffff8082111562000fe6578283fd5b61012091840180860383131562000ffb578384fd5b620010068362001233565b62001012878362000cf2565b815262001023876020840162000cd8565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c082015193508284111562001064578485fd5b620010728785840162000d58565b60c082015260e08201519350828411156200108b578485fd5b620010998785840162000d58565b60e082015261010093508382015183811115620010b4578586fd5b62000fab8882850162000d58565b60158110620010cd57fe5b9052565b60168110620010cd57fe5b60008151808452620010f6816020860160208601620012a1565b601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b8181101562001158578351835292840192918401916001016200113a565b50909695505050505050565b901515815260200190565b90815260200190565b6000602082526200118e602083018451620010d1565b6020830151620011a26040840182620010c2565b506040830151606083015260608301516080830152608083015160a083015260a083015160c083015260c08301516101208060e0850152620011e9610140850183620010dc565b60e08601519250601f196101008187840301818801526200120b8386620010dc565b818901519550828882030185890152620012268187620010dc565b9998505050505050505050565b60405181810167ffffffffffffffff811182821017156200125357600080fd5b604052919050565b600067ffffffffffffffff82111562001272578081fd5b5060209081020190565b600067ffffffffffffffff82111562001293578081fd5b50601f01601f191660200190565b60005b83811015620012be578181015183820152602001620012a4565b83811115620012ce576000848401525b50505050565b60158110620004b157600080fd5b60168110620004b157600080fdfe60806040523480156200001157600080fd5b50604051620013ce380380620013ce833981016040819052620000349162000338565b604081015160001a60f81b6001600160f81b031916158015906200006b5750606081015160001a60f81b6001600160f81b03191615155b80156200008557506000815160158111156200008357fe5b115b80156200009f57506015815160158111156200009d57fe5b105b8015620000bc5750601481602001516014811115620000ba57fe5b105b8015620000dc5750608081015160001a60f81b6001600160f81b03191615155b8015620000ee575060008160e0015151115b8015620001015750600081610100015151115b6200010b57600080fd5b805160008054839290829060ff191660018360158111156200012957fe5b021790555060208201518154829061ff0019166101008360148111156200014c57fe5b021790555060408201516001820155606082015160028201556080820151600382015560a0820151600482015560c0820151805162000196916005840191602090910190620001de565b5060e08201518051620001b4916006840191602090910190620001de565b506101008201518051620001d3916007840191602090910190620001de565b509050505062000475565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022157805160ff191683800117855562000251565b8280016001018555821562000251579182015b828111156200025157825182559160200191906001019062000234565b506200025f92915062000263565b5090565b6200028091905b808211156200025f57600081556001016200026a565b90565b8051601581106200029357600080fd5b92915050565b8051601681106200029357600080fd5b600082601f830112620002ba578081fd5b81516001600160401b03811115620002d0578182fd5b6020620002e6601f8301601f191682016200044e565b92508183528481838601011115620002fd57600080fd5b60005b828110156200031d57848101820151848201830152810162000300565b828111156200032f5760008284860101525b50505092915050565b6000602082840312156200034a578081fd5b81516001600160401b038082111562000361578283fd5b61012091840180860383131562000376578384fd5b62000381836200044e565b6200038d878362000299565b81526200039e876020840162000283565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c0820151935082841115620003df578485fd5b620003ed87858401620002a9565b60c082015260e082015193508284111562000406578485fd5b6200041487858401620002a9565b60e0820152610100935083820151838111156200042f578586fd5b6200043d88828501620002a9565b948201949094529695505050505050565b6040518181016001600160401b03811182821017156200046d57600080fd5b604052919050565b610f4980620004856000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637fc622f41161008c578063c28b16dc11610066578063c28b16dc146101be578063c3a30a6f146101d1578063d13c8f23146101e4578063d6d54164146101f7576100ea565b80637fc622f4146101905780639673484e14610198578063a9c56065146101ab576100ea565b8063476caa00116100c8578063476caa001461014057806359c0161c146101535780636c6071aa146101665780636d4ce63c1461017b576100ea565b8063240b5600146100ef5780632424dccc1461011857806333966d031461012d575b600080fd5b6101026100fd366004610c6c565b6101ff565b60405161010f9190610e30565b60405180910390f35b61012b610126366004610c6c565b61025f565b005b61012b61013b366004610c6c565b610333565b61010261014e366004610c6c565b6103e8565b610102610161366004610c6c565b610441565b61016e61049a565b60405161010f9190610dec565b6101836104f3565b60405161010f9190610e3b565b61016e610749565b61012b6101a6366004610c84565b61079f565b61012b6101b9366004610c6c565b610923565b61012b6101cc366004610c6c565b610969565b61012b6101df366004610c6c565b6109ab565b61012b6101f2366004610c6c565b6109ed565b61016e610aa2565b600081811a60f81b6001600160f81b03191661021a57600080fd5b6000805b600a548110156102585783600a828154811061023657fe5b906000526020600020015414156102505760019150610258565b60010161021e565b5092915050565b8060001a60f81b6001600160f81b03191661027957600080fd5b60005b60095481101561032f57816009828154811061029457fe5b90600052602060002001541415610327576009546000190181101561030157805b600954600019018110156102ff57600981600101815481106102d357fe5b9060005260206000200154600982815481106102eb57fe5b6000918252602090912001556001016102b5565b505b600980548061030c57fe5b6001900381819060005260206000200160009055905561032f565b60010161027c565b5050565b8060001a60f81b6001600160f81b03191661034d57600080fd5b60005b60085481101561032f57816008828154811061036857fe5b906000526020600020015414156103e057600854600019018110156103d557805b600854600019018110156103d357600881600101815481106103a757fe5b9060005260206000200154600882815481106103bf57fe5b600091825260209091200155600101610389565b505b600880548061030c57fe5b600101610350565b600081811a60f81b6001600160f81b03191661040357600080fd5b6000805b60095481101561025857836009828154811061041f57fe5b906000526020600020015414156104395760019150610258565b600101610407565b600081811a60f81b6001600160f81b03191661045c57600080fd5b6000805b60085481101561025857836008828154811061047857fe5b906000526020600020015414156104925760019150610258565b600101610460565b6060600a8054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020905b8154815260200190600101908083116104d4575b505050505090505b90565b6104fb610af8565b60408051610120810190915260008054829060ff16601581111561051b57fe5b601581111561052657fe5b81528154602090910190610100900460ff16601481111561054357fe5b601481111561054e57fe5b815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106135780601f106105e857610100808354040283529160200191610613565b820191906000526020600020905b8154815290600101906020018083116105f657829003601f168201915b505050918352505060068201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b505050918352505060078201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561073b5780601f106107105761010080835404028352916020019161073b565b820191906000526020600020905b81548152906001019060200180831161071e57829003601f168201915b505050505081525050905090565b606060088054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020908154815260200190600101908083116104d4575050505050905090565b60408101516001541480156107c75750606081015160001a60f81b6001600160f81b03191615155b80156107df57506000815160158111156107dd57fe5b115b80156107f757506015815160158111156107f557fe5b105b8015610812575060148160200151601481111561081057fe5b105b80156108315750608081015160001a60f81b6001600160f81b03191615155b8015610842575060008160e0015151115b80156108545750600081610100015151115b61085d57600080fd5b805160008054839290829060ff1916600183601581111561087a57fe5b021790555060208201518154829061ff00191661010083601481111561089c57fe5b021790555060408201516001820155606082015160028201556080820151600382015560a0820151600482015560c082015180516108e4916005840191602090910190610b47565b5060e08201518051610900916006840191602090910190610b47565b50610100820151805161091d916007840191602090910190610b47565b50505050565b61092c81610441565b61096657600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018190555b50565b610972816101ff565b61096657600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80155565b6109b4816103e8565b61096657600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0155565b8060001a60f81b6001600160f81b031916610a0757600080fd5b60005b600a5481101561032f5781600a8281548110610a2257fe5b90600052602060002001541415610a9a57600a5460001901811015610a8f57805b600a5460001901811015610a8d57600a8160010181548110610a6157fe5b9060005260206000200154600a8281548110610a7957fe5b600091825260209091200155600101610a43565b505b600a80548061030c57fe5b600101610a0a565b606060098054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020908154815260200190600101908083116104d4575050505050905090565b604080516101208101909152806000815260200160008152600060208201819052604082018190526060808301829052608083019190915260a0820181905260c0820181905260e09091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b8857805160ff1916838001178555610bb5565b82800160010185558215610bb5579182015b82811115610bb5578251825591602001919060010190610b9a565b50610bc1929150610bc5565b5090565b6104f091905b80821115610bc15760008155600101610bcb565b803560158110610bee57600080fd5b92915050565b803560168110610bee57600080fd5b600082601f830112610c13578081fd5b813567ffffffffffffffff811115610c29578182fd5b610c3c601f8201601f1916602001610eec565b9150808252836020828501011115610c5357600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610c7d578081fd5b5035919050565b600060208284031215610c95578081fd5b813567ffffffffffffffff80821115610cac578283fd5b610120918401808603831315610cc0578384fd5b610cc983610eec565b610cd38783610bf4565b8152610ce28760208401610bdf565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c0820135935082841115610d22578485fd5b610d2e87858401610c03565b60c082015260e0820135935082841115610d46578485fd5b610d5287858401610c03565b60e082015261010093508382013583811115610d6c578586fd5b610d7888828501610c03565b948201949094529695505050505050565b60158110610d9357fe5b9052565b60168110610d9357fe5b60008151808452815b81811015610dc657602081850181015186830182015201610daa565b81811115610dd75782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610e2457835183529284019291840191600101610e08565b50909695505050505050565b901515815260200190565b600060208252610e4f602083018451610d97565b6020830151610e616040840182610d89565b506040830151606083015260608301516080830152608083015160a083015260a083015160c083015260c08301516101208060e0850152610ea6610140850183610da1565b60e08601519250601f19610100818784030181880152610ec68386610da1565b818901519550828882030185890152610edf8187610da1565b9998505050505050505050565b60405181810167ffffffffffffffff81118282101715610f0b57600080fd5b60405291905056fea26469706673582212200add9a2229093fd51d68df0d12b54d1094b456591164167352eb1893863d88f864736f6c63430006080033a2646970667358221220a8a59afe71faab5b23acd44b81b48c5f37f02901712525478459f060d28fa1ae64736f6c63430006080033"

// DeployArtefacts deploys a new Ethereum contract, binding an instance of Artefacts to it.
func DeployArtefacts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Artefacts, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtefactsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArtefactsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Artefacts{ArtefactsCaller: ArtefactsCaller{contract: contract}, ArtefactsTransactor: ArtefactsTransactor{contract: contract}, ArtefactsFilterer: ArtefactsFilterer{contract: contract}}, nil
}

// Artefacts is an auto generated Go binding around an Ethereum contract.
type Artefacts struct {
	ArtefactsCaller     // Read-only binding to the contract
	ArtefactsTransactor // Write-only binding to the contract
	ArtefactsFilterer   // Log filterer for contract events
}

// ArtefactsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArtefactsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtefactsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArtefactsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtefactsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArtefactsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArtefactsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArtefactsSession struct {
	Contract     *Artefacts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArtefactsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArtefactsCallerSession struct {
	Contract *ArtefactsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArtefactsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArtefactsTransactorSession struct {
	Contract     *ArtefactsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArtefactsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArtefactsRaw struct {
	Contract *Artefacts // Generic contract binding to access the raw methods on
}

// ArtefactsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArtefactsCallerRaw struct {
	Contract *ArtefactsCaller // Generic read-only contract binding to access the raw methods on
}

// ArtefactsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArtefactsTransactorRaw struct {
	Contract *ArtefactsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArtefacts creates a new instance of Artefacts, bound to a specific deployed contract.
func NewArtefacts(address common.Address, backend bind.ContractBackend) (*Artefacts, error) {
	contract, err := bindArtefacts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Artefacts{ArtefactsCaller: ArtefactsCaller{contract: contract}, ArtefactsTransactor: ArtefactsTransactor{contract: contract}, ArtefactsFilterer: ArtefactsFilterer{contract: contract}}, nil
}

// NewArtefactsCaller creates a new read-only instance of Artefacts, bound to a specific deployed contract.
func NewArtefactsCaller(address common.Address, caller bind.ContractCaller) (*ArtefactsCaller, error) {
	contract, err := bindArtefacts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArtefactsCaller{contract: contract}, nil
}

// NewArtefactsTransactor creates a new write-only instance of Artefacts, bound to a specific deployed contract.
func NewArtefactsTransactor(address common.Address, transactor bind.ContractTransactor) (*ArtefactsTransactor, error) {
	contract, err := bindArtefacts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArtefactsTransactor{contract: contract}, nil
}

// NewArtefactsFilterer creates a new log filterer instance of Artefacts, bound to a specific deployed contract.
func NewArtefactsFilterer(address common.Address, filterer bind.ContractFilterer) (*ArtefactsFilterer, error) {
	contract, err := bindArtefacts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArtefactsFilterer{contract: contract}, nil
}

// bindArtefacts binds a generic wrapper to an already deployed contract.
func bindArtefacts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtefactsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Artefacts *ArtefactsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Artefacts.Contract.ArtefactsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Artefacts *ArtefactsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artefacts.Contract.ArtefactsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Artefacts *ArtefactsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Artefacts.Contract.ArtefactsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Artefacts *ArtefactsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Artefacts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Artefacts *ArtefactsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Artefacts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Artefacts *ArtefactsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Artefacts.Contract.contract.Transact(opts, method, params...)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Artefacts *ArtefactsCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Artefacts *ArtefactsSession) GetNum() (*big.Int, error) {
	return _Artefacts.Contract.GetNum(&_Artefacts.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Artefacts *ArtefactsCallerSession) GetNum() (*big.Int, error) {
	return _Artefacts.Contract.GetNum(&_Artefacts.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Artefacts *ArtefactsCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Artefacts *ArtefactsSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Artefacts.Contract.GetReference(&_Artefacts.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Artefacts *ArtefactsCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Artefacts.Contract.GetReference(&_Artefacts.CallOpts, _index)
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_Artefacts *ArtefactsCaller) GetReferences(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getReferences")
	return *ret0, err
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_Artefacts *ArtefactsSession) GetReferences() ([][32]byte, error) {
	return _Artefacts.Contract.GetReferences(&_Artefacts.CallOpts)
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_Artefacts *ArtefactsCallerSession) GetReferences() ([][32]byte, error) {
	return _Artefacts.Contract.GetReferences(&_Artefacts.CallOpts)
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_Artefacts *ArtefactsCaller) GetWork(opts *bind.CallOpts, _id [32]byte) (CreativeWorks, error) {
	var (
		ret0 = new(CreativeWorks)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getWork", _id)
	return *ret0, err
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_Artefacts *ArtefactsSession) GetWork(_id [32]byte) (CreativeWorks, error) {
	return _Artefacts.Contract.GetWork(&_Artefacts.CallOpts, _id)
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_Artefacts *ArtefactsCallerSession) GetWork(_id [32]byte) (CreativeWorks, error) {
	return _Artefacts.Contract.GetWork(&_Artefacts.CallOpts, _id)
}

// GetWorkAuthors is a free data retrieval call binding the contract method 0x426aa803.
//
// Solidity: function getWorkAuthors(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsCaller) GetWorkAuthors(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getWorkAuthors", _id)
	return *ret0, err
}

// GetWorkAuthors is a free data retrieval call binding the contract method 0x426aa803.
//
// Solidity: function getWorkAuthors(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsSession) GetWorkAuthors(_id [32]byte) ([][32]byte, error) {
	return _Artefacts.Contract.GetWorkAuthors(&_Artefacts.CallOpts, _id)
}

// GetWorkAuthors is a free data retrieval call binding the contract method 0x426aa803.
//
// Solidity: function getWorkAuthors(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsCallerSession) GetWorkAuthors(_id [32]byte) ([][32]byte, error) {
	return _Artefacts.Contract.GetWorkAuthors(&_Artefacts.CallOpts, _id)
}

// GetWorkContract is a free data retrieval call binding the contract method 0x4d2522a2.
//
// Solidity: function getWorkContract(bytes32 _id) view returns(address)
func (_Artefacts *ArtefactsCaller) GetWorkContract(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getWorkContract", _id)
	return *ret0, err
}

// GetWorkContract is a free data retrieval call binding the contract method 0x4d2522a2.
//
// Solidity: function getWorkContract(bytes32 _id) view returns(address)
func (_Artefacts *ArtefactsSession) GetWorkContract(_id [32]byte) (common.Address, error) {
	return _Artefacts.Contract.GetWorkContract(&_Artefacts.CallOpts, _id)
}

// GetWorkContract is a free data retrieval call binding the contract method 0x4d2522a2.
//
// Solidity: function getWorkContract(bytes32 _id) view returns(address)
func (_Artefacts *ArtefactsCallerSession) GetWorkContract(_id [32]byte) (common.Address, error) {
	return _Artefacts.Contract.GetWorkContract(&_Artefacts.CallOpts, _id)
}

// GetWorkCopyrightHolders is a free data retrieval call binding the contract method 0x86105c65.
//
// Solidity: function getWorkCopyrightHolders(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsCaller) GetWorkCopyrightHolders(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getWorkCopyrightHolders", _id)
	return *ret0, err
}

// GetWorkCopyrightHolders is a free data retrieval call binding the contract method 0x86105c65.
//
// Solidity: function getWorkCopyrightHolders(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsSession) GetWorkCopyrightHolders(_id [32]byte) ([][32]byte, error) {
	return _Artefacts.Contract.GetWorkCopyrightHolders(&_Artefacts.CallOpts, _id)
}

// GetWorkCopyrightHolders is a free data retrieval call binding the contract method 0x86105c65.
//
// Solidity: function getWorkCopyrightHolders(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsCallerSession) GetWorkCopyrightHolders(_id [32]byte) ([][32]byte, error) {
	return _Artefacts.Contract.GetWorkCopyrightHolders(&_Artefacts.CallOpts, _id)
}

// GetWorkPublishers is a free data retrieval call binding the contract method 0x5ffd09d8.
//
// Solidity: function getWorkPublishers(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsCaller) GetWorkPublishers(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getWorkPublishers", _id)
	return *ret0, err
}

// GetWorkPublishers is a free data retrieval call binding the contract method 0x5ffd09d8.
//
// Solidity: function getWorkPublishers(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsSession) GetWorkPublishers(_id [32]byte) ([][32]byte, error) {
	return _Artefacts.Contract.GetWorkPublishers(&_Artefacts.CallOpts, _id)
}

// GetWorkPublishers is a free data retrieval call binding the contract method 0x5ffd09d8.
//
// Solidity: function getWorkPublishers(bytes32 _id) view returns(bytes32[])
func (_Artefacts *ArtefactsCallerSession) GetWorkPublishers(_id [32]byte) ([][32]byte, error) {
	return _Artefacts.Contract.GetWorkPublishers(&_Artefacts.CallOpts, _id)
}

// IsWorkAuthor is a free data retrieval call binding the contract method 0xb2e459cf.
//
// Solidity: function isWorkAuthor(bytes32 _workId, bytes32 _authorId) view returns(bool)
func (_Artefacts *ArtefactsCaller) IsWorkAuthor(opts *bind.CallOpts, _workId [32]byte, _authorId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "isWorkAuthor", _workId, _authorId)
	return *ret0, err
}

// IsWorkAuthor is a free data retrieval call binding the contract method 0xb2e459cf.
//
// Solidity: function isWorkAuthor(bytes32 _workId, bytes32 _authorId) view returns(bool)
func (_Artefacts *ArtefactsSession) IsWorkAuthor(_workId [32]byte, _authorId [32]byte) (bool, error) {
	return _Artefacts.Contract.IsWorkAuthor(&_Artefacts.CallOpts, _workId, _authorId)
}

// IsWorkAuthor is a free data retrieval call binding the contract method 0xb2e459cf.
//
// Solidity: function isWorkAuthor(bytes32 _workId, bytes32 _authorId) view returns(bool)
func (_Artefacts *ArtefactsCallerSession) IsWorkAuthor(_workId [32]byte, _authorId [32]byte) (bool, error) {
	return _Artefacts.Contract.IsWorkAuthor(&_Artefacts.CallOpts, _workId, _authorId)
}

// IsWorkCopyrightHolder is a free data retrieval call binding the contract method 0xf6a50965.
//
// Solidity: function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) view returns(bool)
func (_Artefacts *ArtefactsCaller) IsWorkCopyrightHolder(opts *bind.CallOpts, _workId [32]byte, _copyrightHolderId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "isWorkCopyrightHolder", _workId, _copyrightHolderId)
	return *ret0, err
}

// IsWorkCopyrightHolder is a free data retrieval call binding the contract method 0xf6a50965.
//
// Solidity: function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) view returns(bool)
func (_Artefacts *ArtefactsSession) IsWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (bool, error) {
	return _Artefacts.Contract.IsWorkCopyrightHolder(&_Artefacts.CallOpts, _workId, _copyrightHolderId)
}

// IsWorkCopyrightHolder is a free data retrieval call binding the contract method 0xf6a50965.
//
// Solidity: function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) view returns(bool)
func (_Artefacts *ArtefactsCallerSession) IsWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (bool, error) {
	return _Artefacts.Contract.IsWorkCopyrightHolder(&_Artefacts.CallOpts, _workId, _copyrightHolderId)
}

// IsWorkPublisher is a free data retrieval call binding the contract method 0xef90d9c9.
//
// Solidity: function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) view returns(bool)
func (_Artefacts *ArtefactsCaller) IsWorkPublisher(opts *bind.CallOpts, _workId [32]byte, _publisherId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "isWorkPublisher", _workId, _publisherId)
	return *ret0, err
}

// IsWorkPublisher is a free data retrieval call binding the contract method 0xef90d9c9.
//
// Solidity: function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) view returns(bool)
func (_Artefacts *ArtefactsSession) IsWorkPublisher(_workId [32]byte, _publisherId [32]byte) (bool, error) {
	return _Artefacts.Contract.IsWorkPublisher(&_Artefacts.CallOpts, _workId, _publisherId)
}

// IsWorkPublisher is a free data retrieval call binding the contract method 0xef90d9c9.
//
// Solidity: function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) view returns(bool)
func (_Artefacts *ArtefactsCallerSession) IsWorkPublisher(_workId [32]byte, _publisherId [32]byte) (bool, error) {
	return _Artefacts.Contract.IsWorkPublisher(&_Artefacts.CallOpts, _workId, _publisherId)
}

// AddWork is a paid mutator transaction binding the contract method 0x1d339e24.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_Artefacts *ArtefactsTransactor) AddWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWork", _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x1d339e24.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_Artefacts *ArtefactsSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWork(&_Artefacts.TransactOpts, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x1d339e24.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWork(&_Artefacts.TransactOpts, _work)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x55884f0f.
//
// Solidity: function addWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_Artefacts *ArtefactsTransactor) AddWorkAuthor(opts *bind.TransactOpts, _workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWorkAuthor", _workId, _authorId)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x55884f0f.
//
// Solidity: function addWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_Artefacts *ArtefactsSession) AddWorkAuthor(_workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkAuthor(&_Artefacts.TransactOpts, _workId, _authorId)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x55884f0f.
//
// Solidity: function addWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWorkAuthor(_workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkAuthor(&_Artefacts.TransactOpts, _workId, _authorId)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x0f1ee4a2.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_Artefacts *ArtefactsTransactor) AddWorkCopyrightHolder(opts *bind.TransactOpts, _workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWorkCopyrightHolder", _workId, _copyrightHolderId)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x0f1ee4a2.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_Artefacts *ArtefactsSession) AddWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkCopyrightHolder(&_Artefacts.TransactOpts, _workId, _copyrightHolderId)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x0f1ee4a2.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkCopyrightHolder(&_Artefacts.TransactOpts, _workId, _copyrightHolderId)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0xe17989ea.
//
// Solidity: function addWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_Artefacts *ArtefactsTransactor) AddWorkPublisher(opts *bind.TransactOpts, _workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWorkPublisher", _workId, _publisherId)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0xe17989ea.
//
// Solidity: function addWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_Artefacts *ArtefactsSession) AddWorkPublisher(_workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkPublisher(&_Artefacts.TransactOpts, _workId, _publisherId)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0xe17989ea.
//
// Solidity: function addWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWorkPublisher(_workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkPublisher(&_Artefacts.TransactOpts, _workId, _publisherId)
}

// AmendWork is a paid mutator transaction binding the contract method 0x4e542a6e.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_Artefacts *ArtefactsTransactor) AmendWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "amendWork", _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0x4e542a6e.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_Artefacts *ArtefactsSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AmendWork(&_Artefacts.TransactOpts, _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0x4e542a6e.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_Artefacts *ArtefactsTransactorSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AmendWork(&_Artefacts.TransactOpts, _work)
}

// RemoveWorkAuthor is a paid mutator transaction binding the contract method 0xe229e173.
//
// Solidity: function removeWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_Artefacts *ArtefactsTransactor) RemoveWorkAuthor(opts *bind.TransactOpts, _workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "removeWorkAuthor", _workId, _authorId)
}

// RemoveWorkAuthor is a paid mutator transaction binding the contract method 0xe229e173.
//
// Solidity: function removeWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_Artefacts *ArtefactsSession) RemoveWorkAuthor(_workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.RemoveWorkAuthor(&_Artefacts.TransactOpts, _workId, _authorId)
}

// RemoveWorkAuthor is a paid mutator transaction binding the contract method 0xe229e173.
//
// Solidity: function removeWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_Artefacts *ArtefactsTransactorSession) RemoveWorkAuthor(_workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.RemoveWorkAuthor(&_Artefacts.TransactOpts, _workId, _authorId)
}

// RemoveWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x521ffdce.
//
// Solidity: function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_Artefacts *ArtefactsTransactor) RemoveWorkCopyrightHolder(opts *bind.TransactOpts, _workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "removeWorkCopyrightHolder", _workId, _copyrightHolderId)
}

// RemoveWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x521ffdce.
//
// Solidity: function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_Artefacts *ArtefactsSession) RemoveWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.RemoveWorkCopyrightHolder(&_Artefacts.TransactOpts, _workId, _copyrightHolderId)
}

// RemoveWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x521ffdce.
//
// Solidity: function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_Artefacts *ArtefactsTransactorSession) RemoveWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.RemoveWorkCopyrightHolder(&_Artefacts.TransactOpts, _workId, _copyrightHolderId)
}

// RemoveWorkPublisher is a paid mutator transaction binding the contract method 0x0c99420f.
//
// Solidity: function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_Artefacts *ArtefactsTransactor) RemoveWorkPublisher(opts *bind.TransactOpts, _workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "removeWorkPublisher", _workId, _publisherId)
}

// RemoveWorkPublisher is a paid mutator transaction binding the contract method 0x0c99420f.
//
// Solidity: function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_Artefacts *ArtefactsSession) RemoveWorkPublisher(_workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.RemoveWorkPublisher(&_Artefacts.TransactOpts, _workId, _publisherId)
}

// RemoveWorkPublisher is a paid mutator transaction binding the contract method 0x0c99420f.
//
// Solidity: function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_Artefacts *ArtefactsTransactorSession) RemoveWorkPublisher(_workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _Artefacts.Contract.RemoveWorkPublisher(&_Artefacts.TransactOpts, _workId, _publisherId)
}

// EntitiesABI is the input ABI used to generate the binding from.
const EntitiesABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"name\":\"addEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amendEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntity\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReferences\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isEntityType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntitiesFuncSigs maps the 4-byte function signature to its string representation.
var EntitiesFuncSigs = map[string]string{
	"68a9f4e9": "addEntity((bytes32,string,string,string),uint8)",
	"a81eb545": "amendEntity((bytes32,string,string,string),uint8)",
	"53b66f36": "getEntity(bytes32)",
	"d91ee2e0": "getEntityContract(bytes32)",
	"6ec21da9": "getEntityTypes(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"7a6337fa": "getReferences()",
	"cef8b362": "isEntityType(bytes32,uint8)",
}

// EntitiesBin is the compiled bytecode used for deploying new contracts.
var EntitiesBin = "0x608060405234801561001057600080fd5b506000600255611a1c806100256000396000f3fe60806040523480156200001157600080fd5b5060043610620000a05760003560e01c80637a6337fa116200006f5780637a6337fa146200012c578063a81eb5451462000145578063bc599341146200015c578063cef8b3621462000173578063d91ee2e0146200019957620000a0565b806353b66f3614620000a557806367e0badb14620000d457806368a9f4e914620000ed5780636ec21da91462000106575b600080fd5b620000bc620000b6366004620008f1565b620001bf565b604051620000cb919062000c2b565b60405180910390f35b620000de62000287565b604051620000cb919062000c0b565b62000104620000fe36600462000a04565b6200028d565b005b6200011d62000117366004620008f1565b62000319565b604051620000cb919062000b7e565b62000136620003d1565b604051620000cb919062000bc6565b620001046200015636600462000a04565b62000475565b620000de6200016d366004620008f1565b6200051d565b6200018a620001843660046200090a565b62000557565b604051620000cb919062000c00565b620001b0620001aa366004620008f1565b6200061e565b604051620000cb919062000b6a565b620001c962000712565b6000828152602081905260409020600101546001600160a01b0316620001ee57600080fd5b600082815260208190526040808220600101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b1580156200023f57600080fd5b505afa15801562000254573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526200027e919081019062000939565b9150505b919050565b60025490565b8151620002a39060009063ffffffff6200066216565b15620002bb57620002b5828262000475565b62000315565b60008282604051620002cd906200073d565b620002da92919062000c40565b604051809103906000f080158015620002f7573d6000803e3d6000fd5b50835190915062000312906000908363ffffffff6200067716565b50505b5050565b6000818152602081905260409020600101546060906001600160a01b03166200034157600080fd5b6000828152602081905260408082206001015481516305a2bceb60e51b815291516001600160a01b0390911692839263b4579d609260048083019392829003018186803b1580156200039257600080fd5b505afa158015620003a7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526200027e919081019062000819565b60608060006002015467ffffffffffffffff81118015620003f157600080fd5b506040519080825280602002602001820160405280156200041c578160200160208202803683370190505b50905060005b6002548110156200046f5760018054829081106200043c57fe5b9060005260206000209060020201600001548282815181106200045b57fe5b602090810291909101015260010162000422565b50905090565b81516000908152602081905260409020600101546001600160a01b03166200049c57600080fd5b81516000908152602081905260409081902060010154905163d1f439e560e01b81526001600160a01b0390911690819063d1f439e590620004e4908690869060040162000c40565b600060405180830381600087803b158015620004ff57600080fd5b505af115801562000514573d6000803e3d6000fd5b50505050505050565b60025460009082106200052f57600080fd5b60018054839081106200053e57fe5b9060005260206000209060020201600001549050919050565b6000828152602081905260408120600101546001600160a01b03166200057c57600080fd5b6000838152602081905260409081902060010154905163222f74e760e01b81526001600160a01b0390911690819063222f74e790620005c090869060040162000c14565b60206040518083038186803b158015620005d957600080fd5b505afa158015620005ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620006149190620008cf565b9150505b92915050565b6000818152602081905260408120600101546001600160a01b03166200064357600080fd5b506000908152602081905260409020600101546001600160a01b031690565b60009081526020919091526040902054151590565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b0385161790558015620006b75760019150506200070b565b5060018085018054918201808255600086815260208890526040902081905585919083908110620006e457fe5b600091825260208220600291820201929092559086018054600101905591506200070b9050565b9392505050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b610cf08062000cf783390190565b805180151581146200061857600080fd5b8035600581106200061857600080fd5b600082601f8301126200077d578081fd5b8135620007946200078e8262000c96565b62000c6e565b9150808252836020828501011115620007ac57600080fd5b8060208401602084013760009082016020015292915050565b600082601f830112620007d6578081fd5b8151620007e76200078e8262000c96565b9150808252836020828501011115620007ff57600080fd5b6200081281602084016020860162000cc7565b5092915050565b600060208083850312156200082c578182fd5b825167ffffffffffffffff8082111562000844578384fd5b81850186601f82011262000856578485fd5b805192508183111562000867578485fd5b83830291506200087984830162000c6e565b8381528481019082860184840187018a101562000894578788fd5b8794505b85851015620008c257620008ad8a826200074b565b83526001949094019391860191860162000898565b5098975050505050505050565b600060208284031215620008e1578081fd5b815180151581146200070b578182fd5b60006020828403121562000903578081fd5b5035919050565b600080604083850312156200091d578081fd5b823591506200093084602085016200075c565b90509250929050565b6000602082840312156200094b578081fd5b815167ffffffffffffffff8082111562000963578283fd5b8184016080818703121562000976578384fd5b62000982608062000c6e565b9250805183526020810151828111156200099a578485fd5b620009a887828401620007c5565b602085015250604081015182811115620009c0578485fd5b620009ce87828401620007c5565b604085015250606081015182811115620009e6578485fd5b620009f487828401620007c5565b6060850152509195945050505050565b6000806040838503121562000a17578182fd5b823567ffffffffffffffff8082111562000a2f578384fd5b8185016080818803121562000a42578485fd5b62000a4e608062000c6e565b92508035835260208101358281111562000a66578586fd5b62000a74888284016200076c565b60208501525060408101358281111562000a8c578586fd5b62000a9a888284016200076c565b60408501525060608101358281111562000ab2578586fd5b62000ac0888284016200076c565b6060850152505050809250506200093084602085016200075c565b6000815180845262000af581602086016020860162000cc7565b601f01601f19169290920160200192915050565b60008151835260208201516080602085015262000b2a608085018262000adb565b60408401519150848103604086015262000b45818362000adb565b60608501519250858103606087015262000b60818462000adb565b9695505050505050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b8181101562000bba57835115158352928401929184019160010162000b9a565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101562000bba5783518352928401929184019160010162000be2565b901515815260200190565b90815260200190565b6020810162000c238362000cbb565b825292915050565b6000602082526200070b602083018462000b09565b60006040825262000c55604083018562000b09565b905062000c628362000cbb565b60208301529392505050565b60405181810167ffffffffffffffff8111828210171562000c8e57600080fd5b604052919050565b600067ffffffffffffffff82111562000cad578081fd5b50601f01601f191660200190565b80600581106200028257fe5b60005b8381101562000ce457818101518382015260200162000cca565b8381111562000312575050600091015256fe60806040523480156200001157600080fd5b5060405162000cf038038062000cf08339810160408190526200003491620003b7565b815160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007257506000826020015151115b80156200008457506000826040015151115b80156200009d575060008160048111156200009b57fe5b115b8015620000b657506004816004811115620000b457fe5b105b620000c057600080fd5b60408051600480825260a08201909252906020820160808036833750508151620000f2926004925060200190620001a4565b50600160048260048111156200010457fe5b60ff16815481106200011257fe5b6000918252602080832081830401805460ff601f9094166101000a9384021916941515929092029390931790558351815583820151805185936200015c9260019291019062000250565b50604082015180516200017a91600284019160209091019062000250565b50606082015180516200019891600384019160209091019062000250565b509050505050620004bd565b82805482825590600052602060002090601f016020900481019282156200023e5791602002820160005b838211156200020d57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302620001ce565b80156200023c5782816101000a81549060ff02191690556001016020816000010492830192600103026200020d565b505b506200024c929150620002d1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200029357805160ff1916838001178555620002c3565b82800160010185558215620002c3579182015b82811115620002c3578251825591602001919060010190620002a6565b506200024c929150620002f5565b620002f291905b808211156200024c57805460ff19168155600101620002d8565b90565b620002f291905b808211156200024c5760008155600101620002fc565b8051600581106200032257600080fd5b92915050565b600082601f83011262000339578081fd5b81516001600160401b038111156200034f578182fd5b602062000365601f8301601f1916820162000496565b925081835284818386010111156200037c57600080fd5b60005b828110156200039c5784810182015184820183015281016200037f565b82811115620003ae5760008284860101525b50505092915050565b60008060408385031215620003ca578182fd5b82516001600160401b0380821115620003e1578384fd5b81850160808188031215620003f4578485fd5b62000400608062000496565b92508051835260208101518281111562000418578586fd5b620004268882840162000328565b6020850152506040810151828111156200043e578586fd5b6200044c8882840162000328565b60408501525060608101518281111562000464578586fd5b620004728882840162000328565b6060850152505050809250506200048d846020850162000312565b90509250929050565b6040518181016001600160401b0381118282101715620004b557600080fd5b604052919050565b61082380620004cd6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063222f74e71461005c5780636d4ce63c14610085578063b4579d601461009a578063d1f439e5146100af578063e188130b146100c4575b600080fd5b61006f61006a3660046105ca565b6100d7565b60405161007c9190610751565b60405180910390f35b61008d610149565b60405161007c919061075c565b6100a2610322565b60405161007c919061070b565b6100c26100bd3660046105f0565b61039a565b005b6100c26100d23660046105ca565b610403565b6000808260048111156100e657fe5b1180156100fe575060048260048111156100fc57fe5b105b61010757600080fd5b600482600481111561011557fe5b60ff168154811061012257fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b610151610489565b604080516080810182526000805482526001805484516020600283851615610100026000190190931692909204601f810183900483028201830190965285815293949293818601939092918301828280156101ed5780601f106101c2576101008083540402835291602001916101ed565b820191906000526020600020905b8154815290600101906020018083116101d057829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561027f5780601f106102545761010080835404028352916020019161027f565b820191906000526020600020905b81548152906001019060200180831161026257829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103135780601f106102e857610100808354040283529160200191610313565b820191906000526020600020905b8154815290600101906020018083116102f657829003601f168201915b50505050508152505090505b90565b6060600480548060200260200160405190810160405280929190818152602001828054801561039057602002820191906000526020600020906000905b825461010083900a900460ff16151581526020600192830181810494850194909303909202910180841161035f5790505b5050505050905090565b6103a381610403565b8151600090815560208084015180518593926103c4926001929101906104b4565b50604082015180516103e09160028401916020909101906104b4565b50606082015180516103fc9160038401916020909101906104b4565b5050505050565b600081600481111561041157fe5b1180156104295750600481600481111561042757fe5b105b61043257600080fd5b61043b816100d7565b610486576001600482600481111561044f57fe5b60ff168154811061045c57fe5b90600052602060002090602091828204019190066101000a81548160ff0219169083151502179055505b50565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104f557805160ff1916838001178555610522565b82800160010185558215610522579182015b82811115610522578251825591602001919060010190610507565b5061052e929150610532565b5090565b61031f91905b8082111561052e5760008155600101610538565b80356005811061055b57600080fd5b92915050565b600082601f830112610571578081fd5b813567ffffffffffffffff811115610587578182fd5b61059a601f8201601f19166020016107c6565b91508082528360208285010111156105b157600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156105db578081fd5b8135600581106105e9578182fd5b9392505050565b60008060408385031215610602578081fd5b823567ffffffffffffffff80821115610619578283fd5b8185016080818803121561062b578384fd5b61063560806107c6565b92508035835260208101358281111561064c578485fd5b61065888828401610561565b60208501525060408101358281111561066f578485fd5b61067b88828401610561565b604085015250606081013582811115610692578485fd5b61069e88828401610561565b6060850152505050809250506106b7846020850161054c565b90509250929050565b60008151808452815b818110156106e5576020818501810151868301820152016106c9565b818111156106f65782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610745578351151583529284019291840191600101610727565b50909695505050505050565b901515815260200190565b6000602082528251602083015260208301516080604084015261078260a08401826106c0565b60408501519150601f19808583030160608601526107a082846106c0565b60608701519350818682030160808701526107bb81856106c0565b979650505050505050565b60405181810167ffffffffffffffff811182821017156107e557600080fd5b60405291905056fea2646970667358221220566b1060ffdb4ae5daac633bc59b0a6b7423992070f22fc9ccc75e02d21c44f964736f6c63430006080033a2646970667358221220c1f2bc3450f7da35d5ee783fa723d9ac450d2f2bca6ad469c871b76684dfae9a64736f6c63430006080033"

// DeployEntities deploys a new Ethereum contract, binding an instance of Entities to it.
func DeployEntities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Entities, error) {
	parsed, err := abi.JSON(strings.NewReader(EntitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EntitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Entities{EntitiesCaller: EntitiesCaller{contract: contract}, EntitiesTransactor: EntitiesTransactor{contract: contract}, EntitiesFilterer: EntitiesFilterer{contract: contract}}, nil
}

// Entities is an auto generated Go binding around an Ethereum contract.
type Entities struct {
	EntitiesCaller     // Read-only binding to the contract
	EntitiesTransactor // Write-only binding to the contract
	EntitiesFilterer   // Log filterer for contract events
}

// EntitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntitiesSession struct {
	Contract     *Entities         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntitiesCallerSession struct {
	Contract *EntitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EntitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntitiesTransactorSession struct {
	Contract     *EntitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EntitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntitiesRaw struct {
	Contract *Entities // Generic contract binding to access the raw methods on
}

// EntitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntitiesCallerRaw struct {
	Contract *EntitiesCaller // Generic read-only contract binding to access the raw methods on
}

// EntitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntitiesTransactorRaw struct {
	Contract *EntitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntities creates a new instance of Entities, bound to a specific deployed contract.
func NewEntities(address common.Address, backend bind.ContractBackend) (*Entities, error) {
	contract, err := bindEntities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Entities{EntitiesCaller: EntitiesCaller{contract: contract}, EntitiesTransactor: EntitiesTransactor{contract: contract}, EntitiesFilterer: EntitiesFilterer{contract: contract}}, nil
}

// NewEntitiesCaller creates a new read-only instance of Entities, bound to a specific deployed contract.
func NewEntitiesCaller(address common.Address, caller bind.ContractCaller) (*EntitiesCaller, error) {
	contract, err := bindEntities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntitiesCaller{contract: contract}, nil
}

// NewEntitiesTransactor creates a new write-only instance of Entities, bound to a specific deployed contract.
func NewEntitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*EntitiesTransactor, error) {
	contract, err := bindEntities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntitiesTransactor{contract: contract}, nil
}

// NewEntitiesFilterer creates a new log filterer instance of Entities, bound to a specific deployed contract.
func NewEntitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*EntitiesFilterer, error) {
	contract, err := bindEntities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntitiesFilterer{contract: contract}, nil
}

// bindEntities binds a generic wrapper to an already deployed contract.
func bindEntities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EntitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entities *EntitiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Entities.Contract.EntitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entities *EntitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entities.Contract.EntitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entities *EntitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entities.Contract.EntitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entities *EntitiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Entities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entities *EntitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entities *EntitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entities.Contract.contract.Transact(opts, method, params...)
}

// GetEntity is a free data retrieval call binding the contract method 0x53b66f36.
//
// Solidity: function getEntity(bytes32 _id) view returns((bytes32,string,string,string))
func (_Entities *EntitiesCaller) GetEntity(opts *bind.CallOpts, _id [32]byte) (CreativeEntities, error) {
	var (
		ret0 = new(CreativeEntities)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getEntity", _id)
	return *ret0, err
}

// GetEntity is a free data retrieval call binding the contract method 0x53b66f36.
//
// Solidity: function getEntity(bytes32 _id) view returns((bytes32,string,string,string))
func (_Entities *EntitiesSession) GetEntity(_id [32]byte) (CreativeEntities, error) {
	return _Entities.Contract.GetEntity(&_Entities.CallOpts, _id)
}

// GetEntity is a free data retrieval call binding the contract method 0x53b66f36.
//
// Solidity: function getEntity(bytes32 _id) view returns((bytes32,string,string,string))
func (_Entities *EntitiesCallerSession) GetEntity(_id [32]byte) (CreativeEntities, error) {
	return _Entities.Contract.GetEntity(&_Entities.CallOpts, _id)
}

// GetEntityContract is a free data retrieval call binding the contract method 0xd91ee2e0.
//
// Solidity: function getEntityContract(bytes32 _id) view returns(address)
func (_Entities *EntitiesCaller) GetEntityContract(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getEntityContract", _id)
	return *ret0, err
}

// GetEntityContract is a free data retrieval call binding the contract method 0xd91ee2e0.
//
// Solidity: function getEntityContract(bytes32 _id) view returns(address)
func (_Entities *EntitiesSession) GetEntityContract(_id [32]byte) (common.Address, error) {
	return _Entities.Contract.GetEntityContract(&_Entities.CallOpts, _id)
}

// GetEntityContract is a free data retrieval call binding the contract method 0xd91ee2e0.
//
// Solidity: function getEntityContract(bytes32 _id) view returns(address)
func (_Entities *EntitiesCallerSession) GetEntityContract(_id [32]byte) (common.Address, error) {
	return _Entities.Contract.GetEntityContract(&_Entities.CallOpts, _id)
}

// GetEntityTypes is a free data retrieval call binding the contract method 0x6ec21da9.
//
// Solidity: function getEntityTypes(bytes32 _id) view returns(bool[])
func (_Entities *EntitiesCaller) GetEntityTypes(opts *bind.CallOpts, _id [32]byte) ([]bool, error) {
	var (
		ret0 = new([]bool)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getEntityTypes", _id)
	return *ret0, err
}

// GetEntityTypes is a free data retrieval call binding the contract method 0x6ec21da9.
//
// Solidity: function getEntityTypes(bytes32 _id) view returns(bool[])
func (_Entities *EntitiesSession) GetEntityTypes(_id [32]byte) ([]bool, error) {
	return _Entities.Contract.GetEntityTypes(&_Entities.CallOpts, _id)
}

// GetEntityTypes is a free data retrieval call binding the contract method 0x6ec21da9.
//
// Solidity: function getEntityTypes(bytes32 _id) view returns(bool[])
func (_Entities *EntitiesCallerSession) GetEntityTypes(_id [32]byte) ([]bool, error) {
	return _Entities.Contract.GetEntityTypes(&_Entities.CallOpts, _id)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Entities *EntitiesCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Entities *EntitiesSession) GetNum() (*big.Int, error) {
	return _Entities.Contract.GetNum(&_Entities.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Entities *EntitiesCallerSession) GetNum() (*big.Int, error) {
	return _Entities.Contract.GetNum(&_Entities.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Entities *EntitiesCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Entities *EntitiesSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Entities.Contract.GetReference(&_Entities.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Entities *EntitiesCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Entities.Contract.GetReference(&_Entities.CallOpts, _index)
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_Entities *EntitiesCaller) GetReferences(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getReferences")
	return *ret0, err
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_Entities *EntitiesSession) GetReferences() ([][32]byte, error) {
	return _Entities.Contract.GetReferences(&_Entities.CallOpts)
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_Entities *EntitiesCallerSession) GetReferences() ([][32]byte, error) {
	return _Entities.Contract.GetReferences(&_Entities.CallOpts)
}

// IsEntityType is a free data retrieval call binding the contract method 0xcef8b362.
//
// Solidity: function isEntityType(bytes32 _id, uint8 _type) view returns(bool)
func (_Entities *EntitiesCaller) IsEntityType(opts *bind.CallOpts, _id [32]byte, _type uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "isEntityType", _id, _type)
	return *ret0, err
}

// IsEntityType is a free data retrieval call binding the contract method 0xcef8b362.
//
// Solidity: function isEntityType(bytes32 _id, uint8 _type) view returns(bool)
func (_Entities *EntitiesSession) IsEntityType(_id [32]byte, _type uint8) (bool, error) {
	return _Entities.Contract.IsEntityType(&_Entities.CallOpts, _id, _type)
}

// IsEntityType is a free data retrieval call binding the contract method 0xcef8b362.
//
// Solidity: function isEntityType(bytes32 _id, uint8 _type) view returns(bool)
func (_Entities *EntitiesCallerSession) IsEntityType(_id [32]byte, _type uint8) (bool, error) {
	return _Entities.Contract.IsEntityType(&_Entities.CallOpts, _id, _type)
}

// AddEntity is a paid mutator transaction binding the contract method 0x68a9f4e9.
//
// Solidity: function addEntity((bytes32,string,string,string) _entity, uint8 _entityType) returns()
func (_Entities *EntitiesTransactor) AddEntity(opts *bind.TransactOpts, _entity CreativeEntities, _entityType uint8) (*types.Transaction, error) {
	return _Entities.contract.Transact(opts, "addEntity", _entity, _entityType)
}

// AddEntity is a paid mutator transaction binding the contract method 0x68a9f4e9.
//
// Solidity: function addEntity((bytes32,string,string,string) _entity, uint8 _entityType) returns()
func (_Entities *EntitiesSession) AddEntity(_entity CreativeEntities, _entityType uint8) (*types.Transaction, error) {
	return _Entities.Contract.AddEntity(&_Entities.TransactOpts, _entity, _entityType)
}

// AddEntity is a paid mutator transaction binding the contract method 0x68a9f4e9.
//
// Solidity: function addEntity((bytes32,string,string,string) _entity, uint8 _entityType) returns()
func (_Entities *EntitiesTransactorSession) AddEntity(_entity CreativeEntities, _entityType uint8) (*types.Transaction, error) {
	return _Entities.Contract.AddEntity(&_Entities.TransactOpts, _entity, _entityType)
}

// AmendEntity is a paid mutator transaction binding the contract method 0xa81eb545.
//
// Solidity: function amendEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_Entities *EntitiesTransactor) AmendEntity(opts *bind.TransactOpts, _entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _Entities.contract.Transact(opts, "amendEntity", _entity, _type)
}

// AmendEntity is a paid mutator transaction binding the contract method 0xa81eb545.
//
// Solidity: function amendEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_Entities *EntitiesSession) AmendEntity(_entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _Entities.Contract.AmendEntity(&_Entities.TransactOpts, _entity, _type)
}

// AmendEntity is a paid mutator transaction binding the contract method 0xa81eb545.
//
// Solidity: function amendEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_Entities *EntitiesTransactorSession) AmendEntity(_entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _Entities.Contract.AmendEntity(&_Entities.TransactOpts, _entity, _type)
}

// EntityNodeABI is the input ABI used to generate the binding from.
const EntityNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntityNodeFuncSigs maps the 4-byte function signature to its string representation.
var EntityNodeFuncSigs = map[string]string{
	"e188130b": "addType(uint8)",
	"d1f439e5": "amend((bytes32,string,string,string),uint8)",
	"6d4ce63c": "get()",
	"b4579d60": "getTypes()",
	"222f74e7": "isType(uint8)",
}

// EntityNodeBin is the compiled bytecode used for deploying new contracts.
var EntityNodeBin = "0x60806040523480156200001157600080fd5b5060405162000cf038038062000cf08339810160408190526200003491620003b7565b815160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007257506000826020015151115b80156200008457506000826040015151115b80156200009d575060008160048111156200009b57fe5b115b8015620000b657506004816004811115620000b457fe5b105b620000c057600080fd5b60408051600480825260a08201909252906020820160808036833750508151620000f2926004925060200190620001a4565b50600160048260048111156200010457fe5b60ff16815481106200011257fe5b6000918252602080832081830401805460ff601f9094166101000a9384021916941515929092029390931790558351815583820151805185936200015c9260019291019062000250565b50604082015180516200017a91600284019160209091019062000250565b50606082015180516200019891600384019160209091019062000250565b509050505050620004bd565b82805482825590600052602060002090601f016020900481019282156200023e5791602002820160005b838211156200020d57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302620001ce565b80156200023c5782816101000a81549060ff02191690556001016020816000010492830192600103026200020d565b505b506200024c929150620002d1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200029357805160ff1916838001178555620002c3565b82800160010185558215620002c3579182015b82811115620002c3578251825591602001919060010190620002a6565b506200024c929150620002f5565b620002f291905b808211156200024c57805460ff19168155600101620002d8565b90565b620002f291905b808211156200024c5760008155600101620002fc565b8051600581106200032257600080fd5b92915050565b600082601f83011262000339578081fd5b81516001600160401b038111156200034f578182fd5b602062000365601f8301601f1916820162000496565b925081835284818386010111156200037c57600080fd5b60005b828110156200039c5784810182015184820183015281016200037f565b82811115620003ae5760008284860101525b50505092915050565b60008060408385031215620003ca578182fd5b82516001600160401b0380821115620003e1578384fd5b81850160808188031215620003f4578485fd5b62000400608062000496565b92508051835260208101518281111562000418578586fd5b620004268882840162000328565b6020850152506040810151828111156200043e578586fd5b6200044c8882840162000328565b60408501525060608101518281111562000464578586fd5b620004728882840162000328565b6060850152505050809250506200048d846020850162000312565b90509250929050565b6040518181016001600160401b0381118282101715620004b557600080fd5b604052919050565b61082380620004cd6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063222f74e71461005c5780636d4ce63c14610085578063b4579d601461009a578063d1f439e5146100af578063e188130b146100c4575b600080fd5b61006f61006a3660046105ca565b6100d7565b60405161007c9190610751565b60405180910390f35b61008d610149565b60405161007c919061075c565b6100a2610322565b60405161007c919061070b565b6100c26100bd3660046105f0565b61039a565b005b6100c26100d23660046105ca565b610403565b6000808260048111156100e657fe5b1180156100fe575060048260048111156100fc57fe5b105b61010757600080fd5b600482600481111561011557fe5b60ff168154811061012257fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b610151610489565b604080516080810182526000805482526001805484516020600283851615610100026000190190931692909204601f810183900483028201830190965285815293949293818601939092918301828280156101ed5780601f106101c2576101008083540402835291602001916101ed565b820191906000526020600020905b8154815290600101906020018083116101d057829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561027f5780601f106102545761010080835404028352916020019161027f565b820191906000526020600020905b81548152906001019060200180831161026257829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103135780601f106102e857610100808354040283529160200191610313565b820191906000526020600020905b8154815290600101906020018083116102f657829003601f168201915b50505050508152505090505b90565b6060600480548060200260200160405190810160405280929190818152602001828054801561039057602002820191906000526020600020906000905b825461010083900a900460ff16151581526020600192830181810494850194909303909202910180841161035f5790505b5050505050905090565b6103a381610403565b8151600090815560208084015180518593926103c4926001929101906104b4565b50604082015180516103e09160028401916020909101906104b4565b50606082015180516103fc9160038401916020909101906104b4565b5050505050565b600081600481111561041157fe5b1180156104295750600481600481111561042757fe5b105b61043257600080fd5b61043b816100d7565b610486576001600482600481111561044f57fe5b60ff168154811061045c57fe5b90600052602060002090602091828204019190066101000a81548160ff0219169083151502179055505b50565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104f557805160ff1916838001178555610522565b82800160010185558215610522579182015b82811115610522578251825591602001919060010190610507565b5061052e929150610532565b5090565b61031f91905b8082111561052e5760008155600101610538565b80356005811061055b57600080fd5b92915050565b600082601f830112610571578081fd5b813567ffffffffffffffff811115610587578182fd5b61059a601f8201601f19166020016107c6565b91508082528360208285010111156105b157600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156105db578081fd5b8135600581106105e9578182fd5b9392505050565b60008060408385031215610602578081fd5b823567ffffffffffffffff80821115610619578283fd5b8185016080818803121561062b578384fd5b61063560806107c6565b92508035835260208101358281111561064c578485fd5b61065888828401610561565b60208501525060408101358281111561066f578485fd5b61067b88828401610561565b604085015250606081013582811115610692578485fd5b61069e88828401610561565b6060850152505050809250506106b7846020850161054c565b90509250929050565b60008151808452815b818110156106e5576020818501810151868301820152016106c9565b818111156106f65782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610745578351151583529284019291840191600101610727565b50909695505050505050565b901515815260200190565b6000602082528251602083015260208301516080604084015261078260a08401826106c0565b60408501519150601f19808583030160608601526107a082846106c0565b60608701519350818682030160808701526107bb81856106c0565b979650505050505050565b60405181810167ffffffffffffffff811182821017156107e557600080fd5b60405291905056fea2646970667358221220566b1060ffdb4ae5daac633bc59b0a6b7423992070f22fc9ccc75e02d21c44f964736f6c63430006080033"

// DeployEntityNode deploys a new Ethereum contract, binding an instance of EntityNode to it.
func DeployEntityNode(auth *bind.TransactOpts, backend bind.ContractBackend, _entity CreativeEntities, _entityType uint8) (common.Address, *types.Transaction, *EntityNode, error) {
	parsed, err := abi.JSON(strings.NewReader(EntityNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EntityNodeBin), backend, _entity, _entityType)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EntityNode{EntityNodeCaller: EntityNodeCaller{contract: contract}, EntityNodeTransactor: EntityNodeTransactor{contract: contract}, EntityNodeFilterer: EntityNodeFilterer{contract: contract}}, nil
}

// EntityNode is an auto generated Go binding around an Ethereum contract.
type EntityNode struct {
	EntityNodeCaller     // Read-only binding to the contract
	EntityNodeTransactor // Write-only binding to the contract
	EntityNodeFilterer   // Log filterer for contract events
}

// EntityNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntityNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntityNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntityNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntityNodeSession struct {
	Contract     *EntityNode       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntityNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntityNodeCallerSession struct {
	Contract *EntityNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EntityNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntityNodeTransactorSession struct {
	Contract     *EntityNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EntityNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntityNodeRaw struct {
	Contract *EntityNode // Generic contract binding to access the raw methods on
}

// EntityNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntityNodeCallerRaw struct {
	Contract *EntityNodeCaller // Generic read-only contract binding to access the raw methods on
}

// EntityNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntityNodeTransactorRaw struct {
	Contract *EntityNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntityNode creates a new instance of EntityNode, bound to a specific deployed contract.
func NewEntityNode(address common.Address, backend bind.ContractBackend) (*EntityNode, error) {
	contract, err := bindEntityNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EntityNode{EntityNodeCaller: EntityNodeCaller{contract: contract}, EntityNodeTransactor: EntityNodeTransactor{contract: contract}, EntityNodeFilterer: EntityNodeFilterer{contract: contract}}, nil
}

// NewEntityNodeCaller creates a new read-only instance of EntityNode, bound to a specific deployed contract.
func NewEntityNodeCaller(address common.Address, caller bind.ContractCaller) (*EntityNodeCaller, error) {
	contract, err := bindEntityNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntityNodeCaller{contract: contract}, nil
}

// NewEntityNodeTransactor creates a new write-only instance of EntityNode, bound to a specific deployed contract.
func NewEntityNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*EntityNodeTransactor, error) {
	contract, err := bindEntityNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntityNodeTransactor{contract: contract}, nil
}

// NewEntityNodeFilterer creates a new log filterer instance of EntityNode, bound to a specific deployed contract.
func NewEntityNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*EntityNodeFilterer, error) {
	contract, err := bindEntityNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntityNodeFilterer{contract: contract}, nil
}

// bindEntityNode binds a generic wrapper to an already deployed contract.
func bindEntityNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EntityNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntityNode *EntityNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EntityNode.Contract.EntityNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntityNode *EntityNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityNode.Contract.EntityNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntityNode *EntityNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntityNode.Contract.EntityNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntityNode *EntityNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EntityNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntityNode *EntityNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntityNode *EntityNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntityNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((bytes32,string,string,string))
func (_EntityNode *EntityNodeCaller) Get(opts *bind.CallOpts) (CreativeEntities, error) {
	var (
		ret0 = new(CreativeEntities)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((bytes32,string,string,string))
func (_EntityNode *EntityNodeSession) Get() (CreativeEntities, error) {
	return _EntityNode.Contract.Get(&_EntityNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((bytes32,string,string,string))
func (_EntityNode *EntityNodeCallerSession) Get() (CreativeEntities, error) {
	return _EntityNode.Contract.Get(&_EntityNode.CallOpts)
}

// GetTypes is a free data retrieval call binding the contract method 0xb4579d60.
//
// Solidity: function getTypes() view returns(bool[])
func (_EntityNode *EntityNodeCaller) GetTypes(opts *bind.CallOpts) ([]bool, error) {
	var (
		ret0 = new([]bool)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "getTypes")
	return *ret0, err
}

// GetTypes is a free data retrieval call binding the contract method 0xb4579d60.
//
// Solidity: function getTypes() view returns(bool[])
func (_EntityNode *EntityNodeSession) GetTypes() ([]bool, error) {
	return _EntityNode.Contract.GetTypes(&_EntityNode.CallOpts)
}

// GetTypes is a free data retrieval call binding the contract method 0xb4579d60.
//
// Solidity: function getTypes() view returns(bool[])
func (_EntityNode *EntityNodeCallerSession) GetTypes() ([]bool, error) {
	return _EntityNode.Contract.GetTypes(&_EntityNode.CallOpts)
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_EntityNode *EntityNodeCaller) IsType(opts *bind.CallOpts, _type uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "isType", _type)
	return *ret0, err
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_EntityNode *EntityNodeSession) IsType(_type uint8) (bool, error) {
	return _EntityNode.Contract.IsType(&_EntityNode.CallOpts, _type)
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_EntityNode *EntityNodeCallerSession) IsType(_type uint8) (bool, error) {
	return _EntityNode.Contract.IsType(&_EntityNode.CallOpts, _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_EntityNode *EntityNodeTransactor) AddType(opts *bind.TransactOpts, _type uint8) (*types.Transaction, error) {
	return _EntityNode.contract.Transact(opts, "addType", _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_EntityNode *EntityNodeSession) AddType(_type uint8) (*types.Transaction, error) {
	return _EntityNode.Contract.AddType(&_EntityNode.TransactOpts, _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_EntityNode *EntityNodeTransactorSession) AddType(_type uint8) (*types.Transaction, error) {
	return _EntityNode.Contract.AddType(&_EntityNode.TransactOpts, _type)
}

// Amend is a paid mutator transaction binding the contract method 0xd1f439e5.
//
// Solidity: function amend((bytes32,string,string,string) _entity, uint8 _entityType) returns()
func (_EntityNode *EntityNodeTransactor) Amend(opts *bind.TransactOpts, _entity CreativeEntities, _entityType uint8) (*types.Transaction, error) {
	return _EntityNode.contract.Transact(opts, "amend", _entity, _entityType)
}

// Amend is a paid mutator transaction binding the contract method 0xd1f439e5.
//
// Solidity: function amend((bytes32,string,string,string) _entity, uint8 _entityType) returns()
func (_EntityNode *EntityNodeSession) Amend(_entity CreativeEntities, _entityType uint8) (*types.Transaction, error) {
	return _EntityNode.Contract.Amend(&_EntityNode.TransactOpts, _entity, _entityType)
}

// Amend is a paid mutator transaction binding the contract method 0xd1f439e5.
//
// Solidity: function amend((bytes32,string,string,string) _entity, uint8 _entityType) returns()
func (_EntityNode *EntityNodeTransactorSession) Amend(_entity CreativeEntities, _entityType uint8) (*types.Transaction, error) {
	return _EntityNode.Contract.Amend(&_EntityNode.TransactOpts, _entity, _entityType)
}

// IArtefactABI is the input ABI used to generate the binding from.
const IArtefactABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"addAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"addCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"addPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthors\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCopyrightHolders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublishers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isCopyrightHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isPublisher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"removeAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"removeCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"removePublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArtefactFuncSigs maps the 4-byte function signature to its string representation.
var IArtefactFuncSigs = map[string]string{
	"a9c56065": "addAuthor(bytes32)",
	"c3a30a6f": "addCopyrightHolder(bytes32)",
	"c28b16dc": "addPublisher(bytes32)",
	"9673484e": "amend((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))",
	"6d4ce63c": "get()",
	"7fc622f4": "getAuthors()",
	"d6d54164": "getCopyrightHolders()",
	"6c6071aa": "getPublishers()",
	"59c0161c": "isAuthor(bytes32)",
	"476caa00": "isCopyrightHolder(bytes32)",
	"240b5600": "isPublisher(bytes32)",
	"33966d03": "removeAuthor(bytes32)",
	"2424dccc": "removeCopyrightHolder(bytes32)",
	"d13c8f23": "removePublisher(bytes32)",
}

// IArtefact is an auto generated Go binding around an Ethereum contract.
type IArtefact struct {
	IArtefactCaller     // Read-only binding to the contract
	IArtefactTransactor // Write-only binding to the contract
	IArtefactFilterer   // Log filterer for contract events
}

// IArtefactCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArtefactCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArtefactTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArtefactTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArtefactFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArtefactFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArtefactSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArtefactSession struct {
	Contract     *IArtefact        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArtefactCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArtefactCallerSession struct {
	Contract *IArtefactCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IArtefactTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArtefactTransactorSession struct {
	Contract     *IArtefactTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IArtefactRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArtefactRaw struct {
	Contract *IArtefact // Generic contract binding to access the raw methods on
}

// IArtefactCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArtefactCallerRaw struct {
	Contract *IArtefactCaller // Generic read-only contract binding to access the raw methods on
}

// IArtefactTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArtefactTransactorRaw struct {
	Contract *IArtefactTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArtefact creates a new instance of IArtefact, bound to a specific deployed contract.
func NewIArtefact(address common.Address, backend bind.ContractBackend) (*IArtefact, error) {
	contract, err := bindIArtefact(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArtefact{IArtefactCaller: IArtefactCaller{contract: contract}, IArtefactTransactor: IArtefactTransactor{contract: contract}, IArtefactFilterer: IArtefactFilterer{contract: contract}}, nil
}

// NewIArtefactCaller creates a new read-only instance of IArtefact, bound to a specific deployed contract.
func NewIArtefactCaller(address common.Address, caller bind.ContractCaller) (*IArtefactCaller, error) {
	contract, err := bindIArtefact(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArtefactCaller{contract: contract}, nil
}

// NewIArtefactTransactor creates a new write-only instance of IArtefact, bound to a specific deployed contract.
func NewIArtefactTransactor(address common.Address, transactor bind.ContractTransactor) (*IArtefactTransactor, error) {
	contract, err := bindIArtefact(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArtefactTransactor{contract: contract}, nil
}

// NewIArtefactFilterer creates a new log filterer instance of IArtefact, bound to a specific deployed contract.
func NewIArtefactFilterer(address common.Address, filterer bind.ContractFilterer) (*IArtefactFilterer, error) {
	contract, err := bindIArtefact(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArtefactFilterer{contract: contract}, nil
}

// bindIArtefact binds a generic wrapper to an already deployed contract.
func bindIArtefact(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArtefactABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArtefact *IArtefactRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArtefact.Contract.IArtefactCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArtefact *IArtefactRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArtefact.Contract.IArtefactTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArtefact *IArtefactRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArtefact.Contract.IArtefactTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArtefact *IArtefactCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArtefact.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArtefact *IArtefactTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArtefact.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArtefact *IArtefactTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArtefact.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_IArtefact *IArtefactCaller) Get(opts *bind.CallOpts) (CreativeWorks, error) {
	var (
		ret0 = new(CreativeWorks)
	)
	out := ret0
	err := _IArtefact.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_IArtefact *IArtefactSession) Get() (CreativeWorks, error) {
	return _IArtefact.Contract.Get(&_IArtefact.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_IArtefact *IArtefactCallerSession) Get() (CreativeWorks, error) {
	return _IArtefact.Contract.Get(&_IArtefact.CallOpts)
}

// GetAuthors is a free data retrieval call binding the contract method 0x7fc622f4.
//
// Solidity: function getAuthors() view returns(bytes32[])
func (_IArtefact *IArtefactCaller) GetAuthors(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IArtefact.contract.Call(opts, out, "getAuthors")
	return *ret0, err
}

// GetAuthors is a free data retrieval call binding the contract method 0x7fc622f4.
//
// Solidity: function getAuthors() view returns(bytes32[])
func (_IArtefact *IArtefactSession) GetAuthors() ([][32]byte, error) {
	return _IArtefact.Contract.GetAuthors(&_IArtefact.CallOpts)
}

// GetAuthors is a free data retrieval call binding the contract method 0x7fc622f4.
//
// Solidity: function getAuthors() view returns(bytes32[])
func (_IArtefact *IArtefactCallerSession) GetAuthors() ([][32]byte, error) {
	return _IArtefact.Contract.GetAuthors(&_IArtefact.CallOpts)
}

// GetCopyrightHolders is a free data retrieval call binding the contract method 0xd6d54164.
//
// Solidity: function getCopyrightHolders() view returns(bytes32[])
func (_IArtefact *IArtefactCaller) GetCopyrightHolders(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IArtefact.contract.Call(opts, out, "getCopyrightHolders")
	return *ret0, err
}

// GetCopyrightHolders is a free data retrieval call binding the contract method 0xd6d54164.
//
// Solidity: function getCopyrightHolders() view returns(bytes32[])
func (_IArtefact *IArtefactSession) GetCopyrightHolders() ([][32]byte, error) {
	return _IArtefact.Contract.GetCopyrightHolders(&_IArtefact.CallOpts)
}

// GetCopyrightHolders is a free data retrieval call binding the contract method 0xd6d54164.
//
// Solidity: function getCopyrightHolders() view returns(bytes32[])
func (_IArtefact *IArtefactCallerSession) GetCopyrightHolders() ([][32]byte, error) {
	return _IArtefact.Contract.GetCopyrightHolders(&_IArtefact.CallOpts)
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(bytes32[])
func (_IArtefact *IArtefactCaller) GetPublishers(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IArtefact.contract.Call(opts, out, "getPublishers")
	return *ret0, err
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(bytes32[])
func (_IArtefact *IArtefactSession) GetPublishers() ([][32]byte, error) {
	return _IArtefact.Contract.GetPublishers(&_IArtefact.CallOpts)
}

// GetPublishers is a free data retrieval call binding the contract method 0x6c6071aa.
//
// Solidity: function getPublishers() view returns(bytes32[])
func (_IArtefact *IArtefactCallerSession) GetPublishers() ([][32]byte, error) {
	return _IArtefact.Contract.GetPublishers(&_IArtefact.CallOpts)
}

// IsAuthor is a free data retrieval call binding the contract method 0x59c0161c.
//
// Solidity: function isAuthor(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactCaller) IsAuthor(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IArtefact.contract.Call(opts, out, "isAuthor", _id)
	return *ret0, err
}

// IsAuthor is a free data retrieval call binding the contract method 0x59c0161c.
//
// Solidity: function isAuthor(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactSession) IsAuthor(_id [32]byte) (bool, error) {
	return _IArtefact.Contract.IsAuthor(&_IArtefact.CallOpts, _id)
}

// IsAuthor is a free data retrieval call binding the contract method 0x59c0161c.
//
// Solidity: function isAuthor(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactCallerSession) IsAuthor(_id [32]byte) (bool, error) {
	return _IArtefact.Contract.IsAuthor(&_IArtefact.CallOpts, _id)
}

// IsCopyrightHolder is a free data retrieval call binding the contract method 0x476caa00.
//
// Solidity: function isCopyrightHolder(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactCaller) IsCopyrightHolder(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IArtefact.contract.Call(opts, out, "isCopyrightHolder", _id)
	return *ret0, err
}

// IsCopyrightHolder is a free data retrieval call binding the contract method 0x476caa00.
//
// Solidity: function isCopyrightHolder(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactSession) IsCopyrightHolder(_id [32]byte) (bool, error) {
	return _IArtefact.Contract.IsCopyrightHolder(&_IArtefact.CallOpts, _id)
}

// IsCopyrightHolder is a free data retrieval call binding the contract method 0x476caa00.
//
// Solidity: function isCopyrightHolder(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactCallerSession) IsCopyrightHolder(_id [32]byte) (bool, error) {
	return _IArtefact.Contract.IsCopyrightHolder(&_IArtefact.CallOpts, _id)
}

// IsPublisher is a free data retrieval call binding the contract method 0x240b5600.
//
// Solidity: function isPublisher(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactCaller) IsPublisher(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IArtefact.contract.Call(opts, out, "isPublisher", _id)
	return *ret0, err
}

// IsPublisher is a free data retrieval call binding the contract method 0x240b5600.
//
// Solidity: function isPublisher(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactSession) IsPublisher(_id [32]byte) (bool, error) {
	return _IArtefact.Contract.IsPublisher(&_IArtefact.CallOpts, _id)
}

// IsPublisher is a free data retrieval call binding the contract method 0x240b5600.
//
// Solidity: function isPublisher(bytes32 _id) view returns(bool)
func (_IArtefact *IArtefactCallerSession) IsPublisher(_id [32]byte) (bool, error) {
	return _IArtefact.Contract.IsPublisher(&_IArtefact.CallOpts, _id)
}

// AddAuthor is a paid mutator transaction binding the contract method 0xa9c56065.
//
// Solidity: function addAuthor(bytes32 _authorId) returns()
func (_IArtefact *IArtefactTransactor) AddAuthor(opts *bind.TransactOpts, _authorId [32]byte) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "addAuthor", _authorId)
}

// AddAuthor is a paid mutator transaction binding the contract method 0xa9c56065.
//
// Solidity: function addAuthor(bytes32 _authorId) returns()
func (_IArtefact *IArtefactSession) AddAuthor(_authorId [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.AddAuthor(&_IArtefact.TransactOpts, _authorId)
}

// AddAuthor is a paid mutator transaction binding the contract method 0xa9c56065.
//
// Solidity: function addAuthor(bytes32 _authorId) returns()
func (_IArtefact *IArtefactTransactorSession) AddAuthor(_authorId [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.AddAuthor(&_IArtefact.TransactOpts, _authorId)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0xc3a30a6f.
//
// Solidity: function addCopyrightHolder(bytes32 _copyrightHolderId) returns()
func (_IArtefact *IArtefactTransactor) AddCopyrightHolder(opts *bind.TransactOpts, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "addCopyrightHolder", _copyrightHolderId)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0xc3a30a6f.
//
// Solidity: function addCopyrightHolder(bytes32 _copyrightHolderId) returns()
func (_IArtefact *IArtefactSession) AddCopyrightHolder(_copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.AddCopyrightHolder(&_IArtefact.TransactOpts, _copyrightHolderId)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0xc3a30a6f.
//
// Solidity: function addCopyrightHolder(bytes32 _copyrightHolderId) returns()
func (_IArtefact *IArtefactTransactorSession) AddCopyrightHolder(_copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.AddCopyrightHolder(&_IArtefact.TransactOpts, _copyrightHolderId)
}

// AddPublisher is a paid mutator transaction binding the contract method 0xc28b16dc.
//
// Solidity: function addPublisher(bytes32 _publisherId) returns()
func (_IArtefact *IArtefactTransactor) AddPublisher(opts *bind.TransactOpts, _publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "addPublisher", _publisherId)
}

// AddPublisher is a paid mutator transaction binding the contract method 0xc28b16dc.
//
// Solidity: function addPublisher(bytes32 _publisherId) returns()
func (_IArtefact *IArtefactSession) AddPublisher(_publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.AddPublisher(&_IArtefact.TransactOpts, _publisherId)
}

// AddPublisher is a paid mutator transaction binding the contract method 0xc28b16dc.
//
// Solidity: function addPublisher(bytes32 _publisherId) returns()
func (_IArtefact *IArtefactTransactorSession) AddPublisher(_publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.AddPublisher(&_IArtefact.TransactOpts, _publisherId)
}

// Amend is a paid mutator transaction binding the contract method 0x9673484e.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefact *IArtefactTransactor) Amend(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "amend", _work)
}

// Amend is a paid mutator transaction binding the contract method 0x9673484e.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefact *IArtefactSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.Contract.Amend(&_IArtefact.TransactOpts, _work)
}

// Amend is a paid mutator transaction binding the contract method 0x9673484e.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefact *IArtefactTransactorSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.Contract.Amend(&_IArtefact.TransactOpts, _work)
}

// RemoveAuthor is a paid mutator transaction binding the contract method 0x33966d03.
//
// Solidity: function removeAuthor(bytes32 _id) returns()
func (_IArtefact *IArtefactTransactor) RemoveAuthor(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "removeAuthor", _id)
}

// RemoveAuthor is a paid mutator transaction binding the contract method 0x33966d03.
//
// Solidity: function removeAuthor(bytes32 _id) returns()
func (_IArtefact *IArtefactSession) RemoveAuthor(_id [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.RemoveAuthor(&_IArtefact.TransactOpts, _id)
}

// RemoveAuthor is a paid mutator transaction binding the contract method 0x33966d03.
//
// Solidity: function removeAuthor(bytes32 _id) returns()
func (_IArtefact *IArtefactTransactorSession) RemoveAuthor(_id [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.RemoveAuthor(&_IArtefact.TransactOpts, _id)
}

// RemoveCopyrightHolder is a paid mutator transaction binding the contract method 0x2424dccc.
//
// Solidity: function removeCopyrightHolder(bytes32 _id) returns()
func (_IArtefact *IArtefactTransactor) RemoveCopyrightHolder(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "removeCopyrightHolder", _id)
}

// RemoveCopyrightHolder is a paid mutator transaction binding the contract method 0x2424dccc.
//
// Solidity: function removeCopyrightHolder(bytes32 _id) returns()
func (_IArtefact *IArtefactSession) RemoveCopyrightHolder(_id [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.RemoveCopyrightHolder(&_IArtefact.TransactOpts, _id)
}

// RemoveCopyrightHolder is a paid mutator transaction binding the contract method 0x2424dccc.
//
// Solidity: function removeCopyrightHolder(bytes32 _id) returns()
func (_IArtefact *IArtefactTransactorSession) RemoveCopyrightHolder(_id [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.RemoveCopyrightHolder(&_IArtefact.TransactOpts, _id)
}

// RemovePublisher is a paid mutator transaction binding the contract method 0xd13c8f23.
//
// Solidity: function removePublisher(bytes32 _id) returns()
func (_IArtefact *IArtefactTransactor) RemovePublisher(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "removePublisher", _id)
}

// RemovePublisher is a paid mutator transaction binding the contract method 0xd13c8f23.
//
// Solidity: function removePublisher(bytes32 _id) returns()
func (_IArtefact *IArtefactSession) RemovePublisher(_id [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.RemovePublisher(&_IArtefact.TransactOpts, _id)
}

// RemovePublisher is a paid mutator transaction binding the contract method 0xd13c8f23.
//
// Solidity: function removePublisher(bytes32 _id) returns()
func (_IArtefact *IArtefactTransactorSession) RemovePublisher(_id [32]byte) (*types.Transaction, error) {
	return _IArtefact.Contract.RemovePublisher(&_IArtefact.TransactOpts, _id)
}

// IArtefactFactoryABI is the input ABI used to generate the binding from.
const IArtefactFactoryABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"addWorkAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"addWorkCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"addWorkPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amendWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWork\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"authorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkAuthors\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkCopyrightHolders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkPublishers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"isWorkAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"isWorkCopyrightHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"isWorkPublisher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"removeWorkAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"removeWorkCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"removeWorkPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IArtefactFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IArtefactFactoryFuncSigs = map[string]string{
	"1d339e24": "addWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))",
	"55884f0f": "addWorkAuthor(bytes32,bytes32)",
	"0f1ee4a2": "addWorkCopyrightHolder(bytes32,bytes32)",
	"e17989ea": "addWorkPublisher(bytes32,bytes32)",
	"4e542a6e": "amendWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))",
	"30266537": "getWork(bytes32)",
	"426aa803": "getWorkAuthors(bytes32)",
	"4d2522a2": "getWorkContract(bytes32)",
	"86105c65": "getWorkCopyrightHolders(bytes32)",
	"5ffd09d8": "getWorkPublishers(bytes32)",
	"b2e459cf": "isWorkAuthor(bytes32,bytes32)",
	"f6a50965": "isWorkCopyrightHolder(bytes32,bytes32)",
	"ef90d9c9": "isWorkPublisher(bytes32,bytes32)",
	"e229e173": "removeWorkAuthor(bytes32,bytes32)",
	"521ffdce": "removeWorkCopyrightHolder(bytes32,bytes32)",
	"0c99420f": "removeWorkPublisher(bytes32,bytes32)",
}

// IArtefactFactory is an auto generated Go binding around an Ethereum contract.
type IArtefactFactory struct {
	IArtefactFactoryCaller     // Read-only binding to the contract
	IArtefactFactoryTransactor // Write-only binding to the contract
	IArtefactFactoryFilterer   // Log filterer for contract events
}

// IArtefactFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IArtefactFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArtefactFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IArtefactFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArtefactFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IArtefactFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IArtefactFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IArtefactFactorySession struct {
	Contract     *IArtefactFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IArtefactFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IArtefactFactoryCallerSession struct {
	Contract *IArtefactFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IArtefactFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IArtefactFactoryTransactorSession struct {
	Contract     *IArtefactFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IArtefactFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IArtefactFactoryRaw struct {
	Contract *IArtefactFactory // Generic contract binding to access the raw methods on
}

// IArtefactFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IArtefactFactoryCallerRaw struct {
	Contract *IArtefactFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IArtefactFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IArtefactFactoryTransactorRaw struct {
	Contract *IArtefactFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIArtefactFactory creates a new instance of IArtefactFactory, bound to a specific deployed contract.
func NewIArtefactFactory(address common.Address, backend bind.ContractBackend) (*IArtefactFactory, error) {
	contract, err := bindIArtefactFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IArtefactFactory{IArtefactFactoryCaller: IArtefactFactoryCaller{contract: contract}, IArtefactFactoryTransactor: IArtefactFactoryTransactor{contract: contract}, IArtefactFactoryFilterer: IArtefactFactoryFilterer{contract: contract}}, nil
}

// NewIArtefactFactoryCaller creates a new read-only instance of IArtefactFactory, bound to a specific deployed contract.
func NewIArtefactFactoryCaller(address common.Address, caller bind.ContractCaller) (*IArtefactFactoryCaller, error) {
	contract, err := bindIArtefactFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IArtefactFactoryCaller{contract: contract}, nil
}

// NewIArtefactFactoryTransactor creates a new write-only instance of IArtefactFactory, bound to a specific deployed contract.
func NewIArtefactFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IArtefactFactoryTransactor, error) {
	contract, err := bindIArtefactFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IArtefactFactoryTransactor{contract: contract}, nil
}

// NewIArtefactFactoryFilterer creates a new log filterer instance of IArtefactFactory, bound to a specific deployed contract.
func NewIArtefactFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IArtefactFactoryFilterer, error) {
	contract, err := bindIArtefactFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IArtefactFactoryFilterer{contract: contract}, nil
}

// bindIArtefactFactory binds a generic wrapper to an already deployed contract.
func bindIArtefactFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IArtefactFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArtefactFactory *IArtefactFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArtefactFactory.Contract.IArtefactFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArtefactFactory *IArtefactFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.IArtefactFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArtefactFactory *IArtefactFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.IArtefactFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IArtefactFactory *IArtefactFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IArtefactFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IArtefactFactory *IArtefactFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IArtefactFactory *IArtefactFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.contract.Transact(opts, method, params...)
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_IArtefactFactory *IArtefactFactoryCaller) GetWork(opts *bind.CallOpts, _id [32]byte) (CreativeWorks, error) {
	var (
		ret0 = new(CreativeWorks)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "getWork", _id)
	return *ret0, err
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_IArtefactFactory *IArtefactFactorySession) GetWork(_id [32]byte) (CreativeWorks, error) {
	return _IArtefactFactory.Contract.GetWork(&_IArtefactFactory.CallOpts, _id)
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string))
func (_IArtefactFactory *IArtefactFactoryCallerSession) GetWork(_id [32]byte) (CreativeWorks, error) {
	return _IArtefactFactory.Contract.GetWork(&_IArtefactFactory.CallOpts, _id)
}

// GetWorkAuthors is a free data retrieval call binding the contract method 0x426aa803.
//
// Solidity: function getWorkAuthors(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactoryCaller) GetWorkAuthors(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "getWorkAuthors", _id)
	return *ret0, err
}

// GetWorkAuthors is a free data retrieval call binding the contract method 0x426aa803.
//
// Solidity: function getWorkAuthors(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactorySession) GetWorkAuthors(_id [32]byte) ([][32]byte, error) {
	return _IArtefactFactory.Contract.GetWorkAuthors(&_IArtefactFactory.CallOpts, _id)
}

// GetWorkAuthors is a free data retrieval call binding the contract method 0x426aa803.
//
// Solidity: function getWorkAuthors(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactoryCallerSession) GetWorkAuthors(_id [32]byte) ([][32]byte, error) {
	return _IArtefactFactory.Contract.GetWorkAuthors(&_IArtefactFactory.CallOpts, _id)
}

// GetWorkContract is a free data retrieval call binding the contract method 0x4d2522a2.
//
// Solidity: function getWorkContract(bytes32 _id) view returns(address)
func (_IArtefactFactory *IArtefactFactoryCaller) GetWorkContract(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "getWorkContract", _id)
	return *ret0, err
}

// GetWorkContract is a free data retrieval call binding the contract method 0x4d2522a2.
//
// Solidity: function getWorkContract(bytes32 _id) view returns(address)
func (_IArtefactFactory *IArtefactFactorySession) GetWorkContract(_id [32]byte) (common.Address, error) {
	return _IArtefactFactory.Contract.GetWorkContract(&_IArtefactFactory.CallOpts, _id)
}

// GetWorkContract is a free data retrieval call binding the contract method 0x4d2522a2.
//
// Solidity: function getWorkContract(bytes32 _id) view returns(address)
func (_IArtefactFactory *IArtefactFactoryCallerSession) GetWorkContract(_id [32]byte) (common.Address, error) {
	return _IArtefactFactory.Contract.GetWorkContract(&_IArtefactFactory.CallOpts, _id)
}

// GetWorkCopyrightHolders is a free data retrieval call binding the contract method 0x86105c65.
//
// Solidity: function getWorkCopyrightHolders(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactoryCaller) GetWorkCopyrightHolders(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "getWorkCopyrightHolders", _id)
	return *ret0, err
}

// GetWorkCopyrightHolders is a free data retrieval call binding the contract method 0x86105c65.
//
// Solidity: function getWorkCopyrightHolders(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactorySession) GetWorkCopyrightHolders(_id [32]byte) ([][32]byte, error) {
	return _IArtefactFactory.Contract.GetWorkCopyrightHolders(&_IArtefactFactory.CallOpts, _id)
}

// GetWorkCopyrightHolders is a free data retrieval call binding the contract method 0x86105c65.
//
// Solidity: function getWorkCopyrightHolders(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactoryCallerSession) GetWorkCopyrightHolders(_id [32]byte) ([][32]byte, error) {
	return _IArtefactFactory.Contract.GetWorkCopyrightHolders(&_IArtefactFactory.CallOpts, _id)
}

// GetWorkPublishers is a free data retrieval call binding the contract method 0x5ffd09d8.
//
// Solidity: function getWorkPublishers(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactoryCaller) GetWorkPublishers(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "getWorkPublishers", _id)
	return *ret0, err
}

// GetWorkPublishers is a free data retrieval call binding the contract method 0x5ffd09d8.
//
// Solidity: function getWorkPublishers(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactorySession) GetWorkPublishers(_id [32]byte) ([][32]byte, error) {
	return _IArtefactFactory.Contract.GetWorkPublishers(&_IArtefactFactory.CallOpts, _id)
}

// GetWorkPublishers is a free data retrieval call binding the contract method 0x5ffd09d8.
//
// Solidity: function getWorkPublishers(bytes32 _id) view returns(bytes32[])
func (_IArtefactFactory *IArtefactFactoryCallerSession) GetWorkPublishers(_id [32]byte) ([][32]byte, error) {
	return _IArtefactFactory.Contract.GetWorkPublishers(&_IArtefactFactory.CallOpts, _id)
}

// IsWorkAuthor is a free data retrieval call binding the contract method 0xb2e459cf.
//
// Solidity: function isWorkAuthor(bytes32 _workId, bytes32 _authorId) view returns(bool)
func (_IArtefactFactory *IArtefactFactoryCaller) IsWorkAuthor(opts *bind.CallOpts, _workId [32]byte, _authorId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "isWorkAuthor", _workId, _authorId)
	return *ret0, err
}

// IsWorkAuthor is a free data retrieval call binding the contract method 0xb2e459cf.
//
// Solidity: function isWorkAuthor(bytes32 _workId, bytes32 _authorId) view returns(bool)
func (_IArtefactFactory *IArtefactFactorySession) IsWorkAuthor(_workId [32]byte, _authorId [32]byte) (bool, error) {
	return _IArtefactFactory.Contract.IsWorkAuthor(&_IArtefactFactory.CallOpts, _workId, _authorId)
}

// IsWorkAuthor is a free data retrieval call binding the contract method 0xb2e459cf.
//
// Solidity: function isWorkAuthor(bytes32 _workId, bytes32 _authorId) view returns(bool)
func (_IArtefactFactory *IArtefactFactoryCallerSession) IsWorkAuthor(_workId [32]byte, _authorId [32]byte) (bool, error) {
	return _IArtefactFactory.Contract.IsWorkAuthor(&_IArtefactFactory.CallOpts, _workId, _authorId)
}

// IsWorkCopyrightHolder is a free data retrieval call binding the contract method 0xf6a50965.
//
// Solidity: function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) view returns(bool)
func (_IArtefactFactory *IArtefactFactoryCaller) IsWorkCopyrightHolder(opts *bind.CallOpts, _workId [32]byte, _copyrightHolderId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "isWorkCopyrightHolder", _workId, _copyrightHolderId)
	return *ret0, err
}

// IsWorkCopyrightHolder is a free data retrieval call binding the contract method 0xf6a50965.
//
// Solidity: function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) view returns(bool)
func (_IArtefactFactory *IArtefactFactorySession) IsWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (bool, error) {
	return _IArtefactFactory.Contract.IsWorkCopyrightHolder(&_IArtefactFactory.CallOpts, _workId, _copyrightHolderId)
}

// IsWorkCopyrightHolder is a free data retrieval call binding the contract method 0xf6a50965.
//
// Solidity: function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) view returns(bool)
func (_IArtefactFactory *IArtefactFactoryCallerSession) IsWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (bool, error) {
	return _IArtefactFactory.Contract.IsWorkCopyrightHolder(&_IArtefactFactory.CallOpts, _workId, _copyrightHolderId)
}

// IsWorkPublisher is a free data retrieval call binding the contract method 0xef90d9c9.
//
// Solidity: function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) view returns(bool)
func (_IArtefactFactory *IArtefactFactoryCaller) IsWorkPublisher(opts *bind.CallOpts, _workId [32]byte, _publisherId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "isWorkPublisher", _workId, _publisherId)
	return *ret0, err
}

// IsWorkPublisher is a free data retrieval call binding the contract method 0xef90d9c9.
//
// Solidity: function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) view returns(bool)
func (_IArtefactFactory *IArtefactFactorySession) IsWorkPublisher(_workId [32]byte, _publisherId [32]byte) (bool, error) {
	return _IArtefactFactory.Contract.IsWorkPublisher(&_IArtefactFactory.CallOpts, _workId, _publisherId)
}

// IsWorkPublisher is a free data retrieval call binding the contract method 0xef90d9c9.
//
// Solidity: function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) view returns(bool)
func (_IArtefactFactory *IArtefactFactoryCallerSession) IsWorkPublisher(_workId [32]byte, _publisherId [32]byte) (bool, error) {
	return _IArtefactFactory.Contract.IsWorkPublisher(&_IArtefactFactory.CallOpts, _workId, _publisherId)
}

// AddWork is a paid mutator transaction binding the contract method 0x1d339e24.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWork", _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x1d339e24.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWork(&_IArtefactFactory.TransactOpts, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x1d339e24.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWork(&_IArtefactFactory.TransactOpts, _work)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x55884f0f.
//
// Solidity: function addWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWorkAuthor(opts *bind.TransactOpts, _workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWorkAuthor", _workId, _authorId)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x55884f0f.
//
// Solidity: function addWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWorkAuthor(_workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkAuthor(&_IArtefactFactory.TransactOpts, _workId, _authorId)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x55884f0f.
//
// Solidity: function addWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWorkAuthor(_workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkAuthor(&_IArtefactFactory.TransactOpts, _workId, _authorId)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x0f1ee4a2.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWorkCopyrightHolder(opts *bind.TransactOpts, _workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWorkCopyrightHolder", _workId, _copyrightHolderId)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x0f1ee4a2.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkCopyrightHolder(&_IArtefactFactory.TransactOpts, _workId, _copyrightHolderId)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x0f1ee4a2.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkCopyrightHolder(&_IArtefactFactory.TransactOpts, _workId, _copyrightHolderId)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0xe17989ea.
//
// Solidity: function addWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWorkPublisher(opts *bind.TransactOpts, _workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWorkPublisher", _workId, _publisherId)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0xe17989ea.
//
// Solidity: function addWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWorkPublisher(_workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkPublisher(&_IArtefactFactory.TransactOpts, _workId, _publisherId)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0xe17989ea.
//
// Solidity: function addWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWorkPublisher(_workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkPublisher(&_IArtefactFactory.TransactOpts, _workId, _publisherId)
}

// AmendWork is a paid mutator transaction binding the contract method 0x4e542a6e.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AmendWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "amendWork", _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0x4e542a6e.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefactFactory *IArtefactFactorySession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AmendWork(&_IArtefactFactory.TransactOpts, _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0x4e542a6e.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,bytes32,string,string,string) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AmendWork(&_IArtefactFactory.TransactOpts, _work)
}

// RemoveWorkAuthor is a paid mutator transaction binding the contract method 0xe229e173.
//
// Solidity: function removeWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) RemoveWorkAuthor(opts *bind.TransactOpts, _workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "removeWorkAuthor", _workId, _authorId)
}

// RemoveWorkAuthor is a paid mutator transaction binding the contract method 0xe229e173.
//
// Solidity: function removeWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_IArtefactFactory *IArtefactFactorySession) RemoveWorkAuthor(_workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.RemoveWorkAuthor(&_IArtefactFactory.TransactOpts, _workId, _authorId)
}

// RemoveWorkAuthor is a paid mutator transaction binding the contract method 0xe229e173.
//
// Solidity: function removeWorkAuthor(bytes32 _workId, bytes32 _authorId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) RemoveWorkAuthor(_workId [32]byte, _authorId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.RemoveWorkAuthor(&_IArtefactFactory.TransactOpts, _workId, _authorId)
}

// RemoveWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x521ffdce.
//
// Solidity: function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) RemoveWorkCopyrightHolder(opts *bind.TransactOpts, _workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "removeWorkCopyrightHolder", _workId, _copyrightHolderId)
}

// RemoveWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x521ffdce.
//
// Solidity: function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_IArtefactFactory *IArtefactFactorySession) RemoveWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.RemoveWorkCopyrightHolder(&_IArtefactFactory.TransactOpts, _workId, _copyrightHolderId)
}

// RemoveWorkCopyrightHolder is a paid mutator transaction binding the contract method 0x521ffdce.
//
// Solidity: function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) RemoveWorkCopyrightHolder(_workId [32]byte, _copyrightHolderId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.RemoveWorkCopyrightHolder(&_IArtefactFactory.TransactOpts, _workId, _copyrightHolderId)
}

// RemoveWorkPublisher is a paid mutator transaction binding the contract method 0x0c99420f.
//
// Solidity: function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) RemoveWorkPublisher(opts *bind.TransactOpts, _workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "removeWorkPublisher", _workId, _publisherId)
}

// RemoveWorkPublisher is a paid mutator transaction binding the contract method 0x0c99420f.
//
// Solidity: function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_IArtefactFactory *IArtefactFactorySession) RemoveWorkPublisher(_workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.RemoveWorkPublisher(&_IArtefactFactory.TransactOpts, _workId, _publisherId)
}

// RemoveWorkPublisher is a paid mutator transaction binding the contract method 0x0c99420f.
//
// Solidity: function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) RemoveWorkPublisher(_workId [32]byte, _publisherId [32]byte) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.RemoveWorkPublisher(&_IArtefactFactory.TransactOpts, _workId, _publisherId)
}

// IEntitiesFactoryABI is the input ABI used to generate the binding from.
const IEntitiesFactoryABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amendEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntity\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isEntityType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEntitiesFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IEntitiesFactoryFuncSigs = map[string]string{
	"68a9f4e9": "addEntity((bytes32,string,string,string),uint8)",
	"a81eb545": "amendEntity((bytes32,string,string,string),uint8)",
	"53b66f36": "getEntity(bytes32)",
	"d91ee2e0": "getEntityContract(bytes32)",
	"6ec21da9": "getEntityTypes(bytes32)",
	"cef8b362": "isEntityType(bytes32,uint8)",
}

// IEntitiesFactory is an auto generated Go binding around an Ethereum contract.
type IEntitiesFactory struct {
	IEntitiesFactoryCaller     // Read-only binding to the contract
	IEntitiesFactoryTransactor // Write-only binding to the contract
	IEntitiesFactoryFilterer   // Log filterer for contract events
}

// IEntitiesFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEntitiesFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntitiesFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEntitiesFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntitiesFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEntitiesFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntitiesFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEntitiesFactorySession struct {
	Contract     *IEntitiesFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEntitiesFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEntitiesFactoryCallerSession struct {
	Contract *IEntitiesFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IEntitiesFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEntitiesFactoryTransactorSession struct {
	Contract     *IEntitiesFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IEntitiesFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEntitiesFactoryRaw struct {
	Contract *IEntitiesFactory // Generic contract binding to access the raw methods on
}

// IEntitiesFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEntitiesFactoryCallerRaw struct {
	Contract *IEntitiesFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IEntitiesFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEntitiesFactoryTransactorRaw struct {
	Contract *IEntitiesFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEntitiesFactory creates a new instance of IEntitiesFactory, bound to a specific deployed contract.
func NewIEntitiesFactory(address common.Address, backend bind.ContractBackend) (*IEntitiesFactory, error) {
	contract, err := bindIEntitiesFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEntitiesFactory{IEntitiesFactoryCaller: IEntitiesFactoryCaller{contract: contract}, IEntitiesFactoryTransactor: IEntitiesFactoryTransactor{contract: contract}, IEntitiesFactoryFilterer: IEntitiesFactoryFilterer{contract: contract}}, nil
}

// NewIEntitiesFactoryCaller creates a new read-only instance of IEntitiesFactory, bound to a specific deployed contract.
func NewIEntitiesFactoryCaller(address common.Address, caller bind.ContractCaller) (*IEntitiesFactoryCaller, error) {
	contract, err := bindIEntitiesFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEntitiesFactoryCaller{contract: contract}, nil
}

// NewIEntitiesFactoryTransactor creates a new write-only instance of IEntitiesFactory, bound to a specific deployed contract.
func NewIEntitiesFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IEntitiesFactoryTransactor, error) {
	contract, err := bindIEntitiesFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEntitiesFactoryTransactor{contract: contract}, nil
}

// NewIEntitiesFactoryFilterer creates a new log filterer instance of IEntitiesFactory, bound to a specific deployed contract.
func NewIEntitiesFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IEntitiesFactoryFilterer, error) {
	contract, err := bindIEntitiesFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEntitiesFactoryFilterer{contract: contract}, nil
}

// bindIEntitiesFactory binds a generic wrapper to an already deployed contract.
func bindIEntitiesFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEntitiesFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEntitiesFactory *IEntitiesFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEntitiesFactory.Contract.IEntitiesFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEntitiesFactory *IEntitiesFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.IEntitiesFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEntitiesFactory *IEntitiesFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.IEntitiesFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEntitiesFactory *IEntitiesFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEntitiesFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEntitiesFactory *IEntitiesFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEntitiesFactory *IEntitiesFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.contract.Transact(opts, method, params...)
}

// GetEntity is a free data retrieval call binding the contract method 0x53b66f36.
//
// Solidity: function getEntity(bytes32 _id) view returns((bytes32,string,string,string))
func (_IEntitiesFactory *IEntitiesFactoryCaller) GetEntity(opts *bind.CallOpts, _id [32]byte) (CreativeEntities, error) {
	var (
		ret0 = new(CreativeEntities)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "getEntity", _id)
	return *ret0, err
}

// GetEntity is a free data retrieval call binding the contract method 0x53b66f36.
//
// Solidity: function getEntity(bytes32 _id) view returns((bytes32,string,string,string))
func (_IEntitiesFactory *IEntitiesFactorySession) GetEntity(_id [32]byte) (CreativeEntities, error) {
	return _IEntitiesFactory.Contract.GetEntity(&_IEntitiesFactory.CallOpts, _id)
}

// GetEntity is a free data retrieval call binding the contract method 0x53b66f36.
//
// Solidity: function getEntity(bytes32 _id) view returns((bytes32,string,string,string))
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) GetEntity(_id [32]byte) (CreativeEntities, error) {
	return _IEntitiesFactory.Contract.GetEntity(&_IEntitiesFactory.CallOpts, _id)
}

// GetEntityContract is a free data retrieval call binding the contract method 0xd91ee2e0.
//
// Solidity: function getEntityContract(bytes32 _id) view returns(address)
func (_IEntitiesFactory *IEntitiesFactoryCaller) GetEntityContract(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "getEntityContract", _id)
	return *ret0, err
}

// GetEntityContract is a free data retrieval call binding the contract method 0xd91ee2e0.
//
// Solidity: function getEntityContract(bytes32 _id) view returns(address)
func (_IEntitiesFactory *IEntitiesFactorySession) GetEntityContract(_id [32]byte) (common.Address, error) {
	return _IEntitiesFactory.Contract.GetEntityContract(&_IEntitiesFactory.CallOpts, _id)
}

// GetEntityContract is a free data retrieval call binding the contract method 0xd91ee2e0.
//
// Solidity: function getEntityContract(bytes32 _id) view returns(address)
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) GetEntityContract(_id [32]byte) (common.Address, error) {
	return _IEntitiesFactory.Contract.GetEntityContract(&_IEntitiesFactory.CallOpts, _id)
}

// GetEntityTypes is a free data retrieval call binding the contract method 0x6ec21da9.
//
// Solidity: function getEntityTypes(bytes32 _id) view returns(bool[])
func (_IEntitiesFactory *IEntitiesFactoryCaller) GetEntityTypes(opts *bind.CallOpts, _id [32]byte) ([]bool, error) {
	var (
		ret0 = new([]bool)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "getEntityTypes", _id)
	return *ret0, err
}

// GetEntityTypes is a free data retrieval call binding the contract method 0x6ec21da9.
//
// Solidity: function getEntityTypes(bytes32 _id) view returns(bool[])
func (_IEntitiesFactory *IEntitiesFactorySession) GetEntityTypes(_id [32]byte) ([]bool, error) {
	return _IEntitiesFactory.Contract.GetEntityTypes(&_IEntitiesFactory.CallOpts, _id)
}

// GetEntityTypes is a free data retrieval call binding the contract method 0x6ec21da9.
//
// Solidity: function getEntityTypes(bytes32 _id) view returns(bool[])
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) GetEntityTypes(_id [32]byte) ([]bool, error) {
	return _IEntitiesFactory.Contract.GetEntityTypes(&_IEntitiesFactory.CallOpts, _id)
}

// IsEntityType is a free data retrieval call binding the contract method 0xcef8b362.
//
// Solidity: function isEntityType(bytes32 _id, uint8 _type) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactoryCaller) IsEntityType(opts *bind.CallOpts, _id [32]byte, _type uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "isEntityType", _id, _type)
	return *ret0, err
}

// IsEntityType is a free data retrieval call binding the contract method 0xcef8b362.
//
// Solidity: function isEntityType(bytes32 _id, uint8 _type) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactorySession) IsEntityType(_id [32]byte, _type uint8) (bool, error) {
	return _IEntitiesFactory.Contract.IsEntityType(&_IEntitiesFactory.CallOpts, _id, _type)
}

// IsEntityType is a free data retrieval call binding the contract method 0xcef8b362.
//
// Solidity: function isEntityType(bytes32 _id, uint8 _type) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) IsEntityType(_id [32]byte, _type uint8) (bool, error) {
	return _IEntitiesFactory.Contract.IsEntityType(&_IEntitiesFactory.CallOpts, _id, _type)
}

// AddEntity is a paid mutator transaction binding the contract method 0x68a9f4e9.
//
// Solidity: function addEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntitiesFactory *IEntitiesFactoryTransactor) AddEntity(opts *bind.TransactOpts, _entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntitiesFactory.contract.Transact(opts, "addEntity", _entity, _type)
}

// AddEntity is a paid mutator transaction binding the contract method 0x68a9f4e9.
//
// Solidity: function addEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntitiesFactory *IEntitiesFactorySession) AddEntity(_entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.AddEntity(&_IEntitiesFactory.TransactOpts, _entity, _type)
}

// AddEntity is a paid mutator transaction binding the contract method 0x68a9f4e9.
//
// Solidity: function addEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntitiesFactory *IEntitiesFactoryTransactorSession) AddEntity(_entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.AddEntity(&_IEntitiesFactory.TransactOpts, _entity, _type)
}

// AmendEntity is a paid mutator transaction binding the contract method 0xa81eb545.
//
// Solidity: function amendEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntitiesFactory *IEntitiesFactoryTransactor) AmendEntity(opts *bind.TransactOpts, _entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntitiesFactory.contract.Transact(opts, "amendEntity", _entity, _type)
}

// AmendEntity is a paid mutator transaction binding the contract method 0xa81eb545.
//
// Solidity: function amendEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntitiesFactory *IEntitiesFactorySession) AmendEntity(_entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.AmendEntity(&_IEntitiesFactory.TransactOpts, _entity, _type)
}

// AmendEntity is a paid mutator transaction binding the contract method 0xa81eb545.
//
// Solidity: function amendEntity((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntitiesFactory *IEntitiesFactoryTransactorSession) AmendEntity(_entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.AmendEntity(&_IEntitiesFactory.TransactOpts, _entity, _type)
}

// IEntityABI is the input ABI used to generate the binding from.
const IEntityABI = "[{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEntityFuncSigs maps the 4-byte function signature to its string representation.
var IEntityFuncSigs = map[string]string{
	"e188130b": "addType(uint8)",
	"d1f439e5": "amend((bytes32,string,string,string),uint8)",
	"6d4ce63c": "get()",
	"b4579d60": "getTypes()",
	"222f74e7": "isType(uint8)",
}

// IEntity is an auto generated Go binding around an Ethereum contract.
type IEntity struct {
	IEntityCaller     // Read-only binding to the contract
	IEntityTransactor // Write-only binding to the contract
	IEntityFilterer   // Log filterer for contract events
}

// IEntityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEntityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEntityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEntityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEntitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEntitySession struct {
	Contract     *IEntity          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEntityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEntityCallerSession struct {
	Contract *IEntityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IEntityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEntityTransactorSession struct {
	Contract     *IEntityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IEntityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEntityRaw struct {
	Contract *IEntity // Generic contract binding to access the raw methods on
}

// IEntityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEntityCallerRaw struct {
	Contract *IEntityCaller // Generic read-only contract binding to access the raw methods on
}

// IEntityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEntityTransactorRaw struct {
	Contract *IEntityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEntity creates a new instance of IEntity, bound to a specific deployed contract.
func NewIEntity(address common.Address, backend bind.ContractBackend) (*IEntity, error) {
	contract, err := bindIEntity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEntity{IEntityCaller: IEntityCaller{contract: contract}, IEntityTransactor: IEntityTransactor{contract: contract}, IEntityFilterer: IEntityFilterer{contract: contract}}, nil
}

// NewIEntityCaller creates a new read-only instance of IEntity, bound to a specific deployed contract.
func NewIEntityCaller(address common.Address, caller bind.ContractCaller) (*IEntityCaller, error) {
	contract, err := bindIEntity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEntityCaller{contract: contract}, nil
}

// NewIEntityTransactor creates a new write-only instance of IEntity, bound to a specific deployed contract.
func NewIEntityTransactor(address common.Address, transactor bind.ContractTransactor) (*IEntityTransactor, error) {
	contract, err := bindIEntity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEntityTransactor{contract: contract}, nil
}

// NewIEntityFilterer creates a new log filterer instance of IEntity, bound to a specific deployed contract.
func NewIEntityFilterer(address common.Address, filterer bind.ContractFilterer) (*IEntityFilterer, error) {
	contract, err := bindIEntity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEntityFilterer{contract: contract}, nil
}

// bindIEntity binds a generic wrapper to an already deployed contract.
func bindIEntity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEntityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEntity *IEntityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEntity.Contract.IEntityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEntity *IEntityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntity.Contract.IEntityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEntity *IEntityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEntity.Contract.IEntityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEntity *IEntityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEntity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEntity *IEntityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEntity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEntity *IEntityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEntity.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((bytes32,string,string,string))
func (_IEntity *IEntityCaller) Get(opts *bind.CallOpts) (CreativeEntities, error) {
	var (
		ret0 = new(CreativeEntities)
	)
	out := ret0
	err := _IEntity.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((bytes32,string,string,string))
func (_IEntity *IEntitySession) Get() (CreativeEntities, error) {
	return _IEntity.Contract.Get(&_IEntity.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((bytes32,string,string,string))
func (_IEntity *IEntityCallerSession) Get() (CreativeEntities, error) {
	return _IEntity.Contract.Get(&_IEntity.CallOpts)
}

// GetTypes is a free data retrieval call binding the contract method 0xb4579d60.
//
// Solidity: function getTypes() view returns(bool[])
func (_IEntity *IEntityCaller) GetTypes(opts *bind.CallOpts) ([]bool, error) {
	var (
		ret0 = new([]bool)
	)
	out := ret0
	err := _IEntity.contract.Call(opts, out, "getTypes")
	return *ret0, err
}

// GetTypes is a free data retrieval call binding the contract method 0xb4579d60.
//
// Solidity: function getTypes() view returns(bool[])
func (_IEntity *IEntitySession) GetTypes() ([]bool, error) {
	return _IEntity.Contract.GetTypes(&_IEntity.CallOpts)
}

// GetTypes is a free data retrieval call binding the contract method 0xb4579d60.
//
// Solidity: function getTypes() view returns(bool[])
func (_IEntity *IEntityCallerSession) GetTypes() ([]bool, error) {
	return _IEntity.Contract.GetTypes(&_IEntity.CallOpts)
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_IEntity *IEntityCaller) IsType(opts *bind.CallOpts, _type uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEntity.contract.Call(opts, out, "isType", _type)
	return *ret0, err
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_IEntity *IEntitySession) IsType(_type uint8) (bool, error) {
	return _IEntity.Contract.IsType(&_IEntity.CallOpts, _type)
}

// IsType is a free data retrieval call binding the contract method 0x222f74e7.
//
// Solidity: function isType(uint8 _type) view returns(bool)
func (_IEntity *IEntityCallerSession) IsType(_type uint8) (bool, error) {
	return _IEntity.Contract.IsType(&_IEntity.CallOpts, _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_IEntity *IEntityTransactor) AddType(opts *bind.TransactOpts, _type uint8) (*types.Transaction, error) {
	return _IEntity.contract.Transact(opts, "addType", _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_IEntity *IEntitySession) AddType(_type uint8) (*types.Transaction, error) {
	return _IEntity.Contract.AddType(&_IEntity.TransactOpts, _type)
}

// AddType is a paid mutator transaction binding the contract method 0xe188130b.
//
// Solidity: function addType(uint8 _type) returns()
func (_IEntity *IEntityTransactorSession) AddType(_type uint8) (*types.Transaction, error) {
	return _IEntity.Contract.AddType(&_IEntity.TransactOpts, _type)
}

// Amend is a paid mutator transaction binding the contract method 0xd1f439e5.
//
// Solidity: function amend((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntity *IEntityTransactor) Amend(opts *bind.TransactOpts, _entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntity.contract.Transact(opts, "amend", _entity, _type)
}

// Amend is a paid mutator transaction binding the contract method 0xd1f439e5.
//
// Solidity: function amend((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntity *IEntitySession) Amend(_entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntity.Contract.Amend(&_IEntity.TransactOpts, _entity, _type)
}

// Amend is a paid mutator transaction binding the contract method 0xd1f439e5.
//
// Solidity: function amend((bytes32,string,string,string) _entity, uint8 _type) returns()
func (_IEntity *IEntityTransactorSession) Amend(_entity CreativeEntities, _type uint8) (*types.Transaction, error) {
	return _IEntity.Contract.Amend(&_IEntity.TransactOpts, _entity, _type)
}

// IFactoryABI is the input ABI used to generate the binding from.
const IFactoryABI = "[{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReferences\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IFactoryFuncSigs = map[string]string{
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"7a6337fa": "getReferences()",
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

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_IFactory *IFactoryCaller) GetReferences(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IFactory.contract.Call(opts, out, "getReferences")
	return *ret0, err
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_IFactory *IFactorySession) GetReferences() ([][32]byte, error) {
	return _IFactory.Contract.GetReferences(&_IFactory.CallOpts)
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_IFactory *IFactoryCallerSession) GetReferences() ([][32]byte, error) {
	return _IFactory.Contract.GetReferences(&_IFactory.CallOpts)
}

// IterableDataABI is the input ABI used to generate the binding from.
const IterableDataABI = "[]"

// IterableDataBin is the compiled bytecode used for deploying new contracts.
var IterableDataBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201e16cb4d554323ed35fc71a921b5a636734c1a9c7e850c9663b205061f61c76264736f6c63430006080033"

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
