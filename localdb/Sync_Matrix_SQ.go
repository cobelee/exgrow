package localdb

import (
	c "exgrow/localdb/config"
	"exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"fmt"
)

/* Sync Monthly-Stock database to Quarterly-Stock database

Steps
1. Read Stock-Code-Name from StockMartketM1 database
2. Merge Monthly-Matrix to Quarterly-Matrix
*/
func SyncToSQdb() {

	dbName := c.DBConfig.DBName.StockMarketM1
	names, _ := dbhelp.GetCollectionNames(dbName)
	cCount := len(names)

	for i, name := range names {
		MergeMMtoQM(name)
		fmt.Printf("Synchronizing monthly-bar to quarterly-bar. Stock code:%s    (%d/%d) \r", name, i+1, cCount)
	}
	fmt.Println()
}

/* Merge monthly-Matrix to quarterly-Matrix

Steps
1. Get the instance of stock-data-card
2. New a bar scanner of daily-bar-matrix
3. merge daily-bars to weekly-bar
*/
func MergeMMtoQM(longCode string) {
	card := CreateSMCard(longCode)
	var sddBarArray SDDBarArray
	sddBarArray = card.Matrix
	scanner := NewBarsScanner(sddBarArray)
	for scanner.ScanAQuarter() {
		var barBuffer []o.SDDBar
		barBuffer = card.Matrix[scanner.BufferBoards.LeftLimit:scanner.BufferBoards.RightLimit]

		if merger, e := NewBarsMerger(barBuffer); e == nil {
			wb := merger.CreateLongPeriodBar("Q")
			SaveObjToC(&wb)
		}
	}
}
