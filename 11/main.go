package main

import (
	"log"
	"os"

	"github.com/lovemew67/go-misc/11/command"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Llongfile)
}

func main() {
	if err := command.Execute(); err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
	os.Exit(0)
}
