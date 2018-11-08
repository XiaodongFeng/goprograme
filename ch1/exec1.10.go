package main

import (
	"os"
	"fmt"
	"net/url"
	"net/http"
	"io"
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
	fmt.Printf("main: %.4f\n", time.Since(start).Seconds())
}

func fetch (uri string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	f, err := os.Create(url.QueryEscape(uri))
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()

	if closeError := f.Close(); err == nil {
		err = closeError
	}
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("fetch:url %v, %dbytes, %.4fsec", uri, nbytes, time.Since(start).Seconds())
}
