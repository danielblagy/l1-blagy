package main

import "fmt"

type Set map[int]struct{}

// using a map of empty structs enables us to construct a set,
// meaning every value will be unique, the empty struct data type doesn't take up any memory space
func NewSet(values []int) Set {
	newSet := Set{}
	for _, value := range values {
		newSet[value] = struct{}{}
	}
	return newSet
}

func GetIntersectingSet(s1 Set, s2 Set) Set {
	s := Set{}
	for k, _ := range s1 {
		// if a value is is both s1 and s2, put it in the intersecting set variable
		if _, ok := s2[k]; ok {
			s[k] = struct{}{}
		}
	}
	return s
}

func main() {
	s1 := NewSet([]int{23, 15, 20, 28, 5, 76, 2, 3, 45})
	s2 := NewSet([]int{23, 17, 33, 28, 4, 76, 1, 8, 54})
	fmt.Println(GetIntersectingSet(s1, s2))
}
