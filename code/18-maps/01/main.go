package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["barnabass"])

	v, ok := m["barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["todd"] = 34

	if v, ok := m["James"]; ok {
		fmt.Println("This is the IF PRINT", v)

	}

	for k, v := range m {
		fmt.Println("Indexes", k, "values", v)
	}

	xi := []int{4, 5, 6, 7, 8, 9}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	delete(m, "James")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value: ", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)

}
