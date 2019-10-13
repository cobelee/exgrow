package object

import (
	"exgrow/localdb/config"
	"strings"
	"time"
)

// --------------------------------------------------------------------------

/* 股票单根K线日行情数据，对应 CSV 文件中的单行数据。
 */
type SDBar struct {
	Code                string    `bson:"code"`                // 股票的代码，上证股票以sh开头，深证股票以sz开头
	Date                time.Time `bson:"date"`                // 交易日期
	Open                float64   `bson:"open"`                // 开盘价
	High                float64   `bson:"high"`                // 最高价
	Low                 float64   `bson:"low"`                 // 最低价
	Close               float64   `bson:"close"`               // 收盘价
	Change              float64   `bson:"change"`              // 涨跌幅，复权之后的真实涨跌幅，保证准确。
	Volume              float64   `bson:"volume"`              // 成交量
	Money               float64   `bson:"money"`               // 成交额
	Traded_market_value float64   `bson:"traded_market_value"` // 流通市值
	Market_value        float64   `bson:"market_value"`        // 总市值
	Turnover            float64   `bson:"turnover"`            // 换手率
	Adjust_price        float64   `bson:"adjust_price"`        // 后复权价，复权开始时间为股票上市日，精确到小数点后10位。
	Report_type         string    `bson:"report_type"`         // 最近一期财务报告的类型，3-31对应一季度，6-30对应半年报，9-30对应三季度，12-31对应年报。
	Report_date         time.Time `bson:"report_date"`         // 最近一期财务报告实际发布的日期。
	PE_TTM              float64   `bson:"PE_TTM"`              // 最近12个月市盈率，股价/最近12个月归属母公司的每股收益TTM
	PS_TTM              float64   `bson:"PS_TTM"`              // 最近12个月市销率，股价/最近12个月每股营业收入。
	PC_TTM              float64   `bson:"PC_TTM"`              // 最近12个月市现率，股价/最近12个月每股经营现金流。
	PB                  float64   `bson:"PB"`                  // 市净率，股价/最近期财报每股净资产
	Adjust_price_f      float64   `bson:"adjust_price_f"`      // 前复权价，复权开始时间为股票最近一个交易日，精确到小数点后10位
}

// 此对象存储到的数据库名称。
func (this SDBar) GetDBName() string {
	// config := GetConfig() // 获取项目配置文件
	// return config.DBName.StockMarketRawD1
	return config.DBConfig.DBName.StockMarketRawD1
}

// 此对象存储到数据库中的集合名称。
func (this SDBar) GetCollectionName() string {
	return this.Code
}

// 获取此对象在数据库集合中的键值。
func (this SDBar) MajorKey() time.Time {
	return this.Date.UTC()
}

func (this SDBar) GetSDDBar() SDDBar {
	return SDDBar{
		this.Code,  // 返回股票的代码，上证股票以sh开头，深证股票以sz开头
		this.Date,  // 【交易日期】 交易日
		this.Open,  // 返回本周期的开盘价
		this.High,  // 最高价
		this.Low,   // 最低价
		this.Close, // 收盘价
		this.Change,
		this.Volume,              // 成交量
		this.Money,               // 成交额
		this.Traded_market_value, // 【流通市值】 流通市值，流通市值 / 股价 = 流通股股本
		this.Market_value,        // 【总市值】 总市值，总市值 / 股价 = 总股本
		this.Turnover,            // 【换手率】 成交量 / 流通股本  = 成交量*股价/流通市值
		this.Adjust_price,        // 后复权价，复权开始时间为股票上市日，精确到小数点后10位。
		this.Adjust_price_f,      // 前复权价，复权开始时间为股票最近一个交易日，精确到小数点后10位
		1,                        // 【交易天数】 本周股票实际交易的天数
		"D",                      // 【K线周期】可能为 周K线，月K线，季K线，其值分别 W M Q
	}
}

// 返回股票的代码，上证股票以sh开头，深证股票以sz开头
func (this SDBar) GetCode() string {
	return this.Code
}

// 【交易日期】 交易日
func (this SDBar) GetDate() time.Time {
	return this.Date.UTC()
}

// 返回本周期的开盘价
func (this SDBar) GetOpen() float64 {
	return this.Open
}

// 最高价
func (this SDBar) GetHigh() float64 {
	return this.High
}

// 最低价
func (this SDBar) GetLow() float64 {
	return this.Low
}

// 收盘价
func (this SDBar) GetClose() float64 {
	return this.Close
}

// 成交量
func (this SDBar) GetVolume() float64 {
	return this.Volume
}

// 成交额
func (this SDBar) GetMoney() float64 {
	return this.Money
}

// 【流通市值】 流通市值，流通市值 / 股价 = 流通股股本
func (this SDBar) GetTraded_market_value() float64 {
	return this.Traded_market_value
}

// 【总市值】 总市值，总市值 / 股价 = 总股本
func (this SDBar) GetMarket_value() float64 {
	return this.Market_value
}

// 【换手率】 成交量 / 流通股本  = 成交量*股价/流通市值
func (this SDBar) GetTurnover() float64 {
	return this.Turnover
}

// 后复权价，复权开始时间为股票上市日，精确到小数点后10位。
func (this SDBar) GetAdjust_price() float64 {
	return this.Adjust_price
}

// 前复权价，复权开始时间为股票最近一个交易日，精确到小数点后10位
func (this SDBar) GetAdjust_price_f() float64 {
	return this.Adjust_price_f
}

// 【交易天数】 本周股票实际交易的天数
func (this SDBar) GetTrade_days() int {
	return 1
}

// 【K线周期】可能为 周K线，月K线，季K线，其值分别 W M Q
func (this SDBar) GetPeriod() string {
	return "D"
}

// ---------------------------------------------------------------------------

/* TypifyFieldsLine 将以逗号间隔的 fields， 加入数据类型说明。

 输入的原始数据：code,date,open,high
 格式化后输出为：code.string(),date.date(2006-01-02),open.double(),high.double(),
	该方法对StockData， IndexData文件的标题行转换都适用。
*/
func TypifyFields(line string) string {
	var arrOld []string
	arrOld = strings.Split(line, ",")
	arrNew := make([]string, len(arrOld))

	if !strings.Contains(arrOld[0], "code") {
		return ""
	}

	for i, v := range arrOld {
		switch {
		case strings.Contains(v, "code"):
			arrNew[i] = v + ".string()"
		case strings.Contains(v, "date"):
			arrNew[i] = v + ".date(2006-01-02)"
		case strings.Contains(v, "value"):
			arrNew[i] = v + ".double()"
		case strings.Contains(v, "price"):
			arrNew[i] = v + ".double()"
		case v == "report_type":
			arrNew[i] = v + ".string()"
		case strings.Contains(v, "TTM"):
			arrNew[i] = v + ".double()"
		default:
			arrNew[i] = v + ".double()"
		}
	}
	return strings.Join(arrNew, ",")

}
