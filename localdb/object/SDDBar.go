package object

import (
	"exgrow/localdb/config"
	"time"
)

/* 股票多日行情数据，对应 CSV 文件中的单行数据。
 * 包括周K线、月K线、季K线。
 */
type SDDBar struct {
	Code                string    `bson:"code"`                // 股票的代码，上证股票以sh开头，深证股票以sz开头
	Date                time.Time `bson:"date"`                // 【交易日期】 本周市场的最后一个交易日
	Open                float64   `bson:"open"`                // 开盘价
	High                float64   `bson:"high"`                // 最高价
	Low                 float64   `bson:"low"`                 // 最低价
	Close               float64   `bson:"close"`               // 收盘价
	Change              float64   `bson:"change"`              // 涨跌幅，复权之后的真实涨跌幅，保证准确。
	Volume              float64   `bson:"volume"`              // 成交量
	Money               float64   `bson:"money"`               // 成交额
	Traded_market_value float64   `bson:"traded_market_value"` // 【流通市值】 流通市值，流通市值 / 股价 = 流通股股本
	Market_value        float64   `bson:"market_value"`        // 【总市值】 总市值，总市值 / 股价 = 总股本
	Turnover            float64   `bson:"turnover"`            // 【换手率】 成交量 / 流通股本  = 成交量*股价/流通市值
	Adjust_price        float64   `bson:"adjust_price"`        // 后复权价，复权开始时间为股票上市日，精确到小数点后10位。
	Adjust_price_f      float64   `bson:"adjust_price_f"`      // 前复权价，复权开始时间为股票最近一个交易日，精确到小数点后10位
	Trade_days          int       `bson:"trade_days"`          // 【交易天数】 本周股票实际交易的天数
	Period              string    `bson:"period"`              // 【K线周期】可能为 周K线，月K线，季K线，其值分别 W M Q
}

// 此对象存储到的数据库名称。
func (this *SDDBar) GetDBName() string {

	var dbName string
	// 从项目配置文件中获取

	switch this.Period {
	case "W":
		dbName = config.DBConfig.DBName.StockW1
	case "M":
		dbName = config.DBConfig.DBName.StockM1
	case "Q":
		dbName = config.DBConfig.DBName.StockQ1
	case "Y":
		dbName = config.DBConfig.DBName.StockY1
	default:
		dbName = config.DBConfig.DBName.StockD1
	}
	return dbName
}

// 此对象存储到数据库中的集合名称。
func (this *SDDBar) GetCollectionName() string {
	return this.Code
}

// 获取此对象在数据库集合中的键值。
func (this *SDDBar) MajorKey() time.Time {
	return this.Date.UTC()
}

// 返回股票的代码，上证股票以sh开头，深证股票以sz开头
func (this *SDDBar) GetCode() string {
	return this.Code
}

// 【交易日期】 本周市场的最后一个交易日
func (this *SDDBar) GetDate() time.Time {
	return this.Date.UTC()
}

// 返回本周期的开盘价
func (this *SDDBar) GetOpen() float64 {
	return this.Open
}

// 最高价
func (this *SDDBar) GetHigh() float64 {
	return this.High
}

// 最低价
func (this *SDDBar) GetLow() float64 {
	return this.Low
}

// 收盘价
func (this *SDDBar) GetClose() float64 {
	return this.Close
}

// 成交量
func (this *SDDBar) GetVolume() float64 {
	return this.Volume
}

// 成交额
func (this *SDDBar) GetMoney() float64 {
	return this.Money
}

// 【流通市值】 流通市值，流通市值 / 股价 = 流通股股本
func (this *SDDBar) GetTraded_Market_value() float64 {
	return this.Traded_market_value
}

// 【总市值】 总市值，总市值 / 股价 = 总股本
func (this *SDDBar) GetMarket_value() float64 {
	return this.Market_value
}

// 【换手率】 成交量 / 流通股本  = 成交量*股价/流通市值
func (this *SDDBar) GetTurnover() float64 {
	return this.Turnover
}

// 后复权价，复权开始时间为股票上市日，精确到小数点后10位。
func (this *SDDBar) GetAdjust_price() float64 {
	return this.Adjust_price
}

// 前复权价，复权开始时间为股票最近一个交易日，精确到小数点后10位
func (this *SDDBar) GetAdjust_price_f() float64 {
	return this.Adjust_price_f
}

// 【交易天数】 本周股票实际交易的天数
func (this *SDDBar) GetTrade_days() int {
	return this.Trade_days
}

// 【K线周期】可能为 周K线，月K线，季K线，其值分别 W M Q
func (this *SDDBar) GetPeriod() string {
	return this.Period
}
