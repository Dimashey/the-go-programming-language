package main

import (
	"fmt"
)

func unique(strs []string) []string {
	lastUniqueAddedIndex := 0

	for _, s := range strs {
		if strs[lastUniqueAddedIndex] == s {
			continue
		}

		lastUniqueAddedIndex++
		strs[lastUniqueAddedIndex] = s
	}

	return strs[:lastUniqueAddedIndex+1]
}

func main() {
	fmt.Println(unique([]string{"a", "a", "b", "c", "d", "d", "a"}))
}
