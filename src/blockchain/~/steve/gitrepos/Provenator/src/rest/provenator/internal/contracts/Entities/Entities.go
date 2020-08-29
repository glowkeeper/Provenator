// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Entities

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
var EntitiesBin = "0x608060405234801561001057600080fd5b506000600255612103806100256000396000f3fe60806040523480156200001157600080fd5b5060043610620001005760003560e01c80636ec21da91162000099578063bc599341116200006f578063bc5993411462000227578063cef8b362146200023e578063d090cbc01462000255578063d91ee2e0146200026c5762000100565b80636ec21da914620001d15780637a6337fa14620001f7578063a81eb54514620002105762000100565b806367e0badb11620000db57806367e0badb146200017157806368a9f4e9146200017b5780636c3cf7f714620001945780636e62b52714620001ab5762000100565b80633e30fc9b146200010557806353b66f3614620001345780635caaf3e8146200015a575b600080fd5b6200011c6200011636600462000da9565b62000292565b6040516200012b9190620010a3565b60405180910390f35b6200014b6200014536600462000d21565b62000359565b6040516200012b9190620010bc565b6200011c6200016b36600462000d21565b6200041f565b6200011c620004d0565b620001926200018c36600462000e85565b620004d6565b005b62000192620001a536600462000d53565b62000556565b620001c2620001bc36600462000d53565b620005f8565b6040516200012b919062001098565b620001e8620001e236600462000d21565b620006b5565b6040516200012b919062001016565b620002016200076d565b6040516200012b91906200105e565b620001926200022136600462000e85565b62000811565b6200011c6200023836600462000d21565b62000880565b620001c26200024f36600462000d75565b620008ba565b620002016200026636600462000d21565b62000923565b620002836200027d36600462000d21565b620009db565b6040516200012b919062001002565b6000828152602081905260408120600101546001600160a01b0316620002b757600080fd5b6000838152602081905260409081902060010154905163bc59934160e01b81526001600160a01b0390911690819063bc59934190620002fb908690600401620010a3565b60206040518083038186803b1580156200031457600080fd5b505afa15801562000329573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200034f919062000d3a565b9150505b92915050565b6200036362000ac8565b6000828152602081905260409020600101546001600160a01b03166200038857600080fd5b600082815260208190526040808220600101548151631b53398f60e21b815291516001600160a01b03909116928392636d4ce63c9260048083019392829003018186803b158015620003d957600080fd5b505afa158015620003ee573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000418919081019062000dbc565b9392505050565b6000818152602081905260408120600101546001600160a01b03166200044457600080fd5b600082815260208181526040918290206001015482516367e0badb60e01b815292516001600160a01b039091169283926367e0badb92600480840193829003018186803b1580156200049557600080fd5b505afa158015620004aa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000418919062000d3a565b60025490565b8151620004e69060009062000a1f565b15620004fe57620004f8828262000811565b62000552565b60008282604051620005109062000af3565b6200051d929190620010d1565b604051809103906000f0801580156200053a573d6000803e3d6000fd5b5083519091506200054f906000908362000a34565b50505b5050565b6000828152602081905260409020600101546001600160a01b03166200057b57600080fd5b600082815260208190526040908190206001015490516351366bd360e11b81526001600160a01b0390911690819063a26cd7a690620005bf908590600401620010a3565b600060405180830381600087803b158015620005da57600080fd5b505af1158015620005ef573d6000803e3d6000fd5b50505050505050565b6000828152602081905260408120600101546001600160a01b03166200061d57600080fd5b600083815260208190526040908190206001015490516379ae21e760e11b81526001600160a01b0390911690819063f35c43ce9062000661908690600401620010a3565b60206040518083038186803b1580156200067a57600080fd5b505afa1580156200068f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200034f919062000cff565b6000818152602081905260409020600101546060906001600160a01b0316620006dd57600080fd5b6000828152602081905260408082206001015481516305a2bceb60e51b815291516001600160a01b0390911692839263b4579d609260048083019392829003018186803b1580156200072e57600080fd5b505afa15801562000743573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000418919081019062000bcf565b60608060006002015467ffffffffffffffff811180156200078d57600080fd5b50604051908082528060200260200182016040528015620007b8578160200160208202803683370190505b50905060005b6002548110156200080b576001805482908110620007d857fe5b906000526020600020906002020160000154828281518110620007f757fe5b6020908102919091010152600101620007be565b50905090565b81516000908152602081905260409020600101546001600160a01b03166200083857600080fd5b81516000908152602081905260409081902060010154905163d1f439e560e01b81526001600160a01b0390911690819063d1f439e590620005bf9086908690600401620010d1565b60025460009082106200089257600080fd5b6001805483908110620008a157fe5b9060005260206000209060020201600001549050919050565b6000828152602081905260408120600101546001600160a01b0316620008df57600080fd5b6000838152602081905260409081902060010154905163222f74e760e01b81526001600160a01b0390911690819063222f74e79062000661908690600401620010ac565b6000818152602081905260409020600101546060906001600160a01b03166200094b57600080fd5b600082815260208190526040808220600101548151633d319bfd60e11b815291516001600160a01b03909116928392637a6337fa9260048083019392829003018186803b1580156200099c57600080fd5b505afa158015620009b1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405262000418919081019062000c72565b6000818152602081905260408120600101546001600160a01b031662000a0057600080fd5b506000908152602081905260409020600101546001600160a01b031690565b60009081526020919091526040902054151590565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b038516179055801562000a7457600191505062000418565b506001808501805491820180825560008681526020889052604090208190558591908390811062000aa157fe5b60009182526020822060029182020192909255908601805460010190559150620004189050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b610f39806200119583390190565b805180151581146200035357600080fd5b8035600581106200035357600080fd5b600082601f83011262000b33578081fd5b813562000b4a62000b448262001140565b620010f7565b915080825283602082850101111562000b6257600080fd5b8060208401602084013760009082016020015292915050565b600082601f83011262000b8c578081fd5b815162000b9d62000b448262001140565b915080825283602082850101111562000bb557600080fd5b62000bc881602084016020860162001165565b5092915050565b6000602080838503121562000be2578182fd5b825167ffffffffffffffff81111562000bf9578283fd5b8301601f8101851362000c0a578283fd5b805162000c1b62000b44826200111f565b818152838101908385018584028501860189101562000c38578687fd5b8694505b8385101562000c665762000c51898262000b01565b83526001949094019391850191850162000c3c565b50979650505050505050565b6000602080838503121562000c85578182fd5b825167ffffffffffffffff81111562000c9c578283fd5b8301601f8101851362000cad578283fd5b805162000cbe62000b44826200111f565b818152838101908385018584028501860189101562000cdb578687fd5b8694505b8385101562000c6657805183526001949094019391850191850162000cdf565b60006020828403121562000d11578081fd5b8151801515811462000418578182fd5b60006020828403121562000d33578081fd5b5035919050565b60006020828403121562000d4c578081fd5b5051919050565b6000806040838503121562000d66578081fd5b50508035926020909101359150565b6000806040838503121562000d88578182fd5b8235915060208301356005811062000d9e578182fd5b809150509250929050565b6000806040838503121562000d66578182fd5b60006020828403121562000dce578081fd5b815167ffffffffffffffff8082111562000de6578283fd5b908301906080828603121562000dfa578283fd5b62000e066080620010f7565b8251815260208301518281111562000e1c578485fd5b62000e2a8782860162000b7b565b60208301525060408301518281111562000e42578485fd5b62000e508782860162000b7b565b60408301525060608301518281111562000e68578485fd5b62000e768782860162000b7b565b60608301525095945050505050565b6000806040838503121562000e98578182fd5b823567ffffffffffffffff8082111562000eb0578384fd5b908401906080828703121562000ec4578384fd5b62000ed06080620010f7565b8235815260208301358281111562000ee6578586fd5b62000ef48882860162000b22565b60208301525060408301358281111562000f0c578586fd5b62000f1a8882860162000b22565b60408301525060608301358281111562000f32578586fd5b62000f408882860162000b22565b60608301525080945050505062000f5b846020850162000b12565b90509250929050565b6005811062000f6f57fe5b9052565b6000815180845262000f8d81602086016020860162001165565b601f01601f19169290920160200192915050565b60008151835260208201516080602085015262000fc2608085018262000f73565b90506040830151848203604086015262000fdd828262000f73565b9150506060830151848203606086015262000ff9828262000f73565b95945050505050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b818110156200105257835115158352928401929184019160010162001032565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101562001052578351835292840192918401916001016200107a565b901515815260200190565b90815260200190565b6020810162000353828462000f64565b60006020825262000418602083018462000fa1565b600060408252620010e6604083018562000fa1565b905062000418602083018462000f64565b60405181810167ffffffffffffffff811182821017156200111757600080fd5b604052919050565b600067ffffffffffffffff82111562001136578081fd5b5060209081020190565b600067ffffffffffffffff82111562001157578081fd5b50601f01601f191660200190565b60005b838110156200118257818101518382015260200162001168565b838111156200054f575050600091015256fe60806040523480156200001157600080fd5b5060405162000f3938038062000f398339810160408190526200003491620003a8565b815160001a60f81b7fff0000000000000000000000000000000000000000000000000000000000000016158015906200007257506000826020015151115b80156200008457506000826040015151115b80156200009d575060008160048111156200009b57fe5b115b8015620000b657506004816004811115620000b457fe5b105b620000c057600080fd5b60408051600480825260a08201909252906020820160808036833750508151620000f2926004925060200190620001a4565b50600160048260048111156200010457fe5b60ff16815481106200011257fe5b6000918252602080832081830401805460ff601f9094166101000a9384021916941515929092029390931790558351815583820151805185936200015c9260019291019062000250565b50604082015180516200017a91600284019160209091019062000250565b50606082015180516200019891600384019160209091019062000250565b509050505050620004ad565b82805482825590600052602060002090601f016020900481019282156200023e5791602002820160005b838211156200020d57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302620001ce565b80156200023c5782816101000a81549060ff02191690556001016020816000010492830192600103026200020d565b505b506200024c929150620002d1565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200029357805160ff1916838001178555620002c3565b82800160010185558215620002c3579182015b82811115620002c3578251825591602001919060010190620002a6565b506200024c929150620002ec565b5b808211156200024c57805460ff19168155600101620002d2565b5b808211156200024c5760008155600101620002ed565b8051600581106200031357600080fd5b92915050565b600082601f8301126200032a578081fd5b81516001600160401b0381111562000340578182fd5b602062000356601f8301601f1916820162000486565b925081835284818386010111156200036d57600080fd5b60005b828110156200038d57848101820151848201830152810162000370565b828111156200039f5760008284860101525b50505092915050565b60008060408385031215620003bb578182fd5b82516001600160401b0380821115620003d2578384fd5b9084019060808287031215620003e6578384fd5b620003f2608062000486565b8251815260208301518281111562000408578586fd5b620004168882860162000319565b6020830152506040830151828111156200042e578586fd5b6200043c8882860162000319565b60408301525060608301518281111562000454578586fd5b620004628882860162000319565b6060830152508094505050506200047d846020850162000303565b90509250929050565b6040518181016001600160401b0381118282101715620004a557600080fd5b604052919050565b610a7c80620004bd6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b4579d6011610066578063b4579d6014610120578063bc59934114610135578063d1f439e514610148578063e188130b1461015b578063f35c43ce1461016e5761009e565b8063222f74e7146100a357806367e0badb146100cc5780636d4ce63c146100e15780637a6337fa146100f6578063a26cd7a61461010b575b600080fd5b6100b66100b13660046107e2565b610181565b6040516100c391906109a0565b60405180910390f35b6100d46101f3565b6040516100c391906109ab565b6100e96101f9565b6040516100c391906109b4565b6100fe6103d1565b6040516100c39190610968565b61011e6101193660046107ca565b610429565b005b610128610499565b6040516100c39190610922565b6100d46101433660046107ca565b610510565b61011e610156366004610808565b610540565b61011e6101693660046107e2565b6105a9565b6100b661017c3660046107ca565b61062e565b60008082600481111561019057fe5b1180156101a8575060048260048111156101a657fe5b105b6101b157600080fd5b60048260048111156101bf57fe5b60ff16815481106101cc57fe5b90600052602060002090602091828204019190069054906101000a900460ff169050919050565b60055490565b61020161068e565b604080516080810182526000805482526001805484516020600283851615610100026000190190931692909204601f8101839004830282018301909652858152939492938186019390929183018282801561029d5780601f106102725761010080835404028352916020019161029d565b820191906000526020600020905b81548152906001019060200180831161028057829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561032f5780601f106103045761010080835404028352916020019161032f565b820191906000526020600020905b81548152906001019060200180831161031257829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156103c35780601f10610398576101008083540402835291602001916103c3565b820191906000526020600020905b8154815290600101906020018083116103a657829003601f168201915b505050505081525050905090565b6060600580548060200260200160405190810160405280929190818152602001828054801561041f57602002820191906000526020600020905b81548152602001906001019080831161040b575b5050505050905090565b8060001a60f81b6001600160f81b03191661044357600080fd5b600054811480159061045b57506104598161062e565b155b1561049657600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018190555b50565b6060600480548060200260200160405190810160405280929190818152602001828054801561041f57602002820191906000526020600020906000905b825461010083900a900460ff1615158152602060019283018181049485019490930390920291018084116104d65790505050505050905090565b600554600090821061052157600080fd5b6005828154811061052e57fe5b90600052602060002001549050919050565b610549816105a9565b81516000908155602080840151805185939261056a926001929101906106b9565b50604082015180516105869160028401916020909101906106b9565b50606082015180516105a29160038401916020909101906106b9565b5050505050565b60008160048111156105b757fe5b1180156105cf575060048160048111156105cd57fe5b105b6105d857600080fd5b6105e181610181565b61049657600160048260048111156105f557fe5b60ff168154811061060257fe5b90600052602060002090602091828204019190066101000a81548160ff02191690831515021790555050565b600081811a60f81b6001600160f81b03191661064957600080fd5b6000805b60055481101561068757836005828154811061066557fe5b9060005260206000200154141561067f5760019150610687565b60010161064d565b5092915050565b6040518060800160405280600080191681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106fa57805160ff1916838001178555610727565b82800160010185558215610727579182015b8281111561072757825182559160200191906001019061070c565b50610733929150610737565b5090565b5b808211156107335760008155600101610738565b80356005811061075b57600080fd5b92915050565b600082601f830112610771578081fd5b813567ffffffffffffffff811115610787578182fd5b61079a601f8201601f1916602001610a1f565b91508082528360208285010111156107b157600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156107db578081fd5b5035919050565b6000602082840312156107f3578081fd5b813560058110610801578182fd5b9392505050565b6000806040838503121561081a578081fd5b823567ffffffffffffffff80821115610831578283fd5b9084019060808287031215610844578283fd5b61084e6080610a1f565b82358152602083013582811115610863578485fd5b61086f88828601610761565b602083015250604083013582811115610886578485fd5b61089288828601610761565b6040830152506060830135828111156108a9578485fd5b6108b588828601610761565b6060830152508094505050506108ce846020850161074c565b90509250929050565b60008151808452815b818110156108fc576020818501810151868301820152016108e0565b8181111561090d5782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b8181101561095c57835115158352928401929184019160010161093e565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561095c57835183529284019291840191600101610984565b901515815260200190565b90815260200190565b600060208252825160208301526020830151608060408401526109da60a08401826108d7565b90506040840151601f19808584030160608601526109f883836108d7565b9250606086015191508085840301608086015250610a1682826108d7565b95945050505050565b60405181810167ffffffffffffffff81118282101715610a3e57600080fd5b60405291905056fea2646970667358221220f92eb1360c908b75f4dc4ac1c5c7a481d6ebae9cdcee440cca9ed18514353ea164736f6c63430007000033a26469706673582212200c4a8078fbe0cafd9b89127bf40a3ea32bb26a59b6aab4361725875018cd1d0464736f6c63430007000033"

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
