package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func (xen *XenClaimer) initiateClaiming() {
	concurrentGoroutines := make(chan struct{}, xen.concurrency)
	for _, a := range xen.accounts {
		if !a.Claimed && a.Funded {
			wg.Add(1)
			go func(a *Account) {
				defer wg.Done()
				concurrentGoroutines <- struct{}{}
				hash, err := xen.claimer(a)
				if err != nil {
					fmt.Println(fmt.Sprintf("\n::ERROR CLAIM %s error: %s", a.Address, err.Error()))
					return
				}
				fmt.Println(fmt.Sprintf("\n::INFO CLAIM %s claim request was submitted", a.Address))
				for {
					status, err := xen.checkTxHashStatus(hash)
					if status == 0 {
						fmt.Println(fmt.Sprintf("\n::ERROR CLAIM %s error: %s", a.Address, err.Error()))
						break
					}
					if status == 1 {
						a.Claimed = true
						mutex.Lock()
						xen.saveAccountsState()
						mutex.Unlock()
						fmt.Println(fmt.Sprintf("\n::INFO CLAIM %s claim request was completed", a.Address))
						break
					}
					time.Sleep(1 * time.Second)
				}
				<-concurrentGoroutines
			}(a)
		}
	}
	wg.Wait()
	xen.saveAccountsState()
}

// Xen Claimer
func (x *XenClaimer) claimer(account *Account) (string, error) {

	privateKey, err := crypto.HexToECDSA(account.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("BAD 1 Wallet: %s error: %s \n", x.fundingKey, err.Error())
	}

	gasPrice, err := x.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	gasPrice = new(big.Int).Add(gasPrice, new(big.Int).Div(gasPrice, big.NewInt(5)))

	// Create an authorized transactor and spend 1 unicorn
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(x.chainId)))
	if err != nil {
		return "", err
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = 200000

	term := big.NewInt(int64(x.stakeDays))
	tx, err := x._XEN.ClaimRank(auth, term)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}
