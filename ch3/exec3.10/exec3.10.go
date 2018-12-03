package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	b := &bytes.Buffer{}
	pre := len(s) % 3

	if pre == 0 {
		pre = 3
	}

	b.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}

func main() {
	s := "1234567890"
	fmt.Println(comma(s))
	s = "234567890"
	fmt.Println(comma(s))
	s = "11234567890"
	fmt.Println(comma(s))
}
