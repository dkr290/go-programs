package main

import (
	"fmt"

	"github.com/dkr290/go-programs/code/30-channels/incrementor-puller/incrementor"
	"github.com/dkr290/go-programs/code/30-channels/incrementor-puller/puller"
)

// this is the main function using incrementor with value to increment
func main() {

	c := incrementor.Incr(10)

	// pulling the value in channel and range over it to get the sum
	cSum := puller.Pull(c)

	for n := range cSum {
		fmt.Println(n)
	}

}
