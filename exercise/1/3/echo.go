package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo() {
	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func optimizedEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	echo()
	fmt.Println(time.Since(start).Milliseconds())

	start = time.Now()
	optimizedEcho()
	fmt.Println(time.Since(start).Microseconds())
}
