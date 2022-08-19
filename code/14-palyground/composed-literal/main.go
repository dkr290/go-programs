package main

import "fmt"

func main() {

	//slice
	//composed literal; slice literal
	x := []int{7, 9, 42}
	fmt.Println(x)

	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 0, 100)

	y = append(y, 8)
	y = append(y, 12)
	y = append(y, 43)
	fmt.Println(y)

	fmt.Println(cap(y))

	y1 := make([]int, 0, 10)
	y1 = append(y1, 777)

	for i, v := range y1 {
		fmt.Println(i, "-", v)
	}
	//map
	//struct

}
