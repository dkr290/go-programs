package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {

	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 45; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(3 * time.Microsecond))
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
