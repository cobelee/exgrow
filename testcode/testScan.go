package testcode

import (
	"exgrow/localdb"
	m "exgrow/localdb/maintain"
	"fmt"
)

func TestScan() {
	card := localdb.CreateIdcCard("sh600000")
	scanner := m.NewBarsScanner(card.IdcBarMatrix)

	// for i := 0; i < len(card.Matrix); i++ {
	// 	fmt.Println("%s  %v", card.Matrix[i].Date, card.Matrix[i].Close)
	// }

	// for scanner.Scan() {
	// 	scanner.CurrentIndex()

	// 	sddbar := card.SDDBarMatrix[scanner.CurrentIndex()]
	// 	fmt.Println(sddbar.Date.Format("2006-01-02"))

	// 	// if v, ok := scanner.Token.(localdb.SDDBar); ok {
	// 	// 	fmt.Println(v.Date.Format("2006-01-02"))
	// 	// }
	// }

	for scanner.ScanAMonth() {

		for i, m := range scanner.BarIndexBuffer {
			if i == 0 {
				fmt.Println(card.IdcBarMatrix[m].Date.Format("2006-01"))
			}
			fmt.Printf("%v  ", card.IdcBarMatrix[m].Date.Format("02"))

		}
		fmt.Println()
	}

	// scanner.ScanAWeek()
	// scanner.ScanAWeek()

	// matrix := scanner.BarBuffer
	// matrix[1], matrix[3] = matrix[3], matrix[1]
	// matrix[0], matrix[4] = matrix[4], matrix[0]
	// for i, m := range matrix {
	// 	// year, week := m.Date.ISOWeek()
	// 	fmt.Printf("%v %v  %s\n", i, m.Date, m.Date.Weekday())
	// 	if i == 9 {
	// 		break
	// 	}
	// }
	// fmt.Println("----------------------------")

	// sort.Sort(localdb.SortableBars(scanner.BarBuffer))
	// matrix = scanner.BarBuffer
	// for i, m := range matrix {
	// 	// year, week := m.date.ISOWeek()
	// 	fmt.Printf("%v %v  %s\n", i, m.Date, m.Date.Weekday())
	// 	if i == 9 {
	// 		break
	// 	}
	// }
}
