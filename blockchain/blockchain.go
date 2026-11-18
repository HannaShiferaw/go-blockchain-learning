// Package blockchain is a minimal, in-memory blockchain for learning.
// A chain is a Blockchain value, not a global list: it holds its
// blocks, its mining difficulty, and the rules for adding blocks.
package blockchain

import (
	"fmt"
	"time"

	"go-blockchain/models"
)

// Blockchain is a chain of blocks. Difficulty is how many leading
// zeros a block hash must have; 2 is quick to mine on a laptop.
type Blockchain struct {
	Blocks     []models.Block `json:"blocks"`
	Difficulty int            `json:"difficulty"`
}

// New creates a chain with a mined genesis block.
func New(difficulty int) *Blockchain {
	if difficulty < 0 {
		difficulty = 0
	}
	chain := &Blockchain{Difficulty: difficulty}
	chain.createGenesisBlock()
	return chain
}

// LastBlock returns the newest block, or false when the chain is empty.
func (b *Blockchain) LastBlock() (models.Block, bool) {
	if len(b.Blocks) == 0 {
		return models.Block{}, false
	}
	return b.Blocks[len(b.Blocks)-1], true
}

// Balances works out everyone's balance from all transactions so far.
// "SYSTEM" is not a real account, so it is left out.
func (b *Blockchain) Balances() map[string]float64 {
	balances := make(map[string]float64)
	for _, block := range b.Blocks {
		for _, tx := range block.Transactions {
			if tx.From != "" && tx.From != "SYSTEM" {
				balances[tx.From] -= tx.Amount
			}
			if tx.To != "" {
				balances[tx.To] += tx.Amount
			}
		}
	}
	return balances
}

// AddBlock appends a block with the given transactions, after checking
// them: amounts must be positive, and nobody may spend more than
// their balance. It mines the block first, then appends it.
func (b *Blockchain) AddBlock(txs ...models.Transaction) (models.Block, error) {
	balances := b.Balances()
	for _, tx := range txs {
		if tx.Amount <= 0 {
			return models.Block{}, fmt.Errorf("invalid amount %g in %s", tx.Amount, tx)
		}
		if tx.From != "" && tx.From != "SYSTEM" && balances[tx.From] < tx.Amount {
			return models.Block{}, fmt.Errorf("%s cannot pay %g: balance is %g", tx.From, tx.Amount, balances[tx.From])
		}
		if tx.From != "" && tx.From != "SYSTEM" {
			balances[tx.From] -= tx.Amount
		}
		if tx.To != "" {
			balances[tx.To] += tx.Amount
		}
	}

	prev, ok := b.LastBlock()
	if !ok {
		return models.Block{}, fmt.Errorf("chain has no genesis block")
	}
	block := models.Block{
		Index:        prev.Index + 1,
		Timestamp:    time.Now().UTC(),
		Transactions: txs,
		PrevHash:     prev.Hash,
	}
	block = MineBlock(block, b.Difficulty)
	b.Blocks = append(b.Blocks, block)
	return block, nil
}

// MineBlock finds a Nonce that makes the block hash start with
// Difficulty zeros, and returns the block with that hash.
func MineBlock(block models.Block, difficulty int) models.Block {
	for {
		block.Hash = CalculateHash(block)
		if hasValidProof(block.Hash, difficulty) {
			return block
		}
		block.Nonce++
	}
}

// GenerateBlock builds and mines one block after prevBlock.
// It is kept for the learning demo; new code should use AddBlock.
func GenerateBlock(prevBlock models.Block, txs ...models.Transaction) models.Block {
	block := models.Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().UTC(),
		Transactions: txs,
		PrevHash:     prevBlock.Hash,
	}
	return MineBlock(block, 2)
}
