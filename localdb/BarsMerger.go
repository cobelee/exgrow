// BarsBuffer 处理器
package localdb

import (
	"exgrow/errors"
	o "exgrow/localdb/object"
	"math"
	"sort"
	"time"
)

// 创建多巴合成器
func NewBarsMerger(m []o.SDDBar) (BarsMerger, error) {
	var merger BarsMerger

	merger.Bars = m
	sort.Sort(SortableBars(m)) // 对BarBuffer，按Bar.Date日期从小到大的顺序进行排列。

	merger.Length = len(m)

	if m == nil || len(m) == 0 {
		err := errors.NewError("100009", "No bars data.")
		return merger, &err
	} else {
		return merger, nil
	}
}

// 多巴合成器
type BarsMerger struct {
	Length int
	Bars   []o.SDDBar // 处理对象
}

/* Merge original short term bars to a long period bar.
 *
 * periodType may be "W" ------ Week
 *                   "M" ------ Month
 *                   "Q" ------ Quarter
 *                   "Y" ------ Year
 */
func (this BarsMerger) CreateLongPeriodBar(periodType string) o.SDDBar {
	var longBar = o.SDDBar{}
	longBar.Code = this.GetCode()
	longBar.Date = this.GetDate()
	longBar.Open = this.GetOpen()
	longBar.High = this.GetHigh()
	longBar.Low = this.GetLow()
	longBar.Close = this.GetClose()
	longBar.Change = this.GetChange()
	longBar.Volume = this.GetVolume()
	longBar.Money = this.GetMoney()
	longBar.Traded_market_value = this.GetTradedMarketValue()
	longBar.Market_value = this.GetMarketValue()
	longBar.Turnover = this.GetTurnover()
	longBar.Adjust_price = this.GetAdjustPrice()
	longBar.Adjust_price_f = this.GetAdjustPriceF()
	longBar.Trade_days = this.GetTradeDays()
	longBar.Period = periodType
	return longBar
}

// Get the code of stock.
func (this BarsMerger) GetCode() string {
	return this.Bars[0].Code
}

// Get the last bar's date in the period.
func (this BarsMerger) GetDate() time.Time {
	return this.Bars[this.Length-1].Date
}

// Get the highest price in the period.
func (this BarsMerger) GetHigh() float64 {
	var v float64
	for _, b := range this.Bars {
		if b.High > v {
			v = b.High
		}
	}
	return v
}

// Get low price in the current period
func (this BarsMerger) GetLow() float64 {
	var v float64
	for i, b := range this.Bars {
		if i == 0 {
			v = b.Low
		} else {
			if v > b.Low {
				v = b.Low
			}
		}
	}
	return v
}

// Get open price in this period.
func (this BarsMerger) GetOpen() float64 {
	if this.Length == 0 {
		return 0
	} else {
		return this.Bars[0].Open
	}
}

// Get close price in the period.
func (this BarsMerger) GetClose() float64 {
	if this.Length == 0 {
		return 0
	} else {
		return this.Bars[this.Length-1].Close
	}
}

/* Get the price change in the period.
 *
 * The change means 涨跌幅
 *
 * 1. 外部计算法
 *    （本周期收盘价 - 上周期收盘价）/ 上周期收盘价
 *     公式表示 C5 - C0 / C0
 * 2. 内部计算法
 *    利用本周期首日的收盘价和涨跌幅数据，倒算出上周期收盘价，然后对外部计算法公式进行变化求值。
 *    公式表示为：（CH1+1）* C5/C1 - 1
 *       CH1: 	本周期首日的涨跌幅
 *       C1:		本周期首日的收盘价
 *       C5:		本周期末日的收盘价
 */
func (this BarsMerger) GetChange() float64 {
	var ch1, c1, c5 float64
	ch1 = this.Bars[0].Change
	c1 = this.Bars[0].Close
	c5 = this.Bars[this.Length-1].Close

	out := (ch1+1)*c5/c1 - 1
	return math.Round(out*1e6) / 1e6
}

/* Get the total volume in this period.
 *
 * 算法： 本周期内 volume 总和。
 */
func (this BarsMerger) GetVolume() float64 {
	var total float64
	for _, b := range this.Bars {
		total += b.Volume
	}
	return total
}

/* Get the total money in this period.
 *
 * 算法： 本周期内 成交额 总和。
 * 注意： 不能简单得按本周期成交量与期末股价相乘计算。
 */
func (this BarsMerger) GetMoney() float64 {
	var total float64
	for _, b := range this.Bars {
		total += b.Money
	}
	return total
}

/* Get trade market value in this period
 *
 * It is the last bar's trade-market-value
 */
func (this BarsMerger) GetTradedMarketValue() float64 {
	return this.Bars[this.Length-1].Traded_market_value
}

/* Get market-value in this period
 *
 * It is the last bar's Market-value
 */
func (this BarsMerger) GetMarketValue() float64 {
	return this.Bars[this.Length-1].Market_value
}

/* Get the total money in this period.
 * 【换手率】
 * 算法：  成交量 / 流通股本
 *     成交量： 本周期内成交量总和。
 *     流通股本： 期末流通市值/期末股价
 */
func (this BarsMerger) GetTurnover() float64 {
	var v = this.GetVolume()                              // 成交量
	var m = this.GetTradedMarketValue() / this.GetClose() // 流通股本
	to := v / m
	return to
}

/* Get adjust-price in this period
 *
 * It is the last bar's adjust-price
 */
func (this BarsMerger) GetAdjustPrice() float64 {
	return this.Bars[this.Length-1].Adjust_price
}

/* Get adjust-price-f in this period
 *
 * It is the last bar's adjust-price-f
 */
func (this BarsMerger) GetAdjustPriceF() float64 {
	return this.Bars[this.Length-1].Adjust_price_f
}

// Get count of trade days in this period.
func (this BarsMerger) GetTradeDays() int {
	var total int
	for _, b := range this.Bars {
		total += b.Trade_days
	}
	return total
}

// Get period mark.
func (this BarsMerger) GetOriginalPeriod() string {
	return this.Bars[0].Period
}
