// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract Reverse {

    function reverse(string memory str) external pure returns(string memory) {
        // 先转为字节数组
        bytes memory _byte = bytes(str);
        // 返回 不在原字节数组进行反转
        bytes memory arr = new bytes(_byte.length);
        for(uint i=0; i < _byte.length; i++) {
            arr[i] = _byte[_byte.length - 1 - i];
        }
        return string(arr);
        // for(uint i=0; i< _byte.length/2; i++) {
        //     bytes1 temp = _byte[i];
        //     _byte[i] = _byte[_byte.length - 1 -i];
        //     _byte[_byte.length - 1 -i] = temp;
        // }
        // return string(_byte);
    }


    function reverseGas(string calldata str) external pure returns(string memory) {
        bytes memory _bytes = bytes(str);

        // 原地反转
        uint length = _bytes.length;
        for(uint i=0; i<length/2; i++) {
            // 元组交换
            (_bytes[i], _bytes[length - 1 - i]) = (_bytes[length - 1 - i], _bytes[i]);
        }
        return string(_bytes);
    }
}