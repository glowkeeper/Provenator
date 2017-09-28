pragma solidity ^0.4.11;

// Fake News Premis Object
// Steve Huckle

import "Mortal.sol";
import "Object.sol";
import "Strings.sol";

contract PremisObject is Object, Mortal {

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

  /* function PremisObject() {
  } */

  function setPropertyType(string _propType) onlyOwner {
    propTypes.push(_propType);
  }

  function setObject(string _objectHash, string _category, string _format) onlyOwner {
    premisObjects[_objectHash].category = _category;
    premisObjects[_objectHash].format = _format;
    objects.push(_objectHash);
    ObjectSet(_objectHash, _category, _format);
  }

  function setProperties(string _objectHash, string _type, string _value) onlyOwner {
    premisObjects[_objectHash].properties.push(ObjectProperties(_type,_value));
    ObjectSetProperties(_objectHash, _type, _value);
  }

  function setEvent(string _objectHash, string _eventId) onlyOwner {
    premisObjects[_objectHash].events.push(_eventId);
    ObjectSetEvent(_objectHash, _eventId);
  }

  function setAgent(string _objectHash, string _agentId) onlyOwner {
    premisObjects[_objectHash].agents.push(_agentId);
  }

  function setRights(string _objectHash, string _rightsId) onlyOwner {
    premisObjects[_objectHash].rights.push(_rightsId);
  }

  // These 'Exists?' functions return max(uint256) if not, the index otherwise
  function getPropertyTypeExists(string _propType) constant returns (uint256) {
    return Strings.getIndex(_propType, propTypes);
  }

  function getObjectExists(string _objectHash) constant returns (uint256) {
    return Strings.getIndex(_objectHash, objects);
  }

  function getObjectEventExists(string _objectHash, string _eventId) constant returns (uint256) {
    return Strings.getIndex(_eventId, premisObjects[_objectHash].events);
  }

  function getObjectAgentExists(string _objectHash, string _agentId) constant returns (uint256) {
    return Strings.getIndex(_agentId, premisObjects[_objectHash].agents);
  }

  function getObjectRightsExists(string _objectHash, string _rightsId) constant returns (uint256) {
    return Strings.getIndex(_rightsId, premisObjects[_objectHash].rights);
  }

  function getNumPropTypes() constant returns (uint256) {
    return propTypes.length;
  }

  function getNumObjects() constant returns (uint256) {
    return objects.length;
  }

  function getNumProperties(string _objectHash) constant returns (uint256) {
    return premisObjects[_objectHash].properties.length;
  }

  function getNumEvents(string _objectHash) constant returns (uint256) {
    return premisObjects[_objectHash].events.length;
  }

  function getNumAgents(string _objectHash) constant returns (uint256) {
    return premisObjects[_objectHash].agents.length;
  }

  function getNumRights(string _objectHash) constant returns (uint256) {
    return premisObjects[_objectHash].rights.length;
  }

  function getPropertyTypeName(uint256 _index) constant returns (string) {
    return propTypes[_index];
  }

  function getObject(uint256 _index) constant returns (string) {
    return objects[_index];
  }

  function getCategory(string _objectHash) constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].category);
  }

  function getFormat(string _objectHash) constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].format);
  }

  function getProperties(string _objectHash, uint256 _propertyIndex) constant returns (string, string, string) {
    return (_objectHash, premisObjects[_objectHash].properties[_propertyIndex].propType,
            premisObjects[_objectHash].properties[_propertyIndex].propValue);
  }

  function getEvent(string _objectHash, uint256 _eventIndex) constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].events[_eventIndex]);
  }

  function getAgent(string _objectHash, uint256 _agentIndex) constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].agents[_agentIndex]);
  }

  function getRights(string _objectHash, uint256 _rightsIndex) constant returns (string, string) {
    return (_objectHash, premisObjects[_objectHash].rights[_rightsIndex]);
  }
}
