package main

import "fmt"

var a int = 43

type hotdog int

var b hotdog

func main() {

	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// how to make conversion

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
