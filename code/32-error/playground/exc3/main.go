package main

import (
	"fmt"
)

type customErr struct {
	cust string
}

func (ce customErr) Error() string {

	return fmt.Sprintf("The very specific custom error comes: %v ", ce.cust)
}

func main() {
	c1 := customErr{"ERROR 305 check connection"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e.Error())

}
