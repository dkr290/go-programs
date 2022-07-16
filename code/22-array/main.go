package main

import "fmt"

func main() {

	var x [59]string

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))

	for i := 65; i < 110; i++ {
		x[i-65] = string(i)

	}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x[42]))
}
