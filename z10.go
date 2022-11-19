package main

import "fmt"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroups := make(map[int][]float64)

	for _, t := range temps {
		// for example:
		//	int(32.5) -> 32
		//	32 / 10 = 3.2 -> 3
		//	3 * 10 = 30
		groupKey := int(t) / 10 * 10
		tempGroups[groupKey] = append(tempGroups[groupKey], t)
	}

	fmt.Println(tempGroups)
}
