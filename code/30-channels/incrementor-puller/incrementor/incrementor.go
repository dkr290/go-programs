package incrementor

func Incr(inc int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < inc; i++ {
			c <- i
		}
		close(c)
	}()

	return c

}
