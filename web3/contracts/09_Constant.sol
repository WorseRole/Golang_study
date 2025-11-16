// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// 常量相关的两个关键字，constant（常量） 和 immutable（不变量）。
// 状态变量声明这两个关键字之后，不能在初始化后更改数值。这样做的好处是提升合约的安全性并节省gas
// 另外，只有数值变量可以声明constant和immutable；string和bytes可以声明为constant，但不能为immutable。
contract Constant{

    // constant
    // constant变量必须在声明的时候初始化，之后再也不能改变。尝试改变的话，编译不通过
    uint256 constant CONSTANT_NUM = 10;
    string constant CONSTANT_STRING = "0xAA";
    bytes constant CONSTANT_BYTES = "WTF";
    address constant CONSTANT_ADDRESS = 0x0000000000000000000000000000000000000000;

    // immutable 
    // immutable变量可以在声明时或构造函数中初始化，因此更加灵活。
    // 在Solidity v0.8.21以后，immutable变量不需要显式初始化，
    // 未显式初始化的immutable变量将使用数值类型的初始值（见 8. 变量初始值）。反之，则需要显式初始化。 
    // 若immutable变量既在声明时初始化，又在constructor中初始化，会使用constructor初始化的值。
    // immutable 变量可以在constructor 里初始化，之后不能改变
    uint256 public immutable  IMMUTABLE_NUM = 9999999999;
    // 在`Solidity v0.8.21` 以后，下列变量数值暂为初始值 
    address public immutable IMMUTABLE_ADDRESS;
    uint256 public immutable IMMUTABLE_BLOCK;
    uint256 public immutable IMMUTABLE_TEST;
    
    // 你可以使用全局变量例如address(this)，block.number 或者自定义的函数给immutable变量初始化。
    // 在下面这个例子，我们利用了test()函数给IMMUTABLE_TEST初始化为9：
    constructor() {
        IMMUTABLE_ADDRESS = address(this);
        IMMUTABLE_NUM = 1118;
        IMMUTABLE_TEST = test();
    }
    function test() public pure returns(uint256) {
        uint256 what = 9;
        return what;
    }



}