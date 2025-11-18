// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract BinarySearch {

    function binarySortedSearch(uint256[] calldata array, uint256 value) external pure returns(uint256) {
        if (array.length == 0) return type(uint256).max;
        // first, mid   mid, last
        uint256 first = 0;
        uint256 mid = 0; //(array1.length - 1)/2 + 1;
        uint256 last = array.length-1;

        while(mid < last) {
            if(array[mid] == value) {
                return mid;
            }
            mid = (first + last) / 2 + 1;
            if(array[mid] > value) {
                last = mid;
                mid = (first + last) / 2 + 1;
            }else {
                if(array[mid] < value) {
                    first = mid;
                    mid = (first + last) / 2 + 1;
                }
            }
        }
        return type(uint256).max;
    }
}