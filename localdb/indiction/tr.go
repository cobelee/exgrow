package indiction

import (
	"math"
)

// -----------------------------------------------------------------------------

// 波动指标
type TR1 struct {
}

/* Calulate today's tr value.
   Formula: TR1 := MAX(MAX((HIGH-LOW),ABS(REF(CLOSE,1)-HIGH)),ABS(REF(CLOSE,1)-LOW));
*/
func (this TR1) CalcTodayValue(todayHigh, todayLow float64, yesterdayClose float64) float64 {
	var c1, c2, c3, max float64
	c1 = todayHigh - todayLow
	c2 = math.Abs(yesterdayClose - todayHigh)
	c3 = math.Abs(yesterdayClose - todayLow)
	max = math.Max(math.Max(c1, c2), c3)

	return max
}

// -----------------------------------------------------------------------------

func NewTR1() TR1 {
	return TR1{}
}

// -----------------------------------------------------------------------------

var IndicCalcTR TR1

func init() {
	IndicCalcTR = NewTR1()
}
