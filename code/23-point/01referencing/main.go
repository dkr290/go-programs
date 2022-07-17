package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	fmt.Println(*b) // pointer to the memory address will give the value
	// this is dereferencing of the pointer and give us the value

	// b points to the memory address
}
