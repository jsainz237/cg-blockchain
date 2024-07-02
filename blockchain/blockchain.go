package blockchain

import (
	"time"
)

type Blockchain struct {
	chain       []Block
	pendingData []interface{}
	difficulty  int
	blocktiming int
	reward      int
}

var MTGChain Blockchain

func init() {
	MTGChain = Blockchain{
		blocktiming: 10000,
		reward:      100,
		difficulty:  3,
	}

	// Create the genesis block
	genesisBlock := Block{
		Timestamp: time.Now(),
		Hash:      "000",
	}
	MTGChain.chain = append(MTGChain.chain, genesisBlock)
}

func (bc Blockchain) GetChain() []Block {
	return bc.chain
}

func (bc Blockchain) GetLatestBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) CreateBlock() {
	block := Block{
		Data:         bc.pendingData,
		Timestamp:    time.Now(),
		PreviousHash: bc.GetLatestBlock().Hash,
	}

	block.mineBlock(bc.difficulty)
	bc.chain = append(bc.chain, block)
	bc.pendingData = []interface{}{}
}

func (bc *Blockchain) IsValid() bool {
	// For each block in the chain, check if the hash is valid
	// and the link between blocks is correct

	for i := range bc.chain[1:] {
		currBlock := bc.chain[i]
		prevBlock := bc.chain[i-1]

		if currBlock.Hash != currBlock.calculateHash() || currBlock.PreviousHash != prevBlock.Hash {
			return false
		}
	}

	return true
}
