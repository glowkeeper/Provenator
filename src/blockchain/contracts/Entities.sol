pragma solidity >=0.4.21 <0.8.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";
import "./IterableData.sol";
import "./Enums.sol";
import "./Storage.sol";
import "./IFactory.sol";
import "./IEntities.sol";

contract EntityNode is IEntity {

    uint8 constant typesSize = uint8(EntityTypes.MAX);
    CreativeEntities               entity;
    bytes32[]                      relations;
    bool[]                         entityTypes;

    constructor (CreativeEntities memory _entity, EntityTypes _entityType) public {
        require (
            _entity.id[0] != 0 &&
            bytes(_entity.name).length > 0 &&
            bytes(_entity.email).length > 0 &&
            _entityType > EntityTypes.NONE &&
            _entityType < EntityTypes.MAX
        );

        entityTypes = new bool[](typesSize);
        entity = _entity;
        entityTypes[uint8(_entityType)] = true;
    }

    function amend(CreativeEntities memory _entity, EntityTypes _entityType) override virtual public {

        entity = _entity;
        addType(_entityType);
    }

    function addRelation(bytes32 _relatedId) override virtual public {
        require (
            isType(EntityTypes.AUTHOR) &&
            _relatedId[0] != 0
        );

        if ( ( _relatedId != entity.id ) && !isRelation(_relatedId) ) {
            relations.push(_relatedId);
        }
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

    function get() override virtual public view returns (CreativeEntities memory) {

        return entity;
    }

    function getRelations() override virtual public view returns (bytes32[] memory) {

        return relations;
    }

    function getTypes() override virtual public view returns (bool[] memory) {

        return entityTypes;
    }

    function isRelation(bytes32 _relatedEntityId) override virtual public view returns (bool) {
        require (
            _relatedEntityId[0] != 0
        );

        bool found = false;
        for (uint i = 0; i < relations.length; i++) {
           if ( relations[i] == _relatedEntityId ) {
               found = true;
               break;
           }
        }

        return found;
    }

    function isType(EntityTypes _type) override virtual public view returns (bool) {
        require (
            _type > EntityTypes.NONE &&
            _type < EntityTypes.MAX
        );

        return entityTypes[uint8(_type)];
    }
}

contract Entities is IEntitiesFactory, IFactory {

    uint8 constant maxEntityTypes = uint8(EntityTypes.MAX);
    Data entityStore;
    using IterableData for Data;

    constructor() public {

      entityStore.size = 0;
    }

    function addEntity(CreativeEntities memory _entity, EntityTypes _entityType, bytes32 _parentId) override virtual public {

        if (entityStore.contains(_entity.id)) {

            amendEntity(_entity, _entityType, _parentId);

        } else {

            EntityNode entity = new EntityNode(_entity, _entityType);
            entityStore.insert(_entity.id, address(entity));

            if ( _parentId[0] != 0 ) {

                require ( entityStore.data[_parentId].value != address(0x0) );
                EntityNode parentEntity = EntityNode(entityStore.data[_parentId].value);
                parentEntity.addRelation(_entity.id);
            }
        }
    }

    function amendEntity(CreativeEntities memory _entity, EntityTypes _type, bytes32 _parentId) override virtual public {
        require (
            entityStore.data[_entity.id].value != address(0x0) &&
            entityStore.data[_parentId].value != address(0x0)
        );

        EntityNode entity = EntityNode(entityStore.data[_entity.id].value);
        entity.amend(_entity, _type);

        if ( _parentId[0] != 0 ) {

            require ( entityStore.data[_parentId].value != address(0x0) );
            EntityNode parentEntity = EntityNode(entityStore.data[_parentId].value);
            parentEntity.addRelation(_entity.id);
        }
    }

    function getEntity(bytes32 _id) override virtual public view returns (CreativeEntities memory) {
        require ( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.get();
    }

    function getEntityRelations(bytes32 _id) override virtual public view returns (bytes32[] memory) {
        require (
            entityStore.data[_id].value != address(0x0)
        );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.getRelations();
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

    function isEntityRelation(bytes32 _id, bytes32 _relatedEntityId) override virtual public view returns (bool) {
        require(
            entityStore.data[_id].value != address(0x0)
        );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.isRelation(_relatedEntityId);
    }

    function isEntityType(bytes32 _id, EntityTypes _type) override virtual public view returns (bool) {
        require( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.isType(_type);
    }
}
