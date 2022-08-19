package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

type secertAgent struct {
	person
	secretBadge int
}

func (p person) speak() string {

	return fmt.Sprintf("The person has First Name : %s, Last Name: %s, Age: %d\n", p.fName, p.lName, p.age)
}

func (sa secertAgent) speak() string {

	return fmt.Sprintf("The secret agent %s %s has badge number %d\n", sa.fName, sa.lName, sa.secretBadge)

}

type human interface {
	speak() string
}

func info(h human) {
	fmt.Printf("%v", h.speak())
}

func main() {

	p1 := person{"James", "Bond", 29}

	sa1 := secertAgent{
		person: person{"Miss", "MoneyPenny", 23},

		secretBadge: 2389,
	}
	fmt.Println(p1)
	fmt.Println(sa1)
	info(p1)
	info(sa1)

}
