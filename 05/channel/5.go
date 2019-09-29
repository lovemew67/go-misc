package channel

import (
	"fmt"
)

func Practice5() {
	ch := make(chan int, 4)
	ch <- 2
	ch <- 4
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}