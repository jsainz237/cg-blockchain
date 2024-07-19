package blockchain

import (
	"time"
)

type Blockchain struct {
	Chain       []Block
	PendingData []Transaction
	difficulty  int
	blocktiming int
	reward      int
}

var MTGChain Blockchain

func init() {
	MTGChain = Blockchain{
		PendingData: []Transaction{},
		blocktiming: 10000,
		reward:      100,
		difficulty:  3,
	}

	// Create the genesis block
	genesisBlock := Block{
		Timestamp: time.Now(),
		Hash:      "000",
	}
	MTGChain.Chain = append(MTGChain.Chain, genesisBlock)
}

func (bc Blockchain) GetLatestBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) CreateBlock() {
	block := Block{
		Data:         bc.PendingData,
		Timestamp:    time.Now(),
		PreviousHash: bc.GetLatestBlock().Hash,
	}

	block.mineBlock(bc.difficulty)
	bc.Chain = append(bc.Chain, block)
	bc.PendingData = []Transaction{}
}

func IsValid(chain []Block) bool {
	// For each block in the chain, check if the hash is valid
	// and the link between blocks is correct

	for i := range chain[1:] {
		currBlock := chain[i]
		prevBlock := chain[i-1]

		if currBlock.Hash != currBlock.calculateHash() || currBlock.PreviousHash != prevBlock.Hash {
			return false
		}
	}

	return true
}
