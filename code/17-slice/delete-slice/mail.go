package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	var newarr []int

	newarr = append(newarr, x[0:3]...)
	newarr = append(newarr, x[6:]...)
	fmt.Println(newarr)

}
