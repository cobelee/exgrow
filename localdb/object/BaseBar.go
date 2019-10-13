package object

import (
	"time"
)

type BaseBar interface {
	GetCode() string                 // 返回股票的代码，上证股票以sh开头，深证股票以sz开头
	GetDate() time.Time              // 【交易日期】 本周市场的最后一个交易日
	GetOpen() float64                // 返回本周期的开盘价
	GetHigh() float64                // 最高价
	GetLow() float64                 // 最低价
	GetClose() float64               // 收盘价
	GetVolume() float64              // 成交量
	GetMoney() float64               // 成交额
	GetTraded_market_value() float64 // 【流通市值】 流通市值，流通市值 / 股价 = 流通股股本
	GetMarket_value() float64        // 【总市值】 总市值，总市值 / 股价 = 总股本
	GetTurnover() float64            // 【换手率】 成交量 / 流通股本  = 成交量*股价/流通市值
	GetAdjust_price() float64        // 后复权价，复权开始时间为股票上市日，精确到小数点后10位。
	GetAdjust_price_f() float64      // 前复权价，复权开始时间为股票最近一个交易日，精确到小数点后10位
	GetTrade_days() int              // 【交易天数】 本周股票实际交易的天数
	GetPeriod() string               // 【K线周期】可能为 周K线，月K线，季K线，其值分别 W M Q
}
