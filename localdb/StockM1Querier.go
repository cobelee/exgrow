package localdb

import (
	// "fmt"
	"exgrow/errors"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type Card interface {
	GetDBName() string
	GetCollectionName() string
	GetKeyValue() time.Time
	GetFirstDate() (time.Time, error)
}

// 单个股票历史周K线数据集合结构体
type SMCard struct {
	ShortCode string   // [股票短代码]如 600000
	LongCode  string   // [股票长代码]如 sh600000
	Matrix    []SDDBar // 指定代码的股票全生命期数据矩阵
}

// 此结构体数据的数据库名称。
func (this *SMCard) GetDBName() string {
	// config := GetConfig() // 获取项目配置文件
	// return config.DBName.StockMarketRawD1
	return DBConfig.DBName.StockMarketM1
}

// 此对象存储到数据库中的集合名称。
func (this *SMCard) GetCollectionName() string {
	return this.LongCode
}

// 创建股票数据卡片实例
func CreateSMCard(longCode string) SMCard {
	var stock SMCard
	stock.LongCode = longCode                           // 初始化长代码
	stock.ShortCode = strings.TrimLeft(longCode, "shz") // 初始化短代码
	stock.Matrix = GetSMMatrix(longCode)
	return stock
}

/* 获取股票全生命历史数据
 * 数据按日期从古到今进行排序
 */
func GetSMMatrix(longCode string) []SDDBar {
	var session *mgo.Session
	session = GetSession().Copy()
	defer session.Close()

	dbName := DBConfig.DBName.StockMarketM1
	c := session.DB(dbName).C(longCode)
	var result []SDDBar
	c.Find(nil).Sort("date").All(&result)

	return result
}

func (this *SMCard) GetFirstDate() (time.Time, error) {
	var date time.Time
	if this.Matrix != nil && len(this.Matrix) > 0 {
		date = this.Matrix[0].Date
		return date, nil
	}
	var err = errors.NewError("100008", "The matrix data is valid.")
	return date, &err
}
