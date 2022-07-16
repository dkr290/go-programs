package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		if i%4 != 0 {
			fmt.Printf("When %v is devided by 4 and the remainer is not 0 then the remainer  %v\n", i, i%4)

		}

	}

}
