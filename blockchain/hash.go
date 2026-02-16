package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "go-blockchain/models" // <- import models for Block
)

// Exported function
func CalculateHash(block models.Block) string {
    record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
    hash := sha256.Sum256([]byte(record))
    return hex.EncodeToString(hash[:])
}
