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

func popcountLoop(x uint64) int {
	var sum int
	for i := 0; i<8; i++ {
		sum += int(pc[byte(x>>(uint(i)*8))])
	}
	return sum
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

func main() {
	var x uint64 =  0xfffe8731

	start := time.Now()
	for i:=0; i<10000000; i++ {
		_ = popcount(x)
	}
	fmt.Printf("popcount elapsed: %d ns\n", time.Since(start).Nanoseconds())

	start = time.Now()
	for i:=0; i<10000000; i++ {
		popcountLoop(x)
	}
	fmt.Printf("popcountLoop elapsed: %d ns\n", time.Since(start).Nanoseconds())
}
