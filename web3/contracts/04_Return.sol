// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract Return {

    // 函数输出，包括：返回多种变量，命名式返回，以及利用解构式赋值读取全部或部分返回值。
    // 返回值： return 和returns
    // Solidity 中与函数输出相关的有两个关键字：return 和 returns 它们的区别在于：
    // returns： 跟在函数名后面，用于声明返回的变量类型及变量名。
    // return： 用于函数主题中，返回指定的变量。

    // 返回多个变量     数组类型返回值默认必须用memory修饰
    function returnMultiple() public pure returns(uint256, bool, uint256[3] memory) {
        return(1, true, [uint256(1),2,5]);
    }

    // 命名式返回
    function returnNamed() public pure returns(uint256 _number, bool _bool, uint256[3] memory _array) {
        _number = 2;
        _bool = false;
        _array = [uint256(3),2,1];
    }

    function returnNamed2() public pure returns(uint256 _number, bool _bool, uint256[3] memory _array) {
        return(1, true, [uint256(1),2,5]);
    }

    // 结构式赋值
    function readReturn() public pure {
        // 读取所有返回值
        uint256 _number;
        bool _bool;
        bool _bool2;
        uint256[3] memory _array;
        (_number, _bool, _array) = returnNamed();

        // 读取部分返回值
        (, _bool2, ) = returnNamed();
    }
}