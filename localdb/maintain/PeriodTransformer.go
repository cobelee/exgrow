package maintain

import (
	c "exgrow/localdb/config"
	o "exgrow/localdb/object"
)

/* 将短周期切片转化为长周期切片。

   建议：可将“日”周期转化为“周”、“月”的长周期。
		将“月”周期转化为“季”、“年”的长周期。

*/
func TransformPeriod(source o.IndicBarArray, destPt c.PeriodType) (dest o.IndicBarArray) {
	// Begin to merge
	scanner := NewBarsScanner(source)
	for scanner.ScanByPeriodType(destPt) {
		var barBuffer []o.IndicBar
		barBuffer = source[scanner.BufferBoards.LeftLimit:scanner.BufferBoards.RightLimit]
		if merger, e := NewIndicBarsMerger(barBuffer); e == nil {
			longbar := merger.CreateLongPeriodBar(destPt)
			dest = append(dest, longbar)
		}
	}

	return dest
}
