package main

import (
	"fmt"
	"time"
)

func printMsg(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMsg("Goroutine") // Run in background
	printMsg("Main Thread")  // Run in main thread
}
