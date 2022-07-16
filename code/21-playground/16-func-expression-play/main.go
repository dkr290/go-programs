package main

import "fmt"

func main() {

	half := func(x int) (int, bool) {
		xi := x / 2
		return xi, xi%2 == 0

	}

	x, even := half(29)
	if even {
		fmt.Println("The result of original number / 2  is even ", x)
	} else {
		fmt.Println("The result of original number  /2 is ", x, " is odd")
	}

}
