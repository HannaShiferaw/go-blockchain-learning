package main

import "time"

var Blockchain []Block

func generateBlock(prevBlock Block, data string) Block {
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}

	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}
