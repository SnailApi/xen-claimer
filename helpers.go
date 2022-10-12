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

	xen_abi "github.com/drawrowfly/xen-claimer/xen_abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
	Funded     bool   `json:"funded"`
	Claimed    bool   `json:"claimed"`
}

type XenClaimer struct {
	_XEN               *xen_abi.Store
	chainId            int
	stakeDays          int
	fundingKey         string
	toFundEachWallet   float64
	client             *ethclient.Client
	accountsToCreate   int
	concurrency        int
	concurrencyFunding int
	accounts           []*Account
	walletListPath     string
}

type ChainData struct {
	RpcUrl      string
	ChainId     int
	XenContract string
}

// Ger RPC url, chain id and contract address
func getRpcUrl(chain string) ChainData {
	if chain == "bsc" {
		return ChainData{
			RpcUrl:      os.Getenv("BSC_RPC"),
			ChainId:     56,
			XenContract: "0x2AB0e9e4eE70FFf1fB9D67031E44F6410170d00e",
		}
	}
	return ChainData{
		RpcUrl:      os.Getenv("ETH_RPC"),
		ChainId:     1,
		XenContract: "0x06450dee7fd2fb8e39061434babcfc05599a6fb8",
	}
}

// Get address from private key
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

// Preload existing accounts from file set by using -wl option
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

// Save account state in file set by using -wl option
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

// Withdraw funds from all accounts back to funding key
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

				hash, err := xen.TransferEth(a.PrivateKey, address.String(), big.NewInt(0), false, uint64(0))
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
