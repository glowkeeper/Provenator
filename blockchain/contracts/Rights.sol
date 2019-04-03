pragma solidity ^0.5.0;

contract Rights {

  function setObject(string memory _rightsId, string memory _objectHash) public;
  function setBasis(string memory _rightsId, string memory _basis) public;
  function setCopyrightInfo(string memory _rightsId,
                            string memory _status,
                            string memory _jurisdictionCountryCode,
                            string memory _determinationDate,
                            string memory _note ) public;
  function setRightsGranted(string memory _rightsId, string memory _act, string memory _restriction) public;
  function setAgent(string memory _rightsId, string memory _agentId) public;

  function getRightsExists(string memory _rightsId) public view returns (bool);

  function getNumRights() public view returns (uint256);
  function getRights(uint256 _index) public view returns (string memory);

  function getObject(string memory _rightsId) public view returns (string memory, string memory);
  function getRightsBasis(string memory _rightsId) public view returns (string memory, string memory);
  function getCopyrightStatus(string memory _rightsId) public view returns (string memory, string memory);
  function getCopyrightJurisdiction(string memory _rightsId) public view returns (string memory, string memory);
  function getCopyrightDeterminationDate(string memory _rightsId) public view returns (string memory, string memory);
  function getCopyrightNote(string memory _rightsId) public view returns (string memory, string memory);
  function getGrantedAct(string memory _rightsId) public view returns (string memory, string memory);
  function getGrantedRestriction(string memory _rightsId) public view returns (string memory, string memory);
  function getAgent(string memory _rightsId) public view returns (string memory, string memory);
}
