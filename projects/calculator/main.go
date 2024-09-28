package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outp := make(chan int)
	go func(c chan int) {
		defer close(c)

		select {
		case n := <-firstChan:
			c <- n * n
		case n := <-secondChan:
			c <- n * 3
		case <-stopChan:
		}
	}(outp)

	return outp
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	stop := make(chan struct{})

	result := calculator(c1, c2, stop)

	c1 <- 3

	fmt.Println(<-result)
}
