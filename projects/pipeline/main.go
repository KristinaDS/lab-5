package main

import (
	"fmt"
)

func removeDuplicates(inputStream, outputStream chan string) {
	var previousVal string
	for v := range inputStream {
		if previousVal != v {
			outputStream <- v
			previousVal = v
		}
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	var input string
	fmt.Printf("Введите текст с дубликатами: ")
	fmt.Scanln(&input)

	go func() {
		defer close(inputStream)

		for _, v := range input {
			inputStream <- string(v)
		}
	}()

	fmt.Printf("Текст без дубликатов: ")
	for v := range outputStream {
		fmt.Printf("%s", v)
	}
	fmt.Printf("\n")
}
