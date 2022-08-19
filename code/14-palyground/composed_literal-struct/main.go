package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {

	p1 := person{"Todd", "Mcleod"}
	p2 := person{"Nina", "Simone"}

	fmt.Println(p1)
	fmt.Println(p2)

	a1 := struct {
		breed string
		name  string
	}{
		"French Shepard",
		"Dimitrov",
	}
	fmt.Println(a1)

}
