package object

import (
	"time"
)

// 对象状态监控器的结构体
type ObjectMonitor struct {
	Date     time.Time // 被监控记录的日期
	Modified bool      // 是否被修改标记
}
