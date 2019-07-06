package localdb

import (
	"fmt"
	// "time"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

func TestSDCard() {
	card := CreateSDCard("sh600000")
	scanner := NewBarScanner(card.SDDBarMatrix)

	for scanner.ScanAWeek() {
	}

	matrix := scanner.BarBuffer
	for i, m := range matrix {
		year, week := m.Date.ISOWeek()
		fmt.Printf("%v %v %v %s\n", i, year, week, m.Date.Weekday())
	}
}
