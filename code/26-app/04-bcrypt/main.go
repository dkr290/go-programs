package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPassword := `password123`
	err1 := bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err1 != nil {
		fmt.Println("You can't log in", err1)
		return
	}
	fmt.Println("You are logged in")
}
