package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {

	p := person{
		fname: "Karl",
		lname: "Larson",
	}
	fmt.Println("The firstname and lastname before changeme function is ", p.fname, p.lname)
	changeMe(&p)
	fmt.Println("The firstname and lastname is ", p.fname, p.lname)

}

func changeMe(p *person) {

	p.fname = "James"
	p.lname = "Bond"

	//It can be done this way but also compiler can deriference it
	//(*p).fname = "James"
	//(*p).lname = "Bond"

}
