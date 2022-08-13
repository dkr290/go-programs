package main

import (
	"fmt"

	"github.com/dkr290/go-programs/code/36-playgroud-tests/02/quote"
	"github.com/dkr290/go-programs/code/36-playgroud-tests/02/word"
)

func main() {

	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

}
