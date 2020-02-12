package main

import (
	"context"
	"fmt"
	"log"
	// "time"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
)

type blockIndexMngmt struct {
	timestamp string
	input_ID uint
	from_block uint
	to_block uint
}

type txContent struct {
	from_Addr string
	to_Addr string
	input_ID uint
	txHash string
	blockNumber uint
}

func main() {
	client := ConnectClient("https://rinkeby.infura.io/v3/8e2834b158fa48b0a5fb9ca0f72ce6e6")

	// >>> START saving Mongo Connection to DB
	// clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	// clientMongo, err := mongo.Connect(context.TODO(), clientOptions)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = clientMongo.Ping(context.TODO(), nil)

	// if err != nil {
	// 	fmt.Println("Connect to MongoDB failed !")
	// 	log.Fatal(err)
	// }
	// fmt.Println("blockIndex : Connected to MongoDB!")
	// >>> END saving Mongo Connection to DB

	//
	// Request to API to know the last block index scanned
	// Table blockIndexation { ID (auto inc) ; timestamp ; input_ID (auto inc) ; from_block ; to_block }
	//	Process : get newest timestamp (or max ID if auto inc)
	//			then get associated to_block
	//			set blockIndex = to_block + 1
	//
	// >>> START Mongo get block starting index
	// filter := bson.D{} // Default filter hould become maxID or maxTimestamp
	// var resultIndex blockIndexMngmt

	// collectionBlock := clientMongo.Database("mydb").Collection("blockIndex")
	// err = collectionBlock.FindOne(context.TODO(), filter).Decode(&resultIndex)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Found a single document: %+v\n", resultIndex)

	// findOptions := options.Find()
	// findOptions.SetLimit(1)
	// >>> END Mongo get block starting index

	startBlockIndex := 5954000
	// startBlockIndex = resultIndex.to_block + 1

	// currInputID := 0
	// currInputID := resultIndex.input_ID + 1

	blockIndex := startBlockIndex

	for {
		blockTest := big.NewInt(int64(blockIndex))
		block, err := client.BlockByNumber(context.Background(), blockTest)
		if err != nil {
			log.Fatal(err)
			break
		}

		fmt.Printf("---------------------------------\nBlock Number: %d\n", blockIndex)
		fmt.Printf("Block Hash: %s\n", block.Hash().Hex())
		fmt.Printf("Block timestamp : %s\n---------------------------------\n", block.ReceivedAt)

		for _, tx := range block.Transactions() {
			fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
			fmt.Printf("TX Value: %d\n", tx.Value())
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
			fmt.Println("---------------------------------")


			//
			// Request to API to insert the From and To Address
			// Table UserRelations { ID (auto_inc) ; from_Addr ; to_Addr ; ID_input (same value than current scan in Table blockIndexation) }
			//
			// >>> START current txData saving
			// collectionTx := clientMongo.Database("mydb").Collection("txData")

			// currTx := txContent{string(msg.From().Hex()), string(tx.To().Hex()), uint(currInputID), tx.Hash().Hex(), uint(blockIndex)}

			// insertResult, err := collectionTx.InsertOne(context.TODO(), currTx)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// fmt.Println("txData : Insert success with ID: ", insertResult.InsertedID)
			// >>> END current txData saving
		}

		blockIndex++
	}

	//
	// Request to API to save the last block index scanned
	// Table blockIndexation { ID (auto inc) ; timestamp ; input_ID (auto inc) ; from_block ; to_block }
	//	Insert : currTime ; from_block = startBlockIndex ; to_block = blockIndex
	//
	// >>> START save last block index scanned
	// currTimestamp := time.Now().Format(time.RFC3339)

	// collectionTx := clientMongo.Database("mydb").Collection("blockIndex")

	// finalIndex := blockIndexMngmt{currTimestamp, uint(currInputID), uint(startBlockIndex), uint(blockIndex)}

	// insertResult, err := collectionTx.InsertOne(context.TODO(), finalIndex)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("blockIndex : Insert success with ID: ", insertResult.InsertedID)
	// >>> END save last block index scanned
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