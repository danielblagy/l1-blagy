package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MyMap struct {
	data      map[string]int
	dataMutex sync.Mutex
}

func CreateMyMap() *MyMap {
	return &MyMap{data: map[string]int{}, dataMutex: sync.Mutex{}}
}

func (m *MyMap) GetValue(key string) (int, bool) {
	// lock mutex while reading data from map
	m.dataMutex.Lock()
	value, ok := m.data[key]
	m.dataMutex.Unlock()
	return value, ok
}

func (m *MyMap) SetValue(key string, value int) {
	// lock mutex while writing data to map
	m.dataMutex.Lock()
	m.data[key] = value
	m.dataMutex.Unlock()
}

func (m MyMap) Display() {
	fmt.Println(m.data)
}

func main() {
	keys := []string{"A", "B", "C", "D"}
	m := CreateMyMap()

	for _, key := range keys {
		m.SetValue(key, 0)
	}

	wg := &sync.WaitGroup{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			key := keys[rand.Intn(len(keys))]
			value, _ := m.GetValue(key)
			m.SetValue(key, value+5)
		}()
	}

	wg.Wait()

	m.Display()
}

// without mutex
/*
go run -race z7.go
==================
WARNING: DATA RACE
Read at 0x00c000144480 by goroutine 8:
  runtime.mapaccess2_faststr()
      C:/Program Files/Go/src/runtime/map_faststr.go:108 +0x0
  main.(*MyMap).GetValue()
      D:/dev/wb_internship/l1-blagy/z7.go:19 +0x14b
  main.main.func1()
      D:/dev/wb_internship/l1-blagy/z7.go:48 +0x122

Previous write at 0x00c000144480 by goroutine 7:
  runtime.mapassign_faststr()
      C:/Program Files/Go/src/runtime/map_faststr.go:203 +0x0
  main.(*MyMap).SetValue()
      D:/dev/wb_internship/l1-blagy/z7.go:24 +0x199
  main.main.func1()
      D:/dev/wb_internship/l1-blagy/z7.go:49 +0x1a9

Goroutine 8 (running) created at:
  main.main()
      D:/dev/wb_internship/l1-blagy/z7.go:44 +0x330

Goroutine 7 (finished) created at:
  main.main()
      D:/dev/wb_internship/l1-blagy/z7.go:44 +0x330
==================
==================
WARNING: DATA RACE
Read at 0x00c000116cd0 by goroutine 9:
  main.(*MyMap).GetValue()
      D:/dev/wb_internship/l1-blagy/z7.go:19 +0x158
  main.main.func1()
      D:/dev/wb_internship/l1-blagy/z7.go:48 +0x122

Previous write at 0x00c000116cd0 by goroutine 7:
  main.(*MyMap).SetValue()
      D:/dev/wb_internship/l1-blagy/z7.go:24 +0x1a8
  main.main.func1()
      D:/dev/wb_internship/l1-blagy/z7.go:49 +0x1a9

Goroutine 9 (running) created at:
  main.main()
      D:/dev/wb_internship/l1-blagy/z7.go:44 +0x330

Goroutine 7 (finished) created at:
  main.main()
      D:/dev/wb_internship/l1-blagy/z7.go:44 +0x330
==================
map[A:25 B:30 C:25 D:20]
Found 2 data race(s)
exit status 66
*/
