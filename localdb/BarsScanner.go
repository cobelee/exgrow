package localdb

import (
	"errors"
	t "exgrow/localdb/tools"
	"time"
)

type BarsScanable interface {
	Len() int
	GetBar(int) interface{}
	GetCurrentDate(int) time.Time
}

// 用于缓冲区切片夹板左右界线
type SliceLimit struct {
	LeftLimit  int
	RightLimit int
}

type BarsScanner struct {
	ScanableArray  BarsScanable // Raw data array scanable.
	BarIndexBuffer []int        // Scan buffer
	BufferBoards   SliceLimit   // BufferBoards
	maxTokenSize   int          // Maximum size of a token; modified by tests.
	Token          interface{}  // Last token returned by split.
	start          int          // First non-processed byte in buf.
	index          int          // Current index in buf.
	end            int          // End of data in buf.
	err            error        // Sticky error.
	scanCalled     bool         // Scan has been called; The value will be set to true when first be called.
	done           bool         // Scan has finished.
}

/* Read one day data by default.

   1. 当第一次调用Scan()结束时，索引被置于0位置，读取到的是第一个bar数据。
   2. 当Scan到最后一条记录时，Scan()返回true，done返回true,不引发EOF错误。
      直至再次Scan()，返回false, 并引发EOF错误。
*/
func (this *BarsScanner) Scan() bool {

	var rev = false

	if this.done == true {
		rev = false
		this.Token = nil
		this.err = errors.New("EOF")
	}

	if this.scanCalled == false {
		this.index = 0
		this.scanCalled = true
		this.Token = this.ScanableArray.GetBar(this.index)
		rev = true
	} else {
		if this.done == false {
			this.index++
			this.Token = this.ScanableArray.GetBar(this.index)
			rev = true
		}
	}

	if this.index == this.end {
		this.done = true
	}

	return rev
}

func (this *BarsScanner) scanPeriod(isSamePeriod t.TimeComparator) bool {
	var rev = false
	var t1, ti time.Time
	var i = 0

	if this.ScanableArray.Len() == 0 {
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
			this.BarIndexBuffer = nil
			t1 = this.ScanableArray.GetCurrentDate(this.index)
			this.BarIndexBuffer = append(this.BarIndexBuffer, this.index)
			rev = true
			this.BufferBoards.LeftLimit = this.index
			this.BufferBoards.RightLimit = this.index + 1
		}
		if i > 0 {
			ti = this.ScanableArray.GetCurrentDate(this.index)
			if isSamePeriod(t1, ti) {
				this.BarIndexBuffer = append(this.BarIndexBuffer, this.index)
				this.BufferBoards.RightLimit++
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
func (this *BarsScanner) ScanAWeek() bool {
	return this.scanPeriod(t.IsSameWeek)
}

// Scan a monthly data
func (this *BarsScanner) ScanAMonth() bool {
	return this.scanPeriod(t.IsSameMonth)

}

// Scan a quarterly data
func (this *BarsScanner) ScanAQuarter() bool {
	return this.scanPeriod(t.IsSameQuarter)
}

// 获取当前索引处的K线数据
func (this *BarsScanner) Bar() (interface{}, bool) {
	if this.scanCalled == false {
		return nil, false
	} else {
		return this.Token, true
	}
}

// 获取当前索引左侧第一根K线数据
func (this *BarsScanner) PreviBar() (interface{}, bool) {
	if this.index <= 0 {
		return nil, false
	}
	if this.index >= this.maxTokenSize {
		return nil, false
	}
	return this.ScanableArray.GetBar(this.index - 1), true

}

// 获取包括当前索引及左侧指定数量的多根K线数据
// func (this *BarsScanner) RefBars(count int) ([]SDDBar, bool) {

// 	var start, end int
// 	if this.index-count+1 < 0 {
// 		start = 0
// 	} else {
// 		start = this.index - count + 1
// 	}

// 	end = this.index + 1

// 	b := this.matrix[start:end]

// 	if len(b) < count {
// 		return b, false
// 	} else {
// 		return b, true
// 	}

// }

// Get current index in buf.
func (this *BarsScanner) CurrentIndex() int {
	return this.index
}

// 重置扫描指针索引
func (this *BarsScanner) Reset() {
	this.index = this.start
	this.scanCalled = false
	this.done = false
}

func NewBarsScanner(m BarsScanable) BarsScanner {
	var s BarsScanner
	s.ScanableArray = m
	s.maxTokenSize = m.Len()
	s.Token = nil
	s.start = 0
	s.index = 0
	s.end = s.maxTokenSize - 1
	s.err = nil
	s.scanCalled = false
	s.done = false
	s.BufferBoards.LeftLimit = 0
	s.BufferBoards.RightLimit = 0

	return s
}
