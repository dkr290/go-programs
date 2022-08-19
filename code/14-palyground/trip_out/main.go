package main

import "fmt"

//custom type is like that -  define
type dog int

func (e dog) barnNum() {
	fmt.Println("This is a dog", e)
}

//how to define struct

type person struct {
	fName string
	lName string
}

func (p person) speak() {
	fmt.Println("Hello my name is ", p.fName, p.lName)

}
func main() {

	var d dog
	d = 17
	fmt.Printf("%T\n", d)
	d.barnNum()

	//composite literal; struct literal

	p1 := person{"Dimo", "Dimitrov"}
	p2 := person{"Simeon", "Simeonov"}

	fmt.Println(p1)
	fmt.Println(p2)

	p1.speak()
	p2.speak()

}
