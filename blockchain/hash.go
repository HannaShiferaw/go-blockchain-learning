package main
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)
func calculateHash(block Block) string {
	record := fmt.Sprintf(
		"%d%s%s%s",
		block.Index,
		block.Timestamp,
		block.Data,
		block.PrevHash,
	)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}




