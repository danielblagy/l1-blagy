package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64 // lowercase, won't be accessible in another package
}

// setters and getters
func (p *Point) SetX(x float64) { p.x = x }

func (p *Point) SetY(y float64) { p.y = y }

func (p Point) GetX() float64 { return p.x }

func (p Point) GetY() float64 { return p.y }

func (p *Point) Set(x, y float64) {
	p.x = x
	p.y = y
}

func (p Point) Get() (float64, float64) {
	return p.x, p.y
}

// this method calculates the distance between 2 points
func (p Point) DistanceTo(o *Point) float64 {
	return math.Sqrt(math.Pow(p.x-o.x, 2) + math.Pow(p.y-o.y, 2))
}

func CreatePoint(x, y float64) *Point {
	return &Point{x, y}
}

func main() {
	p1 := CreatePoint(4, 5)
	p2 := CreatePoint(0, -3)

	fmt.Println("p1:", p1.GetX(), p1.GetY())
	fmt.Println("p2:", p2.GetX(), p2.GetY())
	fmt.Println("distance between p1 and p2:", p1.DistanceTo(p2))
}
