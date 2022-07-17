package main

import "fmt"

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // still 5

}

//pass by value and own copy of the value so nothing changed to the original x
func zero(x int) {
	x = 0
}
