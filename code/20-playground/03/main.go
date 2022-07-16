package main

import "fmt"

func main() {

	p1 := struct {
		name      string
		age       int
		friends   map[string]int
		favDrinks []string
	}{
		name: "james",
		age:  91,
		friends: map[string]int{
			"Jorge":  5535,
			"Karoll": 893,
			"Bob":    943},
		favDrinks: []string{"Martini", "Wine"},
	}

	fmt.Println(p1)
	fmt.Printf("The name %v with age %v has friends\n", p1.name, p1.age)
	for n, v := range p1.friends {
		fmt.Printf("%v\t%v\t\n", n, v)
	}
	fmt.Print("And favorite drinks\n")
	for _, v := range p1.favDrinks {
		fmt.Printf(" %v\t\n", v)
	}
}
