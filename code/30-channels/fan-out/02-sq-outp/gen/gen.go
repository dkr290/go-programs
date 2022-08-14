package gen

import "fmt"

func Gen(nums ...int) chan int {
	fmt.Printf("this is %T\n", nums)
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
