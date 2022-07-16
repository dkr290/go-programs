package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	z := 1
	fmt.Printf("%d\t\t%b\n", z, z)
	z_shift := z << 10
	fmt.Printf("%d\t\t%b\n", z_shift, z_shift)

	kb1 := 1024
	mb1 := 1024 * kb1
	gb1 := 1024 * mb1
	fmt.Printf("%d\t\t%b\n", kb1, kb1)
	fmt.Printf("%d\t\t%b\n", mb1, mb1)
	fmt.Printf("%d\t%b\n", gb1, gb1)

	fmt.Println("Printing the values for kb,mb,gb from iota example this time")
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
