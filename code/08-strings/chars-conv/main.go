package main

import "fmt"

func main() {

	for i := 50; i <= 140; i++ {
		fmt.Println(i, "-", string(i), " - ", []byte(string(i)))
	}

	f := 'a'
	fmt.Println(f)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", rune(f))

	l := "a"
	fmt.Println(l)
	fmt.Printf("%T \n", l)

}
