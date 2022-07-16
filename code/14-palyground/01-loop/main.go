package main

import "fmt"

func main() {
	i := 65

	for {

		if i <= 90 {
			fmt.Println(i)
			for j := 0; j < 3; j++ {
				fmt.Printf("\t\t\t%#U\n", i)
			}
		} else {
			break
		}
		i++

	}

}
