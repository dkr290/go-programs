package main

import "fmt"

//how to make structs and relate to variable definition
var x int

type mytype struct {
	firstName string
	lastName  string
	age       int
	myfloat   float64
	myTest    string
}

func main() {
	//annonimous struct
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	fmt.Println(p1.firstName, p1.lastName, p1.age)

}
