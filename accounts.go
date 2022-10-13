package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func (xen *XenClaimer) fundAccounts() {
	nonce, err := xen.getNonce()
	if err != nil {
		fmt.Println(fmt.Sprintf("\n::ERROR error: %s", err.Error()))
		return
	}
	nonce = nonce - 1
	concurrentGoroutines := make(chan struct{}, xen.concurrencyFunding)
	convertTransferAmount := big.NewInt(int64(xen.toFundEachWallet * math.Pow10(18)))
	for _, a := range xen.accounts {
		if !a.Funded && !a.Claimed {
			wg.Add(1)
			go func(a *Account) {
				defer wg.Done()
				concurrentGoroutines <- struct{}{}
				atomic.AddUint64(&nonce, 1)
				hash, err := xen.TransferEth(xen.fundingKey, a.Address, convertTransferAmount, true, nonce)
				if err != nil {
					atomic.AddInt64(&xen.tasks.notCompleted, 1)
					fmt.Println(fmt.Sprintf("\n::ERROR %s error: %s", a.Address, err.Error()))
					return
				}
				fmt.Println(fmt.Sprintf("\n::INFO %s funding request was submitted", a.Address))
				for {
					status, err := xen.checkTxHashStatus(hash)
					if status == 0 {
						atomic.AddInt64(&xen.tasks.notCompleted, 1)
						fmt.Println(fmt.Sprintf("\n::ERROR %s error: %s", a.Address, err.Error()))
						break
					}
					if status == 1 {
						a.Funded = true
						mutex.Lock()
						xen.saveAccountsState()
						mutex.Unlock()
						fmt.Println(fmt.Sprintf("\n::INFO %s funding request was completed", a.Address))
						break
					}
					time.Sleep(2 * time.Second)
				}
				<-concurrentGoroutines
			}(a)

		}
	}
	wg.Wait()
	// Save account list state
	xen.saveAccountsState()
}

// Create N accounts and save them in wallets.json file
func (x *XenClaimer) generateAccounts() {
	//Store wallet details in file
	f, err := os.OpenFile(x.walletListPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < x.accountsToCreate; i++ {
		key, _ := crypto.GenerateKey()
		address := crypto.PubkeyToAddress(key.PublicKey).Hex()
		privateKey := hex.EncodeToString(key.D.Bytes())

		x.accounts = append(x.accounts, &Account{
			PrivateKey:      privateKey,
			Address:         address,
			Funded:          false,
			Claimed:         false,
			RewardClaimed:   false,
			TokenTransfered: false,
		})
	}

	fileData, _ := json.Marshal(x.accounts)

	if _, err := f.Write(fileData); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
