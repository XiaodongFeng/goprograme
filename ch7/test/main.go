package main

import "fmt"

func test(x interface{}) {
	fmt.Println(x)
}
func main() {
	test(1)
}
