pragma solidity ^0.5.0;

// Fake News Premis Object
// Steve Huckle

import "./Object.sol";
import "./Strings.sol";

contract PremisObject is Object {

  struct ObjectProperties {
    string propType;
    string propValue;
  }

  struct ObjectData {
    string category;
    string format;

    ObjectProperties[] properties;
    string[] events;
    string[] agents;
    string[] rights;
  }

  string[] propTypes;
  string[] objects;
  mapping(string => ObjectData) private premisObjects;

  event ObjectSet(string _objectHash, string _category, string _format);
  event ObjectSetProperties(string _objectHash, string _propType, string _propValue);
  event ObjectSetEvent(string _objectHash, string _eventId);
  event ObjectSetRights(string _objectHash, string _rightsId);
  event ObjectSetAgent(string _objectHash, string _agentId);

  constructor() public {
  }

  function setPropertyType(string memory _propType) public {
    propTypes.push(_propType);
  }

  function setObject(string memory _objectHash, string memory _category, string memory _format) public {
    premisObjects[_objectHash].category = _category;
    premisObjects[_objectHash].format = _format;
    objects.push(_objectHash);
    emit ObjectSet(_objectHash, _category, _format);
  }

  function setProperties(string memory _objectHash, string memory _type, string memory _value) public {
    premisObjects[_objectHash].properties.push(ObjectProperties(_type,_value));
    emit ObjectSetProperties(_objectHash, _type, _value);
  }

  function setEvent(string memory _objectHash, string memory _eventId) public {
    premisObjects[_objectHash].events.push(_eventId);
    emit ObjectSetEvent(_objectHash, _eventId);
  }

  function setAgent(string memory _objectHash, string memory _agentId) public {
    premisObjects[_objectHash].agents.push(_agentId);
  }

  function setRights(string memory _objectHash, string memory _rightsId) public {
    premisObjects[_objectHash].rights.push(_rightsId);
  }

  // These 'Exists?' functions return max(uint256) if not, the index otherwise
  function getPropertyTypeExists(string memory _propType) public view returns (bool) {
    uint256 index = Strings.getIndex(_propType, propTypes);
    return index != propTypes.length;
  }

  function getObjectExists(string memory _objectHash) public view returns (bool) {
    uint256 index = Strings.getIndex(_objectHash, objects);
    return index != objects.length;
  }

  function getObjectPropertiesExist(string memory _objectHash, string memory _propType, string memory _propValue) public view returns (bool) {
    bool exists = false;
    if(getObjectExists(_objectHash)) {
      for (uint256 x = 0; x < premisObjects[_objectHash].properties.length; x++) {
        if (Strings.equal(_propType, premisObjects[_objectHash].properties[x].propType)) {
          if(Strings.equal(_propValue, premisObjects[_objectHash].properties[x].propValue)) {
            exists = true;
          }
          break;
        }
      }
    }
    return exists;
  }

  function getObjectEventExists(string memory _objectHash, string memory _eventId) public view returns (bool) {
    uint256 index = Strings.getIndex(_eventId, premisObjects[_objectHash].events);
    return index != premisObjects[_objectHash].events.length;
  }

  function getObjectAgentExists(string memory _objectHash, string memory _agentId) public view returns (bool) {
    uint256 index = Strings.getIndex(_agentId, premisObjects[_objectHash].agents);
    return index != premisObjects[_objectHash].agents.length;
  }

  function getObjectRightsExists(string memory _objectHash, string memory _rightsId) public view returns (bool) {
    uint256 index = Strings.getIndex(_rightsId, premisObjects[_objectHash].rights);
    return index != premisObjects[_objectHash].rights.length;
  }

  function getNumPropTypes() public view returns (uint256) {
    return propTypes.length;
  }

  function getNumObjects() public view returns (uint256) {
    return objects.length;
  }

  function getNumProperties(string memory _objectHash) public view returns (uint256) {
    return premisObjects[_objectHash].properties.length;
  }

  function getNumEvents(string memory _objectHash) public view returns (uint256) {
    return premisObjects[_objectHash].events.length;
  }

  function getNumAgents(string memory _objectHash) public view returns (uint256) {
    return premisObjects[_objectHash].agents.length;
  }

  function getNumRights(string memory _objectHash) public view returns (uint256) {
    return premisObjects[_objectHash].rights.length;
  }

  function getPropertyTypeName(uint256 _index) public view returns (string memory) {
    return propTypes[_index];
  }

  function getObject(uint256 _index) public view returns (string memory) {
    return objects[_index];
  }

  function getCategory(string memory _objectHash) public view returns (string memory, string memory) {
    return (_objectHash, premisObjects[_objectHash].category);
  }

  function getFormat(string memory _objectHash) public view returns (string memory, string memory) {
    return (_objectHash, premisObjects[_objectHash].format);
  }

  function getProperties(string memory _objectHash, uint256 _propertyIndex) public view returns (string memory, string memory, string memory) {
    return (_objectHash, premisObjects[_objectHash].properties[_propertyIndex].propType,
            premisObjects[_objectHash].properties[_propertyIndex].propValue);
  }

  function getEvent(string memory _objectHash, uint256 _eventIndex) public view returns (string memory, string memory) {
    return (_objectHash, premisObjects[_objectHash].events[_eventIndex]);
  }

  function getAgent(string memory _objectHash, uint256 _agentIndex) public view returns (string memory, string memory) {
    return (_objectHash, premisObjects[_objectHash].agents[_agentIndex]);
  }

  function getRights(string memory _objectHash, uint256 _rightsIndex) public view returns (string memory, string memory) {
    return (_objectHash, premisObjects[_objectHash].rights[_rightsIndex]);
  }
}
