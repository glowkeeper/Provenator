pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";
import "./IterableData.sol";
import "./Enums.sol";
import "./Storage.sol";
import "./IFactory.sol";
import "./IEntities.sol";

contract EntityNode is IEntity, IFactory {

    uint8 constant typesSize = uint8(EntityTypes.MAX);
    CreativeEntities               entity;
    bool[]                         entityTypes;
    bytes32[]                      relationStore;

    constructor (CreativeEntities memory _entity, EntityTypes _entityType) public {
        require (
            _entity.id[0] != 0 &&
            bytes(_entity.name).length > 0 &&
            bytes(_entity.email).length > 0 &&
            _entityType > EntityTypes.NONE &&
            _entityType < EntityTypes.MAX
        );

        entityTypes = new bool[](typesSize);
        entityTypes[uint8(_entityType)] = true;
        entity = _entity;
    }

    function amend(CreativeEntities memory _entity, EntityTypes _entityType) override virtual public {

        addType(_entityType);
        entity = _entity;
    }

    function addType(EntityTypes _type) override virtual public {
        require (
            _type > EntityTypes.NONE &&
            _type < EntityTypes.MAX
        );

        if ( !isType(_type) ) {
            entityTypes[uint8(_type)] = true;
        }
    }

    function addRelation(bytes32 _id) override virtual public {
        require ( _id[0] != 0 );

        if( (entity.id != _id ) && (!containsRelation(_id)) ) {
            relationStore.push(_id);
        }
    }

    function get() override virtual public view returns (CreativeEntities memory) {

        return entity;
    }

    function getTypes() override virtual public view returns (bool[] memory) {

        return entityTypes;
    }

    function isType(EntityTypes _type) override virtual public view returns (bool) {
        require (
            _type > EntityTypes.NONE &&
            _type < EntityTypes.MAX
        );

        return entityTypes[uint8(_type)];
    }

    function containsRelation(bytes32 _id) override virtual public view returns (bool) {
        require ( _id[0] != 0 );

        bool found = false;
        for (uint i = 0; i < relationStore.length; i++) {
           if ( relationStore[i] == _id ) {
               found = true;
               break;
           }
        }
        return found;
    }

    function getNum() override virtual public view returns (uint256) {
        return relationStore.length;
    }

    function getReference(uint256 _index) override virtual public view returns (bytes32) {
        require (_index < relationStore.length);

        return relationStore[_index];
    }

    function getReferences() override virtual public view returns (bytes32[] memory) {

        return relationStore;
    }
}

contract Entities is IEntitiesFactory, IFactory {

    uint8 constant maxEntityTypes = uint8(EntityTypes.MAX);
    Data entityStore;
    using IterableData for Data;

    constructor() public {

      entityStore.size = 0;
    }

    function addEntity(CreativeEntities memory _entity, EntityTypes _entityType) override virtual public {

        if (entityStore.contains(_entity.id)) {

            amendEntity(_entity, _entityType);

        } else {

            EntityNode entity = new EntityNode(_entity, _entityType);
            entityStore.insert(_entity.id, address(entity));
        }
    }

    function amendEntity(CreativeEntities memory _entity, EntityTypes _type) override virtual public {
        require ( entityStore.data[_entity.id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_entity.id].value);
        entity.amend(_entity, _type);
    }


    function addEntityRelation(bytes32 _parentId, bytes32 _childId) override virtual public {
        require ( entityStore.data[_parentId].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_parentId].value);
        entity.addRelation(_childId);
    }

    function getEntity(bytes32 _id) override virtual public view returns (CreativeEntities memory) {
        require ( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.get();
    }

    function getEntityTypes(bytes32 _id) override virtual public view returns (bool[] memory) {
        require ( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.getTypes();
    }

    function getEntityContract(bytes32 _id) override virtual public view returns (address) {
        require( entityStore.data[_id].value != address(0x0) );

        return entityStore.data[_id].value;
    }

    function getNum() override virtual public view returns (uint256) {
        return entityStore.size;
    }

    function getReference(uint256 _index) override virtual public view returns (bytes32) {
        require (_index < entityStore.size);

        return entityStore.keys[_index].key;
    }

    function getReferences() override virtual public view returns (bytes32[] memory) {

        bytes32[] memory refs = new bytes32[](entityStore.size);
        for (uint i = 0; i < entityStore.size; i++) {
            refs[i] = entityStore.keys[i].key;
        }
        return refs;
    }

    function isEntityType(bytes32 _id, EntityTypes _type) override virtual public view returns (bool) {
        require( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.isType(_type);
    }

    function containsEntityRelation(bytes32 _parentId, bytes32 _childId) override virtual public view returns (bool) {
        require( entityStore.data[_parentId].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_parentId].value);
        return entity.containsRelation(_childId);
    }

    function getRelationsNum(bytes32 _id) override virtual public view returns (uint256) {
        require( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.getNum();
    }

    function getRelationsReference(bytes32 _id, uint256 _index) override virtual public view returns (bytes32) {
        require( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.getReference(_index);
    }

    function getRelations(bytes32 _id) override virtual public view returns (bytes32[] memory) {
        require( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.getReferences();
    }


}
