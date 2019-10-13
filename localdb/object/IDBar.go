package object

import (
	"exgrow/localdb/config"
	"time"
)

// 指数日行情数据，对应 CSV 文件中的单行数据。
type IDBar struct {
	Index_code string    `bson:"index_code"` // 指数的代码，上证股票以sh开头，深证股票以sz开头
	Date       time.Time `bson:"date"`       // 交易日期
	Open       float64   `bson:"open"`       // 开盘价
	Close      float64   `bson:"close"`      // 收盘价
	Low        float64   `bson:"low"`        // 最低价
	High       float64   `bson:"high"`       // 最高价
	Volume     float64   `bson:"volume"`     // 成交量
	Money      float64   `bson:"money"`      // 成交额
	Change     float64   `bson:"change"`     // 涨跌幅
}

// 此对象存储到的数据库名称。
func (this *IDBar) GetDBName() string {
	// config := GetConfig() // 获取项目配置文件
	// return config.DBName.IndexMarketRawD1
	return config.DBConfig.DBName.IndexMarketRawD1
}

// 此对象存储到数据库中的集合名称。
func (this *IDBar) GetCollectionName() string {
	return this.Index_code
}

// 获取此对象在数据库集合中的键值。
func (this *IDBar) MajorKey() time.Time {
	return this.Date
}
