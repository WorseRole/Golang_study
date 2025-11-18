// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// Solidity有三种方法向其他合约发送ETH，他们是：transfer()，send()和call()，其中call()是被鼓励的用法。

error SendFailed(); // 用send发送ETH失败error
error CallFailed(); // 用call发送ETH失败error

contract SendETH {

    // 构造函数，payable使得部属的时候可以转eth进去
    constructor() payable{}

    // receive 方法，接收eth 时被触发
    receive() external payable { }

    // 用transfer()发送ETH
    function transferETH(address payable _to, uint256 amount) external payable{
        _to.transfer(amount);
    }

    // send()发送ETH
    function sendETH(address payable _to, uint256 amount) external payable{
        // 处理下send的返回值，如果失败，revert交易并发送error
        bool success = _to.send(amount);
        if(!success){
            revert SendFailed();
        }
    }

    // call()发送ETH
    function callETH(address payable _to, uint256 amount) external payable{
        // 处理下call的返回值，如果失败，revert交易并发送error
        (bool success,) = _to.call{value: amount}("");
        if(!success){
            revert CallFailed();
        }
    }
}
/*
call没有gas限制，最为灵活，是最提倡的方法；
transfer有2300 gas限制，但是发送失败会自动revert交易，是次优选择；
send有2300 gas限制，而且发送失败不会自动revert交易，几乎没有人用它。
*/

contract ReceiveETH {
    // 接收合约
    // 收到eth 事件，记录amount 和 gas
    event Log (uint amount, uint gas);

    // receive 方法，接收eth 时被触发
    receive() external payable { 
        emit Log(msg.value, gasleft());
    }
    // 返回合约ETH 余额
    function getBalance() view public returns(uint) {
        return address(this).balance;
    }
}