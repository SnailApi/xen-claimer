package main

import (
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

type Account struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
	Funded     bool   `json:"funded"`
	Claimed    bool   `json:"claimed"`
}

type XenClaimer struct {
	_XEN             *xen_abi.Store
	stakeDays        int
	fundingKey       string
	toFundEachWallet float64
	client           *ethclient.Client
	accountsToCreate int
	concurrency      int
	accounts         []*Account
	walletListPath   string
}

var wg sync.WaitGroup

var mutex = &sync.Mutex{}

const _XEN_contract = "0x06450dee7fd2fb8e39061434babcfc05599a6fb8"

func main() {
	accountsToCreat := flag.Int("a", 2, "Number of accounts to create")
	concurrency := flag.Int("c", 3, "Concurency")
	stakeForDays := flag.Int("d", 100, "Set number of days to stakes")
	walletListPath := flag.String("wl", "wallets.json", "Path to wallets list. Format per string(privateKey publicKey)")
	toFundEachWallet := flag.Float64("f", 0.01, "To fund each wallet")
	toDo := flag.Bool("fund", false, "Create, fund accounts and initiate claim")
	withdrawToFundingKey := flag.Bool("withdraw", false, "Withdraw funds from all acounts back to funding key")

	flag.Parse()
	godotenv.Load()

	ethRpcUrl := os.Getenv("ETH_RPC")
	if ethRpcUrl == "" {
		fmt.Println("\n::ERROR Missing EHT_RPC environment variable\n")
		return
	}
	fundingWallet := os.Getenv("FUNDING_WALLET")
	if ethRpcUrl == "" {
		fmt.Println("\n::ERROR Missing FUNDING_WALLET environment variable\n")
		return
	}

	client, err := ethclient.Dial(ethRpcUrl)
	if err != nil {
		fmt.Println(fmt.Sprintf("\n::ERROR %s\n", err))
		return
	}

	// Create Xen instance
	_XEN, err := xen_abi.NewStore(common.HexToAddress(_XEN_contract), client)
	if err != nil {
		log.Fatalln(err)
	}

	xen := XenClaimer{
		_XEN:             _XEN,
		stakeDays:        *stakeForDays,
		fundingKey:       fundingWallet,
		toFundEachWallet: *toFundEachWallet,
		client:           client,
		accountsToCreate: *accountsToCreat,
		concurrency:      *concurrency,
		accounts:         make([]*Account, 0),
		walletListPath:   *walletListPath,
	}

	// Check if file with wallets exist
	xen.loadExistingAccounts()

	// Withdraw funds from all accounts back to fundingKey
	if *withdrawToFundingKey {
		xen.withdrawToFundingKey()
		return
	}

	if *toDo {
		// Create new list of accounts if file with wallets doesn't exist
		if len(xen.accounts) == 0 {
			xen.generateAccounts()
			fmt.Println(fmt.Sprintf("\n::INFO %d accounts were created", len(xen.accounts)))
		} else {
			fmt.Println(fmt.Sprintf("\n::INFO found %d accounts in provided file", len(xen.accounts)))
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

		// Initiate claiming for accounts with funds
		xen.initiateClaiming()

		return
	}

	fmt.Println("\n::INFO What the fuck should i do? \n")
	flag.Usage()

}
