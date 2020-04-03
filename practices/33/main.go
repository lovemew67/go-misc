package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op == fsnotify.Create {
					fmt.Printf("file: %s, op: %s\n", event.Name, event.Op)
				} else if event.Op == fsnotify.Write {
					fmt.Printf("file: %s, op: %s\n", event.Name, event.Op)
				}
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()
	go func() {
		time.Sleep(5 * time.Second)
		file, err := os.OpenFile("C:/Users/zzh/fsnotify/aaaa.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		defer file.Close()
		_, err = file.WriteString("writes\n")
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
	}()
	if err := watcher.Add("C:/Users/zzh/fsnotify"); err != nil {
		fmt.Println("ERROR", err)
	}
	<-done
}
