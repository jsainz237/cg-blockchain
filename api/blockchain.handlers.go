package api

import (
	bc "mtgbc/blockchain"
	"net/http"

	"github.com/labstack/echo/v4"
)

var getBlockchainHandler = func(c echo.Context) error {
	return c.JSON(http.StatusOK, bc.MTGChain.GetChain())
}

var getLatestBlockHandler = func(c echo.Context) error {
	return c.JSON(http.StatusOK, bc.MTGChain.GetLatestBlock())
}
