package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println([]string(studyGroup))
	sort.Strings([]string(studyGroup))
	fmt.Println(studyGroup)

	s := []string{"Zeno", "John", "Al", "Jenn"}
	sort.Strings(s)
	fmt.Println(s)

	s1 := []string{"Zeno", "John", "Al", "Jenn"}

	sort.Sort(sort.Reverse(sort.StringSlice(s1)))
	fmt.Println(s1)

	// there is also another way
	s2 := []string{"Zeno", "John", "Al", "Jenn"}
	fmt.Println("unsorted", s2)

	sort.StringSlice(s2).Sort()
	fmt.Println("sorted", s2)

}
