package task1

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FindTokenBalance() {
	fmt.Println("find token balance")

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/411f9e3048cb4427b19f3fe66c5b95e2")
	if err != nil {
		log.Fatal(err)
	}
	// 查找token余额 合约地址
	tokenAddress := common.HexToAddress("0xb38dde2F3D6eD620D80E5608f0343e2c97d74229")
	// 在同一个包下 就不用导入
	instance, err := NewMytoken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 要查询的地址（你的钱包地址） 我帐户地址
	accountAddress := common.HexToAddress("0x14503cecD68735b3E02b9Ae849FE5e29A9Bf7229")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, accountAddress)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name = " + name)
	fmt.Println("symbol = " + symbol)
	fmt.Println("balance = " + bal.String())
	fmt.Printf("decimals: %v\n", decimals)
	fmt.Printf("wei: %v\n", bal)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	ethValue := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Println("ethValue = " + ethValue.String())
}
