package object

import (
	"time"
)

// ----------------------------------------------------------------

// SDBar-array implement the BarsScanable-interface
type SDBarArray []SDBar

// The following three method is the interface method with BarsScanable.

func (this SDBarArray) Len() int {
	return len(this)
}

func (this SDBarArray) GetBar(i int) interface{} {
	return this[i]
}

func (this SDBarArray) GetCurrentDate(i int) time.Time {
	return this[i].Date
}
