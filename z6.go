package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	stopSignalRoutine()
	closeChannelRoutine()
	contextRoutine()
}

// having a separate channel for sending a signal to stop the goroutine
// stop signal channels are usually of type struct{} or bool
func stopSignalRoutine() {
	wg := &sync.WaitGroup{}
	doneChannel := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("'stop signal' goroutine started")
	loop:
		for {
			select {
			// when the goroutine receives the stop signal, break out of the loop
			case <-doneChannel:
				break loop

			default:
			}
		}

		fmt.Println("'stop signal' goroutine finished")
	}()

	time.Sleep(time.Millisecond * 10)
	fmt.Println("sending stop signal")
	// send data to signal, or close the channel
	doneChannel <- true

	wg.Wait()
}

func closeChannelRoutine() {
	wg := &sync.WaitGroup{}
	dataChannel := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("'close channel' goroutine started")
		// when dataChannel is closed, the loop stops
		for range dataChannel {

		}
		fmt.Println("'close channel' goroutine finished")
	}()

	// Alternative
	/*go func() {
		defer wg.Done()

		fmt.Println("'close channel' goroutine started")
	loop:
		for {
			// ok is false if the channel is closed
			// in the case of the channel being closed, the default value will be returned in data
			data, ok := <-dataChannel
			if !ok {
				break loop
			}
			fmt.Println(data)
		}

		fmt.Println("'close channel' goroutine finished")
	}()*/

	time.Sleep(time.Millisecond * 10)
	fmt.Println("closing the data channel")
	close(dataChannel)

	wg.Wait()
}

func contextRoutine() {
	wg := &sync.WaitGroup{}
	// create a timeout context which signals when it's done via ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("'context timeout' goroutine started")
	loop:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context is done")
				break loop

			default:
			}
		}

		fmt.Println("'context timeout' goroutine finished")
	}()

	wg.Wait()
}
