package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type foo int

func main() {

	p1 := person{
		fname: "James",
		lname: "Bond",
		age:   20,
	}

	p2 := person{"Miss", "Moneypenny", 18}

	fmt.Println(p1.fname, p1.lname, p1.age)
	fmt.Println(p2.fname, p2.lname, p2.age)

	var myAge foo
	myAge = 91
	fmt.Printf("%T %v \n", myAge, myAge)

}
