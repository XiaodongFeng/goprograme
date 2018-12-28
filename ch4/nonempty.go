package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

	data1 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data1)) // `["one" "three"]`
	fmt.Printf("%q\n", data1)           // `["one" "three" "three"]`
}