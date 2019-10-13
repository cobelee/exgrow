package maintain

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"fmt"
	"time"
)

/* Generate StockD1-database based on StockMarketRawD1 database.

   Indications include EMA, MA, TR etc.

Steps
1. Read Stock-Code-Name from StockMarketRawD1 database
2. Generate basic-adjusted-price and save to 'StockD1' database.
3. Calculate indication EMA、TR、TREMA and save to 'StockD1' database.

   Sequence to invoke method:
   GenerateStockD1() --> GenerateStockBaseCollection()
*/
func GenerateStockD1() {

	// Get LIST of STOCK-CODE
	dbName := c.DBConfig.DBName.StockMarketRawD1
	names, _ := h.GetCollectionNames(dbName)

	// LOOP the LIST of stock-code, generate the StockBase doc.
	cCount := len(names)
	var start, end time.Time
	start = time.Now()
	fmt.Println("Generate basic database --- StockD1 .")
	fmt.Printf("Begin time: %s\n", start.Format("15:04:05"))
	for i, name := range names {
		indicArray := GenerateStockD1Collection(name)

		FillEMA(indicArray)   // Fill indication of EMA
		FillTR(indicArray)    // Fill indication of TR
		FillTREMA(indicArray) // Fill indication of TREMA

		SaveIndicBarArray(indicArray)

		fmt.Printf("Processing stock code:%s    (%d/%d) \r", name, i+1, cCount)
	}
	fmt.Printf("Completed at  %s.\n", end.Format("2006-01-02 15:04:05"))
	fmt.Println("Duration: ", end.Sub(start).String())
	fmt.Println()
}

/* Generate Collection and save to StockD1-database.

Steps
1. Get the instance of SDCard (stock-data-card)
2. New a bar scanner to matrix
3. LOOP the bars, then CALCULATE and FILL indic-data to the array.
*/
func GenerateStockD1Collection(longCode string) o.IndicBarArray {
	card := h.CreateSDCard(longCode)
	scanner := NewBarsScanner(o.SDBarArray(card.Matrix))
	var indicArray o.IndicBarArray
	for scanner.Scan() {
		sdBar := card.Matrix[scanner.CurrentIndex()]

		var indicSet = o.NewIndicBarFromSDBar(sdBar)
		indicArray = append(indicArray, indicSet)
	}
	return indicArray
}

// 将准备好数据的 IndicBarArray 保存到数据库。
func SaveIndicBarArray(indicArray o.IndicBarArray) {
	var length = 0
	length = indicArray.Len()

	for i := 0; i < length; i++ {
		indicBar := indicArray[i]
		h.SaveObjToC(&indicBar)
	}
}
