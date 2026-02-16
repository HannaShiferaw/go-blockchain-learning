package main

import (
    "fmt"
    "go-blockchain/blockchain"
)

func main() {
    blockchain.CreateGenesisBlock()

    block1 := blockchain.GenerateBlock(blockchain.Blockchain[len(blockchain.Blockchain)-1], "Tx: A pays B 10")
    blockchain.Blockchain = append(blockchain.Blockchain, block1)

    block2 := blockchain.GenerateBlock(blockchain.Blockchain[len(blockchain.Blockchain)-1], "Tx: B pays C 5")
    blockchain.Blockchain = append(blockchain.Blockchain, block2)

    for _, block := range blockchain.Blockchain {
        fmt.Println("Index:", block.Index)
        fmt.Println("Data:", block.Data)
        fmt.Println("PrevHash:", block.PrevHash)
        fmt.Println("Hash:", block.Hash)
        fmt.Println("-------------------")
    }
}
