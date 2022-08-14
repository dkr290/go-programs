package puller

import "fmt"

func Pull(c chan int) chan int {
	out := make(chan int)

	go func() {

		var sum int
		for n := range c {
			sum += n
			fmt.Println(sum)
		}
		out <- sum
		close(out)
	}()

	return out
}
