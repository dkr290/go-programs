package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos Dias",
		3: "Bongiorno",
	}

	for k, v := range myGreeting {
		fmt.Println(k, "-", v)
	}

}
