package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)

}
func main() {

	fmt.Println(factorial(4))

}

// 4 * call it again 3 * call again 2 and so on before hitting 0

//4 *3 *2 *1 *1
