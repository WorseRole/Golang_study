// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract MergeSoetedArray {

    function mergeSortedArray(uint256[] calldata array1, uint256[] calldata array2) external pure returns(uint256[] memory) {
        uint256 array1Len = array1.length;
        uint256 array2Len = array2.length;
        // result
        uint256[] memory result = new uint256[](array1Len + array2Len);
        // 双指针
        uint256 i = 0;
        uint256 j = 0;
        // solidy 需要指定数组大小 数组不支持add或者push
        uint256 k = 0;
        while(i < array1Len && j < array2Len) {
            if(array1[i] >= array2[j]) {
                result[k] = array2[j];
                j++;
            } else {
                result[k] = array1[i];
                i++;
            }
            k++;
        }
        while(i < array1Len) {
            result[k] = array1[i];
            i++;
            k++;
        }
        while(j < array2Len) {
            result[k] = array2[j];
            j++;
            k++;
        }
        return result;
    }
}