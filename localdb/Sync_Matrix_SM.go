package localdb

import (
	c "exgrow/localdb/config"
	"exgrow/localdb/dbhelp"
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
	names, _ := dbhelp.GetCollectionNames(dbName)
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
	card := CreateSDCard(code)
	scanner := NewBarsScanner(card.SDDBarMatrix)
	var sddBarArray []o.SDDBar
	sddBarArray = []o.SDDBar(card.SDDBarMatrix)
	for scanner.ScanAMonth() {
		var barBuffer []o.SDDBar
		for _, index := range scanner.BarIndexBuffer {
			sddBar := sddBarArray[index]
			barBuffer = append(barBuffer, sddBar)
		}

		if merger, e := NewBarsMerger(sddBarArray); e == nil {
			wb := merger.CreateLongPeriodBar("M")
			SaveObjToC(&wb)
		}
	}
}