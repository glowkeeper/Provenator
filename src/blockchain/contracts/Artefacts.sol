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

    CreativeWorks       work;
    bytes32[]           authors;
    bytes32[]           copyrightHolders;
    bytes32[]           publishers;

    constructor (CreativeWorks memory _work) public {
        require (
            _work.id[0] != 0 &&
            _work.authorId[0] != 0 &&
            _work.workType > WorksTypes.NONE &&
            _work.workType < WorksTypes.MAX &&
            _work.license < LicenseTypes.MAX &&
            _work.dateCreated[0] != 0 &&
             bytes(_work.name).length > 0 &&
             bytes(_work.description).length > 0
        );

        work = _work;
    }

    function amend(CreativeWorks memory _work) override virtual public {
        require (
            work.id == _work.id &&
            _work.authorId[0] != 0 &&
            _work.workType > WorksTypes.NONE &&
            _work.workType < WorksTypes.MAX &&
            _work.license < LicenseTypes.MAX &&
            _work.dateCreated[0] != 0 &&
             bytes(_work.name).length > 0 &&
             bytes(_work.description).length > 0
        );

        work = _work;
    }

    function addAuthor(bytes32 _authorId) override virtual public {

      if ( !isAuthor(_authorId) ) {
        authors.push(_authorId);
      }
    }

    function addCopyrightHolder(bytes32 _copyrightHolderId) override virtual public {

      if ( !isCopyrightHolder(_copyrightHolderId) ) {
        copyrightHolders.push(_copyrightHolderId);
      }
    }

    function addPublisher(bytes32 _publisherId) override virtual public {

      if ( !isPublisher(_publisherId) ) {
        publishers.push(_publisherId);
      }
    }

    function removeAuthor(bytes32 _id) override virtual public {
        require ( _id[0] != 0 );

        for (uint i = 0; i < authors.length; i++) {
          if ( authors[i] == _id ) {
            if ( i < authors.length -1 ) {
              for (uint j = i; j < authors.length - 1; j++) {
                authors[j] = authors[j+1];
              }
            }
            authors.pop();
            break;
          }
        }
    }

    function removeCopyrightHolder(bytes32 _id) override virtual public {
        require ( _id[0] != 0 );

        for (uint i = 0; i < copyrightHolders.length; i++) {
          if ( copyrightHolders[i] == _id ) {
            if ( i < copyrightHolders.length -1 ) {
              for (uint j = i; j < copyrightHolders.length - 1; j++) {
                copyrightHolders[j] = copyrightHolders[j+1];
              }
            }
            copyrightHolders.pop();
            break;
          }
        }
    }

    function removePublisher(bytes32 _id) override virtual public  {
        require ( _id[0] != 0 );

        for (uint i = 0; i < publishers.length; i++) {
          if ( publishers[i] == _id ) {
            if ( i < publishers.length -1 ) {
              for (uint j = i; j < publishers.length - 1; j++) {
                publishers[j] = publishers[j+1];
              }
            }
            publishers.pop();
            break;
          }
        }
    }

    function get() override virtual public view returns (CreativeWorks memory) {

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
}

contract Artefacts is IArtefactFactory, IFactory {

    Data    workStore;
    using IterableData for Data;

    constructor() public {

        workStore.size = 0;
    }

    function addWork(CreativeWorks memory _work) override virtual public {

        if (workStore.contains(_work.id)) {

            amendWork(_work);

        } else {

            ArtefactNode work = new ArtefactNode(_work);
            workStore.insert(_work.id, address(work));
        }
    }

    function amendWork(CreativeWorks memory _work) override virtual public {
        require ( workStore.data[_work.id].value != address(0x0) );

        ArtefactNode work = ArtefactNode(workStore.data[_work.id].value);
        work.amend(_work);
    }

    function addWorkAuthor(bytes32 _workId, bytes32 _authorId) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.addAuthor(_authorId);
    }

    function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.addCopyrightHolder(_copyrightHolderId);
    }

    function addWorkPublisher(bytes32 _workId, bytes32 _publisherId) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.addPublisher(_publisherId);
    }

    function removeWorkAuthor(bytes32 _workId, bytes32 _authorId) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.removeAuthor(_authorId);
    }

    function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.removeCopyrightHolder(_copyrightHolderId);
    }

    function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId) override virtual public {
      require( workStore.data[_workId].value != address(0x0) );

      ArtefactNode work = ArtefactNode(workStore.data[_workId].value);
      work.removeCopyrightHolder(_publisherId);
    }

    function getWork(bytes32 _id) override virtual public view returns (CreativeWorks memory) {
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
