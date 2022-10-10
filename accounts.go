package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func (xen *XenClaimer) fundAccounts() {
	convertTransferAmount := big.NewInt(int64(xen.toFundEachWallet * math.Pow10(18)))
	for _, a := range xen.accounts {
		if !a.Funded && !a.Claimed {
			hash, err := xen.TransferEth(xen.fundingKey, a.Address, convertTransferAmount)
			if err != nil {
				fmt.Println(fmt.Sprintf("\n::ERROR %s error: %s", a.Address, err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf("\n::INFO %s funding request was submitted", a.Address))
			for {
				status, err := xen.checkTxHashStatus(hash)
				if status == 0 {
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

		}
	}
	// Save account list state
	xen.saveAccountsState()
}

// Create accounts
func (x *XenClaimer) generateAccounts() {
	//Store wallet details in file
	f, err := os.OpenFile("wallets.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < x.accountsToCreate; i++ {
		key, _ := crypto.GenerateKey()
		address := crypto.PubkeyToAddress(key.PublicKey).Hex()
		privateKey := hex.EncodeToString(key.D.Bytes())

		x.accounts = append(x.accounts, &Account{
			PrivateKey: privateKey,
			Address:    address,
			Funded:     false,
			Claimed:    false,
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
