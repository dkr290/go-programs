package main

import (
	"fmt"
)

func calcNum() int {

	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}
	return n

}

func main() {

	fmt.Println("Plerase enter large number")
	b := calcNum()
	fmt.Println("Plerase enter small number")
	s := calcNum()

	fmt.Println("Remainder is", b%s)
}
