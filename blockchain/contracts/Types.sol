pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Enums.sol";

struct Entity {
    string              name;
    string              email;
    string              url;
}

struct CreativeWorks {
    Works               type;
    License             license;
    bytes32             id;
    bytes32             dateCreated;
    bytes32             dateModified;
    string              url;
    string              name;
    string              description;
    Entity              author;
    Entity              copyrightHolder;
    Entity              publisher;
}
