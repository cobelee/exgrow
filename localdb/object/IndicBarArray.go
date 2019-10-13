package object

import (
	"exgrow/errors"
	"time"
)

// ----------------------------------------------------------------

// IndicBar-array implement the BarsScanable-interface
type IndicBarArray []IndicBar

// The following three method is the interface method with IndicBarsScanable.

func (this IndicBarArray) Len() int {
	return len(this)
}

func (this IndicBarArray) GetBar(i int) interface{} {
	return this[i]
}

func (this IndicBarArray) GetCurrentDate(i int) time.Time {
	return this[i].Date
}

/* 获取指定日期(包含)右侧的切片。

 */
func (this IndicBarArray) RightSlice(d time.Time) IndicBarArray {
	var index int
	var s2 IndicBarArray
	length := len(this)
	if length == 0 {
		return this[:]
	}

	// 此处循环不可使用 for range 语句。
	for i := 0; i < length; i++ {
		if this[i].Date.Equal(d) || this[i].Date.After(d) {
			index = i
			break
		}
	}

	if index != 0 {
		s2 = this[index:]
	} else {
		s2 = this[:]
	}
	return s2
}

// Remove the last n bars, n must be <= len(array)
func (this IndicBarArray) RemoveLast(n int) IndicBarArray {
	length := len(this)

	if length >= n {
		return this[0 : length-n]
	} else {
		return this[0:0]
	}
}

// Get the last bar.
func (this IndicBarArray) GetLastBar() IndicBar {
	length := len(this)

	if length > 0 {
		return this[length-1]
	} else {
		return IndicBar{}
	}
}

// Get the last bar's date.
func (this IndicBarArray) GetLastDate() (time.Time, error) {
	length := len(this)
	if length > 0 {
		return this[length-1].Date, nil
	} else {
		strTime := "1900-01-01"
		t, _ := time.Parse("2006-01-02", strTime)
		e := errors.NewError("100009", "No bars data.")
		return t, &e
	}
}
