package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// Simulate work
		time.Sleep(time.Duration(j) * time.Millisecond)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// BUG: this will deadlock if we try to read exactly 5 results.
	// Also, results channel is never closed.
	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
