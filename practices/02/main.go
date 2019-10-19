package main

import (
	"log"
	"os"

	"github.com/lovemew67/go-misc/practices/02/pack1"
	"github.com/lovemew67/go-misc/practices/02/pack2"
)

type people struct {
	name string
	age  int
}

const (
	test = "test"
)

var (
	howard = people{
		name: "howard",
		age:  18,
	}
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("initializing package: 02")
}

func main() {
	handlerName := "main"
	log.Println("Hello World: 02")
	log.Printf("[%s] test: %s", handlerName, test)
	log.Printf("[%s] People: %+v", handlerName, howard)
	pack1.Log()
	pack2.Log()

	logFile, errCreate := os.Create("log")
	if errCreate != nil {
		log.Panicf("[%s] failed to create log file", handlerName)
	}
	fileLogger := log.New(logFile, "FILELOG ", log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lshortfile)
	fileLogger.Printf("[%s] in file", handlerName)
}
