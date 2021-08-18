package main

import (
	"os"

	"github.com/walkergriggs/hellosb/cmd"
)

func main() {
	command := cmd.NewCmdHellosb()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
