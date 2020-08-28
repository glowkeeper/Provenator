pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IEntity {

    function amend(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function addType(EntityTypes _type) virtual public;
    function addRelation(bytes32 _id) virtual public;

    function get() virtual public view returns (CreativeEntities memory);
    function getTypes() virtual public view returns (bool[] memory);

    function isType(EntityTypes _type) virtual public view returns (bool);
    function containsRelation(bytes32 _id) virtual public view returns (bool);
}

abstract contract IEntitiesFactory {

    function addEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function amendEntity(CreativeEntities memory _entity, EntityTypes _type) virtual public;
    function addEntityRelation(bytes32 _parentId, bytes32 _childId) virtual public;

    function getEntity(bytes32 _id) virtual public view returns (CreativeEntities memory);
    function getEntityTypes(bytes32 _id) virtual public view returns (bool[] memory);
    function getEntityContract(bytes32 _id) virtual public view returns (address);

    function isType(bytes32 _id, EntityTypes _type) virtual public view returns (bool);
    function containsRelation(bytes32 _parentId, bytes32 _childId) virtual public view returns (bool);

    function getRelationsNum(bytes32 _id) virtual public view returns (uint256);
    function getRelationsReference(bytes32 _id, uint256 _index) virtual public view returns (bytes32);
    function getRelations(bytes32 _id) virtual public view returns (bytes32[] memory);
}
