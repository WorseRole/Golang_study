package task1

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func FindData() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2")
	if err != nil {
		log.Fatal("连接失败:", err)
	}

	blockNumber := big.NewInt(9984964)
	blockHash := common.HexToHash("0x21876747a63a68b6dcafc61472926ccfaca3387b3be90bf63aae3e607b0ed2ff")
	// 可以调用 BlockReceipts 方法就可以获取指定区块中所有的收据列表。
	// 参数可以是区块的哈希值也可以是区块的高度。
	// 循环遍历集合并获取收集的信息.
	receiptByHash, err := client.BlockReceipts(context.Background(),
		rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal("获取区块失败:", err)
	}

	receiptByNum, err := client.BlockReceipts(context.Background(),
		rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal("获取区块失败:", err)
	}

	// fmt.Println(receiptByHash[0] == receiptByNum[0]) // false？？
	fmt.Println("深度比较结果:", reflect.DeepEqual(receiptByHash[0], receiptByNum[0])) // true

	for _, receipt := range receiptByHash {
		fmt.Println("receipt.Status = " + fmt.Sprint(receipt.Status))
		fmt.Println("receipt.Logs = " + fmt.Sprint(receipt.Logs))
		fmt.Println("receipt.TxHash.Hex() = " + receipt.TxHash.Hex())
		fmt.Println("receipt.TransactionIndex = " + fmt.Sprint(receipt.TransactionIndex))
		fmt.Println("receipt.ContractAddress.Hex() = " + receipt.ContractAddress.Hex())
		break
	}

	txHash := common.HexToHash("0x8241fbfa6f3c656b2d83d34d07a30bd95293b8040ed9d2281577175a39b968c7")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal("获取区块失败:", err)
	}
	fmt.Println("receipt.Status = " + fmt.Sprint(receipt.Status))
	fmt.Println("receipt.Logs = " + fmt.Sprint(receipt.Logs))
	fmt.Println("receipt.TxHash.Hex() = " + receipt.TxHash.Hex())
	fmt.Println("receipt.TransactionIndex = " + fmt.Sprint(receipt.TransactionIndex))
	fmt.Println("receipt.ContractAddress.Hex() = " + receipt.ContractAddress.Hex())

}
