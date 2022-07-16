package main

import (
	"fmt"
)

func main() {
	foo()

}

var x int = 4
var g func() = func() {
	fmt.Println("g from outside")
}

func foo() {
	func() {
		fmt.Println("Print from the inside function")
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	f := func() {
		fmt.Println("Print from the inside function")
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}
	f()

	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	g()
	g = f
	g()
	fmt.Printf("This is g %T\n", g)

	fmt.Println("THis is the main function foo")
}
