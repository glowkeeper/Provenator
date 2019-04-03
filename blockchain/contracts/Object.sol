pragma solidity ^0.5.0;

contract Object {

  function setPropertyType(string memory _propType) public;
  function setObject(string memory _objectHash, string memory _category, string memory _format) public;
  function setProperties(string memory _objectHash, string memory _type, string memory _value) public;
  function setEvent(string memory _objectHash, string memory _eventId) public;
  function setAgent(string memory _objectHash, string memory _agentId) public;
  function setRights(string memory _objectHash, string memory _rightsId) public;

  function getPropertyTypeExists(string memory _propType) public view returns (bool);
  function getObjectExists(string memory _objectHash) public view returns (bool);
  function getObjectPropertiesExist(string memory _objectHash, string memory _propType, string memory _propValue) public view returns (bool);
  function getObjectEventExists(string memory _objectHash, string memory _eventId) public view returns (bool);
  function getObjectAgentExists(string memory _objectHash, string memory _agentId) public view returns (bool);
  function getObjectRightsExists(string memory _objectHash, string memory _rightsId) public view returns (bool);

  function getNumPropTypes() public view returns (uint256);
  function getNumObjects() public view returns (uint256);
  function getNumProperties(string memory _objectHash) public view returns (uint256);
  function getNumEvents(string memory _objectHash) public view returns (uint256);
  function getNumAgents(string memory _objectHash) public view returns (uint256);
  function getNumRights(string memory _objectHash) public view returns (uint256);

  function getPropertyTypeName(uint256 _index) public view returns (string memory);

  function getObject(uint256 _index) public view returns (string memory);
  function getCategory(string memory _objectHash) public view returns (string memory, string memory);
  function getFormat(string memory _objectHash) public view returns (string memory, string memory);
  function getProperties(string memory _objectHash, uint256 _propertyIndex) public view returns (string memory, string memory, string memory);
  function getEvent(string memory _objectHash, uint256 _eventIndex) public view returns (string memory, string memory);
  function getAgent(string memory _objectHash, uint256 _agentIndex) public view returns (string memory, string memory);
  function getRights(string memory _objectHash, uint256 _rightsIndex) public view returns (string memory, string memory);

}
