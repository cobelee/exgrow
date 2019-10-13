package testcode

import (
	"exgrow/localdb"
	h "exgrow/localdb/dbhelp"
	m "exgrow/localdb/maintain"
	"fmt"
	"sort"
)

func TestSort() {
	card := h.CreateSDCard("sh600000")
	scanner := localdb.NewSDDBarScanner(card.SDDBarMatrix)

	scanner.ScanAWeek()
	scanner.ScanAWeek()

	matrix := scanner.BarBuffer
	matrix[1], matrix[3] = matrix[3], matrix[1]
	matrix[0], matrix[4] = matrix[4], matrix[0]
	for i, m := range matrix {
		// year, week := m.Date.ISOWeek()
		fmt.Printf("%v %v  %s\n", i, m.Date, m.Date.Weekday())
		if i == 9 {
			break
		}
	}
	fmt.Println("----------------------------")

	sort.Sort(m.SortableBars(scanner.BarBuffer))
	matrix = scanner.BarBuffer
	for i, m := range matrix {
		// year, week := m.date.ISOWeek()
		fmt.Printf("%v %v  %s\n", i, m.Date, m.Date.Weekday())
		if i == 9 {
			break
		}
	}
}
