pragma solidity >=0.4.21 <0.8.0;
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
    bytes32[]           authors;
    bytes32[]           copyrightHolders;
    bytes32[]           publishers;

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

    function addAuthor(CreativeEntities memory _author) override virtual public {

      entities.addEntity(_author, EntityTypes.AUTHOR);

      if ( !isAuthor(_author.id) ) {
        authors.push(_author.id);
      }
    }

    function addCopyrightHolder(CreativeEntities memory _copyrightHolder) override virtual public {

      entities.addEntity(_copyrightHolder, EntityTypes.COPYRIGHTHOLDER);
      entities.addEntityRelation(id, _copyrightHolder.id);
      entities.addEntityRelation(_copyrightHolder.id, id);

      if ( !isCopyrightHolder(_copyrightHolder.id) ) {
        copyrightHolders.push(_copyrightHolder.id);
      }
    }

    function addPublisher(CreativeEntities memory _publisher) override virtual public {

      entities.addEntity(_publisher, EntityTypes.PUBLISHER);
      entities.addEntityRelation(id, _publisher.id);
      entities.addEntityRelation(_publisher.id, id);

      if ( !isPublisher(_publisher.id) ) {
        publishers.push(_publisher.id);
      }
    }

    function get() override virtual public view returns (Works memory) {

        return work;
    }

    function getAuthors() override virtual public view returns (bytes32[] memory) {

      return authors;
    }

    function getCopyrightHolders() override virtual public view returns (bytes32[] memory) {

      return copyrightHolders;
    }

    function getPublishers() override virtual public view returns (bytes32[] memory) {

      return publishers;
    }

    function isAuthor(bytes32 _id) override virtual public view returns (bool) {
        require ( _id[0] != 0 );

        bool found = false;
        for (uint i = 0; i < authors.length; i++) {
           if ( authors[i] == _id ) {
               found = true;
               break;
           }
        }

        return found;
    }

    function isCopyrightHolder(bytes32 _id) override virtual public view returns (bool) {
        require ( _id[0] != 0 );

        bool found = false;
        for (uint i = 0; i < copyrightHolders.length; i++) {
           if ( copyrightHolders[i] == _id ) {
               found = true;
               break;
           }
        }

        return found;
    }

    function isPublisher(bytes32 _id) override virtual public view returns (bool) {
        require ( _id[0] != 0 );

        bool found = false;
        for (uint i = 0; i < publishers.length; i++) {
           if ( publishers[i] == _id ) {
               found = true;
               break;
           }
        }

        return found;
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
        work.url = _work.url;
        work.name = _work.name;
        work.description = _work.description;

        addAuthor(_work.author);
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

    function addWorkAuthor(bytes32 _workId, CreativeEntities memory _author) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.addAuthor(_author);
    }

    function addWorkCopyrightHolder(bytes32 _workId, CreativeEntities memory _copyrightHolder) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.addCopyrightHolder(_copyrightHolder);
    }

    function addWorkPublisher(bytes32 _workId, CreativeEntities memory _publisher) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.addPublisher(_publisher);
    }

    function getWork(bytes32 _id) override virtual public view returns (Works memory) {
        require ( workStore.data[_id].value != address(0x0) );

        ArtefactNode work = ArtefactNode(workStore.data[_id].value);
        return work.get();
    }

    function getWorkAuthors(bytes32 _id) override virtual public view returns (bytes32[] memory) {
      require ( workStore.data[_id].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_id].value);
      return work.getAuthors();
    }

    function getWorkCopyrightHolders(bytes32 _id) override virtual public view returns (bytes32[] memory) {
      require ( workStore.data[_id].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_id].value);
      return work.getCopyrightHolders();
    }

    function getWorkPublishers(bytes32 _id) override virtual public view returns (bytes32[] memory) {
      require ( workStore.data[_id].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_id].value);
      return work.getPublishers();
    }

    function getWorkContract(bytes32 _id) override virtual public view returns (address) {
        require( workStore.data[_id].value != address(0x0) );

        return workStore.data[_id].value;
    }

    function isWorkAuthor(bytes32 _workId, bytes32 _authorId) override virtual public view returns (bool) {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      return work.isAuthor(_authorId);
    }

    function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) override virtual public view returns (bool) {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      return work.isCopyrightHolder(_copyrightHolderId);
    }

    function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) override virtual public view returns (bool) {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      return work.isPublisher(_publisherId);
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

    function getReferences() override virtual public view returns (bytes32[] memory) {

        bytes32[] memory refs = new bytes32[](workStore.size);
        for (uint i = 0; i < workStore.size; i++) {
            refs[i] = workStore.keys[i].key;
        }
        return refs;
    }
}
