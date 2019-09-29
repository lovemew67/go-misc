package channel

import (
	"fmt"
	"time"
)

func Practice6() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "goroutine 1"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "goroutine 2"
	}()

	// for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("msg from goroutine1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("msg from goroutine2: %s\n", msg2)
		}
	// }
}