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

// IndicBar-array implement the BarsScanable-interface
type IndicBarArray []o.IndicBar

// The following three method is the interface method with BarsScanable.

func (this IndicBarArray) Len() int {
	return len(this)
}

func (this IndicBarArray) GetBar(i int) interface{} {
	return this[i]
}

func (this IndicBarArray) GetCurrentDate(i int) time.Time {
	return this[i].Date
}

// -----------------------------------------------------------------

// 单个股票历史指标数据集合结构体
type IndicDCard struct {
	ShortCode      string        // [股票短代码]如 600000
	LongCode       string        // [股票长代码]如 sh600000
	Matrix         []o.IndicBar  // 指定代码的股票指标数据矩阵
	IndicBarMatrix IndicBarArray // IndicBar 指标期数据矩阵
}

// 此结构体数据的数据库名称。
func (this *IndicDCard) GetDBName() string {
	return c.DBConfig.DBName.IndicationD1
}

// 此对象存储到数据库中的集合名称。
func (this *IndicDCard) GetCollectionName() string {
	return this.LongCode
}

// SDCard 类型值获取最早一天的日期
func (this *IndicDCard) GetFirstDate() (time.Time, error) {
	var date time.Time
	if this.Matrix != nil && len(this.Matrix) > 0 {
		date = this.Matrix[0].Date.UTC()
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}

func (this *IndicDCard) GetLastDate() (time.Time, error) {
	var date time.Time
	var count = len(this.Matrix)
	if this.Matrix != nil && count > 0 {
		date = this.Matrix[count-1].Date.UTC()
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}

func (this *IndicDCard) GetLast10Bar() []o.IndicBar {
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

// 创建股票数据卡片实例
func CreateIndicDCard(longCode string) IndicDCard {
	var card IndicDCard
	card.LongCode = longCode                           // 初始化长代码
	card.ShortCode = strings.TrimLeft(longCode, "shz") // 初始化短代码
	dbName := c.DBConfig.DBName.IndicationD1
	var barArray []o.IndicBar
	dbhelp.GetCollectionData(dbName, longCode, &barArray)
	card.Matrix = barArray
	card.IndicBarMatrix = barArray
	return card
}
