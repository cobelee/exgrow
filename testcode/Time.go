package testcode

import (
	"fmt"
	"time"
)

// 演示如何将时间字符串Parse为ISOTime时间类型
func TestTime() {
	strTime1 := "2019-01-02 11:39:22"
	strTime2 := "2019-01-02 11:45:28"
	t1, _ := time.Parse("2006-01-02 15:04:05", strTime1)
	t2, _ := time.Parse("2006-01-02 15:04:05", strTime2)
	t := t2.Sub(t1)
	fmt.Println(t)
}
