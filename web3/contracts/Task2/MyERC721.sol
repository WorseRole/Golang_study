// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/utils/Counters.sol";


contract MyERC721 {
    using Counters for Counters.Counter;

    // 事件
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
    event ApprovalFroAll(address indexed owner, address indexed operator, bool approved);

    // 状态变量
    string public name;
    string public symbol;

    mapping (uint256 => address) private _owners;
    mapping (address => uint256) private _balances;
    mapping (uint256 => address) private _tokenApprovals;
    mapping (address => mapping(address => bool)) private _operatorApprovals;

    Counters.Counter private _tokenIdCounter;

    constructor(string memory _name, string memory _symbol) {
        name = _name;
        symbol = _symbol;
    }

    // 基础查询函数
    function balanceOf(address owner) public view returns(uint256) {
        require(owner != address(0), "Query for zero address");
        return _balances[owner];
    }

    function ownerOf(uint256 tokenId) public view returns(address) {
        address owner = _owners[tokenId];
        require(owner != address(0), "Token does not exist");
        return owner;
    }

    // 授权函数 
    function approve(address to, uint256 tokenId) public {
        address owner = ownerOf(tokenId);
        require(to != owner, "Approval to current owner");
        require(msg.sender == owner || isApprovedForAll(owner, msg.sender), ("Not owner nor approved for all") );

        _tokenApprovals[tokenId] = to;
        emit Approval(owner, to, tokenId);
    }

    // 通过tokenId 获取 授权者地址
    function getApproved(uint256 tokenId) public view returns(address) {
        require(_owners[tokenId] != address(0), "Token does not exist");
        return  _tokenApprovals[tokenId];
    }

    // 设置操作者授权状态
    function setApprovalForAll(address operator, bool approved) public {
        require(operator != msg.sender, "Approve to caller");
        _operatorApprovals[msg.sender][operator] = approved;
        emit ApprovalFroAll(msg.sender, operator, approved);
    }
 
    function isApprovedForAll(address owner, address operator) public view returns(bool) {
        return _operatorApprovals[owner][operator];
    }

    function transferFrom(address from, address to, uint256 tokenId) public {
        require(_isApprovedOrOwner(msg.sender, tokenId), "Not owner nor approved");
        require(to != address(0), "Transfer to zero address");

        // 清除授权
        _approve(address(0), tokenId);

        // 更新余额
        _balances[from] -= 1;
        _balances[to] += 1;
        _owners[tokenId] = to;

        emit Transfer(from, to, tokenId);
    }


    function safeTransferFrom(address from, address to, uint256 tokenId) public {
        safeTransferFrom(from, to, tokenId, "");
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public {
        transferFrom(from, to, tokenId);
        // require(_checkOnERC721Received(from, to, tokenId, data), "Transfer to non ERC721Receiver");
        require(_checkOnERC721Received(data), "Transfer to non ERC721Receiver");
    }

    function _approve(address to, uint256 tokenId) internal {
        _tokenApprovals[tokenId] = to;
        emit Approval(ownerOf(tokenId), to, tokenId);
    }

    // 铸造函数
    function mint(address to) public returns (uint256) {
        require(to != address(0), "Mint to zero address");

        _tokenIdCounter.increment();
        uint256 tokenId = _tokenIdCounter.current();

        _balances[to] += 1;
        _owners[tokenId] = to;

        emit Transfer(address(0), to, tokenId);
        return tokenId;
    }

    // function _checkOnERC721Received(address from, address to, uint256 tokenId, bytes memory data) internal pure returns (bool) {
    function _checkOnERC721Received(bytes memory data) internal pure returns(bool) {
        if(data.length == 0) {
            return true;
        }
        return true;
    }

    function _isApprovedOrOwner(address spender, uint256 tokenId) internal view returns(bool) {
        address owner = ownerOf(tokenId);
        return (spender == owner ||  
                    isApprovedForAll(owner, spender) || 
                    getApproved(tokenId) == spender) ;
    }

    function tokenURI(uint256 tokenId) public view returns(string memory) {
        require(_owners[tokenId] != address(0), "Token does not exist");
        return string(abi.encodePacked("https://example.com/token/", _toString(tokenId)));
    }

    function _toString(uint256 value) internal pure returns(string memory) {
        if (value == 0) return "0";
        uint256 temp = value;
        uint256 digits;
        while (temp != 0) {
            digits++;
            temp /= 10;
        }
        bytes memory buffer = new bytes(digits);
        while (value != 0) {
            digits -= 1;
            buffer[digits] = bytes1(uint8(48 + uint256(value % 10)));
            value /= 10;
        }
        return string(buffer);
    }


}