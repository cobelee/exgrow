package localdb

import (
	// "fmt"
	"exgrow/errors"
	c "exgrow/localdb/config"
	"exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

// 单个股票历史月K线数据集合结构体
type SMCard struct {
	ShortCode string     // [股票短代码]如 600000
	LongCode  string     // [股票长代码]如 sh600000
	Matrix    []o.SDDBar // 指定代码的股票全生命期数据矩阵
}

type Card interface {
	GetDBName() string
	GetCollectionName() string
	GetKeyValue() time.Time
	GetFirstDate() (time.Time, error)
}

// 创建股票数据卡片实例
func CreateSMCard(longCode string) SMCard {
	var stock SMCard
	stock.LongCode = longCode                           // 初始化长代码
	stock.ShortCode = strings.TrimLeft(longCode, "shz") // 初始化短代码
	dbName := c.DBConfig.DBName.StockMarketM1
	stock.Matrix = GetSMMatrix(dbName, longCode)
	return stock
}

// 此结构体数据的数据库名称。
func (this *SMCard) GetDBName() string {
	// config := GetConfig() // 获取项目配置文件
	// return config.DBName.StockMarketRawD1
	return c.DBConfig.DBName.StockMarketM1
}

// 此对象存储到数据库中的集合名称。
func (this *SMCard) GetCollectionName() string {
	return this.LongCode
}

/* 获取股票全生命历史数据
 * 数据按日期从古到今进行排序
 */
func GetSMMatrix(dbName string, longCode string) []o.SDDBar {
	var session *mgo.Session
	session = dbhelp.GetSession().Copy()
	defer session.Close()

	c := session.DB(dbName).C(longCode)
	var result []o.SDDBar
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
