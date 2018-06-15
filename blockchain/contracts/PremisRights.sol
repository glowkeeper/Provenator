pragma solidity ^0.4.19;

// Fake News Premis Object
// Steve Huckle

import "Rights.sol";
import "Strings.sol";

contract PremisRights is Rights {

  struct CopyrightInfo {
    string status;
    string jurisdictionCountryCode;
    string determinationDate;
    string note;
  }

  struct RightsGranted {
    string act;
    string restriction;
  }

  struct RightsData {
    string basis;
    CopyrightInfo copyrightData;
    RightsGranted rightsGrantedData;
    string objectHash;
    string agentId;
  }

  string[] rights;
  mapping (string => RightsData) premisRights;

  event RightsSetObject(string _rightsId, string _objectHash);
  event RightsSetBasis(string _rightsId, string _basis);
  event RightsSetCopyrightInfo(string _rightsId,
                               string _status,
                               string _jurisdictionCountryCode,
                               string _determinationDate,
                               string _note);
  event RightsSetRightsGranted(string _rightsId, string _act, string _restriction);
  event RightsSetAgent(string _rightsId, string _agentId);

  constructor() public {
  }

  function setObject(string _rightsId, string _objectHash) public {
    rights.push(_rightsId);
    premisRights[_rightsId].objectHash = _objectHash;
    emit RightsSetObject(_rightsId, _objectHash);
  }

  function setBasis(string _rightsId, string _basis) public {
    premisRights[_rightsId].basis = _basis;
    emit RightsSetBasis(_rightsId, _basis);
  }

  function setCopyrightInfo(string _rightsId, string _status, string _jurisdictionCountryCode, string _determinationDate, string _note ) public {
    premisRights[_rightsId].copyrightData.status = _status;
    premisRights[_rightsId].copyrightData.jurisdictionCountryCode = _jurisdictionCountryCode;
    premisRights[_rightsId].copyrightData.determinationDate = _determinationDate;
    premisRights[_rightsId].copyrightData.note = _note;
    emit RightsSetCopyrightInfo(_rightsId, _status, _jurisdictionCountryCode, _determinationDate, _note);
  }

  function setRightsGranted(string _rightsId, string _act, string _restriction) public {
    premisRights[_rightsId].rightsGrantedData.act = _act;
    premisRights[_rightsId].rightsGrantedData.restriction = _restriction;
    emit RightsSetRightsGranted(_rightsId, _act, _restriction);
  }

  function setAgent(string _rightsId, string _agentId) public {
    premisRights[_rightsId].agentId = _agentId;
    emit RightsSetAgent(_rightsId, _agentId);
  }

  // This 'Exists?' function return max(uint256) if not, the index otherwise
  function getRightsExists(string _rightsId) public constant returns (uint256) {
    return Strings.getIndex(_rightsId, rights);
  }

  function getNumRights() public constant returns (uint256) {
    return rights.length;
  }

  function getRights(uint256 _index) public constant returns (string) {
    return rights[_index];
  }

  function getObject(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].objectHash);
  }

  function getRightsBasis(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].basis);
  }

  function getCopyrightStatus(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].copyrightData.status);
  }
  function getCopyrightJurisdiction(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].copyrightData.jurisdictionCountryCode);
  }

  function getCopyrightDeterminationDate(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].copyrightData.determinationDate);
  }

  function getCopyrightNote(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].copyrightData.note);
  }

  function getGrantedAct(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].rightsGrantedData.act);
  }

  function getGrantedRestriction(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].rightsGrantedData.restriction);
  }

  function getAgent(string _rightsId) public constant returns (string, string) {
    return (_rightsId, premisRights[_rightsId].agentId);
  }

}
