package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Println("The value of b is pointer to a", *b)

	*b = 82
	fmt.Println("Change b with *b=82", a)

	changeme(&a)
	fmt.Println("after changme the value of a", a)

}

func changeme(b *int) {

	*b = 91

}
