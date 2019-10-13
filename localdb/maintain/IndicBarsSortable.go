package maintain

import (
	o "exgrow/localdb/object"
)

// IndicBarsSortable
type IndicBarsSortable []o.IndicBar

// 获取数组总长
func (b IndicBarsSortable) Len() int {
	return len(b)
}

//Less():按时间从早到晚排序
func (b IndicBarsSortable) Less(i, j int) bool {
	// 用 before() 函数判断时间的先后
	return b[i].Date.UTC().Before(b[j].Date.UTC())
}

//Swap()
func (b IndicBarsSortable) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
