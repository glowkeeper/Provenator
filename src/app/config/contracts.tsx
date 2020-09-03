
class Config {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'http://127.0.0.1'
  static readonly port = '8545'
  static readonly network = 'rinkeby'

  static readonly checkInterval = 3000
}

class Contracts {

    // rinkeby
    static entitiesAddress = "0x2cAb3f679A12cb03daD25A97Bd516c7F9b5f2ea0"
    static artefactsAddress = "0x1a4704B6F23B811C326b08571003a50c657b72Aa"

    static entitiesABI = [
        "function addEntity(tuple(bytes32 id, string name, string email, string url) _entity, uint8 _entityType)",
        "function amendEntity(tuple(bytes32 id, string name, string email, string url) _entity, uint8 _entityType)",
        "function getEntity(bytes32 _id) view returns (tuple(bytes32 id, string name, string email, string url))",
        "function getEntityTypes(bytes32 _id) view returns (bool[])",
        "function getEntityContract(bytes32 _id) view returns (address)",
        "function getNum() view returns (uint256)",
        "function getReference(uint256 _index) view returns (bytes32)",
        "function getReferences() view returns (bytes32[])",
        "function isEntityType(bytes32 _id, uint8 _type) view returns (bool)",
    ]

    static artefactsABI = [
        "function addWork(tuple(uint8 workType, uint8 license, bytes32 id, bytes32 authorId, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description) _work)",
        "function amendWork(tuple(uint8 workType, uint8 license, bytes32 id, bytes32 authorId, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description) _work)",
        "function addWorkAuthor(bytes32 _workId, bytes32 _authorId)",
        "function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId)",
        "function addWorkPublisher(bytes32 _workId, bytes32 _copyrightHolderId)",
        "function removeWorkAuthor(bytes32 _workId, bytes32 _authorId)",
        "function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _publisherId)",
        "function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId)",
        "function getWork(bytes32 _id) view returns (tuple(uint8 workType, uint8 license, bytes32 id, bytes32 authorId, bytes32 dateCreated, bytes32 dateModified, string url, string name, string description))",
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
