package main

import "fmt"

const (
	m1 = 43
	m2 = 98.94343
	m3 = "This is a test"
)

const (
	l1 int     = 23
	l2 float64 = 44.89
	l3 string  = "James Wind"
)

const (
	year1 = iota + 2019
	year2
	year3
	year4
)

func main() {
	x := 93

	fmt.Println("last 4 years: ", year1, year2, year3, year4)
	fmt.Printf("%d\t%b\t%#x", x, x, x)
	fmt.Println()

	y := x << 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)
	fmt.Println(" ")

	z := ` This is the raw string
	in multiline and just to
	raw string
	literal
	"you see"`
	fmt.Println(z)

	a := (43 == 43)
	b := (43 <= 9)
	c := (46 >= 8)
	d := (4 != 9)
	e := (8 > 1)
	f := (8 < 1)

	fmt.Println(a, b, c, d, e, f)

	fmt.Println(m1, m2, m3)
	fmt.Println(l1, l2, l3)
}
