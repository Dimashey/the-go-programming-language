package main

import "fmt"

func isAnogram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var sourceAlphabetCounter [26]int
	var targetAlphabetCounter [26]int

	for i := 0; i < len(s1); i++ {
		sourceAlphabetCounter[s1[i]-'a']++
		targetAlphabetCounter[s2[i]-'a']++
	}

	for j := 0; j < len(sourceAlphabetCounter); j++ {
		if sourceAlphabetCounter[j] != targetAlphabetCounter[j] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnogram("asd", "das"))
}
