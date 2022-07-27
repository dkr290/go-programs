package main

import "fmt"

func pName() {
	var u string
	fmt.Println("Plerase enter your name")
	arg, err := fmt.Scan(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arg)
	if arg > 1 {
		fmt.Println("Please input only one argument")
	}
	fmt.Println("The username is", u)
}

func main() {
	pName()

}
