package main

import (
	"fmt"

	"github.com/dkr290/go-programs/code/30-channels/channel-pipeline-sqrt/gen"
	"github.com/dkr290/go-programs/code/30-channels/channel-pipeline-sqrt/sq"
)

func main() {

	c := gen.Gen(3, 8)

	out := sq.Sq(c)

	for v := range out {
		fmt.Println(v)
	}

}
