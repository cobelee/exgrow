package object

import (
	c "exgrow/localdb/config"
	"time"
)

// ---------------------------------------------------------------------------

/* Indication record include stock market data and indications data.

  The values of open, high, low, close are adjusted by Adjust-Factor.
* 包括周K线、月K线、季K线。
*/
type IndicBar struct {
	Code       string    `bson:"code"`       // The stock code. such as 'sh600000', 上证股票以sh开头，深证股票以sz开头
	Date       time.Time `bson:"date"`       // 【交易日期】 一般为指定周期市场的最后一个交易日日期。
	Open       float64   `bson:"open"`       // The open price adjusted.
	High       float64   `bson:"high"`       // The high price adjusted.
	Low        float64   `bson:"low"`        // The low price adjusted.
	Close      float64   `bson:"close"`      // The close price adjusted.
	EMA39      float64   `bson:"ema39"`      // Indication: EMA39 39日指数移动平均值
	TR         float64   `bson:"tr"`         // Indication: TR 波动率
	Trade_days int       `bson:"trade_days"` // 【days of trade】 本周股票实际交易的天数
	Period     string    `bson:"period"`     // 【K线周期】可能为 周K线，月K线，季K线，其值分别 W M Q
}

// 此对象存储到的数据库名称。
func (this *IndicBar) GetDBName() string {

	var dbName string
	// 从项目配置文件中获取

	switch this.Period {
	case "D":
		dbName = c.DBConfig.DBName.IndicationD1
	case "W":
		dbName = c.DBConfig.DBName.IndicationW1
	case "M":
		dbName = c.DBConfig.DBName.IndicationM1
	case "Q":
		dbName = c.DBConfig.DBName.IndicationQ1
	case "Y":
		dbName = c.DBConfig.DBName.IndicationY1
	default:
		dbName = c.DBConfig.DBName.IndicationD1
	}
	return dbName
}

// 此对象存储到数据库中的集合名称。
func (this *IndicBar) GetCollectionName() string {
	return this.Code
}

// 获取此对象在数据库集合中的键值。
func (this *IndicBar) MajorKey() time.Time {
	return this.Date.UTC()
}

func NewIndicBar(bar SDDBar) IndicBar {
	var adjust float64
	adjust = bar.Adjust_price / bar.Close

	return IndicBar{
		Code:       bar.Code,
		Date:       bar.Date,
		Open:       bar.Open * adjust,
		High:       bar.High * adjust,
		Low:        bar.Low * adjust,
		Close:      bar.Adjust_price,
		Trade_days: bar.Trade_days,
		Period:     bar.Period,
	}
}
