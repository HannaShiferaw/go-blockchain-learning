package blockchain

import (
    "time"
    "go-blockchain/models"
)

func CreateGenesisBlock() models.Block {
    genesis := models.Block{
        Index:     0,
        Timestamp: time.Now().String(),
        Data:      "Genesis Block",
        PrevHash:  "",
    }
    genesis.Hash = CalculateHash(genesis) // same package
    Blockchain = append(Blockchain, genesis)
    return genesis
}
