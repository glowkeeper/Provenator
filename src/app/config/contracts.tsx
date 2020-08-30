
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
        "function isEntityType(bytes32 _id, uint8 _type) view returns (bool)",
        "function containsEntityRelation(bytes32 _parentId, bytes32 _childId) view returns (bool)",
        "function getRelationsNum(bytes32 _id) view returns (uint256)",
        "function getRelationsReference(bytes32 _id, uint256 _index) view returns (bytes32)",
        "function getRelations(bytes32 _id) view returns (bytes32[])"
    ]

    static artefactsABI = [
        "function addWork(tuple(uint8 workType, uint8 license, bytes32 id, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description, tuple(bytes32 id, string name, string email, string url) author) _work)",
        "function amendWork(tuple(uint8 workType, uint8 license, bytes32 id, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description, tuple(bytes32 id, string name, string email, string url) author) _work)",
        "function addWorkAuthor(bytes32 _workId, tuple(bytes32 id, string name, string email, string url) _author)",
        "function addWorkCopyrightHolder(bytes32 _workId, tuple(bytes32 id, string name, string email, string url) _copyrightHolder)",
        "function addWorkPublisher(bytes32 _workId, tuple(bytes32 id, string name, string email, string url) _publisher)",
        "function removeWorkAuthor(bytes32 _workId, bytes32 _authorId)",
        "function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId)",
        "function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId)",
        "function getWork(bytes32 _id) view returns (tuple(uint8 workType, uint8 license, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description))",
        "function getWorkAuthors(bytes32 _id) view returns (bytes32[])",
        "function getWorkCopyrightHolders(bytes32 _id) view returns (bytes32[])",
        "function getWorkPublishers(bytes32 _id) view returns (bytes32[])",
        "function getWorkContract(bytes32 _id) view returns (address)",
        "function isWorkAuthor(bytes32 _workId, bytes32 _authorId) view returns (bool)",
        "function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) view returns (bool)",
        "function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) view returns (bool)",
        "function getNum() view returns (uint256)",
        "function getReference(uint256 _index) view returns (bytes32)",
        "function getReferences() view returns (bytes32[])"
    ]
}

export { Config, Contracts }
