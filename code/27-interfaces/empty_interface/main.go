package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a any) {

	fmt.Println(a)

}

func main() {

	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meou"}, true}

	specs(fido)
	specs(fifi)
}
