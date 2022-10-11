package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (x *XenClaimer) checkTxHashStatus(txHash string) (int, error) {
	xHash := common.HexToHash(txHash)
	tx, err := x.client.TransactionReceipt(context.Background(), xHash)
	if err != nil {
		return -1, fmt.Errorf("Transaction Pending")
	}

	if tx.Status == 0 {
		return 0, fmt.Errorf("Transaction Failed")
	}
	return 1, fmt.Errorf("Transaction Completed")
}

func (x *XenClaimer) TransferEth(pKey string, toAddress string, amount *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		return "", fmt.Errorf("BAD 1 Wallet: %s error: %s \n", pKey, err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := x.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(1)
		return "", err
	}

	gasPrice, err := x.client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(3)
	}
	gasPrice = new(big.Int).Add(gasPrice, new(big.Int).Div(gasPrice, big.NewInt(6)))

	// If amount == 0 then we need to transfer whole balance to the specified address
	if amount.Cmp(big.NewInt(0)) == 0 {
		balance, err := x.client.BalanceAt(context.Background(), fromAddress, nil)
		if err != nil {
			return "", err
		}

		price := new(big.Int).Mul(gasPrice, big.NewInt(21000))

		amount = new(big.Int).Sub(balance, price)
		if err != nil {
			fmt.Println(3)
			return "", err
		}
	}

	gasLimit := uint64(21000)
	var data []byte
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), amount, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(x.chainId))), privateKey)
	if err != nil {
		return "", err
	}

	err = x.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), err
}

func (x *XenClaimer) getNonce() (uint64, error) {
	privateKey, err := crypto.HexToECDSA(x.fundingKey)
	if err != nil {
		return 0, fmt.Errorf("BAD 1 Wallet: %s error: %s \n", x.fundingKey, err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := x.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

func (x *XenClaimer) getBalance() (*big.Float, error) {
	privateKey, err := crypto.HexToECDSA(x.fundingKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Bad private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	balance, err := x.client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return nil, err
	}
	humanBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	return humanBalance, nil
}
