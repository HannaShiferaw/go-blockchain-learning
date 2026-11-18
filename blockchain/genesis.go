package blockchain

import (
	"time"

	"go-blockchain/models"
)

// createGenesisBlock makes the first block of the chain: index 0,
// no previous hash, no transactions. It is mined like any other
// block, and it is only ever created once, by New.
func (b *Blockchain) createGenesisBlock() models.Block {
	genesis := models.Block{
		Index:        0,
		Timestamp:    time.Now().UTC(),
		Transactions: nil,
		PrevHash:     "",
	}
	genesis = MineBlock(genesis, b.Difficulty)
	b.Blocks = append(b.Blocks, genesis)
	return genesis
}

// CreateGenesisBlock is kept for backwards compatibility with the
// first version of this project. New code should use New instead.
func CreateGenesisBlock() models.Block {
	chain := New(2)
	return chain.Blocks[0]
}
