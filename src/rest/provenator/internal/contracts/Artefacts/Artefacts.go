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
var ArtefactNodeBin = "0x60806040523480156200001157600080fd5b50604051620014a8380380620014a8833981016040819052620000349162000408565b604081015160001a60f81b6001600160f81b031916158015906200006b5750606081015160001a60f81b6001600160f81b03191615155b80156200008557506000815160158111156200008357fe5b115b80156200009f57506015815160158111156200009d57fe5b105b8015620000bc5750601481602001516014811115620000ba57fe5b105b8015620000dc5750608081015160001a60f81b6001600160f81b03191615155b8015620000ee575060008160e0015151115b8015620001015750600081610100015151115b6200010b57600080fd5b805160008054839290829060ff191660018360158111156200012957fe5b021790555060208201518154829061ff0019166101008360148111156200014c57fe5b021790555060408201516001820155606082015160028201556080820151600382015560a0820151600482015560c0820151805162000196916005840191602090910190620002ae565b5060e08201518051620001b4916006840191602090910190620002ae565b506101008201518051620001d3916007840191602090910190620002ae565b5050506060810151620001ef906001600160e01b03620001f616565b5062000545565b6200020a816001600160e01b036200024816565b6200024557600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018190555b50565b600081811a60f81b6001600160f81b0319166200026457600080fd5b6000805b600854811015620002a75783600882815481106200028257fe5b906000526020600020015414156200029e5760019150620002a7565b60010162000268565b5092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002f157805160ff191683800117855562000321565b8280016001018555821562000321579182015b828111156200032157825182559160200191906001019062000304565b506200032f92915062000333565b5090565b6200035091905b808211156200032f57600081556001016200033a565b90565b8051601581106200036357600080fd5b92915050565b8051601681106200036357600080fd5b600082601f8301126200038a578081fd5b81516001600160401b03811115620003a0578182fd5b6020620003b6601f8301601f191682016200051e565b92508183528481838601011115620003cd57600080fd5b60005b82811015620003ed578481018201518482018301528101620003d0565b82811115620003ff5760008284860101525b50505092915050565b6000602082840312156200041a578081fd5b81516001600160401b038082111562000431578283fd5b61012091840180860383131562000446578384fd5b62000451836200051e565b6200045d878362000369565b81526200046e876020840162000353565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c0820151935082841115620004af578485fd5b620004bd8785840162000379565b60c082015260e0820151935082841115620004d6578485fd5b620004e48785840162000379565b60e082015261010093508382015183811115620004ff578586fd5b6200050d8882850162000379565b948201949094529695505050505050565b6040518181016001600160401b03811182821017156200053d57600080fd5b604052919050565b610f5380620005556000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637fc622f41161008c578063c28b16dc11610066578063c28b16dc146101be578063c3a30a6f146101d1578063d13c8f23146101e4578063d6d54164146101f7576100ea565b80637fc622f4146101905780639673484e14610198578063a9c56065146101ab576100ea565b8063476caa00116100c8578063476caa001461014057806359c0161c146101535780636c6071aa146101665780636d4ce63c1461017b576100ea565b8063240b5600146100ef5780632424dccc1461011857806333966d031461012d575b600080fd5b6101026100fd366004610c76565b6101ff565b60405161010f9190610e3a565b60405180910390f35b61012b610126366004610c76565b61025f565b005b61012b61013b366004610c76565b610333565b61010261014e366004610c76565b6103e8565b610102610161366004610c76565b610441565b61016e61049a565b60405161010f9190610df6565b6101836104f3565b60405161010f9190610e45565b61016e610749565b61012b6101a6366004610c8e565b61079f565b61012b6101b9366004610c76565b610931565b61012b6101cc366004610c76565b610973565b61012b6101df366004610c76565b6109b5565b61012b6101f2366004610c76565b6109f7565b61016e610aac565b600081811a60f81b6001600160f81b03191661021a57600080fd5b6000805b600a548110156102585783600a828154811061023657fe5b906000526020600020015414156102505760019150610258565b60010161021e565b5092915050565b8060001a60f81b6001600160f81b03191661027957600080fd5b60005b60095481101561032f57816009828154811061029457fe5b90600052602060002001541415610327576009546000190181101561030157805b600954600019018110156102ff57600981600101815481106102d357fe5b9060005260206000200154600982815481106102eb57fe5b6000918252602090912001556001016102b5565b505b600980548061030c57fe5b6001900381819060005260206000200160009055905561032f565b60010161027c565b5050565b8060001a60f81b6001600160f81b03191661034d57600080fd5b60005b60085481101561032f57816008828154811061036857fe5b906000526020600020015414156103e057600854600019018110156103d557805b600854600019018110156103d357600881600101815481106103a757fe5b9060005260206000200154600882815481106103bf57fe5b600091825260209091200155600101610389565b505b600880548061030c57fe5b600101610350565b600081811a60f81b6001600160f81b03191661040357600080fd5b6000805b60095481101561025857836009828154811061041f57fe5b906000526020600020015414156104395760019150610258565b600101610407565b600081811a60f81b6001600160f81b03191661045c57600080fd5b6000805b60085481101561025857836008828154811061047857fe5b906000526020600020015414156104925760019150610258565b600101610460565b6060600a8054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020905b8154815260200190600101908083116104d4575b505050505090505b90565b6104fb610b02565b60408051610120810190915260008054829060ff16601581111561051b57fe5b601581111561052657fe5b81528154602090910190610100900460ff16601481111561054357fe5b601481111561054e57fe5b815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106135780601f106105e857610100808354040283529160200191610613565b820191906000526020600020905b8154815290600101906020018083116105f657829003601f168201915b505050918352505060068201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b505050918352505060078201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561073b5780601f106107105761010080835404028352916020019161073b565b820191906000526020600020905b81548152906001019060200180831161071e57829003601f168201915b505050505081525050905090565b606060088054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020908154815260200190600101908083116104d4575050505050905090565b60408101516001541480156107c75750606081015160001a60f81b6001600160f81b03191615155b80156107df57506000815160158111156107dd57fe5b115b80156107f757506015815160158111156107f557fe5b105b8015610812575060148160200151601481111561081057fe5b105b80156108315750608081015160001a60f81b6001600160f81b03191615155b8015610842575060008160e0015151115b80156108545750600081610100015151115b61085d57600080fd5b805160008054839290829060ff1916600183601581111561087a57fe5b021790555060208201518154829061ff00191661010083601481111561089c57fe5b021790555060408201516001820155606082015160028201556080820151600382015560a0820151600482015560c082015180516108e4916005840191602090910190610b51565b5060e08201518051610900916006840191602090910190610b51565b50610100820151805161091d916007840191602090910190610b51565b5090505061092e8160600151610931565b50565b61093a81610441565b61092e57600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b61097c816101ff565b61092e57600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80155565b6109be816103e8565b61092e57600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0155565b8060001a60f81b6001600160f81b031916610a1157600080fd5b60005b600a5481101561032f5781600a8281548110610a2c57fe5b90600052602060002001541415610aa457600a5460001901811015610a9957805b600a5460001901811015610a9757600a8160010181548110610a6b57fe5b9060005260206000200154600a8281548110610a8357fe5b600091825260209091200155600101610a4d565b505b600a80548061030c57fe5b600101610a14565b606060098054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020908154815260200190600101908083116104d4575050505050905090565b604080516101208101909152806000815260200160008152600060208201819052604082018190526060808301829052608083019190915260a0820181905260c0820181905260e09091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b9257805160ff1916838001178555610bbf565b82800160010185558215610bbf579182015b82811115610bbf578251825591602001919060010190610ba4565b50610bcb929150610bcf565b5090565b6104f091905b80821115610bcb5760008155600101610bd5565b803560158110610bf857600080fd5b92915050565b803560168110610bf857600080fd5b600082601f830112610c1d578081fd5b813567ffffffffffffffff811115610c33578182fd5b610c46601f8201601f1916602001610ef6565b9150808252836020828501011115610c5d57600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610c87578081fd5b5035919050565b600060208284031215610c9f578081fd5b813567ffffffffffffffff80821115610cb6578283fd5b610120918401808603831315610cca578384fd5b610cd383610ef6565b610cdd8783610bfe565b8152610cec8760208401610be9565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c0820135935082841115610d2c578485fd5b610d3887858401610c0d565b60c082015260e0820135935082841115610d50578485fd5b610d5c87858401610c0d565b60e082015261010093508382013583811115610d76578586fd5b610d8288828501610c0d565b948201949094529695505050505050565b60158110610d9d57fe5b9052565b60168110610d9d57fe5b60008151808452815b81811015610dd057602081850181015186830182015201610db4565b81811115610de15782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610e2e57835183529284019291840191600101610e12565b50909695505050505050565b901515815260200190565b600060208252610e59602083018451610da1565b6020830151610e6b6040840182610d93565b506040830151606083015260608301516080830152608083015160a083015260a083015160c083015260c08301516101208060e0850152610eb0610140850183610dab565b60e08601519250601f19610100818784030181880152610ed08386610dab565b818901519550828882030185890152610ee98187610dab565b9998505050505050505050565b60405181810167ffffffffffffffff81118282101715610f1557600080fd5b60405291905056fea26469706673582212202e89a15c9665622f4bc5622a5fee231740d5d45421e4097a87a89e68e81842b564736f6c63430006080033"

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
var ArtefactsBin = "0x608060405234801561001057600080fd5b5060006002556127ce806100256000396000f3fe60806040523480156200001157600080fd5b50600436106200013c5760003560e01c80635ffd09d811620000bd578063bc599341116200007b578063bc59934114620002a8578063e17989ea14620002bf578063e229e17314620002d6578063ef90d9c914620002ed578063f6a509651462000304576200013c565b80635ffd09d8146200023157806367e0badb14620002485780637a6337fa146200026157806386105c65146200026b578063b2e459cf1462000282576200013c565b8063426aa803116200010b578063426aa80314620001b75780634d2522a214620001dd5780634e542a6e1462000203578063521ffdce146200014157806355884f0f146200021a576200013c565b80630c99420f14620001415780630f1ee4a2146200015a5780631d339e241462000171578063302665371462000188575b600080fd5b620001586200015236600462000e83565b6200031b565b005b620001586200016b36600462000e83565b620003bd565b620001586200018236600462000ea5565b62000426565b6200019f6200019936600462000e6a565b620004b4565b604051620001ae919062001178565b60405180910390f35b620001ce620001c836600462000e6a565b6200057a565b604051620001ae91906200111e565b620001f4620001ee36600462000e6a565b62000632565b604051620001ae91906200110a565b620001586200021436600462000ea5565b62000676565b620001586200022b36600462000e83565b6200071d565b620001ce6200024236600462000e6a565b62000786565b62000252620007ff565b604051620001ae91906200116f565b620001ce62000805565b620001ce6200027c36600462000e6a565b620008a9565b620002996200029336600462000e83565b62000922565b604051620001ae919062001164565b62000252620002b936600462000e6a565b620009e9565b62000158620002d036600462000e83565b62000a23565b62000158620002e736600462000e83565b62000a8c565b62000299620002fe36600462000e83565b62000af5565b620002996200031536600462000e83565b62000b5d565b6000828152602081905260409020600101546001600160a01b03166200034057600080fd5b60008281526020819052604090819020600101549051630909373360e21b81526001600160a01b03909116908190632424dccc90620003849085906004016200116f565b600060405180830381600087803b1580156200039f57600080fd5b505af1158015620003b4573d6000803e3d6000fd5b50505050505050565b6000828152602081905260409020600101546001600160a01b0316620003e257600080fd5b6000828152602081905260409081902060010154905163c3a30a6f60e01b81526001600160a01b0390911690819063c3a30a6f90620003849085906004016200116f565b60408101516200043f9060009063ffffffff62000bc516565b156200045657620004508162000676565b620004b1565b600081604051620004679062000c6e565b62000473919062001178565b604051809103906000f08015801562000490573d6000803e3d6000fd5b506040830151909150620004ae906000908363ffffffff62000bda16565b50505b50565b620004be62000c7c565b6000828152602081905260409020600101546001600160a01b0316620004e357600080fd5b600082815260208190526040808220600101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b1580156200053457600080fd5b505afa15801562000549573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000573919081019062000fbc565b9392505050565b6000818152602081905260409020600101546060906001600160a01b0316620005a257600080fd5b600082815260208190526040808220600101548151631ff188bd60e21b815291516001600160a01b03909116928392637fc622f49260048083019392829003018186803b158015620005f357600080fd5b505afa15801562000608573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000573919081019062000dac565b6000818152602081905260408120600101546001600160a01b03166200065757600080fd5b506000908152602081905260409020600101546001600160a01b031690565b6040808201516000908152602081905220600101546001600160a01b03166200069e57600080fd5b60408082015160009081526020819052819020600101549051634b39a42760e11b81526001600160a01b03909116908190639673484e90620006e590859060040162001178565b600060405180830381600087803b1580156200070057600080fd5b505af115801562000715573d6000803e3d6000fd5b505050505050565b6000828152602081905260409020600101546001600160a01b03166200074257600080fd5b6000828152602081905260409081902060010154905163a9c5606560e01b81526001600160a01b0390911690819063a9c5606590620003849085906004016200116f565b6000818152602081905260409020600101546060906001600160a01b0316620007ae57600080fd5b60008281526020819052604080822060010154815163363038d560e11b815291516001600160a01b03909116928392636c6071aa9260048083019392829003018186803b158015620005f357600080fd5b60025490565b60608060006002015467ffffffffffffffff811180156200082557600080fd5b5060405190808252806020026020018201604052801562000850578160200160208202803683370190505b50905060005b600254811015620008a35760018054829081106200087057fe5b9060005260206000209060020201600001548282815181106200088f57fe5b602090810291909101015260010162000856565b50905090565b6000818152602081905260409020600101546060906001600160a01b0316620008d157600080fd5b6000828152602081905260408082206001015481516335b5505960e21b815291516001600160a01b0390911692839263d6d541649260048083019392829003018186803b158015620005f357600080fd5b6000828152602081905260408120600101546001600160a01b03166200094757600080fd5b60008381526020819052604090819020600101549051631670058760e21b81526001600160a01b039091169081906359c0161c906200098b9086906004016200116f565b60206040518083038186803b158015620009a457600080fd5b505afa158015620009b9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620009df919062000e48565b9150505b92915050565b6002546000908210620009fb57600080fd5b600180548390811062000a0a57fe5b9060005260206000209060020201600001549050919050565b6000828152602081905260409020600101546001600160a01b031662000a4857600080fd5b600082815260208190526040908190206001015490516330a2c5b760e21b81526001600160a01b0390911690819063c28b16dc90620003849085906004016200116f565b6000828152602081905260409020600101546001600160a01b031662000ab157600080fd5b600082815260208190526040908190206001015490516333966d0360e01b81526001600160a01b039091169081906333966d0390620003849085906004016200116f565b6000828152602081905260408120600101546001600160a01b031662000b1a57600080fd5b60008381526020819052604090819020600101549051621205ab60e91b81526001600160a01b0390911690819063240b5600906200098b9086906004016200116f565b6000828152602081905260408120600101546001600160a01b031662000b8257600080fd5b600083815260208190526040908190206001015490516223b65560e91b81526001600160a01b0390911690819063476caa00906200098b9086906004016200116f565b60009081526020919091526040902054151590565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000c1a57600191505062000573565b506001808501805491820180825560008681526020889052604090208190558591908390811062000c4757fe5b60009182526020822060029182020192909255908601805460010190559150620005739050565b6114a880620012f183390190565b604080516101208101909152806000815260200160008152600060208201819052604082018190526060808301829052608083019190915260a0820181905260c0820181905260e09091015290565b8035620009e381620012d4565b8051620009e381620012d4565b8035620009e381620012e2565b8051620009e381620012e2565b600082601f83011262000d10578081fd5b813562000d2762000d21826200127c565b62001233565b915080825283602082850101111562000d3f57600080fd5b8060208401602084013760009082016020015292915050565b600082601f83011262000d69578081fd5b815162000d7a62000d21826200127c565b915080825283602082850101111562000d9257600080fd5b62000da5816020840160208601620012a1565b5092915050565b6000602080838503121562000dbf578182fd5b825167ffffffffffffffff81111562000dd6578283fd5b80840185601f82011262000de8578384fd5b8051915062000dfb62000d21836200125b565b828152838101908285018585028401860189101562000e18578687fd5b8693505b8484101562000e3c57805183526001939093019291850191850162000e1c565b50979650505050505050565b60006020828403121562000e5a578081fd5b8151801515811462000573578182fd5b60006020828403121562000e7c578081fd5b5035919050565b6000806040838503121562000e96578081fd5b50508035926020909101359150565b60006020828403121562000eb7578081fd5b813567ffffffffffffffff8082111562000ecf578283fd5b61012091840180860383131562000ee4578384fd5b62000eef8362001233565b62000efb878362000ce5565b815262000f0c876020840162000ccb565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c082013593508284111562000f4d578485fd5b62000f5b8785840162000cff565b60c082015260e082013593508284111562000f74578485fd5b62000f828785840162000cff565b60e08201526101009350838201358381111562000f9d578586fd5b62000fab8882850162000cff565b948201949094529695505050505050565b60006020828403121562000fce578081fd5b815167ffffffffffffffff8082111562000fe6578283fd5b61012091840180860383131562000ffb578384fd5b620010068362001233565b62001012878362000cf2565b815262001023876020840162000cd8565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c082015193508284111562001064578485fd5b620010728785840162000d58565b60c082015260e08201519350828411156200108b578485fd5b620010998785840162000d58565b60e082015261010093508382015183811115620010b4578586fd5b62000fab8882850162000d58565b60158110620010cd57fe5b9052565b60168110620010cd57fe5b60008151808452620010f6816020860160208601620012a1565b601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b8181101562001158578351835292840192918401916001016200113a565b50909695505050505050565b901515815260200190565b90815260200190565b6000602082526200118e602083018451620010d1565b6020830151620011a26040840182620010c2565b506040830151606083015260608301516080830152608083015160a083015260a083015160c083015260c08301516101208060e0850152620011e9610140850183620010dc565b60e08601519250601f196101008187840301818801526200120b8386620010dc565b818901519550828882030185890152620012268187620010dc565b9998505050505050505050565b60405181810167ffffffffffffffff811182821017156200125357600080fd5b604052919050565b600067ffffffffffffffff82111562001272578081fd5b5060209081020190565b600067ffffffffffffffff82111562001293578081fd5b50601f01601f191660200190565b60005b83811015620012be578181015183820152602001620012a4565b83811115620012ce576000848401525b50505050565b60158110620004b157600080fd5b60168110620004b157600080fdfe60806040523480156200001157600080fd5b50604051620014a8380380620014a8833981016040819052620000349162000408565b604081015160001a60f81b6001600160f81b031916158015906200006b5750606081015160001a60f81b6001600160f81b03191615155b80156200008557506000815160158111156200008357fe5b115b80156200009f57506015815160158111156200009d57fe5b105b8015620000bc5750601481602001516014811115620000ba57fe5b105b8015620000dc5750608081015160001a60f81b6001600160f81b03191615155b8015620000ee575060008160e0015151115b8015620001015750600081610100015151115b6200010b57600080fd5b805160008054839290829060ff191660018360158111156200012957fe5b021790555060208201518154829061ff0019166101008360148111156200014c57fe5b021790555060408201516001820155606082015160028201556080820151600382015560a0820151600482015560c0820151805162000196916005840191602090910190620002ae565b5060e08201518051620001b4916006840191602090910190620002ae565b506101008201518051620001d3916007840191602090910190620002ae565b5050506060810151620001ef906001600160e01b03620001f616565b5062000545565b6200020a816001600160e01b036200024816565b6200024557600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018190555b50565b600081811a60f81b6001600160f81b0319166200026457600080fd5b6000805b600854811015620002a75783600882815481106200028257fe5b906000526020600020015414156200029e5760019150620002a7565b60010162000268565b5092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002f157805160ff191683800117855562000321565b8280016001018555821562000321579182015b828111156200032157825182559160200191906001019062000304565b506200032f92915062000333565b5090565b6200035091905b808211156200032f57600081556001016200033a565b90565b8051601581106200036357600080fd5b92915050565b8051601681106200036357600080fd5b600082601f8301126200038a578081fd5b81516001600160401b03811115620003a0578182fd5b6020620003b6601f8301601f191682016200051e565b92508183528481838601011115620003cd57600080fd5b60005b82811015620003ed578481018201518482018301528101620003d0565b82811115620003ff5760008284860101525b50505092915050565b6000602082840312156200041a578081fd5b81516001600160401b038082111562000431578283fd5b61012091840180860383131562000446578384fd5b62000451836200051e565b6200045d878362000369565b81526200046e876020840162000353565b602082015260408201516040820152606082015160608201526080820151608082015260a082015160a082015260c0820151935082841115620004af578485fd5b620004bd8785840162000379565b60c082015260e0820151935082841115620004d6578485fd5b620004e48785840162000379565b60e082015261010093508382015183811115620004ff578586fd5b6200050d8882850162000379565b948201949094529695505050505050565b6040518181016001600160401b03811182821017156200053d57600080fd5b604052919050565b610f5380620005556000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637fc622f41161008c578063c28b16dc11610066578063c28b16dc146101be578063c3a30a6f146101d1578063d13c8f23146101e4578063d6d54164146101f7576100ea565b80637fc622f4146101905780639673484e14610198578063a9c56065146101ab576100ea565b8063476caa00116100c8578063476caa001461014057806359c0161c146101535780636c6071aa146101665780636d4ce63c1461017b576100ea565b8063240b5600146100ef5780632424dccc1461011857806333966d031461012d575b600080fd5b6101026100fd366004610c76565b6101ff565b60405161010f9190610e3a565b60405180910390f35b61012b610126366004610c76565b61025f565b005b61012b61013b366004610c76565b610333565b61010261014e366004610c76565b6103e8565b610102610161366004610c76565b610441565b61016e61049a565b60405161010f9190610df6565b6101836104f3565b60405161010f9190610e45565b61016e610749565b61012b6101a6366004610c8e565b61079f565b61012b6101b9366004610c76565b610931565b61012b6101cc366004610c76565b610973565b61012b6101df366004610c76565b6109b5565b61012b6101f2366004610c76565b6109f7565b61016e610aac565b600081811a60f81b6001600160f81b03191661021a57600080fd5b6000805b600a548110156102585783600a828154811061023657fe5b906000526020600020015414156102505760019150610258565b60010161021e565b5092915050565b8060001a60f81b6001600160f81b03191661027957600080fd5b60005b60095481101561032f57816009828154811061029457fe5b90600052602060002001541415610327576009546000190181101561030157805b600954600019018110156102ff57600981600101815481106102d357fe5b9060005260206000200154600982815481106102eb57fe5b6000918252602090912001556001016102b5565b505b600980548061030c57fe5b6001900381819060005260206000200160009055905561032f565b60010161027c565b5050565b8060001a60f81b6001600160f81b03191661034d57600080fd5b60005b60085481101561032f57816008828154811061036857fe5b906000526020600020015414156103e057600854600019018110156103d557805b600854600019018110156103d357600881600101815481106103a757fe5b9060005260206000200154600882815481106103bf57fe5b600091825260209091200155600101610389565b505b600880548061030c57fe5b600101610350565b600081811a60f81b6001600160f81b03191661040357600080fd5b6000805b60095481101561025857836009828154811061041f57fe5b906000526020600020015414156104395760019150610258565b600101610407565b600081811a60f81b6001600160f81b03191661045c57600080fd5b6000805b60085481101561025857836008828154811061047857fe5b906000526020600020015414156104925760019150610258565b600101610460565b6060600a8054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020905b8154815260200190600101908083116104d4575b505050505090505b90565b6104fb610b02565b60408051610120810190915260008054829060ff16601581111561051b57fe5b601581111561052657fe5b81528154602090910190610100900460ff16601481111561054357fe5b601481111561054e57fe5b815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106135780601f106105e857610100808354040283529160200191610613565b820191906000526020600020905b8154815290600101906020018083116105f657829003601f168201915b505050918352505060068201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b505050918352505060078201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561073b5780601f106107105761010080835404028352916020019161073b565b820191906000526020600020905b81548152906001019060200180831161071e57829003601f168201915b505050505081525050905090565b606060088054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020908154815260200190600101908083116104d4575050505050905090565b60408101516001541480156107c75750606081015160001a60f81b6001600160f81b03191615155b80156107df57506000815160158111156107dd57fe5b115b80156107f757506015815160158111156107f557fe5b105b8015610812575060148160200151601481111561081057fe5b105b80156108315750608081015160001a60f81b6001600160f81b03191615155b8015610842575060008160e0015151115b80156108545750600081610100015151115b61085d57600080fd5b805160008054839290829060ff1916600183601581111561087a57fe5b021790555060208201518154829061ff00191661010083601481111561089c57fe5b021790555060408201516001820155606082015160028201556080820151600382015560a0820151600482015560c082015180516108e4916005840191602090910190610b51565b5060e08201518051610900916006840191602090910190610b51565b50610100820151805161091d916007840191602090910190610b51565b5090505061092e8160600151610931565b50565b61093a81610441565b61092e57600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b61097c816101ff565b61092e57600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80155565b6109be816103e8565b61092e57600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0155565b8060001a60f81b6001600160f81b031916610a1157600080fd5b60005b600a5481101561032f5781600a8281548110610a2c57fe5b90600052602060002001541415610aa457600a5460001901811015610a9957805b600a5460001901811015610a9757600a8160010181548110610a6b57fe5b9060005260206000200154600a8281548110610a8357fe5b600091825260209091200155600101610a4d565b505b600a80548061030c57fe5b600101610a14565b606060098054806020026020016040519081016040528092919081815260200182805480156104e857602002820191906000526020600020908154815260200190600101908083116104d4575050505050905090565b604080516101208101909152806000815260200160008152600060208201819052604082018190526060808301829052608083019190915260a0820181905260c0820181905260e09091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b9257805160ff1916838001178555610bbf565b82800160010185558215610bbf579182015b82811115610bbf578251825591602001919060010190610ba4565b50610bcb929150610bcf565b5090565b6104f091905b80821115610bcb5760008155600101610bd5565b803560158110610bf857600080fd5b92915050565b803560168110610bf857600080fd5b600082601f830112610c1d578081fd5b813567ffffffffffffffff811115610c33578182fd5b610c46601f8201601f1916602001610ef6565b9150808252836020828501011115610c5d57600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610c87578081fd5b5035919050565b600060208284031215610c9f578081fd5b813567ffffffffffffffff80821115610cb6578283fd5b610120918401808603831315610cca578384fd5b610cd383610ef6565b610cdd8783610bfe565b8152610cec8760208401610be9565b602082015260408201356040820152606082013560608201526080820135608082015260a082013560a082015260c0820135935082841115610d2c578485fd5b610d3887858401610c0d565b60c082015260e0820135935082841115610d50578485fd5b610d5c87858401610c0d565b60e082015261010093508382013583811115610d76578586fd5b610d8288828501610c0d565b948201949094529695505050505050565b60158110610d9d57fe5b9052565b60168110610d9d57fe5b60008151808452815b81811015610dd057602081850181015186830182015201610db4565b81811115610de15782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610e2e57835183529284019291840191600101610e12565b50909695505050505050565b901515815260200190565b600060208252610e59602083018451610da1565b6020830151610e6b6040840182610d93565b506040830151606083015260608301516080830152608083015160a083015260a083015160c083015260c08301516101208060e0850152610eb0610140850183610dab565b60e08601519250601f19610100818784030181880152610ed08386610dab565b818901519550828882030185890152610ee98187610dab565b9998505050505050505050565b60405181810167ffffffffffffffff81118282101715610f1557600080fd5b60405291905056fea26469706673582212202e89a15c9665622f4bc5622a5fee231740d5d45421e4097a87a89e68e81842b564736f6c63430006080033a26469706673582212203df087dbc421c8a4052722d11d453801784e2a9e30132e13f0eb773fc2db44b164736f6c63430006080033"

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
const EntitiesABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"name\":\"addEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedId\",\"type\":\"bytes32\"}],\"name\":\"addEntityRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amendEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntity\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityRelations\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReferences\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedEntityId\",\"type\":\"bytes32\"}],\"name\":\"isEntityRelation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isEntityType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntitiesFuncSigs maps the 4-byte function signature to its string representation.
var EntitiesFuncSigs = map[string]string{
	"68a9f4e9": "addEntity((bytes32,string,string,string),uint8)",
	"6c3cf7f7": "addEntityRelation(bytes32,bytes32)",
	"a81eb545": "amendEntity((bytes32,string,string,string),uint8)",
	"53b66f36": "getEntity(bytes32)",
	"d91ee2e0": "getEntityContract(bytes32)",
	"b24aa090": "getEntityRelations(bytes32)",
	"6ec21da9": "getEntityTypes(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"7a6337fa": "getReferences()",
	"941c0169": "isEntityRelation(bytes32,bytes32)",
	"cef8b362": "isEntityType(bytes32,uint8)",
}

// EntitiesBin is the compiled bytecode used for deploying new contracts.
var EntitiesBin = "0x608060405234801561001057600080fd5b506000600255611fb9806100256000396000f3fe60806040523480156200001157600080fd5b5060043610620000c45760003560e01c8063941c0169116200007b578063941c01691462000180578063a81eb54514620001a6578063b24aa09014620001bd578063bc59934114620001d4578063cef8b36214620001eb578063d91ee2e0146200020257620000c4565b806353b66f3614620000c957806367e0badb14620000f857806368a9f4e914620001115780636c3cf7f7146200012a5780636ec21da914620001415780637a6337fa1462000167575b600080fd5b620000e0620000da36600462000b83565b62000228565b604051620000ef919062000ef5565b60405180910390f35b62000102620002ee565b604051620000ef919062000edc565b620001286200012236600462000cbd565b620002f4565b005b620001286200013b36600462000b9c565b62000380565b620001586200015236600462000b83565b6200044a565b604051620000ef919062000e4f565b6200017162000502565b604051620000ef919062000e97565b620001976200019136600462000b9c565b620005a6565b604051620000ef919062000ed1565b62000128620001b736600462000cbd565b6200066d565b62000171620001ce36600462000b83565b620006dc565b62000102620001e536600462000b83565b62000794565b62000197620001fc36600462000bbe565b620007ce565b620002196200021336600462000b83565b62000837565b604051620000ef919062000e3b565b6200023262000924565b6000828152602081905260409020600101546001600160a01b03166200025757600080fd5b600082815260208190526040808220600101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b158015620002a857600080fd5b505afa158015620002bd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620002e7919081019062000bf2565b9392505050565b60025490565b81516200030a9060009063ffffffff6200087b16565b1562000322576200031c82826200066d565b6200037c565b6000828260405162000334906200094f565b6200034192919062000f0a565b604051809103906000f0801580156200035e573d6000803e3d6000fd5b50835190915062000379906000908363ffffffff6200089016565b50505b5050565b6000828152602081905260409020600101546001600160a01b031615801590620003c357506000818152602081905260409020600101546001600160a01b031615155b620003cd57600080fd5b600082815260208190526040908190206001015490516351366bd360e11b81526001600160a01b0390911690819063a26cd7a6906200041190859060040162000edc565b600060405180830381600087803b1580156200042c57600080fd5b505af115801562000441573d6000803e3d6000fd5b50505050505050565b6000818152602081905260409020600101546060906001600160a01b03166200047257600080fd5b6000828152602081905260408082206001015481516305a2bceb60e51b815291516001600160a01b0390911692839263b4579d609260048083019392829003018186803b158015620004c357600080fd5b505afa158015620004d8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620002e7919081019062000a2b565b60608060006002015467ffffffffffffffff811180156200052257600080fd5b506040519080825280602002602001820160405280156200054d578160200160208202803683370190505b50905060005b600254811015620005a05760018054829081106200056d57fe5b9060005260206000209060020201600001548282815181106200058c57fe5b602090810291909101015260010162000553565b50905090565b6000828152602081905260408120600101546001600160a01b0316620005cb57600080fd5b60008381526020819052604090819020600101549051636d4a255d60e01b81526001600160a01b03909116908190636d4a255d906200060f90869060040162000edc565b60206040518083038186803b1580156200062857600080fd5b505afa1580156200063d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000663919062000b61565b9150505b92915050565b81516000908152602081905260409020600101546001600160a01b03166200069457600080fd5b81516000908152602081905260409081902060010154905163d1f439e560e01b81526001600160a01b0390911690819063d1f439e59062000411908690869060040162000f0a565b6000818152602081905260409020600101546060906001600160a01b03166200070457600080fd5b60008281526020819052604080822060010154815163616e1f1160e11b815291516001600160a01b0390911692839263c2dc3e229260048083019392829003018186803b1580156200075557600080fd5b505afa1580156200076a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620002e7919081019062000ad1565b6002546000908210620007a657600080fd5b6001805483908110620007b557fe5b9060005260206000209060020201600001549050919050565b6000828152602081905260408120600101546001600160a01b0316620007f357600080fd5b6000838152602081905260409081902060010154905163222f74e760e01b81526001600160a01b0390911690819063222f74e7906200060f90869060040162000ee5565b6000818152602081905260408120600101546001600160a01b03166200085c57600080fd5b506000908152602081905260409020600101546001600160a01b031690565b60009081526020919091526040902054151590565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b0385161790558015620008d0576001915050620002e7565b5060018085018054918201808255600086815260208890526040902081905585919083908110620008fd57fe5b60009182526020822060029182020192909255908601805460010190559150620002e79050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b610fb68062000fce83390190565b805180151581146200066757600080fd5b8035600581106200066757600080fd5b600082601f8301126200098f578081fd5b8135620009a6620009a08262000f79565b62000f30565b9150808252836020828501011115620009be57600080fd5b8060208401602084013760009082016020015292915050565b600082601f830112620009e8578081fd5b8151620009f9620009a08262000f79565b915080825283602082850101111562000a1157600080fd5b62000a2481602084016020860162000f9e565b5092915050565b6000602080838503121562000a3e578182fd5b825167ffffffffffffffff81111562000a55578283fd5b80840185601f82011262000a67578384fd5b8051915062000a7a620009a08362000f58565b828152838101908285018585028401860189101562000a97578687fd5b8693505b8484101562000ac55762000ab089826200095d565b83526001939093019291850191850162000a9b565b50979650505050505050565b6000602080838503121562000ae4578182fd5b825167ffffffffffffffff81111562000afb578283fd5b80840185601f82011262000b0d578384fd5b8051915062000b20620009a08362000f58565b828152838101908285018585028401860189101562000b3d578687fd5b8693505b8484101562000ac557805183526001939093019291850191850162000b41565b60006020828403121562000b73578081fd5b81518015158114620002e7578182fd5b60006020828403121562000b95578081fd5b5035919050565b6000806040838503121562000baf578081fd5b50508035926020909101359150565b6000806040838503121562000bd1578182fd5b8235915060208301356005811062000be7578182fd5b809150509250929050565b60006020828403121562000c04578081fd5b815167ffffffffffffffff8082111562000c1c578283fd5b8184016080818703121562000c2f578384fd5b62000c3b608062000f30565b92508051835260208101518281111562000c53578485fd5b62000c6187828401620009d7565b60208501525060408101518281111562000c79578485fd5b62000c8787828401620009d7565b60408501525060608101518281111562000c9f578485fd5b62000cad87828401620009d7565b6060850152509195945050505050565b6000806040838503121562000cd0578182fd5b823567ffffffffffffffff8082111562000ce8578384fd5b8185016080818803121562000cfb578485fd5b62000d07608062000f30565b92508035835260208101358281111562000d1f578586fd5b62000d2d888284016200097e565b60208501525060408101358281111562000d45578586fd5b62000d53888284016200097e565b60408501525060608101358281111562000d6b578586fd5b62000d79888284016200097e565b60608501525050508092505062000d9484602085016200096e565b90509250929050565b6005811062000da857fe5b9052565b6000815180845262000dc681602086016020860162000f9e565b601f01601f19169290920160200192915050565b60008151835260208201516080602085015262000dfb608085018262000dac565b60408401519150848103604086015262000e16818362000dac565b60608501519250858103606087015262000e31818462000dac565b9695505050505050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b8181101562000e8b57835115158352928401929184019160010162000e6b565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101562000e8b5783518352928401929184019160010162000eb3565b901515815260200190565b90815260200190565b6020810162000667828462000d9d565b600060208252620002e7602083018462000dda565b60006040825262000f1f604083018562000dda565b9050620002e7602083018462000d9d565b60405181810167ffffffffffffffff8111828210171562000f5057600080fd5b604052919050565b600067ffffffffffffffff82111562000f6f578081fd5b5060209081020190565b600067ffffffffffffffff82111562000f90578081fd5b50601f01601f191660200190565b60005b8381101562000fbb57818101518382015260200162000fa1565b8381111562000379575050600091015256fe60806040523480156200001157600080fd5b5060405162000fb638038062000fb6833981016040819052620000349162000496565b815160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007257506000826020015151115b80156200008457506000826040015151115b80156200009d575060008160048111156200009b57fe5b115b8015620000b657506004816004811115620000b457fe5b105b620000c057600080fd5b60408051600480825260a08201909252906020820160808036833750508151620000f292600592506020019062000283565b5081516000908155602080840151805185939262000116926001929101906200032f565b5060408201518051620001349160028401916020909101906200032f565b5060608201518051620001529160038401916020909101906200032f565b506200016b91508290506001600160e01b036200017316565b50506200059c565b60008160048111156200018257fe5b1180156200019c575060048160048111156200019a57fe5b105b620001a657600080fd5b620001ba816001600160e01b036200020b16565b620002085760016005826004811115620001d057fe5b60ff1681548110620001de57fe5b90600052602060002090602091828204019190066101000a81548160ff0219169083151502179055505b50565b6000808260048111156200021b57fe5b11801562000235575060048260048111156200023357fe5b105b6200023f57600080fd5b60058260048111156200024e57fe5b60ff16815481106200025c57fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b82805482825590600052602060002090601f016020900481019282156200031d5791602002820160005b83821115620002ec57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302620002ad565b80156200031b5782816101000a81549060ff0219169055600101602081600001049283019260010302620002ec565b505b506200032b929150620003b0565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200037257805160ff1916838001178555620003a2565b82800160010185558215620003a2579182015b82811115620003a257825182559160200191906001019062000385565b506200032b929150620003d4565b620003d191905b808211156200032b57805460ff19168155600101620003b7565b90565b620003d191905b808211156200032b5760008155600101620003db565b8051600581106200040157600080fd5b92915050565b600082601f83011262000418578081fd5b81516001600160401b038111156200042e578182fd5b602062000444601f8301601f1916820162000575565b925081835284818386010111156200045b57600080fd5b60005b828110156200047b5784810182015184820183015281016200045e565b828111156200048d5760008284860101525b50505092915050565b60008060408385031215620004a9578182fd5b82516001600160401b0380821115620004c0578384fd5b81850160808188031215620004d3578485fd5b620004df608062000575565b925080518352602081015182811115620004f7578586fd5b620005058882840162000407565b6020850152506040810151828111156200051d578586fd5b6200052b8882840162000407565b60408501525060608101518281111562000543578586fd5b620005518882840162000407565b6060850152505050809250506200056c8460208501620003f1565b90509250929050565b6040518181016001600160401b03811182821017156200059457600080fd5b604052919050565b610a0a80620005ac6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063b4579d601161005b578063b4579d60146100f3578063c2dc3e2214610108578063d1f439e51461011d578063e188130b1461013057610088565b8063222f74e71461008d5780636d4a255d146100b65780636d4ce63c146100c9578063a26cd7a6146100de575b600080fd5b6100a061009b366004610779565b610143565b6040516100ad9190610938565b60405180910390f35b6100a06100c4366004610761565b6101b5565b6100d1610215565b6040516100ad9190610943565b6100f16100ec366004610761565b6103ee565b005b6100fb610462565b6040516100ad91906108ba565b6101106104da565b6040516100ad9190610900565b6100f161012b36600461079f565b610531565b6100f161013e366004610779565b61059b565b60008082600481111561015257fe5b11801561016a5750600482600481111561016857fe5b105b61017357600080fd5b600582600481111561018157fe5b60ff168154811061018e57fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b600081811a60f81b6001600160f81b0319166101d057600080fd5b6000805b60045481101561020e5783600482815481106101ec57fe5b90600052602060002001541415610206576001915061020e565b6001016101d4565b5092915050565b61021d610620565b604080516080810182526000805482526001805484516020600283851615610100026000190190931692909204601f810183900483028201830190965285815293949293818601939092918301828280156102b95780601f1061028e576101008083540402835291602001916102b9565b820191906000526020600020905b81548152906001019060200180831161029c57829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561034b5780601f106103205761010080835404028352916020019161034b565b820191906000526020600020905b81548152906001019060200180831161032e57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103df5780601f106103b4576101008083540402835291602001916103df565b820191906000526020600020905b8154815290600101906020018083116103c257829003601f168201915b50505050508152505090505b90565b6103f86001610143565b801561041357508060001a60f81b6001600160f81b03191615155b61041c57600080fd5b610425816101b5565b61045f57600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018190555b50565b606060058054806020026020016040519081016040528092919081815260200182805480156104d057602002820191906000526020600020906000905b825461010083900a900460ff16151581526020600192830181810494850194909303909202910180841161049f5790505b5050505050905090565b606060048054806020026020016040519081016040528092919081815260200182805480156104d057602002820191906000526020600020905b815481526020019060010190808311610514575050505050905090565b8151600090815560208084015180518593926105529260019291019061064b565b506040820151805161056e91600284019160209091019061064b565b506060820151805161058a91600384019160209091019061064b565b509050506105978161059b565b5050565b60008160048111156105a957fe5b1180156105c1575060048160048111156105bf57fe5b105b6105ca57600080fd5b6105d381610143565b61045f57600160058260048111156105e757fe5b60ff16815481106105f457fe5b90600052602060002090602091828204019190066101000a81548160ff02191690831515021790555050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061068c57805160ff19168380011785556106b9565b828001600101855582156106b9579182015b828111156106b957825182559160200191906001019061069e565b506106c59291506106c9565b5090565b6103eb91905b808211156106c557600081556001016106cf565b8035600581106106f257600080fd5b92915050565b600082601f830112610708578081fd5b813567ffffffffffffffff81111561071e578182fd5b610731601f8201601f19166020016109ad565b915080825283602082850101111561074857600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610772578081fd5b5035919050565b60006020828403121561078a578081fd5b813560058110610798578182fd5b9392505050565b600080604083850312156107b1578081fd5b823567ffffffffffffffff808211156107c8578283fd5b818501608081880312156107da578384fd5b6107e460806109ad565b9250803583526020810135828111156107fb578485fd5b610807888284016106f8565b60208501525060408101358281111561081e578485fd5b61082a888284016106f8565b604085015250606081013582811115610841578485fd5b61084d888284016106f8565b60608501525050508092505061086684602085016106e3565b90509250929050565b60008151808452815b8181101561089457602081850181015186830182015201610878565b818111156108a55782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b818110156108f45783511515835292840192918401916001016108d6565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156108f45783518352928401929184019160010161091c565b901515815260200190565b6000602082528251602083015260208301516080604084015261096960a084018261086f565b60408501519150601f1980858303016060860152610987828461086f565b60608701519350818682030160808701526109a2818561086f565b979650505050505050565b60405181810167ffffffffffffffff811182821017156109cc57600080fd5b60405291905056fea26469706673582212205100593a9f8921e27d0c343bdc93f44527705437c50ca3dcfea2a9933c0e82fc64736f6c63430006080033a2646970667358221220e2eaeb4c2bd7a75c43ea7bb00aa9bb49aac9723579e92396b60ea38f6449cd7d64736f6c63430006080033"

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

// GetEntityRelations is a free data retrieval call binding the contract method 0xb24aa090.
//
// Solidity: function getEntityRelations(bytes32 _id) view returns(bytes32[])
func (_Entities *EntitiesCaller) GetEntityRelations(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getEntityRelations", _id)
	return *ret0, err
}

// GetEntityRelations is a free data retrieval call binding the contract method 0xb24aa090.
//
// Solidity: function getEntityRelations(bytes32 _id) view returns(bytes32[])
func (_Entities *EntitiesSession) GetEntityRelations(_id [32]byte) ([][32]byte, error) {
	return _Entities.Contract.GetEntityRelations(&_Entities.CallOpts, _id)
}

// GetEntityRelations is a free data retrieval call binding the contract method 0xb24aa090.
//
// Solidity: function getEntityRelations(bytes32 _id) view returns(bytes32[])
func (_Entities *EntitiesCallerSession) GetEntityRelations(_id [32]byte) ([][32]byte, error) {
	return _Entities.Contract.GetEntityRelations(&_Entities.CallOpts, _id)
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

// IsEntityRelation is a free data retrieval call binding the contract method 0x941c0169.
//
// Solidity: function isEntityRelation(bytes32 _id, bytes32 _relatedEntityId) view returns(bool)
func (_Entities *EntitiesCaller) IsEntityRelation(opts *bind.CallOpts, _id [32]byte, _relatedEntityId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "isEntityRelation", _id, _relatedEntityId)
	return *ret0, err
}

// IsEntityRelation is a free data retrieval call binding the contract method 0x941c0169.
//
// Solidity: function isEntityRelation(bytes32 _id, bytes32 _relatedEntityId) view returns(bool)
func (_Entities *EntitiesSession) IsEntityRelation(_id [32]byte, _relatedEntityId [32]byte) (bool, error) {
	return _Entities.Contract.IsEntityRelation(&_Entities.CallOpts, _id, _relatedEntityId)
}

// IsEntityRelation is a free data retrieval call binding the contract method 0x941c0169.
//
// Solidity: function isEntityRelation(bytes32 _id, bytes32 _relatedEntityId) view returns(bool)
func (_Entities *EntitiesCallerSession) IsEntityRelation(_id [32]byte, _relatedEntityId [32]byte) (bool, error) {
	return _Entities.Contract.IsEntityRelation(&_Entities.CallOpts, _id, _relatedEntityId)
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

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _relatedId) returns()
func (_Entities *EntitiesTransactor) AddEntityRelation(opts *bind.TransactOpts, _parentId [32]byte, _relatedId [32]byte) (*types.Transaction, error) {
	return _Entities.contract.Transact(opts, "addEntityRelation", _parentId, _relatedId)
}

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _relatedId) returns()
func (_Entities *EntitiesSession) AddEntityRelation(_parentId [32]byte, _relatedId [32]byte) (*types.Transaction, error) {
	return _Entities.Contract.AddEntityRelation(&_Entities.TransactOpts, _parentId, _relatedId)
}

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _relatedId) returns()
func (_Entities *EntitiesTransactorSession) AddEntityRelation(_parentId [32]byte, _relatedId [32]byte) (*types.Transaction, error) {
	return _Entities.Contract.AddEntityRelation(&_Entities.TransactOpts, _parentId, _relatedId)
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
const EntityNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_relatedId\",\"type\":\"bytes32\"}],\"name\":\"addRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelations\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_relatedEntityId\",\"type\":\"bytes32\"}],\"name\":\"isRelation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntityNodeFuncSigs maps the 4-byte function signature to its string representation.
var EntityNodeFuncSigs = map[string]string{
	"a26cd7a6": "addRelation(bytes32)",
	"e188130b": "addType(uint8)",
	"d1f439e5": "amend((bytes32,string,string,string),uint8)",
	"6d4ce63c": "get()",
	"c2dc3e22": "getRelations()",
	"b4579d60": "getTypes()",
	"6d4a255d": "isRelation(bytes32)",
	"222f74e7": "isType(uint8)",
}

// EntityNodeBin is the compiled bytecode used for deploying new contracts.
var EntityNodeBin = "0x60806040523480156200001157600080fd5b5060405162000fb638038062000fb6833981016040819052620000349162000496565b815160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007257506000826020015151115b80156200008457506000826040015151115b80156200009d575060008160048111156200009b57fe5b115b8015620000b657506004816004811115620000b457fe5b105b620000c057600080fd5b60408051600480825260a08201909252906020820160808036833750508151620000f292600592506020019062000283565b5081516000908155602080840151805185939262000116926001929101906200032f565b5060408201518051620001349160028401916020909101906200032f565b5060608201518051620001529160038401916020909101906200032f565b506200016b91508290506001600160e01b036200017316565b50506200059c565b60008160048111156200018257fe5b1180156200019c575060048160048111156200019a57fe5b105b620001a657600080fd5b620001ba816001600160e01b036200020b16565b620002085760016005826004811115620001d057fe5b60ff1681548110620001de57fe5b90600052602060002090602091828204019190066101000a81548160ff0219169083151502179055505b50565b6000808260048111156200021b57fe5b11801562000235575060048260048111156200023357fe5b105b6200023f57600080fd5b60058260048111156200024e57fe5b60ff16815481106200025c57fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b82805482825590600052602060002090601f016020900481019282156200031d5791602002820160005b83821115620002ec57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302620002ad565b80156200031b5782816101000a81549060ff0219169055600101602081600001049283019260010302620002ec565b505b506200032b929150620003b0565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200037257805160ff1916838001178555620003a2565b82800160010185558215620003a2579182015b82811115620003a257825182559160200191906001019062000385565b506200032b929150620003d4565b620003d191905b808211156200032b57805460ff19168155600101620003b7565b90565b620003d191905b808211156200032b5760008155600101620003db565b8051600581106200040157600080fd5b92915050565b600082601f83011262000418578081fd5b81516001600160401b038111156200042e578182fd5b602062000444601f8301601f1916820162000575565b925081835284818386010111156200045b57600080fd5b60005b828110156200047b5784810182015184820183015281016200045e565b828111156200048d5760008284860101525b50505092915050565b60008060408385031215620004a9578182fd5b82516001600160401b0380821115620004c0578384fd5b81850160808188031215620004d3578485fd5b620004df608062000575565b925080518352602081015182811115620004f7578586fd5b620005058882840162000407565b6020850152506040810151828111156200051d578586fd5b6200052b8882840162000407565b60408501525060608101518281111562000543578586fd5b620005518882840162000407565b6060850152505050809250506200056c8460208501620003f1565b90509250929050565b6040518181016001600160401b03811182821017156200059457600080fd5b604052919050565b610a0a80620005ac6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063b4579d601161005b578063b4579d60146100f3578063c2dc3e2214610108578063d1f439e51461011d578063e188130b1461013057610088565b8063222f74e71461008d5780636d4a255d146100b65780636d4ce63c146100c9578063a26cd7a6146100de575b600080fd5b6100a061009b366004610779565b610143565b6040516100ad9190610938565b60405180910390f35b6100a06100c4366004610761565b6101b5565b6100d1610215565b6040516100ad9190610943565b6100f16100ec366004610761565b6103ee565b005b6100fb610462565b6040516100ad91906108ba565b6101106104da565b6040516100ad9190610900565b6100f161012b36600461079f565b610531565b6100f161013e366004610779565b61059b565b60008082600481111561015257fe5b11801561016a5750600482600481111561016857fe5b105b61017357600080fd5b600582600481111561018157fe5b60ff168154811061018e57fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b600081811a60f81b6001600160f81b0319166101d057600080fd5b6000805b60045481101561020e5783600482815481106101ec57fe5b90600052602060002001541415610206576001915061020e565b6001016101d4565b5092915050565b61021d610620565b604080516080810182526000805482526001805484516020600283851615610100026000190190931692909204601f810183900483028201830190965285815293949293818601939092918301828280156102b95780601f1061028e576101008083540402835291602001916102b9565b820191906000526020600020905b81548152906001019060200180831161029c57829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561034b5780601f106103205761010080835404028352916020019161034b565b820191906000526020600020905b81548152906001019060200180831161032e57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103df5780601f106103b4576101008083540402835291602001916103df565b820191906000526020600020905b8154815290600101906020018083116103c257829003601f168201915b50505050508152505090505b90565b6103f86001610143565b801561041357508060001a60f81b6001600160f81b03191615155b61041c57600080fd5b610425816101b5565b61045f57600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018190555b50565b606060058054806020026020016040519081016040528092919081815260200182805480156104d057602002820191906000526020600020906000905b825461010083900a900460ff16151581526020600192830181810494850194909303909202910180841161049f5790505b5050505050905090565b606060048054806020026020016040519081016040528092919081815260200182805480156104d057602002820191906000526020600020905b815481526020019060010190808311610514575050505050905090565b8151600090815560208084015180518593926105529260019291019061064b565b506040820151805161056e91600284019160209091019061064b565b506060820151805161058a91600384019160209091019061064b565b509050506105978161059b565b5050565b60008160048111156105a957fe5b1180156105c1575060048160048111156105bf57fe5b105b6105ca57600080fd5b6105d381610143565b61045f57600160058260048111156105e757fe5b60ff16815481106105f457fe5b90600052602060002090602091828204019190066101000a81548160ff02191690831515021790555050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061068c57805160ff19168380011785556106b9565b828001600101855582156106b9579182015b828111156106b957825182559160200191906001019061069e565b506106c59291506106c9565b5090565b6103eb91905b808211156106c557600081556001016106cf565b8035600581106106f257600080fd5b92915050565b600082601f830112610708578081fd5b813567ffffffffffffffff81111561071e578182fd5b610731601f8201601f19166020016109ad565b915080825283602082850101111561074857600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610772578081fd5b5035919050565b60006020828403121561078a578081fd5b813560058110610798578182fd5b9392505050565b600080604083850312156107b1578081fd5b823567ffffffffffffffff808211156107c8578283fd5b818501608081880312156107da578384fd5b6107e460806109ad565b9250803583526020810135828111156107fb578485fd5b610807888284016106f8565b60208501525060408101358281111561081e578485fd5b61082a888284016106f8565b604085015250606081013582811115610841578485fd5b61084d888284016106f8565b60608501525050508092505061086684602085016106e3565b90509250929050565b60008151808452815b8181101561089457602081850181015186830182015201610878565b818111156108a55782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b818110156108f45783511515835292840192918401916001016108d6565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156108f45783518352928401929184019160010161091c565b901515815260200190565b6000602082528251602083015260208301516080604084015261096960a084018261086f565b60408501519150601f1980858303016060860152610987828461086f565b60608701519350818682030160808701526109a2818561086f565b979650505050505050565b60405181810167ffffffffffffffff811182821017156109cc57600080fd5b60405291905056fea26469706673582212205100593a9f8921e27d0c343bdc93f44527705437c50ca3dcfea2a9933c0e82fc64736f6c63430006080033"

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

// GetRelations is a free data retrieval call binding the contract method 0xc2dc3e22.
//
// Solidity: function getRelations() view returns(bytes32[])
func (_EntityNode *EntityNodeCaller) GetRelations(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "getRelations")
	return *ret0, err
}

// GetRelations is a free data retrieval call binding the contract method 0xc2dc3e22.
//
// Solidity: function getRelations() view returns(bytes32[])
func (_EntityNode *EntityNodeSession) GetRelations() ([][32]byte, error) {
	return _EntityNode.Contract.GetRelations(&_EntityNode.CallOpts)
}

// GetRelations is a free data retrieval call binding the contract method 0xc2dc3e22.
//
// Solidity: function getRelations() view returns(bytes32[])
func (_EntityNode *EntityNodeCallerSession) GetRelations() ([][32]byte, error) {
	return _EntityNode.Contract.GetRelations(&_EntityNode.CallOpts)
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

// IsRelation is a free data retrieval call binding the contract method 0x6d4a255d.
//
// Solidity: function isRelation(bytes32 _relatedEntityId) view returns(bool)
func (_EntityNode *EntityNodeCaller) IsRelation(opts *bind.CallOpts, _relatedEntityId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "isRelation", _relatedEntityId)
	return *ret0, err
}

// IsRelation is a free data retrieval call binding the contract method 0x6d4a255d.
//
// Solidity: function isRelation(bytes32 _relatedEntityId) view returns(bool)
func (_EntityNode *EntityNodeSession) IsRelation(_relatedEntityId [32]byte) (bool, error) {
	return _EntityNode.Contract.IsRelation(&_EntityNode.CallOpts, _relatedEntityId)
}

// IsRelation is a free data retrieval call binding the contract method 0x6d4a255d.
//
// Solidity: function isRelation(bytes32 _relatedEntityId) view returns(bool)
func (_EntityNode *EntityNodeCallerSession) IsRelation(_relatedEntityId [32]byte) (bool, error) {
	return _EntityNode.Contract.IsRelation(&_EntityNode.CallOpts, _relatedEntityId)
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

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _relatedId) returns()
func (_EntityNode *EntityNodeTransactor) AddRelation(opts *bind.TransactOpts, _relatedId [32]byte) (*types.Transaction, error) {
	return _EntityNode.contract.Transact(opts, "addRelation", _relatedId)
}

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _relatedId) returns()
func (_EntityNode *EntityNodeSession) AddRelation(_relatedId [32]byte) (*types.Transaction, error) {
	return _EntityNode.Contract.AddRelation(&_EntityNode.TransactOpts, _relatedId)
}

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _relatedId) returns()
func (_EntityNode *EntityNodeTransactorSession) AddRelation(_relatedId [32]byte) (*types.Transaction, error) {
	return _EntityNode.Contract.AddRelation(&_EntityNode.TransactOpts, _relatedId)
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
const IEntitiesFactoryABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedId\",\"type\":\"bytes32\"}],\"name\":\"addEntityRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amendEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntity\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityRelations\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedEntityId\",\"type\":\"bytes32\"}],\"name\":\"isEntityRelation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isEntityType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEntitiesFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IEntitiesFactoryFuncSigs = map[string]string{
	"68a9f4e9": "addEntity((bytes32,string,string,string),uint8)",
	"6c3cf7f7": "addEntityRelation(bytes32,bytes32)",
	"a81eb545": "amendEntity((bytes32,string,string,string),uint8)",
	"53b66f36": "getEntity(bytes32)",
	"d91ee2e0": "getEntityContract(bytes32)",
	"b24aa090": "getEntityRelations(bytes32)",
	"6ec21da9": "getEntityTypes(bytes32)",
	"941c0169": "isEntityRelation(bytes32,bytes32)",
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

// GetEntityRelations is a free data retrieval call binding the contract method 0xb24aa090.
//
// Solidity: function getEntityRelations(bytes32 _id) view returns(bytes32[])
func (_IEntitiesFactory *IEntitiesFactoryCaller) GetEntityRelations(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "getEntityRelations", _id)
	return *ret0, err
}

// GetEntityRelations is a free data retrieval call binding the contract method 0xb24aa090.
//
// Solidity: function getEntityRelations(bytes32 _id) view returns(bytes32[])
func (_IEntitiesFactory *IEntitiesFactorySession) GetEntityRelations(_id [32]byte) ([][32]byte, error) {
	return _IEntitiesFactory.Contract.GetEntityRelations(&_IEntitiesFactory.CallOpts, _id)
}

// GetEntityRelations is a free data retrieval call binding the contract method 0xb24aa090.
//
// Solidity: function getEntityRelations(bytes32 _id) view returns(bytes32[])
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) GetEntityRelations(_id [32]byte) ([][32]byte, error) {
	return _IEntitiesFactory.Contract.GetEntityRelations(&_IEntitiesFactory.CallOpts, _id)
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

// IsEntityRelation is a free data retrieval call binding the contract method 0x941c0169.
//
// Solidity: function isEntityRelation(bytes32 _id, bytes32 _relatedEntityId) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactoryCaller) IsEntityRelation(opts *bind.CallOpts, _id [32]byte, _relatedEntityId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "isEntityRelation", _id, _relatedEntityId)
	return *ret0, err
}

// IsEntityRelation is a free data retrieval call binding the contract method 0x941c0169.
//
// Solidity: function isEntityRelation(bytes32 _id, bytes32 _relatedEntityId) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactorySession) IsEntityRelation(_id [32]byte, _relatedEntityId [32]byte) (bool, error) {
	return _IEntitiesFactory.Contract.IsEntityRelation(&_IEntitiesFactory.CallOpts, _id, _relatedEntityId)
}

// IsEntityRelation is a free data retrieval call binding the contract method 0x941c0169.
//
// Solidity: function isEntityRelation(bytes32 _id, bytes32 _relatedEntityId) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) IsEntityRelation(_id [32]byte, _relatedEntityId [32]byte) (bool, error) {
	return _IEntitiesFactory.Contract.IsEntityRelation(&_IEntitiesFactory.CallOpts, _id, _relatedEntityId)
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

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _relatedId) returns()
func (_IEntitiesFactory *IEntitiesFactoryTransactor) AddEntityRelation(opts *bind.TransactOpts, _parentId [32]byte, _relatedId [32]byte) (*types.Transaction, error) {
	return _IEntitiesFactory.contract.Transact(opts, "addEntityRelation", _parentId, _relatedId)
}

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _relatedId) returns()
func (_IEntitiesFactory *IEntitiesFactorySession) AddEntityRelation(_parentId [32]byte, _relatedId [32]byte) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.AddEntityRelation(&_IEntitiesFactory.TransactOpts, _parentId, _relatedId)
}

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _relatedId) returns()
func (_IEntitiesFactory *IEntitiesFactoryTransactorSession) AddEntityRelation(_parentId [32]byte, _relatedId [32]byte) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.AddEntityRelation(&_IEntitiesFactory.TransactOpts, _parentId, _relatedId)
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
const IEntityABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_relatedId\",\"type\":\"bytes32\"}],\"name\":\"addRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelations\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_relatedEntityId\",\"type\":\"bytes32\"}],\"name\":\"isRelation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEntityFuncSigs maps the 4-byte function signature to its string representation.
var IEntityFuncSigs = map[string]string{
	"a26cd7a6": "addRelation(bytes32)",
	"e188130b": "addType(uint8)",
	"d1f439e5": "amend((bytes32,string,string,string),uint8)",
	"6d4ce63c": "get()",
	"c2dc3e22": "getRelations()",
	"b4579d60": "getTypes()",
	"6d4a255d": "isRelation(bytes32)",
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

// GetRelations is a free data retrieval call binding the contract method 0xc2dc3e22.
//
// Solidity: function getRelations() view returns(bytes32[])
func (_IEntity *IEntityCaller) GetRelations(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IEntity.contract.Call(opts, out, "getRelations")
	return *ret0, err
}

// GetRelations is a free data retrieval call binding the contract method 0xc2dc3e22.
//
// Solidity: function getRelations() view returns(bytes32[])
func (_IEntity *IEntitySession) GetRelations() ([][32]byte, error) {
	return _IEntity.Contract.GetRelations(&_IEntity.CallOpts)
}

// GetRelations is a free data retrieval call binding the contract method 0xc2dc3e22.
//
// Solidity: function getRelations() view returns(bytes32[])
func (_IEntity *IEntityCallerSession) GetRelations() ([][32]byte, error) {
	return _IEntity.Contract.GetRelations(&_IEntity.CallOpts)
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

// IsRelation is a free data retrieval call binding the contract method 0x6d4a255d.
//
// Solidity: function isRelation(bytes32 _relatedEntityId) view returns(bool)
func (_IEntity *IEntityCaller) IsRelation(opts *bind.CallOpts, _relatedEntityId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEntity.contract.Call(opts, out, "isRelation", _relatedEntityId)
	return *ret0, err
}

// IsRelation is a free data retrieval call binding the contract method 0x6d4a255d.
//
// Solidity: function isRelation(bytes32 _relatedEntityId) view returns(bool)
func (_IEntity *IEntitySession) IsRelation(_relatedEntityId [32]byte) (bool, error) {
	return _IEntity.Contract.IsRelation(&_IEntity.CallOpts, _relatedEntityId)
}

// IsRelation is a free data retrieval call binding the contract method 0x6d4a255d.
//
// Solidity: function isRelation(bytes32 _relatedEntityId) view returns(bool)
func (_IEntity *IEntityCallerSession) IsRelation(_relatedEntityId [32]byte) (bool, error) {
	return _IEntity.Contract.IsRelation(&_IEntity.CallOpts, _relatedEntityId)
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

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _relatedId) returns()
func (_IEntity *IEntityTransactor) AddRelation(opts *bind.TransactOpts, _relatedId [32]byte) (*types.Transaction, error) {
	return _IEntity.contract.Transact(opts, "addRelation", _relatedId)
}

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _relatedId) returns()
func (_IEntity *IEntitySession) AddRelation(_relatedId [32]byte) (*types.Transaction, error) {
	return _IEntity.Contract.AddRelation(&_IEntity.TransactOpts, _relatedId)
}

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _relatedId) returns()
func (_IEntity *IEntityTransactorSession) AddRelation(_relatedId [32]byte) (*types.Transaction, error) {
	return _IEntity.Contract.AddRelation(&_IEntity.TransactOpts, _relatedId)
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
