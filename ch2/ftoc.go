package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fTOC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fTOC(boilingF))
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(*f())
	fmt.Println(*f())
}

func fTOC(f float64) float64 {
	return (f-32)*5/9
}

func f() *int {
	v := 1
	return &v
}