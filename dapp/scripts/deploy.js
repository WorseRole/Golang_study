// scripts/deploy.js
const hre = require("hardhat");

async function main() {
  // 获取合约工厂
  const MyToken = await hre.ethers.getContractFactory("MyToken");
  
  // 部署合约，传入构造函数参数（initialOwner地址）
  // 注意：这里将部署者钱包（即私钥对应的地址）设为合约的初始所有者
  const initialOwner = (await hre.ethers.getSigners())[0].address;
  const myToken = await MyToken.deploy(initialOwner);

  // 等待合约部署完成
  await myToken.waitForDeployment();

  const contractAddress = await myToken.getAddress();
  console.log("MyToken 合约部署成功！地址:", contractAddress);
  console.log("合约所有者（initialOwner）:", initialOwner);


  // MyToken 合约部署成功！
  // 地址: 0xb38dde2F3D6eD620D80E5608f0343e2c97d74229
  // 合约所有者（initialOwner）: 0x14503cecD68735b3E02b9Ae849FE5e29A9Bf7229
  
  return contractAddress;
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});