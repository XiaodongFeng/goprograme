package main

import "fmt"

func unique(strs []string) []string {
	w := 0 // index of last written string
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}

func main() {
	strs := []string{"1","1","1","2","3","3","4"}
	strs = unique(strs)
	fmt.Println(strs)
}
