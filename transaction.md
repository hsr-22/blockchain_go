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