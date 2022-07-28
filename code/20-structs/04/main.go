package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p person) fullName() string {

	return p.fname + " " + p.lname

}

func main() {
	p1 := person{
		fname: "James",
		lname: "Bond",
		age:   31,
	}
	p2 := person{
		fname: "Miss",
		lname: "MoneyPenny",
		age:   24,
	}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

}
