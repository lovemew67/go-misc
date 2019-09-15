package main

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
}

func stillWork() {
	log.Println("still work")
}

func main() {
	handlerName := "main"
	log.Printf("[%s] Hello World: 03", handlerName)

	defer func() {
		log.Println("c")
		if err := recover(); err != nil {
			log.Println(err)
		}
		log.Println("d")
		stillWork()
	}()

	log.Println("a")
	log.Panicln("pannnnnnnnnnnnnnnnnnnic")
	log.Println("b")
}
