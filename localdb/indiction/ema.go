package indiction

// -----------------------------------------------------------------------------

// 指数移动平均值
type EMA struct {
	N int // sadfqewf
}

// Calulate today's ema39 value.
func (this EMA) CalcTodayValue(todayClose float64, yesterdayEMA float64) float64 {
	var todayValue float64
	if yesterdayEMA == 0 {
		return todayClose
	}

	todayValue = 2*todayClose/float64(this.N+1) + float64(this.N-1)*yesterdayEMA/float64(this.N+1)
	return todayValue
}

// -----------------------------------------------------------------------------

func NewEMA() EMA {
	return EMA{
		N: 39,
	}
}
