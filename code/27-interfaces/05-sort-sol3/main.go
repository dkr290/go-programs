package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Fname string
	Lname string
	Age   int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s: %d", p.Fname, p.Lname, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	// sort strings
	n := []int{5, 2, 6, 3, 1, 4}

	sort.Ints(n)
	fmt.Println(n)

	//sort custom type struct

	people := []Person{
		{Fname: "Bob", Lname: "Marley", Age: 31},
		{Fname: "John", Lname: "Helmes", Age: 81},
		{Fname: "Lil", Lname: "Framen", Age: 32},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	for _, v := range people {
		fmt.Println("Name", v.Fname, "- Last Name:", v.Lname, "- Age:", v.Age)
	}

}
