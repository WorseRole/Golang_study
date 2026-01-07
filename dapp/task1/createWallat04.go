package task1

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func CreateWallet() {
	// 首先生成一个新的钱包，我们需要导入go-ethereumcrypto包，该包提供了用于生成随机私钥的GenerateKey方法。
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("生成钱包失败: ", err)
	}
	fmt.Println("crypto.GenerateKey(), privateKey = ", privateKey)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("hexutil.Encode(privateKeyBytes)[2:]" + hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 去掉‘0x04’
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位

}
