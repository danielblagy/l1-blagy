package main

import "fmt"

func main() {
	stringsArray := []string{"cat", "cat", "dog", "cat", "tree"}
	// using a map of empty structs enables us to construct a set,
	// meaning every value will be unique, the empty struct data type doesn't take up any memory space
	set := make(map[string]struct{}, len(stringsArray))

	for _, v := range stringsArray {
		set[v] = struct{}{}
	}

	fmt.Println(set)
}
