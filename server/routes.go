package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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

	router.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	router.GET("/blockchain/chain", GetBlockchainHandler(*serverContext))
	router.GET("/blockchain/latest", GetLatestBlockHandler(*serverContext))
}
