package maintain

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	indic "exgrow/localdb/indiction"
	o "exgrow/localdb/object"
	"fmt"
	"time"
)

// Fill the indications in IndicationsD1.
func SmartFillIndices(pt c.PeriodType) {

	dbName := c.GetStockDBNameByPeriodType(pt)
	names, _ := h.GetCollectionNames(dbName) // Get collection-names from Stock database

	// LOOP the LIST of stock-code, fill the indication in indicBar.
	cCount := len(names)
	var omArray []o.ObjectMonitor
	var start, end time.Time
	start = time.Now()
	fmt.Println("Generate indications for the database ", dbName)
	fmt.Println("Begin time: ", start.Format("15:04:05"))
	// go range 不保证执行顺序，此处先处理哪个集合并不重要，故用range
	for i, name := range names {
		indicCard := h.CreateStockCard(name, pt)
		indicArray := indicCard.IndicBarMatrix
		length := indicArray.Len()
		// om 跟踪器严重依赖顺序，不可用range,故用 经典 for 循环
		omArray = []o.ObjectMonitor{} // Clear slice, this code line is important.
		for j := 0; j < length; j++ {
			om := o.ObjectMonitor{
				Date:     indicArray[j].Date,
				Modified: false,
			}
			omArray = append(omArray, om)
		}

		SmartFillEMA(indicArray, omArray)           // Fill indication of EMA
		SmartFillTR(indicArray, omArray)            // Fill indication of TR
		SmartFillTREMA(indicArray, omArray)         // Fill indication of TREMA
		SmartSaveIndicBarArray(indicArray, omArray) //Save indication data to db
		fmt.Printf("Processing stock code: %s    (%d/%d) \r", name, i+1, cCount)
	}
	end = time.Now()
	fmt.Println(dbName, " is Completed at ", end.Format("15:04:05"))
	fmt.Println("Duration: ", end.Sub(start).String())
	fmt.Println()

}

// 将准备好数据的 IndicBarArray 保存到数据库。
func SmartSaveIndicBarArray(indicArray o.IndicBarArray, omArray []o.ObjectMonitor) {
	var length = 0
	length = indicArray.Len()

	// sort.SliceStable(indicArray, func(i int, j int) bool {
	// 	return indicArray[i].Date.Before(indicArray[j].Date)
	// })

	// sort.SliceStable(omArray, func(i int, j int) bool {
	// 	return omArray[i].Date.Before(omArray[j].Date)
	// })

	for i := 0; i < length; i++ {
		if indicArray[i].Date.Equal(omArray[i].Date) && omArray[i].Modified {
			indicBar := indicArray[i]
			h.SaveObjToC(&indicBar)
		}
	}
}

// Generate indication EMA for indicArray
func SmartFillEMA(indicArray o.IndicBarArray, omArray []o.ObjectMonitor) {
	var length = 0
	length = indicArray.Len()

	var todayClose, lastEma, todayEma float64
	for i := 0; i < length; i++ {
		if indicArray[i].EMA39 == 0 {
			todayClose = indicArray[i].Close
			if i == 0 {
				lastEma = indicArray[i].Close
			} else {
				lastEma = indicArray[i-1].EMA39
			}
			todayEma = indic.IndicCalcEMA.CalcTodayValue(todayClose, lastEma)
			indicArray[i].EMA39 = todayEma
			omArray[i].Modified = true
		}
	}

}

// Generate indication TR for indicArray
func SmartFillTR(indicArray o.IndicBarArray, omArray []o.ObjectMonitor) {
	var length = 0
	length = indicArray.Len()

	var todayHigh, todayLow, lastClose, todayTR float64
	for i := 0; i < length; i++ {
		if indicArray[i].TR == 0 {
			todayHigh = indicArray[i].High
			todayLow = indicArray[i].Low
			if i == 0 {
				lastClose = indicArray[i].Close
			} else {
				lastClose = indicArray[i-1].Close
			}

			todayTR = indic.IndicCalcTR.CalcTodayValue(todayHigh, todayLow, lastClose)
			indicArray[i].TR = todayTR
			omArray[i].Modified = true
		}
	}
}

// Generate indication TREMA for indicArray
func SmartFillTREMA(indicArray o.IndicBarArray, omArray []o.ObjectMonitor) {
	var length = 0
	length = indicArray.Len()

	var todayTR, lastEma, todayEma float64
	for i := 0; i < length; i++ {
		if indicArray[i].TREMA == 0 {
			todayTR = indicArray[i].TR
			if i == 0 {
				lastEma = indicArray[i].TR
			} else {
				lastEma = indicArray[i-1].TREMA
			}

			todayEma = indic.IndicCalcEMA.CalcTodayValue(todayTR, lastEma)
			indicArray[i].TREMA = todayEma
			omArray[i].Modified = true
		}
	}

}
