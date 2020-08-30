pragma solidity >=0.4.21 <0.8.0;
pragma experimental ABIEncoderV2;

import "./Enums.sol";

// Input/Output
struct CreativeEntities {
    bytes32             id;
    string              name;
    string              email;
    string              url;
}


// What we receive
struct CreativeWorks {
    WorksTypes          workType;
    LicenseTypes        license;
    bytes32             id;
    bytes32             dateCreated;
    bytes32             dateModified;
    string              url;
    string              name;
    string              description;
    CreativeEntities    author;
}

// what we store and output
struct Works {
    WorksTypes          workType;
    LicenseTypes        license;
    bytes32             dateCreated;
    bytes32             dateModified;
    string              url;
    string              name;
    string              description;
}
