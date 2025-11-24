// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract BeggingConstract {
    // 允许用户向合约地址发送以太币。
    // 记录每个捐赠者的地址和捐赠金额。
    // 允许合约所有者提取所有捐赠的资金。
    address public owner;
    // donorsMap 只是一个记账系统 ， 真正的ETH 是存储在合约地址的余额中的 所以withdraw()函数提取的是合约地址的实际ETH余额，不是从donorsMap中提取
    mapping(address => uint256) donorsMap;

    // 事件记录捐赠和提款
    event Donated(address indexed donor, uint256 amount);
    event Withdraw(address indexed woner, uint256 amount);

    constructor() {
        owner = msg.sender;
    }

    // 捐赠函数 - 需要payable 来接收ETH
    function donate() public payable  {
        require(msg.value > 0, "Donotation amount must be greater than 0");
        // 只是记录
            // ETH自动进入合约余额，不需要手动操作！
        donorsMap[msg.sender] += msg.value;
        emit Donated(msg.sender, msg.value);
    }

    modifier onlyOwner {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // 允许合约所有者提取所有资金
    function withdraw() public onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");

        // 使用 call 而不是 transfer , 更安全
        // 从合约余额中扣除balance，向 owner 地址转入balance 数量的ETH
        // owner.call: 对owner地址进行底层调用 
        // {value: balance}: 发送指定数量的ETH 
        // "": 空数据，表示不调用任何函数
        // (bool success, ): 接收调用结果（成功/失败）
        (bool success, ) = owner.call{value: balance}("");
        require(success, "withdrawal failed");

        // 等于上面的 call 但是 不如call 安全可靠
        // payable (owner).transfer(balance);


        emit Withdraw(owner, balance);
    }

    // 允许查询某个地址的捐赠金额
    function getDonation(address account) public view returns(uint256) {
        return donorsMap[account];
    }

    // 获取合约余额
    function getContractBalance() public view returns(uint256) {
        return address(this).balance;
    }

    // 接收ETH 的回退函数
    receive() external payable {
        donorsMap[msg.sender] += msg.value;
        emit Donated(msg.sender, msg.value);
     }

}