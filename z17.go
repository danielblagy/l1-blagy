package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{34, 5, 33, 2, 97, 33, -22, 35, 1, 3, 56, 76, 8, 4}
	sort.Ints(a)
	fmt.Println("input:", a)
	fmt.Println("output:", binarySearch(a, 8))
}

func binarySearch(arr []int, value int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < value {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != value {
		return -1
	} else {
		return low
	}
}
