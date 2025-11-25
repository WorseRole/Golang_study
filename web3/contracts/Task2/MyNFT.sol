// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

//  已在测试环境 进行铸造
// 合约地址: 0x6c611EA59184B8828DaEbd588fB4C35950e1fc61
// Token ID: 1
// 交易哈希: 0x19791b72fc6d4f88e185e1ac927ae1bc6ae77b3e31df6de10c642a42c9c6031a

contract MyNFT is ERC721, ERC721URIStorage, Ownable {

    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    // Ownable 设置合约所有者
    constructor() ERC721("LYTXTNFT", "LYTXTNFT") Ownable(msg.sender) {}

    function mintNFT(address to, string memory _tokenURI) public returns(uint256) {
        _tokenIds.increment();
        uint256 newTokenId = _tokenIds.current();

        _mint(to, newTokenId);
        _setTokenURI(newTokenId, _tokenURI);

        return newTokenId;
    }

    // 重写
    function tokenURI(uint256 tokenId) public view override (ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId) public view override (ERC721, ERC721URIStorage) returns (bool) {
        return super.supportsInterface(interfaceId);
    }

}