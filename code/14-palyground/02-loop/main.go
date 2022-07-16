package main

import (
	"fmt"
	"time"
)

func main() {
	bd := 1985
	currentTime := time.Now()

	y := currentTime.Year()

	for bd <= y {

		fmt.Println("Year", bd)

		bd++
	}

}
