package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	messageChannel := make(chan int)
	// the amount of workers is set via command line argument
	workersAmount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("workers amount must be an integer")
		os.Exit(1)
	}

	// stop signal channel is used to send a signal when program exits,
	// so that goroutines can be shut down properly
	stopSignalChannel := make(chan os.Signal, 1)
	signal.Notify(stopSignalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	count := 0

	startWorkers(workersAmount, messageChannel, &wg)

	fmt.Println("workers started")

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			// when program is being closed with Ctrl + C, os.Signal will be sent
			// via stopSignalChannel
			case <-stopSignalChannel:
				// close messageChannel will cause the for range loop in
				// receiver goroutine to stop
				close(messageChannel)
				break loop

			// default prevents a goroutine from blocking
			// if no channels inside the select statement have incoming data, the default
			// code will be run
			default:
				messageChannel <- count
			}

			count++
		}
		fmt.Println("main loop stopped")
	}()

	wg.Wait()
	fmt.Println("program exiting")
}

func startWorkers(n int, messageChannel chan int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("worker", i, "has started")
			// the for range construct means that the data from messageChannel
			// will be read until the channel is closed
			for message := range messageChannel {
				fmt.Println("worker", i, ":", message)
			}
			fmt.Println("worker", i, "has finished")
		}(i)
	}
}
