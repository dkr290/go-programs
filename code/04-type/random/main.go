package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)

	}

	for i := 0; i < 5; i++ {
		a := rand.Intn(8)
		fmt.Printf("%d\n", a)

	}

	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f /", 100*rand.Float32())
	}
}
