package main

import (
	"fmt"
	"sync"

	"github.com/dkr290/go-programs/code/30-channels/fan-out/02-sq-outp/gen"
	"github.com/dkr290/go-programs/code/30-channels/fan-out/02-sq-outp/sq"
)

func main() {
	in := gen.Gen(2, 3)

	//FAN OUT
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq.Sq(in)
	c2 := sq.Sq(in)

	// FAN IN
	// Consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func merge(cs ...chan int) chan int {
	fmt.Printf("for cs: %T\n", cs) // just FYI

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed
FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.
PATTERN
there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.
source:
https://blog.golang.org/pipelines
*/
