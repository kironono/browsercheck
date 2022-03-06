package main

import (
	"os"

	"github.com/kironono/browsercheck/cmd/browsercheck/commands"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
