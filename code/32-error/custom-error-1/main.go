package main

import (
	"fmt"
	"log"
	"os"
)

// init function to set output for logging to that particular file
func init() {

	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)

}

func main() {

	_, err := os.Open("no_file.txt")
	if err != nil {
		log.Println("eroor happened", err)
		//log.Fatalln(err)
	}
}
