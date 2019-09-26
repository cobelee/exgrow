/*
localdb

*/
package localdb

import (
	"bufio"
	"exgrow/alistock"
	"exgrow/localdb/dbhelp"
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SyncStockCode() {

	var err error
	var session *mgo.Session
	session = dbhelp.GetSession() // 初始化mgo.Session 并获取Session

	c := session.DB("aliStockDB").C("stock_code")

	var mapStock, mapStockFiltered map[string]alistock.StockCode
	// Begin to fetch stockCode from aliStock
	WriteLog("Begin to fetch stockCode from aliStock.")
	fmt.Println("Begin to fetch stockCode from aliStock.")
	mapStock = alistock.GetStockMap()
	mapStockFiltered = make(map[string]alistock.StockCode)

	// 将mapStock中非上海非深圳证券交易所的品种过滤掉，并输出到 mapStockFiltered。
	var count int = 0
	for index, stock := range mapStock {
		if stock.Market_ch == "上海证券交易所" || stock.Market_ch == "深圳证券交易所" {
			mapStockFiltered[index] = stock
			count++
		}
	}

	WriteLog("Begin to sync to localdb.")
	fmt.Println("Fetch is completed.  Begin to sync to localdb.")
	nSucc := 0
	nFail := 0
	for _, stock := range mapStockFiltered {
		seletor := bson.M{"code": stock.Code}
		_, err = c.Upsert(seletor, stock) // 采用upsert防止重复插入。
		if err != nil {
			nFail++
		} else {
			nSucc++
		}
	}

	writeSyncResultToDBLog(nSucc, nFail)
	fmt.Printf("Sync is completed.\n")

}

// Log the sync result to log.txt
func writeSyncResultToDBLog(nSucc, nFail int) {

	f, _ := os.OpenFile("localdb/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()

	w := bufio.NewWriter(f)
	cTime := time.Now().Format("2006-01-02 15:04:05")
	lineStr := fmt.Sprintf("%s > StockCodes has be inserted to localdb. Success: %v, Fail: %v.\n", cTime, nSucc, nFail)
	w.WriteString(lineStr)
	w.Flush()
	fmt.Printf("The total count of stocks updated is %v\n", nSucc+nFail)
}
