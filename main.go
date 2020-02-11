package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/ethereum/go-ethereum/ethclient"
)

type personAndRelations struct {
	person common.Address
	relations []relationAndWeight
}

type relationAndWeight struct {
	relation common.Address
	weight uint64
}


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

	fmt.Printf("-----------\nBlock Hash: %s\n", block.Hash().Hex())
	fmt.Printf("Block timestamp : %s\n-----------\n", block.ReceivedAt)

	for _, tx := range block.Transactions() {
		fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
		fmt.Printf("TX Value: %s\n", tx.Value().String())
		fmt.Printf("TX Gas: %d\n", tx.Gas())
		fmt.Printf("TX Gas Price: %d\n", tx.GasPrice().Uint64())
		fmt.Printf("TX Nonce: %d\n", tx.Nonce())
		fmt.Printf("TX To: %s\n", tx.To().Hex())

		// Doc : https://github.com/ethereum/go-ethereum/blob/master/core/types/transaction.go
		// type Message struct {
		// 	to         *common.Address
		// 	from       common.Address
		// 	nonce      uint64
		// 	amount     *big.Int
		// 	gasLimit   uint64
		// 	gasPrice   *big.Int
		// 	data       []byte
		// 	checkNonce bool
		// }
		msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("TX From : %s\n", msg.From().Hex())

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Receipt Status: %d\n", receipt.Status)
		fmt.Println("---")
	}

	// Grab block by hash then iterate over transactions by index
	blockHash := common.HexToHash("0x01e6b1caed7765220448df0979018f5613728ff1273f5de2f137393f4d583e5e")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
	}

	// Grab a transaction by it's individual hash
	txHash := common.HexToHash("0xa9a42eefa76655e5298996813138e6c33fac6f89506ae233c2f0b7a4e699ed68")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("Pending?: %v\n", isPending)
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

func base64Encode(input []byte) []byte {
	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(b64, input)

	return b64
}