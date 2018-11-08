package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"io"
	"io/ioutil"
)


func main() {
	ch := make(chan string)
	start := time.Now()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("main: %.4fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetch url %s error, %v\n", url, err)
		return
	}
	secs:= time.Since(start).Seconds()
	ch <- fmt.Sprintf("fetch url %s, tocal %7d bytes, elapsed %.4fs", url, nbytes, secs)
}