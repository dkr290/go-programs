package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fmt.Println("This is a text")
	fmt.Fprintf(os.Stdout, "This is another text ")
	_, err := io.WriteString(os.Stdout, "Hello World")
	if err != nil {
		log.Fatal(err)
	}

}
