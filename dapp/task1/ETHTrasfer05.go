package task1

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

func ETHTransfer() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2")
	if err != nil {
		log.Fatal("连接失败:", err)
	}

	// 从私钥导出发送方地址
	prvKey := ""
	privateKey, err := crypto.HexToECDSA(prvKey)
	if err != nil {
		log.Fatal("私钥解析失败: ", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥推导失败, err:", err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress = " + fromAddress.Hex())

	// 获取账户Nonce（交易序列号）
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取账户Nonce(交易序列号)失败, err:", err)
	}
	fmt.Println("nonce = " + fmt.Sprint(nonce))

	// 构建交易参数
	value := big.NewInt(1000000000000000) // 0.001ETH  in wei (1eth = 10^18)
	gasLimit := uint64(21000)             // 标准ETH转账的Gas上限
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取Gas Price失败, err:", err)
	}
	fmt.Println("value = " + fmt.Sprint(value))
	fmt.Println("gasLimit = " + fmt.Sprint(gasLimit))
	fmt.Println("gasPrice = " + fmt.Sprint(gasPrice))

	// 创建交易对象
	toAddress := common.HexToAddress("0xa5ead3165e1a5ef00D4f9E74f5baBCEe8709FC16")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	fmt.Println("toAddress = " + toAddress.Hex())

	// 获取网络ID
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("获取网络ID失败, err:", err)
	}
	fmt.Println("chainId = " + chainId.String())

	// 对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal("签名交易失败, err:", err)
	}
	fmt.Println("signedTx = " + signedTx.Hash().Hex())

	// 发送交易到网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}
