pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Enums.sol";

// Input/Output
struct Entities {
    bytes32             id;
    string              name;
    string              email;
    string              url;
}

struct CreativeWorks {
    WorksTypes          workType;
    LicenseTypes        license;
    bytes32             id;
    bytes32             dateCreated;
    bytes32             dateModified;
    string              url;
    string              name;
    string              description;
    Entities            author;
    Entities            copyrightHolder;
    Entities            publisher;
}

// what we store
struct Entity {
    EntityTypes         entityType;
    string              name;
    string              email;
    string              url;
}

struct Works {
    WorksTypes          workType;
    LicenseTypes        license;
    bytes32             dateCreated;
    bytes32             dateModified;
    bytes32             author;
    bytes32             copyrightHolder;
    bytes32             publisher;
    string              url;
    string              name;
    string              description;
}
