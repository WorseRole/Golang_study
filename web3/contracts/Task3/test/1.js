// scripts/interact.js
const hre = require("hardhat");

async function main() {
  const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
  
  // 获取合约实例
  const BeggingContract = await hre.ethers.getContractFactory("BeggingContract");
  const beggingContract = BeggingContract.attach(contractAddress);

  // 获取当前乞讨次数
  const begCount = await beggingContract.begCount();
  console.log("当前乞讨次数:", begCount.toString());

  // 获取合约余额
  const balance = await hre.ethers.provider.getBalance(contractAddress);
  console.log("合约余额:", hre.ethers.formatEther(balance), "ETH");

  // 测试乞讨功能（如果需要支付ETH）
  // const [owner] = await hre.ethers.getSigners();
  // const tx = await beggingContract.beg({ value: hre.ethers.parseEther("0.001") });
  // await tx.wait();
  // console.log("乞讨成功！");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });