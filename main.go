package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client := ConnectClient("rinkeby.infura.io/v3/8e2834b158fa48b0a5fb9ca0f72ce6e6")

	// Get Block Number
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := header.Number

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block time: " + block.Time().String())
}

// Connexion à un noeud geth Rinkeby via Infura
func ConnectClient(string url) *Client {
	var client, err = ethclient.Dial(url) //"rinkeby.infura.io/v3/8e2834b158fa48b0a5fb9ca0f72ce6e6"
	if err != nil {
		fmt.Println("Error while connecting to infura")
		log.Fatal(err)
	}
	return client
}
