// BarsBuffer 处理器
package localdb

import (
	// "fmt"
	"exgrow/errors"
	"sort"
	"time"
)

type BarsProcesser struct {
	Length int
	Bars   []SDDBar // 处理对象
}

func (this BarsProcesser) Date() time.Time {
	return this.Bars[this.Length-1].Date
}

func (this BarsProcesser) High() float64 {
	var v float64
	for _, b := range this.Bars {
		if b.High > v {
			v = b.High
		}
	}
	return v
}

// Get low price in the current period
func (this BarsProcesser) Low() float64 {
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
func (this BarsProcesser) Open() float64 {
	if this.Length == 0 {
		return 0
	} else {
		return this.Bars[0].Open
	}
}

// Get close price in the period.
func (this BarsProcesser) Close() float64 {
	if this.Length == 0 {
		return 0
	} else {
		return this.Bars[this.Length-1].Close
	}
}

// 创建多巴处理器
func NewBarsProcesser(m []SDDBar) (BarsProcesser, error) {
	var p BarsProcesser

	p.Bars = m
	sort.Sort(SortableBars(m)) // 对BarBuffer，按Bar.Date日期从小到大的顺序进行排列。

	p.Length = len(m)

	if m == nil || len(m) == 0 {
		err := errors.NewError("100009", "No bars data.")
		return p, &err
	} else {
		return p, nil
	}
}
