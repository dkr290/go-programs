package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "The quick brown fox jumps over the lazy dog and some other string words for separation"

	s1 := strings.Fields(str)

	fmt.Printf("Splitted in slice %v\n", s1)

	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()
	str2 := "GO1|The ABC of Go|25"

	s2 := strings.Split(str2, "|")

	for _, val1 := range s2 {
		fmt.Printf("%s ", val1)
	}

	fmt.Println()

	fmt.Println(strings.Join(s2, ";"))

	p := 10

	for p > 0 {
		fmt.Println(p)
		p--
	}

	str3 := "Go is a beautiful language!"
	for pos, char := range str3 {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

	fmt.Println()

	str4 := "Chinese: 日本語"
	for pos, char := range str4 {
		fmt.Printf("character %c starts at byte position %d \n", char, pos)

	}

	fmt.Println()
	fmt.Println("index int(rune) rune char bytes")
	for index, rune := range str4 {
		fmt.Printf("%-2d %d %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}

}
