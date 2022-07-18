package main

import "fmt"

func main() {

	i := 0
	w(&i)
	fmt.Println("For loop and break below")
	j := 0
	f_break(&j)
	fmt.Println("This is another way")
	b := 0
	f_continue(&b)

}

func w(i *int) {

	// this is like while loop
	for *i < 10 {
		fmt.Println(*i)
		*i++
	}

}

func f_break(i *int) {
	// this is for loop only with break otherwise will continuosly loop
	for {
		if *i >= 10 {
			break
		}
		fmt.Println(*i)
		*i++

	}
}

func f_continue(i *int) {

	for {
		*i++
		if *i%2 == 0 {
			continue
		}
		fmt.Println(*i)
		if *i > 50 {
			break
		}
	}
}
