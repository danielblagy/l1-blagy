package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	output := make([]int, len(input))

	// use WaitGroup to prevent the main function from exiting before all the goroutines are done
	wg := sync.WaitGroup{}

	for i, n := range input {
		// add one to the WaitGroup counter
		wg.Add(1)
		go func(i int, n int) {
			// Decrement the WaitGroup counter when we're done
			defer wg.Done()
			// this is fine since we're always accessing different parts of the underlying array
			output[i] = n * n
		}(i, n)
	}
	// wait for the counter to go to zero, indicating all the goroutines are finished
	wg.Wait()

	// there's no practical benefit of concurrent execution here, since the calculation is too simple

	fmt.Println(output)
}
