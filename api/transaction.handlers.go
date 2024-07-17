package api

import (
	"fmt"
	bc "mtgbc/blockchain"
	network "mtgbc/network"
	"net/http"

	"github.com/google/uuid"
)

type TransactionHandlers struct{}

type AddTransactionArgs struct {
	Transaction bc.Transaction
}

type AddTransactionResponse struct {
	PendingTransactions int
	Transaction         bc.Transaction
}

func (th *TransactionHandlers) Add(
	r *http.Request,
	args *AddTransactionArgs,
	reply *AddTransactionResponse,
) error {
	tx := args.Transaction
	tx.Id = uuid.New().String()

	numPending := bc.MTGChain.AddTransaction(tx)

	for _, node := range network.MTGNetwork.ConnectionPool {
		_, _, err := CallRPC(node, "Transaction.Sync", &AddTransactionArgs{Transaction: tx})
		if err != nil {
			return fmt.Errorf("could not sync transaction")
		}
	}

	*reply = AddTransactionResponse{numPending, tx}
	return nil
}

func (th *TransactionHandlers) Sync(
	r *http.Request,
	args *AddTransactionArgs,
	reply *string,
) error {
	tx := args.Transaction
	bc.MTGChain.AddTransaction(tx)
	*reply = "Synced"
	return nil
}

type WinnerResponse struct {
	Winner      string
	Tie         bool
	WinningCard bc.Card
}

func (th *TransactionHandlers) Winner(
	r *http.Request,
	args *string,
	reply *WinnerResponse,
) error {
	txId := *args
	tx, err := bc.MTGChain.GetTransaction(txId)

	if err != nil {
		return fmt.Errorf("transaction not found")
	}

	winner, tie, winningCard := tx.CalculateWinner()
	*reply = WinnerResponse{winner, tie, winningCard}
	return nil
}
