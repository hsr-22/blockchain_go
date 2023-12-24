package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const subsidy = 10

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded) // The gob package implements an efficient binary format encoding and decoding scheme
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

type TXInput struct { // An input references an output from a previous transaction
	Txid      []byte
	Vout      int
	ScriptSig string // The ScriptSig is the script provided by the spender of the previous output, combined with the ScriptPubKey of the output they are trying to spend
}

type TXOutput struct { // An output contains the value to be transferred, and specifies the conditions that must be fulfilled for those satoshis to be further spent
	Value        int
	StrictPubKey string
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}
