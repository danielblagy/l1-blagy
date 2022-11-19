package main

import "fmt"

func main() {
	fmt.Println(reverse1("Hello World!"))
	fmt.Println(reverse2("Hello World!"))
}

// making a rune slice and swapping the letters
func reverse1(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

// making a rune slice and appending to it in reverse order
func reverse2(s string) string {
	ri := []rune(s)
	ro := make([]rune, 0, len(ri))

	for i := len(ri) - 1; i >= 0; i-- {
		ro = append(ro, ri[i])
	}

	return string(ro)
}
