package main

import "fmt"

type person struct {
	Fname string
	Lname string
	Age   int
}

func (p *person) speak() {

	fmt.Printf("First Name %s, Last Nmae %s and Age is %d\n", p.Fname, p.Lname, p.Age)

}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	h1 := person{
		Fname: "James",
		Lname: "Helr",
		Age:   21,
	}

	h1.speak()
	saySomething(&h1)
	//This one will now work becuase saySomethinhg shoud accept &h1 not h1
	//cannot use h1 (variable of type person) as human value in
	//argument to saySomething: person does not implement human (method speak has pointer receiver)
	//saySomething(h1)
}
