package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("The firstname", p.first, "the lastname", p.last, "and the age", p.age)
}
func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   21,
	}
	p1.speak()

}
