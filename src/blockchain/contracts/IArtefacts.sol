pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IArtefact {

    function get() virtual public view returns (CreativeWorks memory);
    function amend(CreativeWorks memory _work) virtual public;
}

abstract contract IArtefactFactory {

    function addWork(CreativeWorks memory _work) virtual public;
    function amendWork(CreativeWorks memory _work) virtual public;
    function getWork(bytes32 _id) virtual public view returns (CreativeWorks memory);
    function getWorkContract(bytes32 _id) virtual public view returns (address);
}
