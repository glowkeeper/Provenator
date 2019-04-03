pragma solidity ^0.5.0;

contract Agent {

  function setAgent(string memory _agentId, string memory _name, string memory _type) public;
  function setObject(string memory _agentId, string memory _objectHash) public;
  function setEvent(string memory _agentId, string memory _eventId) public;
  function setRights(string memory _agentId, string memory _rightsId) public;

  function getAgentExists(string memory _agentId) public view returns (bool);
  function getAgentObjectExists(string memory _agentId, string memory _objectHash) public view returns (bool);
  function getAgentEventExists(string memory _agentId, string memory _eventId) public view returns (bool);
  function getAgentRightsExists(string memory _agentId, string memory _rightsId) public view returns (bool);

  function getNumAgents() public view returns (uint256);
  function getNumObjects(string memory _agentId) public view returns (uint256);
  function getNumEvents(string memory _agentId) public view returns (uint256);
  function getNumRights(string memory _agentId) public view returns (uint256);

  function getAgent(uint256 _index) public view returns (string memory);

  function getName(string memory _agentId) public view returns (string memory, string memory);
  function getType(string memory _agentId) public view returns (string memory, string memory);
  function getObject(string memory _agentId, uint256 _objectIndex) public view returns (string memory, string memory);
  function getEvent(string memory _agentId, uint256 _eventIndex) public view returns (string memory, string memory);
  function getRights(string memory _agentId, uint256 _rightsIndex) public view returns (string memory, string memory);

}
