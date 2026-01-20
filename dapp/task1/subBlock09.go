package task1

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubBlock() {
	fmt.Println("sub block")
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/411f9e3048cb4427b19f3fe66c5b95e2")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("block.Hash().Hex() = " + block.Hash().Hex())
			fmt.Println("block.Number().Uint64() = " + fmt.Sprint(block.Number().Uint64()))
			fmt.Println("block.Time() = " + fmt.Sprint(block.Time()))
			fmt.Println("block.Nonce() = " + fmt.Sprint(block.Nonce()))
			fmt.Println("len(block.Transactions()) = " + fmt.Sprint(len(block.Transactions())))
		}
	}

}
