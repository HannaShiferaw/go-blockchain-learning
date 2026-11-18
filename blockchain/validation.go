package blockchain

import (
	"fmt"
)

// Validate checks the whole chain and returns the first problem it
// finds, or nil when the chain is valid. It checks the genesis block
// too — the first version only checked blocks after it — plus the
// index order, every hash, every PrevHash link, and the proof of work.
func (b *Blockchain) Validate() error {
	if len(b.Blocks) == 0 {
		return fmt.Errorf("chain is empty: no genesis block")
	}
	for i, block := range b.Blocks {
		if block.Index != i {
			return fmt.Errorf("invalid index at position %d: block says %d", i, block.Index)
		}
		if block.Hash != CalculateHash(block) {
			return fmt.Errorf("invalid hash at index %d", block.Index)
		}
		if !hasValidProof(block.Hash, b.Difficulty) {
			return fmt.Errorf("invalid proof of work at index %d", block.Index)
		}
		if i == 0 {
			if block.PrevHash != "" {
				return fmt.Errorf("genesis block must have an empty PrevHash")
			}
			continue
		}
		if block.PrevHash != b.Blocks[i-1].Hash {
			return fmt.Errorf("invalid PrevHash at index %d", block.Index)
		}
	}
	return nil
}

// IsBlockchainValid reports whether the chain is valid.
// It is kept for backwards compatibility; Validate gives the reason.
func (b *Blockchain) IsBlockchainValid() bool {
	return b.Validate() == nil
}
