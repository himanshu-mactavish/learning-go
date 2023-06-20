package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (square *Square) Area() float64 {
	return square.side * square.side
}

func (circle *Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	var areaIntf Shaper
	square := new(Square)
	square.side = 25.0
	circle := new(Circle)
	areaIntf = square
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Println("Area of sqaure is :", t)
	case *Circle:
		fmt.Println("Area of Circle is :", t)
	default:
		fmt.Println("Unknown Type")
	}
	circle.radius = 2.0
	fmt.Println(circle.Area())
}
