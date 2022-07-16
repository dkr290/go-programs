package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"Fname":"James","Lname":"Bond","Age":32},{"Fname":"Miss","Lname":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{}

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Print all the data", people)

	for i, v := range people {
		fmt.Println("\nPerson number ", i)
		fmt.Println(v.Fname, v.Lname, v.Age)
	}

}
