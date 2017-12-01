pragma solidity ^0.4.11;

contract Event {

  function setEventType(string _eventType);
  function setEvent(string _eventId, string _type, string _time, string _objectHash, string _agentId);

  function getEventTypeExists(string _eventType) constant returns (uint256);
  function getEventExists(string _eventId) constant returns (uint256);

  function getNumEventTypes() constant returns (uint256);
  function getEventType(uint256 _eventTypeIndex) constant returns (string);

  function getNumEvents() constant returns (uint256);
  function getEvent(uint256 _eventIndex) constant returns (string);

  function getObject(string _eventId) public constant returns (string, string);
  function getType(string _eventId) public constant returns (string, string);
  function getTime(string _eventId) public constant returns (string, string);
  function getAgent(string _eventId) public constant returns (string, string);

}
