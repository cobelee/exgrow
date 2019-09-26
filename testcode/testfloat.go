package testcode

import (
	"fmt"
	"math"
)

func Testfloat() {
	var ch1, c5, c1 = 0.008897, 11.11, 11.34
	var out = (ch1+1)*c5/c1 - 1

	fmt.Println(math.Round(out*1e6) / 1e6)
	// fmt.Printf("%.6f", out)
}
