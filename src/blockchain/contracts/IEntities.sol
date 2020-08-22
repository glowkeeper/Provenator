pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IEntity {

    function get() virtual public view returns (CreativeEntities memory);
    function getType() virtual public view returns (EntityTypes);
    function amend(CreativeEntities memory _entity, EntityTypes _type) virtual public;
}

abstract contract IEntitiesFactory {

    function addEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function amendEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function getEntity(bytes32 _id) virtual public view returns (CreativeEntities memory);
    function getEntityType(bytes32 _id) virtual public view returns (EntityTypes);
    function getEntityContract(bytes32 _id) virtual public view returns (address);
}
