package main

import (
	"os"

	"github.com/pokeh/golang-blockchain/cli"
)

func main() {
	// failsafe to properly close DB
	defer os.Exit(0)

	cli := cli.CommandLine{}
	cli.Run()
}
