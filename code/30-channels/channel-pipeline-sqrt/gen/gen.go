package gen

// it does range ofver n and then pass them to the channel and returns channel
func Gen(nums ...int) chan int {

	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}
