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
	WorkType        uint8
	License         uint8
	Id              [32]byte
	DateCreated     [32]byte
	DateModified    [32]byte
	Url             string
	Name            string
	Description     string
	Author          CreativeEntities
	CopyrightHolder CreativeEntities
	Publisher       CreativeEntities
}

// ArtefactNodeABI is the input ABI used to generate the binding from.
const ArtefactNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_entityFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArtefactNodeFuncSigs maps the 4-byte function signature to its string representation.
var ArtefactNodeFuncSigs = map[string]string{
	"adef5f6b": "amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))",
	"6d4ce63c": "get()",
}

// ArtefactNodeBin is the compiled bytecode used for deploying new contracts.
var ArtefactNodeBin = "0x60806040523480156200001157600080fd5b5060405162001ae838038062001ae883398101604081905262000034916200059d565b604082015160001a60f81b6001600160f81b0319166200005357600080fd5b600080546001600160a01b0319166001600160a01b0383161790556040820151600155620000818262000089565b505062000841565b6000815160158111156200009957fe5b118015620000b45750601581516015811115620000b257fe5b105b8015620000d15750600081602001516014811115620000cf57fe5b115b8015620000ee5750601481602001516014811115620000ec57fe5b105b80156200010e5750606081015160001a60f81b6001600160f81b03191615155b801562000120575060008160c0015151115b801562000132575060008160e0015151115b80156200015457506101008101515160001a60f81b6001600160f81b03191615155b6200015e57600080fd5b80516002805460ff191660018360158111156200017757fe5b021790555060208101516002805461ff0019166101008360148111156200019a57fe5b02179055506060810151600355608081015160045561010081015151600555610120810151516006556101408101515160075560a08101518051620001e891600891602090910190620003a1565b5060c081015180516200020491600991602090910190620003a1565b5060e081015180516200022091600a91602090910190620003a1565b506000546101008201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e9916200025a9160019060040162000764565b600060405180830381600087803b1580156200027557600080fd5b505af11580156200028a573d6000803e3d6000fd5b50506006549150600090501a60f81b6001600160f81b0319161562000317576000546101208201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e991620002e29160029060040162000764565b600060405180830381600087803b158015620002fd57600080fd5b505af115801562000312573d6000803e3d6000fd5b505050505b60075460001a60f81b6001600160f81b031916156200039e576000546101408201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e991620003699160039060040162000764565b600060405180830381600087803b1580156200038457600080fd5b505af115801562000399573d6000803e3d6000fd5b505050505b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003e457805160ff191683800117855562000414565b8280016001018555821562000414579182015b8281111562000414578251825591602001919060010190620003f7565b506200042292915062000426565b5090565b6200044391905b808211156200042257600081556001016200042d565b90565b80516001600160a01b03811681146200045e57600080fd5b92915050565b8051601581106200045e57600080fd5b8051601681106200045e57600080fd5b600082601f83011262000495578081fd5b81516001600160401b03811115620004ab578182fd5b620004c0601f8201601f1916602001620007e7565b9150808252836020828501011115620004d857600080fd5b620004eb8160208401602086016200080e565b5092915050565b60006080828403121562000504578081fd5b620005106080620007e7565b8251815260208301519091506001600160401b03808211156200053257600080fd5b620005408583860162000484565b602084015260408401519150808211156200055a57600080fd5b620005688583860162000484565b604084015260608401519150808211156200058257600080fd5b50620005918482850162000484565b60608301525092915050565b60008060408385031215620005b0578182fd5b82516001600160401b0380821115620005c7578384fd5b610160918501808703831315620005dc578485fd5b620005e783620007e7565b620005f3888362000474565b815262000604886020840162000464565b602082015260408201516040820152606082015160608201526080820151608082015260a08201519350828411156200063b578586fd5b620006498885840162000484565b60a082015260c082015193508284111562000662578586fd5b620006708885840162000484565b60c082015260e082015193508284111562000689578586fd5b620006978885840162000484565b60e082015261010093508382015183811115620006b2578687fd5b620006c089828501620004f2565b858301525061012093508382015183811115620006db578687fd5b620006e989828501620004f2565b85830152506101409350838201518381111562000704578687fd5b6200071289828501620004f2565b8583015250809550505050506200072d846020850162000446565b90509250929050565b60008151808452620007508160208601602086016200080e565b601f01601f19169290920160200192915050565b600060408252835160408301526020840151608060608401526200078c60c084018262000736565b60408601519150603f1980858303016080860152620007ac828462000736565b60608801519350818682030160a0870152620007c9818562000736565b94505050505060058310620007da57fe5b8260208301529392505050565b6040518181016001600160401b03811182821017156200080657600080fd5b604052919050565b60005b838110156200082b57818101518382015260200162000811565b838111156200083b576000848401525b50505050565b61129780620008516000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063adef5f6b14610059575b600080fd5b61004361006e565b60405161005091906110e3565b60405180910390f35b61006c610067366004610ea4565b6107f3565b005b610076610b06565b61007e610b06565b600254819060ff16601581111561009157fe5b9081601581111561009e57fe5b9052506002546020820190610100900460ff1660148111156100bc57fe5b908160148111156100c957fe5b90525060018054604083810191909152600354606084015260045460808401526008805482516020600295831615610100026000190190921694909404601f810182900482028501820190935282845290919083018282801561016d5780601f106101425761010080835404028352916020019161016d565b820191906000526020600020905b81548152906001019060200180831161015057829003601f168201915b505050505060a08201526009805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101fd5780601f106101d2576101008083540402835291602001916101fd565b820191906000526020600020905b8154815290600101906020018083116101e057829003601f168201915b505050505060c0820152600a805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561028d5780601f106102625761010080835404028352916020019161028d565b820191906000526020600020905b81548152906001019060200180831161027057829003601f168201915b505050505060e0820152600080546005546040516306c8f71760e51b81526001600160a01b039092169163d91ee2e0916102c9916004016110ae565b60206040518083038186803b1580156102e157600080fd5b505afa1580156102f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103199190610d9a565b90506001816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561035657600080fd5b505afa15801561036a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038e9190610dc8565b600481111561039957fe5b146103a357600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b1580156103dc57600080fd5b505afa1580156103f0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104189190810190610de7565b61010083015260065460001a60f81b6001600160f81b031916156105c1576000546006546040516306c8f71760e51b81526001600160a01b039092169163d91ee2e091610467916004016110ae565b60206040518083038186803b15801561047f57600080fd5b505afa158015610493573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b79190610d9a565b90506002816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156104f457600080fd5b505afa158015610508573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061052c9190610dc8565b600481111561053757fe5b1461054157600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b15801561057a57600080fd5b505afa15801561058e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105b69190810190610de7565b610120830152610605565b610120820180516000908190526040805160208082018352838252845181019190915281518082018352838152845183015281519081019091529081529051606001525b60075460001a60f81b6001600160f81b031916156107a8576000546007546040516306c8f71760e51b81526001600160a01b039092169163d91ee2e09161064e916004016110ae565b60206040518083038186803b15801561066657600080fd5b505afa15801561067a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069e9190610d9a565b90506003816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156106db57600080fd5b505afa1580156106ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107139190610dc8565b600481111561071e57fe5b1461072857600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b15801561076157600080fd5b505afa158015610775573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261079d9190810190610de7565b6101408301526107ec565b610140820180516000908190526040805160208082018352838252845181019190915281518082018352838152845183015281519081019091529081529051606001525b5090505b90565b80604001516001541461080557600080fd5b61080e81610811565b50565b60008151601581111561082057fe5b118015610839575060158151601581111561083757fe5b105b8015610854575060008160200151601481111561085257fe5b115b801561086f575060148160200151601481111561086d57fe5b105b801561088e5750606081015160001a60f81b6001600160f81b03191615155b801561089f575060008160c0015151115b80156108b0575060008160e0015151115b80156108d157506101008101515160001a60f81b6001600160f81b03191615155b6108da57600080fd5b80516002805460ff191660018360158111156108f257fe5b021790555060208101516002805461ff00191661010083601481111561091457fe5b02179055506060810151600355608081015160045561010081015151600555610120810151516006556101408101515160075560a0810151805161096091600891602090910190610b74565b5060c0810151805161097a91600991602090910190610b74565b5060e0810151805161099491600a91602090910190610b74565b506000546101008201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e9916109cc916001906004016110b7565b600060405180830381600087803b1580156109e657600080fd5b505af11580156109fa573d6000803e3d6000fd5b50506006549150600090501a60f81b6001600160f81b03191615610a82576000546101208201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e991610a4f916002906004016110b7565b600060405180830381600087803b158015610a6957600080fd5b505af1158015610a7d573d6000803e3d6000fd5b505050505b60075460001a60f81b6001600160f81b0319161561080e576000546101408201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e991610ad1916003906004016110b7565b600060405180830381600087803b158015610aeb57600080fd5b505af1158015610aff573d6000803e3d6000fd5b5050505050565b604080516101608101909152806000815260200160008152600060208201819052604082018190526060808301919091526080820181905260a0820181905260c082015260e001610b55610bf2565b8152602001610b62610bf2565b8152602001610b6f610bf2565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bb557805160ff1916838001178555610be2565b82800160010185558215610be2579182015b82811115610be2578251825591602001919060010190610bc7565b50610bee929150610c1d565b5090565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b6107f091905b80821115610bee5760008155600101610c23565b803560158110610c4657600080fd5b92915050565b803560168110610c4657600080fd5b600082601f830112610c6b578081fd5b8135610c7e610c798261120d565b6111e6565b9150808252836020828501011115610c9557600080fd5b8060208401602084013760009082016020015292915050565b600082601f830112610cbe578081fd5b8151610ccc610c798261120d565b9150808252836020828501011115610ce357600080fd5b610cf4816020840160208601611231565b5092915050565b600060808284031215610d0c578081fd5b610d1660806111e6565b905081358152602082013567ffffffffffffffff80821115610d3757600080fd5b610d4385838601610c5b565b60208401526040840135915080821115610d5c57600080fd5b610d6885838601610c5b565b60408401526060840135915080821115610d8157600080fd5b50610d8e84828501610c5b565b60608301525092915050565b600060208284031215610dab578081fd5b81516001600160a01b0381168114610dc1578182fd5b9392505050565b600060208284031215610dd9578081fd5b815160058110610dc1578182fd5b600060208284031215610df8578081fd5b815167ffffffffffffffff80821115610e0f578283fd5b81840160808187031215610e21578384fd5b610e2b60806111e6565b925080518352602081015182811115610e42578485fd5b610e4e87828401610cae565b602085015250604081015182811115610e65578485fd5b610e7187828401610cae565b604085015250606081015182811115610e88578485fd5b610e9487828401610cae565b6060850152509195945050505050565b600060208284031215610eb5578081fd5b813567ffffffffffffffff80821115610ecc578283fd5b610160918401808603831315610ee0578384fd5b610ee9836111e6565b610ef38783610c4c565b8152610f028760208401610c37565b602082015260408201356040820152606082013560608201526080820135608082015260a0820135935082841115610f38578485fd5b610f4487858401610c5b565b60a082015260c0820135935082841115610f5c578485fd5b610f6887858401610c5b565b60c082015260e0820135935082841115610f80578485fd5b610f8c87858401610c5b565b60e082015261010093508382013583811115610fa6578586fd5b610fb288828501610cfb565b858301525061012093508382013583811115610fcc578586fd5b610fd888828501610cfb565b858301525061014093508382013583811115610ff2578586fd5b610ffe88828501610cfb565b948201949094529695505050505050565b6015811061101957fe5b9052565b6016811061101957fe5b6000815180845261103f816020860160208601611231565b601f01601f19169290920160200192915050565b6000815183526020820151608060208501526110726080850182611027565b60408401519150848103604086015261108b8183611027565b6060850151925085810360608701526110a48184611027565b9695505050505050565b90815260200190565b6000604082526110ca6040830185611053565b9050600583106110d657fe5b8260208301529392505050565b6000602082526110f760208301845161101d565b6020830151611109604084018261100f565b506040830151606083015260608301516080830152608083015160a083015260a08301516101608060c0850152611144610180850183611027565b60c08601519250601f19808683030160e08701526111628285611027565b60e0880151945061010092508187820301838801526111818186611027565b838901519550610120935082888203018489015261119f8187611053565b91505082880151945061014092508187820301838801526111c08186611053565b8389015195508288820301858901526111d98187611053565b9998505050505050505050565b60405181810167ffffffffffffffff8111828210171561120557600080fd5b604052919050565b600067ffffffffffffffff821115611223578081fd5b50601f01601f191660200190565b60005b8381101561124c578181015183820152602001611234565b8381111561125b576000848401525b5050505056fea2646970667358221220661583b4b216ff3830102a4dc0f6ac670628addaf3e1336204ef50a49716567564736f6c63430006080033"

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
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
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
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
func (_ArtefactNode *ArtefactNodeSession) Get() (CreativeWorks, error) {
	return _ArtefactNode.Contract.Get(&_ArtefactNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
func (_ArtefactNode *ArtefactNodeCallerSession) Get() (CreativeWorks, error) {
	return _ArtefactNode.Contract.Get(&_ArtefactNode.CallOpts)
}

// Amend is a paid mutator transaction binding the contract method 0xadef5f6b.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_ArtefactNode *ArtefactNodeTransactor) Amend(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.contract.Transact(opts, "amend", _work)
}

// Amend is a paid mutator transaction binding the contract method 0xadef5f6b.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_ArtefactNode *ArtefactNodeSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.Contract.Amend(&_ArtefactNode.TransactOpts, _work)
}

// Amend is a paid mutator transaction binding the contract method 0xadef5f6b.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_ArtefactNode *ArtefactNodeTransactorSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _ArtefactNode.Contract.Amend(&_ArtefactNode.TransactOpts, _work)
}

// ArtefactsABI is the input ABI used to generate the binding from.
const ArtefactsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entityFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amendWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWork\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArtefactsFuncSigs maps the 4-byte function signature to its string representation.
var ArtefactsFuncSigs = map[string]string{
	"5ed47d96": "addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))",
	"f8a3a2e0": "amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"30266537": "getWork(bytes32)",
	"4d2522a2": "getWorkContract(bytes32)",
}

// ArtefactsBin is the compiled bytecode used for deploying new contracts.
var ArtefactsBin = "0x608060405234801561001057600080fd5b506040516128aa3803806128aa83398101604081905261002f9161006a565b6001600160a01b03811661004257600080fd5b600080546001600160a01b0319166001600160a01b0392909216919091178155600355610098565b60006020828403121561007b578081fd5b81516001600160a01b0381168114610091578182fd5b9392505050565b612803806100a76000396000f3fe60806040523480156200001157600080fd5b50600436106200006a5760003560e01c806330266537146200006f5780634d2522a2146200009e5780635ed47d9614620000c457806367e0badb14620000dd578063bc59934114620000f6578063f8a3a2e0146200010d575b600080fd5b62000086620000803660046200073c565b62000124565b60405162000095919062000c08565b60405180910390f35b620000b5620000af3660046200073c565b620001ed565b60405162000095919062000beb565b620000db620000d536600462000755565b62000232565b005b620000e7620002ce565b60405162000095919062000bff565b620000e7620001073660046200073c565b620002d4565b620000db6200011e36600462000755565b6200030e565b6200012e62000464565b600082815260016020819052604090912001546001600160a01b03166200015457600080fd5b6000828152600160208190526040808320909101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b158015620001a757600080fd5b505afa158015620001bc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620001e69190810190620008db565b9392505050565b6000818152600160208190526040822001546001600160a01b03166200021257600080fd5b50600090815260016020819052604090912001546001600160a01b031690565b60408101516200024b9060019063ffffffff620003b816565b1562000262576200025c816200030e565b620002cb565b6000805460405183916001600160a01b0316906200028090620004d8565b6200028d92919062000c1d565b604051809103906000f080158015620002aa573d6000803e3d6000fd5b506040830151909150620002c8906001908363ffffffff620003d016565b50505b50565b60035490565b6003546000908210620002e657600080fd5b6002805483908110620002f557fe5b9060005260206000209060020201600001549050919050565b60408082015160009081526001602081905291902001546001600160a01b03166200033857600080fd5b604080820151600090815260016020819052908290200154905163adef5f6b60e01b81526001600160a01b0390911690819063adef5f6b906200038090859060040162000c08565b600060405180830381600087803b1580156200039b57600080fd5b505af1158015620003b0573d6000803e3d6000fd5b505050505050565b60008181526020839052604090205415155b92915050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000410576001915050620001e6565b50600180850180549182018082556000868152602088905260409020819055859190839081106200043d57fe5b60009182526020822060029182020192909255908601805460010190559150620001e69050565b604080516101608101909152806000815260200160008152600060208201819052604082018190526060808301919091526080820181905260a0820181905260c082015260e001620004b5620004e6565b8152602001620004c4620004e6565b8152602001620004d3620004e6565b905290565b611ae88062000ce683390190565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b8035620003ca8162000cc9565b8051620003ca8162000cc9565b8035620003ca8162000cd7565b8051620003ca8162000cd7565b600082601f83011262000556578081fd5b81356200056d620005678262000c71565b62000c49565b91508082528360208285010111156200058557600080fd5b8060208401602084013760009082016020015292915050565b600082601f830112620005af578081fd5b8151620005c0620005678262000c71565b9150808252836020828501011115620005d857600080fd5b620005eb81602084016020860162000c96565b5092915050565b60006080828403121562000604578081fd5b62000610608062000c49565b905081358152602082013567ffffffffffffffff808211156200063257600080fd5b620006408583860162000545565b602084015260408401359150808211156200065a57600080fd5b620006688583860162000545565b604084015260608401359150808211156200068257600080fd5b50620006918482850162000545565b60608301525092915050565b600060808284031215620006af578081fd5b620006bb608062000c49565b905081518152602082015167ffffffffffffffff80821115620006dd57600080fd5b620006eb858386016200059e565b602084015260408401519150808211156200070557600080fd5b62000713858386016200059e565b604084015260608401519150808211156200072d57600080fd5b5062000691848285016200059e565b6000602082840312156200074e578081fd5b5035919050565b60006020828403121562000767578081fd5b813567ffffffffffffffff808211156200077f578283fd5b61016091840180860383131562000794578384fd5b6200079f8362000c49565b620007ab87836200052b565b8152620007bc876020840162000511565b602082015260408201356040820152606082013560608201526080820135608082015260a0820135935082841115620007f3578485fd5b620008018785840162000545565b60a082015260c08201359350828411156200081a578485fd5b620008288785840162000545565b60c082015260e082013593508284111562000841578485fd5b6200084f8785840162000545565b60e0820152610100935083820135838111156200086a578586fd5b6200087888828501620005f2565b85830152506101209350838201358381111562000893578586fd5b620008a188828501620005f2565b858301525061014093508382013583811115620008bc578586fd5b620008ca88828501620005f2565b948201949094529695505050505050565b600060208284031215620008ed578081fd5b815167ffffffffffffffff8082111562000905578283fd5b6101609184018086038313156200091a578384fd5b620009258362000c49565b62000931878362000538565b81526200094287602084016200051e565b602082015260408201516040820152606082015160608201526080820151608082015260a082015193508284111562000979578485fd5b62000987878584016200059e565b60a082015260c0820151935082841115620009a0578485fd5b620009ae878584016200059e565b60c082015260e0820151935082841115620009c7578485fd5b620009d5878584016200059e565b60e082015261010093508382015183811115620009f0578586fd5b620009fe888285016200069d565b85830152506101209350838201518381111562000a19578586fd5b62000a27888285016200069d565b85830152506101409350838201518381111562000a42578586fd5b620008ca888285016200069d565b6015811062000a5b57fe5b9052565b6016811062000a5b57fe5b6000815180845262000a8481602086016020860162000c96565b601f01601f19169290920160200192915050565b60008151835260208201516080602085015262000ab9608085018262000a6a565b60408401519150848103604086015262000ad4818362000a6a565b60608501519250858103606087015262000aef818462000a6a565b9695505050505050565b600061016062000b0b84845162000a5f565b602083015162000b1f602086018262000a50565b5060408301516040850152606083015160608501526080830151608085015260a08301518160a086015262000b578286018262000a6a565b60c0850151925085810360c087015262000b72818462000a6a565b91505060e0840151915084810360e086015262000b90818362000a6a565b61010092508285015191508581038387015262000bae818362000a98565b92505050610120808401518583038287015262000bcc838262000a98565b9150506101409150818401518582038387015262000aef828262000a98565b6001600160a01b0391909116815260200190565b90815260200190565b600060208252620001e6602083018462000af9565b60006040825262000c32604083018562000af9565b905060018060a01b03831660208301529392505050565b60405181810167ffffffffffffffff8111828210171562000c6957600080fd5b604052919050565b600067ffffffffffffffff82111562000c88578081fd5b50601f01601f191660200190565b60005b8381101562000cb357818101518382015260200162000c99565b8381111562000cc3576000848401525b50505050565b60158110620002cb57600080fd5b60168110620002cb57600080fdfe60806040523480156200001157600080fd5b5060405162001ae838038062001ae883398101604081905262000034916200059d565b604082015160001a60f81b6001600160f81b0319166200005357600080fd5b600080546001600160a01b0319166001600160a01b0383161790556040820151600155620000818262000089565b505062000841565b6000815160158111156200009957fe5b118015620000b45750601581516015811115620000b257fe5b105b8015620000d15750600081602001516014811115620000cf57fe5b115b8015620000ee5750601481602001516014811115620000ec57fe5b105b80156200010e5750606081015160001a60f81b6001600160f81b03191615155b801562000120575060008160c0015151115b801562000132575060008160e0015151115b80156200015457506101008101515160001a60f81b6001600160f81b03191615155b6200015e57600080fd5b80516002805460ff191660018360158111156200017757fe5b021790555060208101516002805461ff0019166101008360148111156200019a57fe5b02179055506060810151600355608081015160045561010081015151600555610120810151516006556101408101515160075560a08101518051620001e891600891602090910190620003a1565b5060c081015180516200020491600991602090910190620003a1565b5060e081015180516200022091600a91602090910190620003a1565b506000546101008201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e9916200025a9160019060040162000764565b600060405180830381600087803b1580156200027557600080fd5b505af11580156200028a573d6000803e3d6000fd5b50506006549150600090501a60f81b6001600160f81b0319161562000317576000546101208201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e991620002e29160029060040162000764565b600060405180830381600087803b158015620002fd57600080fd5b505af115801562000312573d6000803e3d6000fd5b505050505b60075460001a60f81b6001600160f81b031916156200039e576000546101408201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e991620003699160039060040162000764565b600060405180830381600087803b1580156200038457600080fd5b505af115801562000399573d6000803e3d6000fd5b505050505b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003e457805160ff191683800117855562000414565b8280016001018555821562000414579182015b8281111562000414578251825591602001919060010190620003f7565b506200042292915062000426565b5090565b6200044391905b808211156200042257600081556001016200042d565b90565b80516001600160a01b03811681146200045e57600080fd5b92915050565b8051601581106200045e57600080fd5b8051601681106200045e57600080fd5b600082601f83011262000495578081fd5b81516001600160401b03811115620004ab578182fd5b620004c0601f8201601f1916602001620007e7565b9150808252836020828501011115620004d857600080fd5b620004eb8160208401602086016200080e565b5092915050565b60006080828403121562000504578081fd5b620005106080620007e7565b8251815260208301519091506001600160401b03808211156200053257600080fd5b620005408583860162000484565b602084015260408401519150808211156200055a57600080fd5b620005688583860162000484565b604084015260608401519150808211156200058257600080fd5b50620005918482850162000484565b60608301525092915050565b60008060408385031215620005b0578182fd5b82516001600160401b0380821115620005c7578384fd5b610160918501808703831315620005dc578485fd5b620005e783620007e7565b620005f3888362000474565b815262000604886020840162000464565b602082015260408201516040820152606082015160608201526080820151608082015260a08201519350828411156200063b578586fd5b620006498885840162000484565b60a082015260c082015193508284111562000662578586fd5b620006708885840162000484565b60c082015260e082015193508284111562000689578586fd5b620006978885840162000484565b60e082015261010093508382015183811115620006b2578687fd5b620006c089828501620004f2565b858301525061012093508382015183811115620006db578687fd5b620006e989828501620004f2565b85830152506101409350838201518381111562000704578687fd5b6200071289828501620004f2565b8583015250809550505050506200072d846020850162000446565b90509250929050565b60008151808452620007508160208601602086016200080e565b601f01601f19169290920160200192915050565b600060408252835160408301526020840151608060608401526200078c60c084018262000736565b60408601519150603f1980858303016080860152620007ac828462000736565b60608801519350818682030160a0870152620007c9818562000736565b94505050505060058310620007da57fe5b8260208301529392505050565b6040518181016001600160401b03811182821017156200080657600080fd5b604052919050565b60005b838110156200082b57818101518382015260200162000811565b838111156200083b576000848401525b50505050565b61129780620008516000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063adef5f6b14610059575b600080fd5b61004361006e565b60405161005091906110e3565b60405180910390f35b61006c610067366004610ea4565b6107f3565b005b610076610b06565b61007e610b06565b600254819060ff16601581111561009157fe5b9081601581111561009e57fe5b9052506002546020820190610100900460ff1660148111156100bc57fe5b908160148111156100c957fe5b90525060018054604083810191909152600354606084015260045460808401526008805482516020600295831615610100026000190190921694909404601f810182900482028501820190935282845290919083018282801561016d5780601f106101425761010080835404028352916020019161016d565b820191906000526020600020905b81548152906001019060200180831161015057829003601f168201915b505050505060a08201526009805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101fd5780601f106101d2576101008083540402835291602001916101fd565b820191906000526020600020905b8154815290600101906020018083116101e057829003601f168201915b505050505060c0820152600a805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561028d5780601f106102625761010080835404028352916020019161028d565b820191906000526020600020905b81548152906001019060200180831161027057829003601f168201915b505050505060e0820152600080546005546040516306c8f71760e51b81526001600160a01b039092169163d91ee2e0916102c9916004016110ae565b60206040518083038186803b1580156102e157600080fd5b505afa1580156102f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103199190610d9a565b90506001816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561035657600080fd5b505afa15801561036a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038e9190610dc8565b600481111561039957fe5b146103a357600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b1580156103dc57600080fd5b505afa1580156103f0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104189190810190610de7565b61010083015260065460001a60f81b6001600160f81b031916156105c1576000546006546040516306c8f71760e51b81526001600160a01b039092169163d91ee2e091610467916004016110ae565b60206040518083038186803b15801561047f57600080fd5b505afa158015610493573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b79190610d9a565b90506002816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156104f457600080fd5b505afa158015610508573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061052c9190610dc8565b600481111561053757fe5b1461054157600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b15801561057a57600080fd5b505afa15801561058e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105b69190810190610de7565b610120830152610605565b610120820180516000908190526040805160208082018352838252845181019190915281518082018352838152845183015281519081019091529081529051606001525b60075460001a60f81b6001600160f81b031916156107a8576000546007546040516306c8f71760e51b81526001600160a01b039092169163d91ee2e09161064e916004016110ae565b60206040518083038186803b15801561066657600080fd5b505afa15801561067a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069e9190610d9a565b90506003816001600160a01b03166315dae03e6040518163ffffffff1660e01b815260040160206040518083038186803b1580156106db57600080fd5b505afa1580156106ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107139190610dc8565b600481111561071e57fe5b1461072857600080fd5b806001600160a01b0316636d4ce63c6040518163ffffffff1660e01b815260040160006040518083038186803b15801561076157600080fd5b505afa158015610775573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261079d9190810190610de7565b6101408301526107ec565b610140820180516000908190526040805160208082018352838252845181019190915281518082018352838152845183015281519081019091529081529051606001525b5090505b90565b80604001516001541461080557600080fd5b61080e81610811565b50565b60008151601581111561082057fe5b118015610839575060158151601581111561083757fe5b105b8015610854575060008160200151601481111561085257fe5b115b801561086f575060148160200151601481111561086d57fe5b105b801561088e5750606081015160001a60f81b6001600160f81b03191615155b801561089f575060008160c0015151115b80156108b0575060008160e0015151115b80156108d157506101008101515160001a60f81b6001600160f81b03191615155b6108da57600080fd5b80516002805460ff191660018360158111156108f257fe5b021790555060208101516002805461ff00191661010083601481111561091457fe5b02179055506060810151600355608081015160045561010081015151600555610120810151516006556101408101515160075560a0810151805161096091600891602090910190610b74565b5060c0810151805161097a91600991602090910190610b74565b5060e0810151805161099491600a91602090910190610b74565b506000546101008201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e9916109cc916001906004016110b7565b600060405180830381600087803b1580156109e657600080fd5b505af11580156109fa573d6000803e3d6000fd5b50506006549150600090501a60f81b6001600160f81b03191615610a82576000546101208201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e991610a4f916002906004016110b7565b600060405180830381600087803b158015610a6957600080fd5b505af1158015610a7d573d6000803e3d6000fd5b505050505b60075460001a60f81b6001600160f81b0319161561080e576000546101408201516040516368a9f4e960e01b81526001600160a01b03909216916368a9f4e991610ad1916003906004016110b7565b600060405180830381600087803b158015610aeb57600080fd5b505af1158015610aff573d6000803e3d6000fd5b5050505050565b604080516101608101909152806000815260200160008152600060208201819052604082018190526060808301919091526080820181905260a0820181905260c082015260e001610b55610bf2565b8152602001610b62610bf2565b8152602001610b6f610bf2565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bb557805160ff1916838001178555610be2565b82800160010185558215610be2579182015b82811115610be2578251825591602001919060010190610bc7565b50610bee929150610c1d565b5090565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b6107f091905b80821115610bee5760008155600101610c23565b803560158110610c4657600080fd5b92915050565b803560168110610c4657600080fd5b600082601f830112610c6b578081fd5b8135610c7e610c798261120d565b6111e6565b9150808252836020828501011115610c9557600080fd5b8060208401602084013760009082016020015292915050565b600082601f830112610cbe578081fd5b8151610ccc610c798261120d565b9150808252836020828501011115610ce357600080fd5b610cf4816020840160208601611231565b5092915050565b600060808284031215610d0c578081fd5b610d1660806111e6565b905081358152602082013567ffffffffffffffff80821115610d3757600080fd5b610d4385838601610c5b565b60208401526040840135915080821115610d5c57600080fd5b610d6885838601610c5b565b60408401526060840135915080821115610d8157600080fd5b50610d8e84828501610c5b565b60608301525092915050565b600060208284031215610dab578081fd5b81516001600160a01b0381168114610dc1578182fd5b9392505050565b600060208284031215610dd9578081fd5b815160058110610dc1578182fd5b600060208284031215610df8578081fd5b815167ffffffffffffffff80821115610e0f578283fd5b81840160808187031215610e21578384fd5b610e2b60806111e6565b925080518352602081015182811115610e42578485fd5b610e4e87828401610cae565b602085015250604081015182811115610e65578485fd5b610e7187828401610cae565b604085015250606081015182811115610e88578485fd5b610e9487828401610cae565b6060850152509195945050505050565b600060208284031215610eb5578081fd5b813567ffffffffffffffff80821115610ecc578283fd5b610160918401808603831315610ee0578384fd5b610ee9836111e6565b610ef38783610c4c565b8152610f028760208401610c37565b602082015260408201356040820152606082013560608201526080820135608082015260a0820135935082841115610f38578485fd5b610f4487858401610c5b565b60a082015260c0820135935082841115610f5c578485fd5b610f6887858401610c5b565b60c082015260e0820135935082841115610f80578485fd5b610f8c87858401610c5b565b60e082015261010093508382013583811115610fa6578586fd5b610fb288828501610cfb565b858301525061012093508382013583811115610fcc578586fd5b610fd888828501610cfb565b858301525061014093508382013583811115610ff2578586fd5b610ffe88828501610cfb565b948201949094529695505050505050565b6015811061101957fe5b9052565b6016811061101957fe5b6000815180845261103f816020860160208601611231565b601f01601f19169290920160200192915050565b6000815183526020820151608060208501526110726080850182611027565b60408401519150848103604086015261108b8183611027565b6060850151925085810360608701526110a48184611027565b9695505050505050565b90815260200190565b6000604082526110ca6040830185611053565b9050600583106110d657fe5b8260208301529392505050565b6000602082526110f760208301845161101d565b6020830151611109604084018261100f565b506040830151606083015260608301516080830152608083015160a083015260a08301516101608060c0850152611144610180850183611027565b60c08601519250601f19808683030160e08701526111628285611027565b60e0880151945061010092508187820301838801526111818186611027565b838901519550610120935082888203018489015261119f8187611053565b91505082880151945061014092508187820301838801526111c08186611053565b8389015195508288820301858901526111d98187611053565b9998505050505050505050565b60405181810167ffffffffffffffff8111828210171561120557600080fd5b604052919050565b600067ffffffffffffffff821115611223578081fd5b50601f01601f191660200190565b60005b8381101561124c578181015183820152602001611234565b8381111561125b576000848401525b5050505056fea2646970667358221220661583b4b216ff3830102a4dc0f6ac670628addaf3e1336204ef50a49716567564736f6c63430006080033a2646970667358221220bc671fca53ff3fb54860c674c0cb926d54f57e0fa923a181380dbcd40636826464736f6c63430006080033"

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

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
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
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
func (_Artefacts *ArtefactsSession) GetWork(_id [32]byte) (CreativeWorks, error) {
	return _Artefacts.Contract.GetWork(&_Artefacts.CallOpts, _id)
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
func (_Artefacts *ArtefactsCallerSession) GetWork(_id [32]byte) (CreativeWorks, error) {
	return _Artefacts.Contract.GetWork(&_Artefacts.CallOpts, _id)
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

// AddWork is a paid mutator transaction binding the contract method 0x5ed47d96.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsTransactor) AddWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "addWork", _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x5ed47d96.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWork(&_Artefacts.TransactOpts, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x5ed47d96.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsTransactorSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AddWork(&_Artefacts.TransactOpts, _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0xf8a3a2e0.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsTransactor) AmendWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.contract.Transact(opts, "amendWork", _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0xf8a3a2e0.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AmendWork(&_Artefacts.TransactOpts, _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0xf8a3a2e0.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_Artefacts *ArtefactsTransactorSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _Artefacts.Contract.AmendWork(&_Artefacts.TransactOpts, _work)
}

// EntitiesABI is the input ABI used to generate the binding from.
const EntitiesABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"name\":\"addEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amendEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntity\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityType\",\"outputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntitiesFuncSigs maps the 4-byte function signature to its string representation.
var EntitiesFuncSigs = map[string]string{
	"68a9f4e9": "addEntity((bytes32,string,string,string),uint8)",
	"a81eb545": "amendEntity((bytes32,string,string,string),uint8)",
	"53b66f36": "getEntity(bytes32)",
	"d91ee2e0": "getEntityContract(bytes32)",
	"641efe75": "getEntityType(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
}

// EntitiesBin is the compiled bytecode used for deploying new contracts.
var EntitiesBin = "0x608060405234801561001057600080fd5b506000600255611407806100256000396000f3fe60806040523480156200001157600080fd5b5060043610620000885760003560e01c806368a9f4e9116200006357806368a9f4e914620000fb578063a81eb5451462000114578063bc599341146200012b578063d91ee2e014620001425762000088565b806353b66f36146200008d578063641efe7514620000bc57806367e0badb14620000e2575b600080fd5b620000a46200009e36600462000636565b62000168565b604051620000b39190620008da565b60405180910390f35b620000d3620000cd36600462000636565b6200022e565b604051620000b39190620008c5565b620000ec620002df565b604051620000b39190620008bc565b620001126200010c36600462000739565b620002e5565b005b620001126200012536600462000739565b62000371565b620000ec6200013c36600462000636565b62000419565b620001596200015336600462000636565b62000453565b604051620000b39190620008a8565b6200017262000543565b6000828152602081905260409020600101546001600160a01b03166200019757600080fd5b600082815260208190526040808220600101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b158015620001e857600080fd5b505afa158015620001fd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526200022791908101906200066e565b9392505050565b6000818152602081905260408120600101546001600160a01b03166200025357600080fd5b60008281526020818152604091829020600101548251630aed701f60e11b815292516001600160a01b039091169283926315dae03e92600480840193829003018186803b158015620002a457600080fd5b505afa158015620002b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200022791906200064f565b60025490565b8151620002fb9060009063ffffffff6200049716565b1562000313576200030d828262000371565b6200036d565b6000828260405162000325906200056e565b62000332929190620008ef565b604051809103906000f0801580156200034f573d6000803e3d6000fd5b5083519091506200036a906000908363ffffffff620004af16565b50505b5050565b81516000908152602081905260409020600101546001600160a01b03166200039857600080fd5b81516000908152602081905260409081902060010154905163d1f439e560e01b81526001600160a01b0390911690819063d1f439e590620003e09086908690600401620008ef565b600060405180830381600087803b158015620003fb57600080fd5b505af115801562000410573d6000803e3d6000fd5b50505050505050565b60025460009082106200042b57600080fd5b60018054839081106200043a57fe5b9060005260206000209060020201600001549050919050565b6000818152602081905260408120600101546001600160a01b03166200047857600080fd5b506000908152602081905260409020600101546001600160a01b031690565b60008181526020839052604090205415155b92915050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b0385161790558015620004ef57600191505062000227565b50600180850180549182018082556000868152602088905260409020819055859190839081106200051c57fe5b60009182526020822060029182020192909255908601805460010190559150620002279050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b610a1b80620009b783390190565b8035620004a981620009a8565b600082601f8301126200059a578081fd5b8135620005b1620005ab8262000946565b6200091e565b9150808252836020828501011115620005c957600080fd5b8060208401602084013760009082016020015292915050565b600082601f830112620005f3578081fd5b815162000604620005ab8262000946565b91508082528360208285010111156200061c57600080fd5b6200062f8160208401602086016200096b565b5092915050565b60006020828403121562000648578081fd5b5035919050565b60006020828403121562000661578081fd5b81516200022781620009a8565b60006020828403121562000680578081fd5b815167ffffffffffffffff8082111562000698578283fd5b81840160808187031215620006ab578384fd5b620006b760806200091e565b925080518352602081015182811115620006cf578485fd5b620006dd87828401620005e2565b602085015250604081015182811115620006f5578485fd5b6200070387828401620005e2565b6040850152506060810151828111156200071b578485fd5b6200072987828401620005e2565b6060850152509195945050505050565b600080604083850312156200074c578081fd5b823567ffffffffffffffff8082111562000764578283fd5b8185016080818803121562000777578384fd5b6200078360806200091e565b9250803583526020810135828111156200079b578485fd5b620007a98882840162000589565b602085015250604081013582811115620007c1578485fd5b620007cf8882840162000589565b604085015250606081013582811115620007e7578485fd5b620007f58882840162000589565b6060850152505050809250506200081084602085016200057c565b90509250929050565b60008151808452620008338160208601602086016200096b565b601f01601f19169290920160200192915050565b60008151835260208201516080602085015262000868608085018262000819565b60408401519150848103604086015262000883818362000819565b6060850151925085810360608701526200089e818462000819565b9695505050505050565b6001600160a01b0391909116815260200190565b90815260200190565b60208101620008d4836200099a565b91905290565b60006020825262000227602083018462000847565b60006040825262000904604083018562000847565b905062000911836200099a565b8260208301529392505050565b60405181810167ffffffffffffffff811182821017156200093e57600080fd5b604052919050565b600067ffffffffffffffff8211156200095d578081fd5b50601f01601f191660200190565b60005b83811015620009885781810151838201526020016200096e565b838111156200036a5750506000910152565b60058110620009a557fe5b50565b60058110620009a557600080fdfe60806040523480156200001157600080fd5b5060405162000a1b38038062000a1b83398101604081905262000034916200029f565b60008160048111156200004357fe5b1180156200005d575060048160048111156200005b57fe5b105b6200006757600080fd5b6000805482919060ff191660018360048111156200008157fe5b02179055506200009a826001600160e01b03620000a216565b5050620003a5565b805160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615801590620000e057506000816020015151115b8015620000f257506000816040015151115b620000fc57600080fd5b805160015560208082015180516200011992600292019062000155565b5060408101518051620001359160039160209091019062000155565b5060608101518051620001519160049160209091019062000155565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019857805160ff1916838001178555620001c8565b82800160010185558215620001c8579182015b82811115620001c8578251825591602001919060010190620001ab565b50620001d6929150620001da565b5090565b620001f791905b80821115620001d65760008155600101620001e1565b90565b8051600581106200020a57600080fd5b92915050565b600082601f83011262000221578081fd5b81516001600160401b0381111562000237578182fd5b60206200024d601f8301601f191682016200037e565b925081835284818386010111156200026457600080fd5b60005b828110156200028457848101820151848201830152810162000267565b82811115620002965760008284860101525b50505092915050565b60008060408385031215620002b2578182fd5b82516001600160401b0380821115620002c9578384fd5b81850160808188031215620002dc578485fd5b620002e860806200037e565b92508051835260208101518281111562000300578586fd5b6200030e8882840162000210565b60208501525060408101518281111562000326578586fd5b620003348882840162000210565b6040850152506060810151828111156200034c578586fd5b6200035a8882840162000210565b606085015250505080925050620003758460208501620001fa565b90509250929050565b6040518181016001600160401b03811182821017156200039d57600080fd5b604052919050565b61066680620003b56000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806315dae03e146100465780636d4ce63c14610064578063d1f439e514610079575b600080fd5b61004e61008e565b60405161005b919061058b565b60405180910390f35b61006c610098565b60405161005b919061059f565b61008c610087366004610470565b61026d565b005b60005460ff165b90565b6100a061032f565b604080516080810182526001805482526002805484516020828516156101000260001901909216839004601f810183900483028201830190965285815293949293818601939092918301828280156101395780601f1061010e57610100808354040283529160200191610139565b820191906000526020600020905b81548152906001019060200180831161011c57829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156101cb5780601f106101a0576101008083540402835291602001916101cb565b820191906000526020600020905b8154815290600101906020018083116101ae57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561025f5780601f106102345761010080835404028352916020019161025f565b820191906000526020600020905b81548152906001019060200180831161024257829003601f168201915b505050505081525050905090565b80600481111561027957fe5b60005460ff16600481111561028a57fe5b1461029457600080fd5b61029d826102a1565b5050565b805160001a60f81b6001600160f81b031916158015906102c657506000816020015151115b80156102d757506000816040015151115b6102e057600080fd5b805160015560208082015180516102fb92600292019061035a565b50604081015180516103159160039160209091019061035a565b506060810151805161029d9160049160209091019061035a565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061039b57805160ff19168380011785556103c8565b828001600101855582156103c8579182015b828111156103c85782518255916020019190600101906103ad565b506103d49291506103d8565b5090565b61009591905b808211156103d457600081556001016103de565b80356005811061040157600080fd5b92915050565b600082601f830112610417578081fd5b813567ffffffffffffffff81111561042d578182fd5b610440601f8201601f1916602001610609565b915080825283602082850101111561045757600080fd5b8060208401602084013760009082016020015292915050565b60008060408385031215610482578182fd5b823567ffffffffffffffff80821115610499578384fd5b818501608081880312156104ab578485fd5b6104b56080610609565b9250803583526020810135828111156104cc578586fd5b6104d888828401610407565b6020850152506040810135828111156104ef578586fd5b6104fb88828401610407565b604085015250606081013582811115610512578586fd5b61051e88828401610407565b60608501525050508092505061053784602085016103f2565b90509250929050565b60008151808452815b8181101561056557602081850181015186830182015201610549565b818111156105765782602083870101525b50601f01601f19169290920160200192915050565b602081016005831061059957fe5b91905290565b600060208252825160208301526020830151608060408401526105c560a0840182610540565b60408501519150601f19808583030160608601526105e38284610540565b60608701519350818682030160808701526105fe8185610540565b979650505050505050565b60405181810167ffffffffffffffff8111828210171561062857600080fd5b60405291905056fea264697066735822122033821d4d8730780cfd8277423c84656701b05f61e5ab0e2f7d1f3f4b9dafd7f964736f6c63430006080033a26469706673582212206cce98e3f0c3d57db7ef75a21345a42389b9ec40b5d8db5ebefd046f314a139064736f6c63430006080033"

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

// GetEntityType is a free data retrieval call binding the contract method 0x641efe75.
//
// Solidity: function getEntityType(bytes32 _id) view returns(uint8)
func (_Entities *EntitiesCaller) GetEntityType(opts *bind.CallOpts, _id [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Entities.contract.Call(opts, out, "getEntityType", _id)
	return *ret0, err
}

// GetEntityType is a free data retrieval call binding the contract method 0x641efe75.
//
// Solidity: function getEntityType(bytes32 _id) view returns(uint8)
func (_Entities *EntitiesSession) GetEntityType(_id [32]byte) (uint8, error) {
	return _Entities.Contract.GetEntityType(&_Entities.CallOpts, _id)
}

// GetEntityType is a free data retrieval call binding the contract method 0x641efe75.
//
// Solidity: function getEntityType(bytes32 _id) view returns(uint8)
func (_Entities *EntitiesCallerSession) GetEntityType(_id [32]byte) (uint8, error) {
	return _Entities.Contract.GetEntityType(&_Entities.CallOpts, _id)
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
const EntityNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_entityType\",\"type\":\"uint8\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntityNodeFuncSigs maps the 4-byte function signature to its string representation.
var EntityNodeFuncSigs = map[string]string{
	"d1f439e5": "amend((bytes32,string,string,string),uint8)",
	"6d4ce63c": "get()",
	"15dae03e": "getType()",
}

// EntityNodeBin is the compiled bytecode used for deploying new contracts.
var EntityNodeBin = "0x60806040523480156200001157600080fd5b5060405162000a1b38038062000a1b83398101604081905262000034916200029f565b60008160048111156200004357fe5b1180156200005d575060048160048111156200005b57fe5b105b6200006757600080fd5b6000805482919060ff191660018360048111156200008157fe5b02179055506200009a826001600160e01b03620000a216565b5050620003a5565b805160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615801590620000e057506000816020015151115b8015620000f257506000816040015151115b620000fc57600080fd5b805160015560208082015180516200011992600292019062000155565b5060408101518051620001359160039160209091019062000155565b5060608101518051620001519160049160209091019062000155565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019857805160ff1916838001178555620001c8565b82800160010185558215620001c8579182015b82811115620001c8578251825591602001919060010190620001ab565b50620001d6929150620001da565b5090565b620001f791905b80821115620001d65760008155600101620001e1565b90565b8051600581106200020a57600080fd5b92915050565b600082601f83011262000221578081fd5b81516001600160401b0381111562000237578182fd5b60206200024d601f8301601f191682016200037e565b925081835284818386010111156200026457600080fd5b60005b828110156200028457848101820151848201830152810162000267565b82811115620002965760008284860101525b50505092915050565b60008060408385031215620002b2578182fd5b82516001600160401b0380821115620002c9578384fd5b81850160808188031215620002dc578485fd5b620002e860806200037e565b92508051835260208101518281111562000300578586fd5b6200030e8882840162000210565b60208501525060408101518281111562000326578586fd5b620003348882840162000210565b6040850152506060810151828111156200034c578586fd5b6200035a8882840162000210565b606085015250505080925050620003758460208501620001fa565b90509250929050565b6040518181016001600160401b03811182821017156200039d57600080fd5b604052919050565b61066680620003b56000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806315dae03e146100465780636d4ce63c14610064578063d1f439e514610079575b600080fd5b61004e61008e565b60405161005b919061058b565b60405180910390f35b61006c610098565b60405161005b919061059f565b61008c610087366004610470565b61026d565b005b60005460ff165b90565b6100a061032f565b604080516080810182526001805482526002805484516020828516156101000260001901909216839004601f810183900483028201830190965285815293949293818601939092918301828280156101395780601f1061010e57610100808354040283529160200191610139565b820191906000526020600020905b81548152906001019060200180831161011c57829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156101cb5780601f106101a0576101008083540402835291602001916101cb565b820191906000526020600020905b8154815290600101906020018083116101ae57829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561025f5780601f106102345761010080835404028352916020019161025f565b820191906000526020600020905b81548152906001019060200180831161024257829003601f168201915b505050505081525050905090565b80600481111561027957fe5b60005460ff16600481111561028a57fe5b1461029457600080fd5b61029d826102a1565b5050565b805160001a60f81b6001600160f81b031916158015906102c657506000816020015151115b80156102d757506000816040015151115b6102e057600080fd5b805160015560208082015180516102fb92600292019061035a565b50604081015180516103159160039160209091019061035a565b506060810151805161029d9160049160209091019061035a565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061039b57805160ff19168380011785556103c8565b828001600101855582156103c8579182015b828111156103c85782518255916020019190600101906103ad565b506103d49291506103d8565b5090565b61009591905b808211156103d457600081556001016103de565b80356005811061040157600080fd5b92915050565b600082601f830112610417578081fd5b813567ffffffffffffffff81111561042d578182fd5b610440601f8201601f1916602001610609565b915080825283602082850101111561045757600080fd5b8060208401602084013760009082016020015292915050565b60008060408385031215610482578182fd5b823567ffffffffffffffff80821115610499578384fd5b818501608081880312156104ab578485fd5b6104b56080610609565b9250803583526020810135828111156104cc578586fd5b6104d888828401610407565b6020850152506040810135828111156104ef578586fd5b6104fb88828401610407565b604085015250606081013582811115610512578586fd5b61051e88828401610407565b60608501525050508092505061053784602085016103f2565b90509250929050565b60008151808452815b8181101561056557602081850181015186830182015201610549565b818111156105765782602083870101525b50601f01601f19169290920160200192915050565b602081016005831061059957fe5b91905290565b600060208252825160208301526020830151608060408401526105c560a0840182610540565b60408501519150601f19808583030160608601526105e38284610540565b60608701519350818682030160808701526105fe8185610540565b979650505050505050565b60405181810167ffffffffffffffff8111828210171561062857600080fd5b60405291905056fea264697066735822122033821d4d8730780cfd8277423c84656701b05f61e5ab0e2f7d1f3f4b9dafd7f964736f6c63430006080033"

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

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_EntityNode *EntityNodeCaller) GetType(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _EntityNode.contract.Call(opts, out, "getType")
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_EntityNode *EntityNodeSession) GetType() (uint8, error) {
	return _EntityNode.Contract.GetType(&_EntityNode.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_EntityNode *EntityNodeCallerSession) GetType() (uint8, error) {
	return _EntityNode.Contract.GetType(&_EntityNode.CallOpts)
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
const IArtefactABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IArtefactFuncSigs maps the 4-byte function signature to its string representation.
var IArtefactFuncSigs = map[string]string{
	"adef5f6b": "amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))",
	"6d4ce63c": "get()",
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
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
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
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
func (_IArtefact *IArtefactSession) Get() (CreativeWorks, error) {
	return _IArtefact.Contract.Get(&_IArtefact.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
func (_IArtefact *IArtefactCallerSession) Get() (CreativeWorks, error) {
	return _IArtefact.Contract.Get(&_IArtefact.CallOpts)
}

// Amend is a paid mutator transaction binding the contract method 0xadef5f6b.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefact *IArtefactTransactor) Amend(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.contract.Transact(opts, "amend", _work)
}

// Amend is a paid mutator transaction binding the contract method 0xadef5f6b.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefact *IArtefactSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.Contract.Amend(&_IArtefact.TransactOpts, _work)
}

// Amend is a paid mutator transaction binding the contract method 0xadef5f6b.
//
// Solidity: function amend((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefact *IArtefactTransactorSession) Amend(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefact.Contract.Amend(&_IArtefact.TransactOpts, _work)
}

// IArtefactFactoryABI is the input ABI used to generate the binding from.
const IArtefactFactoryABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"addWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"_work\",\"type\":\"tuple\"}],\"name\":\"amendWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWork\",\"outputs\":[{\"components\":[{\"internalType\":\"enumWorksTypes\",\"name\":\"workType\",\"type\":\"uint8\"},{\"internalType\":\"enumLicenseTypes\",\"name\":\"license\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateCreated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dateModified\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"author\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"copyrightHolder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"publisher\",\"type\":\"tuple\"}],\"internalType\":\"structCreativeWorks\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getWorkContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IArtefactFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IArtefactFactoryFuncSigs = map[string]string{
	"5ed47d96": "addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))",
	"f8a3a2e0": "amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))",
	"30266537": "getWork(bytes32)",
	"4d2522a2": "getWorkContract(bytes32)",
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
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
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
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
func (_IArtefactFactory *IArtefactFactorySession) GetWork(_id [32]byte) (CreativeWorks, error) {
	return _IArtefactFactory.Contract.GetWork(&_IArtefactFactory.CallOpts, _id)
}

// GetWork is a free data retrieval call binding the contract method 0x30266537.
//
// Solidity: function getWork(bytes32 _id) view returns((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)))
func (_IArtefactFactory *IArtefactFactoryCallerSession) GetWork(_id [32]byte) (CreativeWorks, error) {
	return _IArtefactFactory.Contract.GetWork(&_IArtefactFactory.CallOpts, _id)
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

// AddWork is a paid mutator transaction binding the contract method 0x5ed47d96.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AddWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "addWork", _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x5ed47d96.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactorySession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWork(&_IArtefactFactory.TransactOpts, _work)
}

// AddWork is a paid mutator transaction binding the contract method 0x5ed47d96.
//
// Solidity: function addWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AddWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AddWork(&_IArtefactFactory.TransactOpts, _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0xf8a3a2e0.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactor) AmendWork(opts *bind.TransactOpts, _work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.contract.Transact(opts, "amendWork", _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0xf8a3a2e0.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactorySession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AmendWork(&_IArtefactFactory.TransactOpts, _work)
}

// AmendWork is a paid mutator transaction binding the contract method 0xf8a3a2e0.
//
// Solidity: function amendWork((uint8,uint8,bytes32,bytes32,bytes32,string,string,string,(bytes32,string,string,string),(bytes32,string,string,string),(bytes32,string,string,string)) _work) returns()
func (_IArtefactFactory *IArtefactFactoryTransactorSession) AmendWork(_work CreativeWorks) (*types.Transaction, error) {
	return _IArtefactFactory.Contract.AmendWork(&_IArtefactFactory.TransactOpts, _work)
}

// IEntitiesFactoryABI is the input ABI used to generate the binding from.
const IEntitiesFactoryABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amendEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntity\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getEntityType\",\"outputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEntitiesFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IEntitiesFactoryFuncSigs = map[string]string{
	"68a9f4e9": "addEntity((bytes32,string,string,string),uint8)",
	"a81eb545": "amendEntity((bytes32,string,string,string),uint8)",
	"53b66f36": "getEntity(bytes32)",
	"d91ee2e0": "getEntityContract(bytes32)",
	"641efe75": "getEntityType(bytes32)",
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

// GetEntityType is a free data retrieval call binding the contract method 0x641efe75.
//
// Solidity: function getEntityType(bytes32 _id) view returns(uint8)
func (_IEntitiesFactory *IEntitiesFactoryCaller) GetEntityType(opts *bind.CallOpts, _id [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IEntitiesFactory.contract.Call(opts, out, "getEntityType", _id)
	return *ret0, err
}

// GetEntityType is a free data retrieval call binding the contract method 0x641efe75.
//
// Solidity: function getEntityType(bytes32 _id) view returns(uint8)
func (_IEntitiesFactory *IEntitiesFactorySession) GetEntityType(_id [32]byte) (uint8, error) {
	return _IEntitiesFactory.Contract.GetEntityType(&_IEntitiesFactory.CallOpts, _id)
}

// GetEntityType is a free data retrieval call binding the contract method 0x641efe75.
//
// Solidity: function getEntityType(bytes32 _id) view returns(uint8)
func (_IEntitiesFactory *IEntitiesFactoryCallerSession) GetEntityType(_id [32]byte) (uint8, error) {
	return _IEntitiesFactory.Contract.GetEntityType(&_IEntitiesFactory.CallOpts, _id)
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
const IEntityABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"_entity\",\"type\":\"tuple\"},{\"internalType\":\"enumEntityTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structCreativeEntities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"enumEntityTypes\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEntityFuncSigs maps the 4-byte function signature to its string representation.
var IEntityFuncSigs = map[string]string{
	"d1f439e5": "amend((bytes32,string,string,string),uint8)",
	"6d4ce63c": "get()",
	"15dae03e": "getType()",
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

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_IEntity *IEntityCaller) GetType(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IEntity.contract.Call(opts, out, "getType")
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_IEntity *IEntitySession) GetType() (uint8, error) {
	return _IEntity.Contract.GetType(&_IEntity.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(uint8)
func (_IEntity *IEntityCallerSession) GetType() (uint8, error) {
	return _IEntity.Contract.GetType(&_IEntity.CallOpts)
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

// IterableDataABI is the input ABI used to generate the binding from.
const IterableDataABI = "[]"

// IterableDataBin is the compiled bytecode used for deploying new contracts.
var IterableDataBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c27c006fe91b3e0c8af6114d4785a6afaf10f5b0833ef2e140b30d2e3ebade9464736f6c63430006080033"

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
