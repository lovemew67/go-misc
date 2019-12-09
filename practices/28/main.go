package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	orilog "log"
)

func init() {
	orilog.SetFlags(orilog.LstdFlags | orilog.LUTC | orilog.Lshortfile | orilog.Lmicroseconds)
	orilog.Println("Hello World: 28")

	log.SetOutput(os.Stdout)
	log.Println("Hello World: 28, one")

	log.SetFormatter(&log.TextFormatter{})
	log.Println("Hello World: 28, two")

	log.SetFormatter(&log.JSONFormatter{})
	log.Println("Hello World: 28, three")

	log.SetLevel(log.DebugLevel)
}

func main() {
	log.WithFields(log.Fields{
		"msg": "Hello World: 28",
	}).Trace("WTF")
}
