package main

import (
	"strings"
	"fmt"
)

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	// check if '.' can find
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	s:= "abcd"
	s = basename2(s)
	fmt.Println(s)
}
