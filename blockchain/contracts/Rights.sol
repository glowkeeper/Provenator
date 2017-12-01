pragma solidity ^0.4.11;

contract Rights {

  function setObject(string _rightsId, string _objectHash);
  function setBasis(string _rightsId, string _basis);
  function setCopyrightInfo(string _rightsId,
                            string _status,
                            string _jurisdictionCountryCode,
                            string _determinationDate,
                            string _note );
  function setRightsGranted(string _rightsId, string _act, string _restriction);
  function setAgent(string _rightsId, string _agentId);

  function getRightsExists(string _rightsId) constant returns (uint256);

  function getNumRights() constant returns (uint256);
  function getRights(uint256 _index) constant returns (string);

  function getObject(string _rightsId) constant returns (string, string);
  function getRightsBasis(string _rightsId) constant returns (string, string);
  function getCopyrightStatus(string _rightsId) constant returns (string, string);
  function getCopyrightJurisdiction(string _rightsId) constant returns (string, string);
  function getCopyrightDeterminationDate(string _rightsId) constant returns (string, string);
  function getCopyrightNote(string _rightsId) constant returns (string, string);
  function getGrantedAct(string _rightsId) constant returns (string, string);
  function getGrantedRestriction(string _rightsId) constant returns (string, string);
  function getAgent(string _rightsId) constant returns (string, string);
}
