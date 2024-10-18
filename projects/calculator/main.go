package main

import (
	"fmt"
	"time"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outp := make(chan int)
	go func(ch chan int) {
		defer close(ch)

		select {
		case n := <-firstChan:
			ch <- n * n
		case n := <-secondChan:
			ch <- n * 3
		case <-stopChan:
		}
	}(outp)

	return outp
}

func main() {
	firstCh := make(chan int)
	secondCh := make(chan int)
	stopCh := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		firstCh <- 100
	}()

	go func() {
		time.Sleep(2 * time.Second)
		secondCh <- 7
	}()

	output := calculator(firstCh, secondCh, stopCh)

	select {
	case result := <-output:
		fmt.Println("Результат:", result)
	case <-time.After(5 * time.Second):
		fmt.Println("Время ожидания истекло")
	}
}
