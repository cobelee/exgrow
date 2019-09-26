package localdb

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	i "exgrow/localdb/indiction"
	o "exgrow/localdb/object"
	t "exgrow/localdb/tools"
	"fmt"
)

/* Generate IndicationD1-database based on StockMarketRawD1 database.

   Indications about stock include EMA, MA, TR etc.

Steps
1. Read Stock-Code-Name from StockMarketRawD1 database
2. Generate indication-base-instance and save to 'IndicationsD1' database.
3. Calculate indication EMA and save to 'IndicationD1' database.

   Sequence to invoke method:
   GenerateIndicationD1() --> GenerateIndicBaseCollection()
*/
func GenerateIndicationD1() {

	// Get LIST of STOCK-CODE
	dbName := c.DBConfig.DBName.StockMarketRawD1
	names, _ := h.GetCollectionNames(dbName)

	// LOOP the LIST of stock-code, generate the IndicBase doc.
	cCount := len(names)
	for i, name := range names {
		GenerateIndicBaseCollection(name)
		fmt.Printf("Synchronizing indic-set collections to IndicationD1. Stock code:%s    (%d/%d) \r", name, i+1, cCount)
	}
	fmt.Println()
}

/* Generate IndicCollection and save to IndicationD1-database.

Steps
1. Get the instance of SDCard (stock-data-card)
2. New a bar scanner to daily-bar-matrix
3. LOOP the bars, then CALCULATE and FILL indic-data to the bars.
*/
func GenerateIndicBaseCollection(longCode string) {
	card := CreateSDCard(longCode)

	scanner := NewBarsScanner(card.SDDBarMatrix)
	for scanner.Scan() {
		var sddBar o.SDDBar
		sddBar = card.SDDBarMatrix[scanner.CurrentIndex()]

		var indicSet = o.NewIndicBar(sddBar)

		// EndowEMA(&scanner, &indicSet)

		// GenerateIndicSet(bar)
		h.SaveIndicSetToC(&indicSet)
		// if bar, ok := scanner.Bar(); ok {

		// }
	}
}

// Fill the collections in IndicationsD1 with EMA-indication data.
func FillIndices() {
	dbName := c.DBConfig.DBName.IndicationD1
	names, _ := h.GetCollectionNames(dbName) // Get collection-names from IndicationsD1 database
	// cCount := len(names)

	for _, name := range names {
		FillEMA(name)
		break
	}

}

func FillEMA(code string) {

	sdCard := CreateSDCard(code)
	sdArray := sdCard.SDDBarMatrix

	indicCard := CreateIndicDCard(code)
	indicArray := indicCard.IndicBarMatrix

	// 此处需完善，若两个数组没有同步一致性，需先同步。
	if t.CheckSynchronism(sdArray, indicArray) {
		fmt.Println("The both array have synchronism.")
	}

	var length = 0
	length = indicArray.Len()

	emaCalculater := i.NewEMA()
	var todayClose, lastEma, todayEma float64
	for i := 0; i < length; i++ {
		todayClose = sdArray[i].Adjust_price
		if i == 0 {
			lastEma = sdArray[i].Close
		} else {
			lastEma = indicArray[i-1].EMA39
		}

		todayEma = emaCalculater.CalcTodayValue(todayClose, lastEma)
		indicArray[i].EMA39 = todayEma
		fmt.Printf("%4d  %v\n", i, todayEma)
	}

	// scanner := NewBarsScanner(indicCard.IndicBarMatrix)
	// var bars []o.IndicBar
	// bars = []o.IndicBar(indicCard.IndicBarMatrix)

	// // ema := indiction.NewEMA()

	// row := 0
	// for scanner.Scan() {
	// 	bar := bars[scanner.CurrentIndex()]
	// 	// bar.EMA39 = ema.CalcTodayValue(bar.)
	// 	//TODO: 需要开发行情数据查询接口
	// 	fmt.Printf("%s %s %v\n", bar.Code, bar.Date.Format("2006-01-02"), bar.EMA39)
	// 	row++
	// 	if row == 20 {
	// 		break

	// 	}
	// }
}
