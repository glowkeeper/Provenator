pragma solidity ^0.5.0;

// Fake News Premis Object
// Steve Huckle

import "./Event.sol";
import "./Strings.sol";

contract PremisEvent is Event {

  struct EventData {
    string eventType;
    string time;
    string objectHash;
    string agentId;
  }

  string[] private eventTypes;
  string[] events;
  mapping(string => EventData) private premisEvents;

  event EventSet(string _eventId, string _type, string _time, string _objectHash, string _agentId);

  constructor() public {
  }

  function setEventType(string memory _eventType) public {
    eventTypes.push(_eventType);
  }

  function setEvent(string memory _eventId, string memory _type, string memory _time, string memory _objectHash, string memory _agentId) public {
    premisEvents[_eventId].eventType = _type;
    premisEvents[_eventId].time = _time;
    premisEvents[_eventId].objectHash = _objectHash;
    premisEvents[_eventId].agentId = _agentId;
    events.push(_eventId);
    emit EventSet(_eventId, _type, _time, _objectHash, _agentId);
  }

  // These 'Exists?' functions return max(uint256) if not, the index otherwise
  function getEventTypeExists(string memory _eventType) public view returns (bool) {
      uint256 index = Strings.getIndex(_eventType, eventTypes);
      return index != eventTypes.length;
  }


  function getEventExists(string memory _eventId) public view returns (bool) {
    uint256 index = Strings.getIndex(_eventId, events);
    return index != events.length;
  }


  function getNumEventTypes() public view returns (uint256) {
    return eventTypes.length;
  }

  function getEventType(uint256 _eventTypeIndex) public view returns (string memory) {
    return eventTypes[_eventTypeIndex];
  }

  function getNumEvents() public view returns (uint256) {
    return events.length;
  }

  function getEvent(uint256 _eventIndex) public view returns (string memory) {
    return events[_eventIndex];
  }

  function getObject(string memory _eventId) public view returns (string memory, string memory) {
    return (_eventId, premisEvents[_eventId].objectHash);
  }

  function getType(string memory _eventId) public view returns (string memory, string memory) {
    return (_eventId, premisEvents[_eventId].eventType);
  }

  function getTime(string memory _eventId) public view returns (string memory, string memory) {
    return (_eventId, premisEvents[_eventId].time);
  }

  function getAgent(string memory _eventId) public view returns (string memory, string memory) {
    return (_eventId, premisEvents[_eventId].agentId);
  }

}
