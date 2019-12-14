package main

import (
	"log"
)

// type node struct {
// 	From    string
// 	To      string
// 	Enabled bool
// }

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile | log.Lmicroseconds)
	log.Println("Hello World: 26")
}

func main() {
	// construct forwardMap from database records
	forwardMap := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
		"D": "+886999111222",
		"E": "D",
	}
	ansMap := map[string]string{}
	for k := range forwardMap {
		current := k
		forwardPath := []string{k}
		loop := true
		// for loop {
		// 	if dest, ok := forwardMap[current]; ok {
		// 		if dest != k {
		// 			forwardPath = append(forwardPath, dest)
		// 			current = dest
		// 		} else {
		// 			loop = false
		// 		}
		// 	} else {
		// 		loop = false
		// 	}
		// }

		// FIXME: fix loop - 1 2 3 2
		for loop {
			dest, ok := forwardMap[current]
			if ok && dest != k {
				forwardPath = append(forwardPath, dest)
				current = dest
			} else {
				loop = false
			}
		}
		ansMap[k] = forwardPath[len(forwardPath)-1]
	}
	log.Printf("%+v", ansMap)
}
