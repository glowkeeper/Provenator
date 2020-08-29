pragma solidity >=0.4.21 <0.8.0;
pragma experimental ABIEncoderV2;

contract Migrations {
  address public owner;
  uint public last_completed_migration;

  constructor() virtual public {
    owner = msg.sender;
  }

  modifier restricted() {
    if (msg.sender == owner) _;
  }

  function setCompleted(uint completed) virtual public restricted {
    last_completed_migration = completed;
  }

  function upgrade(address new_address) virtual public restricted {
    Migrations upgraded = Migrations(new_address);
    upgraded.setCompleted(last_completed_migration);
  }
}
