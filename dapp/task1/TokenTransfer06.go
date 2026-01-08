package task1

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func TokenTransfer() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2")
	if err != nil {
		log.Fatal("连接失败:", err)
	}

	// 获取账户私钥
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	// 获取账户公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 获取账户地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress = " + fromAddress.Hex())

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce = " + fmt.Sprint(nonce))

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 要被转账的地址
	toAddress := common.HexToAddress("0xa5ead3165e1a5ef00D4f9E74f5baBCEe8709FC16")

	// 让我们将代币合约地址赋予标志。
	tokenAddress := common.HexToAddress("0xb38dde2F3D6eD620D80E5608f0343e2c97d74229")

	// 函数名将是传递函数的名称，即ERC-20规范中的transfer和参数类型。
	// 第一个参数类型是address（令牌的接收者），
	// 第二个类型是uint256（要发送的代币数量）。
	// 不需要没有空格和参数名称。我们还需要用字节切片格式。
	trasferFnSignature := []byte("transfer(address,uint256)")

	// 我们现在队列go-ethereum导入crypto/sha3包以生成函数签名的Keccak256存储。
	// 然后我们只使用前4个字节来获取方法ID。
	hash := sha3.NewLegacyKeccak256()
	hash.Write(trasferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println("hexutil.Encode(methodID) = " + hexutil.Encode(methodID))

	// 接下来，我们需要将给我们发送代币的地址左填充到 32 字节。
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println("hexutil.Encode(paddedAddress) = " + hexutil.Encode(paddedAddress))

	// 接下来我们确定要发送多少代币，在这个例子中是 1 个，并且我们需要在big.Int中格式化为 wei。
	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10) // 1 RDT token

	// 代币量还需要填充到32个字节。
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	// 接下来我们只需将方法 ID，填充后的地址和填充后的字节量，即可成为我们数据字段的字节片。
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 燃气上限制将取决于交易数据的大小和智能合约执行必须的计算步骤。幸运的是，客户端提供了EstimateGas方法，
	// 它可以为我们提示所需的燃气量。该函数从ethereum包中获取CallMsg结构，我们在其中指定数据和地址。
	// 将返回我们提示的完成交易所需的估计燃气上限。
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,   // 发送者
		To:   &tokenAddress, // 正确：指向合约地址
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasLimit = " + fmt.Sprint(gasLimit))

	// 接下来我们需要做的是构建交易类型，这类似于在 ETH 交互部分中的，除了_to_字段将是代币智能合约地址。
	// 这常让人困惑。我们还必须在调用中包含 0 ETH 的值字段并看到刚刚生成的数据字节。
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	// 下一步是使用发件人的私钥对事务进行签名。SignTx方法需要EIP155igner，需要我们先从客户端获取链ID。
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 最后发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx send: %s", signedTx.Hash().Hex())

}
