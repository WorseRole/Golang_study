// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract FunctionTypes {
    // pure 和 view 关键字
    // solidity 引入这两个关键字主要是因为以太坊需要支付 gas fee。
    // 合约的状态变量存储在链上，gas fee 很贵，如果计算不改变链上状态，就可以不用付gas。
    // 包含pure 和view 关键字的函数是不改写链上状态的，因此用户直接调用它们是不需要付 gas的
    // （注意： 合约中非pure/view 函数调用 pure/view 函数时需要付gas ）

    
    // 以太坊中，以下语句被视为修改链上状态：
    // 1. 写入状态变量
    // 2. 释放事件
    // 3. 创建其他合约
    // 4. 使用 selfdestruct
    // 5. 通过调用发送以太币
    // 6. 调用任何未标记 view 或 pure 的函数
    // 7. 使用低级调用（low-level calls）
    // 8. 使用包含某些操作码的内联汇编

    // pure 函数既不能读取也不能写入链上的状态变量
    // view 函数能读取但不能写入链上的状态变量
    // 非 pure 或view 的函数既可以读取也可以写入状态变量。

    // 在合约里定义一个状态变量 number 5
    uint256 public number = 5;

    // 定义一个 add 函数
    function add() external {
        number = number + 1;
    }
    // 如果 add() 函数被标记为pure，比如 function.add() external pure，就会报错。因为pure 是不配合读取合约里的状态变量的，更不配改写。
    // 那可以给函数传递一个参数 _number,然后让他返回 _number + 1, 这个操作不会读取或写入状态变量
    // pure: 纯纯牛马
    function addPure(uint _number) external pure returns (uint256 new_number) {
        new_number = _number + 1;
    }

    // 如果add() 函数被标记为 view，比如 function add() external view, 也会报错。因为view 能读取，但不能够改写状态变量。
    // 可以改写下函数，读取但是不改写number，返回一个新的变量
    // view: 看客
    function addView() external view returns (uint256 new_number) {
        new_number = number + 1;
    }

    
    // internal v.s. external
    // 定义一个 internal 的minus() 函数，每次调用使得number 变量减少1。由于internal 函数只能由合约内部调用，
    // 必须再定义一个external 的 minusCall() 函数，通过它简洁调用内部的 minus() 函数

    // internal 内部函数
    function minus() internal {
        number = number - 1;
    }
    // 合约内的函数可以调用内部函数
    function minusCall() external {
        minus();
    }



    // 定义一个 external payable的 minusPayable() 函数，间接地调用minus(), 
    //  并且返回合约里的 ETH 余额（this 关键字可以让我们引用合约地址）。可以在调用minusPayable()时往合约里转入1 个ETH
    // payable
    // payable: 递钱，能给合约支付eth 的函数
    function minusPayable() external payable returns (uint256 balance) {
        minus();
        balance = address(this).balance;
    }


}