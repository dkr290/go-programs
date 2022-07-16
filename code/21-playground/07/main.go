package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Jane", " Doe"))
	fmt.Println(gree1t("Jane", " Doe"))
	fmt.Println(greet2t(" Jane", " Doe"))

}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

//return in a named value
func gree1t(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return s
}

//return multiply values

func greet2t(fname, lastname string) (string, string) {

	return fmt.Sprint(fname, lastname), fmt.Sprint(lastname, fname)
}
