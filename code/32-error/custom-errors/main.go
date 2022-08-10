package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("names.txt")
	if err != nil {
		//fmt.Println("error happened", err)
		log.Println("error happened", err)
		//log.Fatal()
		//log.Panic()
		//panic(err)
		return
	}
	defer f2.Close()

	bs, err := ioutil.ReadAll(f2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
