package task01

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 首先 我先写一个交易 然后拿到交易的hash 然后再拿到区块信息
// 0xa5ead3165e1a5ef00D4f9E74f5baBCEe8709FC16

const (
	URL        = "https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2"
	privateKey = ""
	toAddress  = "0xa5ead3165e1a5ef00d4f9e74f5babcee8709fc16"
)

/*
>>> 开始查询区块信息...
区块高度: 9984964
区块哈希: 0x21876747a63a68b6dcafc61472926ccfaca3387b3be90bf63aae3e607b0ed2ff
区块时间戳: 1767631524 (Unix时间戳)
交易数量: 140
矿工地址: 0x13CB6AE34A13a0977F4d7101eBc24B87Bb23F0d5
--- 查询完毕 ---
>>> 开始发送交易...
交易发送成功！
交易哈希: 0x8241fbfa6f3c656b2d83d34d07a30bd95293b8040ed9d2281577175a39b968c7
--- 交易完毕 ---
*/

func QueryBlock(client *ethclient.Client) {
	fmt.Println(">>> 开始查询区块信息...")
	// 获取最新区块号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("获取最新区块头失败: ", err)
	}
	blockNumber := header.Number

	// 获取完整的区块信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("获取区块失败: ", err)
	}
	// 打印关键信息
	fmt.Printf("区块高度: %d\n", block.Number().Int64())
	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("区块时间戳: %d (Unix时间戳)\n", block.Time())
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))
	fmt.Printf("矿工地址: %s\n", block.Coinbase().Hex())
	fmt.Println("--- 查询完毕 ---")
}

func SendTransaction(client *ethclient.Client) {
	fmt.Println(">>> 开始发送交易...")

	// 1. 从私钥导出发送方地址
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("私钥解析失败: ", err)
	}
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥推导失败")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 2. 获取账户Nonce（交易序列号）
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取Nonce失败: ", err)
	}

	// 3. 获取当前网络建议的Gas价格
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatal("获取Gas Tip失败: ", err)
	}
	gasFeeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取Gas Fee失败: ", err)
	}

	// 4. 构建交易参数
	value := big.NewInt(1000000000000000) // 转账金额：0.001 ETH (单位：Wei, 1 ETH = 10^18 Wei)
	gasLimit := uint64(21000)             // 标准ETH转账的Gas上限

	// 5. 创建交易对象 (EIP-1559动态费用交易类型)
	recipient := common.HexToAddress(toAddress)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(11155111), // Sepolia 网络的 ChainID
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &recipient,
		Value:     value,
		Data:      nil, // 普通转账，Data为空
	})

	// 6. 对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(11155111)), privateKeyECDSA)
	if err != nil {
		log.Fatal("交易签名失败: ", err)
	}

	// 7. 发送交易到网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("发送交易失败: ", err)
	}

	// 8. 输出交易哈希（这是交易的唯一标识）
	fmt.Printf("交易发送成功！\n交易哈希: %s\n", signedTx.Hash().Hex())
	fmt.Println("--- 交易完毕 ---")

}
