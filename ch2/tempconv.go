package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.15
	Freezing Celsius = 0
	Boiling Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
