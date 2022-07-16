package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	name, x := sum("James", xi...)
	//x:= sum() it could be that it is not passin 0  so passing 0 or more variables
	fmt.Println("The total is", x)
	fmt.Println(name)
}

// func (r receiver ) identifier(parameter(s)) (return(s)) { code }
// the final argument if many should be vari like
func sum(s string, x ...int) (string, int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, " we are adding", v, "To the total which is now", sum)

	}
	return s, sum
}
