package main

import "fmt"

type person struct {
	firstNmae   string
	lastName    string
	favIceCream []string
}

func main() {
	p1 := person{
		firstNmae:   "John",
		lastName:    "DEpp",
		favIceCream: []string{"Vanilla", "chocolate"},
	}

	p2 := person{
		firstNmae:   "John",
		lastName:    "Blum",
		favIceCream: []string{"chocolate", "martini"},
	}

	fmt.Printf("The first person Firstname: %s Lastname: %s and favorite icecream: %s\n", p1.firstNmae, p1.lastName, p1.favIceCream)
	fmt.Printf("The first person Firstname: %s Lastname: %s and favorite icecream: %s\n", p2.firstNmae, p2.lastName, p2.favIceCream)

	fmt.Printf("The person Firstname: %s Lastname: %s\n", p1.firstNmae, p1.lastName)
	for _, v := range p1.favIceCream {
		fmt.Println("\t\t\thas favorite icecream flavour:", v)
	}
	fmt.Printf("The person Firstname: %s Lastname: %s\n", p2.firstNmae, p2.lastName)
	for _, v := range p2.favIceCream {
		fmt.Println("\t\t\thas favorite icecream flavour:", v)
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m)

	for _, v := range m {

		fmt.Println(v.firstNmae)
		fmt.Println(v.lastName)
		for i, val := range v.favIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("---------------")

	}
}
