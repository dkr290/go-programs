package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	forWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	tr1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		forWheel: true,
	}

	sd1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "pink",
		},
		luxury: true,
	}

	fmt.Println(tr1)
	fmt.Println(sd1)

	fmt.Printf("The truck has %v doors and is of %v color and for wheels is:%v\n", tr1.doors, tr1.color, tr1.forWheel)
	fmt.Printf("The sedan has %v doors and is color %v and is luxuary %v\n", sd1.doors, sd1.color, sd1.luxury)
}
