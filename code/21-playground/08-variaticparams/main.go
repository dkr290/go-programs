package main

import "fmt"

func main() {

	n := average(43, 56, 87, 18, 45, 59)

	fmt.Println(n)

	// variatic arguments

	data := []float64{43, 56, 7, 89, 23, 32, 21}

	n1 := average(data...)
	fmt.Println(n1)

	n2 := average1(data)
	fmt.Println(n2)

}

//variatic parametes  function
func average(sf ...float64) float64 {

	fmt.Println(sf)
	fmt.Printf("%T\n", sf)

	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))

}

// alternative way to pass the same slice is

func average1(sf []float64) float64 {

	fmt.Println(sf)
	fmt.Printf("%T\n", sf)

	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))

}
