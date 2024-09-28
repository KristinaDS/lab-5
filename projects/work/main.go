package main

import (
	"fmt"
	"sync"
)

func work() {
	fmt.Println("work func is working")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
}
