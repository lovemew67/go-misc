package main

import (
	"log"
)

type persion struct {
	name string
	age  int
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
}

func stillWork() {
	log.Println("still need to work")
	howard := persion{
		name: "howard",
		age:  30,
	}
	log.Printf("%v", howard)
	log.Printf("%+v", howard)
	log.Printf("%#v", howard)
}

func main() {
	handlerName := "main"
	log.Printf("[%s] Hello World: 04", handlerName)

	try(
		func() {
			panic("foo")
		},
		func(e interface{}) {
			log.Printf("%+v", e)
			stillWork()
		},
	)
}

// source: https://www.douban.com/note/238705941/
func try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
