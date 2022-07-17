package main

import (
	"fmt"

	"github.com/dkr290/go-programs/code/01-scope/visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
