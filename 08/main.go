package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type counter struct {
	sync.RWMutex
	count int
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 08")
}

func main() {

	// init counter
	count := counter{
		count: 0,
	}

	// // init file logger 1
	// logFile, errCreate := os.Create("log1")
	// if errCreate != nil {
	// 	log.Panicf("failed to create log file 1")
	// }
	// fileLogger1 := log.New(logFile, "FILELOG ", log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lshortfile)

	// init file logger 2
	logFile, errCreate := os.Create("log2")
	if errCreate != nil {
		log.Panicf("failed to create log file 2")
	}
	fileLogger2 := log.New(logFile, "FILELOG ", log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lshortfile)

	go func() {
		for {
			time.Sleep(time.Second * 5)
			count.Lock()
			// fileLogger1.Printf("increate counter from: %d to %d", count.count, count.count+1)
			log.Printf("increate counter from: %d to %d", count.count, count.count+1)
			count.count++
			count.Unlock()
		}
	}()

	for {
		fileLogger2.Println(">>> type 'count' to get current count")
		consolescanner := bufio.NewScanner(os.Stdin)

		// by default, bufio.Scanner scans newline-separated lines
		for consolescanner.Scan() {
			input := consolescanner.Text()
			fileLogger2.Printf("got: %s", input)
			if strings.Contains(input, "count") {
				count.RLock()
				fileLogger2.Printf("curent count: %d", count.count)
				count.RUnlock()
			} else {
				fileLogger2.Println("invalid command")
			}
			fileLogger2.Println(">>> type 'count' to get current count")
		}

		// check once at the end to see if any errors
		// were encountered (the Scan() method will
		// return false as soon as an error is encountered)
		if err := consolescanner.Err(); err != nil {
			fileLogger2.Printf(">>> err: %+v", err)
		}
	}
}
