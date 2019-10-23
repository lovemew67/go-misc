package main

import (
	"log"
	
  	"github.com/lovemew67/go-misc/practices/25/command"
)


func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 26")
}


func main() {
	command.Execute()
}