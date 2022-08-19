package main

import "fmt"

func main() {

	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	//to add the record

	m["fleming_jan"] = []string{"BJams", "Geography", "Biking"}

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("For ", k, "the favorite things are")

		for i, val := range v {
			fmt.Printf("%d - \t %v\n", i, val)
		}
		fmt.Println()

	}

}
