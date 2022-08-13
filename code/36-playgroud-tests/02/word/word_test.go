package word

import (
	"fmt"
	"testing"

	"github.com/dkr290/go-programs/code/36-playgroud-tests/02/quote"
)

func TestUseCount(t *testing.T) {

	m := UseCount("one two three three four four four")

	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		case "four":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}

	}

}

func TestCount(t *testing.T) {
	n := Count("one two three")

	if n != 3 {
		t.Error("got", n, "expect", 3)
	}

}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output:
	//3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
