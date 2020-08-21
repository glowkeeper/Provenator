pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Enums.sol";
import "./Types.sol";

struct IndexedValue {
    uint keyIndex;
    address value;
}

struct KeyFlag {
    bytes32 key;
    bool deleted;
}

struct Data {
    mapping(bytes32 => IndexedValue) data;
    KeyFlag[] keys;
    uint size;
}
