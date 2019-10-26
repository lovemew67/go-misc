package main

import (
	"os"

	"github.com/lovemew67/go-misc/projects/templates/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
