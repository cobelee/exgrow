package localdb

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	m "exgrow/localdb/maintain"
	o "exgrow/localdb/object"
	"fmt"
	// "time"
)

// Sync Daily-Stock database to Weekly-Stock database
func SyncToSWdb() {

	dbName := c.DBConfig.DBName.StockMarketRawD1
	names, _ := h.GetCollectionNames(dbName)
	cCount := len(names)

	for i, name := range names {
		MergeDMtoWM(name)
		fmt.Printf("Synchronizing daily-bar to weekly-bar. Stock code:%s    (%d/%d) \r", name, i+1, cCount)
	}
	fmt.Println()
}

/* Merge Daily-Matrix to Weekly-Matrix

Steps
1. Get the instance of stock-data-card
2. New a bar scanner of daily-bar-matrix
3. merge daily-bars to weekly-bar
*/
func MergeDMtoWM(code string) {
	card := h.CreateSDCard(code)
	//scanner := NewSDDBarScanner(card.SDDBarMatrix)
	scanner := m.NewBarsScanner(card.SDDBarMatrix)
	for scanner.ScanAWeek() {
		var barBuffer []o.SDDBar
		barBuffer = card.SDDBarMatrix[scanner.BufferBoards.LeftLimit:scanner.BufferBoards.RightLimit]

		if merger, e := m.NewBarsMerger(barBuffer); e == nil {
			wb := merger.CreateLongPeriodBar("W")
			h.SaveObjToC(&wb)
		}
	}
}

// code := "sh600000"
// card := CreateSDCard(code)
// scanner := NewBarScanner(card.SDDBarMatrix)

// for scanner.ScanAWeek() {
// 	scanner.ScanAWeek()

// 	if proc, e := NewBarsMerger(scanner.BarBuffer); e == nil {
// 		wBar := proc.CreateLongPeriodBar("W")
// 		fmt.Println(wBar)
// 	}

// 	break
// }
// fmt.Println("----------------------------------\n")

// matrix := scanner.BarBuffer
// for _, m := range matrix {
// 	// year, week := m.Date.ISOWeek()
// 	// fmt.Printf("%v %v %v %s\n", i, year, week, m.Date.Weekday())
// 	fmt.Println(m)
// }
// fmt.Println("----------------------------------\n")
// scanner.ScanAWeek()
// if proc, e := NewBarsMerger(scanner.BarBuffer); e == nil {
// 	wBar := proc.CreateLongPeriodBar("W")
// 	fmt.Println(wBar)
// }
// fmt.Println("----------------------------------\n")

// matrix = scanner.BarBuffer
// for _, m := range matrix {
// 	// year, week := m.Date.ISOWeek()
// 	// fmt.Printf("%v %v %v %s\n", i, year, week, m.Date.Weekday())
// 	fmt.Println(m)
// }
