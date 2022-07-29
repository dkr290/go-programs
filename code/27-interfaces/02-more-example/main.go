package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side

}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
}

func displayInfo(shape Shape) {
	fmt.Println(shape.area())

}

func main() {

	s1 := Square{
		side: 10,
	}
	s2 := Square{
		side: 81,
	}

	c1 := Circle{
		radius: 87.32,
	}
	c2 := Circle{
		radius: 93.92,
	}

	fmt.Println("Information for the square")
	displayInfo(s1)
	displayInfo(s2)
	fmt.Println("Information for the Circle")
	displayInfo(c1)
	displayInfo(c2)

}
