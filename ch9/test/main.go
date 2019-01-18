package main

import (
	"fmt"
	"time"
)

var x, y int

func main() {
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
	time.Sleep(1*time.Second)
}
