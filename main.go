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
	walletListPath := flag.String("wl", "wallets.json", "Path to wallets list. Format per string(privateKey publicKey)")
	toFundEachWallet := flag.Float64("f", 0.01, "To fund each wallet")
	toDo := flag.Bool("fund", false, "Create, fund accounts and initiate claim")
	withdrawToFundingKey := flag.Bool("withdraw", false, "Withdraw funds from all acounts back to funding key")

	flag.Parse()
	godotenv.Load()

	client, err := ethclient.Dial(os.Getenv("ETH_RPC"))
	if err != nil {
		log.Fatal(err)
	}

	// Create Xen instance
	_XEN, err := xen_abi.NewStore(common.HexToAddress(_XEN_contract), client)
	if err != nil {
		log.Fatalln(err)
	}

	xen := XenClaimer{
		_XEN:             _XEN,
		fundingKey:       os.Getenv("FUNDING_WALLET"),
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
			fmt.Println("\n" + fmt.Sprintf("::INFO %d accounts were created", len(xen.accounts)))
		} else {
			fmt.Println("\n" + fmt.Sprintf("::INFO found %d accounts in provided file", len(xen.accounts)))
		}

		// Get funding account balance
		fundingAccountBalance := xen.getBalance()

		fmt.Println("\n" + fmt.Sprintf("::INFO funding account balance %f", fundingAccountBalance))

		fmt.Println("\n" + fmt.Sprintf("::INFO amount to transfer on each wallet %f", xen.toFundEachWallet))

		// Fund accounts by using funding key
		xen.fundAccounts()

		// Initiate claiming for accounts with funds
		xen.initiateClaiming()

		return
	}

	fmt.Println("\n::INFO What the fuck should i do? \n")
	flag.Usage()

}
