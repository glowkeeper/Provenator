
class Config {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'http://127.0.0.1'
  static readonly port = '8545'
  static readonly network = 'rinkeby'

  static readonly checkInterval = 3000
}

class Contracts {

    // rinkeby
    static entitiesAddress = "0xa435af28590d0fA707B17aa6b7E381FBa6DA4159"
    static artefactsAddress = "0xaF1130Cc84e7e24C58959b706fd47d3E896b1ef0"

    static entitiesABI = [
        "function addEntity(tuple(bytes32 id, string name, string email, string url) _entity, uint8 _entityType)",
        "function amendEntity(tuple(bytes32 id, string name, string email, string url) _entity, uint8 _entityType)",
        "function getEntity(bytes32 _id) view returns (tuple(bytes32 id, string name, string email, string url))",
        "function getEntityType(bytes32 _id) view returns (uint8)",
        "function getEntityContract(bytes32 _id) view returns (address)",
        "function getNum() view returns (uint256)",
        "function getReference(uint256 _index) view returns (bytes32)"
    ]

    static artefactsABI = [
        "function addWork(tuple(uint8 workType, uint8 license, bytes32 id, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description, tuple(bytes32 id, string name, string email, string url) author, tuple(bytes32 id, string name, string email, string url) copyrightHolder, tuple(bytes32 id, string name, string email, string url) publisher) _work)",
        "function amendWork(tuple(uint8 workType, uint8 license, bytes32 id, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description, tuple(bytes32 id, string name, string email, string url) author, tuple(bytes32 id, string name, string email, string url) copyrightHolder, tuple(bytes32 id, string name, string email, string url) publisher) _work)",
        "function getWork(bytes32 _id) view returns (tuple(uint8 workType, uint8 license, bytes32 id, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description, tuple(bytes32 id, string name, string email, string url) author, tuple(bytes32 id, string name, string email, string url) copyrightHolder, tuple(bytes32 id, string name, string email, string url) publisher) _work)",
        "function getWorkContract(bytes32 _id) view returns (address)",
        "function getNum() view returns (uint256)",
        "function getReference(uint256 _index) view returns (bytes32)"
    ]

}

export { Config, Contracts }
