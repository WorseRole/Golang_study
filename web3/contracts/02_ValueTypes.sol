// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// 值类型
// 变量类型包括
// 1. 值类型     布尔型、整数型等等，这类变量赋值时候直接传递数值
// 2. 引用类型   数组和结构体，这类变量占空间大，赋值时候直接传递地址（类似指针）
// 3. 映射类型  Solidity中存储键值对的数据结构，可以理解为哈希表
contract ValueTypes{

    // 布尔
    bool public  _bool = true;

    bool public _bool1 = !_bool;    // 取非 false
    bool public _bool2 = _bool && _bool1;   // 与   false
    bool public _bool3 = _bool || _bool1;   // 或   true
    bool public _bool4 = _bool == _bool1;   // 相等 false
    bool public _bool5 = _bool != _bool1;   // 不相等   true


    // 整型
    int public _int = -1;   // 整数，包括负数
    uint public _uint = 1; // 无符号整数
    uint256 public _number = 20251116;  // 256位无符号整数

    // 常用的整型运算
    // 比较运算符(返回布尔值)： <=, <, ==, !=, >=, >
    // 算术运算符： +, -, *, /, %, **(幂)
    uint256 public _number1 = _number + 1;
    uint256 public _number2 = 2**2;
    uint256 public _number3 = 7 % 2;
    bool public _numberbool = _number2 > _number3; 


    // 地址类型
    // 普通地址：存储一个20字节的值（以太坊地址的大小）
    // payable address： 比普通地址多了transfer 和send 两个成员方法，用于接收转账
    address public _address = 0x7A58c0Be72BE218B41C608b7Fe7C5bB630736C71;
    address payable public _address1 = payable(_address);   // payable address，可以转账、查余额
    // 地址类型的成员
    uint256 public balance = _address1.balance; // balance of address


    // 字节数组
    // 字节数组分为 定长 和 不定长 字节数组
    // 定长字节数组： 属于值类型，数组长度在声明之后不能改变。根据字节数组的长度分为 bytes1, bytes8, bytes32等类型。定长字节数组最多存储32 bytes数据，即bytes32
    // 不定长字节数组：属于引用类型（之后的章节介绍），数组长度在声明之后可以改变，包括bytes等
    // 固定长度的字节数组
    bytes32 public _byte32 = "MiniSolidity";
    bytes1 public _byte = _byte32[0];



    // 枚举
    // 枚举是Solidity 中用户定义的数据类型。它主要用于为 uint 分配名称，使程序易于阅读和维护。它与 C 语言中的 enum 类似，使用名称来代替从 0 开始的 uint
    // 用enum 将 uint 0, 1, 2表示为Buy, Hold, Sell
    enum ActionSet {
        Buy, Hold, Sell
    }
    // 创建enum 变量 action
    ActionSet action = ActionSet.Buy;

    // 枚举可以显示地和 uint 相互转换，炳辉检查转换的无符号整数是否在枚举的长度内，否则会报错:
    // enum 可以和 uint 显示的转换
    function enumToUint() external view returns(uint) {
        return uint(action);
    }
    //enum 是一个比较冷门的数据类型，几乎没什么人用。



}