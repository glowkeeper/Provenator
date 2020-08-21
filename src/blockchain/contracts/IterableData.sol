pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Storage.sol";

library IterableData {

    function insert(Data storage self, bytes32 key, address value) internal returns (bool replaced) {
        uint keyIndex = self.data[key].keyIndex;
        self.data[key].value = value;
        if (keyIndex > 0)
            return true;
        else {
            keyIndex = self.keys.length;
            self.keys.push();
            self.data[key].keyIndex = keyIndex + 1;
            self.keys[keyIndex].key = key;
            self.size++;
            return false;
        }
    }

    function remove(Data storage self, bytes32 key) internal returns (bool success) {
        uint keyIndex = self.data[key].keyIndex;
        if (keyIndex == 0)
            return false;
        delete self.data[key];
        self.keys[keyIndex - 1].deleted = true;
        self.size --;
    }

    function contains(Data storage self, bytes32 key) internal view returns (bool) {
        return self.data[key].keyIndex > 0;
    }

    function iterate_start(Data storage self) internal view returns (uint keyIndex) {
        return iterate_next(self, uint(-1));
    }

    function iterate_valid(Data storage self, uint keyIndex) internal view returns (bool) {
        return keyIndex < self.keys.length;
    }

    function iterate_next(Data storage self, uint keyIndex) internal view returns (uint r_keyIndex) {
        keyIndex++;
        while (keyIndex < self.keys.length && self.keys[keyIndex].deleted)
            keyIndex++;
        return keyIndex;
    }

    function iterate_get(Data storage self, uint keyIndex) internal view returns (bytes32 key, address value) {
        key = self.keys[keyIndex].key;
        value = self.data[key].value;
    }
}
