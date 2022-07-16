package main

import "fmt"

func main() {

	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("after foo x", &x)
	fmt.Println("after foo x", x)
}

func foo(y *int) {

	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y passed by value to foo function and changed in foo", y)
	fmt.Println("y passed by value to foo function and changed in foo", *y)

}
