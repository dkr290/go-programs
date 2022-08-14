package gen

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {

	c := Gen(3)

	for n := range c {
		if n != 3 {
			t.Error("got", n, "expected", 3)
		}
	}

}

func ExampleGen() {
	c := Gen(3)
	for n := range c {
		fmt.Println(n)
	}
}

func BenchmarkGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gen(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}
