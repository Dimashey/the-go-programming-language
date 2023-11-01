package main

import (
	"bytes"
	"fmt"
)

const (
	commaUniCode    = 44 // decimal number of "," character
	fullStopUniCode = 46 // decimal number of "." character
)

func comma(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		if i%3 == 0 && i != 0 {
			buf.WriteByte(commaUniCode)
		}

		buf.WriteByte(s[i])

		if s[i] == fullStopUniCode {
			buf.Write([]byte(s[i+1:]))

			break
		}
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("12345.6789"))
}
