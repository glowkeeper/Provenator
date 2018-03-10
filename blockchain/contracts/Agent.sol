pragma solidity ^0.4.19;

contract Agent {

  function setAgent(string _agentId, string _name, string _type) public;
  function setObject(string _agentId, string _objectHash) public;
  function setEvent(string _agentId, string _eventId) public;
  function setRights(string _agentId, string _rightsId) public;

  function getAgentExists(string _agentId) public constant returns (uint256);
  function getAgentObjectExists(string _agentId, string _objectHash) public constant returns (uint256);
  function getAgentEventExists(string _agentId, string _eventId) public constant returns (uint256);
  function getAgentRightsExists(string _agentId, string _rightsId) public constant returns (uint256);

  function getNumAgents() public constant returns (uint256);
  function getNumObjects(string _agentId) public constant returns (uint256);
  function getNumEvents(string _agentId) public constant returns (uint256);
  function getNumRights(string _agentId) public constant returns (uint256);

  function getAgent(uint256 _index) public constant returns (string);

  function getName(string _agentId) public constant returns (string, string);
  function getType(string _agentId) public constant returns (string, string);
  function getObject(string _agentId, uint256 _objectIndex) public constant returns (string, string);
  function getEvent(string _agentId, uint256 _eventIndex) public constant returns (string, string);
  function getRights(string _agentId, uint256 _rightsIndex) public constant returns (string, string);

}
