package main

import (
"fmt"
"os"
"net/http"
"strings"
"io/ioutil"
)

func main() {
	for _,url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exec1.8: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		s :=resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "exec1.8: status: %s, url: %s, %v\n", s, url, err)
			os.Exit(1)
		}
		fmt.Printf("status: %s, body: %s\n", s, b)
	}
}
