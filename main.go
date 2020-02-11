package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	_ "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client := ConnectClient("https://rinkeby.infura.io/v3/8e2834b158fa48b0a5fb9ca0f72ce6e6")

	// Get Block Number
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := header.Number.String()
	fmt.Println(blockNumber)

	blockTest := big.NewInt(5952590)
	fmt.Println(blockTest)

	block, err := client.BlockByNumber(context.Background(), blockTest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Hash().Hex())

	for _, tx := range block.Transactions() {
		fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
		fmt.Printf("TX Value: %s\n", tx.Value().String())
		fmt.Printf("TX Gas: %d\n", tx.Gas())
		fmt.Printf("TX Gas Price: %d\n", tx.GasPrice().Uint64())
		fmt.Printf("TX Nonce: %d\n", tx.Nonce())
		//fmt.Printf("TX Data: %v\n", tx.Data())
		fmt.Printf("TX To: %s\n", tx.To().Hex())

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Receipt Status: %d\n", receipt.Status)
		fmt.Println("---")
	}
}

// Connexion à un noeud geth Rinkeby via Infura
func ConnectClient(url string) *Client {
	client, err := ethclient.Dial(url) //"rinkeby.infura.io/v3/8e2834b158fa48b0a5fb9ca0f72ce6e6"
	if err != nil {
		fmt.Println("Error while connecting to infura")
		log.Fatal(err)
	}
	return client
}
