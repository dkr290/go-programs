package main

import "fmt"

func main() {

	m := make([]string, 1, 25)
	fmt.Println(m)
	changme(m)

	fmt.Println(m) // will print the same text

	fmt.Println("Example with the map")

	m1 := make(map[string]int)
	fmt.Println(m1)
	changeme1(m1)
	fmt.Println(m1)

}

func changme(z []string) {
	z[0] = "This is a text"
	fmt.Println(z)

}

func changeme1(z map[string]int) {
	z["KEY1"] = 83

}
