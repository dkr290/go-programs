package main

import "fmt"

func main() {

	//maps are for key vlues

	// they are reference types

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v, ok := m["k1"]
	fmt.Println("ok?:", ok, v)

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//other way of declaration
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//some other examples

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good Morning."
	myGreeting["Miss"] = "Bonjour."

	fmt.Println(myGreeting)

	fmt.Println(myGreeting == nil)

	//with initialization

	var myGreeting1 = map[string]string{
		"Tim": "Good Morning",
	}

	fmt.Println(myGreeting1)

	myGreeting2 := map[string]string{
		"Tim":  "Good Morning",
		"Miss": "Bonjour",
	}

	fmt.Println(myGreeting2)
	fmt.Println(len(myGreeting2))

	mygreeting3 := map[string]string{

		"zero": "Good Morning",
		"one":  "Bonjour",
		"two":  "Boenos dias",
		"four": "Bongiorno",
	}
	for k, v := range mygreeting3 {
		fmt.Printf("The key %s  value :  %s\n", k, v)
	}

	delete(mygreeting3, "two")

	fmt.Println("##################  after delete ########################## ")

	for k, v := range mygreeting3 {
		fmt.Printf("The key %s  value :  %s\n", k, v)
	}

}
