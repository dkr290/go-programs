package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname       string
	Lname       string
	Age         int
	notExported int
}

type Person1 struct {
	FirstName string
	LastName  string `json:"-"`
	Age       int    `json:"wisdom score"`
}

func main() {

	p1 := person{
		Fname:       "James",
		Lname:       "Bond",
		Age:         20,
		notExported: 007,
	}

	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

	p2 := Person1{
		FirstName: "Jimmy",
		LastName:  "Wiskey",
		Age:       29,
	}
	bs1, _ := json.Marshal(p2)
	fmt.Println(string(bs1))
}
