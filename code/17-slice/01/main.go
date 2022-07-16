package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)
	//deleting from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	g := make([]int, 10, 100)
	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(cap(g))
	g[0] = 42
	g[9] = 999
	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(cap(g))

	g = append(g, 3423)
	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(cap(g))

	n := make([]int, 10, 12)
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))
	n[0] = 42
	n[9] = 999

	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))

	n = append(n, 3423)
	n = append(n, 3424)
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))

	n = append(n, 3424)
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))

	n = append(n, 3424)
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))

	//other type of slices

	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}

	fmt.Println(xp)

}
