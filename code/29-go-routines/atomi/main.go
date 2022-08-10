package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var counter int64

const gs int = 100

var wg sync.WaitGroup

func main() {

	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(gs)
	for i := 0; i < gs; i++ {

		go func() {

			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()

			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Count", counter)

}
