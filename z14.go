package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	var a interface{}

	// randomly choose a value
	rand.Seed(time.Now().Unix())
	switch rand.Intn(5) {
	case 0:
		a = 15
	case 1:
		a = true
	case 2:
		a = "hello"
	case 3:
		a = float64(14.5)
	case 4:
		a = struct{ a, b, c int }{2, 3, 4}
	}

	fmt.Println("a =", a)

	// using type assertions and switch
	switch a.(type) {
	case int:
		fmt.Println("a is an int")
	case bool:
		fmt.Println("a is a bool")
	case string:
		fmt.Println("a is a string")
	case float64:
		fmt.Println("a is a float")
	case struct{ a, b, c int }:
		fmt.Println("a is a struct")

	default:
		fmt.Println("a is of an unknown")
	}

	// using fmt package
	fmt.Printf("the type of a derived with fmt: %T\n", a)

	// using reflect package
	fmt.Println("the type of a derived with reflect:", reflect.TypeOf(a).String())
}
