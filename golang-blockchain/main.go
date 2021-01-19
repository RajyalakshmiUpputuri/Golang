package main

import (
	"os"

	"github.com/RajyalakshmiUpputuri/Golang/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
