package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBlockchainHandler(sc ServerContext) ServerContextHandler {
	blockchain := sc.Blockchain
	return func(e echo.Context) error {
		return e.JSON(http.StatusOK, blockchain.GetChain())
	}
}

func GetLatestBlockHandler(sc ServerContext) ServerContextHandler {
	blockchain := sc.Blockchain
	block := blockchain.GetLatestBlock()
	return func(e echo.Context) error {
		return e.JSON(http.StatusOK, block)
	}
}
