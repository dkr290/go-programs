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

func (p person) pSpeak() {

	fmt.Printf("The person has First Name : %s, Last Name: %s, Age: %d\n", p.fName, p.lName, p.age)
}

func (sa secertAgent) saSpeak() {

	fmt.Printf("The secret agent %s %s has badge number %d\n", sa.fName, sa.lName, sa.secretBadge)

}

func main() {

	p1 := person{"James", "Bond", 29}

	sa1 := secertAgent{
		person: person{"Miss", "MoneyPenny", 23},

		secretBadge: 2389,
	}
	fmt.Println(p1)
	fmt.Println(sa1)
	p1.pSpeak()
	sa1.saSpeak()
	sa1.pSpeak()

}
