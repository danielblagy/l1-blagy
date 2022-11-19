package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// the data channel
	messageChannel := make(chan int)
	// used to send 'done' signal
	doneChannel := make(chan struct{})

	// the duration of execution is set via command line arguments
	setDuration, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("workers amount must be an integer")
		os.Exit(1)
	}

	// get the starting time
	startTime := time.Now()
	// start sender and receiver goroutines
	wg.Add(2)
	go startSender(messageChannel, doneChannel, &wg)
	go startReceiver(messageChannel, doneChannel, &wg)

	// sleep for the set duration in seconds, then send the 'done' signal by closing the channel
	time.Sleep(time.Duration(setDuration) * time.Second)
	close(doneChannel)

	wg.Wait()
	fmt.Println("program exiting after", time.Since(startTime).Seconds(), "seconds")
}

// send data until there's a signal send via doneChannel
func startSender(messageChannel chan int, doneChannel chan struct{}, wg *sync.WaitGroup) {
loop:
	for i := 0; ; i++ {
		select {
		// break out of the loop once the 'done' signal is sent
		case <-doneChannel:
			break loop

		// on each iteration of the loop, new data will be sent
		// (the buffer size of messageChannel is 1, so the goroutine will be blocked untli the data
		//  is read)
		default:
			messageChannel <- i
		}
	}
	wg.Done()
}

// receive data until there's a signal send via doneChannel
func startReceiver(messageChannel chan int, doneChannel chan struct{}, wg *sync.WaitGroup) {
loop:
	for i := 0; ; i++ {
		select {
		// break out of the loop once the 'done' signal is sent
		case <-doneChannel:
			break loop

		// when there's new data in message channel, print it
		case message := <-messageChannel:
			fmt.Println(message)
		}
	}
	wg.Done()
}
