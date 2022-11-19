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

// source creates a channel and pushes the input values into it
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

// processor gets values from input channel, computes new values based on the input,
// and puts them into output channel
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

// sink gets values from the input channel and prints them
func sink(inChannel <-chan int) {
	for value := range inChannel {
		fmt.Println(value)
	}
}
