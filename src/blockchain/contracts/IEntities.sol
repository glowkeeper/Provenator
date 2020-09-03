pragma solidity >=0.4.21 <0.8.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IEntity {

    function amend(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function addType(EntityTypes _type) virtual public;

    function get() virtual public view returns (CreativeEntities memory);
    function getTypes() virtual public view returns (bool[] memory);

    function isType(EntityTypes _type) virtual public view returns (bool);
}

abstract contract IEntitiesFactory {

    function addEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function amendEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;

    function getEntity(bytes32 _id) virtual public view returns (CreativeEntities memory);
    function getEntityTypes(bytes32 _id) virtual public view returns (bool[] memory);
    function getEntityContract(bytes32 _id) virtual public view returns (address);

    function isEntityType(bytes32 _id, EntityTypes _type) virtual public view returns (bool);
}
