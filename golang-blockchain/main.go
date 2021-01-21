package main

import (
	"os"

	"github.com/RajyalakshmiUpputuri/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
