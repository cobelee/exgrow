package object

import (
	"time"
)

// ----------------------------------------------------------------

// SDDBar-array implement the BarsScanable-interface
type SDDBarArray []SDDBar

// The following three method is the interface method with BarsScanable.

func (this SDDBarArray) Len() int {
	return len(this)
}

func (this SDDBarArray) GetBar(i int) interface{} {
	return this[i]
}

func (this SDDBarArray) GetCurrentDate(i int) time.Time {
	return this[i].Date
}
