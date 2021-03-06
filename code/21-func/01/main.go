package main

import "fmt"

func main() {

	foo()
	bar("James")

	s1 := woo("MoneyPenny")
	fmt.Println(s1)

	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

}

//func

// func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {

	fmt.Println("hello from foo")

}

//Everything is GO is pass by VALUE
func bar(s string) {

	fmt.Println("Hello ", s)

}

func woo(s string) string {
	return fmt.Sprint("Hello from woo ", s)

}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true

	return a, b
}
