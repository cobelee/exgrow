package maintain

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	indic "exgrow/localdb/indiction"
	o "exgrow/localdb/object"
	"fmt"
	"time"
)

func FillIndications(pt string, mode string) {
	if mode == "increment" {
		switch pt {
		case "d", "D", "Day":
			SmartFillIndices(c.Period_D)
		case "w", "W", "week":
			SmartFillIndices(c.Period_W)
		case "m", "M", "month":
			SmartFillIndices(c.Period_M)
		case "q", "Q", "quarter":
			SmartFillIndices(c.Period_Q)
		case "y", "Y", "year":
			SmartFillIndices(c.Period_Y)
		case "all", "All", "ALL":
			SmartFillIndices(c.Period_D)
			SmartFillIndices(c.Period_W)
			SmartFillIndices(c.Period_M)
			SmartFillIndices(c.Period_Q)
			SmartFillIndices(c.Period_Y)
		default:
			SmartFillIndices(c.Period_D)
		}

	}

	if mode == "rerun" {
		switch pt {
		case "d", "D", "Day":
			FillIndices(c.Period_D)
		case "w", "W", "week":
			FillIndices(c.Period_W)
		case "m", "M", "month":
			FillIndices(c.Period_M)
		case "q", "Q", "quarter":
			FillIndices(c.Period_Q)
		case "y", "Y", "year":
			FillIndices(c.Period_Y)
		case "all", "All", "ALL":
			FillIndices(c.Period_D)
			FillIndices(c.Period_W)
			FillIndices(c.Period_M)
			FillIndices(c.Period_Q)
			FillIndices(c.Period_Y)
		default:
			FillIndices(c.Period_D)
		}
	}

}

// Fill the indications in IndicationsD1.
func FillIndices(pt c.PeriodType) {

	dbName := c.GetStockDBNameByPeriodType(pt)
	names, _ := h.GetCollectionNames(dbName) // Get collection-names from Stock database

	// LOOP the LIST of stock-code, fill the indication in indicBar.
	cCount := len(names)
	var start, end time.Time
	start = time.Now()
	fmt.Println("Generate indications and save to ", dbName, " database.")
	fmt.Println("Begin time: ", start.Format("15:04:05"))
	for i, name := range names {
		indicCard := h.CreateStockCard(name, pt)
		indicArray := indicCard.IndicBarMatrix

		FillEMA(indicArray)           // Fill indication of EMA
		FillTR(indicArray)            // Fill indication of TR
		FillTREMA(indicArray)         // Fill indication of TREMA
		SaveIndicBarArray(indicArray) //Save indication data to db
		fmt.Printf("Processing stock code: %s    (%d/%d) \r", name, i+1, cCount)
	}
	end = time.Now()
	fmt.Println(dbName, " is Completed at ", end.Format("2006-01-02 15:04:05"))
	fmt.Println("Duration: ", end.Sub(start).String())
	fmt.Println()

}

// Generate indication EMA for indicArray
func FillEMA(indicArray o.IndicBarArray) {
	var length = 0
	length = indicArray.Len()

	var todayClose, lastEma, todayEma float64
	for i := 0; i < length; i++ {
		todayClose = indicArray[i].Close
		if i == 0 {
			lastEma = indicArray[i].Close
		} else {
			lastEma = indicArray[i-1].EMA39
		}

		todayEma = indic.IndicCalcEMA.CalcTodayValue(todayClose, lastEma)
		indicArray[i].EMA39 = todayEma
		// fmt.Printf("%4d  %s  %v\n", i, indicArray[i].Date, todayEma)
	}
}

// Generate indication TR for indicArray
func FillTR(indicArray o.IndicBarArray) {
	var length = 0
	length = indicArray.Len()

	var todayHigh, todayLow, lastClose, todayTR float64
	for i := 0; i < length; i++ {
		todayHigh = indicArray[i].High
		todayLow = indicArray[i].Low
		if i == 0 {
			lastClose = indicArray[i].Close
		} else {
			lastClose = indicArray[i-1].Close
		}

		todayTR = indic.IndicCalcTR.CalcTodayValue(todayHigh, todayLow, lastClose)
		indicArray[i].TR = todayTR
	}
}

// Generate indication TREMA for indicArray
func FillTREMA(indicArray o.IndicBarArray) {
	var length = 0
	length = indicArray.Len()

	var todayTR, lastEma, todayEma float64
	for i := 0; i < length; i++ {
		todayTR = indicArray[i].TR
		if i == 0 {
			lastEma = indicArray[i].TR
		} else {
			lastEma = indicArray[i-1].TREMA
		}

		todayEma = indic.IndicCalcEMA.CalcTodayValue(todayTR, lastEma)
		indicArray[i].TREMA = todayEma
	}

}
