package main

import "fmt"

const t string = "This is the tax declaration"

func main() {
	const q = 42

	fmt.Println("t - ", t)
	fmt.Println("q - ", q)

	const (
		pi       = 3.14
		language = "Go"
	)

	fmt.Println(pi)
	fmt.Println(language)

}
