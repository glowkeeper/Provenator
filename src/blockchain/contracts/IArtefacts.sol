pragma solidity >=0.4.21 <0.8.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IArtefact {

    function amend(CreativeWorks memory _work) virtual public;

    function addAuthor(bytes32 _authorId) virtual public;
    function addCopyrightHolder(bytes32 _copyrightHolderId) virtual public;
    function addPublisher(bytes32 _publisherId) virtual public;

    function removeAuthor(bytes32 _id) virtual public;
    function removeCopyrightHolder(bytes32 _id) virtual public;
    function removePublisher(bytes32 _id) virtual public;

    function get() virtual public view returns (CreativeWorks memory);
    function getAuthors() virtual public view returns (bytes32[] memory);
    function getCopyrightHolders() virtual public view returns (bytes32[] memory);
    function getPublishers() virtual public view returns (bytes32[] memory);

    function isAuthor(bytes32 _id) virtual public view returns (bool);
    function isCopyrightHolder(bytes32 _id) virtual public view returns (bool);
    function isPublisher(bytes32 _id) virtual public view returns (bool);
}

abstract contract IArtefactFactory {

    function addWork(CreativeWorks memory _work) virtual public;
    function amendWork(CreativeWorks memory _work) virtual public;

    function addWorkAuthor(bytes32 _workId, bytes32 _authorId) virtual public;
    function addWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) virtual public;
    function addWorkPublisher(bytes32 _workId, bytes32 _publisherId) virtual public;

    function removeWorkAuthor(bytes32 _workId, bytes32 _authorId) virtual public;
    function removeWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) virtual public;
    function removeWorkPublisher(bytes32 _workId, bytes32 _publisherId) virtual public;

    function getWork(bytes32 _id) virtual public view returns (CreativeWorks memory);
    function getWorkAuthors(bytes32 _id) virtual public view returns (bytes32[] memory);
    function getWorkCopyrightHolders(bytes32 _id) virtual public view returns (bytes32[] memory);
    function getWorkPublishers(bytes32 _id) virtual public view returns (bytes32[] memory);
    function getWorkContract(bytes32 _id) virtual public view returns (address);

    function isWorkAuthor(bytes32 _workId, bytes32 _authorId) virtual public view returns (bool);
    function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) virtual public view returns (bool);
    function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) virtual public view returns (bool);
}
