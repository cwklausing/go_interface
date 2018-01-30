package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	square := square{sideLength: 12}
	triangle := triangle{height: 5, base: 12}

	printArea(square)
	printArea(triangle)
}
func printArea(g shape) {
	fmt.Println("Area:", g.getArea())
}
