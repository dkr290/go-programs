package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		fmt.Println(err)
		return

	}

	fmt.Println(string(bs))

}

func toJSON(a interface{}) ([]byte, error) {

	bs, err := json.Marshal(a)
	if err != nil {
		//return []byte{}, fmt.Errorf("There was and error in to json %v", err)
		return []byte{}, errors.New(fmt.Sprintf("There was an error in the json %v", err))
	}
	return bs, nil

}
