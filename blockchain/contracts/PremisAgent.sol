pragma solidity ^0.4.11;

// Fake News Premis Object
// Steve Huckle

import "Mortal.sol";
import "Agent.sol";
import "Strings.sol";

contract PremisAgent is Agent, Mortal {

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

  /* function PremisAgent() {
  } */

  function setAgent(string _agentId, string _name, string _type) onlyOwner {
    premisAgents[_agentId].name = _name;
    premisAgents[_agentId].agentType = _type;
    agents.push(_agentId);
    AgentSet(_agentId, _name, _type);
  }

  function setObject(string _agentId, string _objectHash) onlyOwner {
    premisAgents[_agentId].objects.push(_objectHash);
    AgentSetObject(_agentId, _objectHash);
  }

  function setEvent(string _agentId, string _eventId) onlyOwner {
    premisAgents[_agentId].events.push(_eventId);
    AgentSetEvent(_agentId, _eventId);
  }

  function setRights(string _agentId, string _rightsId) onlyOwner {
    premisAgents[_agentId].rights.push(_rightsId);
    AgentSetRights(_agentId, _rightsId);
  }

  // These 'Exists?' functions return max(uint256) if not, the index otherwise
  function getAgentExists(string _agentId) constant returns (uint256) {
    return Strings.getIndex(_agentId, agents);
  }

  function getAgentObjectExists(string _agentId, string _objectHash) constant returns (uint256) {
    return Strings.getIndex(_objectHash, premisAgents[_agentId].objects);
  }

  function getAgentEventExists(string _agentId, string _eventId) constant returns (uint256) {
    return Strings.getIndex(_eventId, premisAgents[_agentId].events);
  }

  function getAgentRightsExists(string _agentId, string _rightsId) constant returns (uint256) {
    return Strings.getIndex(_rightsId, premisAgents[_agentId].rights);
  }

  function getNumAgents() constant returns (uint256) {
    return agents.length;
  }

  function getNumObjects(string _agentId) constant returns (uint256) {
    return premisAgents[_agentId].objects.length;
  }

  function getNumEvents(string _agentId) constant returns (uint256) {
    return premisAgents[_agentId].events.length;
  }

  function getNumRights(string _agentId) constant returns (uint256) {
    return premisAgents[_agentId].rights.length;
  }

  function getAgent(uint256 _index) constant returns (string) {
    return agents[_index];
  }

  function getName(string _agentId) constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].name);
  }

  function getType(string _agentId) constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].agentType);
  }

  function getObject(string _agentId, uint256 _objectIndex) constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].objects[_objectIndex]);
  }

  function getEvent(string _agentId, uint256 _eventIndex) constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].events[_eventIndex]);
  }

  function getRights(string _agentId, uint256 _rightsIndex) constant returns (string, string) {
    return (_agentId, premisAgents[_agentId].rights[_rightsIndex]);
  }
}
