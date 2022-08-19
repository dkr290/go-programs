package main

import "fmt"

func main() {

	x := map[string]int{"Todd": 45, "Nina": 25, "Patrick": 27}

	for k, v := range x {
		fmt.Println(k, "-", v)
	}

	for k, _ := range x {

		fmt.Println(k, "-", x[k])
	}

	y := make(map[string]string)
	y["James"] = "007"
	fmt.Println(y["James"])

	x["Nina"] = 28
	fmt.Println(x["Nina"])
}
