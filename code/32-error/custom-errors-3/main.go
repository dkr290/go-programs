package main

import (
	"fmt"
	"log"
)

func main() {

	_, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error) {

	if f < 0 {
		ErrNortgateMath := fmt.Errorf("nordgate math again: square root of negative number: %v", f)
		return 0, ErrNortgateMath
	}

	return 42, nil

}
