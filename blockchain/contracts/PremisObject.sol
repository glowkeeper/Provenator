pragma solidity ^0.4.19;

// Fake News Premis Object
// Steve Huckle

import "Object.sol";
import "Strings.sol";

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

  function setPropertyType(string _propType) public {
    propTypes.push(_propType);
  }

  function setObject(string _objectHash, string _category, string _format) public {
    premisObjects[_objectHash].category = _category;
    premisObjects[_objectHash].format = _format;
    objects.push(_objectHash);
    emit ObjectSet(_objectHash, _category, _format);
  }

  function setProperties(string _objectHash, string _type, string _value) public {
    premisObjects[_objectHash].properties.push(ObjectProperties(_type,_value));
    emit ObjectSetProperties(_objectHash, _type, _value);
  }

  function setEvent(string _objectHash, string _eventId) public {
    premisObjects[_objectHash].events.push(_eventId);
    emit ObjectSetEvent(_objectHash, _eventId);
  }

  function setAgent(string _objectHash, string _agentId) public {
    premisObjects[_objectHash].agents.push(_agentId);
  }

  function setRights(string _objectHash, string _rightsId) public {
    premisObjects[_objectHash].rights.push(_rightsId);
  }

  // These 'Exists?' functions return max(uint256) if not, the index otherwise
  function getPropertyTypeExists(string _propType) public constant returns (uint256) {
    return Strings.getIndex(_propType, propTypes);
  }

  function getObjectExists(string _objectHash) public constant returns (uint256) {
    return Strings.getIndex(_objectHash, objects);
  }

  function getObjectEventExists(string _objectHash, string _eventId) public constant returns (uint256) {
    return Strings.getIndex(_eventId, premisObjects[_objectHash].events);
  }

  function getObjectAgentExists(string _objectHash, string _agentId) public constant returns (uint256) {
    return Strings.getIndex(_agentId, premisObjects[_objectHash].agents);
  }

  function getObjectRightsExists(string _objectHash, string _rightsId) public constant returns (uint256) {
    return Strings.getIndex(_rightsId, premisObjects[_objectHash].rights);
  }

  function getNumPropTypes() public constant returns (uint256) {
    return propTypes.length;
  }

  function getNumObjects() public constant returns (uint256) {
    return objects.length;
  }

  function getNumProperties(string _objectHash) public constant returns (uint256) {
    return premisObjects[_objectHash].properties.length;
  }

  function getNumEvents(string _objectHash) public constant returns (uint256) {
    return premisObjects[_objectHash].events.length;
  }

  function getNumAgents(string _objectHash) public constant returns (uint256) {
    return premisObjects[_objectHash].agents.length;
  }

  function getNumRights(string _objectHash) public constant returns (uint256) {
    return premisObjects[_objectHash].rights.length;
  }

  function getPropertyTypeName(uint256 _index) public constant returns (string) {
    return propTypes[_index];
  }

  function getObject(uint256 _index) public constant returns (string) {
    return objects[_index];
  }

  function getCategory(string _objectHash) public constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].category);
  }

  function getFormat(string _objectHash) public constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].format);
  }

  function getProperties(string _objectHash, uint256 _propertyIndex) public constant returns (string, string, string) {
    return (_objectHash, premisObjects[_objectHash].properties[_propertyIndex].propType,
            premisObjects[_objectHash].properties[_propertyIndex].propValue);
  }

  function getEvent(string _objectHash, uint256 _eventIndex) public constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].events[_eventIndex]);
  }

  function getAgent(string _objectHash, uint256 _agentIndex) public constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].agents[_agentIndex]);
  }

  function getRights(string _objectHash, uint256 _rightsIndex) public constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].rights[_rightsIndex]);
  }
}
