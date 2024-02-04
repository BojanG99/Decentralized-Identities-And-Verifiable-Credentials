package did_blockchain

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Tester() {
	// Your code using the Ethereum client package here
	fmt.Println("Hello Ethereum!")

	// Example: Connect to an Ethereum node
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/cb91be5166984b73ac6e94685fd74987")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("")
	}

	fmt.Println(block.Number())
	addr := "0x1585e8eBc846F815C12EA633AEAF08D1A6d565dd"
	bal, err := client.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
	if err != nil {
		log.Fatal("")
	}
	fmt.Println(bal)
	// Now you can use 'client' to interact with the Ethereum blockchain
	// (e.g., query balances, send transactions, etc.)
}
