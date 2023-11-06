package main

import "fmt"

func rotate(s []int, x int) {
	l := len(s)

	tmp := make([]int, x)
	copy(tmp, s[:x])
	copy(s, s[x:])
	copy(s[l-x:], tmp)
}

func main() {
	s := []int{1, 2, 3, 4}
	rotate(s, 2)
	fmt.Println(s)
}
