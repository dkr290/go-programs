package main

import "fmt"

func main() {

	var x [5]int

	x[3] = 43

	fmt.Println(x)
	fmt.Println(len(x))

	//slices

	y := []int{1, 2, 3, 4, 5}
	fmt.Println(y)

	for i, v := range y {
		fmt.Println("Index is ", i, "value ", v)
	}
	fmt.Println(y[1])
	fmt.Println(y[2:4])

	ind := 0
	for ind < len(y) {
		fmt.Println(ind, y[ind])
		ind++
	}
	fmt.Println("-------------------------------")
	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}
	fmt.Println("---append---------")
	apd := []int{1, 2, 3, 4, 778, 38, 93}
	fmt.Println(apd)
	apd = append(apd, 77, 88, 99, 1014)
	fmt.Println(apd)

	apd1 := []int{345, 678, 432}
	apd = append(apd, apd1...)
	fmt.Println(apd)
}
