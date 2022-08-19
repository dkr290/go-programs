package main

import "fmt"

type persone struct {
	fname string
	lname string
}

func main() {
	p1 := persone{"Todd", "McLeod"}
	p2 := persone{"Nina", "Dimitrova"}

	fmt.Println(p1)
	fmt.Println(p2)

	m := meaning()
	fmt.Println(m)
}

func meaning() int {
	return 42

}
