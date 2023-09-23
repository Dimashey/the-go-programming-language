package main

import (
	"fmt"
	"io"
	"net/http"
	u "net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	parsedURL, err := u.Parse(url)
	if err != nil {
		ch <- "Invalid URL"
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	file, err := os.Create(fmt.Sprintf("%s.html", parsedURL.Host))
	if err != nil {
		ch <- "Error while creating file"
	}

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	file.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
