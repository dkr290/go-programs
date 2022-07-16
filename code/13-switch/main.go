package main

import (
	"fmt"
)

func main() {

	switch {
	case false:
		fmt.Println("This should not run")

	case (9 == 4):
		fmt.Println("This should not run")
	case (2 == 2):
		fmt.Println("THis runs")
		fallthrough
	case (3 == 3):
		fmt.Println("This should also run last")
		fallthrough
	case (7 == 9):
		fmt.Println("not true 7 == 9")
		fallthrough
	case (11 == 14):
		fmt.Println("not thue 11 ==14")
		fallthrough //if previous is truth then print next one also
	case (15 == 15):
		fmt.Println("true 15==15")
	default:
		fmt.Println("This prints by default if nothing else prints")

	}
	n := "Bond"
	switch n {
	case "James":
		fmt.Println("This will not print")
	case "Bond", "MoneyPenny", "test":
		fmt.Println("James Bond")
	case "K":
		fmt.Println("This prints K")
	default:
		fmt.Println("This prints the default")
	}

}
