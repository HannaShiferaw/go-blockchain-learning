// Package models holds the data types of the blockchain:
// a transaction and a block that groups transactions together.
package models

import (
	"fmt"
	"time"
)

// Transaction is one payment: From pays To an Amount.
// From is "SYSTEM" for money created out of nothing, for example
// the first payment in the demo.
type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// String renders a transaction the way the demo prints it,
// for example "A pays B 10".
func (tx Transaction) String() string {
	return fmt.Sprintf("%s pays %s %g", tx.From, tx.To, tx.Amount)
}

// Block is one link in the chain. Its Hash covers every other field,
// including Nonce, so changing any of them breaks the chain.
type Block struct {
	Index        int           `json:"index"`
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prev_hash"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"`
}
