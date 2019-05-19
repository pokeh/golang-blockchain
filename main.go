package main

import (
	"os"

	"github.com/pokeh/golang-blockchain/wallet"
)

func main() {
	// failsafe to properly close DB
	defer os.Exit(0)

	w := wallet.MakeWallet()
	w.Address()
}
