package main

import (
	"fmt"
)

var x int
var y float64

var myint int8 = -127

func main() {
	x = 42
	y = 428.535453

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(myint)

}
