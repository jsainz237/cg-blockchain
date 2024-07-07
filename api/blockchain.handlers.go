package api

import (
	"encoding/json"
	"io"
	bc "mtgbc/blockchain"
	network "mtgbc/network"
	"net/http"

	"github.com/labstack/echo/v4"
)

var getBlockchainHandler = func(c echo.Context) error {
	return c.JSON(http.StatusOK, bc.MTGChain)
}

var getLatestBlockHandler = func(c echo.Context) error {
	return c.JSON(http.StatusOK, bc.MTGChain.GetLatestBlock())
}

var consensusHandler = func(c echo.Context) error {
	var blockchains []bc.Blockchain

	// Get the blockchain from all connected nodes
	for _, node := range network.MTGNetwork.ConnectionPool {
		resp, _ := http.Get("http://" + node + "/blockchain")
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

			return c.String(http.StatusOK, "Blockchain replaced")
		}
	}

	return c.String(http.StatusOK, "Blockchain is authoritative")
}
