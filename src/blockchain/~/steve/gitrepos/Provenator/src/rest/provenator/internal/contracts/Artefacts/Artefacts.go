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
	DateCreated  [32]byte
	DateModified [32]byte
	Url          string
	Name         string
	Description  string
	Author       CreativeEntities
}

// Works is an auto generated low-level Go binding around an user-defined struct.
type Works struct {
	WorkType     uint8
	License      uint8
	DateCreated  [32]byte
	DateModified [32]byte
	Url          string
	Name         string
	Description  string
}

// ArtefactNodeABI is the input ABI used to generate the binding from.
const ArtefactNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_entityFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_author\",\"type\":\"tuple\"}],\"name\":\"addAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_copyrightHolder\",\"type\":\"tuple\"}],\"name\":\"addCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_publisher\",\"type\":\"tuple\"}],\"name\":\"addPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthors\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCopyrightHolders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublishers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isCopyrightHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isPublisher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArtefactNodeFuncSigs maps the 4-byte function signature to its string representation.
var ArtefactNodeFuncSigs = map[string]string{
	"880a8bce": "addAuthor((bytes32,string,string,string))",
	"0d464d44": "addCopyrightHolder((bytes32,string,string,string))",
	"49952c49": "addPublisher((bytes32,string,string,string))",
	"7502f03c": "amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)))",
	"6d4ce63c": "get()",
	"7fc622f4": "getAuthors()",
	"d6d54164": "getCopyrightHolders()",
	"6c6071aa": "getPublishers()",
	"59c0161c": "isAuthor(bytes32)",
	"476caa00": "isCopyrightHolder(bytes32)",
	"240b5600": "isPublisher(bytes32)",
}

// ArtefactNodeBin is the compiled bytecode used for deploying new contracts.
var ArtefactNodeBin = "0x60806040523480156200001157600080fd5b50604051620018243803806200182483398101604081905262000034916200051b565b604082015160001a60f81b6001600160f81b0319166200005357600080fd5b600080546001600160a01b0319166001600160a01b0383161790556040820151600155620000818262000089565b50506200076d565b6000815160158111156200009957fe5b118015620000b45750601581516015811115620000b257fe5b105b8015620000d15750600081602001516014811115620000cf57fe5b115b8015620000ee5750601481602001516014811115620000ec57fe5b105b80156200010e5750606081015160001a60f81b6001600160f81b03191615155b801562000120575060008160c0015151115b801562000132575060008160e0015151115b80156200015457506101008101515160001a60f81b6001600160f81b03191615155b6200015e57600080fd5b80516002805460ff191660018360158111156200017757fe5b021790555060208101516002805461ff0019166101008360148111156200019a57fe5b02179055506060810151600355608081015160045560a08101518051620001ca916005916020909101906200032f565b5060c08101518051620001e6916006916020909101906200032f565b5060e0810151805162000202916007916020909101906200032f565b50610100810151620002149062000217565b50565b6000546040516368a9f4e960e01b81526001600160a01b03909116906368a9f4e9906200024c9084906001906004016200068f565b600060405180830381600087803b1580156200026757600080fd5b505af11580156200027c573d6000803e3d6000fd5b505082516200028e92509050620002c9565b620002145751600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b600081811a60f81b6001600160f81b031916620002e557600080fd5b6000805b600854811015620003285783600882815481106200030357fe5b906000526020600020015414156200031f576001915062000328565b600101620002e9565b5092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200037257805160ff1916838001178555620003a2565b82800160010185558215620003a2579182015b82811115620003a257825182559160200191906001019062000385565b50620003b0929150620003b4565b5090565b5b80821115620003b05760008155600101620003b5565b80516001600160a01b0381168114620003e357600080fd5b92915050565b805160158110620003e357600080fd5b805160168110620003e357600080fd5b600082601f8301126200041a578081fd5b81516001600160401b0381111562000430578182fd5b62000445601f8201601f191660200162000713565b91508082528360208285010111156200045d57600080fd5b620003288160208401602086016200073a565b60006080828403121562000482578081fd5b6200048e608062000713565b8251815260208301519091506001600160401b0380821115620004b057600080fd5b620004be8583860162000409565b60208401526040840151915080821115620004d857600080fd5b620004e68583860162000409565b604084015260608401519150808211156200050057600080fd5b506200050f8482850162000409565b60608301525092915050565b600080604083850312156200052e578182fd5b82516001600160401b038082111562000545578384fd5b81850191506101208083880312156200055c578485fd5b620005678162000713565b9050620005758784620003f9565b8152620005868760208501620003e9565b602082015260408301516040820152606083015160608201526080830151608082015260a083015182811115620005bb578586fd5b620005c98882860162000409565b60a08301525060c083015182811115620005e1578586fd5b620005ef8882860162000409565b60c08301525060e08301518281111562000607578586fd5b620006158882860162000409565b60e08301525061010080840151838111156200062f578687fd5b6200063d8982870162000470565b828401525050809450505050620006588460208501620003cb565b90509250929050565b600081518084526200067b8160208601602086016200073a565b601f01601f19169290920160200192915050565b60006040825283516040830152602084015160806060840152620006b760c084018262000661565b90506040850151603f1980858403016080860152620006d7838362000661565b925060608701519150808584030160a086015250620006f7828262000661565b92505050600583106200070657fe5b8260208301529392505050565b6040518181016001600160401b03811182821017156200073257600080fd5b604052919050565b60005b83811015620007575781810151838201526020016200073d565b8381111562000767576000848401525b50505050565b6110a7806200077d6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80636c6071aa116100715780636c6071aa146101255780636d4ce63c1461013a5780637502f03c1461014f5780637fc622f414610162578063880a8bce1461016a578063d6d541641461017d576100a9565b80630d464d44146100ae578063240b5600146100c3578063476caa00146100ec57806349952c49146100ff57806359c0161c14610112575b600080fd5b6100c16100bc366004610d1c565b610185565b005b6100d66100d1366004610d04565b6102fe565b6040516100e39190610f04565b60405180910390f35b6100d66100fa366004610d04565b610360565b6100c161010d366004610d1c565b6103b9565b6100d6610120366004610d04565b61052f565b61012d610588565b6040516100e39190610ec0565b6101426105e0565b6040516100e39190610f9a565b6100c161015d366004610d57565b610821565b61012d61083c565b6100c1610178366004610d1c565b610892565b61012d61093e565b6000546040516368a9f4e960e01b81526001600160a01b03909116906368a9f4e9906101b8908490600290600401610f1d565b600060405180830381600087803b1580156101d257600080fd5b505af11580156101e6573d6000803e3d6000fd5b50506000546001548451604051636c3cf7f760e01b81526001600160a01b039093169450636c3cf7f7935061021d92600401610f0f565b600060405180830381600087803b15801561023757600080fd5b505af115801561024b573d6000803e3d6000fd5b50506000548351600154604051636c3cf7f760e01b81526001600160a01b039093169450636c3cf7f7935061028292600401610f0f565b600060405180830381600087803b15801561029c57600080fd5b505af11580156102b0573d6000803e3d6000fd5b505050506102c18160000151610360565b6102fb578051600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af01555b50565b600081811a60f81b6001600160f81b03191661031957600080fd5b6000805b600a548110156103575783600a828154811061033557fe5b9060005260206000200154141561034f5760019150610357565b60010161031d565b5090505b919050565b600081811a60f81b6001600160f81b03191661037b57600080fd5b6000805b60095481101561035757836009828154811061039757fe5b906000526020600020015414156103b15760019150610357565b60010161037f565b6000546040516368a9f4e960e01b81526001600160a01b03909116906368a9f4e9906103ec908490600390600401610f1d565b600060405180830381600087803b15801561040657600080fd5b505af115801561041a573d6000803e3d6000fd5b50506000546001548451604051636c3cf7f760e01b81526001600160a01b039093169450636c3cf7f7935061045192600401610f0f565b600060405180830381600087803b15801561046b57600080fd5b505af115801561047f573d6000803e3d6000fd5b50506000548351600154604051636c3cf7f760e01b81526001600160a01b039093169450636c3cf7f793506104b692600401610f0f565b600060405180830381600087803b1580156104d057600080fd5b505af11580156104e4573d6000803e3d6000fd5b505050506104f581600001516102fe565b6102fb5751600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80155565b600081811a60f81b6001600160f81b03191661054a57600080fd5b6000805b60085481101561035757836008828154811061056657fe5b906000526020600020015414156105805760019150610357565b60010161054e565b6060600a8054806020026020016040519081016040528092919081815260200182805480156105d657602002820191906000526020600020905b8154815260200190600101908083116105c2575b5050505050905090565b6105e8610b08565b6040805160e0810190915260028054829060ff16601581111561060757fe5b601581111561061257fe5b81528154602090910190610100900460ff16601481111561062f57fe5b601481111561063a57fe5b81526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106eb5780601f106106c0576101008083540402835291602001916106eb565b820191906000526020600020905b8154815290600101906020018083116106ce57829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561077f5780601f106107545761010080835404028352916020019161077f565b820191906000526020600020905b81548152906001019060200180831161076257829003601f168201915b505050918352505060058201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156108135780601f106107e857610100808354040283529160200191610813565b820191906000526020600020905b8154815290600101906020018083116107f657829003601f168201915b505050505081525050905090565b80604001516001541461083357600080fd5b6102fb81610994565b606060088054806020026020016040519081016040528092919081815260200182805480156105d657602002820191906000526020600020908154815260200190600101908083116105c2575050505050905090565b6000546040516368a9f4e960e01b81526001600160a01b03909116906368a9f4e9906108c5908490600190600401610f1d565b600060405180830381600087803b1580156108df57600080fd5b505af11580156108f3573d6000803e3d6000fd5b50505050610904816000015161052f565b6102fb5751600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b606060098054806020026020016040519081016040528092919081815260200182805480156105d657602002820191906000526020600020908154815260200190600101908083116105c2575050505050905090565b6000815160158111156109a357fe5b1180156109bc57506015815160158111156109ba57fe5b105b80156109d757506000816020015160148111156109d557fe5b115b80156109f257506014816020015160148111156109f057fe5b105b8015610a115750606081015160001a60f81b6001600160f81b03191615155b8015610a22575060008160c0015151115b8015610a33575060008160e0015151115b8015610a5457506101008101515160001a60f81b6001600160f81b03191615155b610a5d57600080fd5b80516002805460ff19166001836015811115610a7557fe5b021790555060208101516002805461ff001916610100836014811115610a9757fe5b02179055506060810151600355608081015160045560a08101518051610ac591600591602090910190610b45565b5060c08101518051610adf91600691602090910190610b45565b5060e08101518051610af991600791602090910190610b45565b506102fb816101000151610892565b6040805160e08101909152806000815260200160008152600060208201819052604082015260608082018190526080820181905260a09091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b8657805160ff1916838001178555610bb3565b82800160010185558215610bb3579182015b82811115610bb3578251825591602001919060010190610b98565b50610bbf929150610bc3565b5090565b5b80821115610bbf5760008155600101610bc4565b803560158110610be757600080fd5b92915050565b803560168110610be757600080fd5b600082601f830112610c0c578081fd5b813567ffffffffffffffff811115610c22578182fd5b610c35601f8201601f1916602001611034565b9150808252836020828501011115610c4c57600080fd5b8060208401602084013760009082016020015292915050565b600060808284031215610c76578081fd5b610c806080611034565b905081358152602082013567ffffffffffffffff80821115610ca157600080fd5b610cad85838601610bfc565b60208401526040840135915080821115610cc657600080fd5b610cd285838601610bfc565b60408401526060840135915080821115610ceb57600080fd5b50610cf884828501610bfc565b60608301525092915050565b600060208284031215610d15578081fd5b5035919050565b600060208284031215610d2d578081fd5b813567ffffffffffffffff811115610d43578182fd5b610d4f84828501610c65565b949350505050565b600060208284031215610d68578081fd5b813567ffffffffffffffff80821115610d7f578283fd5b8184019150610120808387031215610d95578384fd5b610d9e81611034565b9050610daa8684610bed565b8152610db98660208501610bd8565b602082015260408301356040820152606083013560608201526080830135608082015260a083013582811115610ded578485fd5b610df987828601610bfc565b60a08301525060c083013582811115610e10578485fd5b610e1c87828601610bfc565b60c08301525060e083013582811115610e33578485fd5b610e3f87828601610bfc565b60e0830152506101008084013583811115610e58578586fd5b610e6488828701610c65565b918301919091525095945050505050565b60008151808452815b81811015610e9a57602081850181015186830182015201610e7e565b81811115610eab5782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610ef857835183529284019291840191600101610edc565b50909695505050505050565b901515815260200190565b918252602082015260400190565b60006040825283516040830152602084015160806060840152610f4360c0840182610e75565b90506040850151603f1980858403016080860152610f618383610e75565b925060608701519150808584030160a086015250610f7f8282610e75565b9250505060058310610f8d57fe5b8260208301529392505050565b600060208252610faa8351611066565b6020830152610fbc602084015161105b565b60408301526040830151606083015260608301516080830152608083015160e060a0840152610fef610100840182610e75565b905060a0840151601f19808584030160c086015261100d8383610e75565b925060c08601519150808584030160e08601525061102b8282610e75565b95945050505050565b60405181810167ffffffffffffffff8111828210171561105357600080fd5b604052919050565b806015811061035b57fe5b806016811061035b57fefea2646970667358221220292cf56a1e01c881170cdd76ae875a14deddc93a2b9e0b8e8932840562a3d5d864736f6c63430007000033"

// DeployArtefactNode deploys a new Ethereum contract, binding an instance of ArtefactNode to it.
func DeployArtefactNode(auth *bind.TransactOpts, backend bind.ContractBackend, _work CreativeWorks, _entityFactory common.Address) (common.Address, *types.Transaction, *ArtefactNode, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtefactNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArtefactNodeBin), backend, _work, _entityFactory)
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
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_ArtefactNode *ArtefactNodeCaller) Get(opts *bind.CallOpts) (Works, error) {
	var (
		ret0 = new(Works)
	)
	out := ret0
	err := _ArtefactNode.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_ArtefactNode *ArtefactNodeSession) Get() (Works, error) {
	return _ArtefactNode.Contract.Get(&_ArtefactNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_ArtefactNode *ArtefactNodeCallerSession) Get() (Works, error) {
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

// AddAuthor is a paid mutator transaction binding the contract method 0x880a8bce.
//
// Solidity: function addAuthor((bytes32,string,string,string) _author) returns()
func (_ArtefactNode *ArtefactNodeTransactor) AddAuthor(opts *bind.TransactOpts, _author CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "addAuthor", _author)
}

// AddAuthor is a paid mutator transaction binding the contract method 0x880a8bce.
//
// Solidity: function addAuthor((bytes32,string,string,string) _author) returns()
func (_ArtefactNode *ArtefactNodeSession) AddAuthor(_author CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddAuthor(&_ArtefactNode.TransactOpts, _author)
}

// AddAuthor is a paid mutator transaction binding the contract method 0x880a8bce.
//
// Solidity: function addAuthor((bytes32,string,string,string) _author) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) AddAuthor(_author CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddAuthor(&_ArtefactNode.TransactOpts, _author)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0x0d464d44.
//
// Solidity: function addCopyrightHolder((bytes32,string,string,string) _copyrightHolder) returns()
func (_ArtefactNode *ArtefactNodeTransactor) AddCopyrightHolder(opts *bind.TransactOpts, _copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "addCopyrightHolder", _copyrightHolder)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0x0d464d44.
//
// Solidity: function addCopyrightHolder((bytes32,string,string,string) _copyrightHolder) returns()
func (_ArtefactNode *ArtefactNodeSession) AddCopyrightHolder(_copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddCopyrightHolder(&_ArtefactNode.TransactOpts, _copyrightHolder)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0x0d464d44.
//
// Solidity: function addCopyrightHolder((bytes32,string,string,string) _copyrightHolder) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) AddCopyrightHolder(_copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddCopyrightHolder(&_ArtefactNode.TransactOpts, _copyrightHolder)
}

// AddPublisher is a paid mutator transaction binding the contract method 0x49952c49.
//
// Solidity: function addPublisher((bytes32,string,string,string) _publisher) returns()
func (_ArtefactNode *ArtefactNodeTransactor) AddPublisher(opts *bind.TransactOpts, _publisher CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "addPublisher", _publisher)
}

// AddPublisher is a paid mutator transaction binding the contract method 0x49952c49.
//
// Solidity: function addPublisher((bytes32,string,string,string) _publisher) returns()
func (_ArtefactNode *ArtefactNodeSession) AddPublisher(_publisher CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddPublisher(&_ArtefactNode.TransactOpts, _publisher)
}

// AddPublisher is a paid mutator transaction binding the contract method 0x49952c49.
//
// Solidity: function addPublisher((bytes32,string,string,string) _publisher) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) AddPublisher(_publisher CreativeEntities) (*types.Transaction, error) {
	return _ArtefactNode.Contract.AddPublisher(&_ArtefactNode.TransactOpts, _publisher)
}

// Amend is a paid mutator transaction binding the contract method 0x7502f03c.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_ArtefactNode *ArtefactNodeTransactor) Amend(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "amend", _work)
}

// Amend is a paid mutator transaction binding the contract method 0x7502f03c.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_ArtefactNode *ArtefactNodeSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.Contract.Amend(&_ArtefactNode.TransactOpts, _work)
}

// Amend is a paid mutator transaction binding the contract method 0x7502f03c.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.Contract.Amend(&_ArtefactNode.TransactOpts, _work)
}

// ArtefactsABI is the input ABI used to generate the binding from.
const ArtefactsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entityFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_author\",\"type\":\"tuple\"}],\"name\":\"addWorkAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_copyrightHolder\",\"type\":\"tuple\"}],\"name\":\"addWorkCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_publisher\",\"type\":\"tuple\"}],\"name\":\"addWorkPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amendWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReferences\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWork\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkAuthors\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkCopyrightHolders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkPublishers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"isWorkAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"isWorkCopyrightHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"isWorkPublisher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArtefactsFuncSigs maps the 4-byte function signature to its string representation.
var ArtefactsFuncSigs = map[string]string{
	"42e9f1e0": "addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)))",
	"5f04d064": "addWorkAuthor(bytes32,(bytes32,string,string,string))",
	"a4a9051e": "addWorkCopyrightHolder(bytes32,(bytes32,string,string,string))",
	"0f0063d8": "addWorkPublisher(bytes32,(bytes32,string,string,string))",
	"14f806c8": "amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)))",
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
}

// ArtefactsBin is the compiled bytecode used for deploying new contracts.
var ArtefactsBin = "0x608060405234801561001057600080fd5b50604051612d47380380612d4783398101604081905261002f9161006a565b6001600160a01b03811661004257600080fd5b600080546001600160a01b0319166001600160a01b0392909216919091178155600355610098565b60006020828403121561007b578081fd5b81516001600160a01b0381168114610091578182fd5b9392505050565b612ca0806100a76000396000f3fe60806040523480156200001157600080fd5b5060043610620001185760003560e01c806367e0badb11620000a5578063b2e459cf116200006f578063b2e459cf146200025e578063bc5993411462000284578063ef90d9c9146200029b578063f6a5096514620002b25762000118565b806367e0badb146200020d5780637a6337fa146200022657806386105c651462000230578063a4a9051e14620002475762000118565b806342e9f1e011620000e757806342e9f1e014620001a25780634d2522a214620001b95780635f04d06414620001df5780635ffd09d814620001f65762000118565b80630f0063d8146200011d57806314f806c8146200013657806330266537146200014d578063426aa803146200017c575b600080fd5b620001346200012e36600462000e35565b620002c9565b005b620001346200014736600462000e7e565b6200036c565b620001646200015e36600462000dfa565b62000416565b604051620001739190620012d6565b60405180910390f35b620001936200018d36600462000dfa565b620004e1565b60405162000173919062001226565b62000134620001b336600462000e7e565b6200059c565b620001d0620001ca36600462000dfa565b6200062c565b60405162000173919062001212565b62000134620001f036600462000e35565b62000671565b620001936200020736600462000dfa565b620006db565b6200021762000757565b60405162000173919062001277565b620001936200075d565b620001936200024136600462000dfa565b62000801565b620001346200025836600462000e35565b6200087d565b620002756200026f36600462000e13565b620008e7565b6040516200017391906200126c565b620002176200029536600462000dfa565b620009ae565b62000275620002ac36600462000e13565b620009e8565b62000275620002c336600462000e13565b62000a50565b600082815260016020819052604090912001546001600160a01b0316620002ef57600080fd5b600082815260016020819052604091829020015490516349952c4960e01b81526001600160a01b039091169081906349952c49906200033390859060040162001280565b600060405180830381600087803b1580156200034e57600080fd5b505af115801562000363573d6000803e3d6000fd5b50505050505050565b60408082015160009081526001602081905291902001546001600160a01b03166200039657600080fd5b6040808201516000908152600160208190529082902001549051631d40bc0f60e21b81526001600160a01b03909116908190637502f03c90620003de90859060040162001295565b600060405180830381600087803b158015620003f957600080fd5b505af11580156200040e573d6000803e3d6000fd5b505050505050565b6200042062000b68565b600082815260016020819052604090912001546001600160a01b03166200044657600080fd5b6000828152600160208190526040808320909101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b1580156200049957600080fd5b505afa158015620004ae573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620004d8919081019062000fb1565b9150505b919050565b600081815260016020819052604090912001546060906001600160a01b03166200050a57600080fd5b6000828152600160208190526040808320909101548151631ff188bd60e21b815291516001600160a01b03909116928392637fc622f49260048083019392829003018186803b1580156200055d57600080fd5b505afa15801562000572573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620004d8919081019062000d3f565b6040810151620005af9060019062000ab8565b15620005c657620005c0816200036c565b62000629565b6000805460405183916001600160a01b031690620005e49062000ba5565b620005f1929190620012aa565b604051809103906000f0801580156200060e573d6000803e3d6000fd5b50604083015190915062000626906001908362000acd565b50505b50565b6000818152600160208190526040822001546001600160a01b03166200065157600080fd5b50600090815260016020819052604090912001546001600160a01b031690565b600082815260016020819052604090912001546001600160a01b03166200069757600080fd5b6000828152600160208190526040918290200154905163440545e760e11b81526001600160a01b0390911690819063880a8bce906200033390859060040162001280565b600081815260016020819052604090912001546060906001600160a01b03166200070457600080fd5b600082815260016020819052604080832090910154815163363038d560e11b815291516001600160a01b03909116928392636c6071aa9260048083019392829003018186803b1580156200055d57600080fd5b60035490565b60608060016002015467ffffffffffffffff811180156200077d57600080fd5b50604051908082528060200260200182016040528015620007a8578160200160208202803683370190505b50905060005b600354811015620007fb576002805482908110620007c857fe5b906000526020600020906002020160000154828281518110620007e757fe5b6020908102919091010152600101620007ae565b50905090565b600081815260016020819052604090912001546060906001600160a01b03166200082a57600080fd5b60008281526001602081905260408083209091015481516335b5505960e21b815291516001600160a01b0390911692839263d6d541649260048083019392829003018186803b1580156200055d57600080fd5b600082815260016020819052604090912001546001600160a01b0316620008a357600080fd5b60008281526001602081905260409182902001549051630351935160e21b81526001600160a01b03909116908190630d464d44906200033390859060040162001280565b6000828152600160208190526040822001546001600160a01b03166200090c57600080fd5b60008381526001602081905260409182902001549051631670058760e21b81526001600160a01b039091169081906359c0161c906200095090869060040162001277565b60206040518083038186803b1580156200096957600080fd5b505afa1580156200097e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620009a4919062000dd8565b9150505b92915050565b6003546000908210620009c057600080fd5b6002805483908110620009cf57fe5b9060005260206000209060020201600001549050919050565b6000828152600160208190526040822001546001600160a01b031662000a0d57600080fd5b60008381526001602081905260409182902001549051621205ab60e91b81526001600160a01b0390911690819063240b5600906200095090869060040162001277565b6000828152600160208190526040822001546001600160a01b031662000a7557600080fd5b600083815260016020819052604091829020015490516223b65560e91b81526001600160a01b0390911690819063476caa00906200095090869060040162001277565b60009081526020919091526040902054151590565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000b0d57600191505062000b61565b506001808501805491820180825560008681526020889052604090208190558591908390811062000b3a57fe5b6000918252602082206002918202019290925590860180546001019055915062000b619050565b9392505050565b6040805160e08101909152806000815260200160008152600060208201819052604082015260608082018190526080820181905260a09091015290565b611824806200144783390190565b8035620009a8816200142a565b8051620009a8816200142a565b8035620009a88162001438565b8051620009a88162001438565b600082601f83011262000bf8578081fd5b813562000c0f62000c0982620013ba565b62001371565b915080825283602082850101111562000c2757600080fd5b8060208401602084013760009082016020015292915050565b600082601f83011262000c51578081fd5b815162000c6262000c0982620013ba565b915080825283602082850101111562000c7a57600080fd5b62000c8d816020840160208601620013f7565b5092915050565b60006080828403121562000ca6578081fd5b62000cb2608062001371565b905081358152602082013567ffffffffffffffff8082111562000cd457600080fd5b62000ce28583860162000be7565b6020840152604084013591508082111562000cfc57600080fd5b62000d0a8583860162000be7565b6040840152606084013591508082111562000d2457600080fd5b5062000d338482850162000be7565b60608301525092915050565b6000602080838503121562000d52578182fd5b825167ffffffffffffffff81111562000d69578283fd5b8301601f8101851362000d7a578283fd5b805162000d8b62000c098262001399565b818152838101908385018584028501860189101562000da8578687fd5b8694505b8385101562000dcc57805183526001949094019391850191850162000dac565b50979650505050505050565b60006020828403121562000dea578081fd5b8151801515811462000b61578182fd5b60006020828403121562000e0c578081fd5b5035919050565b6000806040838503121562000e26578081fd5b50508035926020909101359150565b6000806040838503121562000e48578182fd5b82359150602083013567ffffffffffffffff81111562000e66578182fd5b62000e748582860162000c94565b9150509250929050565b60006020828403121562000e90578081fd5b813567ffffffffffffffff8082111562000ea8578283fd5b818401915061012080838703121562000ebf578384fd5b62000eca8162001371565b905062000ed8868462000bcd565b815262000ee9866020850162000bb3565b602082015260408301356040820152606083013560608201526080830135608082015260a08301358281111562000f1e578485fd5b62000f2c8782860162000be7565b60a08301525060c08301358281111562000f44578485fd5b62000f528782860162000be7565b60c08301525060e08301358281111562000f6a578485fd5b62000f788782860162000be7565b60e083015250610100808401358381111562000f92578586fd5b62000fa08882870162000c94565b918301919091525095945050505050565b60006020828403121562000fc3578081fd5b815167ffffffffffffffff8082111562000fdb578283fd5b9083019060e0828603121562000fef578283fd5b62000ffb60e062001371565b62001007868462000bda565b815262001018866020850162000bc0565b6020820152604083015160408201526060830151606082015260808301518281111562001043578485fd5b620010518782860162000c40565b60808301525060a08301518281111562001069578485fd5b620010778782860162000c40565b60a08301525060c0830151828111156200108f578485fd5b6200109d8782860162000c40565b60c08301525095945050505050565b60158110620010b757fe5b9052565b60168110620010b757fe5b60008151808452620010e0816020860160208601620013f7565b601f01601f19169290920160200192915050565b600081518352602082015160806020850152620011156080850182620010c6565b905060408301518482036040860152620011308282620010c6565b915050606083015184820360608601526200114c8282620010c6565b95945050505050565b600061012062001167848451620010bb565b60208301516200117b6020860182620010ac565b5060408301516040850152606083015160608501526080830151608085015260a08301518160a0860152620011b382860182620010c6565b91505060c083015184820360c0860152620011cf8282620010c6565b91505060e083015184820360e0860152620011eb8282620010c6565b9150506101008084015185830382870152620012088382620010f4565b9695505050505050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b81811015620012605783518352928401929184019160010162001242565b50909695505050505050565b901515815260200190565b90815260200190565b60006020825262000b616020830184620010f4565b60006020825262000b61602083018462001155565b600060408252620012bf604083018562001155565b905060018060a01b03831660208301529392505050565b600060208252620012e88351620013eb565b6020830152620012fc6020840151620013df565b60408301526040830151606083015260608301516080830152608083015160e060a084015262001331610100840182620010c6565b905060a0840151601f19808584030160c0860152620013518383620010c6565b925060c08601519150808584030160e0860152506200114c8282620010c6565b60405181810167ffffffffffffffff811182821017156200139157600080fd5b604052919050565b600067ffffffffffffffff821115620013b0578081fd5b5060209081020190565b600067ffffffffffffffff821115620013d1578081fd5b50601f01601f191660200190565b8060158110620004dc57fe5b8060168110620004dc57fe5b60005b8381101562001414578181015183820152602001620013fa565b8381111562001424576000848401525b50505050565b601581106200062957600080fd5b601681106200062957600080fdfe60806040523480156200001157600080fd5b50604051620018243803806200182483398101604081905262000034916200051b565b604082015160001a60f81b6001600160f81b0319166200005357600080fd5b600080546001600160a01b0319166001600160a01b0383161790556040820151600155620000818262000089565b50506200076d565b6000815160158111156200009957fe5b118015620000b45750601581516015811115620000b257fe5b105b8015620000d15750600081602001516014811115620000cf57fe5b115b8015620000ee5750601481602001516014811115620000ec57fe5b105b80156200010e5750606081015160001a60f81b6001600160f81b03191615155b801562000120575060008160c0015151115b801562000132575060008160e0015151115b80156200015457506101008101515160001a60f81b6001600160f81b03191615155b6200015e57600080fd5b80516002805460ff191660018360158111156200017757fe5b021790555060208101516002805461ff0019166101008360148111156200019a57fe5b02179055506060810151600355608081015160045560a08101518051620001ca916005916020909101906200032f565b5060c08101518051620001e6916006916020909101906200032f565b5060e0810151805162000202916007916020909101906200032f565b50610100810151620002149062000217565b50565b6000546040516368a9f4e960e01b81526001600160a01b03909116906368a9f4e9906200024c9084906001906004016200068f565b600060405180830381600087803b1580156200026757600080fd5b505af11580156200027c573d6000803e3d6000fd5b505082516200028e92509050620002c9565b620002145751600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b600081811a60f81b6001600160f81b031916620002e557600080fd5b6000805b600854811015620003285783600882815481106200030357fe5b906000526020600020015414156200031f576001915062000328565b600101620002e9565b5092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200037257805160ff1916838001178555620003a2565b82800160010185558215620003a2579182015b82811115620003a257825182559160200191906001019062000385565b50620003b0929150620003b4565b5090565b5b80821115620003b05760008155600101620003b5565b80516001600160a01b0381168114620003e357600080fd5b92915050565b805160158110620003e357600080fd5b805160168110620003e357600080fd5b600082601f8301126200041a578081fd5b81516001600160401b0381111562000430578182fd5b62000445601f8201601f191660200162000713565b91508082528360208285010111156200045d57600080fd5b620003288160208401602086016200073a565b60006080828403121562000482578081fd5b6200048e608062000713565b8251815260208301519091506001600160401b0380821115620004b057600080fd5b620004be8583860162000409565b60208401526040840151915080821115620004d857600080fd5b620004e68583860162000409565b604084015260608401519150808211156200050057600080fd5b506200050f8482850162000409565b60608301525092915050565b600080604083850312156200052e578182fd5b82516001600160401b038082111562000545578384fd5b81850191506101208083880312156200055c578485fd5b620005678162000713565b9050620005758784620003f9565b8152620005868760208501620003e9565b602082015260408301516040820152606083015160608201526080830151608082015260a083015182811115620005bb578586fd5b620005c98882860162000409565b60a08301525060c083015182811115620005e1578586fd5b620005ef8882860162000409565b60c08301525060e08301518281111562000607578586fd5b620006158882860162000409565b60e08301525061010080840151838111156200062f578687fd5b6200063d8982870162000470565b828401525050809450505050620006588460208501620003cb565b90509250929050565b600081518084526200067b8160208601602086016200073a565b601f01601f19169290920160200192915050565b60006040825283516040830152602084015160806060840152620006b760c084018262000661565b90506040850151603f1980858403016080860152620006d7838362000661565b925060608701519150808584030160a086015250620006f7828262000661565b92505050600583106200070657fe5b8260208301529392505050565b6040518181016001600160401b03811182821017156200073257600080fd5b604052919050565b60005b83811015620007575781810151838201526020016200073d565b8381111562000767576000848401525b50505050565b6110a7806200077d6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80636c6071aa116100715780636c6071aa146101255780636d4ce63c1461013a5780637502f03c1461014f5780637fc622f414610162578063880a8bce1461016a578063d6d541641461017d576100a9565b80630d464d44146100ae578063240b5600146100c3578063476caa00146100ec57806349952c49146100ff57806359c0161c14610112575b600080fd5b6100c16100bc366004610d1c565b610185565b005b6100d66100d1366004610d04565b6102fe565b6040516100e39190610f04565b60405180910390f35b6100d66100fa366004610d04565b610360565b6100c161010d366004610d1c565b6103b9565b6100d6610120366004610d04565b61052f565b61012d610588565b6040516100e39190610ec0565b6101426105e0565b6040516100e39190610f9a565b6100c161015d366004610d57565b610821565b61012d61083c565b6100c1610178366004610d1c565b610892565b61012d61093e565b6000546040516368a9f4e960e01b81526001600160a01b03909116906368a9f4e9906101b8908490600290600401610f1d565b600060405180830381600087803b1580156101d257600080fd5b505af11580156101e6573d6000803e3d6000fd5b50506000546001548451604051636c3cf7f760e01b81526001600160a01b039093169450636c3cf7f7935061021d92600401610f0f565b600060405180830381600087803b15801561023757600080fd5b505af115801561024b573d6000803e3d6000fd5b50506000548351600154604051636c3cf7f760e01b81526001600160a01b039093169450636c3cf7f7935061028292600401610f0f565b600060405180830381600087803b15801561029c57600080fd5b505af11580156102b0573d6000803e3d6000fd5b505050506102c18160000151610360565b6102fb578051600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af01555b50565b600081811a60f81b6001600160f81b03191661031957600080fd5b6000805b600a548110156103575783600a828154811061033557fe5b9060005260206000200154141561034f5760019150610357565b60010161031d565b5090505b919050565b600081811a60f81b6001600160f81b03191661037b57600080fd5b6000805b60095481101561035757836009828154811061039757fe5b906000526020600020015414156103b15760019150610357565b60010161037f565b6000546040516368a9f4e960e01b81526001600160a01b03909116906368a9f4e9906103ec908490600390600401610f1d565b600060405180830381600087803b15801561040657600080fd5b505af115801561041a573d6000803e3d6000fd5b50506000546001548451604051636c3cf7f760e01b81526001600160a01b039093169450636c3cf7f7935061045192600401610f0f565b600060405180830381600087803b15801561046b57600080fd5b505af115801561047f573d6000803e3d6000fd5b50506000548351600154604051636c3cf7f760e01b81526001600160a01b039093169450636c3cf7f793506104b692600401610f0f565b600060405180830381600087803b1580156104d057600080fd5b505af11580156104e4573d6000803e3d6000fd5b505050506104f581600001516102fe565b6102fb5751600a80546001810182556000919091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80155565b600081811a60f81b6001600160f81b03191661054a57600080fd5b6000805b60085481101561035757836008828154811061056657fe5b906000526020600020015414156105805760019150610357565b60010161054e565b6060600a8054806020026020016040519081016040528092919081815260200182805480156105d657602002820191906000526020600020905b8154815260200190600101908083116105c2575b5050505050905090565b6105e8610b08565b6040805160e0810190915260028054829060ff16601581111561060757fe5b601581111561061257fe5b81528154602090910190610100900460ff16601481111561062f57fe5b601481111561063a57fe5b81526020016001820154815260200160028201548152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106eb5780601f106106c0576101008083540402835291602001916106eb565b820191906000526020600020905b8154815290600101906020018083116106ce57829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561077f5780601f106107545761010080835404028352916020019161077f565b820191906000526020600020905b81548152906001019060200180831161076257829003601f168201915b505050918352505060058201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156108135780601f106107e857610100808354040283529160200191610813565b820191906000526020600020905b8154815290600101906020018083116107f657829003601f168201915b505050505081525050905090565b80604001516001541461083357600080fd5b6102fb81610994565b606060088054806020026020016040519081016040528092919081815260200182805480156105d657602002820191906000526020600020908154815260200190600101908083116105c2575050505050905090565b6000546040516368a9f4e960e01b81526001600160a01b03909116906368a9f4e9906108c5908490600190600401610f1d565b600060405180830381600087803b1580156108df57600080fd5b505af11580156108f3573d6000803e3d6000fd5b50505050610904816000015161052f565b6102fb5751600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b606060098054806020026020016040519081016040528092919081815260200182805480156105d657602002820191906000526020600020908154815260200190600101908083116105c2575050505050905090565b6000815160158111156109a357fe5b1180156109bc57506015815160158111156109ba57fe5b105b80156109d757506000816020015160148111156109d557fe5b115b80156109f257506014816020015160148111156109f057fe5b105b8015610a115750606081015160001a60f81b6001600160f81b03191615155b8015610a22575060008160c0015151115b8015610a33575060008160e0015151115b8015610a5457506101008101515160001a60f81b6001600160f81b03191615155b610a5d57600080fd5b80516002805460ff19166001836015811115610a7557fe5b021790555060208101516002805461ff001916610100836014811115610a9757fe5b02179055506060810151600355608081015160045560a08101518051610ac591600591602090910190610b45565b5060c08101518051610adf91600691602090910190610b45565b5060e08101518051610af991600791602090910190610b45565b506102fb816101000151610892565b6040805160e08101909152806000815260200160008152600060208201819052604082015260608082018190526080820181905260a09091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b8657805160ff1916838001178555610bb3565b82800160010185558215610bb3579182015b82811115610bb3578251825591602001919060010190610b98565b50610bbf929150610bc3565b5090565b5b80821115610bbf5760008155600101610bc4565b803560158110610be757600080fd5b92915050565b803560168110610be757600080fd5b600082601f830112610c0c578081fd5b813567ffffffffffffffff811115610c22578182fd5b610c35601f8201601f1916602001611034565b9150808252836020828501011115610c4c57600080fd5b8060208401602084013760009082016020015292915050565b600060808284031215610c76578081fd5b610c806080611034565b905081358152602082013567ffffffffffffffff80821115610ca157600080fd5b610cad85838601610bfc565b60208401526040840135915080821115610cc657600080fd5b610cd285838601610bfc565b60408401526060840135915080821115610ceb57600080fd5b50610cf884828501610bfc565b60608301525092915050565b600060208284031215610d15578081fd5b5035919050565b600060208284031215610d2d578081fd5b813567ffffffffffffffff811115610d43578182fd5b610d4f84828501610c65565b949350505050565b600060208284031215610d68578081fd5b813567ffffffffffffffff80821115610d7f578283fd5b8184019150610120808387031215610d95578384fd5b610d9e81611034565b9050610daa8684610bed565b8152610db98660208501610bd8565b602082015260408301356040820152606083013560608201526080830135608082015260a083013582811115610ded578485fd5b610df987828601610bfc565b60a08301525060c083013582811115610e10578485fd5b610e1c87828601610bfc565b60c08301525060e083013582811115610e33578485fd5b610e3f87828601610bfc565b60e0830152506101008084013583811115610e58578586fd5b610e6488828701610c65565b918301919091525095945050505050565b60008151808452815b81811015610e9a57602081850181015186830182015201610e7e565b81811115610eab5782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610ef857835183529284019291840191600101610edc565b50909695505050505050565b901515815260200190565b918252602082015260400190565b60006040825283516040830152602084015160806060840152610f4360c0840182610e75565b90506040850151603f1980858403016080860152610f618383610e75565b925060608701519150808584030160a086015250610f7f8282610e75565b9250505060058310610f8d57fe5b8260208301529392505050565b600060208252610faa8351611066565b6020830152610fbc602084015161105b565b60408301526040830151606083015260608301516080830152608083015160e060a0840152610fef610100840182610e75565b905060a0840151601f19808584030160c086015261100d8383610e75565b925060c08601519150808584030160e08601525061102b8282610e75565b95945050505050565b60405181810167ffffffffffffffff8111828210171561105357600080fd5b604052919050565b806015811061035b57fe5b806016811061035b57fefea2646970667358221220292cf56a1e01c881170cdd76ae875a14deddc93a2b9e0b8e8932840562a3d5d864736f6c63430007000033a2646970667358221220db09e9cbe1cbf2757b87b1bee486517271bccc786e558ceced727027e5b4b33664736f6c63430007000033"

// DeployArtefacts deploys a new Ethereum contract, binding an instance of Artefacts to it.
func DeployArtefacts(auth *bind.TransactOpts, backend bind.ContractBackend, _entityFactory common.Address) (common.Address, *types.Transaction, *Artefacts, error) {
	parsed, err := abi.JSON(strings.NewReader(ArtefactsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArtefactsBin), backend, _entityFactory)
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
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_Artefacts *ArtefactsCaller) GetWork(opts *bind.CallOpts, _id [32]byte) (Works, error) {
	var (
		ret0 = new(Works)
	)
	out := ret0
	err := _Artefacts.contract.Call(opts, out, "getWork", _id)
	return *ret0, err
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_Artefacts *ArtefactsSession) GetWork(_id [32]byte) (Works, error) {
	return _Artefacts.Contract.GetWork(&_Artefacts.CallOpts, _id)
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_Artefacts *ArtefactsCallerSession) GetWork(_id [32]byte) (Works, error) {
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

// AddWork is a paid mutator transaction binding the contract method 0x42e9f1e0.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsTransactor) AddWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWork", _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x42e9f1e0.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWork(&_Artefacts.TransactOpts, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x42e9f1e0.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWork(&_Artefacts.TransactOpts, _work)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x5f04d064.
//
// Solidity: function addWorkAuthor(bytes32 _workId, (bytes32,string,string,string) _author) returns()
func (_Artefacts *ArtefactsTransactor) AddWorkAuthor(opts *bind.TransactOpts, _workId [32]byte, _author CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWorkAuthor", _workId, _author)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x5f04d064.
//
// Solidity: function addWorkAuthor(bytes32 _workId, (bytes32,string,string,string) _author) returns()
func (_Artefacts *ArtefactsSession) AddWorkAuthor(_workId [32]byte, _author CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkAuthor(&_Artefacts.TransactOpts, _workId, _author)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x5f04d064.
//
// Solidity: function addWorkAuthor(bytes32 _workId, (bytes32,string,string,string) _author) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWorkAuthor(_workId [32]byte, _author CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkAuthor(&_Artefacts.TransactOpts, _workId, _author)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0xa4a9051e.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, (bytes32,string,string,string) _copyrightHolder) returns()
func (_Artefacts *ArtefactsTransactor) AddWorkCopyrightHolder(opts *bind.TransactOpts, _workId [32]byte, _copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWorkCopyrightHolder", _workId, _copyrightHolder)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0xa4a9051e.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, (bytes32,string,string,string) _copyrightHolder) returns()
func (_Artefacts *ArtefactsSession) AddWorkCopyrightHolder(_workId [32]byte, _copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkCopyrightHolder(&_Artefacts.TransactOpts, _workId, _copyrightHolder)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0xa4a9051e.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, (bytes32,string,string,string) _copyrightHolder) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWorkCopyrightHolder(_workId [32]byte, _copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkCopyrightHolder(&_Artefacts.TransactOpts, _workId, _copyrightHolder)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0x0f0063d8.
//
// Solidity: function addWorkPublisher(bytes32 _workId, (bytes32,string,string,string) _publisher) returns()
func (_Artefacts *ArtefactsTransactor) AddWorkPublisher(opts *bind.TransactOpts, _workId [32]byte, _publisher CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWorkPublisher", _workId, _publisher)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0x0f0063d8.
//
// Solidity: function addWorkPublisher(bytes32 _workId, (bytes32,string,string,string) _publisher) returns()
func (_Artefacts *ArtefactsSession) AddWorkPublisher(_workId [32]byte, _publisher CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkPublisher(&_Artefacts.TransactOpts, _workId, _publisher)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0x0f0063d8.
//
// Solidity: function addWorkPublisher(bytes32 _workId, (bytes32,string,string,string) _publisher) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWorkPublisher(_workId [32]byte, _publisher CreativeEntities) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWorkPublisher(&_Artefacts.TransactOpts, _workId, _publisher)
}

// AmendWork is a paid mutator transaction binding the contract method 0x14f806c8.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsTransactor) AmendWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "amendWork", _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0x14f806c8.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AmendWork(&_Artefacts.TransactOpts, _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0x14f806c8.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsTransactorSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AmendWork(&_Artefacts.TransactOpts, _work)
}

// EntitiesABI is the input ABI used to generate the binding from.
const EntitiesABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"name\":\"addEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_childId\",\"type\":\"bytes32\"}],\"name\":\"addEntityRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amendEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_childId\",\"type\":\"bytes32\"}],\"name\":\"containsEntityRelation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntity\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReferences\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getRelations\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getRelationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRelationsReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isEntityType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntitiesFuncSigs maps the 4-byte function signature to its string representation.
var EntitiesFuncSigs = map[string]string{
	"68a9f4e9": "addEntity((bytes32,string,string,string),uint8)",
	"6c3cf7f7": "addEntityRelation(bytes32,bytes32)",
	"a81eb545": "amendEntity((bytes32,string,string,string),uint8)",
	"6e62b527": "containsEntityRelation(bytes32,bytes32)",
	"53b66f36": "getEntity(bytes32)",
	"d91ee2e0": "getEntityContract(bytes32)",
	"6ec21da9": "getEntityTypes(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"7a6337fa": "getReferences()",
	"d090cbc0": "getRelations(bytes32)",
	"5caaf3e8": "getRelationsNum(bytes32)",
	"3e30fc9b": "getRelationsReference(bytes32,uint256)",
	"cef8b362": "isEntityType(bytes32,uint8)",
}

// EntitiesBin is the compiled bytecode used for deploying new contracts.
var EntitiesBin = "0x608060405234801561001057600080fd5b5060006002556120f4806100256000396000f3fe60806040523480156200001157600080fd5b5060043610620001005760003560e01c80636ec21da91162000099578063bc599341116200006f578063bc5993411462000227578063cef8b362146200023e578063d090cbc01462000255578063d91ee2e0146200026c5762000100565b80636ec21da914620001d15780637a6337fa14620001f7578063a81eb54514620002105762000100565b806367e0badb11620000db57806367e0badb146200017157806368a9f4e9146200017b5780636c3cf7f714620001945780636e62b52714620001ab5762000100565b80633e30fc9b146200010557806353b66f3614620001345780635caaf3e8146200015a575b600080fd5b6200011c6200011636600462000da4565b62000292565b6040516200012b919062001086565b60405180910390f35b6200014b6200014536600462000d21565b62000359565b6040516200012b9190620010a4565b6200011c6200016b36600462000d21565b6200041f565b6200011c620004d0565b620001926200018c36600462000e80565b620004d6565b005b62000192620001a536600462000d53565b62000556565b620001c2620001bc36600462000d53565b620005f8565b6040516200012b91906200107b565b620001e8620001e236600462000d21565b620006b5565b6040516200012b919062000ff9565b620002016200076d565b6040516200012b919062001041565b620001926200022136600462000e80565b62000811565b6200011c6200023836600462000d21565b62000880565b620001c26200024f36600462000d75565b620008ba565b620002016200026636600462000d21565b62000923565b620002836200027d36600462000d21565b620009db565b6040516200012b919062000fe5565b6000828152602081905260408120600101546001600160a01b0316620002b757600080fd5b6000838152602081905260409081902060010154905163bc59934160e01b81526001600160a01b0390911690819063bc59934190620002fb90869060040162001086565b60206040518083038186803b1580156200031457600080fd5b505afa15801562000329573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200034f919062000d3a565b9150505b92915050565b6200036362000ac8565b6000828152602081905260409020600101546001600160a01b03166200038857600080fd5b600082815260208190526040808220600101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b158015620003d957600080fd5b505afa158015620003ee573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000418919081019062000db7565b9392505050565b6000818152602081905260408120600101546001600160a01b03166200044457600080fd5b600082815260208181526040918290206001015482516367e0badb60e01b815292516001600160a01b039091169283926367e0badb92600480840193829003018186803b1580156200049557600080fd5b505afa158015620004aa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000418919062000d3a565b60025490565b8151620004e69060009062000a1f565b15620004fe57620004f8828262000811565b62000552565b60008282604051620005109062000af3565b6200051d929190620010b9565b604051809103906000f0801580156200053a573d6000803e3d6000fd5b5083519091506200054f906000908362000a34565b50505b5050565b6000828152602081905260409020600101546001600160a01b03166200057b57600080fd5b600082815260208190526040908190206001015490516351366bd360e11b81526001600160a01b0390911690819063a26cd7a690620005bf90859060040162001086565b600060405180830381600087803b158015620005da57600080fd5b505af1158015620005ef573d6000803e3d6000fd5b50505050505050565b6000828152602081905260408120600101546001600160a01b03166200061d57600080fd5b600083815260208190526040908190206001015490516379ae21e760e11b81526001600160a01b0390911690819063f35c43ce906200066190869060040162001086565b60206040518083038186803b1580156200067a57600080fd5b505afa1580156200068f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200034f919062000cff565b6000818152602081905260409020600101546060906001600160a01b0316620006dd57600080fd5b6000828152602081905260408082206001015481516305a2bceb60e51b815291516001600160a01b0390911692839263b4579d609260048083019392829003018186803b1580156200072e57600080fd5b505afa15801562000743573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000418919081019062000bcf565b60608060006002015467ffffffffffffffff811180156200078d57600080fd5b50604051908082528060200260200182016040528015620007b8578160200160208202803683370190505b50905060005b6002548110156200080b576001805482908110620007d857fe5b906000526020600020906002020160000154828281518110620007f757fe5b6020908102919091010152600101620007be565b50905090565b81516000908152602081905260409020600101546001600160a01b03166200083857600080fd5b81516000908152602081905260409081902060010154905163d1f439e560e01b81526001600160a01b0390911690819063d1f439e590620005bf9086908690600401620010b9565b60025460009082106200089257600080fd5b6001805483908110620008a157fe5b9060005260206000209060020201600001549050919050565b6000828152602081905260408120600101546001600160a01b0316620008df57600080fd5b6000838152602081905260409081902060010154905163222f74e760e01b81526001600160a01b0390911690819063222f74e790620006619086906004016200108f565b6000818152602081905260409020600101546060906001600160a01b03166200094b57600080fd5b600082815260208190526040808220600101548151633d319bfd60e11b815291516001600160a01b03909116928392637a6337fa9260048083019392829003018186803b1580156200099c57600080fd5b505afa158015620009b1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000418919081019062000c72565b6000818152602081905260408120600101546001600160a01b031662000a0057600080fd5b506000908152602081905260409020600101546001600160a01b031690565b60009081526020919091526040902054151590565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000a7457600191505062000418565b506001808501805491820180825560008681526020889052604090208190558591908390811062000aa157fe5b60009182526020822060029182020192909255908601805460010190559150620004189050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b610f39806200118683390190565b805180151581146200035357600080fd5b8035600581106200035357600080fd5b600082601f83011262000b33578081fd5b813562000b4a62000b448262001131565b620010e8565b915080825283602082850101111562000b6257600080fd5b8060208401602084013760009082016020015292915050565b600082601f83011262000b8c578081fd5b815162000b9d62000b448262001131565b915080825283602082850101111562000bb557600080fd5b62000bc881602084016020860162001156565b5092915050565b6000602080838503121562000be2578182fd5b825167ffffffffffffffff81111562000bf9578283fd5b8301601f8101851362000c0a578283fd5b805162000c1b62000b448262001110565b818152838101908385018584028501860189101562000c38578687fd5b8694505b8385101562000c665762000c51898262000b01565b83526001949094019391850191850162000c3c565b50979650505050505050565b6000602080838503121562000c85578182fd5b825167ffffffffffffffff81111562000c9c578283fd5b8301601f8101851362000cad578283fd5b805162000cbe62000b448262001110565b818152838101908385018584028501860189101562000cdb578687fd5b8694505b8385101562000c6657805183526001949094019391850191850162000cdf565b60006020828403121562000d11578081fd5b8151801515811462000418578182fd5b60006020828403121562000d33578081fd5b5035919050565b60006020828403121562000d4c578081fd5b5051919050565b6000806040838503121562000d66578081fd5b50508035926020909101359150565b6000806040838503121562000d88578182fd5b8235915062000d9b846020850162000b12565b90509250929050565b6000806040838503121562000d66578182fd5b60006020828403121562000dc9578081fd5b815167ffffffffffffffff8082111562000de1578283fd5b908301906080828603121562000df5578283fd5b62000e016080620010e8565b8251815260208301518281111562000e17578485fd5b62000e258782860162000b7b565b60208301525060408301518281111562000e3d578485fd5b62000e4b8782860162000b7b565b60408301525060608301518281111562000e63578485fd5b62000e718782860162000b7b565b60608301525095945050505050565b6000806040838503121562000e93578182fd5b823567ffffffffffffffff8082111562000eab578384fd5b908401906080828703121562000ebf578384fd5b62000ecb6080620010e8565b8235815260208301358281111562000ee1578586fd5b62000eef8882860162000b22565b60208301525060408301358281111562000f07578586fd5b62000f158882860162000b22565b60408301525060608301358281111562000f2d578586fd5b62000f3b8882860162000b22565b60608301525080945050505062000d9b846020850162000b12565b6000815180845262000f7081602086016020860162001156565b601f01601f19169290920160200192915050565b60008151835260208201516080602085015262000fa5608085018262000f56565b90506040830151848203604086015262000fc0828262000f56565b9150506060830151848203606086015262000fdc828262000f56565b95945050505050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b818110156200103557835115158352928401929184019160010162001015565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101562001035578351835292840192918401916001016200105d565b901515815260200190565b90815260200190565b60208101600583106200109e57fe5b91905290565b60006020825262000418602083018462000f84565b600060408252620010ce604083018562000f84565b905060058310620010db57fe5b8260208301529392505050565b60405181810167ffffffffffffffff811182821017156200110857600080fd5b604052919050565b600067ffffffffffffffff82111562001127578081fd5b5060209081020190565b600067ffffffffffffffff82111562001148578081fd5b50601f01601f191660200190565b60005b838110156200117357818101518382015260200162001159565b838111156200054f575050600091015256fe60806040523480156200001157600080fd5b5060405162000f3938038062000f398339810160408190526200003491620003a8565b815160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007257506000826020015151115b80156200008457506000826040015151115b80156200009d575060008160048111156200009b57fe5b115b8015620000b657506004816004811115620000b457fe5b105b620000c057600080fd5b60408051600480825260a08201909252906020820160808036833750508151620000f2926004925060200190620001a4565b50600160048260048111156200010457fe5b60ff16815481106200011257fe5b6000918252602080832081830401805460ff601f9094166101000a9384021916941515929092029390931790558351815583820151805185936200015c9260019291019062000250565b50604082015180516200017a91600284019160209091019062000250565b50606082015180516200019891600384019160209091019062000250565b509050505050620004ad565b82805482825590600052602060002090601f016020900481019282156200023e5791602002820160005b838211156200020d57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302620001ce565b80156200023c5782816101000a81549060ff02191690556001016020816000010492830192600103026200020d565b505b506200024c929150620002d1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200029357805160ff1916838001178555620002c3565b82800160010185558215620002c3579182015b82811115620002c3578251825591602001919060010190620002a6565b506200024c929150620002ec565b5b808211156200024c57805460ff19168155600101620002d2565b5b808211156200024c5760008155600101620002ed565b8051600581106200031357600080fd5b92915050565b600082601f8301126200032a578081fd5b81516001600160401b0381111562000340578182fd5b602062000356601f8301601f1916820162000486565b925081835284818386010111156200036d57600080fd5b60005b828110156200038d57848101820151848201830152810162000370565b828111156200039f5760008284860101525b50505092915050565b60008060408385031215620003bb578182fd5b82516001600160401b0380821115620003d2578384fd5b9084019060808287031215620003e6578384fd5b620003f2608062000486565b8251815260208301518281111562000408578586fd5b620004168882860162000319565b6020830152506040830151828111156200042e578586fd5b6200043c8882860162000319565b60408301525060608301518281111562000454578586fd5b620004628882860162000319565b6060830152508094505050506200047d846020850162000303565b90509250929050565b6040518181016001600160401b0381118282101715620004a557600080fd5b604052919050565b610a7c80620004bd6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b4579d6011610066578063b4579d6014610120578063bc59934114610135578063d1f439e514610148578063e188130b1461015b578063f35c43ce1461016e5761009e565b8063222f74e7146100a357806367e0badb146100cc5780636d4ce63c146100e15780637a6337fa146100f6578063a26cd7a61461010b575b600080fd5b6100b66100b13660046107e2565b610181565b6040516100c391906109a0565b60405180910390f35b6100d46101f3565b6040516100c391906109ab565b6100e96101f9565b6040516100c391906109b4565b6100fe6103d1565b6040516100c39190610968565b61011e6101193660046107ca565b610429565b005b610128610499565b6040516100c39190610922565b6100d46101433660046107ca565b610510565b61011e610156366004610808565b610540565b61011e6101693660046107e2565b6105a9565b6100b661017c3660046107ca565b61062e565b60008082600481111561019057fe5b1180156101a8575060048260048111156101a657fe5b105b6101b157600080fd5b60048260048111156101bf57fe5b60ff16815481106101cc57fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b60055490565b61020161068e565b604080516080810182526000805482526001805484516020600283851615610100026000190190931692909204601f8101839004830282018301909652858152939492938186019390929183018282801561029d5780601f106102725761010080835404028352916020019161029d565b820191906000526020600020905b81548152906001019060200180831161028057829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561032f5780601f106103045761010080835404028352916020019161032f565b820191906000526020600020905b81548152906001019060200180831161031257829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103c35780601f10610398576101008083540402835291602001916103c3565b820191906000526020600020905b8154815290600101906020018083116103a657829003601f168201915b505050505081525050905090565b6060600580548060200260200160405190810160405280929190818152602001828054801561041f57602002820191906000526020600020905b81548152602001906001019080831161040b575b5050505050905090565b8060001a60f81b6001600160f81b03191661044357600080fd5b600054811480159061045b57506104598161062e565b155b1561049657600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018190555b50565b6060600480548060200260200160405190810160405280929190818152602001828054801561041f57602002820191906000526020600020906000905b825461010083900a900460ff1615158152602060019283018181049485019490930390920291018084116104d65790505050505050905090565b600554600090821061052157600080fd5b6005828154811061052e57fe5b90600052602060002001549050919050565b610549816105a9565b81516000908155602080840151805185939261056a926001929101906106b9565b50604082015180516105869160028401916020909101906106b9565b50606082015180516105a29160038401916020909101906106b9565b5050505050565b60008160048111156105b757fe5b1180156105cf575060048160048111156105cd57fe5b105b6105d857600080fd5b6105e181610181565b61049657600160048260048111156105f557fe5b60ff168154811061060257fe5b90600052602060002090602091828204019190066101000a81548160ff02191690831515021790555050565b600081811a60f81b6001600160f81b03191661064957600080fd5b6000805b60055481101561068757836005828154811061066557fe5b9060005260206000200154141561067f5760019150610687565b60010161064d565b5092915050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106fa57805160ff1916838001178555610727565b82800160010185558215610727579182015b8281111561072757825182559160200191906001019061070c565b50610733929150610737565b5090565b5b808211156107335760008155600101610738565b80356005811061075b57600080fd5b92915050565b600082601f830112610771578081fd5b813567ffffffffffffffff811115610787578182fd5b61079a601f8201601f1916602001610a1f565b91508082528360208285010111156107b157600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156107db578081fd5b5035919050565b6000602082840312156107f3578081fd5b813560058110610801578182fd5b9392505050565b6000806040838503121561081a578081fd5b823567ffffffffffffffff80821115610831578283fd5b9084019060808287031215610844578283fd5b61084e6080610a1f565b82358152602083013582811115610863578485fd5b61086f88828601610761565b602083015250604083013582811115610886578485fd5b61089288828601610761565b6040830152506060830135828111156108a9578485fd5b6108b588828601610761565b6060830152508094505050506108ce846020850161074c565b90509250929050565b60008151808452815b818110156108fc576020818501810151868301820152016108e0565b8181111561090d5782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b8181101561095c57835115158352928401929184019160010161093e565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561095c57835183529284019291840191600101610984565b901515815260200190565b90815260200190565b600060208252825160208301526020830151608060408401526109da60a08401826108d7565b90506040840151601f19808584030160608601526109f883836108d7565b9250606086015191508085840301608086015250610a1682826108d7565b95945050505050565b60405181810167ffffffffffffffff81118282101715610a3e57600080fd5b60405291905056fea2646970667358221220f92eb1360c908b75f4dc4ac1c5c7a481d6ebae9cdcee440cca9ed18514353ea164736f6c63430007000033a26469706673582212200c4a8078fbe0cafd9b89127bf40a3ea32bb26a59b6aab4361725875018cd1d0464736f6c63430007000033"

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

// ContainsEntityRelation is a free data retrieval call binding the contract method 0x6e62b527.
//
// Solidity: function containsEntityRelation(bytes32 _parentId, bytes32 _childId) view returns(bool)
func (_Entities *EntitiesCaller) ContainsEntityRelation(opts *bind.CallOpts, _parentId [32]byte, _childId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "containsEntityRelation", _parentId, _childId)
	return *ret0, err
}

// ContainsEntityRelation is a free data retrieval call binding the contract method 0x6e62b527.
//
// Solidity: function containsEntityRelation(bytes32 _parentId, bytes32 _childId) view returns(bool)
func (_Entities *EntitiesSession) ContainsEntityRelation(_parentId [32]byte, _childId [32]byte) (bool, error) {
	return _Entities.Contract.ContainsEntityRelation(&_Entities.CallOpts, _parentId, _childId)
}

// ContainsEntityRelation is a free data retrieval call binding the contract method 0x6e62b527.
//
// Solidity: function containsEntityRelation(bytes32 _parentId, bytes32 _childId) view returns(bool)
func (_Entities *EntitiesCallerSession) ContainsEntityRelation(_parentId [32]byte, _childId [32]byte) (bool, error) {
	return _Entities.Contract.ContainsEntityRelation(&_Entities.CallOpts, _parentId, _childId)
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

// GetRelations is a free data retrieval call binding the contract method 0xd090cbc0.
//
// Solidity: function getRelations(bytes32 _id) view returns(bytes32[])
func (_Entities *EntitiesCaller) GetRelations(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getRelations", _id)
	return *ret0, err
}

// GetRelations is a free data retrieval call binding the contract method 0xd090cbc0.
//
// Solidity: function getRelations(bytes32 _id) view returns(bytes32[])
func (_Entities *EntitiesSession) GetRelations(_id [32]byte) ([][32]byte, error) {
	return _Entities.Contract.GetRelations(&_Entities.CallOpts, _id)
}

// GetRelations is a free data retrieval call binding the contract method 0xd090cbc0.
//
// Solidity: function getRelations(bytes32 _id) view returns(bytes32[])
func (_Entities *EntitiesCallerSession) GetRelations(_id [32]byte) ([][32]byte, error) {
	return _Entities.Contract.GetRelations(&_Entities.CallOpts, _id)
}

// GetRelationsNum is a free data retrieval call binding the contract method 0x5caaf3e8.
//
// Solidity: function getRelationsNum(bytes32 _id) view returns(uint256)
func (_Entities *EntitiesCaller) GetRelationsNum(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getRelationsNum", _id)
	return *ret0, err
}

// GetRelationsNum is a free data retrieval call binding the contract method 0x5caaf3e8.
//
// Solidity: function getRelationsNum(bytes32 _id) view returns(uint256)
func (_Entities *EntitiesSession) GetRelationsNum(_id [32]byte) (*big.Int, error) {
	return _Entities.Contract.GetRelationsNum(&_Entities.CallOpts, _id)
}

// GetRelationsNum is a free data retrieval call binding the contract method 0x5caaf3e8.
//
// Solidity: function getRelationsNum(bytes32 _id) view returns(uint256)
func (_Entities *EntitiesCallerSession) GetRelationsNum(_id [32]byte) (*big.Int, error) {
	return _Entities.Contract.GetRelationsNum(&_Entities.CallOpts, _id)
}

// GetRelationsReference is a free data retrieval call binding the contract method 0x3e30fc9b.
//
// Solidity: function getRelationsReference(bytes32 _id, uint256 _index) view returns(bytes32)
func (_Entities *EntitiesCaller) GetRelationsReference(opts *bind.CallOpts, _id [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getRelationsReference", _id, _index)
	return *ret0, err
}

// GetRelationsReference is a free data retrieval call binding the contract method 0x3e30fc9b.
//
// Solidity: function getRelationsReference(bytes32 _id, uint256 _index) view returns(bytes32)
func (_Entities *EntitiesSession) GetRelationsReference(_id [32]byte, _index *big.Int) ([32]byte, error) {
	return _Entities.Contract.GetRelationsReference(&_Entities.CallOpts, _id, _index)
}

// GetRelationsReference is a free data retrieval call binding the contract method 0x3e30fc9b.
//
// Solidity: function getRelationsReference(bytes32 _id, uint256 _index) view returns(bytes32)
func (_Entities *EntitiesCallerSession) GetRelationsReference(_id [32]byte, _index *big.Int) ([32]byte, error) {
	return _Entities.Contract.GetRelationsReference(&_Entities.CallOpts, _id, _index)
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
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _childId) returns()
func (_Entities *EntitiesTransactor) AddEntityRelation(opts *bind.TransactOpts, _parentId [32]byte, _childId [32]byte) (*types.Transaction, error) {
	return _Entities.contract.Transact(opts, "addEntityRelation", _parentId, _childId)
}

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _childId) returns()
func (_Entities *EntitiesSession) AddEntityRelation(_parentId [32]byte, _childId [32]byte) (*types.Transaction, error) {
	return _Entities.Contract.AddEntityRelation(&_Entities.TransactOpts, _parentId, _childId)
}

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _childId) returns()
func (_Entities *EntitiesTransactorSession) AddEntityRelation(_parentId [32]byte, _childId [32]byte) (*types.Transaction, error) {
	return _Entities.Contract.AddEntityRelation(&_Entities.TransactOpts, _parentId, _childId)
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
const EntityNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"addRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"containsRelation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReferences\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntityNodeFuncSigs maps the 4-byte function signature to its string representation.
var EntityNodeFuncSigs = map[string]string{
	"a26cd7a6": "addRelation(bytes32)",
	"e188130b": "addType(uint8)",
	"d1f439e5": "amend((bytes32,string,string,string),uint8)",
	"f35c43ce": "containsRelation(bytes32)",
	"6d4ce63c": "get()",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"7a6337fa": "getReferences()",
	"b4579d60": "getTypes()",
	"222f74e7": "isType(uint8)",
}

// EntityNodeBin is the compiled bytecode used for deploying new contracts.
var EntityNodeBin = "0x60806040523480156200001157600080fd5b5060405162000f3938038062000f398339810160408190526200003491620003a8565b815160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007257506000826020015151115b80156200008457506000826040015151115b80156200009d575060008160048111156200009b57fe5b115b8015620000b657506004816004811115620000b457fe5b105b620000c057600080fd5b60408051600480825260a08201909252906020820160808036833750508151620000f2926004925060200190620001a4565b50600160048260048111156200010457fe5b60ff16815481106200011257fe5b6000918252602080832081830401805460ff601f9094166101000a9384021916941515929092029390931790558351815583820151805185936200015c9260019291019062000250565b50604082015180516200017a91600284019160209091019062000250565b50606082015180516200019891600384019160209091019062000250565b509050505050620004ad565b82805482825590600052602060002090601f016020900481019282156200023e5791602002820160005b838211156200020d57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302620001ce565b80156200023c5782816101000a81549060ff02191690556001016020816000010492830192600103026200020d565b505b506200024c929150620002d1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200029357805160ff1916838001178555620002c3565b82800160010185558215620002c3579182015b82811115620002c3578251825591602001919060010190620002a6565b506200024c929150620002ec565b5b808211156200024c57805460ff19168155600101620002d2565b5b808211156200024c5760008155600101620002ed565b8051600581106200031357600080fd5b92915050565b600082601f8301126200032a578081fd5b81516001600160401b0381111562000340578182fd5b602062000356601f8301601f1916820162000486565b925081835284818386010111156200036d57600080fd5b60005b828110156200038d57848101820151848201830152810162000370565b828111156200039f5760008284860101525b50505092915050565b60008060408385031215620003bb578182fd5b82516001600160401b0380821115620003d2578384fd5b9084019060808287031215620003e6578384fd5b620003f2608062000486565b8251815260208301518281111562000408578586fd5b620004168882860162000319565b6020830152506040830151828111156200042e578586fd5b6200043c8882860162000319565b60408301525060608301518281111562000454578586fd5b620004628882860162000319565b6060830152508094505050506200047d846020850162000303565b90509250929050565b6040518181016001600160401b0381118282101715620004a557600080fd5b604052919050565b610a7c80620004bd6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b4579d6011610066578063b4579d6014610120578063bc59934114610135578063d1f439e514610148578063e188130b1461015b578063f35c43ce1461016e5761009e565b8063222f74e7146100a357806367e0badb146100cc5780636d4ce63c146100e15780637a6337fa146100f6578063a26cd7a61461010b575b600080fd5b6100b66100b13660046107e2565b610181565b6040516100c391906109a0565b60405180910390f35b6100d46101f3565b6040516100c391906109ab565b6100e96101f9565b6040516100c391906109b4565b6100fe6103d1565b6040516100c39190610968565b61011e6101193660046107ca565b610429565b005b610128610499565b6040516100c39190610922565b6100d46101433660046107ca565b610510565b61011e610156366004610808565b610540565b61011e6101693660046107e2565b6105a9565b6100b661017c3660046107ca565b61062e565b60008082600481111561019057fe5b1180156101a8575060048260048111156101a657fe5b105b6101b157600080fd5b60048260048111156101bf57fe5b60ff16815481106101cc57fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b60055490565b61020161068e565b604080516080810182526000805482526001805484516020600283851615610100026000190190931692909204601f8101839004830282018301909652858152939492938186019390929183018282801561029d5780601f106102725761010080835404028352916020019161029d565b820191906000526020600020905b81548152906001019060200180831161028057829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561032f5780601f106103045761010080835404028352916020019161032f565b820191906000526020600020905b81548152906001019060200180831161031257829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103c35780601f10610398576101008083540402835291602001916103c3565b820191906000526020600020905b8154815290600101906020018083116103a657829003601f168201915b505050505081525050905090565b6060600580548060200260200160405190810160405280929190818152602001828054801561041f57602002820191906000526020600020905b81548152602001906001019080831161040b575b5050505050905090565b8060001a60f81b6001600160f81b03191661044357600080fd5b600054811480159061045b57506104598161062e565b155b1561049657600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018190555b50565b6060600480548060200260200160405190810160405280929190818152602001828054801561041f57602002820191906000526020600020906000905b825461010083900a900460ff1615158152602060019283018181049485019490930390920291018084116104d65790505050505050905090565b600554600090821061052157600080fd5b6005828154811061052e57fe5b90600052602060002001549050919050565b610549816105a9565b81516000908155602080840151805185939261056a926001929101906106b9565b50604082015180516105869160028401916020909101906106b9565b50606082015180516105a29160038401916020909101906106b9565b5050505050565b60008160048111156105b757fe5b1180156105cf575060048160048111156105cd57fe5b105b6105d857600080fd5b6105e181610181565b61049657600160048260048111156105f557fe5b60ff168154811061060257fe5b90600052602060002090602091828204019190066101000a81548160ff02191690831515021790555050565b600081811a60f81b6001600160f81b03191661064957600080fd5b6000805b60055481101561068757836005828154811061066557fe5b9060005260206000200154141561067f5760019150610687565b60010161064d565b5092915050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106fa57805160ff1916838001178555610727565b82800160010185558215610727579182015b8281111561072757825182559160200191906001019061070c565b50610733929150610737565b5090565b5b808211156107335760008155600101610738565b80356005811061075b57600080fd5b92915050565b600082601f830112610771578081fd5b813567ffffffffffffffff811115610787578182fd5b61079a601f8201601f1916602001610a1f565b91508082528360208285010111156107b157600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156107db578081fd5b5035919050565b6000602082840312156107f3578081fd5b813560058110610801578182fd5b9392505050565b6000806040838503121561081a578081fd5b823567ffffffffffffffff80821115610831578283fd5b9084019060808287031215610844578283fd5b61084e6080610a1f565b82358152602083013582811115610863578485fd5b61086f88828601610761565b602083015250604083013582811115610886578485fd5b61089288828601610761565b6040830152506060830135828111156108a9578485fd5b6108b588828601610761565b6060830152508094505050506108ce846020850161074c565b90509250929050565b60008151808452815b818110156108fc576020818501810151868301820152016108e0565b8181111561090d5782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b8181101561095c57835115158352928401929184019160010161093e565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561095c57835183529284019291840191600101610984565b901515815260200190565b90815260200190565b600060208252825160208301526020830151608060408401526109da60a08401826108d7565b90506040840151601f19808584030160608601526109f883836108d7565b9250606086015191508085840301608086015250610a1682826108d7565b95945050505050565b60405181810167ffffffffffffffff81118282101715610a3e57600080fd5b60405291905056fea2646970667358221220f92eb1360c908b75f4dc4ac1c5c7a481d6ebae9cdcee440cca9ed18514353ea164736f6c63430007000033"

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

// ContainsRelation is a free data retrieval call binding the contract method 0xf35c43ce.
//
// Solidity: function containsRelation(bytes32 _id) view returns(bool)
func (_EntityNode *EntityNodeCaller) ContainsRelation(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "containsRelation", _id)
	return *ret0, err
}

// ContainsRelation is a free data retrieval call binding the contract method 0xf35c43ce.
//
// Solidity: function containsRelation(bytes32 _id) view returns(bool)
func (_EntityNode *EntityNodeSession) ContainsRelation(_id [32]byte) (bool, error) {
	return _EntityNode.Contract.ContainsRelation(&_EntityNode.CallOpts, _id)
}

// ContainsRelation is a free data retrieval call binding the contract method 0xf35c43ce.
//
// Solidity: function containsRelation(bytes32 _id) view returns(bool)
func (_EntityNode *EntityNodeCallerSession) ContainsRelation(_id [32]byte) (bool, error) {
	return _EntityNode.Contract.ContainsRelation(&_EntityNode.CallOpts, _id)
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

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_EntityNode *EntityNodeCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_EntityNode *EntityNodeSession) GetNum() (*big.Int, error) {
	return _EntityNode.Contract.GetNum(&_EntityNode.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_EntityNode *EntityNodeCallerSession) GetNum() (*big.Int, error) {
	return _EntityNode.Contract.GetNum(&_EntityNode.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_EntityNode *EntityNodeCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_EntityNode *EntityNodeSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _EntityNode.Contract.GetReference(&_EntityNode.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_EntityNode *EntityNodeCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _EntityNode.Contract.GetReference(&_EntityNode.CallOpts, _index)
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_EntityNode *EntityNodeCaller) GetReferences(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "getReferences")
	return *ret0, err
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_EntityNode *EntityNodeSession) GetReferences() ([][32]byte, error) {
	return _EntityNode.Contract.GetReferences(&_EntityNode.CallOpts)
}

// GetReferences is a free data retrieval call binding the contract method 0x7a6337fa.
//
// Solidity: function getReferences() view returns(bytes32[])
func (_EntityNode *EntityNodeCallerSession) GetReferences() ([][32]byte, error) {
	return _EntityNode.Contract.GetReferences(&_EntityNode.CallOpts)
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

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _id) returns()
func (_EntityNode *EntityNodeTransactor) AddRelation(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _EntityNode.contract.Transact(opts, "addRelation", _id)
}

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _id) returns()
func (_EntityNode *EntityNodeSession) AddRelation(_id [32]byte) (*types.Transaction, error) {
	return _EntityNode.Contract.AddRelation(&_EntityNode.TransactOpts, _id)
}

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _id) returns()
func (_EntityNode *EntityNodeTransactorSession) AddRelation(_id [32]byte) (*types.Transaction, error) {
	return _EntityNode.Contract.AddRelation(&_EntityNode.TransactOpts, _id)
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
const IArtefactABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_author\",\"type\":\"tuple\"}],\"name\":\"addAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_copyrightHolder\",\"type\":\"tuple\"}],\"name\":\"addCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_publisher\",\"type\":\"tuple\"}],\"name\":\"addPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthors\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCopyrightHolders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublishers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isCopyrightHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"isPublisher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IArtefactFuncSigs maps the 4-byte function signature to its string representation.
var IArtefactFuncSigs = map[string]string{
	"880a8bce": "addAuthor((bytes32,string,string,string))",
	"0d464d44": "addCopyrightHolder((bytes32,string,string,string))",
	"49952c49": "addPublisher((bytes32,string,string,string))",
	"7502f03c": "amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)))",
	"6d4ce63c": "get()",
	"7fc622f4": "getAuthors()",
	"d6d54164": "getCopyrightHolders()",
	"6c6071aa": "getPublishers()",
	"59c0161c": "isAuthor(bytes32)",
	"476caa00": "isCopyrightHolder(bytes32)",
	"240b5600": "isPublisher(bytes32)",
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
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_IArtefact *IArtefactCaller) Get(opts *bind.CallOpts) (Works, error) {
	var (
		ret0 = new(Works)
	)
	out := ret0
	err := _IArtefact.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_IArtefact *IArtefactSession) Get() (Works, error) {
	return _IArtefact.Contract.Get(&_IArtefact.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_IArtefact *IArtefactCallerSession) Get() (Works, error) {
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

// AddAuthor is a paid mutator transaction binding the contract method 0x880a8bce.
//
// Solidity: function addAuthor((bytes32,string,string,string) _author) returns()
func (_IArtefact *IArtefactTransactor) AddAuthor(opts *bind.TransactOpts, _author CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "addAuthor", _author)
}

// AddAuthor is a paid mutator transaction binding the contract method 0x880a8bce.
//
// Solidity: function addAuthor((bytes32,string,string,string) _author) returns()
func (_IArtefact *IArtefactSession) AddAuthor(_author CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.Contract.AddAuthor(&_IArtefact.TransactOpts, _author)
}

// AddAuthor is a paid mutator transaction binding the contract method 0x880a8bce.
//
// Solidity: function addAuthor((bytes32,string,string,string) _author) returns()
func (_IArtefact *IArtefactTransactorSession) AddAuthor(_author CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.Contract.AddAuthor(&_IArtefact.TransactOpts, _author)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0x0d464d44.
//
// Solidity: function addCopyrightHolder((bytes32,string,string,string) _copyrightHolder) returns()
func (_IArtefact *IArtefactTransactor) AddCopyrightHolder(opts *bind.TransactOpts, _copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "addCopyrightHolder", _copyrightHolder)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0x0d464d44.
//
// Solidity: function addCopyrightHolder((bytes32,string,string,string) _copyrightHolder) returns()
func (_IArtefact *IArtefactSession) AddCopyrightHolder(_copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.Contract.AddCopyrightHolder(&_IArtefact.TransactOpts, _copyrightHolder)
}

// AddCopyrightHolder is a paid mutator transaction binding the contract method 0x0d464d44.
//
// Solidity: function addCopyrightHolder((bytes32,string,string,string) _copyrightHolder) returns()
func (_IArtefact *IArtefactTransactorSession) AddCopyrightHolder(_copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.Contract.AddCopyrightHolder(&_IArtefact.TransactOpts, _copyrightHolder)
}

// AddPublisher is a paid mutator transaction binding the contract method 0x49952c49.
//
// Solidity: function addPublisher((bytes32,string,string,string) _publisher) returns()
func (_IArtefact *IArtefactTransactor) AddPublisher(opts *bind.TransactOpts, _publisher CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "addPublisher", _publisher)
}

// AddPublisher is a paid mutator transaction binding the contract method 0x49952c49.
//
// Solidity: function addPublisher((bytes32,string,string,string) _publisher) returns()
func (_IArtefact *IArtefactSession) AddPublisher(_publisher CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.Contract.AddPublisher(&_IArtefact.TransactOpts, _publisher)
}

// AddPublisher is a paid mutator transaction binding the contract method 0x49952c49.
//
// Solidity: function addPublisher((bytes32,string,string,string) _publisher) returns()
func (_IArtefact *IArtefactTransactorSession) AddPublisher(_publisher CreativeEntities) (*types.Transaction, error) {
	return _IArtefact.Contract.AddPublisher(&_IArtefact.TransactOpts, _publisher)
}

// Amend is a paid mutator transaction binding the contract method 0x7502f03c.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefact *IArtefactTransactor) Amend(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "amend", _work)
}

// Amend is a paid mutator transaction binding the contract method 0x7502f03c.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefact *IArtefactSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.Contract.Amend(&_IArtefact.TransactOpts, _work)
}

// Amend is a paid mutator transaction binding the contract method 0x7502f03c.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefact *IArtefactTransactorSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.Contract.Amend(&_IArtefact.TransactOpts, _work)
}

// IArtefactFactoryABI is the input ABI used to generate the binding from.
const IArtefactFactoryABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_author\",\"type\":\"tuple\"}],\"name\":\"addWorkAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_copyrightHolder\",\"type\":\"tuple\"}],\"name\":\"addWorkCopyrightHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_publisher\",\"type\":\"tuple\"}],\"name\":\"addWorkPublisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amendWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWork\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkAuthors\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkCopyrightHolders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkPublishers\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_authorId\",\"type\":\"bytes32\"}],\"name\":\"isWorkAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_copyrightHolderId\",\"type\":\"bytes32\"}],\"name\":\"isWorkCopyrightHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_workId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publisherId\",\"type\":\"bytes32\"}],\"name\":\"isWorkPublisher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IArtefactFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IArtefactFactoryFuncSigs = map[string]string{
	"42e9f1e0": "addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)))",
	"5f04d064": "addWorkAuthor(bytes32,(bytes32,string,string,string))",
	"a4a9051e": "addWorkCopyrightHolder(bytes32,(bytes32,string,string,string))",
	"0f0063d8": "addWorkPublisher(bytes32,(bytes32,string,string,string))",
	"14f806c8": "amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)))",
	"30266537": "getWork(bytes32)",
	"426aa803": "getWorkAuthors(bytes32)",
	"4d2522a2": "getWorkContract(bytes32)",
	"86105c65": "getWorkCopyrightHolders(bytes32)",
	"5ffd09d8": "getWorkPublishers(bytes32)",
	"b2e459cf": "isWorkAuthor(bytes32,bytes32)",
	"f6a50965": "isWorkCopyrightHolder(bytes32,bytes32)",
	"ef90d9c9": "isWorkPublisher(bytes32,bytes32)",
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
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_IArtefactFactory *IArtefactFactoryCaller) GetWork(opts *bind.CallOpts, _id [32]byte) (Works, error) {
	var (
		ret0 = new(Works)
	)
	out := ret0
	err := _IArtefactFactory.contract.Call(opts, out, "getWork", _id)
	return *ret0, err
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_IArtefactFactory *IArtefactFactorySession) GetWork(_id [32]byte) (Works, error) {
	return _IArtefactFactory.Contract.GetWork(&_IArtefactFactory.CallOpts, _id)
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,string,string,string))
func (_IArtefactFactory *IArtefactFactoryCallerSession) GetWork(_id [32]byte) (Works, error) {
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

// AddWork is a paid mutator transaction binding the contract method 0x42e9f1e0.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWork", _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x42e9f1e0.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWork(&_IArtefactFactory.TransactOpts, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x42e9f1e0.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWork(&_IArtefactFactory.TransactOpts, _work)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x5f04d064.
//
// Solidity: function addWorkAuthor(bytes32 _workId, (bytes32,string,string,string) _author) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWorkAuthor(opts *bind.TransactOpts, _workId [32]byte, _author CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWorkAuthor", _workId, _author)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x5f04d064.
//
// Solidity: function addWorkAuthor(bytes32 _workId, (bytes32,string,string,string) _author) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWorkAuthor(_workId [32]byte, _author CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkAuthor(&_IArtefactFactory.TransactOpts, _workId, _author)
}

// AddWorkAuthor is a paid mutator transaction binding the contract method 0x5f04d064.
//
// Solidity: function addWorkAuthor(bytes32 _workId, (bytes32,string,string,string) _author) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWorkAuthor(_workId [32]byte, _author CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkAuthor(&_IArtefactFactory.TransactOpts, _workId, _author)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0xa4a9051e.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, (bytes32,string,string,string) _copyrightHolder) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWorkCopyrightHolder(opts *bind.TransactOpts, _workId [32]byte, _copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWorkCopyrightHolder", _workId, _copyrightHolder)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0xa4a9051e.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, (bytes32,string,string,string) _copyrightHolder) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWorkCopyrightHolder(_workId [32]byte, _copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkCopyrightHolder(&_IArtefactFactory.TransactOpts, _workId, _copyrightHolder)
}

// AddWorkCopyrightHolder is a paid mutator transaction binding the contract method 0xa4a9051e.
//
// Solidity: function addWorkCopyrightHolder(bytes32 _workId, (bytes32,string,string,string) _copyrightHolder) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWorkCopyrightHolder(_workId [32]byte, _copyrightHolder CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkCopyrightHolder(&_IArtefactFactory.TransactOpts, _workId, _copyrightHolder)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0x0f0063d8.
//
// Solidity: function addWorkPublisher(bytes32 _workId, (bytes32,string,string,string) _publisher) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWorkPublisher(opts *bind.TransactOpts, _workId [32]byte, _publisher CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWorkPublisher", _workId, _publisher)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0x0f0063d8.
//
// Solidity: function addWorkPublisher(bytes32 _workId, (bytes32,string,string,string) _publisher) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWorkPublisher(_workId [32]byte, _publisher CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkPublisher(&_IArtefactFactory.TransactOpts, _workId, _publisher)
}

// AddWorkPublisher is a paid mutator transaction binding the contract method 0x0f0063d8.
//
// Solidity: function addWorkPublisher(bytes32 _workId, (bytes32,string,string,string) _publisher) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWorkPublisher(_workId [32]byte, _publisher CreativeEntities) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWorkPublisher(&_IArtefactFactory.TransactOpts, _workId, _publisher)
}

// AmendWork is a paid mutator transaction binding the contract method 0x14f806c8.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AmendWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "amendWork", _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0x14f806c8.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactorySession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AmendWork(&_IArtefactFactory.TransactOpts, _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0x14f806c8.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AmendWork(&_IArtefactFactory.TransactOpts, _work)
}

// IEntitiesFactoryABI is the input ABI used to generate the binding from.
const IEntitiesFactoryABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_childId\",\"type\":\"bytes32\"}],\"name\":\"addEntityRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amendEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_childId\",\"type\":\"bytes32\"}],\"name\":\"containsEntityRelation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntity\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getRelations\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getRelationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRelationsReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isEntityType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEntitiesFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IEntitiesFactoryFuncSigs = map[string]string{
	"68a9f4e9": "addEntity((bytes32,string,string,string),uint8)",
	"6c3cf7f7": "addEntityRelation(bytes32,bytes32)",
	"a81eb545": "amendEntity((bytes32,string,string,string),uint8)",
	"6e62b527": "containsEntityRelation(bytes32,bytes32)",
	"53b66f36": "getEntity(bytes32)",
	"d91ee2e0": "getEntityContract(bytes32)",
	"6ec21da9": "getEntityTypes(bytes32)",
	"d090cbc0": "getRelations(bytes32)",
	"5caaf3e8": "getRelationsNum(bytes32)",
	"3e30fc9b": "getRelationsReference(bytes32,uint256)",
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

// ContainsEntityRelation is a free data retrieval call binding the contract method 0x6e62b527.
//
// Solidity: function containsEntityRelation(bytes32 _parentId, bytes32 _childId) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactoryCaller) ContainsEntityRelation(opts *bind.CallOpts, _parentId [32]byte, _childId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "containsEntityRelation", _parentId, _childId)
	return *ret0, err
}

// ContainsEntityRelation is a free data retrieval call binding the contract method 0x6e62b527.
//
// Solidity: function containsEntityRelation(bytes32 _parentId, bytes32 _childId) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactorySession) ContainsEntityRelation(_parentId [32]byte, _childId [32]byte) (bool, error) {
	return _IEntitiesFactory.Contract.ContainsEntityRelation(&_IEntitiesFactory.CallOpts, _parentId, _childId)
}

// ContainsEntityRelation is a free data retrieval call binding the contract method 0x6e62b527.
//
// Solidity: function containsEntityRelation(bytes32 _parentId, bytes32 _childId) view returns(bool)
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) ContainsEntityRelation(_parentId [32]byte, _childId [32]byte) (bool, error) {
	return _IEntitiesFactory.Contract.ContainsEntityRelation(&_IEntitiesFactory.CallOpts, _parentId, _childId)
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

// GetRelations is a free data retrieval call binding the contract method 0xd090cbc0.
//
// Solidity: function getRelations(bytes32 _id) view returns(bytes32[])
func (_IEntitiesFactory *IEntitiesFactoryCaller) GetRelations(opts *bind.CallOpts, _id [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "getRelations", _id)
	return *ret0, err
}

// GetRelations is a free data retrieval call binding the contract method 0xd090cbc0.
//
// Solidity: function getRelations(bytes32 _id) view returns(bytes32[])
func (_IEntitiesFactory *IEntitiesFactorySession) GetRelations(_id [32]byte) ([][32]byte, error) {
	return _IEntitiesFactory.Contract.GetRelations(&_IEntitiesFactory.CallOpts, _id)
}

// GetRelations is a free data retrieval call binding the contract method 0xd090cbc0.
//
// Solidity: function getRelations(bytes32 _id) view returns(bytes32[])
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) GetRelations(_id [32]byte) ([][32]byte, error) {
	return _IEntitiesFactory.Contract.GetRelations(&_IEntitiesFactory.CallOpts, _id)
}

// GetRelationsNum is a free data retrieval call binding the contract method 0x5caaf3e8.
//
// Solidity: function getRelationsNum(bytes32 _id) view returns(uint256)
func (_IEntitiesFactory *IEntitiesFactoryCaller) GetRelationsNum(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "getRelationsNum", _id)
	return *ret0, err
}

// GetRelationsNum is a free data retrieval call binding the contract method 0x5caaf3e8.
//
// Solidity: function getRelationsNum(bytes32 _id) view returns(uint256)
func (_IEntitiesFactory *IEntitiesFactorySession) GetRelationsNum(_id [32]byte) (*big.Int, error) {
	return _IEntitiesFactory.Contract.GetRelationsNum(&_IEntitiesFactory.CallOpts, _id)
}

// GetRelationsNum is a free data retrieval call binding the contract method 0x5caaf3e8.
//
// Solidity: function getRelationsNum(bytes32 _id) view returns(uint256)
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) GetRelationsNum(_id [32]byte) (*big.Int, error) {
	return _IEntitiesFactory.Contract.GetRelationsNum(&_IEntitiesFactory.CallOpts, _id)
}

// GetRelationsReference is a free data retrieval call binding the contract method 0x3e30fc9b.
//
// Solidity: function getRelationsReference(bytes32 _id, uint256 _index) view returns(bytes32)
func (_IEntitiesFactory *IEntitiesFactoryCaller) GetRelationsReference(opts *bind.CallOpts, _id [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "getRelationsReference", _id, _index)
	return *ret0, err
}

// GetRelationsReference is a free data retrieval call binding the contract method 0x3e30fc9b.
//
// Solidity: function getRelationsReference(bytes32 _id, uint256 _index) view returns(bytes32)
func (_IEntitiesFactory *IEntitiesFactorySession) GetRelationsReference(_id [32]byte, _index *big.Int) ([32]byte, error) {
	return _IEntitiesFactory.Contract.GetRelationsReference(&_IEntitiesFactory.CallOpts, _id, _index)
}

// GetRelationsReference is a free data retrieval call binding the contract method 0x3e30fc9b.
//
// Solidity: function getRelationsReference(bytes32 _id, uint256 _index) view returns(bytes32)
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) GetRelationsReference(_id [32]byte, _index *big.Int) ([32]byte, error) {
	return _IEntitiesFactory.Contract.GetRelationsReference(&_IEntitiesFactory.CallOpts, _id, _index)
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
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _childId) returns()
func (_IEntitiesFactory *IEntitiesFactoryTransactor) AddEntityRelation(opts *bind.TransactOpts, _parentId [32]byte, _childId [32]byte) (*types.Transaction, error) {
	return _IEntitiesFactory.contract.Transact(opts, "addEntityRelation", _parentId, _childId)
}

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _childId) returns()
func (_IEntitiesFactory *IEntitiesFactorySession) AddEntityRelation(_parentId [32]byte, _childId [32]byte) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.AddEntityRelation(&_IEntitiesFactory.TransactOpts, _parentId, _childId)
}

// AddEntityRelation is a paid mutator transaction binding the contract method 0x6c3cf7f7.
//
// Solidity: function addEntityRelation(bytes32 _parentId, bytes32 _childId) returns()
func (_IEntitiesFactory *IEntitiesFactoryTransactorSession) AddEntityRelation(_parentId [32]byte, _childId [32]byte) (*types.Transaction, error) {
	return _IEntitiesFactory.Contract.AddEntityRelation(&_IEntitiesFactory.TransactOpts, _parentId, _childId)
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
const IEntityABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"addRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"containsRelation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTypes\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"isType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEntityFuncSigs maps the 4-byte function signature to its string representation.
var IEntityFuncSigs = map[string]string{
	"a26cd7a6": "addRelation(bytes32)",
	"e188130b": "addType(uint8)",
	"d1f439e5": "amend((bytes32,string,string,string),uint8)",
	"f35c43ce": "containsRelation(bytes32)",
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

// ContainsRelation is a free data retrieval call binding the contract method 0xf35c43ce.
//
// Solidity: function containsRelation(bytes32 _id) view returns(bool)
func (_IEntity *IEntityCaller) ContainsRelation(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEntity.contract.Call(opts, out, "containsRelation", _id)
	return *ret0, err
}

// ContainsRelation is a free data retrieval call binding the contract method 0xf35c43ce.
//
// Solidity: function containsRelation(bytes32 _id) view returns(bool)
func (_IEntity *IEntitySession) ContainsRelation(_id [32]byte) (bool, error) {
	return _IEntity.Contract.ContainsRelation(&_IEntity.CallOpts, _id)
}

// ContainsRelation is a free data retrieval call binding the contract method 0xf35c43ce.
//
// Solidity: function containsRelation(bytes32 _id) view returns(bool)
func (_IEntity *IEntityCallerSession) ContainsRelation(_id [32]byte) (bool, error) {
	return _IEntity.Contract.ContainsRelation(&_IEntity.CallOpts, _id)
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

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _id) returns()
func (_IEntity *IEntityTransactor) AddRelation(opts *bind.TransactOpts, _id [32]byte) (*types.Transaction, error) {
	return _IEntity.contract.Transact(opts, "addRelation", _id)
}

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _id) returns()
func (_IEntity *IEntitySession) AddRelation(_id [32]byte) (*types.Transaction, error) {
	return _IEntity.Contract.AddRelation(&_IEntity.TransactOpts, _id)
}

// AddRelation is a paid mutator transaction binding the contract method 0xa26cd7a6.
//
// Solidity: function addRelation(bytes32 _id) returns()
func (_IEntity *IEntityTransactorSession) AddRelation(_id [32]byte) (*types.Transaction, error) {
	return _IEntity.Contract.AddRelation(&_IEntity.TransactOpts, _id)
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
var IterableDataBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209ae6fd29a3cb6dc0da5b558e6aa30f697081dedd4dd5c556138a19f60965655964736f6c63430007000033"

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
