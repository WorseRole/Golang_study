// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// 在合约中创建新合约
// 在以太坊链上，用户（外部账户，EOA）可以创建智能合约，智能合约同样也可以创建新的智能合约。
// 去中心化交易所uniswap就是利用工厂合约（PairFactory）创建了无数个币对合约（Pair）。
// 这一讲，我会用简化版的uniswap讲如何通过合约创建合约。

// 有两种方法可以在合约中创建新合约，create和create2
contract Pair{
    address public factory; // 工厂合约地址
    address public token0;  // 代币1
    address public token1;  // 代币2

    constructor() payable {
        factory = msg.sender;   // 记录创建者（工厂合约）记录创建这个 Pair 的工厂地址
    }

    // 只能被工厂合约调用（权限控制）
    function initialize(address _token0, address _token1) external {
        require(msg.sender == factory, 'UniswapV2: FORBIDDEN');
        token0 = _token0;
        token1 = _token1;
    }
}

contract PairFactory {
    // 通过两个代币地址查 Pair 地址
    mapping (address => mapping (address => address)) public getPair;
    // 保存所有Pair 地址
    address[] public allPairs;

    function craetePair(address tokenA, address tokenB) external returns(address pairAddr) {
        // 创建新合约
        Pair pair = new Pair();
        // 调用新合约的initialize 方法
        pair.initialize(tokenA, tokenB);
        // 更新map 地址
        pairAddr = address(pair);
        allPairs.push(pairAddr);
        // 双向映射
        getPair[tokenA][tokenB] = pairAddr;
        getPair[tokenB][tokenA] = pairAddr;
    }

}