# Go Blockchain Learning

A minimal blockchain in Go, for learning how one works. No outside
dependencies — just the standard library.

Each block holds real payments (`from` pays `to` an amount), a hash
that covers every field in the block, and a proof of work: the hash
must start with a number of zeros. Change any payment, and the hash
no longer matches, so validation fails.

## Run it

You need Go 1.25 or newer.

```sh
go run . demo
```

The demo adds three payments, prints the chain, says it is valid,
then tampers with a copy and shows that the tampered chain is
invalid. The chain is saved to `blockchain.json`.

Other commands:

```sh
go run . add A B 10    # A pays B 10 (A must have the money)
go run . list          # print the saved chain
go run . validate      # check the saved chain is valid
```

To start over, delete `blockchain.json`.

## How it works

- `models/block.go` — `Transaction` and `Block`
- `blockchain/hash.go` — SHA-256 hashing; every field, including
  the nonce, is hashed, with `|` separators
- `blockchain/blockchain.go` — the `Blockchain` type, mining, adding
  blocks, and balance checks (nobody can spend more than they have)
- `blockchain/validation.go` — checks the genesis block, the index
  order, every hash, every link, and the proof of work
- `blockchain/storage.go` — save and load the chain as JSON
- `main.go` — the `demo`, `add`, `list` and `validate` commands

## Tests

```sh
go test ./...
go vet ./...
```

The tests cover hashing (including the nonce), the genesis block,
adding payments, overspending, tamper detection, proof of work,
and saving/loading.
