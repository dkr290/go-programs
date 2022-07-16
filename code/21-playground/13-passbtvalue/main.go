package main

import "fmt"

//this is pass by value and everything is pass by value

// pass by value and if we cwant to change on the other scope we pass the address

func main() {
	age := 44
	fmt.Println(&age)

	changeme(&age)

	fmt.Println(&age) // printing the address

	fmt.Println(age)

}

func changeme(z *int) {
	fmt.Println("Function changeme ------------")

	fmt.Println(z)  // this is the address
	fmt.Println(*z) // this is 44

	*z = 28
	fmt.Println(z)  // will give adress
	fmt.Println(*z) // it will give 28
}
