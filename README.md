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

## Options

```
-a          Number of accounts to create (default 2)
-c          Concurency to claim XEN (default 10)
-cw         Concurency to fund wallets (default 10)
-d          Set number of days to stakes (default 100)
-f          To fund each wallet (default 0.01)
-fund       Create, fund accounts and initiate claim
-withdraw   Withdraw funds from all acounts back to funding key
-wl         Path to wallets list (default "wallets.json")
-rpc        HTTP RPC url
-fw         Private key for wallet that will be used to fund newly created wallets
```

If you will set RPC for binance chain then the claimer will work with binance chain

## Example

Create 100 wallets, fund each with 0.01 ETH from FUNDING_KEY_WALLET with concurrency set to 5, claim XEN with the duration 32 days with set concurency to 10

```
xen-claimer -fund -a 100 -c 10 -cw 5 -d 2 -f 0.01 -rpc https://ETH.RPC.COM -fw PRIVATE_KEY
```
