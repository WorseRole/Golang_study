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

// 当使用BlockByNumber方法获取到完整的区块信息之后，可以调用区块实例的Transactions方法来读取块中的交易，
// 该方法返回一个Transaction类型的列表。
func FindTrasaction() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2")
	if err != nil {
		log.Fatal("连接失败:", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 调用客户端的BlockByNumber方法来获得完整区块。您可以读取该区块的所有内容和元数据，
	// 例如，区块号、区块计时器、区块摘要、区块入口以及交易列表等等。
	blockNumber := big.NewInt(9984964)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Println(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println("tx.Hash().Hex() = " + tx.Hash().Hex())
		fmt.Println("tx.Value().String() = " + tx.Value().String())
		fmt.Println("tx.Gas() = " + fmt.Sprint(tx.Gas()))
		fmt.Println("tx.GasPrice().Uint64() = " + fmt.Sprint(tx.GasPrice().Uint64()))
		fmt.Println("tx.Nonce() = " + fmt.Sprint(tx.Nonce()))
		fmt.Println(tx.Data())
		fmt.Println("tx.To().Hex() = " + tx.To().Hex())

		// 为了读取发送方的地址，我们需要在事务上调用AsMessage，它返回一个Message类型，
		// 其中包含一个返回发送方（from）地址的函数。AsMessage方法需要EIP155签名者。
		if sender, err := types.Sender(types.LatestSignerForChainID(chainId), tx); err == nil {
			fmt.Println("sender = ", sender.Hex())
		} else {
			log.Fatal(err)
		}

		fmt.Println("tx.To() = " + fmt.Sprint(tx.To()))

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receipt.Status = ", receipt.Status)
		fmt.Println(receipt.Logs)
		break
	}

	blockHash := common.HexToHash("0x21876747a63a68b6dcafc61472926ccfaca3387b3be90bf63aae3e607b0ed2ff")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// 在无法获取块的情况下查找事务的另一种方式是调用客户端的TransactionInBlock方法。
	// 该方法仅接受块分布和块内事务的索引值。
	// 调用TransactionCount来了解有多少个事务。
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("tx.Hash().Hex() = " + tx.Hash().Hex())
		fmt.Println("tx.To() = " + fmt.Sprint(tx.To()))
		break
	}

	// 可以使用TransactionByHash在给定具体事务哈希值的情况下直接查询单个事务。
	txHash := common.HexToHash("0x8241fbfa6f3c656b2d83d34d07a30bd95293b8040ed9d2281577175a39b968c7")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("isPending = " + fmt.Sprint(isPending))
	fmt.Println("tx.Hash().Hex() = " + tx.Hash().Hex())
}
