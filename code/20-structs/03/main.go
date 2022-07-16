package main

import "fmt"

//how to make structs and relate to variable definition
var x int

type foo int

var y foo

const bar int = 42
const r = 82

type mytype struct {
	firstName string
	lastName  string
	age       int
	myfloat   float64
	myTest    string
}

func main() {
	p1 := mytype{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
		myfloat:   38.3,
		myTest:    "Thisi s another string",
	}

	fmt.Println(p1.firstName, p1.lastName, p1.age)

	fmt.Println(bar)
	fmt.Println(r)

	y = foo(bar)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", r)

	y = r
	fmt.Printf("%T\t%v\n", y, y)
	fmt.Printf("%T\n", int(y))

}
