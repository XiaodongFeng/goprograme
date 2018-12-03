package popcount

import (
	"goprograme/ch2/popcount"
	"fmt"
)

func main() {
	xr := popcount.PopCount(0xffffff)
	fmt.Print(xr)
}
