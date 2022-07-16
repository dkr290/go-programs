package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}

	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("thje type of callback variable %T\n", callback)
	return xs
}

func main() {

	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 2
	})

	fmt.Println(xs) // [3 4]

}
