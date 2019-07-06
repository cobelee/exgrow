package testcode

import (
	"exgrow/localdb"
	"fmt"
)

func TestSDCard() {
	card := localdb.CreateSDCard("sh600000")

	fmt.Println(card.SDDBarMatrix)

	scanner := localdb.NewBarScanner(card.SDDBarMatrix)

	fmt.Println(scanner.Bar())
	fmt.Println("------------------------")
	scanner.Scan()
	fmt.Println(scanner.Bar())
	fmt.Println("------------------------")
	scanner.ScanAWeek()
	fmt.Println(scanner.BarBuffer)
}
