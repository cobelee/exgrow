package testcode

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	_ "exgrow/localdb/object"
	"fmt"
)

func TestArray() {
	indicCard := h.CreateStockCard("sh600000", c.Period_D)
	indicArray := indicCard.IndicBarMatrix

	var i int
	for _, bar := range indicArray {
		fmt.Printf("%p  %s\n", &indicArray[i], bar.Date)
		i++
		if i == 5 {
			break
		}
	}

	var indicBar interface{}
	indicBar = &indicArray[0]
	fmt.Printf("%p\n", indicBar)
	indicBar = &indicArray[1]
	fmt.Printf("%p\n", indicBar)
}
