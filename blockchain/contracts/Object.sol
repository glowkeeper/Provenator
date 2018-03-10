pragma solidity ^0.4.19;

contract Object {

  function setPropertyType(string _propType) public;
  function setObject(string _objectHash, string _category, string _format) public;
  function setProperties(string _objectHash, string _type, string _value) public;
  function setEvent(string _objectHash, string _eventId) public;
  function setAgent(string _objectHash, string _agentId) public;
  function setRights(string _objectHash, string _rightsId) public;

  function getPropertyTypeExists(string _propType) public constant returns (uint256);
  function getObjectExists(string _objectHash) public constant returns (uint256);
  function getObjectEventExists(string _objectHash, string _eventId) public constant returns (uint256);
  function getObjectAgentExists(string _objectHash, string _agentId) public constant returns (uint256);
  function getObjectRightsExists(string _objectHash, string _rightsId) public constant returns (uint256);

  function getNumPropTypes() public constant returns (uint256);
  function getNumObjects() public constant returns (uint256);
  function getNumProperties(string _objectHash) public constant returns (uint256);
  function getNumEvents(string _objectHash) public constant returns (uint256);
  function getNumAgents(string _objectHash) public constant returns (uint256);
  function getNumRights(string _objectHash) public constant returns (uint256);

  function getPropertyTypeName(uint256 _index) public constant returns (string);

  function getObject(uint256 _index) public constant returns (string);
  function getCategory(string _objectHash) public constant returns (string, string);
  function getFormat(string _objectHash) public constant returns (string, string);
  function getProperties(string _objectHash, uint256 _propertyIndex) public constant returns (string, string, string);
  function getEvent(string _objectHash, uint256 _eventIndex) public constant returns (string, string);
  function getAgent(string _objectHash, uint256 _agentIndex) public constant returns (string, string);
  function getRights(string _objectHash, uint256 _rightsIndex) public constant returns (string, string);

}
