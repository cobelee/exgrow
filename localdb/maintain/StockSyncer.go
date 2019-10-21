package maintain

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"fmt"
	"time"
)

func SyncFromStockD1(pt string) {
	switch pt {
	case "w", "W", "week":
		BeginSync(o.DTW)
	case "m", "M", "month":
		BeginSync(o.DTM)
	case "q", "Q", "quarter":
		BeginSync(o.MTQ)
	case "y", "Y", "year":
		BeginSync(o.MTY)
	case "all", "All", "ALL":
		BeginSync(o.DTW)
		BeginSync(o.DTM)
		BeginSync(o.MTQ)
		BeginSync(o.MTY)
	default:
		fmt.Print("    Wrong value of flag p.\n    the available value is w, m, q or y.\n")
	}
}

/* Sync Stock database base on syncType

   Note the order of synchronization in the production environment.
   The right order is DTW, DTM, MTQ, MTY.
*/
func BeginSync(st o.SyncType) {

	dbName := c.DBConfig.DBName.StockD1
	names, _ := h.GetCollectionNames(dbName)
	cCount := len(names)
	var start, end time.Time
	start = time.Now()
	fmt.Println("Sync stock base information from database StockD1. SyncType is ", st.ToString())
	fmt.Println("Begin time: ", start.Format("2006-01-02 15:04:05"))
	for i, name := range names {
		IncrementSync(name, st)
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

// 增量同步
func IncrementSync(code string, st o.SyncType) {

	// L62-78 获取目标库已同步的状态信息。
	var destCard o.IndicCard
	switch st {
	case o.DTW:
		destCard = h.CreateStockCard(code, c.Period_W)
	case o.DTM:
		destCard = h.CreateStockCard(code, c.Period_M)
	case o.MTQ:
		destCard = h.CreateStockCard(code, c.Period_Q)
	case o.MTY:
		destCard = h.CreateStockCard(code, c.Period_Y)
	default:
		destCard = h.CreateStockCard(code, c.Period_W)
	}
	var dest1, dest2 o.IndicBarArray
	dest1 = destCard.IndicBarMatrix
	dest2 = dest1.RemoveLast(2) // 最后一周期可能是非完整周期，需剔除。
	lastDate, e := dest2.GetLastDate()

	// L81-97 在源数据库中，获取指定日期右侧的有效部分数据。
	var sourceCard o.IndicCard
	switch st {
	case o.DTW, o.DTM:
		sourceCard = h.CreateStockCard(code, c.Period_D)
	case o.MTQ, o.MTY:
		sourceCard = h.CreateStockCard(code, c.Period_M)
	default:
		sourceCard = h.CreateStockCard(code, c.Period_D)
	}
	var source1, source2 o.IndicBarArray
	source1 = sourceCard.IndicBarMatrix
	if e != nil {
		source2 = source1
	} else {
		beginDate := lastDate.AddDate(0, 0, 1)
		source2 = source1.RightSlice(beginDate)
	}

	var destArray o.IndicBarArray
	switch st {
	case o.DTW:
		destArray = TransformPeriod(source2, c.Period_W)
	case o.DTM:
		destArray = TransformPeriod(source2, c.Period_M)
	case o.MTQ:
		destArray = TransformPeriod(source2, c.Period_Q)
	case o.MTY:
		destArray = TransformPeriod(source2, c.Period_Y)
	default:
		destArray = TransformPeriod(source2, c.Period_W)
	}

	destArrayLen := len(destArray)
	for i := 0; i < destArrayLen; i++ {
		bar := destArray[i]
		h.SaveObjToC(&bar)
	}
}
