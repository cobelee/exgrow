package object

import (
	"exgrow/errors"
	c "exgrow/localdb/config"
	"time"
)

// 单个股票历史指标数据集合结构体
type IndicCard struct {
	ShortCode      string        // [股票短代码]如 600000
	LongCode       string        // [股票长代码]如 sh600000
	Period         c.PeriodType  // 数据卡片周期类型
	Matrix         []IndicBar    // 指定代码的股票指标数据矩阵
	IndicBarMatrix IndicBarArray // IndicBar 指标期数据矩阵
}

// 此结构体数据的数据库名称。
func (this *IndicCard) GetDBName() string {
	return c.DBConfig.DBName.StockD1
}

// 此对象存储到数据库中的集合名称。
func (this *IndicCard) GetCollectionName() string {
	return this.LongCode
}

// SDCard 类型值获取最早一天的日期
func (this *IndicCard) GetFirstDate() (time.Time, error) {
	var date time.Time
	if this.Matrix != nil && len(this.Matrix) > 0 {
		date = this.Matrix[0].Date.UTC()
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}

func (this *IndicCard) GetLastDate() (time.Time, error) {
	var date time.Time
	var count = len(this.Matrix)
	if this.Matrix != nil && count > 0 {
		date = this.Matrix[count-1].Date.UTC()
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}

func (this *IndicCard) GetLast10Bar() []IndicBar {
	var bar []IndicBar
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
