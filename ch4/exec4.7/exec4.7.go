package main

import (
	"unicode/utf8"
	"fmt"
)

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

// ex4.7
// Reverse all the runes, and then the entire slice. The runes' bytes end up in
// the right order.
func revUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	rev(b)
	return b
}


func main() {
	b := []byte("hello，世界")
	fmt.Println(b, string(b))
	b=revUTF8(b)
	fmt.Println(b, string(b))
}
