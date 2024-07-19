package blockchain

import (
	"errors"
	"strconv"
)

type Card struct {
	Value string
	Suit  string
}

type Transaction struct {
	Id     string
	CardP1 Card
	CardP2 Card
}

var cardValues = make(map[string]int)

func init() {
	for i := 1; i <= 10; i++ {
		cardValues[strconv.Itoa(i)] = i
	}

	cardValues["J"] = 11
	cardValues["Q"] = 12
	cardValues["K"] = 13
	cardValues["A"] = 14
}

// Return the winner, whether it was a tie, and the winning card
func (t Transaction) CalculateWinner() (string, bool, Card) {
	if t.CardP1.Value == t.CardP2.Value {
		return "", true, t.CardP1
	}

	if cardValues[t.CardP1.Value] > cardValues[t.CardP2.Value] {
		return "Player 1", false, t.CardP1
	}

	return "Player 2", false, t.CardP2
}

func (bc *Blockchain) AddTransaction(t Transaction) int {
	bc.PendingData = append(bc.PendingData, t)
	return len(bc.PendingData)
}

// In an actual application, I'd use a read-only database to store the blocks
// and transactions, but for timesake just reading through blockchain iteratively
func (bc *Blockchain) GetTransaction(txId string) (Transaction, error) {
	// search pending transactions first
	for _, tx := range bc.PendingData {
		if tx.Id == txId {
			return tx, nil
		}
	}

	// search blockchain transactions
	for _, block := range bc.Chain {
		for _, tx := range block.Data {
			if tx.Id == txId {
				return tx, nil
			}
		}
	}

	return Transaction{}, errors.New("Transaction not found")
}
