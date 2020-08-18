pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Enums.sol";

struct Person {
    address             id;
    string              url;
    string              name;
    string              email;
}

struct CreativeWorks {
    bytes32             type;
    bytes32             id;
    bytes32             dateCreated;
    bytes32             dateModified;
    string              url;
    string              name;
    string              email;
    string              description;
    string              headline;
    Person              copyrightHolder;
    Person              author;
    Person              publisher;
}
