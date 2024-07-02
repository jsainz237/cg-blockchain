package server

import (
	bc "mtgbc/blockchain"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AddTransactionResponse struct {
	PendingTransactions int
	Transaction         bc.Transaction
}

var addTransactionHandler = func(c echo.Context) error {
	tx := new(bc.Transaction)
	tx.Id = uuid.New().String()

	if err := c.Bind(tx); err != nil {
		return c.String(http.StatusBadRequest, "Invalid transaction")
	}

	numPending := bc.MTGChain.AddTransaction(*tx)
	return c.JSON(http.StatusOK, AddTransactionResponse{numPending, *tx})
}

var getTransactionHandler = func(c echo.Context) error {
	txId := c.Param("transactionId")
	tx, err := bc.MTGChain.GetTransaction(txId)

	if err != nil {
		return c.String(http.StatusNotFound, "Transaction not found")
	}

	return c.JSON(http.StatusOK, tx)
}

var getWinnerHandler = func(c echo.Context) error {
	txId := c.Param("transactionId")
	tx, err := bc.MTGChain.GetTransaction(txId)

	if err != nil {
		return c.String(http.StatusNotFound, "Transaction not found")
	}

	winner, tie, winningCard := tx.CalculateWinner()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"winner":      winner,
		"tie":         tie,
		"winningCard": winningCard,
	})
}
