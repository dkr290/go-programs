package main

import "fmt"

func main() {

	//maps are for key vlues

	// they are reference types

	// defining maps

	/*
		     m := make(map[string]int)
			 m := map[string]int{} // to create not nil map
		     m := map[string]int{"key1":3,"key2":3}
			 var m = make(map[string][string])

	*/

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

	fmt.Println("#############################  another  ############################")

	mygreeting4 := map[int]string{

		1: "Good Morning",
		2: "Bonjour",
		3: "Boenos dias",
		4: "Bongiorno",
	}

	fmt.Println(mygreeting4)

	if val, exists := mygreeting4[2]; exists {
		delete(mygreeting4, 2)
		fmt.Println("val", val)
		fmt.Println("exists", exists)
	} else {
		fmt.Println("The value does not exists")
		fmt.Println("val", val)
		fmt.Println("exists", exists)
	}
}
