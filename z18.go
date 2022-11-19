package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count      int
	countMutex sync.Mutex
}

// using mutex to prevent race conditions
func (c *Counter) Increment() {
	c.countMutex.Lock()
	c.count++
	c.countMutex.Unlock()
}

func (c Counter) Display() {
	fmt.Println(c.count)
}

func main() {
	counter := Counter{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	counter.Display()
}
