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

// IdcBar-array implement the BarsScanable-interface
type IdcBarArray []o.IndicBar

// The following three method is the implementation of the BarsScanable-interface.
// Len()  GetBar()  GetCurrentDate()

func (this IdcBarArray) Len() int {
	return len(this)
}

func (this IdcBarArray) GetBar(i int) interface{} {
	return this[i]
}

func (this IdcBarArray) GetCurrentDate(i int) time.Time {
	return this[i].Date
}

// -----------------------------------------------------------------

// 单个股票历史指标数据集合结构体
type IdcCard struct {
	ShortCode    string       // [股票短代码]如 600000
	LongCode     string       // [股票长代码]如 sh600000
	Matrix       []o.IndicBar // 指定代码的股票全生命期数据矩阵
	IdcBarMatrix IdcBarArray  // SDDBar股票全生命期数据矩阵
}

// 此结构体数据的数据库名称。
func (this *IdcCard) GetDBName() string {
	// config := GetConfig() // 获取项目配置文件
	// return config.DBName.StockMarketRawD1
	return c.DBConfig.DBName.StockMarketRawD1
}

// 此对象存储到数据库中的集合名称。
func (this *IdcCard) GetCollectionName() string {
	return this.LongCode
}

// IdcCard 类型值获取最早一天的日期
func (this *IdcCard) GetFirstDate() (time.Time, error) {
	var date time.Time
	if this.Matrix != nil && len(this.Matrix) > 0 {
		date = this.Matrix[0].Date.UTC()
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}

func (this *IdcCard) GetLastDate() (time.Time, error) {
	var date time.Time
	var count = len(this.Matrix)
	if this.Matrix != nil && count > 0 {
		date = this.Matrix[count-1].Date.UTC()
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}

func (this *IdcCard) GetLast10Bar() []o.IndicBar {
	var bar []o.IndicBar
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

// Create card instance including indication-data
func CreateIdcCard(longCode string) IdcCard {
	var card IdcCard
	card.LongCode = longCode                           // 初始化长代码
	card.ShortCode = strings.TrimLeft(longCode, "shz") // 初始化短代码
	dbName := c.DBConfig.DBName.IndicationD1

	var idcArray []o.IndicBar
	dbhelp.GetCollectionData(dbName, longCode, &idcArray)
	card.Matrix = idcArray
	card.IdcBarMatrix = card.Matrix //  GetIdcBarMatrix(stock.Matrix)
	return card
}

// 获取BaseBarMatrix
// func GetIdcBarMatrix(idcBars []IndicSet) IdcBarArray {
// 	iba := idcBars
// 	return iba

// }
