package main

import "fmt"

func main() {
	if x := 32; x != 0 {
		fmt.Println(x)
	}

	x := 41
	if x == 40 {
		fmt.Println("Our value is 40")
	} else if x == 41 {
		fmt.Println("Our value is 41")
	} else if x == 42 {
		fmt.Println("Our value is 42")
	} else if x == 43 {
		fmt.Println("Our value is 43")
	} else {
		fmt.Println("Our value is not 40,41,42 or 43")
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}
