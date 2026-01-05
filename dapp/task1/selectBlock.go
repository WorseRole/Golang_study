package task1

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

// 有两种方式查询区块信息。
// 区块头
// 调用客户端的 HeaderByNumber 来返回有关一个区块的头信息。若传入 nil，它将返回最新的区块头。
// 完整区块
// 调用客户端的 BlockByNumber 方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。

// 有两种方式查询区块信息.
// 一种是通过 HeaderByNumber 查询区块头,另一种是通过 BlockByNumber 查询完整区块.

// 查询交易
// 有两种方式查询交易信息。
// 一种是通过 TransactionByHash 查询交易信息,另一种是通过 TransactionByBlockHashAndIndex 查询交易信息。

func SelectBlock() {
	fmt.Println("select block")
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/i-HtbuxBqV8BpZlCF4zSN")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(9705177)

	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	fmt.Println(header.Number.Uint64())     // 9705177
	fmt.Println(header.Time)                // 1764091428
	fmt.Println(header.Difficulty.Uint64()) // 0
	fmt.Println(header.Hash().Hex())        // 0x1764b5fa60015d827f290ab4ab7ed9a71f6e668b3277d51bf227a4846c3e08ed

	if err != nil {
		log.Fatal(err)
	}
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 9705177
	fmt.Println(block.Time())                // 1764091428
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0x1764b5fa60015d827f290ab4ab7ed9a71f6e668b3277d51bf227a4846c3e08ed
	fmt.Println(len(block.Transactions()))   // 145
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 145
}
