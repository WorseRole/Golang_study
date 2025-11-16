// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DataStorage {
    // 引用类型：数组array 和 结构体 struct 由于这类变量比较复杂，占用存储空间大，我们在使用时必须要声明数据存储的位置

    // 数据存储位置有三类： storage，memory 和calldata。
    // 不同存储位置的gas 成本不同。
    // storage 类型的数据存在链上，类似计算机的硬盘，消耗gas 多；
    // memory 和 calldata类型的临时存在内存里，消耗gas 少。整体消耗gas 从多到少依次为：storage > memory > calldata
    // 1. storage: 合约里的状态变量默认都是storage，存储在链上。
    // 2. memory： 函数里的参数和临时变量一般用memory，存储在内存中，不上链/尤其是如果返回数据类型是变长的情况下，必须加memory 修饰，例如：string，bytes，array和自定义结构
    // 3. calldata：和memory 类似，存储在内存中，不上链。与memory 的不同点在于calldata 变量不能修改（immutable），一般用于函数的参数。

    function fCalldata(uint[] calldata _x) public pure returns (uint[] calldata) {
        // 参数为calldata 数组，不能被修改
        //  _x[0] = 0 // 这样会报错
         return(_x);
    }

    // 数据位置和赋值规则
    // 在不同存储类型相互赋值时候，有时会产生独立的副本（修改新变量不会影响原变量），有时会产生引用（修改新变量会影响原变量）。规则如下：
    // 赋值本质上是创建引用指向本地，因此修改本体或者是引用，变化可以被同步：
        // storage（合约的状态变量）赋值给本地storage（函数里的）时候，会创建引用，改变新变量会影响原变量。例子：
    uint[] x = [1, 2, 3];
    function fStorage() public {
        // s声明一个storage的变量 xStorage，指向x。 修改xStorage也会影响x
        uint[] storage xStorage = x;
        xStorage[0] = 100;
    }

        // memory 赋值非 memory，会创建引用，改变新变量会影响原变量
    // 其他情况下，赋值创建的是本体的副本，即对二者之一的修改，并不会同步到另一方。这有时会设计到开发中的问题，
    // 比如从storage中读取数据，赋值给memory，然后修改mermory的数据，但如果没有将memory的数据赋值回storage，
    // 那么storage的数据是不会改变的。
    function fMemory() public view  {
        uint[] memory xMemory = x;
        xMemory[0] = 100;
    }

}

// 变量的作用域
// Solidity中变量按作用域划分有三种，分别是状态变量（state variable），局部变量（local variable）和全局变量(global variable)
contract Variables {

    uint public x = 1;
    uint public y;
    string public z;

    // 1. 状态变量
    // 状态变量是数据存储在链上的变量，所有合约内函数都可以访问，gas消耗高。状态变量在合约内、函数外声明：
    function foo() external {
        // 可以在函数里更改状态变量的值
        x = 5;
        y = 2;
        z = "0xAA";
    }

    // 2. 局部变量
    // 局部变量是尽在函数执行过程中有效的变量，函数退出后，变量无效。局部变量的数据存储在内存里，不上链，gas低。
    // 局部变量在函数内声明：
    function bar() external pure returns(uint) {
        uint xx = 1;
        uint yy = 3;
        uint zz = xx + yy;
        return(zz);
    }

    // 3. 全局变量
    // 全局变量是全局范围工作的变量，都是solidity 预留关键字。它们可以在函数内不声明直接使用：
    function gloab() external view returns(address, uint, bytes memory) {
        address sender = msg.sender;    // 请求发起地址
        uint blockNum = block.number;   // 当前区块高度
        bytes memory data = msg.data;   // 请求数据
        return(sender, blockNum, data);
    }
    // 下面是一些常用的全局变量
    // blockhash(uint blockNumber): (bytes32) 给定区块的哈希值 – 只适用于最近的256个区块, 不包含当前区块。
    // block.coinbase: (address payable) 当前区块矿工的地址
    // block.gaslimit: (uint) 当前区块的gaslimit
    // block.number: (uint) 当前区块的number
    // block.timestamp: (uint) 当前区块的时间戳，为unix纪元以来的秒
    // gasleft(): (uint256) 剩余 gas
    // msg.data: (bytes calldata) 完整call data
    // msg.sender: (address payable) 消息发送者 (当前 caller)
    // msg.sig: (bytes4) calldata的前四个字节 (function identifier)
    // msg.value: (uint) 当前交易发送的 wei 值
    // block.blobbasefee: (uint) 当前区块的blob基础费用。这是Cancun升级新增的全局变量。
    // blobhash(uint index): (bytes32) 返回跟当前交易关联的第 index 个blob的版本化哈希（第一个字节为版本号，当前为0x01，后面接KZG承诺的SHA256哈希的最后31个字节）。若当前交易不包含blob，则返回空字节。这是Cancun升级新增的全局变量。


    // 4. 全局变量 - 以太单位与时间单位
    // 以太单位
    // Solidity中不存在小数点，以0代替为小数点，来确保交易的精确度，并且防止精度的损失，利用以太单位可以避免误算的问题，方便程序员在合约中处理货币交易。
    // wei:1
    // gwei:1e9 = 1000000000
    // ether: 1e18 = 1000000000000000000
    function weiUint() external pure returns(uint) {
        assert(1 wei == 1e0);
        assert(1 wei == 1);
        return 1 wei;
    }
    function gweiUint() external pure returns(uint) {
        assert(1 gwei == 1e9);
        assert(1 gwei == 1000000000);
        return 1 gwei;
    }
    function etherUnit() external pure returns(uint) {
        assert(1 ether == 1e18);
        assert(1 ether == 1000000000000000000);
        return 1 ether;
    }

    // 时间单位
    // 可以在合约中规定一个操作必须在一周内完成，或者某个时间在一个月后发生。这样就能让合约的执行可以更加精确，不会因为技术上的误差而影响合约的结果。
    // 因此，时间单位在 Solidity 中是一个重要的概念，有助于提高合约的可读性和可可维护性。
    // seconds: 1
    // minutes: 60 seconds = 60
    // hours: 60 minutes = 3600
    // days: 60 hours = 86400
    // weeks: 7 days = 604800
    function secondsUnit() external pure returns(uint) {
        assert(1 seconds == 1);
        return 1 seconds;
    }

    function minutesUnit() external pure returns(uint) {
        assert(1 minutes == 60);
        assert(1 minutes == 60 seconds);
        return 1 minutes;
    }
    function hoursUnit() external pure returns(uint) {
        assert(1 hours == 3600);
        assert(1 hours == 60 minutes);
        return 1 hours;
    }
    function daysUnit() external pure returns(uint) {
        assert(1 days == 86400);
        assert(1 days == 24 hours);
        return 1 days;
    }
    function weeksUnit() external  pure returns(uint) {
        assert(1 weeks == 604800);
        assert(1 weeks == 7 days);
        return 1 weeks;
    }

}