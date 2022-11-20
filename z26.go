package main

import (
	"fmt"
	"unicode"
)

func allCharUnique(s string) bool {
	// use a map of empty structs to check if the character is already present in s
	uniqueChars := map[rune]struct{}{}
	for _, c := range s {
		// make case-insensitive
		c = unicode.ToLower(c)
		// check if the character is already present in s
		if _, ok := uniqueChars[c]; ok {
			return false
		}
		uniqueChars[c] = struct{}{}
	}
	return true
}

func main() {
	s1 := "Hello World!" // false
	s2 := "abcdtorNA"    // false
	s3 := "ahdliem60x🤩"  // true

	fmt.Println("s1:", s1, "-", allCharUnique(s1))
	fmt.Println("s2:", s2, "-", allCharUnique(s2))
	fmt.Println("s3:", s3, "-", allCharUnique(s3))
}
