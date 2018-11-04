package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	slices := []int{1, 2, 3, 4}
	fmt.Println(strings.Join(os.Args[:], " "))
	fmt.Println(os.Args[:])
	fmt.Println(slices[:])
}
