package channel

import (
	"log"
	"os"
	"sync"
)

func Practice1() {
	logger := log.New(os.Stdout, "", 0)
	logger.Println("Practice 1")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Println("from goroutine 1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Println("from goroutine 2")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Println("from goroutine 3")
	}()

	logger.Println("from practice 1")
	wg.Wait()
}