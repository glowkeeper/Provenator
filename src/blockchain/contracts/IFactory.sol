pragma solidity >=0.4.21 <0.8.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IFactory {

    function getNum() virtual public view returns (uint256);
    function getReference(uint256 _index) virtual public view returns (bytes32);
    function getReferences() virtual public view returns (bytes32[] memory);
}
