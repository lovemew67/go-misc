package pack1

import (
	"log"
)

func Log() {
	handlerName := "log"
	log.Printf("[%s] print log", handlerName)
}