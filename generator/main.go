package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// here we use generator function that returns channel, this channel will be filled of strings
// pushed into the channel from 100 goroutines.
func generator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		wg.Add(100)

		// heavy task!
		for i := 0; i < 100; i++ {
			go func() {
				defer wg.Done()
				time.Sleep(5 * time.Second)
				out <- "5 seconds passed"
			}()
		}

		wg.Wait()
	}()
	return out
}

func main() {
	for v := range generator() {
		fmt.Println(v)
	}
}
