package main

import (
	"fmt"

	"github.com/dkr290/go-programs/code/35-examples/sumutil"
)

func main() {

	fmt.Println(sumutil.Sum(2, 3, 4, 5, 6, 7, 78, 9))
	fmt.Println(sumutil.Sum(2, 3, 4, 5, 83))
}
