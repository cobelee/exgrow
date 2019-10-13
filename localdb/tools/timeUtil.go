package tools

import (
	"time"
)

// 时间比较器 --- 声明了一个函数类型
// Determine whether the two dates are in the same period.
type TimeComparator func(t1 time.Time, t2 time.Time) bool

// Determine whether the two dates are the same week.
func IsSameWeek(t1 time.Time, t2 time.Time) bool {
	var y1, w1, y2, w2 int
	y1, w1 = t1.ISOWeek()
	y2, w2 = t2.ISOWeek()
	if y1 == y2 && w1 == w2 {
		return true
	} else {
		return false
	}
}

// Determine whether the two dates are the same month
func IsSameMonth(t1 time.Time, t2 time.Time) bool {
	var m1, m2 string
	m1 = t1.Format("0601")
	m2 = t2.Format("0601")
	return m1 == m2
}

// Determine whether the two dates are the same quarter.
func IsSameQuarter(t1 time.Time, t2 time.Time) bool {
	var isSame bool = false
	if t1.Year() == t2.Year() {
		if quarter(t1) == quarter(t2) {
			isSame = true
		}
	}
	return isSame
}

// Determine whether the two dates are the same year.
func IsSameYear(t1 time.Time, t2 time.Time) bool {
	var isSame bool = false
	if t1.Year() == t2.Year() {
		isSame = true
	}
	return isSame
}

// Return the quarter number in a year.
func quarter(t time.Time) int {
	month := int(t.Month())
	q := (month-1)/3 + 1
	return q
}
