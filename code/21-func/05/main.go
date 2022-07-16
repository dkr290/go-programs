package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {

	fmt.Println("I am", s.firstName, s.lastName, "the secretAgent Speak")
}

func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, "the person speak")
}

type human interface {
	speak()
}

type hotdog int

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("It was passed into barrrrrrrr", h.(person).firstName)

	case secretAgent:
		fmt.Println("It was passed into barrrrrrrrrr", h.(secretAgent).firstName)

	}
	fmt.Println("It was passed into bar", h)
}

func main() {

	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			firstName: "Miss",
			lastName:  "Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		firstName: "Dr.",
		lastName:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
