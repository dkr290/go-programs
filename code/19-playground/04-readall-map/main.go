package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/AllenDowney/OlinPyShop/master/wordsEn.txt")
	if err != nil {
		fmt.Println(err)
	}
	//close response body after main function
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("%s", body)
}
