package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type DoubleZero struct {
	person
	fname         string
	licenseToKill bool
}

func main() {

	p1 := DoubleZero{
		person: person{
			fname: "James",
			lname: "Bond",
			age:   30,
		},
		fname:         "This is overriding James Bond",
		licenseToKill: true,
	}

	p2 := DoubleZero{
		person: person{
			fname: "Miss",
			lname: "Moneypenny",
			age:   27,
		},

		fname:         "Seeme will be overwriten",
		licenseToKill: true,
	}

	fmt.Println(p1.fname, p1.lname, p1.licenseToKill, p1.person.fname)
	fmt.Println(p2.fname, p2.lname, p2.licenseToKill)

}
