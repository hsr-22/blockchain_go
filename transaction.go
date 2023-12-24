package main

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
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
