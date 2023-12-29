# Transactions

Every new transaction must have at least one input and output.

## Transaction Outputs

Outputs are where "coins" are stored. Each output comes with an unlocking script, which determines the logic of unlocking the output.

> In Bitcoin, the value field stores the number of satoshis, not the number of BTC. A satoshi is a hundred millionth of a bitcoin (0.00000001 BTC), thus this is the smallest unit of currency in Bitcoin (like a cent).

## Transaction Inputs

Input references an output from a previous transaction and provides data (`ScriptSig`) that is used in the output's unlocking script to unlock it and use its value to create new outputs.

In Bitcoin, outputs come before inputs (*egg or the chicken?*)

## Coinbase Transaction

A coinbase transaction is a special type of transactions, which doesn’t require previously existing outputs. It creates outputs (i.e., “coins”) out of nowhere. The egg without a chicken. This is the reward miners get for mining new blocks.

# Full Lifecycle of a Transaction

1. In the beginning, there’s the genesis block that contains a coinbase transaction. There are no real inputs in coinbase transactions, so [signing](./address.md) is not necessary. The output of the coinbase transaction contains a hashed public key 
(RIPEMD16(SHA256(PubKey)) algorithms are used).

2. When one sends coins, a transaction is created. Inputs of the transaction will reference outputs from previous transaction(s). Every input will store a public key (not hashed) and a signature of the whole transaction.

3. Other nodes in the Bitcoin network that receive the transaction will verify it. Besides other things, they will check that: the hash of the public key in an input matches the hash of the referenced output (this ensures that the sender spends only coins belonging to them); the signature is correct (this ensures that the transaction is created by the real owner of the coins).

4. When a miner node is ready to mine a new block, it’ll put the transaction in a block and start mining it.

5. When the block is mined, every other node in the network receives a message saying the block is mined and adds the block to the blockchain.

6. After a block is added to the blockchain, the transaction is completed, its outputs can be referenced in new transactions.