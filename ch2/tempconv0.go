package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100.0
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5+32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32)*5/9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	fmt.Printf("AbsoluteZero=%g°C,%g°F，Freezing=%.2f°C，%.2f°F，BoilingC=%.2f°C，%.2f°F\n",
		AbsoluteZeroC,CToF(AbsoluteZeroC),FreezingC,CToF(FreezingC),BoilingC,CToF(BoilingC))
	fmt.Print("AbsoluteZero=", AbsoluteZeroC.String())
}
