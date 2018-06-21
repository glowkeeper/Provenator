pragma solidity ^0.4.24;

contract Rights {

  function setObject(string _rightsId, string _objectHash) public;
  function setBasis(string _rightsId, string _basis) public;
  function setCopyrightInfo(string _rightsId,
                            string _status,
                            string _jurisdictionCountryCode,
                            string _determinationDate,
                            string _note ) public;
  function setRightsGranted(string _rightsId, string _act, string _restriction) public;
  function setAgent(string _rightsId, string _agentId) public;

  function getRightsExists(string _rightsId) public constant returns (bool);

  function getNumRights() public constant returns (uint256);
  function getRights(uint256 _index) public constant returns (string);

  function getObject(string _rightsId) public constant returns (string, string);
  function getRightsBasis(string _rightsId) public constant returns (string, string);
  function getCopyrightStatus(string _rightsId) public constant returns (string, string);
  function getCopyrightJurisdiction(string _rightsId) public constant returns (string, string);
  function getCopyrightDeterminationDate(string _rightsId) public constant returns (string, string);
  function getCopyrightNote(string _rightsId) public constant returns (string, string);
  function getGrantedAct(string _rightsId) public constant returns (string, string);
  function getGrantedRestriction(string _rightsId) public constant returns (string, string);
  function getAgent(string _rightsId) public constant returns (string, string);
}
