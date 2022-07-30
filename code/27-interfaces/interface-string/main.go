package main

import "fmt"

func main() {

	var name any = "String123"
	str, ok := name.(string)

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("value is not a string")
	}

	var val interface{} = 91
	fmt.Printf("%T\n", val)
	fmt.Printf("%T\n", val.(int))

	fmt.Println(val.(int) + 6)

}
