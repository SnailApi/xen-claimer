package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (xen *XenClaimer) getAddressFromPKey(pKey string) (common.Address, error) {
	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("Bad private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress, nil
}

func (x *XenClaimer) loadExistingAccounts() {
	f, err := os.ReadFile(x.walletListPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(f, &x.accounts)
	if err != nil {
		log.Fatal(err)
	}
}

func (x *XenClaimer) saveAccountsState() {
	content, err := json.Marshal(x.accounts)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(x.walletListPath, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (xen *XenClaimer) withdrawToFundingKey() {
	address, err := xen.getAddressFromPKey(xen.fundingKey)
	if err != nil {
		log.Fatalln(err)
	}

	concurrentGoroutines := make(chan struct{}, xen.concurrency)
	for _, a := range xen.accounts {
		if a.Funded {
			wg.Add(1)
			go func(a *Account) {
				defer wg.Done()
				concurrentGoroutines <- struct{}{}

				hash, err := xen.TransferEth(a.PrivateKey, address.String(), big.NewInt(0))
				if err != nil {
					fmt.Println(fmt.Sprintf("\n::ERROR %s error: %s", a.Address, err.Error()))
					return
				}
				fmt.Println(fmt.Sprintf("\n::INFO %s withdraw request submitted", a.Address))
				for {
					status, err := xen.checkTxHashStatus(hash)
					if status == 0 {
						fmt.Println(fmt.Sprintf("\n::ERROR %s error: %s", a.Address, err.Error()))
						break
					}
					if status == 1 {
						a.Funded = false
						mutex.Lock()
						xen.saveAccountsState()
						mutex.Unlock()
						fmt.Println(fmt.Sprintf("\n::INFO %s withdraw request completed", a.Address))
						break
					}
					time.Sleep(2 * time.Second)
				}
				<-concurrentGoroutines
			}(a)
		}
	}
	wg.Wait()
}
