pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IEntity {

    function get() virtual public view returns (CreativeEntities memory);
    function getType() virtual public view returns (EntityTypes);
    function amend(CreativeEntities memory _entity) virtual public;
}

abstract contract IEntitiesFactory {

    function addEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function amendEntity(CreativeEntities memory _entity) virtual public;
    function getEntity(bytes32 _id) virtual public view returns (CreativeEntities memory);
    function getEntityContract(bytes32 _id) virtual public view returns (address);
}

abstract contract IMedia {

    function get() virtual public view returns (CreativeWorks memory);
    function amend(CreativeWorks memory _work) virtual public;
}

abstract contract IMediaFactory {

    function addWork(CreativeWorks memory _work) virtual public;
    function amendWork(CreativeWorks memory _work) virtual public;
    function getWork(bytes32 _id) virtual public view returns (CreativeWorks memory);
    function getWorkContract(bytes32 _id) virtual public view returns (address);
}
