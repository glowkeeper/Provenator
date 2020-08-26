pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";
import "./IterableData.sol";
import "./Enums.sol";
import "./Storage.sol";
import "./IFactory.sol";
import "./IArtefacts.sol";
import "./IEntities.sol";
import "./Entities.sol";

contract ArtefactNode is IArtefact {

    IEntitiesFactory    entities;
    bytes32             id;
    Works               work;

    constructor (CreativeWorks memory _work, address _entityFactory) public {
        require ( _work.id[0] != 0 );

        entities = IEntitiesFactory(_entityFactory);
        id = _work.id;
        set(_work);
    }

    function amend(CreativeWorks memory _work) override virtual public {
        require ( id == _work.id );

        set(_work);
    }

    function get() override virtual public view returns (CreativeWorks memory) {

        CreativeWorks memory works;

        works.workType = work.workType;
        works.license = work.license;
        works.id = id;
        works.dateCreated  = work.dateCreated;
        works.dateModified = work.dateModified;
        works.url = work.url;
        works.name = work.name;
        works.description = work.description;

        EntityNode entity = EntityNode(entities.getEntityContract(work.author));
        require ( entity.isType(EntityTypes.AUTHOR) );
        works.author = entity.get();

        if ( work.copyrightHolder[0] != 0 ) {
            entity = EntityNode(entities.getEntityContract(work.copyrightHolder));
            require ( entity.isType(EntityTypes.COPYRIGHTHOLDER) );
            works.copyrightHolder = entity.get();
        } else {
            works.copyrightHolder.id = bytes32(0);
            works.copyrightHolder.name = "";
            works.copyrightHolder.email = "";
            works.copyrightHolder.url = "";
        }

        if ( work.publisher[0] != 0 ) {
            entity = EntityNode(entities.getEntityContract(work.publisher));
            require ( entity.isType(EntityTypes.PUBLISHER) );
            works.publisher = entity.get();
        } else {
            works.publisher.id = bytes32(0);
            works.publisher.name = "";
            works.publisher.email = "";
            works.publisher.url = "";
        }

        return works;
    }

    function set(CreativeWorks memory _work) private {
        require (
            _work.workType > WorksTypes.NONE &&
            _work.workType < WorksTypes.MAX &&
            _work.license > LicenseTypes.NONE &&
            _work.license < LicenseTypes.MAX &&
            _work.dateCreated[0] != 0 &&
             bytes(_work.name).length > 0 &&
             bytes(_work.description).length > 0 &&
             _work.author.id[0] != 0
        );

        work.workType = _work.workType;
        work.license = _work.license;
        work.dateCreated = _work.dateCreated;
        work.dateModified = _work.dateModified;
        work.author = _work.author.id;
        work.copyrightHolder = _work.copyrightHolder.id;
        work.publisher = _work.publisher.id;
        work.url = _work.url;
        work.name = _work.name;
        work.description = _work.description;

        entities.addEntity(_work.author, EntityTypes.AUTHOR);

        if ( work.copyrightHolder[0] != 0 ) {
            entities.addEntity(_work.copyrightHolder, EntityTypes.COPYRIGHTHOLDER);
            entities.addEntityRelation(work.author, work.copyrightHolder);
            entities.addEntityRelation(work.copyrightHolder, work.author);
        }

        if ( work.publisher[0] != 0 ) {
            entities.addEntity(_work.publisher, EntityTypes.PUBLISHER);
            entities.addEntityRelation(work.author, work.publisher);
            entities.addEntityRelation(work.publisher, work.author);
        }
    }
}

contract Artefacts is IArtefactFactory, IFactory {

    IEntitiesFactory entities;
    Data workStore;
    using IterableData for Data;

    constructor(address _entityFactory) public {
        require (
            _entityFactory != address(0x0)
        );

        entities = IEntitiesFactory(_entityFactory);
        workStore.size = 0;
    }

    function addWork(CreativeWorks memory _work) override virtual public {

        if (workStore.contains(_work.id)) {

            amendWork(_work);

        } else {

            ArtefactNode work = new ArtefactNode(_work, address(entities));
            workStore.insert(_work.id, address(work));
        }
    }

    function amendWork(CreativeWorks memory _work) override virtual public {
        require ( workStore.data[_work.id].value != address(0x0) );

        ArtefactNode work = ArtefactNode(workStore.data[_work.id].value);
        work.amend(_work);
    }

    function getWork(bytes32 _id) override virtual public view returns (CreativeWorks memory) {
        require ( workStore.data[_id].value != address(0x0)
        );

        ArtefactNode work = ArtefactNode(workStore.data[_id].value);
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
