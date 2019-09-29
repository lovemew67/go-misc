package channel

import (
	"fmt"
	"strings"
)

func Practice7() {
	data := []string{
        "The yellow fish swims slowly in the water",
        "The brown dog barks loudly after a drink from its water bowl",
        "The dark bird of prey lands on a small tree after hunting for fish",
	}
	
	histogram := map[string]int{}
	wordsChannel := make(chan string)

	go func() {
		defer close(wordsChannel)
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				wordsChannel <- word
			}
		}
	}()

	for {
		word, isOpened := <- wordsChannel
		if !isOpened {
			break
		}
		histogram[word]++
	}

	for k, v := range histogram {
		fmt.Printf("%s\t%d\n", k, v)
	}
}