package main

import (
	"drylang/cli"
	"os"
)

func main() {
	os.Exit(cli.RunCLI(os.Args[1:]))
}
