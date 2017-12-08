pragma solidity ^0.4.11;

// Fake News Premis Object
// Steve Huckle

import "Event.sol";
import "Strings.sol";

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

  /* function PremisEvent() {
  } */

  function setEventType(string _eventType) public {
    eventTypes.push(_eventType);
  }

  function setEvent(string _eventId, string _type, string _time, string _objectHash, string _agentId) public {
    premisEvents[_eventId].eventType = _type;
    premisEvents[_eventId].time = _time;
    premisEvents[_eventId].objectHash = _objectHash;
    premisEvents[_eventId].agentId = _agentId;
    events.push(_eventId);
    EventSet(_eventId, _type, _time, _objectHash, _agentId);
  }

  // These 'Exists?' functions return max(uint256) if not, the index otherwise
  function getEventTypeExists(string _eventType) public constant returns (uint256) {
    return Strings.getIndex(_eventType, eventTypes);
  }


  function getEventExists(string _eventId) public constant returns (uint256) {
    return Strings.getIndex(_eventId, events);
  }


  function getNumEventTypes() public constant returns (uint256) {
    return eventTypes.length;
  }

  function getEventType(uint256 _eventTypeIndex) public constant returns (string) {
    return eventTypes[_eventTypeIndex];
  }

  function getNumEvents() public constant returns (uint256) {
    return events.length;
  }

  function getEvent(uint256 _eventIndex) public constant returns (string) {
    return events[_eventIndex];
  }

  function getObject(string _eventId) public constant returns (string, string) {
    return (_eventId, premisEvents[_eventId].objectHash);
  }

  function getType(string _eventId) public constant returns (string, string) {
    return (_eventId, premisEvents[_eventId].eventType);
  }

  function getTime(string _eventId) public constant returns (string, string) {
    return (_eventId, premisEvents[_eventId].time);
  }

  function getAgent(string _eventId) public constant returns (string, string) {
    return (_eventId, premisEvents[_eventId].agentId);
  }

}
