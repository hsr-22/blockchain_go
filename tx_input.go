package main

import "bytes"

type TXInput struct { // An input references an output from a previous transaction
	Txid      []byte
	Vout      int // The index of the output being referenced
	Signature []byte
	PubKey    []byte
}

// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Equal(lockingHash, pubKeyHash)
}
