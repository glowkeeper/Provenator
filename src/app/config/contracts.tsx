
class Config {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'http://127.0.0.1'
  static readonly port = '8545'
  static readonly network = 'rinkeby'

  static readonly checkInterval = 3000
}

class Contracts {

    // rinkeby
    static entitiesAddress = "0x9A08c375ecb50c8A52fB434D3cc3927dff056DF1"
    static artefactsAddress = "0xAbd55D232D281Ce74a0BF2B3C7532156603878a1"

    static entitiesABI = [
        "function addEntity(tuple(bytes32 id, string name, string email, string url) _entity, uint8 _entityType)",
        "function amendEntity(tuple(bytes32 id, string name, string email, string url) _entity)",
        "function getEntity(bytes32 _id) view returns (tuple(bytes32 id, string name, string email, string url))",
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
