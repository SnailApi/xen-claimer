# xen-claimer

Claimer for https://xen.network/mainnet

Following tool creates n amount of accounts, funds them with ETH by using FUNDING_WALLET and claims XEN. At the end of execution you will have wallet.json file with the following format:

```json
[
  {
    "private_key": "private_key_data",
    "address": "wallet_address",
    "funded": true, // true if wallet was already funded
    "claimed": true // true if wallet has claimed XEN
  }
]
```

## Setup

Before usig it, you need to have 2 environment variable set, for example by creating .env file in the same folder where you will execute the tool from with following data

```
ETH_RPC=ETH_RPC_URL
BSC_RPC=BSC_RPC_URL
FUNDING_WALLET=FUNDING_WALLET_PRIVATE_KEY
```

## Options

```
-a          Number of accounts to create (default 2)
-c          Concurency to claim XEN (default 10)
-cw         Concurency to fund wallets (default 10)
-chain      Select chain to work with: eth or bsc (default "eth")
-d          Set number of days to stakes (default 100)
-f          To fund each wallet (default 0.01)
-fund       Create, fund accounts and initiate claim
-withdraw   Withdraw funds from all acounts back to funding key
-wl         Path to wallets list (default "wallets.json")
```

## Example

Create 100 wallets, fund each with 0.01 ETH from FUNDING_KEY_WALLET with concurrency set to 5, claim XEN with the duration 32 days with set concurency to 10

```
xen-claimer -fund -a 100 -c 10 -cw 5 -d 2 -f 0.01 -chain eth
```
