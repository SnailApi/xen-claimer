package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	xen_abi "github.com/drawrowfly/xen-claimer/xen_abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

var mutex = &sync.Mutex{}

func main() {
	accountsToCreat := flag.Int("a", 2, "Number of accounts to create")
	concurrencyFunding := flag.Int("cw", 10, "Concurency to fund wallets")
	concurrency := flag.Int("c", 10, "Concurency to claim XEN")
	stakeForDays := flag.Int("d", 100, "Set number of days to stakes")
	walletListPath := flag.String("wl", "wallets.json", "Path to wallets list")
	toFundEachWallet := flag.Float64("f", 0.01, "To fund each wallet")
	transferXenTo := flag.String("transfer-to", "", "Transfer claimed xen to the address")
	userFpcUrl := flag.String("rpc", "", "HTTP RPC url")
	userFundingWallet := flag.String("fw", "", "Private key for wallet that will be used to fund newly created wallets")
	// Main actions
	fundAndStakeXen := flag.Bool("fund", false, "Create, fund accounts and initiate XEN claim")
	claimAndTransferXen := flag.Bool("claim", false, "Claim XEN reward")
	withdrawToFundingKey := flag.Bool("withdraw", false, "Withdraw funds from all acounts back to funding key")
	claimDetails := flag.Bool("cd", false, "Get wallet claimed details, such as terms of claiming and etc")

	flag.Parse()
	godotenv.Load()

	rpc := os.Getenv("RPC")
	if rpc == "" {
		rpc = *userFpcUrl
	}
	if rpc == "" {
		fmt.Println("\n::ERROR Missing RPC environment variable\n")
		return
	}

	fundingWallet := os.Getenv("FUNDING_WALLET")
	if fundingWallet == "" {
		fundingWallet = *userFundingWallet
	}
	if *fundAndStakeXen || *withdrawToFundingKey {
		if fundingWallet == "" {
			fmt.Println("\n::ERROR Missing FUNDING_WALLET environment variable\n")
			return
		}
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Println(fmt.Sprintf("\n::ERROR %s\n", err))
		return
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println(fmt.Sprintf("\n::ERROR %s\n", err))
		return
	}

	chainInfo := getChainHelperData(int(chainId.Int64()))

	// Create Xen instance
	_XEN, err := xen_abi.NewStore(common.HexToAddress(chainInfo.XenContract), client)
	if err != nil {
		log.Fatalln(err)
	}

	xen := XenClaimer{
		_XEN:               _XEN,
		stakeDays:          *stakeForDays,
		fundingKey:         fundingWallet,
		toFundEachWallet:   *toFundEachWallet,
		client:             client,
		accountsToCreate:   *accountsToCreat,
		concurrency:        *concurrency,
		concurrencyFunding: *concurrencyFunding,
		accounts:           make([]*Account, 0),
		walletListPath:     *walletListPath,
		transferXenTo:      common.HexToAddress(*transferXenTo),
		chain:              chainInfo,
		tasks: Tasks{
			completed:       0,
			notCompleted:    0,
			totalTransfered: uint64(0),
		},
	}

	fmt.Println(fmt.Sprintf("\n::XEN CLAIMER START :: CHAIN %s :: CHAIN ID %d\n", chainInfo.ChainName, chainInfo.ChainId))

	// Check if file with wallets exist
	xen.loadExistingAccounts()

	// Start main actions

	// Get wallet claim details, crank, time and etc
	if *claimDetails {
		xen.getXenWalletClaimingDetails()

		fmt.Println(fmt.Sprintf("\n::XEN CLAIMER COMPLETED claim details task was completed :: COMPLETED: %d  NOT COMPLETED: %d \n", xen.tasks.completed, xen.tasks.notCompleted))
		return
	}

	// Withdraw funds from all accounts back to fundingKey
	if *withdrawToFundingKey {
		xen.withdrawToFundingKey()

		fmt.Println(fmt.Sprintf("\n::XEN CLAIMER COMPLETED withdraw all funds back to funding wallet was completed :: COMPLETED: %d  NOT COMPLETED: %d \n", xen.tasks.completed, xen.tasks.notCompleted))
		return
	}

	if *fundAndStakeXen {
		// Create new list of accounts if file with wallets doesn't exist
		if len(xen.accounts) == 0 {
			xen.generateAccounts()
			fmt.Println(fmt.Sprintf("\n::INFO %d accounts were created", len(xen.accounts)))
		} else {
			completed := 0
			for _, v := range xen.accounts {
				if v.Claimed {
					completed += 1
				}
			}
			notCompleted := len(xen.accounts) - completed
			if xen.accountsToCreate > notCompleted {
				xen.accountsToCreate = xen.accountsToCreate - notCompleted
				xen.generateAccounts()
			} else {
				xen.accountsToCreate = 0
			}
			//xen.accountsToCreate = xen.accountsToCreate - completed
			fmt.Println(fmt.Sprintf("\n::INFO found %d accounts in provided file :: TO CREATE %d", len(xen.accounts), xen.accountsToCreate))
		}

		// Get funding account balance
		fundingAccountBalance, err := xen.getBalance()
		if err != nil {
			fmt.Println(fmt.Sprintf("\n::ERROR %s\n", err))
			return
		}

		fmt.Println(fmt.Sprintf("\n::INFO funding account balance %f", fundingAccountBalance))

		fmt.Println(fmt.Sprintf("\n::INFO amount to transfer to each wallet %f", xen.toFundEachWallet))

		// Fund accounts by using funding key
		xen.fundAccounts()

		// Initiate XEN claiming for accounts with funds
		xen.initiateXenClaiming()

		fmt.Println(fmt.Sprintf("\n::XEN CLAIMER COMPLETED fund and claim task was completed :: COMPLETED: %d  NOT COMPLETED: %d \n", xen.tasks.completed, xen.tasks.notCompleted))
		return
	}

	if *claimAndTransferXen {

		if *transferXenTo == "" {
			fmt.Println("\n::ERROR Missing transfer-to wallet address\n")
		}
		// Initiate XEN reward claiming
		xen.initiateXenRewardClaiming()
		xen.initiateXenTransferToWallet()

		fmt.Println(fmt.Sprintf("\n::XEN CLAIMER COMPLETED claim and transfer task was completed :: COMPLETED: %d :: NOT COMPLETED: %d :: TRANSFERED :: %d \n", xen.tasks.completed, xen.tasks.notCompleted, xen.tasks.totalTransfered))
		return
	}

	fmt.Println("\n::INFO What the fuck should i do? \n")
	flag.Usage()

}
