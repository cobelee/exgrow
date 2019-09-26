// BarsSortable
package localdb

import (
	o "exgrow/localdb/object"
)

type SortableBars []o.SDDBar

// 获取数组总长
func (b SortableBars) Len() int {
	return len(b)
}

//Less():按时间从早到晚排序
func (b SortableBars) Less(i, j int) bool {
	// 用 before() 函数判断时间的先后
	return b[i].Date.UTC().Before(b[j].Date.UTC())

}

//Swap()
func (b SortableBars) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
