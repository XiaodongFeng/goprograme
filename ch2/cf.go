package main

import (
	"os"
	"strconv"
	"fmt"
	"goprograme/ch2/exec2.1"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v \n", err)
			os.Exit(1)
		}

		f := tempconv2.Fahrenheit(t)
		c := tempconv2.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f,tempconv2.FToC(f) , c, tempconv2.CToF(c))
	}
}
