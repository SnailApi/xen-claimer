package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Initiate XEN claiming process
func (xen *XenClaimer) initiateXenClaiming() {
	concurrentGoroutines := make(chan struct{}, xen.concurrency)
	for _, a := range xen.accounts {
		if !a.Claimed && a.Funded {
			wg.Add(1)
			go func(a *Account) {
				defer wg.Done()
				concurrentGoroutines <- struct{}{}
				hash, err := xen.xenClaimer(a)
				if err != nil {
					atomic.AddInt64(&xen.tasks.notCompleted, 1)
					fmt.Println(fmt.Sprintf("\n::ERROR CLAIM %s error: %s", a.Address, err.Error()))
					return
				}
				fmt.Println(fmt.Sprintf("\n::INFO CLAIM %s claim request was submitted", a.Address))
				for {
					status, err := xen.checkTxHashStatus(hash)
					if status == 0 {
						atomic.AddInt64(&xen.tasks.notCompleted, 1)
						fmt.Println(fmt.Sprintf("\n::ERROR CLAIM %s error: %s", a.Address, err.Error()))
						break
					}
					if status == 1 {
						a.Claimed = true
						mutex.Lock()
						xen.saveAccountsState()
						mutex.Unlock()
						atomic.AddInt64(&xen.tasks.completed, 1)
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

// XEN claimer
func (x *XenClaimer) xenClaimer(account *Account) (string, error) {

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
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(x.chain.ChainId)))
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

// Initiate XEN reward claiming process
func (xen *XenClaimer) initiateXenRewardClaiming() {
	concurrentGoroutines := make(chan struct{}, xen.concurrency)
	for _, a := range xen.accounts {
		if a.Claimed && !a.RewardClaimed {
			wg.Add(1)
			go func(a *Account) {
				defer wg.Done()
				concurrentGoroutines <- struct{}{}
				hash, err := xen.xenRewardClaimer(a)
				if err != nil {
					atomic.AddInt64(&xen.tasks.notCompleted, 1)
					fmt.Println(fmt.Sprintf("\n::ERROR REWARD CLAIM %s error: %s", a.Address, err.Error()))
					return
				}
				fmt.Println(fmt.Sprintf("\n::INFO REWARD CLAIM %s claim request was submitted", a.Address))
				for {
					status, err := xen.checkTxHashStatus(hash)
					if status == 0 {
						atomic.AddInt64(&xen.tasks.notCompleted, 1)
						fmt.Println(fmt.Sprintf("\n::ERROR REWARD CLAIM %s error: %s", a.Address, err.Error()))
						break
					}
					if status == 1 {
						a.RewardClaimed = true
						mutex.Lock()
						xen.saveAccountsState()
						mutex.Unlock()
						fmt.Println(fmt.Sprintf("\n::INFO REWARD CLAIM %s claim request was completed", a.Address))
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

// XEN reward claimer
func (x *XenClaimer) xenRewardClaimer(account *Account) (string, error) {

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
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(x.chain.ChainId)))
	if err != nil {
		return "", err
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(x.chain.GasLimitDefault)

	tx, err := x._XEN.ClaimMintReward(auth)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

// Get XEN claiming details
func (xen *XenClaimer) getXenWalletClaimingDetails() {
	concurrentGoroutines := make(chan struct{}, xen.concurrency)
	for _, a := range xen.accounts {
		if a.Claimed {
			wg.Add(1)
			go func(a *Account) {
				defer wg.Done()
				concurrentGoroutines <- struct{}{}
				details, err := xen.xenWalletClaimingDetails(a)
				if err != nil {
					atomic.AddInt64(&xen.tasks.notCompleted, 1)
					fmt.Println(fmt.Sprintf("\n::ERROR DETAILS %s error: %s", a.Address, err.Error()))
					return
				}
				atomic.AddInt64(&xen.tasks.completed, 1)
				fmt.Println(fmt.Sprintf("::INFO DETAILS: %s RANK: %d TERM: %d", details.User.String(), details.Rank.Int64(), details.Term.Int64()))
				<-concurrentGoroutines
			}(a)
		}
	}
	wg.Wait()
	xen.saveAccountsState()
}

// XEN address claiming details
func (xen *XenClaimer) xenWalletClaimingDetails(account *Account) (_XEN_UserMints, error) {

	address := common.HexToAddress(account.Address)

	details, err := xen._XEN.UserMints(nil, address)
	if err != nil {
		return _XEN_UserMints{}, err
	}

	return details, nil
}

// Get XEN address balance
func (xen *XenClaimer) getXenBalance() {
	concurrentGoroutines := make(chan struct{}, xen.concurrency)
	for _, a := range xen.accounts {
		if a.RewardClaimed {
			wg.Add(1)
			go func(a *Account) {
				defer wg.Done()
				concurrentGoroutines <- struct{}{}
				balance, err := xen.xenBalance(a)
				if err != nil {
					fmt.Println(fmt.Sprintf("\n::ERROR DETAILS %s error: %s", a.Address, err.Error()))
					return
				}
				humanBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
				fmt.Println(fmt.Sprintf("::INFO DETAILS: %s BALANCE: %f", a.Address, humanBalance))
				<-concurrentGoroutines
			}(a)
		}
	}
	wg.Wait()
	xen.saveAccountsState()
}

// XEN address balance
func (xen *XenClaimer) xenBalance(account *Account) (*big.Int, error) {

	address := common.HexToAddress(account.Address)

	balance, err := xen._XEN.BalanceOf(nil, address)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// Initiate XEN token transfer to provided wallet
func (xen *XenClaimer) initiateXenTransferToWallet() {
	concurrentGoroutines := make(chan struct{}, xen.concurrency)
	for _, a := range xen.accounts {
		if a.RewardClaimed && !a.TokenTransfered {
			wg.Add(1)
			go func(a *Account) {
				defer wg.Done()
				concurrentGoroutines <- struct{}{}
				hash, err, transferred := xen.xenTransferToWallet(a)
				if err != nil {
					atomic.AddInt64(&xen.tasks.notCompleted, 1)
					fmt.Println(fmt.Sprintf("\n::ERROR TRANSFER %s error: %s", a.Address, err.Error()))
					return
				}
				fmt.Println(fmt.Sprintf("\n::INFO TRANSFER %s transfer request was submitted", a.Address))
				for {
					status, err := xen.checkTxHashStatus(hash)
					if status == 0 {
						atomic.AddInt64(&xen.tasks.notCompleted, 1)
						fmt.Println(fmt.Sprintf("\n::ERROR TRANSFER %s error: %s", a.Address, err.Error()))
						break
					}
					if status == 1 {
						a.TokenTransfered = true
						mutex.Lock()
						xen.saveAccountsState()
						mutex.Unlock()
						atomic.AddInt64(&xen.tasks.completed, 1)
						atomic.AddUint64(&xen.tasks.totalTransfered, transferred)
						fmt.Println(fmt.Sprintf("\n::INFO TRANSFER %s transfer request was completed", a.Address))
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

// XEN token transfer
func (x *XenClaimer) xenTransferToWallet(account *Account) (string, error, uint64) {

	privateKey, err := crypto.HexToECDSA(account.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("BAD 1 Wallet: %s error: %s \n", x.fundingKey, err.Error()), 0
	}

	gasPrice, err := x.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err, 0
	}

	gasPrice = new(big.Int).Add(gasPrice, new(big.Int).Div(gasPrice, big.NewInt(5)))

	// Create an authorized transactor and spend 1 unicorn
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(x.chain.ChainId)))
	if err != nil {
		return "", err, 0
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(x.chain.GasLimitDefault)

	// Get XEN token balance
	balance, err := x.xenBalance(account)
	if err != nil {
		return "", err, 0
	}
	if balance.Cmp(big.NewInt(0)) == 0 {
		return "", fmt.Errorf("xen balance is 0"), 0
	}
	tx, err := x._XEN.Transfer(auth, x.transferXenTo, balance)
	if err != nil {
		return "", err, 0
	}

	humanBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))

	humanBalanceUint, _ := humanBalance.Uint64()
	return tx.Hash().String(), nil, humanBalanceUint
}
