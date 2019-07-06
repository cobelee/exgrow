package testcode

import (
	"fmt"
	"time"
)

// 演示如何将时间字符串Parse为ISOTime时间类型
func testTime() {
	strTime := "2019-01-02"
	t, e := time.Parse("2006-01-02", strTime)
	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println(t)
}
