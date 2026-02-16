package blockchain

import (
    "time"
    "go-blockchain/models" // import Block struct
)

// Exported Blockchain variable
var Blockchain []models.Block

func GenerateBlock(prevBlock models.Block, data string) models.Block {
    newBlock := models.Block{
        Index:     prevBlock.Index + 1,
        Timestamp: time.Now().String(),
        Data:      data,
        PrevHash:  prevBlock.Hash,
    }
    newBlock.Hash = CalculateHash(newBlock) // same package, can call directly
    return newBlock
}
