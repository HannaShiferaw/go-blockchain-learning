package blockchain

import (
	"encoding/json"
	"fmt"
	"os"
)

// Save writes the chain to a JSON file, so it survives restarts.
func (b *Blockchain) Save(path string) error {
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Errorf("could not encode chain: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("could not write %s: %w", path, err)
	}
	return nil
}

// Load reads a chain back from a JSON file and checks that it is valid.
func Load(path string) (*Blockchain, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", path, err)
	}
	var chain Blockchain
	if err := json.Unmarshal(data, &chain); err != nil {
		return nil, fmt.Errorf("could not decode %s: %w", path, err)
	}
	if err := chain.Validate(); err != nil {
		return nil, fmt.Errorf("saved chain in %s is invalid: %w", path, err)
	}
	return &chain, nil
}
