package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var total int
			for x := 0; x <= i; x++ {
				total += x
			}
			sum <- total
		}(i)
	}

	go func() {
		wg.Wait()
		close(sum)
	}()

	var total int
	for size := range sum {
		fmt.Println(size)
		total += size
	}
	fmt.Println("done")
}
