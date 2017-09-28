pragma solidity ^0.4.11;

import "Owned.sol";

contract Mortal is Owned {

  function kill() onlyOwner {
    suicide(owner);
  }
}
