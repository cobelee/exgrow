package localdb

import (
	"errors"
)

type BarScanner struct {
	matrix       []SDDBar    //
	BarBuffer    []SDDBar    // Scan 操作后的数据
	maxTokenSize int         // Maximum size of a token; modified by tests.
	token        interface{} // Last token returned by split.
	start        int         // First non-processed byte in buf.
	index        int         // Current index in buf.
	end          int         // End of data in buf.
	err          error       // Sticky error.
	scanCalled   bool        // Scan has been called; buffer is in use.
	done         bool        // Scan has finished.
}

// 默认读取一日数据
func (this *BarScanner) Scan() bool {
	this.scanCalled = true

	if this.index == this.end {
		this.done = true
		this.err = errors.New("EOF")
		return false
	}

	this.index++
	return true
}

// 读取一周数据
func (this *BarScanner) ScanAWeek() bool {
	ok := false
	var year, week, yeari, weeki int
	var i = 0

	if this.done {
		return ok
	}

	for this.index <= this.end {
		if i == 0 {
			this.BarBuffer = nil
			year, week = this.matrix[this.index].Date.ISOWeek()
			this.BarBuffer = append(this.BarBuffer, this.matrix[this.index])
			ok = true
		}
		if i > 0 {
			yeari, weeki = this.matrix[this.index].Date.ISOWeek()
			if year == yeari && week == weeki {
				this.BarBuffer = append(this.BarBuffer, this.matrix[this.index])
			} else {
				break
			}
		}
		this.index++
		i++
	}

	if this.index > this.end {
		this.done = true
		this.err = errors.New("EOF")
		ok = false
	}

	this.scanCalled = ok

	return ok

}

// 获取当前索引处的K线数据
func (this *BarScanner) Bar() SDDBar {
	if this.scanCalled == false {
		return SDDBar{}
	} else {
		return this.matrix[this.index]
	}
}

// 重置扫描指针索引
func (this *BarScanner) Reset() {
	this.index = this.start
	this.done = false
}

func NewBarScanner(m []SDDBar) BarScanner {
	var s BarScanner
	s.matrix = m
	s.maxTokenSize = len(m)
	s.token = nil
	s.start = 0
	s.index = 0
	s.end = s.maxTokenSize - 1
	s.err = nil
	s.scanCalled = false
	s.done = false

	return s
}
