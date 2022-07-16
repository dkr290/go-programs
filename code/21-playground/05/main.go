package main

import (
	"fmt"
	"math"
)

type square struct {
	lenght float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius

}

func (s square) area() float64 {

	return s.lenght * s.lenght

}

type shape interface {
	area() float64
}

func info(s shape) {

	fmt.Println("The area  is", s.area())

}

func main() {
	circ := circle{
		radius: 12.889,
	}
	sqr := square{
		lenght: 18.3,
	}

	info(circ)
	info(sqr)

}
