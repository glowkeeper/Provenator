import Web3Handler from './web3Handler'

class ContractHandler {

  static premisObjectAbi = [
    {
      "constant": true,
      "inputs": [
        {
          "name": "_index",
          "type": "uint256"
        }
      ],
      "name": "getPropertyTypeName",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getObjectRightsExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "getCategory",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_index",
          "type": "uint256"
        }
      ],
      "name": "getObject",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getNumObjects",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_propType",
          "type": "string"
        }
      ],
      "name": "getPropertyTypeExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "kill",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_type",
          "type": "string"
        },
        {
          "name": "_value",
          "type": "string"
        }
      ],
      "name": "setProperties",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "setEvent",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "getObjectExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "setRights",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "getNumProperties",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "getNumAgents",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "getObjectEventExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "getFormat",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_propType",
          "type": "string"
        }
      ],
      "name": "setPropertyType",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "getObjectAgentExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getNumPropTypes",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_rightsIndex",
          "type": "uint256"
        }
      ],
      "name": "getRights",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_eventIndex",
          "type": "uint256"
        }
      ],
      "name": "getEvent",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "getNumRights",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_propertyIndex",
          "type": "uint256"
        }
      ],
      "name": "getProperties",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "getNumEvents",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_agentIndex",
          "type": "uint256"
        }
      ],
      "name": "getAgent",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "setAgent",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_category",
          "type": "string"
        },
        {
          "name": "_format",
          "type": "string"
        }
      ],
      "name": "setObject",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_objectHash",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_category",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_format",
          "type": "string"
        }
      ],
      "name": "ObjectSet",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_objectHash",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_propType",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_propValue",
          "type": "string"
        }
      ],
      "name": "ObjectSetProperties",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_objectHash",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "ObjectSetEvent",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_objectHash",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "ObjectSetRights",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_objectHash",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "ObjectSetAgent",
      "type": "event"
    }
  ]

  static premisAgentAbi = [
    {
      "constant": false,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "setObject",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "getAgentObjectExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "getNumObjects",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_name",
          "type": "string"
        },
        {
          "name": "_type",
          "type": "string"
        }
      ],
      "name": "setAgent",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_index",
          "type": "uint256"
        }
      ],
      "name": "getAgent",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "getAgentExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "kill",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "setEvent",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getAgentRightsExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "setRights",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "getName",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "getAgentEventExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "getType",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_rightsIndex",
          "type": "uint256"
        }
      ],
      "name": "getRights",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_eventIndex",
          "type": "uint256"
        }
      ],
      "name": "getEvent",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "getNumRights",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "getNumEvents",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getNumAgents",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_agentId",
          "type": "string"
        },
        {
          "name": "_objectIndex",
          "type": "uint256"
        }
      ],
      "name": "getObject",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_address",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_name",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_type",
          "type": "string"
        }
      ],
      "name": "AgentSet",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_address",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "AgentSetObject",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_address",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "AgentSetEvent",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_address",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "AgentSetRights",
      "type": "event"
    }
  ]

  static premisEventAbi = [
    {
      "constant": true,
      "inputs": [
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "getObject",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_eventTypeIndex",
          "type": "uint256"
        }
      ],
      "name": "getEventType",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "kill",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getNumEvents",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_eventType",
          "type": "string"
        }
      ],
      "name": "setEventType",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_eventIndex",
          "type": "uint256"
        }
      ],
      "name": "getEvent",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "getEventExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "getAgent",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "getType",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_eventId",
          "type": "string"
        }
      ],
      "name": "getTime",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_eventType",
          "type": "string"
        }
      ],
      "name": "getEventTypeExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_eventId",
          "type": "string"
        },
        {
          "name": "_type",
          "type": "string"
        },
        {
          "name": "_time",
          "type": "string"
        },
        {
          "name": "_objectHash",
          "type": "string"
        },
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "setEvent",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getNumEventTypes",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_eventId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_type",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_time",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_objectHash",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "EventSet",
      "type": "event"
    }
  ]

  static premisRightsAbi = [
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getObject",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getGrantedRestriction",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getRightsBasis",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        },
        {
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "setObject",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_index",
          "type": "uint256"
        }
      ],
      "name": "getRights",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getCopyrightDeterminationDate",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getCopyrightStatus",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getGrantedAct",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "kill",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        },
        {
          "name": "_act",
          "type": "string"
        },
        {
          "name": "_restriction",
          "type": "string"
        }
      ],
      "name": "setRightsGranted",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        },
        {
          "name": "_basis",
          "type": "string"
        }
      ],
      "name": "setBasis",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        },
        {
          "name": "_status",
          "type": "string"
        },
        {
          "name": "_jurisdictionCountryCode",
          "type": "string"
        },
        {
          "name": "_determinationDate",
          "type": "string"
        },
        {
          "name": "_note",
          "type": "string"
        }
      ],
      "name": "setCopyrightInfo",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getAgent",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getRightsExists",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getCopyrightJurisdiction",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getNumRights",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        },
        {
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "setAgent",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "payable": false,
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_rightsId",
          "type": "string"
        }
      ],
      "name": "getCopyrightNote",
      "outputs": [
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "type": "function"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_rightsId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_objectHash",
          "type": "string"
        }
      ],
      "name": "RightsSetObject",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_rightsId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_basis",
          "type": "string"
        }
      ],
      "name": "RightsSetBasis",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_rightsId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_status",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_jurisdictionCountryCode",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_determinationDate",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_note",
          "type": "string"
        }
      ],
      "name": "RightsSetCopyrightInfo",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_rightsId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_act",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_restriction",
          "type": "string"
        }
      ],
      "name": "RightsSetRightsGranted",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_rightsId",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_agentId",
          "type": "string"
        }
      ],
      "name": "RightsSetAgent",
      "type": "event"
    }
  ]

  /*  Deploying PremisAgent...
  ... 0xce2b44653f144853945dfb86ea220f6ba9ccb4cafb03d202c843b86625ea6016
  PremisAgent: 0x934a94e376079206f5a2d213a6152f0164d29535
  Deploying PremisEvent...
  ... 0xe6276f874817b15a83c0cc96cb4915289c21620140e1faed3ac7312bcb67aeb8
  PremisEvent: 0x762895147173daaff705066f5573162c95f617d0
  Deploying PremisObject...
  ... 0x9bb89c921fcd4c79abdd123ec114ecc9399026e304d0e8a78ac1260822e0e279
  PremisObject: 0xda14e434e957e9c7e053b5d8fb80e6ed2dc174ee
  Deploying PremisRights...
  ... 0xf49023f401f09be368799eab5aa7ee351ca389ea7ca90fe2df928c0b1cc086d2
  PremisRights: 0x973b9658d34bbd3dc06be257020d93900cf32073

  */

  static premisObjectContractAddress = '0xda14e434e957e9c7e053b5d8fb80e6ed2dc174ee'
  static premisEventContractAddress = '0x762895147173daaff705066f5573162c95f617d0'
  static premisAgentContractAddress = '0x934a94e376079206f5a2d213a6152f0164d29535'
  static premisRightsContractAddress = '0x973b9658d34bbd3dc06be257020d93900cf32073'

  constructor (_web3Handler) {
    //console.log(_web3)

    this.web3Handler = _web3Handler
    const thisWeb3 = this.web3Handler.getWeb3()

    this.premisObjectContract = thisWeb3.eth.contract(ContractHandler.premisObjectAbi)
    this.premisObject = this.premisObjectContract.at(ContractHandler.premisObjectContractAddress)

    this.premisEventContract = thisWeb3.eth.contract(ContractHandler.premisEventAbi)
    this.premisEvent = this.premisEventContract.at(ContractHandler.premisEventContractAddress)

    this.premisAgentContract = thisWeb3.eth.contract(ContractHandler.premisAgentAbi)
    this.premisAgent = this.premisAgentContract.at(ContractHandler.premisAgentContractAddress)

    this.premisRightsContract = thisWeb3.eth.contract(ContractHandler.premisRightsAbi)
    this.premisRights = this.premisRightsContract.at(ContractHandler.premisRightsContractAddress)
  }

  getPremisObject () {
    return this.premisObject
  }

  getPremisAgent () {
    return this.premisAgent
  }

  getPremisEvent () {
    return this.premisEvent
  }

  getPremisRights () {
    return this.premisRights
  }
}

export default ContractHandler
