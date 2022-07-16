package main

import "fmt"

func main() {

	x, even := f1(29)
	if even {
		fmt.Println("The number /2  is even ", x)
	} else {
		fmt.Println("The number /2  ", x, " is odd")
	}

}

func f1(x int) (int, bool) {

	xi := x / 2
	var even bool
	if xi%2 == 0 {

		even = true
	}

	return xi, even

}
