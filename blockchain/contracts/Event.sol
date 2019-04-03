pragma solidity ^0.5.0;

contract Event {

  function setEventType(string memory _eventType) public;
  function setEvent(string memory _eventId, string memory _type, string memory _time, string memory _objectHash, string memory _agentId) public;

  function getEventTypeExists(string memory _eventType) public view returns (bool);
  function getEventExists(string memory _eventId) public view returns (bool);

  function getNumEventTypes() public view returns (uint256);
  function getEventType(uint256 _eventTypeIndex) public view returns (string memory);

  function getNumEvents() public view returns (uint256);
  function getEvent(uint256 _eventIndex) public view returns (string memory);

  function getObject(string memory _eventId) public view returns (string memory, string memory);
  function getType(string memory _eventId) public view returns (string memory, string memory);
  function getTime(string memory _eventId) public view returns (string memory, string memory);
  function getAgent(string memory _eventId) public view returns (string memory, string memory);

}
