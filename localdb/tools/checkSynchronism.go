package tools

import (
	"time"
)

type Syncable interface {
	Len() int
	GetCurrentDate(int) time.Time
}

/* 测试两个数组是否具有同步一致性。

*  两个数组同步一致，需测试以下两项：
*  1、两个数组记录数一致。
*  2、两个数组中对应记录的日期一致。
 */
func CheckSynchronism(ba1, ba2 Syncable) bool {

	var isSameLen = false
	var isDateFit = true
	var size = 0

	size = ba1.Len()
	if ba1.Len() == ba2.Len() {
		isSameLen = true
	}

	for i := 0; i < size; i++ {
		if ba1.GetCurrentDate(i) != ba2.GetCurrentDate(i) {
			isDateFit = false
			break
		}
	}

	if isSameLen && isDateFit {
		return true
	}
	return false

}
