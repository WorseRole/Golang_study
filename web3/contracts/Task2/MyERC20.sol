// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

/*
合约包含以下标准 ERC20 功能：
balanceOf：查询账户余额。
transfer：转账。
approve 和 transferFrom：授权和代扣转账。
使用 event 记录转账和授权操作。
提供 mint 函数，允许合约所有者增发代币。
*/
contract MyERC20 {
    // 3. 使用 event 记录转账和授权操作。
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    modifier onlyOwner() {
        require(msg.sender == _owner, "ERC20: caller is not the owner");
        _;
    }

    // 状态变量
    // 账户余额映射
    mapping (address account => uint256) private _balances;
    // 授权额度映射
    mapping (address account => mapping (address spender => uint256)) private _allowances;
    // 总供应量
    uint256 private _totalSupply;
    // 代币名称
    string private _name;
    // 代币类型
    string private _symbol;

    // 发币的话 还是保证自己发
    address private _owner;

    // 构造  设置代币的名称和符号 ("MyToken, MTK)
    constructor (string memory name_, string memory symbol_) {
        _name = name_;
        _symbol = symbol_;
        _owner = msg.sender;
    }

    function name() public view returns(string memory) {
        return _name;
    }
    function symbol() public view returns(string memory) {
        return _symbol;
    }
    // 显示和计算精度: 
        // 18位小数是Ethereum 上最常见的标准（像 ETH、USDT、USDC 都用18）
        // 这意味着 1 个完整的代币 = 10^18 个最小单位
        // 例如 1.5 MTK 在合约内部存储为 1500000000000000000
    // 钱包正确显示 
        // 没有 decimals 函数的话 你的余额 1500000000000000000
        // 有 decimal 函数的话，你的余额 1.5 MTK
    function decimals() public pure returns(uint8) {
        return 18;
    }
    function totalSupply() public view returns(uint256) {
        return _totalSupply;
    }

    // 核心功能：
    // 0. 查询账户余额
    function balanceOf (address account) view external returns(uint256) {
        return _balances[account];
    }
    // 1. 转账
    function transfer(address to, uint256 value) public returns(bool) {
        address owner = msg.sender;
        _transfer(owner, to, value);
        return true;
    }

    function _transfer(address from, address to, uint256 value) internal {
        // require(条件为false,  就断了)
        // 比如 from 为 address(0)   != 就为false了
        require(from != address(0), "Transfer from zero address");
        require(to != address(0), "Transfer to zero address");
        _update(from, to, value);
    }

    function _update(address from, address to, uint256 value) internal  {
        // from 为 address(0) 的话 就是铸币了
        if(from == address(0)) {
            _totalSupply += value;
        } else  {
            // 看够不够转出
            uint256 fromBalance = _balances[from];
            if (fromBalance < value) {
                // 终止
                revert("Insufficient balance");
            }
            // 默认算术溢出检查
            unchecked {
                _balances[from] -= value;
            }
        }

        // 转出地址为0 的话，就是纯消耗币了
        if(to == address(0)) {
            unchecked {
                _totalSupply -= value;
            }
        } else {
            unchecked {
                _balances[to] += value;
            }
        }

        emit Transfer(from, to, value);
    }

    // 2. approve 和 transferFrom：授权和代扣转账。
    // 2.1 approve 授权
    function approve(address spender, uint256 value) public returns (bool) {
        address owner = msgSender();
        _approve(owner, spender, value, true);
        return true;
    }
    // 对 _allowance[owner][spender] = value 嵌套授权
    function _approve(address owner, address spender, uint256 value, bool emitEvent) internal {
        require(owner != address(0), "Approval from zero address");
        require(spender != address(0), "Approval to zero address");
        _allowances[owner][spender] = value;
        if(emitEvent) {
            emit Approval(owner, spender, value);
        }
    }

    // 2.2 代扣转账 代扣 顾名思义 带你扣，替你扣，用你自己替你扣，更像是做了层监管 我合约这块替你扣，所以 spender = msg.sender
    function transferFrom(address from, address to, uint256 value) public returns(bool){
        address spender = msgSender();
        // 检查授权代扣钱够不够，并更新授权代扣额度
        _spendAllowance(from, spender, value);
        // 转账
        _transfer(from, to, value);
        return true;
    }

    // 代扣转账授权 判断+更新
    function _spendAllowance(address owner, address spender, uint256 value) internal {
        // 检查代扣授权的钱够不够
        uint256 currentAllowance = _allowances[owner][spender];
        if(currentAllowance < value) {
            revert();
        }
        // 重新更新授权
        unchecked {
            _approve(owner, spender, currentAllowance - value, false);
        }
    }

    function mint(address account, uint256 value) external onlyOwner returns(bool) {
        _mint(account, value);
        return true;
    }

    // 提供 mint 函数，允许合约所有者增发代币。     external 是否要用成 internal 或者判断一下
    function _mint(address account, uint256 value) internal {
        // check address
        require(account != address(0), "Mint to zero address");
        _update(address(0), account, value);
    }

    function msgSender() internal view returns (address) {
        return msg.sender;
    }

}