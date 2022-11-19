package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 3, 34, 8, 9, 0, 6, 23, 6, 5}

	c1 := source(numbers...)
	c2 := processor(c1)
	sink(c2)
}

func source(numbers ...int) <-chan int {
	outChannel := make(chan int)

	go func() {
		for _, number := range numbers {
			outChannel <- number
		}
		close(outChannel)
	}()

	return outChannel
}

func processor(inChannel <-chan int) <-chan int {
	outChannel := make(chan int)

	go func() {
		for number := range inChannel {
			outChannel <- number * 2
		}
		close(outChannel)
	}()

	return outChannel
}

func sink(inChannel <-chan int) {
	for value := range inChannel {
		fmt.Println(value)
	}
}
