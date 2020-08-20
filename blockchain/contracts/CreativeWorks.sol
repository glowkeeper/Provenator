pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";
import "./IterableData.sol";
import "./Enums.sol";
import "./Storage.sol";
import "./IFactory.sol";
import "./ICreativeWorks.sol";

contract EntityNode is IEntity {

    EntityTypes type;
    bytes32     id
    Entity      entity;

    constructor (Entities memory _entity, EntityTypes _type) public {
        require (
            _type > EntityTypes.NONE &&
            _type < EntityTypes.MAX &&
            _work.id[0] != 0
        );

        type = _type;
        id = _work.id;
        set(_entity)
    }

    function get() virtual public view returns (Entities memory) {

        Entities memory entities;

        entities.id = id
        entities.name = entity.name;
        entities.email = entity.email;
        entities.url = entity.url;

        return entities;
    }

    function getType() virtual public view returns (EntityTypes) {
        return type;
    }

    function amend(Entities memory _entity) virtual public {
        require ( id == _work.id );

        set(_entity)
    }

    function set(Entities memory _entity) virtual private {
        require (
             bytes(_entity.name).length > 0 &&
             bytes(_entity.email).length > 0
        );

        entity.type = type;
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

    function addEntity(Entities memory _entity, EntityTypes _type) override virtual public {

        EntityNode entity = new EntityNode(_entity, _type);
        entityStore.insert(_entity.id, address(entity));
    }

    function amendEntity(Entities memory _entity) override virtual public {
        require ( entityStore.data[_entity.id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_entity.id].value);
        entity.amend(_entity);
    }

    function getEntity(bytes32 _id) override virtual public view returns (Entities memory) {
        require ( entityStore.data[_entity.id].value != address(0x0) );

        EntityNode entity = EntityNode(entityStore.data[_entity.id].value);
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

contract WorksNode is IWorks {

    IEntitiesFactory    entities;
    bytes32             id;
    Works               work;

    constructor (CreativeWorks memory _work) public {
        require ( _work.id[0] != 0 );

        id = _work.id;
        setWork(_work)
    }

    function get() virtual public view returns (CreativeWorks memory) {

        CreativeWorks memory works;

        works.type = work.type;
        works.license = work.license;
        works.id = id;
        works.dateCreated  = work.dateCreated;
        works.dateModified = work.dateModified;
        works.url = work.url;
        works.name = work.name;
        works.description = work.description;

        works.author = entities.getEntity(work.author);

        if ( work.copyrightHolder[0] != 0 ) {
            EntityNode entity = EntityNode(entities.getEntityContract(work.copyrightHolder);
            require ( entity.getType() == EntityTypes.COPYRIGHTHOLDER );
            works.copyrightHolder = entity.get();
        } else {
            works.copyrightHolder.id = bytes32(0);
            works.copyrightHolder.name = "";
            works.copyrightHolder.email = "";
            works.copyrightHolder.url = "";
        }

        if ( work.publisher[0] != 0 ) {
            EntityNode entity = EntityNode(entities.getEntityContract(work.publisher);
            require ( entity.getType() == EntityTypes.PUBLISHER );
            works.publisher = entity.get();
        } else {
            works.publisher.id = bytes32(0);
            works.publisher.name = "";
            works.publisher.email = "";
            works.publisher.url = "";
        }

        return works;
    }

    function amend(CreativeWorks memory _work) virtual public {
        require ( id == _work.id );

        set(_work)
    }

    function set(CreativeWorks memory _work) override virtual private {
        require (
            _work.type > WorksTypes.NONE &&
            _work.type < WorksTypes.MAX &&
            _work.license > LicenseTypes.NONE &&
            _work.license < LicenseTypes.MAX &&
            _work.dateCreated[0] != 0 &&
             bytes(_work.name).length > 0 &&
             bytes(_work.description).length > 0 &&
             work.author.id[0] != 0 &&
        );

        work.type = _work.type;
        work.license = _work.license;
        work.dateCreated = _work.dateCreated;
        work.dateModified = _work.dateModified;
        work.author = _work.author.id;
        work.copyrightHolder = _work.copyrightHolder.id;
        work.publisher = _work.publisher.id;
        work.url = _work.type;
        work.name = _work.type;
        work.description = _work.type;

        entities.addEntity(_work.author, EntityTypes.AUTHOR)

        if ( work.copyrightHolder[0] != 0 ) {
            entities.addEntity(_work.copyrightHolder, EntityTypes.COPYRIGHTHOLDER);
        }

        if ( work.publisher[0] != 0 ) {
            entities.addEntity(_work.publisher, EntityTypes.PUBLISHER);
        }
    }
}

contract Works is IWorksFactory, IFactory {

    IEntitiesFactory entities;
    Data workStore;
    using IterableData for Data;

    constructor(address _entityFactor) public {
        require (
            _entityFactor != address(0x0)
        );

        entities = IEntitiesFactory(_entityFactor);
        workStore.size = 0;
    }

    function addWork(CreativeWorks memory _work) override virtual public {
        WorksNode work = new WorksNode(_work, entities);
        workStore.insert(_work.id, address(work));
    }

    function amendWork(CreativeWorks memory _work) virtual public {
        require ( workStore.data[_work.id].value != address(0x0) );

        WorksNode work = WorksNode(workStore.data[_work.id].value);
        work.amend(_work);
    }

    function getWork(bytes32 _id) override virtual public view returns (CreativeWorks memory) {
        require ( workStore.data[_id].value != address(0x0)
        );

        WorksNode work = WorksNode(workStore.data[_id].value);
        return work.get();
    }

    function getWorkContract(bytes32 _id) override virtual public view returns (address) {
        require( workStore.data[_id].value != address(0x0) );

        return workStore.data[_id].value;
    }

    function getNum() override virtual public view returns (uint256)
    {
        return workStore.size;
    }

    function getReference(uint256 _index) override virtual public view returns (bytes32)
    {
        require (_index < workStore.size);

        return workStore.keys[_index].key;
    }
}
