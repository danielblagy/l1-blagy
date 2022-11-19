package main

import "fmt"

func main() {
	a := []int{34, 5, 33, 2, 97, 33, -22, 34, 1, 3, 56, 76, 8, 4}
	fmt.Println("input:", a)
	quickSort(a, 0, len(a)-1)
	fmt.Println("output:", a)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}
