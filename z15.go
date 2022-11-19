package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(length int) string {
	return strings.Repeat("💖", length)
}

// justString will not contain 30 characters, but 30 bytes
func someFunc() {
	v := createHugeString(30)
	//justString = v[:30]
	// to get 30 bytes, we need to convert string to a slice of runes
	justString := string([]rune(v)[:30])
	fmt.Println(justString)
}

func main() {
	someFunc()
}
