const hre = require("hardhat");

async function main() {
  // 获取合约工厂
  const BeggingContract = await hre.ethers.getContractFactory("BeggingContract");
  
  // 部署合约
  const beggingContract = await BeggingContract.deploy();
  await beggingContract.waitForDeployment();
  
  // 获取合约地址
  const address = await beggingContract.getAddress();
  console.log("乞讨合约已部署到:", address);
  
  // 返回合约实例，方便后续交互
  return beggingContract;
}

// 执行部署
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});