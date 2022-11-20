package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// using math/big package for big numbers

	a := big.NewFloat(math.Pow(2, 22))
	b := big.NewFloat(math.Pow(2, 25))

	PrintMul(a, b)
	PrintDiv(a, b)
	PrintAdd(a, b)
	PrintSub(a, b)
}

func PrintMul(a *big.Float, b *big.Float) {
	result := big.NewFloat(0.0)
	result.Mul(a, b)
	fmt.Println("a * b =", result)
}

func PrintDiv(a *big.Float, b *big.Float) {
	result := big.NewFloat(0.0)
	result.Quo(a, b)
	fmt.Println("a / b =", result)
}

func PrintAdd(a *big.Float, b *big.Float) {
	result := big.NewFloat(0.0)
	result.Add(a, b)
	fmt.Println("a + b =", result)
}

func PrintSub(a *big.Float, b *big.Float) {
	result := big.NewFloat(0.0)
	result.Sub(a, b)
	fmt.Println("a - b =", result)
}
