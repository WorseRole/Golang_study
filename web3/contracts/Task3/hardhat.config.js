require("dotenv").config(); // ← 添加这一行！
require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.21",
  networks: {
    // Hardhat 内置的本地开发网络，启动后默认URL
    localhost: {
      url: "http://127.0.0.1:8545", // 启动 'npx hardhat node' 后的默认URL[citation:5]
    },
    sepolia: {
      url: process.env.SEPOLIA_RPC_URL, // 从环境变量获取 Sepolia 测试网的 RPC URL
      accounts: [process.env.PRIVATE_KEY],
    }
  }
}
