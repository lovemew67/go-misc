package channel

import (
	"log"
	"os"
	"sync"
)

func Practice3() {
	logger := log.New(os.Stdout, "", 0)
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
		wg.Add(1)
		go func() {
			defer wg.Done()
			logger.Printf("print from goroutine: %d", <-ch)
		}()
	}

	logger.Println("print from practice3")
	wg.Wait()
}