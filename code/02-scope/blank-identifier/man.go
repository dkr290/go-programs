package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//no error checking wih blank identifer
	res, _ := http.Get("http://www.geekwiseacademy.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
