package indiction

// -----------------------------------------------------------------------------

/* 指数移动平均值

   指数移动平均，又叫平滑移动平均
   用法：Y=EMA(X,N),求X的N日指数移动平均
   算法：若Y=EMA(X,N),则Y=(1/N)*X + (1-1/N)*Y', 其中Y'表示上一周期Y值。

   带权重的算法：
   若Y=EMA(X,N,W),则Y=[(1+W)*X + (N-1)*Y'] / (N+W), 其中Y'表示上一周期Y值。
*/
type EMA struct {
	N int // 值的个数, 默认39
	W int // 当前周期值的权重，默认为1。若为 0 表示不加权重。建议权重值不可过大。
}

// Calulate today's ema39 value.
func (this EMA) CalcTodayValue(x float64, yesterdayEMA float64) float64 {
	var todayValue float64
	if yesterdayEMA == 0 {
		return x
	}

	todayValue = (float64(1+this.W)*x + float64(this.N-1)*yesterdayEMA) / float64(this.N+this.W)
	return todayValue
}

// -----------------------------------------------------------------------------

func NewEMA() EMA {
	return EMA{
		N: 39,
		W: 1,
	}
}

// -----------------------------------------------------------------------------

var IndicCalcEMA EMA

func init() {
	IndicCalcEMA = NewEMA()
}
