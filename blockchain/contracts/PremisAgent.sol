pragma solidity ^0.5.0;

// Fake News Premis Object
// Steve Huckle

import "./Agent.sol";
import "./Strings.sol";

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

  function setAgent(string memory _agentId, string memory _name, string memory _type) public {
    premisAgents[_agentId].name = _name;
    premisAgents[_agentId].agentType = _type;
    agents.push(_agentId);
    emit AgentSet(_agentId, _name, _type);
  }

  function setObject(string memory _agentId, string memory _objectHash) public {
    premisAgents[_agentId].objects.push(_objectHash);
    emit AgentSetObject(_agentId, _objectHash);
  }

  function setEvent(string memory _agentId, string memory _eventId) public {
    premisAgents[_agentId].events.push(_eventId);
    emit AgentSetEvent(_agentId, _eventId);
  }

  function setRights(string memory _agentId, string memory _rightsId) public {
    premisAgents[_agentId].rights.push(_rightsId);
    emit AgentSetRights(_agentId, _rightsId);
  }

  // These 'Exists?' functions return max(uint256) if not, the index otherwise
  function getAgentExists(string memory _agentId) public view returns (bool) {
    uint256 index = Strings.getIndex(_agentId, agents);
    return index != agents.length;
  }

  function getAgentObjectExists(string memory _agentId, string memory _objectHash) public view returns (bool) {
    uint256 index = Strings.getIndex(_objectHash, premisAgents[_agentId].objects);
    return index !=  premisAgents[_agentId].objects.length;
  }

  function getAgentEventExists(string memory _agentId, string memory _eventId) public view returns (bool) {
    uint256 index = Strings.getIndex(_eventId, premisAgents[_agentId].events);
    return index != premisAgents[_agentId].events.length;
  }

  function getAgentRightsExists(string memory _agentId, string memory _rightsId) public view returns (bool) {
    uint256 index = Strings.getIndex(_rightsId, premisAgents[_agentId].rights);
    return index != premisAgents[_agentId].rights.length;
  }

  function getNumAgents() public view returns (uint256) {
    return agents.length;
  }

  function getNumObjects(string memory _agentId) public view returns (uint256) {
    return premisAgents[_agentId].objects.length;
  }

  function getNumEvents(string memory _agentId) public view returns (uint256) {
    return premisAgents[_agentId].events.length;
  }

  function getNumRights(string memory _agentId) public view returns (uint256) {
    return premisAgents[_agentId].rights.length;
  }

  function getAgent(uint256 _index) public view returns (string memory) {
    return agents[_index];
  }

  function getName(string memory _agentId) public view returns (string memory, string memory) {
    return (_agentId, premisAgents[_agentId].name);
  }

  function getType(string memory _agentId) public view returns (string memory, string memory) {
    return (_agentId, premisAgents[_agentId].agentType);
  }

  function getObject(string memory _agentId, uint256 _objectIndex) public view returns (string memory, string memory) {
    return (_agentId, premisAgents[_agentId].objects[_objectIndex]);
  }

  function getEvent(string memory _agentId, uint256 _eventIndex) public view returns (string memory, string memory) {
    return (_agentId, premisAgents[_agentId].events[_eventIndex]);
  }

  function getRights(string memory _agentId, uint256 _rightsIndex) public view returns (string memory, string memory) {
    return (_agentId, premisAgents[_agentId].rights[_rightsIndex]);
  }
}
