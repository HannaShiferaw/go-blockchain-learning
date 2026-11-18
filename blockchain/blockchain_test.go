package blockchain

import (
	"path/filepath"
	"strings"
	"testing"

	"go-blockchain/models"
)

func TestCalculateHashIncludesNonce(t *testing.T) {
	block := models.Block{Index: 1, Transactions: []models.Transaction{{From: "A", To: "B", Amount: 10}}}
	first := CalculateHash(block)
	block.Nonce = 1
	if CalculateHash(block) == first {
		t.Error("changing Nonce did not change the hash")
	}
}

func TestCalculateHashIsStable(t *testing.T) {
	block := models.Block{Index: 1, Transactions: []models.Transaction{{From: "A", To: "B", Amount: 10}}}
	if CalculateHash(block) != CalculateHash(block) {
		t.Error("hashing the same block twice gave different results")
	}
}

func TestNewChainIsValid(t *testing.T) {
	chain := New(1)
	if err := chain.Validate(); err != nil {
		t.Fatalf("fresh chain is invalid: %v", err)
	}
	if len(chain.Blocks) != 1 || chain.Blocks[0].Index != 0 {
		t.Fatalf("fresh chain should have only a genesis block, got %+v", chain.Blocks)
	}
}

func TestAddBlockAndBalances(t *testing.T) {
	chain := New(1)
	if _, err := chain.AddBlock(models.Transaction{From: "SYSTEM", To: "A", Amount: 100}); err != nil {
		t.Fatalf("adding SYSTEM payment: %v", err)
	}
	if _, err := chain.AddBlock(models.Transaction{From: "A", To: "B", Amount: 10}); err != nil {
		t.Fatalf("adding payment: %v", err)
	}
	balances := chain.Balances()
	if balances["A"] != 90 || balances["B"] != 10 {
		t.Errorf("balances = %v, want A=90 B=10", balances)
	}
	if err := chain.Validate(); err != nil {
		t.Errorf("chain is invalid after adding blocks: %v", err)
	}
}

func TestAddBlockRejectsOverspending(t *testing.T) {
	chain := New(1)
	if _, err := chain.AddBlock(models.Transaction{From: "A", To: "B", Amount: 10}); err == nil {
		t.Error("A paying 10 with a balance of 0 should be rejected")
	}
	if len(chain.Blocks) != 1 {
		t.Errorf("rejected block was still added: %d blocks", len(chain.Blocks))
	}
}

func TestAddBlockRejectsBadAmount(t *testing.T) {
	chain := New(1)
	if _, err := chain.AddBlock(models.Transaction{From: "SYSTEM", To: "A", Amount: 0}); err == nil {
		t.Error("an amount of 0 should be rejected")
	}
}

func TestTamperingIsDetected(t *testing.T) {
	chain := New(1)
	if _, err := chain.AddBlock(models.Transaction{From: "SYSTEM", To: "A", Amount: 100}); err != nil {
		t.Fatalf("adding payment: %v", err)
	}
	chain.Blocks[1].Transactions[0].Amount = 10000
	if err := chain.Validate(); err == nil {
		t.Error("tampered chain passed validation")
	}
}

func TestGenesisTamperingIsDetected(t *testing.T) {
	chain := New(1)
	chain.Blocks[0].PrevHash = "forged"
	if err := chain.Validate(); err == nil {
		t.Error("tampered genesis block passed validation")
	}
}

func TestMinedBlocksHaveProofOfWork(t *testing.T) {
	chain := New(2)
	if _, err := chain.AddBlock(models.Transaction{From: "SYSTEM", To: "A", Amount: 5}); err != nil {
		t.Fatalf("adding payment: %v", err)
	}
	for _, block := range chain.Blocks {
		if !strings.HasPrefix(block.Hash, "00") {
			t.Errorf("block %d hash %s does not start with 00", block.Index, block.Hash)
		}
	}
}

func TestSaveAndLoad(t *testing.T) {
	chain := New(1)
	if _, err := chain.AddBlock(models.Transaction{From: "SYSTEM", To: "A", Amount: 100}); err != nil {
		t.Fatalf("adding payment: %v", err)
	}
	path := filepath.Join(t.TempDir(), "chain.json")
	if err := chain.Save(path); err != nil {
		t.Fatalf("saving: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("loading: %v", err)
	}
	if len(loaded.Blocks) != len(chain.Blocks) {
		t.Fatalf("loaded %d blocks, want %d", len(loaded.Blocks), len(chain.Blocks))
	}
	if loaded.Blocks[1].Hash != chain.Blocks[1].Hash {
		t.Error("loaded chain hash does not match the saved chain")
	}
}
