package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for file, dup := range counts {
		for line, n := range dup {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		if _, ok := counts[fileName]; !ok {
			counts[fileName] = make(map[string]int)
		}

		counts[fileName][input.Text()]++
	}
}
