package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/AllenDowney/OlinPyShop/master/wordsEn.txt")
	if err != nil {
		fmt.Println(err)
	}
	//close response body after main function
	defer resp.Body.Close()

	//scan the page
	sc := bufio.NewScanner(resp.Body)
	//set the split function for scanning operation
	sc.Split(bufio.ScanLines)

	//create slice to hold scanners

	buckets := make([]int, 200)

	for sc.Scan() {
		n := HasBucket(sc.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:122])

}

func HasBucket(w string) int {
	return int(w[0])
}
