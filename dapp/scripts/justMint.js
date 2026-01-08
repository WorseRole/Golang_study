// scripts/justMint.js
const hre = require("hardhat");

async function main() {
  console.log("🪙 为Go手动转账测试准备RDT代币余额\n");
  
  const [deployer] = await hre.ethers.getSigners();
  console.log("操作账户:", deployer.address);
  
  // 连接到已部署的合约 
  const tokenAddress = "0xb38dde2F3D6eD620D80E5608f0343e2c97d74229";
  const MyToken = await hre.ethers.getContractFactory("MyToken");
  const token = await MyToken.attach(tokenAddress);
  
  // 检查当前余额
  const currentBalance = await token.balanceOf(deployer.address);
  console.log(`当前RDT余额: ${hre.ethers.formatUnits(currentBalance, 18)} RDT`);
  
  // 如果余额不足，铸造新代币
  const desiredBalance = hre.ethers.parseEther("10000"); // 目标：10,000 RDT
  if (currentBalance < desiredBalance) {
    console.log("\n🔄 铸造代币...");
    
    // 计算需要支付的ETH（根据RATE = 100000000）
    // 1 ETH = 100,000,000 RDT
    // 需要 10,000 RDT = 10,000 / 100,000,000 = 0.0001 ETH
    const ethNeeded = hre.ethers.parseEther("0.0001");
    
    // 但合约要求 MIN_ETH = 0.001 ETH，所以支付最小值
    const ethToSend = hre.ethers.parseEther("0.001"); // 合约要求的最低值
    
    console.log(`支付: ${hre.ethers.formatEther(ethToSend)} ETH`);
    console.log(`将获得: ${hre.ethers.formatEther(ethToSend * 100000000n)} RDT`);
    
    const mintTx = await token.mint({
      value: ethToSend
    });
    
    console.log(`铸造交易哈希: ${mintTx.hash}`);
    console.log("等待确认...");
    await mintTx.wait();
    console.log("✅ 铸造成功！");
    
    // 显示新余额
    const newBalance = await token.balanceOf(deployer.address);
    console.log(`\n💰 新RDT余额: ${hre.ethers.formatUnits(newBalance, 18)} RDT`);
    console.log(`相当于: ${hre.ethers.formatEther(newBalance)} RDT (无小数格式化)`);
  } else {
    console.log("✅ 余额已足够，无需铸造");
  }
  
  console.log("\n🎯 现在可以在Go中测试手动转账了！");
  console.log("使用以下地址进行测试：");
  console.log(`合约地址: ${tokenAddress}`);
  console.log(`你的地址: ${deployer.address}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });