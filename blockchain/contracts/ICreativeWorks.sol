pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Types.sol";

abstract contract IEntity {

    function get() virtual public view returns (Entities memory);
    function getType() virtual public view returns (EntityTypes);
    function amend(Entities memory _entity) virtual public;
    function set(Entities memory _entity) virtual private;
}

abstract contract IEntitiesFactory {

    function addEntity(Entities memory _entity, EntityTypes _type) virtual public;
    function amendEntity(Entities memory _entity) virtual public;
    function getEntity(bytes32 _id) virtual public view returns (Entities memory);
    function getEntityContract(bytes32 _id) virtual public view returns (address);
}

abstract contract IWorks {

    function get() virtual public view returns (CreativeWorks memory);
    function amend(CreativeWorks memory _work) virtual public;
    function set(CreativeWorks memory _work) virtual private;
}

abstract contract IWorksFactory {

    function addWork(CreativeWorks memory _work) virtual public;
    function amendWork(CreativeWorks memory _work) virtual public;
    function getWork(bytes32 _id) virtual public view returns (CreativeWorks memory)
    function getWorkContract(bytes32 _id) virtual public view returns (address);
}
