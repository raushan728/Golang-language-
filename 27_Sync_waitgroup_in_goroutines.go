package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //Important: call Done when work finishes

	for i := 1; i <= 3; i++ {
		fmt.Printf("Worker %d doing task %d\n", id, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	//Launch 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1) // bataya ki ek aur goroutine aa rahi hai
		go worker(i, &wg)
	}

	wg.Wait() //Jab tak 3 baar Done() nahi hota, ruk ja
	fmt.Println("All workers done!")
}
