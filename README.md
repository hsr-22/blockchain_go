# Blockchain Implementation in Go

This project is a simple implementation of a blockchain in Go, inspired by the tutorial series by Jeiwan: [Building Blockchain in Go](https://jeiwan.net/posts/building-blockchain-in-go-part-1/).

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Blockchain](#blockchain)
- [Proof of Work](#proof-of-work)
- [Transactions](#transactions)
- [Wallets](#wallets)
- [CLI](#cli)
- [Usage](#usage)
- [License](#license)

## Introduction

This project demonstrates the core concepts of blockchain technology, including blocks, transactions, proof of work, and wallets. It uses BoltDB for persistent storage and implements basic functionalities such as creating a blockchain, mining new blocks, and handling transactions.

## Project Structure
```

.gitignore
.vscode/
launch.json
settings.json

address.md

base58.go

block.go

blockchain.go

cli.go

database.md

go.mod

go.sum

hashing.md

keys.md

LICENSE

main.go

pow.go

transaction.go

transaction.md

tx_input.go

tx_output.go

utils.go

wallet.go

wallets.go

```

## Dependencies

The project relies on the following dependencies:

- [BoltDB](https://github.com/boltdb/bolt): A key/value store for Go.
- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto): Additional cryptographic packages for Go.

These dependencies are specified in the [go.mod](go.mod) file.

## Blockchain

The core of the blockchain is implemented in [blockchain.go](blockchain.go). The `Blockchain` struct manages the chain of blocks and interactions with the database.

```go
// Blockchain implements interactions with a DB
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}
````

Blocks are created and added to the blockchain using the `MineBlock` method, which includes transaction validation.

```go
func (bc *Blockchain) MineBlock(transactions []*Transaction) {
    // Implementation details...
}
```

## Proof of Work

The proof of work algorithm is implemented in

pow.go

. It ensures that blocks are mined with a certain level of difficulty.

```go
// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}
```

The `Run` method performs the proof of work by finding a valid nonce.

```go
func (pow *ProofOfWork) Run() (int, []byte) {
    // Implementation details...
}
```

## Transactions

Transactions are the core of the blockchain's functionality. They are implemented in

transaction.go

and consist of inputs and outputs.

```go
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}
```

Coinbase transactions, which reward miners, are created using the

NewCoinbaseTX

function.

```go
func NewCoinbaseTX(to, data string) *Transaction {
    // Implementation details...
}
```

## Wallets

Wallets manage private and public keys and generate addresses. They are implemented in

wallet.go

and

wallets.go

.

```go
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}
```

The

NewWallet

function creates a new wallet with a unique key pair.

```go
func NewWallet() *Wallet {
    // Implementation details...
}
```

## CLI

The command-line interface (CLI) is implemented in

cli.go

. It provides commands to interact with the blockchain, such as creating a blockchain, mining blocks, and checking balances.

```go
type CLI struct{}
```

The `Run` method parses command-line arguments and executes the corresponding commands.

```go
func (cli *CLI) Run() {
    // Implementation details...
}
```

## Usage

To build and run the project, use the following commands:

```sh
go build -o blockchain
./blockchain
```

The CLI supports the following commands:

- `createblockchain -address ADDRESS`: Create a blockchain and send the genesis block reward to `ADDRESS`.
- `createwallet`: Generate a new key-pair and save it into the wallet file.
- `getbalance -address ADDRESS`: Get the balance of `ADDRESS`.
- `listaddresses`: List all addresses from the wallet file.
- `printchain`: Print all the blocks of the blockchain.
- `send -from FROM -to TO -amount AMOUNT`: Send `AMOUNT` of coins from `FROM` address to `TO`.

## License

This project is licensed under the MIT License. See the

LICENSE

file for details.
