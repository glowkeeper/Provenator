pragma solidity ^0.4.24;

contract Event {

  function setEventType(string _eventType) public;
  function setEvent(string _eventId, string _type, string _time, string _objectHash, string _agentId) public;

  function getEventTypeExists(string _eventType) public constant returns (uint256);
  function getEventExists(string _eventId) public constant returns (uint256);

  function getNumEventTypes() public constant returns (uint256);
  function getEventType(uint256 _eventTypeIndex) public constant returns (string);

  function getNumEvents() public constant returns (uint256);
  function getEvent(uint256 _eventIndex) public constant returns (string);

  function getObject(string _eventId) public constant returns (string, string);
  function getType(string _eventId) public constant returns (string, string);
  function getTime(string _eventId) public constant returns (string, string);
  function getAgent(string _eventId) public constant returns (string, string);

}
