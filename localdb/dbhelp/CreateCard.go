package dbhelp

import (
	// "fmt"
	c "exgrow/localdb/config"
	o "exgrow/localdb/object"
	"strings"
)

// 创建股票数据卡片实例
func CreateSDCard(longCode string) o.SDCard {
	var stock o.SDCard
	stock.LongCode = longCode                           // 初始化长代码
	stock.ShortCode = strings.TrimLeft(longCode, "shz") // 初始化短代码
	dbName := c.DBConfig.DBName.StockMarketRawD1
	var sdbarArray []o.SDBar
	GetCollectionData(dbName, longCode, &sdbarArray)
	stock.Matrix = sdbarArray
	stock.BaseBarMatrix = GetBaseBarMatrix(stock.Matrix)
	stock.SDDBarMatrix = GetSDDBarMatrix(stock.Matrix)
	return stock
}

// -------------------------------------------------------------------

/* Create Indication Card instance.

   e.g. longCode: sh600000					Stock code.
		pt : "D", "W", "M", "Q", "Y" etc.		Period type.
*/
func CreateStockCard(longCode string, pt c.PeriodType) o.IndicCard {
	var card o.IndicCard
	card.LongCode = longCode                           // 初始化长代码
	card.ShortCode = strings.TrimLeft(longCode, "shz") // 初始化短代码
	card.Period = pt                                   // 初始化周期

	var dbName string
	dbName = c.GetStockDBNameByPeriodType(pt)

	var barArray []o.IndicBar
	GetCollectionData(dbName, longCode, &barArray)
	card.Matrix = barArray
	card.IndicBarMatrix = barArray
	return card
}

// -----------------------------------------------------------------

// 获取BaseBarMatrix
func GetBaseBarMatrix(sdBars []o.SDBar) []o.BaseBar {
	var baseBars = []o.BaseBar{}
	length := len(sdBars)
	for i := 0; i < length; i++ {
		var baseBar o.BaseBar = sdBars[i]
		baseBars = append(baseBars, baseBar)
	}
	return baseBars

}

// 获取BaseBarMatrix
func GetSDDBarMatrix(sdBars []o.SDBar) []o.SDDBar {
	var sddBars = []o.SDDBar{}
	length := len(sddBars)
	for i := 0; i < length; i++ {
		sddBar := sdBars[i].GetSDDBar()
		sddBars = append(sddBars, sddBar)
	}
	return sddBars

}
