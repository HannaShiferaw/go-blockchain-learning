// Command go-blockchain-learning is a minimal blockchain you can
// play with from the terminal. It creates a genesis block, adds
// payments, mines blocks, and checks that nobody tampered with it.
// The chain is saved to blockchain.json, so it survives restarts.
//
// Usage:
//
//	go run . demo                        # the guided demo (default)
//	go run . add <from> <to> <amount>    # add a payment
//	go run . list                        # print the chain
//	go run . validate                    # check the chain is valid
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"go-blockchain/blockchain"
	"go-blockchain/models"
)

const chainFile = "blockchain.json"

const difficulty = 2

// loadOrNew reads the saved chain, or starts a fresh one.
func loadOrNew() *blockchain.Blockchain {
	if chain, err := blockchain.Load(chainFile); err == nil {
		return chain
	}
	return blockchain.New(difficulty)
}

func save(chain *blockchain.Blockchain) {
	if err := chain.Save(chainFile); err != nil {
		fmt.Println("Warning: could not save the chain:", err)
	}
}

func printChain(chain *blockchain.Blockchain) {
	for _, block := range chain.Blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Timestamp:", block.Timestamp.Format("2006-01-02 15:04:05"))
		if len(block.Transactions) == 0 {
			fmt.Println("Transactions: (genesis block)")
		} else {
			for _, tx := range block.Transactions {
				fmt.Println("Transaction:", tx)
			}
		}
		fmt.Println("PrevHash:", block.PrevHash)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("Nonce:", block.Nonce)
		fmt.Println("-------------------")
	}
}

func reportValid(chain *blockchain.Blockchain) {
	if err := chain.Validate(); err != nil {
		fmt.Println("Blockchain is invalid ❌:", err)
	} else {
		fmt.Println("Blockchain is valid ✅")
	}
}

// demo reproduces the original example: two payments, a validity
// check, then a tampered copy that fails the check.
func demo(chain *blockchain.Blockchain) {
	if _, err := chain.AddBlock(models.Transaction{From: "SYSTEM", To: "A", Amount: 100}); err != nil {
		fmt.Println("Could not add the first payment:", err)
		return
	}
	if _, err := chain.AddBlock(models.Transaction{From: "A", To: "B", Amount: 10}); err != nil {
		fmt.Println("Could not add the second payment:", err)
		return
	}
	if _, err := chain.AddBlock(models.Transaction{From: "B", To: "C", Amount: 5}); err != nil {
		fmt.Println("Could not add the third payment:", err)
		return
	}
	save(chain)

	printChain(chain)
	reportValid(chain)

	fmt.Println("\nTampering with block 1...")
	tampered := *chain
	tampered.Blocks = append([]models.Block{}, chain.Blocks...)
	tampered.Blocks[1].Transactions = []models.Transaction{{From: "A", To: "B", Amount: 10000}}
	reportValid(&tampered)
}

func run(args []string) error {
	chain := loadOrNew()
	if len(args) == 0 || args[0] == "demo" {
		demo(chain)
		return nil
	}

	switch args[0] {
	case "list":
		printChain(chain)
		return nil
	case "validate":
		reportValid(chain)
		return nil
	case "add":
		if len(args) != 4 {
			return errors.New("usage: add <from> <to> <amount>")
		}
		amount, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			return fmt.Errorf("amount %q is not a number", args[3])
		}
		if _, err := chain.AddBlock(models.Transaction{From: args[1], To: args[2], Amount: amount}); err != nil {
			return err
		}
		save(chain)
		fmt.Println("Payment added and block mined.")
		printChain(chain)
		return nil
	default:
		return fmt.Errorf("unknown command %q: use demo, add, list or validate", args[0])
	}
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Usage: go run . [demo | add <from> <to> <amount> | list | validate]")
		os.Exit(1)
	}
}
