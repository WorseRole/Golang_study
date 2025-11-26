// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.21;

contract BeggingContract {
    // 状态变量
    address public owner;
    mapping (address => uint256) donorsMap;

    event Donated(address indexed donor, uint256 amount);
    event Withdraw(address indexed owner, uint256 amount);

    constructor() {
        owner = msg.sender;
    }

    // 捐赠人调就不用传参数了
    function donate() public payable {
        require(msg.value > 0, "Donotation amount must be greater than 0");
        donorsMap[msg.sender] += msg.value;

        emit Donated(msg.sender, msg.value);
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    function withdraw() public onlyOwner{
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");
        
        (bool success, ) = owner.call{value: balance}("");
        require(success, "withdraw failed");

        emit Withdraw(owner, balance);
    }

    function getDonation(address account) public view returns(uint256) {
        return donorsMap[account];
    }

    function getContractBalance() public view returns(uint256) {
        return address(this).balance;
    }

    // 回退函数
    receive() external payable {
        donorsMap[msg.sender] += msg.value;
        emit Donated(msg.sender, msg.value);
    }


}