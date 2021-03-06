package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://raw.githubusercontent.com/AllenDowney/OlinPyShop/master/wordsEn.txt")
	if err != nil {
		fmt.Println(err)
	}
	//close response body after main function
	defer resp.Body.Close()
	words := make(map[string]string)

	//scan the page
	sc := bufio.NewScanner(resp.Body)
	//set the split function for scanning operation
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		words[sc.Text()] = ""

	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standart input:", err)
	}

	i := 0
	// prints only 200 words and not entire stuff
	for w, _ := range words {
		fmt.Println(w)
		if i == 200 {
			break
		}
		i++
	}

}
