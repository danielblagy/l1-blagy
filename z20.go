package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "snow dog sun"
	fmt.Println("input:", text)

	// split the text by space and get an array of words
	textSlice := strings.Split(text, " ")

	// make a slice for reversed text
	reversedSlice := make([]string, 0, len(textSlice))

	// append to reversedSlice in reverse order
	for i := len(textSlice) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, textSlice[i])
	}

	// join reversedSlice, adding a space in between the elements
	reversedText := strings.Join(reversedSlice, " ")
	fmt.Println("output:", reversedText)
}
