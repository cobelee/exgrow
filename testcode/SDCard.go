package testcode

import (
	_ "exgrow/localdb"
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"fmt"
)

func TestSDCard() {
	card := h.CreateStockCard("sh600000", c.Period_D)
	var slice1, slice2 o.IndicBarArray
	slice1 = card.IndicBarMatrix[0:20]

	for i, s := range slice1 {
		fmt.Printf("%v  %v\n", i, s)
	}
	fmt.Println()

	slice2 = slice1.RemoveLast(2)

	for i, s := range slice2 {
		fmt.Printf("%v  %v\n", i, s)
	}

	fmt.Println()

	fmt.Println(slice2.GetLastDate())
}
