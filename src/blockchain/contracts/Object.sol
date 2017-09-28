pragma solidity ^0.4.11;

contract Object {

  function setPropertyType(string _propType);
  function setObject(string _objectHash, string _category, string _format);
  function setProperties(string _objectHash, string _type, string _value);
  function setEvent(string _objectHash, string _eventId);
  function setAgent(string _objectHash, string _agentId);
  function setRights(string _objectHash, string _rightsId);

  function getPropertyTypeExists(string _propType) constant returns (uint256);
  function getObjectExists(string _objectHash) constant returns (uint256);
  function getObjectEventExists(string _objectHash, string _eventId) constant returns (uint256);
  function getObjectAgentExists(string _objectHash, string _agentId) constant returns (uint256);
  function getObjectRightsExists(string _objectHash, string _rightsId) constant returns (uint256);

  function getNumPropTypes() constant returns (uint256);
  function getNumObjects() constant returns (uint256);
  function getNumProperties(string _objectHash) constant returns (uint256);
  function getNumEvents(string _objectHash) constant returns (uint256);
  function getNumAgents(string _objectHash) constant returns (uint256);
  function getNumRights(string _objectHash) constant returns (uint256);

  function getPropertyTypeName(uint256 _index) constant returns (string);

  function getObject(uint256 _index) constant returns (string);
  function getCategory(string _objectHash) constant returns (string, string);
  function getFormat(string _objectHash) constant returns (string, string);
  function getProperties(string _objectHash, uint256 _propertyIndex) constant returns (string, string, string);
  function getEvent(string _objectHash, uint256 _eventIndex) constant returns (string, string);
  function getAgent(string _objectHash, uint256 _agentIndex) constant returns (string, string);
  function getRights(string _objectHash, uint256 _rightsIndex) constant returns (string, string);

}
