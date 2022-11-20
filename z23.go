package main

import "fmt"

func main() {
	a := []int{34, 5, 33, 2, 97, 33, -22, 35, 1, 3, 56, 76, 8, 4}

	fmt.Println(a, "len:", len(a))
	a = deleteAt(a, 2)
	fmt.Println(a, "len:", len(a))
}

func deleteAt(s []int, index int) []int {
	// Append function used to append elements to a slice
	// first parameter as the slice to which the elements
	// are to be added/appended second parameter is the
	// element(s) to be appended into the slice
	// e.g. if s = [1, 2, 3, 4, 5]
	// and we call deleteAt(s, 2)
	// s[:index] == [1, 2]
	// s[index + 1:] == [4, 5]
	// append == [1, 2] + [4, 5] == [1, 2, 4, 5]
	return append(s[:index], s[index+1:]...)
}
