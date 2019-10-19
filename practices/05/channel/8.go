package channel

import (
	"fmt"
	"sync"
)

func Practice8() {
	ch := make(chan int, 20)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {	
		wg.Add(1)
		go func(i int) {
			fmt.Printf("going to send id: %d\n", i)
			ch <- i
			fmt.Printf("sent id: %d\n\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(ch)
	for {
		id, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("received from ch id: %d\n", id)
	}
}