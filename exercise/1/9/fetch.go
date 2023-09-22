package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const protocol = "https://"

func getURL(url string) string {
	if strings.HasPrefix(url, protocol) {
		return url
	}

	return protocol + url
}

func main() {
	for _, url := range os.Args[1:] {
		u := getURL(url)

		resp, err := http.Get(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Status: %s\n", resp.Status)
		bytes, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil || bytes == 0 {
			fmt.Fprintf(os.Stderr, "fetch: coping %v. copied %d byte(s)\n", err, bytes)
			os.Exit(1)
		}
	}
}
