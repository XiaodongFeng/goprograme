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
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "exec1.8: url: %s, %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
