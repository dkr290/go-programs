package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {

	num := Years(20)
	if num != 140 {
		t.Error("expected", 140, "got", num)
	}
}

func TestYearsTwo(t *testing.T) {

	num := YearsTwo(2)
	if num != 14 {
		t.Error("expected", 14, "got", num)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}

}

func BenchmarkYearsTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}

}

func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70

}

func ExampleYearsTwo() {

	fmt.Println(YearsTwo(10))
	//Output:
	//70

}
