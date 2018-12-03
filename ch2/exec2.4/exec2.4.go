package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popcountshiftvalue(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
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
	var x uint64 = 0xFF0A0011
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		popcount(x)
	}
	fmt.Printf("popcount elapsed: %d ns\n", time.Since(start).Nanoseconds())

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		popcountshiftvalue(x)
	}
	fmt.Printf("popcountshiftvalue elapsed: %d ns\n", time.Since(start).Nanoseconds())
}
