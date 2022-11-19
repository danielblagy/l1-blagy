package main

import (
	"fmt"
	"sync"
)

func main() {
	input := [5]int{2, 4, 6, 8, 10}

	var sum int = 0
	wg := sync.WaitGroup{}
	// use mutex to prevent a data race while concurrently modifying the same variable
	sumMutex := sync.Mutex{}

	for _, n := range input {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			// locking the mutex will block other goroutines from modifying the data
			// until the mutex is unlocked
			sumMutex.Lock()
			sum += n * n
			sumMutex.Unlock()
		}(n)
	}
	wg.Wait()

	fmt.Println(sum)
}

// not using a mutex will cause a race condition
/*
go run -race z3.go
==================
WARNING: DATA RACE
Read at 0x00c00012c078 by goroutine 8:
  main.main.func1()
      D:/dev/wb_internship/l1-blagy/z3.go:22 +0x7c
  main.main.func2()
      D:/dev/wb_internship/l1-blagy/z3.go:24 +0x47

Previous write at 0x00c00012c078 by goroutine 7:
  main.main.func1()
      D:/dev/wb_internship/l1-blagy/z3.go:22 +0x9a
  main.main.func2()

  main.main()
      D:/dev/wb_internship/l1-blagy/z3.go:18 +0xfc

Goroutine 7 (finished) created at:
  main.main()
      D:/dev/wb_internship/l1-blagy/z3.go:18 +0xfc
==================
220
Found 1 data race(s)
exit status 66
*/
