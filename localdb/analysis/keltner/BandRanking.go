//

package keltner

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"fmt"
	"math"
	"sort"
	"time"
)

// --------------------------------------------------------------

type SortableRecord []KCRecord

// 获取数组总长
func (b SortableRecord) Len() int {
	return len(b)
}

//Less():按时间从早到晚排序
func (b SortableRecord) Less(i, j int) bool {
	// 用 before() 函数判断时间的先后
	return b[i].KCD < b[j].KCD

}

//Swap()
func (b SortableRecord) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// ----------------------------------------------------------------

type SortMethod func(i int, j int) bool

// ----------------------------------------------------------------

func BandRanking() {
	dbName := c.DBConfig.DBName.StockD1
	names, _ := h.GetCollectionNames(dbName)
	var filterBarArray o.IndicBarArray
	nameCount := len(names)
	for i, name := range names {
		stockDCard := h.CreateStockCard(name, c.Period_D)
		barArray := stockDCard.IndicBarMatrix

		if barArray.Len() == 0 {
			continue
		}

		lastBar := barArray.GetLastBar()
		due := time.Now().Sub(lastBar.Date)
		if due.Hours()/24 < 15 {
			filterBarArray = append(filterBarArray, lastBar)
		}
		fmt.Printf("Calculating now, Please wait a while. %v%%\r", int(i*100/nameCount))
	}
	fmt.Println("Processing completed. 100%")

	var ranking []KCRecord
	for _, stock := range filterBarArray {
		record := KCRecord{
			Code:   stock.Code,
			CloseT: stock.CloseT,
			KCD:    math.Ceil((stock.Close-stock.EMA39)*100/stock.TREMA+0.5) / 100,
		}
		ranking = append(ranking, record)
	}

	fmt.Println()

	// 按从小到大顺序排序，输出头20个。
	sort.Slice(ranking, func(i int, j int) bool {
		return ranking[i].KCD < ranking[j].KCD
	})
	fmt.Println("     Stock     BandCount    PriceT")
	head20 := ranking[0:20]
	for _, stock := range head20 {
		fmt.Printf("%10s    %10.2f%10.2f\n", stock.Code, stock.KCD, stock.CloseT)
	}
	fmt.Println()

	// 逆序后，输出头20个。
	fmt.Println("     Stock     BandCount    PriceT")
	var i, length, forCount int
	forCount = 20
	length = len(ranking)
	if length < 20 {
		forCount = length
	}
	for i = 1; i <= forCount; i++ {
		fmt.Printf("%10s    %10.2f%10.2f\n", ranking[length-i].Code,
			ranking[length-i].KCD, ranking[length-i].CloseT)
	}

}
