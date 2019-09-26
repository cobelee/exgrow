package localdb

import (
	// "fmt"
	"exgrow/errors"
	c "exgrow/localdb/config"
	"exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"strings"
	"time"
)

// ----------------------------------------------------------------

// SDDBar-array implement the BarsScanable-interface
type SDDBarArray []o.SDDBar

// The following three method is the interface method with BarsScanable.

func (this SDDBarArray) Len() int {
	return len(this)
}

func (this SDDBarArray) GetBar(i int) interface{} {
	return this[i]
}

func (this SDDBarArray) GetCurrentDate(i int) time.Time {
	return this[i].Date
}

// -----------------------------------------------------------------

// 单个股票历史数据集合结构体
type SDCard struct {
	ShortCode     string      // [股票短代码]如 600000
	LongCode      string      // [股票长代码]如 sh600000
	Matrix        []o.SDBar   // 指定代码的股票全生命期数据矩阵
	BaseBarMatrix []BaseBar   // BaseBar股票全生命期数据矩阵
	SDDBarMatrix  SDDBarArray // SDDBar股票全生命期数据矩阵
}

// 此结构体数据的数据库名称。
func (this *SDCard) GetDBName() string {
	// config := GetConfig() // 获取项目配置文件
	// return config.DBName.StockMarketRawD1
	return c.DBConfig.DBName.StockMarketRawD1
}

// 此对象存储到数据库中的集合名称。
func (this *SDCard) GetCollectionName() string {
	return this.LongCode
}

// SDCard 类型值获取最早一天的日期
func (this *SDCard) GetFirstDate() (time.Time, error) {
	var date time.Time
	if this.Matrix != nil && len(this.Matrix) > 0 {
		date = this.Matrix[0].Date.UTC()
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}

func (this *SDCard) GetLastDate() (time.Time, error) {
	var date time.Time
	var count = len(this.Matrix)
	if this.Matrix != nil && count > 0 {
		date = this.Matrix[count-1].Date.UTC()
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}

func (this *SDCard) GetLast10Bar() []o.SDBar {
	var bar []o.SDBar
	var count = len(this.Matrix)
	var i, j int
	if count > 0 {
		i = count - 1
		for i >= 0 && j < 10 {
			bar = append(bar, this.Matrix[i])
			i--
			j++
		}
	}
	return bar
}

// -------------------------------------------------------------------

// 创建股票数据卡片实例
func CreateSDCard(longCode string) SDCard {
	var stock SDCard
	stock.LongCode = longCode                           // 初始化长代码
	stock.ShortCode = strings.TrimLeft(longCode, "shz") // 初始化短代码
	dbName := c.DBConfig.DBName.StockMarketRawD1
	var sdbarArray []o.SDBar
	dbhelp.GetCollectionData(dbName, longCode, &sdbarArray)
	stock.Matrix = sdbarArray
	stock.BaseBarMatrix = GetBaseBarMatrix(stock.Matrix)
	stock.SDDBarMatrix = GetSDDBarMatrix(stock.Matrix)
	return stock
}

// 获取BaseBarMatrix
func GetBaseBarMatrix(sdBars []o.SDBar) []BaseBar {
	var baseBars = []BaseBar{}
	for _, b := range sdBars {
		var baseBar BaseBar = b
		baseBars = append(baseBars, baseBar)
	}
	return baseBars

}

// 获取BaseBarMatrix
func GetSDDBarMatrix(sdBars []o.SDBar) []o.SDDBar {
	var sddBars = []o.SDDBar{}
	for _, b := range sdBars {
		sddBars = append(sddBars, b.GetSDDBar())
	}
	return sddBars

}
