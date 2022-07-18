package main

import "fmt"

func zero(z int) { //passing by value x from main this is totally different variable
	fmt.Printf("%p\n", &z) // new variable new address
	fmt.Println(&z)
	z = 0 // so this is new memory address so nothing to do with x from function below
}

func zero1(z *int) {
	fmt.Println(z) // same address as the y from below
	*z = 0

}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	zero(x)        // pass by value so the result is 5 and not 0
	fmt.Println(x) // this is still 5 and

	fmt.Println("---------------------------------------")
	fmt.Println("Using pointers")

	y := 5
	fmt.Printf("%p\n", &y)
	zero1(&y)

	fmt.Println(y) // this time passing the & address to the pointer z from zero1 and y is 0 becuase of the pointer
}
