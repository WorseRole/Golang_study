// hardhat.config.js
require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

module.exports = {
  solidity: "0.8.20", // 务必与你合约中的 pragma 版本一致
  networks: {
    // 配置Sepolia测试网
    sepolia: {
      url: `https://sepolia.infura.io/v3/${process.env.SEPOLIA_PRIVATE_KEY}`, // 或使用Infura等其他节点服务
      accounts: [process.env.PRIVATE_KEY] // 部署者钱包私钥
    },
    // 本地开发网络
    localhost: {
      url: "http://127.0.0.1:8545"
    }
  }
};