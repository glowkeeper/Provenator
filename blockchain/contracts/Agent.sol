pragma solidity ^0.4.11;

contract Agent {

  function setAgent(string _agentId, string _name, string _type);
  function setObject(string _agentId, string _objectHash);
  function setEvent(string _agentId, string _eventId);
  function setRights(string _agentId, string _rightsId);

  function getAgentExists(string _agentId) constant returns (uint256);
  function getAgentObjectExists(string _agentId, string _objectHash) constant returns (uint256);
  function getAgentEventExists(string _agentId, string _eventId) constant returns (uint256);
  function getAgentRightsExists(string _agentId, string _rightsId) constant returns (uint256);

  function getNumAgents() constant returns (uint256);
  function getNumObjects(string _agentId) constant returns (uint256);
  function getNumEvents(string _agentId) constant returns (uint256);
  function getNumRights(string _agentId) constant returns (uint256);

  function getAgent(uint256 _index) constant returns (string);

  function getName(string _agentId) constant returns (string, string);
  function getType(string _agentId) constant returns (string, string);
  function getObject(string _agentId, uint256 _objectIndex) constant returns (string, string);
  function getEvent(string _agentId, uint256 _eventIndex) constant returns (string, string);
  function getRights(string _agentId, uint256 _rightsIndex) constant returns (string, string);

}
