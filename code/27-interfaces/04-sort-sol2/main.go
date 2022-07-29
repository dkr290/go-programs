package main

import (
	"fmt"
	"sort"
)

type people []string

func (a people) Len() int           { return len(a) }
func (a people) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a people) Less(i, j int) bool { return a[i] < a[j] }

func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Unsorted looks like that", studyGroup)
	sort.Sort(people(studyGroup))
	fmt.Println("Sorted", studyGroup)

}
