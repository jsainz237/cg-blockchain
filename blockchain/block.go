package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Data         []interface{}
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	nonce        int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.nonce)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mineBlock(difficulty int) Block {
	// Hash must start with `difficulty` number of 0s (PoW)
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.nonce++
		b.Hash = b.calculateHash()
	}

	return *b
}
