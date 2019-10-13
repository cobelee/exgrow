package maintain

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	"fmt"
	"time"
)

func RemoveStockRecordSinceDate(t time.Time) {

	removeFromStock(c.Period_D, t)
	removeFromStock(c.Period_W, t)
	removeFromStock(c.Period_M, t)
	removeFromStock(c.Period_Q, t)
	removeFromStock(c.Period_Y, t)
}

func removeFromStock(pt c.PeriodType, t time.Time) {
	// Get LIST of STOCK-CODE
	dbName := c.GetStockDBNameByPeriodType(pt)
	names, _ := h.GetCollectionNames(dbName)

	// LOOP the LIST of stock-code, generate the StockBase doc.
	cCount := len(names)
	var start, end time.Time
	start = time.Now()
	fmt.Println("Removing dbDoc from database ", dbName, ".")
	fmt.Printf("Begin time: %s\n", start.Format("15:04:05"))
	for i, name := range names {
		h.RemoveAll(dbName, name, t)
		fmt.Printf("Processing stock code:%s    (%d/%d) \r", name, i+1, cCount)
	}
	end = time.Now()
	fmt.Printf("Completed at  %s.\n", end.Format("2006-01-02 15:04:05"))
	fmt.Println("Duration: ", end.Sub(start).String())
	fmt.Println()
}
