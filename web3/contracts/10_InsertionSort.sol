// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract InsertionSort {
    // if else 
    function ifElseTest(uint256 _number) public pure returns(bool) {
        if(_number == 0) {
            return true;
        } else {
            return false;
        }
    }
    // for loop
    function forLoopTest() public pure returns(uint256) {
        uint sum = 0;
        for(uint i=0; i<10; i++) {
            sum += i;
        }
        return sum;
    }

    function whileTest() public pure returns(uint256) {
        uint sum = 0;
        uint i=0;
        while(i<10) {
            sum+=i;
            i++;
        }
        return sum;
    }

    function doWhileTest() public pure returns(uint256) {
        uint sum = 0;
        uint i = 0;
        do {
            sum += i;
            i++;
        } 
        while (i<10);
        return sum;
    }

    // 三元运算符
    function ternaryTest(uint256 x, uint y) public pure returns(uint256) {
        // return the max of x and y
        return x >= y ? x : y;
    }


    function insertionSort(uint[] memory a) public pure returns(uint[] memory) {
        for(uint i = 1; i < a.length; i++) {
            uint temp = a[i];
            uint j = i;
            while( (j>=1) && (temp<a[j-1])) {
                a[j] = a[j-1];
                j--;
            }
            a[j] = temp;
        }
        return a;
    }
    function insertionSort1(uint[] memory a) public pure returns(uint[] memory) {
        for(uint i = 1; i < a.length; i++) {
            uint temp = a[i];
            uint j = i;
            for(j = i; j >= 1; j--) {
                if(temp < a[j-1]) {
                    a[j] = a[j-1];
                    continue ;
                }
            }
            // Solidity中最常用的变量类型是uint，也就是无符号整数，取到负值的话，会报underflow错误。
            // 而在插入算法中，变量j有可能会取到-1，引起报错。 <在找到最头部的话 j就是0了>

            if(j == 0 ) {
                a[j] = temp;
            } else {
                a[j-1] = temp;
            }
        }
        return a;
    }
}