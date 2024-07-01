package server

import (
	bc "mtgbc/blockchain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ServerContext struct {
	Blockchain *bc.Blockchain
	Router     *echo.Echo
}

func StartNode(serverContext *ServerContext, port string) *ServerContext {
	router := echo.New()
	router.HideBanner = true

	serverContext.Router = router

	initRoutes(serverContext)

	router.Logger.Fatal(router.Start(port))
	return serverContext
}

func initRoutes(serverContext *ServerContext) {
	router := serverContext.Router
	blockchain := serverContext.Blockchain

	router.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	router.GET("/blockchain/chain", func(c echo.Context) error {
		chain := blockchain.GetChain()
		return c.JSON(http.StatusOK, chain)
	})
	router.GET("/blockchain/latest", func(c echo.Context) error {
		block := blockchain.GetLatestBlock()
		return c.JSON(http.StatusOK, block)
	})
}
