package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e
	f
)

const (
	_ = iota //0
	n = iota * 10
	p = iota * 10
)

const (
	_  = iota             //0
	KB = 1 << (iota * 10) //1 <<  1* 10
	MB = 1 << (iota * 10) //1 <<  2 *10
	GB = 1 << (iota * 10) // 1 << 3 * 10
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println(n)
	fmt.Println(p)

	fmt.Printf("Decimal KB %d", KB)
	fmt.Printf("\tBinary %b\n", KB)
	fmt.Printf("Decimal MB %d", MB)
	fmt.Printf("\t Binary %b\n", MB)
	fmt.Println(GB)
	fmt.Printf("%b\n", GB)

}
