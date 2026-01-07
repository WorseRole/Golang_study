package main

import (
	"dapp/task1"
)

// import (
// 	"dapp/task01"
// 	"log"

// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// 测试
// func main() {
// 	// task1.SelectBlock()
// 	task1.FindTransactions()
// }

// const (
// 	URL = "https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2"
// )

func main() {
	// task1.FindTransactions()
	// task1.FindBlock()
	// task1.FindTrasaction()
	// task1.FindData()
	// task1.CreateWallet()
	task1.ETHTransfer()

	/*
	   // 1. 连接到Sepolia网络
	   client, err := ethclient.Dial(URL)

	   	if err != nil {
	   		log.Fatal("连接客户端失败: ", err)
	   	}

	   defer client.Close()

	   // 2. 查询区块信息（例如查询最新区块）
	   task01.QueryBlock(client)

	   // 3. 发送一笔转账交易
	   task01.SendTransaction(client)
	*/
}
