package channel

import (
	"log"
	"os"
)

func Practice2() {
	logger := log.New(os.Stdout, "", 0)
	message := make(chan string)

	go func() {
		message <- "hello from channel"
	}()

	logger.Println("wait for channel")
	msg := <- message
	logger.Println(msg)
}