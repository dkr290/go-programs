package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type DoubleZero struct {
	person
	licensetoKill bool
}

func main() {

	p1 := DoubleZero{
		person: person{
			fname: "James",
			lname: "Bond",
			age:   30,
		},
		licensetoKill: true,
	}

	p2 := DoubleZero{
		person: person{
			fname: "Miss",
			lname: "Moneypenny",
			age:   23,
		},
	}

	fmt.Println(p1.fname, p1.lname, p1.age, p1.licensetoKill)
	fmt.Println(p2.fname, p2.licensetoKill)
}
