package blockchain

import "fmt"
import "go-blockchain/models"

// Check if blockchain is valid
func IsBlockchainValid(chain []models.Block) bool {
    for i := 1; i < len(chain); i++ {
        currentBlock := chain[i]
        prevBlock := chain[i-1]

        // Check hash
        if currentBlock.Hash != CalculateHash(currentBlock) {
            fmt.Println("Invalid hash at index", currentBlock.Index)
            return false
        }

        // Check prev hash
        if currentBlock.PrevHash != prevBlock.Hash {
            fmt.Println("Invalid PrevHash at index", currentBlock.Index)
            return false
        }
    }
    return true
}
