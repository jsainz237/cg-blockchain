package server

import (
	bc "mtgbc/blockchain"

	"github.com/labstack/echo/v4"
)

type ServerContext struct {
	Blockchain *bc.Blockchain
	Router     *echo.Echo
}

type ServerContextHandler = func(e echo.Context) error
