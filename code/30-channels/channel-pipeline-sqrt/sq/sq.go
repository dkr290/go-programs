package sq

func Sq(in chan int) chan int {

	out := make(chan int)
	go func() {

		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
