package main

import (
	"fmt"
	"os"
)

func main() {

	var o string = os.Getenv("GOOS")

	fmt.Println("The operating system is ", o)

	path := os.Getenv("GOPATH")
	fmt.Println("The path is", path)
}
