package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(c shape) {
	fmt.Println("area is", c.area())
}

func main() {

	c := circle{
		radius: 5,
	}
	// if we try to do the info with passing only the c withrout passing the address will fail
	//cannot use c (variable of type circle) as shape value in argument to info: circle does not implement shape (method area has pointer receiver)

	//info(c) // the metod has pointer receiver
	info(&c)
	fmt.Println(c.area())
}
