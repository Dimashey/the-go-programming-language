package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()

		if word == "done" {
			break
		}

		wordfreq[word]++
	}

	fmt.Printf("word\tcount\n")

	for w, c := range wordfreq {
		fmt.Printf("%s\t%d\n", w, c)
	}
}
