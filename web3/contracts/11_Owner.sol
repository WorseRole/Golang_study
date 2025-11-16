// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// 构造函数和修饰器
contract Owner {

    address owner;

    // 构造函数
    // 构造函数（constructor）是一种特殊的函数，每个合约可以定义一个，并在部署合约的时候自动运行一次。
    // 它可以用来初始化合约的一些参数，例如初始化合约的owner地址
    constructor(address initialOwner) {
        owner = initialOwner;   // 在部署合约的时候，将owner设置为传入的 initialOwner 地址
    }
    // 注意：构造函数在不同的Solidity版本中的语法并不一致，
    // 在Solidity 0.4.22之前，构造函数不使用 constructor 而是使用与合约名同名的函数作为构造函数而使用，
    // 由于这种旧写法容易使开发者在书写时发生疏漏（例如合约名叫 Parents，构造函数名写成 parents），
    // 使得构造函数变成普通函数，引发漏洞，所以0.4.22版本及之后，采用了全新的 constructor 写法


    // 修饰器
    // 修饰器（modifier）是Solidity特有的语法，类似于面向对象编程中的装饰器（decorator），
    // 声明函数拥有的特性，并减少代码冗余。
    // 它就像钢铁侠的智能盔甲，穿上它的函数会带有某些特定的行为。
    // modifier的主要使用场景是运行函数前的检查，例如地址，变量，余额等
    // 定义modifier
    modifier onlyOwner {
        require(msg.sender == owner);   // 检查调用者是否为 owner 地址
        _;  // 如果是的话，继续运行函数主体；否则报错并revert 交易
    }
    // 带有onlyOwner 修饰符的函数只能被 owner 地址调用，比如下面这个例子：
    // 定义一个changeOwner 函数，运行它可以改变合约的owner，
    // 但是由于onlyOwner 修饰符的存在，只有原先的owner 可以调用，别人调用就会报错。
    // 这也是最常用的控制智能合约全线的方法。

    /* 报错信息：

    transact to Owner.changeOwner errored: Error occurred: revert.

revert
	The transaction has been reverted to the initial state.
Note: The called function should be payable if you send value and the value you send should be less than your current balance.
If the transaction failed for not having enough gas, try increasing the gas limit gently.
*/
    function changeOwner(address _newOwner) external onlyOwner {
        owner = _newOwner;
    }

    // 1. 在 Remix 上编译并部署代码,在合约部署时传入 initialOwner 变量。
    // 2. 点击 owner 按钮查看当前 owner 变量。
    // 3. 以 owner 地址的用户身份，调用 changeOwner 函数，交易成功。
    // 4. 以非 owner 地址的用户身份，调用 changeOwner 函数，交易失败，因为modifier onlyOwner 的检查语句不满足。

    // 控制合约权限的Ownable合约



}