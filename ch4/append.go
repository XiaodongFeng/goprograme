package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendNInt(x []int, y ... int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z=x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func main() {
	x := make([]int, 0)
	for i := 0; i < 10; i++ {
		x = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(x), x)
	}

	y := make([]int, 0)
	for i := 0; i < 10; i++ {
		y = appendNInt(y, i, i*i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	}
}
