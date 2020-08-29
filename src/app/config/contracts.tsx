
class Config {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'http://127.0.0.1'
  static readonly port = '8545'
  static readonly network = 'rinkeby'

  static readonly checkInterval = 3000
}

class Contracts {

    // rinkeby
    static entitiesAddress = "0x9BF1F90733Ec4eE006271CFD33647ED652E8C46D"
    static artefactsAddress = "0x9B8CCCD12f5359924236975CD1309E9cEFa598dE"

    static entitiesABI = [
        "function addEntity(tuple(bytes32 id, string name, string email, string url) _entity, uint8 _entityType)",
        "function amendEntity(tuple(bytes32 id, string name, string email, string url) _entity, uint8 _entityType)",
        "function addEntityRelation(bytes32 _parentId, bytes32 _childId)",
        "function getEntity(bytes32 _id) view returns (tuple(bytes32 id, string name, string email, string url))",
        "function getEntityTypes(bytes32 _id) view returns (bool[])",
        "function getEntityContract(bytes32 _id) view returns (address)",
        "function getNum() view returns (uint256)",
        "function getReference(uint256 _index) view returns (bytes32)",
        "function getReferences() view returns (bytes32[])",
        "function isType(bytes32 _id, uint8 _type) view returns (bool)",
        "function containsRelation(bytes32 _parentId, bytes32 _childId) view returns (bool)",
        "function getRelationsNum(bytes32 _id) view returns (uint256)",
        "function getRelationsReference(bytes32 _id, uint256 _index) view returns (bytes32)",
        "function getRelations(bytes32 _id) view returns (bytes32[])"
    ]

    static artefactsABI = [
        "function addWork(tuple(uint8 workType, uint8 license, bytes32 id, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description, tuple(bytes32 id, string name, string email, string url) author, tuple(bytes32 id, string name, string email, string url) copyrightHolder, tuple(bytes32 id, string name, string email, string url) publisher) _work)",
        "function amendWork(tuple(uint8 workType, uint8 license, bytes32 id, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description, tuple(bytes32 id, string name, string email, string url) author, tuple(bytes32 id, string name, string email, string url) copyrightHolder, tuple(bytes32 id, string name, string email, string url) publisher) _work)",
        "function getWork(bytes32 _id) view returns (tuple(uint8 workType, uint8 license, bytes32 id, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description, tuple(bytes32 id, string name, string email, string url) author, tuple(bytes32 id, string name, string email, string url) copyrightHolder, tuple(bytes32 id, string name, string email, string url) publisher) _work)",
        "function getWorkContract(bytes32 _id) view returns (address)",
        "function getNum() view returns (uint256)",
        "function getReference(uint256 _index) view returns (bytes32)",
        "function getReferences() view returns (bytes32[])"
    ]
}

export { Config, Contracts }
