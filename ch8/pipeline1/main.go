package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
			if x==10 {
				close(naturals)
				break
			}
		}
	}()

	go func() {
		for {
			x,ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
		time.Sleep(1 * time.Second)
	}
}
