package main

import (
	"time"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popcount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popcountclearrightmost(x uint64) int {
	count := 0
	for x!=0 {
		x &= x-1
		count++
	}
	return count
}


func main() {
	var x uint64 = 0xFFE09D34
	start := time.Now()
	for i:=0; i< 1000000; i++ {
		popcount(x)
	}
	fmt.Printf("popcount elapsed %dns\n", time.Since(start).Nanoseconds())

	start = time.Now()
	for i:=0; i< 1000000; i++ {
		popcountclearrightmost(x)
	}
	fmt.Printf("popcountclearrightmost elapsed %dns\n", time.Since(start).Nanoseconds())
}
