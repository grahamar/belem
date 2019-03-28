package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"

	"github.com/grahamar/belem/root"

	// commands
	_ "github.com/grahamar/belem/login"
	_ "github.com/grahamar/belem/version"
)

func main() {
	log.SetHandler(cli.Default)

	args := os.Args[1:]
	root.Command.SetArgs(args)

	if err := root.Command.Execute(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
