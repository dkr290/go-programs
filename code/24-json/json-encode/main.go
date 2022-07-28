package main

import (
	"encoding/json"
	"os"
)

type person struct {
	Fname string
	Lname string
	Age   int
}

func main() {

	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Age:   32,
	}

	json.NewEncoder(os.Stdout).Encode(p1)

}
