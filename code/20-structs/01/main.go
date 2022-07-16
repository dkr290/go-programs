package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	firstName string
	ltk       bool
}

func main() {

	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	p2 := person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		age:       56,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)

	sa1 := secretAgent{
		person: person{
			firstName: "james",
			lastName:  "Bond",
			age:       24,
		},
		firstName: "something collision",
		ltk:       true,
	}
	//specify only the embeded type sa1.person.firstName if there is the same name inside one and the other type
	fmt.Println(sa1.person.firstName, sa1.lastName, sa1.firstName, sa1.age, sa1.ltk)

}
