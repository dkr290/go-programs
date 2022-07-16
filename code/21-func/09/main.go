package main

import "fmt"

func main() {

	a := incrementor()
	b := incrementor()
	fmt.Println("The value is ", a())
	fmt.Println("The value is ", a())
	fmt.Println("The value is ", a())
	fmt.Println("The value is ", a())
	fmt.Println("The value is ", a())
	fmt.Println("The value is ", a())
	fmt.Println("The value is ", a())

	fmt.Println("The value is ", b())
	fmt.Println("The value is ", b())
	fmt.Println("The value is ", b())
	fmt.Println("The value is ", b())

}

func incrementor() func() int {

	var x int
	return func() int {
		x++
		return x
	}
}
