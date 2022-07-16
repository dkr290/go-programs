package main

import (
	"encoding/json"
	"fmt"
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
	p2 := person{
		Fname: "Miss",
		Lname: "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	//func Marshal(v any) ([]byte, error)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs))
}
