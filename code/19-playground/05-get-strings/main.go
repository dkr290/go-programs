package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		fmt.Println(err)
	}

	// scan the page

	scanner := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	// Set the splic function for the scanniong operation

	scanner.Split(bufio.ScanWords)
	//create slice of slice to hold slices of words

	buckets := make([][]string, 12)
	fmt.Println(buckets)

	//create slices to hold words

	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	fmt.Println(buckets)

	// over the words

	for scanner.Scan() {
		word := scanner.Text()
		n := HasBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	//print the word in one of the buckets
	//fmt.Println(buckets[8])
}

func HasBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)

	}
	return sum % buckets

}
