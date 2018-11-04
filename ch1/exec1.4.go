package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	if len(files) == 0 {
		countLine(os.Stdin, counts, foundIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLine(f, counts, foundIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%v\t%s\t%d\n", foundIn[line], line, n)
		}
	}
}

func countLine(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !in(f.Name(), foundIn[line]) {
			foundIn[input.Text()] = append(foundIn[line], f.Name())
		}
	}
}

func in(filename string, strings []string) bool {
	for _,s := range strings {
		if s==filename {
			return true
		}
	}
	return false
}
