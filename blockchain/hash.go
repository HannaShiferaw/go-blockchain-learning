package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"go-blockchain/models"
)

// CalculateHash returns the SHA-256 hash of a block. Every field is
// included, separated by "|", so no two different blocks can produce
// the same input string by accident, and Nonce is really hashed.
func CalculateHash(block models.Block) string {
	txs := make([]string, len(block.Transactions))
	for i, tx := range block.Transactions {
		txs[i] = tx.String()
	}
	record := fmt.Sprintf("%d|%s|%s|%s|%d",
		block.Index,
		block.Timestamp.Format(time.RFC3339Nano),
		strings.Join(txs, ","),
		block.PrevHash,
		block.Nonce,
	)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// hasValidProof reports whether hash starts with Difficulty zeros.
func hasValidProof(hash string, difficulty int) bool {
	if difficulty <= 0 {
		return true
	}
	return strings.HasPrefix(hash, strings.Repeat("0", difficulty))
}
