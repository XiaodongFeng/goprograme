package main

import (
	"fmt"
	"os"
	"log"
)

var cwd string
func init() {
	cwd, err := os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y) // compile error: x and y are not visible here

	var x uint8 = 101   //01100101
	var y uint8 = 98  //01100010
	fmt.Printf("%08b\n", x&^y)

	var f float32 = 16777216 // 1 << 24
	fmt.Println(f+1)
	fmt.Println(f == f+1) // "true"!
}

func f() int {
	return 1
}

func g(int) int {
	return 2
}
