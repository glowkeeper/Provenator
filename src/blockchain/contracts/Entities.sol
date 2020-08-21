pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";
import "./IterableData.sol";
import "./Enums.sol";
import "./Storage.sol";
import "./IFactory.sol";
import "./IEntities.sol";

contract EntityNode is IEntity {

    EntityTypes entityType;
    bytes32     id;
    Entity      entity;

    constructor (CreativeEntities memory _entity, EntityTypes _entityType) public {
        require (
            _entityType > EntityTypes.NONE &&
            _entityType < EntityTypes.MAX &&
            _entity.id[0] != 0
        );

        entityType = _entityType;
        id = _entity.id;
        set(_entity);
    }

    function get() override virtual public view returns (CreativeEntities memory) {

        CreativeEntities memory entities;

        entities.id = id;
        entities.name = entity.name;
        entities.email = entity.email;
        entities.url = entity.url;

        return entities;
    }

    function getType() override virtual public view returns (EntityTypes) {
        return entityType;
    }

    function amend(CreativeEntities memory _entity) override virtual public {
        require ( id == _entity.id );

        set(_entity);
    }

    function set(CreativeEntities memory _entity) private {
        require (
             bytes(_entity.name).length > 0 &&
             bytes(_entity.email).length > 0
        );

        entity.entityType = entityType;
        entity.name = _entity.name;
        entity.email = _entity.email;
        entity.url = _entity.url;
    }
}

contract Entities is IEntitiesFactory, IFactory {

    Data entityStore;
    using IterableData for Data;

    constructor() public {

      entityStore.size = 0;
    }

    function addEntity(CreativeEntities memory _entity, EntityTypes _entityType) override virtual public {

        if (entityStore.contains(_entity.id)) {

            amendEntity(_entity);

        } else {

            EntityNode entity = new EntityNode(_entity, _entityType);
            entityStore.insert(_entity.id, address(entity));
        }
    }

    function amendEntity(CreativeEntities memory _entity) override virtual public {
        require ( entityStore.data[_entity.id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_entity.id].value);
        entity.amend(_entity);
    }

    function getEntity(bytes32 _id) override virtual public view returns (CreativeEntities memory) {
        require ( entityStore.data[_id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_id].value);
        return entity.get();
    }

    function getEntityContract(bytes32 _id) override virtual public view returns (address) {
        require( entityStore.data[_id].value != address(0x0) );

        return entityStore.data[_id].value;
    }

    function getNum() override virtual public view returns (uint256)
    {
        return entityStore.size;
    }

    function getReference(uint256 _index) override virtual public view returns (bytes32)
    {
        require (_index < entityStore.size);

        return entityStore.keys[_index].key;
    }
}
