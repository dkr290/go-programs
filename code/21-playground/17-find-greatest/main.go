package main

import "fmt"

func main() {

	x := []int{1, 5, 8, 10, 44, 32}
	fmt.Println("The greatest number is ", findG(x...))

}

func findG(x ...int) int {

	var g int
	for _, v := range x {

		if v > g {
			g = v
		}
	}
	return g
}
