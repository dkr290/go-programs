package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type DoubleZero struct {
	person
	licenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("Greeting from person")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Greeting from DoubleZero")
}

func main() {
	p1 := person{
		fname: "Ian",
		lname: "Flemming",
		age:   32,
	}

	p2 := DoubleZero{
		person: person{
			fname: "James",
			lname: "Bond",
			age:   41,
		},
		licenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
