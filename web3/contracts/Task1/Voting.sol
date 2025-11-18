// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// ✅ 创建一个名为Voting的合约，包含以下功能：
// 一个mapping来存储候选人的得票数
// 一个vote函数，允许用户投票给某个候选人
// 一个getVotes函数，返回某个候选人的得票数
// 一个resetVotes函数，重置所有候选人的得票数
contract Voting {

    mapping (address => uint) public votes;
    address[] public keys;

    function vote(address adr) public {
        if(votes[adr] == 0) {
            keys.push(adr); 
        }
        votes[adr] += 1;
    }

    function getVotes(address adr) view  external  returns(uint) {
        return votes[adr];
    }

    function resetVotes() external {
        for(uint i =0; i<keys.length; i++) {
            votes[keys[i]] = 0;
        }
    }
}