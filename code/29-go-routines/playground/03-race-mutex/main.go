package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mux sync.Mutex

func main() {

	fmt.Println("Number of CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Add(2)
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()

}

func incrementor(s string) {
	for i := 0; i < 45; i++ {
		mux.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mux.Unlock()
	}
	wg.Done()
	fmt.Println("Number of CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
}
