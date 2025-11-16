// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract InitValue{
    // 在Solidity 中，声明但没赋值的变量都有它的初始值或默认值。
    // 值类型初始值
    /*
    boolean: false
    string: ""
    int: 0
    uint: 0
    enum: 枚举中的第一个元素
    address: 0x0000000000000000000000000000000000000000 (或 address(0))
    function
        internal: 空白函数
        external: 空白函数
    */
    // 可以用public 变量的getter 函数验证上面写的初始值是否正确：
    bool public _bool;
    string public _string;
    int public _int;
    uint public _uint;
    address public _address;
    enum ActionSet {Buy, Hold, Sell}
    ActionSet public _enum; // 第一个内容Buy 为0
    function fi() internal {}   // internal 空白函数
    function fe() external {}   // external 空白函数



    // 引用类型初始值
    /*
    映射mapping: 所有元素都为其默认值的mapping
    结构体struct: 所有成员设为其默认值的结构体
    数组array
        动态数组: []
        静态数组（定长）: 所有成员设为其默认值的静态数组
    */
    uint[8] public _staticArray; // 所有成员设为其默认值的静态数组[0, 0, 0, 0, 0, 0, 0, 0]
    uint[] public _dynamicArray; // `[]`
    mapping (uint => address) public _mapping;  // 所有元素都为其默认值的mapping
    // 所有成员设为其默认值的结构体 0，0
    struct Student{
        uint256 id;
        uint score;
    }
    Student public student;

    // delete 操作符
    // delete a 会让变量a 的值变为初始值
    bool public _bool2 = true;
    function d() external {
        delete _bool2; // delete 会让bool2 变为默认值 false
    }
}