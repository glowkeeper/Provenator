pragma solidity ^0.4.24;

// Fake News Premis Object
// Steve Huckle

import "Agent.sol";
import "Strings.sol";

contract PremisAgent is Agent {

  struct AgentData {
    string name;
    string agentType;

    string[] objects;
    string[] events;
    string[] rights;
  }

  string[] agents;
  mapping (string => AgentData) private premisAgents;

  event AgentSet(string _address, string _name, string _type);
  event AgentSetObject(string _address, string _objectHash);
  event AgentSetEvent(string _address, string _eventId);
  event AgentSetRights(string _address, string _rightsId);

  constructor() public {
  }

  function setAgent(string _agentId, string _name, string _type) public {
    premisAgents[_agentId].name = _name;
    premisAgents[_agentId].agentType = _type;
    agents.push(_agentId);
    emit AgentSet(_agentId, _name, _type);
  }

  function setObject(string _agentId, string _objectHash) public {
    premisAgents[_agentId].objects.push(_objectHash);
    emit AgentSetObject(_agentId, _objectHash);
  }

  function setEvent(string _agentId, string _eventId) public {
    premisAgents[_agentId].events.push(_eventId);
    emit AgentSetEvent(_agentId, _eventId);
  }

  function setRights(string _agentId, string _rightsId) public {
    premisAgents[_agentId].rights.push(_rightsId);
    emit AgentSetRights(_agentId, _rightsId);
  }

  // These 'Exists?' functions return max(uint256) if not, the index otherwise
  function getAgentExists(string _agentId) public constant returns (uint256) {
    return Strings.getIndex(_agentId, agents);
  }

  function getAgentObjectExists(string _agentId, string _objectHash) public constant returns (uint256) {
    return Strings.getIndex(_objectHash, premisAgents[_agentId].objects);
  }

  function getAgentEventExists(string _agentId, string _eventId) public constant returns (uint256) {
    return Strings.getIndex(_eventId, premisAgents[_agentId].events);
  }

  function getAgentRightsExists(string _agentId, string _rightsId) public constant returns (uint256) {
    return Strings.getIndex(_rightsId, premisAgents[_agentId].rights);
  }

  function getNumAgents() public constant returns (uint256) {
    return agents.length;
  }

  function getNumObjects(string _agentId) public constant returns (uint256) {
    return premisAgents[_agentId].objects.length;
  }

  function getNumEvents(string _agentId) public constant returns (uint256) {
    return premisAgents[_agentId].events.length;
  }

  function getNumRights(string _agentId) public constant returns (uint256) {
    return premisAgents[_agentId].rights.length;
  }

  function getAgent(uint256 _index) public constant returns (string) {
    return agents[_index];
  }

  function getName(string _agentId) public constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].name);
  }

  function getType(string _agentId) public constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].agentType);
  }

  function getObject(string _agentId, uint256 _objectIndex) public constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].objects[_objectIndex]);
  }

  function getEvent(string _agentId, uint256 _eventIndex) public constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].events[_eventIndex]);
  }

  function getRights(string _agentId, uint256 _rightsIndex) public constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].rights[_rightsIndex]);
  }
}
