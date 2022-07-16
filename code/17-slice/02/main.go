package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 5)

	fmt.Println("--------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("--------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Len: %v capacity %v value %v\n", len(mySlice), cap(mySlice), mySlice[i])

	}

	m1 := make([]string, 3, 7)
	m1 = append(m1, "Monday", "Tuesday", "Wednesday")
	m2 := make([]string, 3, 7)
	m2 = append(m2, "Thursday", "Friday")

	weekDays := append(m1, m2...)
	fmt.Println(weekDays)
	sl1 := make([]int, 10, 20)
	fmt.Println(sl1)

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	s2 = append(s2[1:2], s2[7:]...)
	fmt.Println(s2)
}
