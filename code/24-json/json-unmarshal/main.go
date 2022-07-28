package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Fname string
	Lname string
	Age   int
}

func main() {

	var p1 Person
	fmt.Println(p1.Fname)
	fmt.Println(p1.Lname)
	fmt.Println(p1.Age)

	var jsonBlob = []byte(`{"Fname":"James","Lname":"Bond","Age":29}`)

	err := json.Unmarshal(jsonBlob, &p1)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("------------------------------------------------------")
	fmt.Println(p1.Fname)
	fmt.Println(p1.Lname)
	fmt.Println(p1.Age)

}
