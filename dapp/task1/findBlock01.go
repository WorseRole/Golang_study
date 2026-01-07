package task1

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func FindBlock() {

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2")
	if err != nil {
		log.Fatal("连接失败:", err)
	}

	// 调用客户端的HeaderByNumber来返回有关一个区块的头信息。若指定nil，将返回最新的区块头。
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("调用客户端的HeaderByNumber来返回有关一个区块的头信息。若指定nil，将返回最新的区块头")
	fmt.Println("header.Number.String() = " + header.Number.String())

	// 调用客户端的BlockByNumber方法来获得完整区块。您可以读取该区块的所有内容和元数据，
	// 例如，区块号、区块计时器、区块摘要、区块入口以及交易列表等等。
	blockNumber := big.NewInt(9984964)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("调用客户端的BlockByNumber方法来获得完整区块。您可以读取该区块的所有内容和元数据")
	fmt.Println("block.Number().String() = " + block.Number().String())
	fmt.Println("block.Time() = " + fmt.Sprint(block.Time()))
	fmt.Println("block.Difficulty().Uint64() = " + fmt.Sprint(block.Difficulty().Uint64()))
	fmt.Println("block.Coinbase().Hex() = " + block.Coinbase().Hex())
	fmt.Println("block.Hash().Hex() = " + block.Hash().Hex())
	fmt.Println("len(block.Transactions()) = " + fmt.Sprint(len(block.Transactions())))

	// 查询交易
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)

}
