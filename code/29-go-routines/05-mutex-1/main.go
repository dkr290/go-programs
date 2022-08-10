package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

const gs int = 100

var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(gs)
	for i := 0; i < gs; i++ {

		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Count", counter)

}
