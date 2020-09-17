pragma solidity >=0.4.21 <0.8.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IEntity {

    function amend(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function addRelation(bytes32 _relatedId) virtual public;
    function addType(EntityTypes _type) virtual public;

    function get() virtual public view returns (CreativeEntities memory);
    function getRelations() virtual public view returns (bytes32[] memory);
    function getTypes() virtual public view returns (bool[] memory);

    function isRelation(bytes32 _relatedEntityId) virtual public view returns (bool);
    function isType(EntityTypes _type) virtual public view returns (bool);
}

abstract contract IEntitiesFactory {

    function addEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function addEntityRelation(bytes32 _parentId, bytes32 _relatedId) virtual public;
    function amendEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;

    function getEntity(bytes32 _id) virtual public view returns (CreativeEntities memory);
    function getEntityRelations(bytes32 _id) virtual public view returns (bytes32[] memory);
    function getEntityTypes(bytes32 _id) virtual public view returns (bool[] memory);
    function getEntityContract(bytes32 _id) virtual public view returns (address);

    function isEntityRelation(bytes32 _id, bytes32 _relatedEntityId) virtual public view returns (bool);
    function isEntityType(bytes32 _id, EntityTypes _type) virtual public view returns (bool);
}
