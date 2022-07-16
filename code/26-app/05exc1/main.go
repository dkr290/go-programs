package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Fname string
	Mname string
	Lname string
	Age   int
	Id    uint32
}

func main() {

	u1 := user{
		Fname: "James",
		Mname: "Telon",
		Lname: "Bond",
		Age:   34,
		Id:    3467867,
	}
	u2 := user{
		Fname: "Henry",
		Mname: "Still",
		Lname: "Bolem",
		Age:   23,
		Id:    34678456,
	}

	users := []user{u1, u2}

	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error", err)
	}
	os.Stdout.Write(b)
	fmt.Println(" ")
	fmt.Println(string(b))

}
