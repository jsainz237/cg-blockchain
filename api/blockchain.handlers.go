package api

import (
	"encoding/json"
	"io"
	bc "mtgbc/blockchain"
	network "mtgbc/network"
	"net/http"
)

type BlockchainHandlers struct{}

func (bh *BlockchainHandlers) GetBlockchain(r *http.Request, args *struct{}, reply *bc.Blockchain) error {
	*reply = bc.MTGChain
	return nil
}

func (bh *BlockchainHandlers) GetLatestBlock(r *http.Request, args *struct{}, reply *bc.Block) error {
	*reply = bc.MTGChain.GetLatestBlock()
	return nil
}

type ConsensusResponse struct {
	Authoritive bool
	Replaced    bool
}

func (bh *BlockchainHandlers) Consensus(r *http.Request, args *struct{}, reply *ConsensusResponse) error {
	var blockchains []bc.Blockchain

	// Get the blockchain from all connected nodes
	for _, node := range network.MTGNetwork.ConnectionPool {
		resp, _ := http.Get(node + "/blockchain")
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		blockchain := bc.Blockchain{}
		json.Unmarshal(body, &blockchain)
		blockchains = append(blockchains, blockchain)
	}

	// Replace the current blockchain with the longest one if it is valid
	for _, blockchain := range blockchains {
		if len(blockchain.Chain) > len(bc.MTGChain.Chain) && bc.IsValid(blockchain.Chain) {
			bc.MTGChain.Chain = blockchain.Chain
			bc.MTGChain.PendingData = blockchain.PendingData

			*reply = ConsensusResponse{Authoritive: false, Replaced: true}
			return nil
		}
	}

	*reply = ConsensusResponse{Authoritive: true, Replaced: false}
	return nil
}
