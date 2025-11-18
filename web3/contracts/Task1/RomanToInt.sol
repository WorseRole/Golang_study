// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract RomanToInt{

/*
    mapping (bytes1 => uint256) private symbolValues ;

    constructor() {
        // 在构造函数中初始化映射
        symbolValues['I'] = 1;
        symbolValues['V'] = 5;
        symbolValues['X'] = 10;
        symbolValues['L'] = 50;
        symbolValues['C'] = 100;
        symbolValues['D'] = 500;
        symbolValues['M'] = 1000;
    }

    function romanToInt(string calldata roman) external view returns(uint256) {
        bytes memory romanBytes = bytes(roman);
        uint256 result = 0;
        uint256 n = romanBytes.length;
        for(uint256 i = 0; i < n; i++) {
            uint256 value = symbolValues[romanBytes[i]];
            // 检查是否是减法
            if(i<n-1 && value < symbolValues[romanBytes[i+1]]) {
                result -= value;
            } else {
                result += value;
            }
        }
        return result;
    }
*/

    function romanToInt(string calldata roman) external pure returns (uint256) {
        bytes memory romanBytes = bytes(roman);
        uint256 result = 0;
        uint256 n = romanBytes.length;
        uint256 prevValue = 0;

        // 从右向左遍历 因为solidity i 不能为负数 会报错 所以需要大一位 不能用 for (uint256 i = n-1; i >= 0; i--) {
        for (uint256 i = n; i > 0; i--) {
            uint256 currentValue = getValue(romanBytes[i - 1]);
            if(currentValue < prevValue) {
                result -= currentValue;
            } else {
                result += currentValue;
            }
            prevValue = currentValue;
        }
        return result;

    } 

    function getValue(bytes1 romanChar) internal pure returns (uint256) {
        if (romanChar == 'I') return 1;
        if (romanChar == 'V') return 5;
        if (romanChar == 'X') return 10;
        if (romanChar == 'L') return 50;
        if (romanChar == 'C') return 100;
        if (romanChar == 'D') return 500;
        if (romanChar == 'M') return 1000;
        revert("Invalid Roman character.");
    }
    

}