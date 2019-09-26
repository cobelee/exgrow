package test

import (
	q "exgrow/localdb/query"
	"fmt"
	"time"
)

func PrintSddBar() {
	var t time.Time
	strTime := "2019-05-16"
	t, _ = time.Parse("2006-01-02", strTime)

	// today := time.Date(2019, time.May, 15, 0, 0, 0, 0, time.UTC)

	bar := q.GetSDDBar("sh600000", "D", t)
	fmt.Print(bar)
}
