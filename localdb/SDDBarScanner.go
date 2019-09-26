/* Scanner for SDDBar

   Provide the functions to scan sddbars by different period.
   Return the result scaned.

   *** 本扫描器已被否决，请用 BarsScanner 替代。 ***

*/

package localdb

import (
	"errors"
	o "exgrow/localdb/object"
	t "exgrow/localdb/tools"
	"time"
)

type SDDBarScanner struct {
	matrix       []o.SDDBar  // Raw data array of stock-bar
	BarBuffer    []o.SDDBar  // Scan buffer
	maxTokenSize int         // Maximum size of a token; modified by tests.
	Token        interface{} // Last token returned by split.
	start        int         // First non-processed byte in buf.
	index        int         // Current index in buf.
	end          int         // End of data in buf.
	err          error       // Sticky error.
	scanCalled   bool        // Scan has been called; The value will be set to true when first be called.
	done         bool        // Scan has finished.
}

/* Read one day data by default.

   1. 当第一次调用Scan()结束时，索引被置于0位置，读取到的是第一个bar数据。
   2. 当Scan到最后一条记录时，Scan()返回true，done返回true,不引发EOF错误。
      直至再次Scan()，返回false, 并引发EOF错误。
*/
func (this *SDDBarScanner) Scan() bool {

	var rev = false

	if this.done == true {
		rev = false
		this.Token = nil
		this.err = errors.New("EOF")
	}

	if this.scanCalled == false {
		this.index = 0
		this.scanCalled = true
		this.Token = this.matrix[this.index]
		rev = true
	} else {
		if this.done == false {
			this.index++
			this.Token = this.matrix[this.index]
			rev = true
		}
	}

	if this.index == this.end {
		this.done = true
	}

	return rev
}

func (this *SDDBarScanner) scanPeriod(isSamePeriod t.TimeComparator) bool {
	var rev = false
	var t1, ti time.Time
	var i = 0

	if len(this.matrix) == 0 {
		rev = false
		this.scanCalled = true
		this.done = true
		this.err = errors.New("EOF")
		return rev
	}

	if this.done {
		rev = false
		this.err = errors.New("EOF")
		return rev
	}

	if this.scanCalled == false {
		this.index = 0
		this.scanCalled = true
	}

	for this.index <= this.end {
		if i == 0 {
			this.BarBuffer = nil
			t1 = this.matrix[this.index].Date
			this.BarBuffer = append(this.BarBuffer, this.matrix[this.index])
			rev = true
		}
		if i > 0 {
			ti = this.matrix[this.index].Date
			if isSamePeriod(t1, ti) {
				this.BarBuffer = append(this.BarBuffer, this.matrix[this.index])
			} else {
				break
			}
		}
		if this.index < this.end {
			this.index++
		} else {
			this.done = true
			break
		}

		i++
	}

	return rev
}

// Scan a weekly data.
func (this *SDDBarScanner) ScanAWeek() bool {
	return this.scanPeriod(t.IsSameWeek)
}

// Scan a monthly data
func (this *SDDBarScanner) ScanAMonth() bool {
	return this.scanPeriod(t.IsSameMonth)

}

// Scan a quarterly data
func (this *SDDBarScanner) ScanAQuarter() bool {
	return this.scanPeriod(t.IsSameQuarter)
}

// 获取当前索引处的K线数据
func (this *SDDBarScanner) Bar() (o.SDDBar, bool) {
	if this.scanCalled == false {
		return o.SDDBar{}, false
	} else {
		return this.matrix[this.index], true
	}
}

// 获取当前索引左侧第一根K线数据
func (this *SDDBarScanner) PreviBar() (o.SDDBar, bool) {
	if this.index <= 0 {
		return o.SDDBar{}, false
	}
	if this.index >= this.maxTokenSize {
		return o.SDDBar{}, false
	}
	return this.matrix[this.index-1], true

}

// 获取包括当前索引及左侧指定数量的多根K线数据
func (this *SDDBarScanner) RefBars(count int) ([]o.SDDBar, bool) {

	var start, end int
	if this.index-count+1 < 0 {
		start = 0
	} else {
		start = this.index - count + 1
	}

	end = this.index + 1

	b := this.matrix[start:end]

	if len(b) < count {
		return b, false
	} else {
		return b, true
	}

}

// 重置扫描指针索引
func (this *SDDBarScanner) Reset() {
	this.index = this.start
	this.scanCalled = false
	this.done = false
}

func NewSDDBarScanner(m []o.SDDBar) SDDBarScanner {
	var s SDDBarScanner
	s.matrix = m
	s.maxTokenSize = len(m)
	s.Token = nil
	s.start = 0
	s.index = 0
	s.end = s.maxTokenSize - 1
	s.err = nil
	s.scanCalled = false
	s.done = false

	return s
}
