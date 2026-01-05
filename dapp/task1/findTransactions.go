package task1

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FindTransactions() {
	fmt.Println("=== 开始查找交易 ===")

	// 1. 连接客户端
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer client.Close()

	// 2. 获取网络 ChainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("获取 ChainID 失败:", err)
	}
	fmt.Printf("网络 ChainID: %d\n", chainID.Uint64())

	// 3. 查询指定区块 (你之前查询的高度)
	blockNumber := big.NewInt(9984964)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("获取区块失败:", err)
	}
	fmt.Printf("正在分析区块 #%d，内含 %d 笔交易\n", blockNumber, len(block.Transactions()))

	// 4. 遍历并打印该区块中的交易详情 (示例：只处理前3笔避免输出过多)
	fmt.Println("\n--- 区块内交易列表 ---")
	for i, tx := range block.Transactions() {
		if i >= 3 { // 限制输出，避免刷屏
			fmt.Printf("... (已截断，共 %d 笔交易)\n", len(block.Transactions()))
			break
		}
		fmt.Printf("\n交易 #%d:\n", i)
		fmt.Printf("  Hash: %s\n", tx.Hash().Hex())
		fmt.Printf("  From: ")
		if sender, err := types.Sender(types.LatestSignerForChainID(chainID), tx); err == nil {
			fmt.Printf("%s\n", sender.Hex())
		} else {
			fmt.Printf("获取失败: %v\n", err)
		}
		if tx.To() != nil {
			fmt.Printf("  To:   %s\n", tx.To().Hex())
		} else {
			fmt.Printf("  To:   (合约创建)\n")
		}
		fmt.Printf("  Value: %s Wei\n", tx.Value().String())
	}

	// 5. 按区块哈希查询交易数量 (你遇到的错误点，现已放入函数内)
	fmt.Println("\n--- 根据区块哈希查询 ---")
	// 注意：这里应使用你查询的区块的哈希，而不是交易哈希
	correctBlockHash := block.Hash() // 直接使用上面查询到的区块的哈希
	// 如果你想手动指定，可以用：common.HexToHash("0x21876747a63a68b6dcafc61472926ccfaca3387b3be90bf63aae3e607b0ed2ff")

	count, err := client.TransactionCount(context.Background(), correctBlockHash)
	if err != nil {
		log.Fatal("根据区块哈希获取交易数量失败:", err)
	}
	fmt.Printf("区块 %s 包含 %d 笔交易\n", correctBlockHash.Hex(), count)

	// 6. 根据你的交易哈希查询特定交易 (你发送的那笔)
	fmt.Println("\n--- 根据交易哈希查询 ---")
	yourTxHash := common.HexToHash("0x8241fbfa6f3c656b2d83d34d07a30bd95293b8040ed9d2281577175a39b968c7")
	tx, isPending, err := client.TransactionByHash(context.Background(), yourTxHash)
	if err != nil {
		log.Fatal("根据交易哈希查询失败:", err)
	}
	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())
	fmt.Printf("是否在待处理池中: %v\n", isPending)
	if !isPending {
		// 获取交易收据，查看状态和Gas使用情况
		receipt, err := client.TransactionReceipt(context.Background(), yourTxHash)
		if err != nil {
			log.Fatal("获取交易收据失败:", err)
		}
		fmt.Printf("交易状态 (1=成功, 0=失败): %d\n", receipt.Status)
		fmt.Printf("实际使用的 Gas: %d\n", receipt.GasUsed)
	}
	fmt.Println("=== 查找完成 ===")
}
