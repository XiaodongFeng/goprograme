package main

import (
	"fmt"
	"os"
)

func main() {
	for index, args := range os.Args[:] {
		fmt.Printf("%d: %s\n", index, args)
	}
}
