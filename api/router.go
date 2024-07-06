package api

import (
	"net/http"
	"os"

	_ "mtgbc/env"

	"github.com/labstack/echo/v4"
)

var Router *echo.Echo

func init() {
	Router = echo.New()
	Router.HideBanner = true
}

func Startserver() {
	port := os.Getenv("PORT")

	initRoutes()
	Router.Logger.Fatal(Router.Start(":" + port))
}

func initRoutes() {
	Router.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	Router.GET("/blockchain/chain", getBlockchainHandler)
	Router.GET("/blockchain/latest", getLatestBlockHandler)

	Router.POST("/transaction", addTransactionHandler)
	Router.GET("/transaction/:transactionId", getTransactionHandler)
	Router.GET("/transaction/:transactionId/winner", getWinnerHandler)
}
