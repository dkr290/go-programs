package main

import "fmt"

type foo int

func main() {

	var myAge foo
	myAge = 91
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29

	fmt.Printf("%T %v \n", yourAge, yourAge)

	fmt.Println(int(myAge) + yourAge)

}
