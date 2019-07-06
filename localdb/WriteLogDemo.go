/*
	Just test program during Dev.
*/

package localdb

import (
	"bufio"
	"exgrow/alistock"
	"fmt"
	"os"
)

// The demo show how to write text to a local file  with bufio.NewWrite()
//  and fmt.Fprintf()
func WriteLogDemo() {
	var mapStock, mapStockFiltered map[string]alistock.StockCode

	mapStockFiltered = make(map[string]alistock.StockCode)
	mapStock = alistock.GetStockMap()

	f, _ := os.Create("localdb/stockCode.txt")
	defer f.Close()

	var count int = 0
	w := bufio.NewWriter(f)

	for index, stock := range mapStock {
		if stock.Market_ch == "上海证券交易所" || stock.Market_ch == "深圳证券交易所" {
			mapStockFiltered[index] = stock
			count++
			lineStr := fmt.Sprintf("%9v   %9v      %19v  %19v\n", index, stock.Code, stock.Name_ch_abbr, stock.Market_ch)
			fmt.Fprintf(w, lineStr)

			//			fmt.Printf("%9v   %9v      %19v  %19v\n", index, stock.Code, stock.Name_ch_abbr, stock.Market_ch)
		}

	}
	w.WriteString("-----------------------------------")
	lineStr2 := fmt.Sprintf("This is the end, the total count is %v\n", count)
	w.WriteString(lineStr2)
	w.Flush()
	fmt.Printf("The total count of stocks is %v\n", count)
}
