package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPUs begin", runtime.NumCPU())
	fmt.Println("Goroutines begin", runtime.NumGoroutine())
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("First one:", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Second one", i)
		}
		wg.Done()
	}()
	fmt.Println("CPUs medium", runtime.NumCPU())
	fmt.Println("Goroutines medium", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
}
