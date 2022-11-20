package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	startTime := time.Now()
	for time.Since(startTime) < duration {

	}
}

func main() {
	startTime := time.Now()
	Sleep(time.Second * 3)
	elapsedTime := time.Since(startTime).Seconds()
	fmt.Println("time elapsed:", elapsedTime)
}
