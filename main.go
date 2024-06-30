package main

import (
	"encoding/json"
	"fmt"
	bc "mtgbc/blockchain"
)

func main() {
	blockchain := bc.CreateBlockchain()

	// Print the blockchain
	data, _ := json.Marshal(blockchain.GetChain())
	fmt.Println(string(data))
}
