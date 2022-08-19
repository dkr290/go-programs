package main

import "fmt"

// func (receiver) identifier(parameters) (returns) { code }

type person struct {
	fname string
	lname string
}

func (p person) speak() {

	fmt.Println("Hello, my name is ", p.fname, p.lname)

}
func main() {

	p1 := person{"Simon", "McLr"}
	p2 := person{"James", "Bond"}

	fmt.Println(p1)
	fmt.Println(p2)

	p1.speak()
	p2.speak()
}
