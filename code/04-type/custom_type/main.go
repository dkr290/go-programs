package main

import (
	"fmt"
	"unicode/utf8"
)

//defining custom type with underlying type int
type TZ int

func main() {

	//define variables of type TZ
	var a, b TZ = 2, 9
	c := a + b

	s := "asSASA ddd dsjkdsjs dk"
	s1 := "asSASA ddd dsjkdsjsこん dk"
	// difference between runes and bytes with some characters whuch could be more bytes
	fmt.Println("bytes =", len(s))
	fmt.Println("runes =", utf8.RuneCountInString(s))
	fmt.Println("bytes =", len(s1))
	fmt.Println("runes =", utf8.RuneCountInString(s1))
	fmt.Printf("c has the value %d", c)

}
