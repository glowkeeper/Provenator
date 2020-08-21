pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IFactory {

    function getNum() virtual public view returns (uint256);
    function getReference(uint256 _index) virtual public view returns (bytes32);
}
