package maintain

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"fmt"
	"time"
)

/* Sync Stock database RawD1 to StockD1

 */
func BeginSyncRTD() {
	dbNameRaw := c.DBConfig.DBName.StockMarketRawD1
	namesRaw, _ := h.GetCollectionNames(dbNameRaw)

	cCount := len(namesRaw)
	var start, end time.Time
	start = time.Now()
	fmt.Println("Sync database StockMarketRawD1 to StockD1.")
	fmt.Println("Begin time: ", start.Format("2006-01-02 15:04:05"))
	for i, name := range namesRaw {
		IncrementSyncRTD(name)
		fmt.Printf("Stock code: %s    (%d/%d) \r", name, i+1, cCount)
	}
	end = time.Now()
	fmt.Println("Synchronization Completed at ", end.Format("2006-01-02 15:04:05"))
	fmt.Println("Duration: ", end.Sub(start).String())
	fmt.Println()
}

/* Merge Daily-Matrix to Weekly-Matrix

Steps
1. Get the instance of stock-data-card
2. New a bar scanner of daily-bar-matrix
3. merge daily-bars to weekly-bar
*/
// func MergeDMtoWM(code string) {
// 	card := h.CreateIndicCard(code, o.Period_D)
// 	//scanner := NewSDDBarScanner(card.SDDBarMatrix)
// 	scanner := NewBarsScanner(card.IndicBarMatrix)

// 	for scanner.ScanAWeek() {
// 		var barBuffer []o.IndicBar
// 		barBuffer = card.Matrix[scanner.BufferBoards.LeftLimit:scanner.BufferBoards.RightLimit]

// 		if merger, e := NewBarsMerger(barBuffer); e == nil {
// 			wb := merger.CreateLongPeriodBar("W")
// 			h.SaveObjToC(&wb)
// 		}
// 	}
// }

// 仅同步增量部分
// 同步结束，不保证两库完全一致
func IncrementSyncRTD(code string) {

	// L62-78 获取目标库已同步的状况信息。
	var sourceCard o.SDCard
	var destCard o.IndicCard

	sourceCard = h.CreateSDCard(code)
	destCard = h.CreateStockCard(code, c.Period_D)

	sdBarArray := sourceCard.Matrix
	indicBarArray := destCard.IndicBarMatrix

	var countSource, countDest int
	countSource = len(sdBarArray)
	countDest = indicBarArray.Len()

	if countSource <= countDest {
		return
	}

	// 未来代码优化，考虑将以下循环单独开发库差异检测功能。
	// 此处仅考虑增量同步。
	beginSyncIndex := countDest
	for i := 0; i < countDest; i++ {
		if !sdBarArray[i].Date.Equal(indicBarArray[i].Date) {
			beginSyncIndex = i
		}
	}

	sdBarArraySource := sdBarArray[beginSyncIndex:]
	length := len(sdBarArraySource)
	for i := 0; i < length; i++ {
		indicBar := o.NewIndicBarFromSDBar(sdBarArraySource[i])
		h.SaveObjToC(&indicBar)
	}

}
