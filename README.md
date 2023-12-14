# Filecoin inscription mint

## What's this?

Just like what BRC-20 project `ordi` did.
> { "p": "brc-20", "op": "deploy", "tick": "ordi", "max": "21000000", "lim": "1000" }

## How to use it?

In the current version, you need to change the variable in the `main.go` to use it.
> data:,{"p":"fil-20","op":"mint","tick":"fils","amt":"1000"}

I'll improve it just like using a CLI later.

## Public API

```text
Ankr: https://www.ankr.com/rpc/filecoin/  
Glif: https://api.node.glif.io/
ChainStack: https://chainstack.com/labs/#filecoin
```

> Local node: ws://127.0.0.1:1234/rpc/v0

## Filecoin Mainnet RPC URL List

https://chainlist.org/chain/314

