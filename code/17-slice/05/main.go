package main

import "fmt"

func main() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	fmt.Println(x1)
	fmt.Println(x2)

	xss := [][]string{x1, x2}
	fmt.Println(xss)

	for i, xs := range xss {
		fmt.Println("record", i)
		for j, val := range xs {

			fmt.Printf("\t index position: %v \t %v \n", j, val)
		}

	}

}
