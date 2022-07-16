package main

import "fmt"

func main() {

	switch {
	case 2 == 4:
		fmt.Println("This will not print")

	case 4 == 4:
		fmt.Println("This will print")
		fallthrough
	case 4 >= 4:
		fmt.Println("This will also pring")
	default:
		fmt.Println("This is by default")
	}

}
