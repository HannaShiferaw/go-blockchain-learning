
package main

import "time"

func createGenesisBlock() Block {
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}

	genesis.Hash = calculateHash(genesis)
	return genesis
}
