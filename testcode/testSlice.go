package testcode

import (
	"fmt"
)

func Testslice() {
	var a = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	index := 5
	start := index - 4
	end := index - 3
	b := a[start:end]

	fmt.Println(b)
}
