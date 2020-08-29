pragma solidity >=0.4.21 <0.8.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IArtefact {

    function amend(CreativeWorks memory _work) virtual public;

    function addAuthor(CreativeEntities memory _author) virtual public;
    function addCopyrightHolder(CreativeEntities memory _copyrightHolder) virtual public;
    function addPublisher(CreativeEntities memory _publisher) virtual public;

    function get() virtual public view returns (Works memory);
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

    function addWorkAuthor(bytes32 _workId, CreativeEntities memory _author) virtual public;
    function addWorkCopyrightHolder(bytes32 _workId, CreativeEntities memory _copyrightHolder) virtual public;
    function addWorkPublisher(bytes32 _workId, CreativeEntities memory _publisher) virtual public;

    function getWork(bytes32 _id) virtual public view returns (Works memory);
    function getWorkAuthors(bytes32 _id) virtual public view returns (bytes32[] memory);
    function getWorkCopyrightHolders(bytes32 _id) virtual public view returns (bytes32[] memory);
    function getWorkPublishers(bytes32 _id) virtual public view returns (bytes32[] memory);
    function getWorkContract(bytes32 _id) virtual public view returns (address);

    function isWorkAuthor(bytes32 _workId, bytes32 _authorId) virtual public view returns (bool);
    function isWorkCopyrightHolder(bytes32 _workId, bytes32 _copyrightHolderId) virtual public view returns (bool);
    function isWorkPublisher(bytes32 _workId, bytes32 _publisherId) virtual public view returns (bool);
}
