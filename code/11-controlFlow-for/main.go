package main

import (
	"fmt"
)

func main() {
	// Simple for loop
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	// Nesting loops

	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tThe inner loop is  %d\n", j)
		}
	}

	//loop for condition statement

	x := 1

	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("good.")
	y := 1

	for {
		if y > 10 {
			break
		}
		fmt.Println(y)
		y++

	}
	fmt.Println("-------------------   ")
	w := 41 % 40
	l := 83 % 40
	fmt.Println(w, l)

	fmt.Println("-----------------------")

	// printing even numbers

	f := 1

	for {
		f++
		if f > 100 {
			break
		}
		if f%2 != 0 {
			continue

		}

		fmt.Println(f)

	}
	mynum := 33
	for {
		if mynum <= 122 {
			fmt.Printf("The number %d \t%#x is with UTF %#U\n", mynum, mynum, mynum)
		} else {
			break
		}
		mynum++
	}

}
