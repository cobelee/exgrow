package localdb

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	m "exgrow/localdb/maintain"
	o "exgrow/localdb/object"
	"fmt"
)

/* Sync Daily-Stock database to Monthly-Stock database

Steps
1. Read Stock-Code-Name from StockMartketRawD1 database
2. Merge Daily-Matrix to Monthly-Matrix
*/
func SyncToSMdb() {

	dbName := c.DBConfig.DBName.StockMarketRawD1
	names, _ := h.GetCollectionNames(dbName)
	cCount := len(names)

	for i, name := range names {
		MergeDMtoMM(name)
		fmt.Printf("Synchronizing daily-bar to monthly-bar. Stock code:%s    (%d/%d) \r", name, i+1, cCount)
	}
	fmt.Println()
}

/* Merge Daily-Matrix to Weekly-Matrix

Steps
1. Get the instance of stock-data-card
2. New a bar scanner of daily-bar-matrix
3. merge daily-bars to weekly-bar
*/
func MergeDMtoMM(code string) {
	card := h.CreateSDCard(code)
	scanner := m.NewBarsScanner(card.SDDBarMatrix)
	// var sddBarArray []o.SDDBar
	// sddBarArray = []o.SDDBar(card.SDDBarMatrix)
	for scanner.ScanAMonth() {
		var barBuffer []o.SDDBar
		barBuffer = card.SDDBarMatrix[scanner.BufferBoards.LeftLimit:scanner.BufferBoards.RightLimit]

		if merger, e := m.NewBarsMerger(barBuffer); e == nil {
			wb := merger.CreateLongPeriodBar("M")
			h.SaveObjToC(&wb)
		}
	}
}
