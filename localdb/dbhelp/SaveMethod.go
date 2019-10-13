package dbhelp

import (
	"exgrow/errors"
	o "exgrow/localdb/object"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBDoc interface {
	GetDBName() string
	GetCollectionName() string
	MajorKey() time.Time
}

// SaveObjToC 将Obj数据保存到mongo数据库
//
// @param  object DBObject
func SaveObjToC(obj DBDoc) error {
	var session *mgo.Session
	session = GetSession()
	dbName := obj.GetDBName()
	if dbName == "" {
		err := errors.NewError("100007", "Have not identify dbName.")
		return &err
	}
	cName := obj.GetCollectionName()
	c := session.DB(dbName).C(cName)

	seletor := bson.M{"date": obj.MajorKey()}
	c.Upsert(seletor, obj) // 采用upsert防止重复插入。
	return nil
}

// --------------------------------------------------------------------------

type AnaDBDoc interface {
	GetDBName() string
	GetCollectionName() string
	MajorKey() string
}

// SaveObjToC 将Obj数据保存到mongo数据库
//
// @param  object DBObject
func SaveAnaObjToC(obj AnaDBDoc) error {
	var session *mgo.Session
	session = GetSession()
	dbName := obj.GetDBName()
	if dbName == "" {
		err := errors.NewError("100007", "Have not identify dbName.")
		return &err
	}
	cName := obj.GetCollectionName()
	c := session.DB(dbName).C(cName)

	seletor := bson.M{"code": obj.MajorKey()}
	c.Upsert(seletor, obj) // 采用upsert防止重复插入。
	return nil
}

// --------------------------------------------------------------------------

// ParseToDBObject 将文本格式的数据行，转化为数据对象。
// 该方法用于将cvs文本数据，转化为程序数据对象，方便导入数据库。
//
// @param line []byte byte类型的文本数据行
func ParseToDBObject(line []byte) (DBDoc, error) {
	var array []string
	// 去除换行符
	strLine := string(line)
	strLine = strings.Replace(strLine, "\n", "", -1)
	array = strings.Split(strLine, ",")

	if strings.Contains(array[0], "code") {
		e := errors.NewError("100000", "This line is comment text, cannot be parsed to DBObject.")
		return nil, &e
	}

	var fieldCount = len(array)

	if fieldCount < 9 {
		e := errors.NewError("100005", "This is not data line, cannot be parsed to DBObject.")
		return nil, &e
	}

	if fieldCount == 20 {
		var sd o.SDBar

		sd.Code = array[0]
		sd.Date, _ = time.Parse("2006-01-02", array[1])
		sd.Open, _ = strconv.ParseFloat(array[2], 64)
		sd.High, _ = strconv.ParseFloat(array[3], 64)
		sd.Low, _ = strconv.ParseFloat(array[4], 64)
		sd.Close, _ = strconv.ParseFloat(array[5], 64)
		sd.Change, _ = strconv.ParseFloat(array[6], 64)
		sd.Volume, _ = strconv.ParseFloat(array[7], 64)
		sd.Money, _ = strconv.ParseFloat(array[8], 64)
		sd.Traded_market_value, _ = strconv.ParseFloat(array[9], 64)
		sd.Market_value, _ = strconv.ParseFloat(array[10], 64)
		sd.Turnover, _ = strconv.ParseFloat(array[11], 64)
		sd.Adjust_price, _ = strconv.ParseFloat(array[12], 64)
		sd.Report_type = array[13]
		sd.Report_date, _ = time.Parse("2006-01-02", array[14])
		sd.PE_TTM, _ = strconv.ParseFloat(array[15], 64)
		sd.PS_TTM, _ = strconv.ParseFloat(array[16], 64)
		sd.PC_TTM, _ = strconv.ParseFloat(array[17], 64)
		sd.PB, _ = strconv.ParseFloat(array[18], 64)
		sd.Adjust_price_f, _ = strconv.ParseFloat(array[19], 64)

		return &sd, nil
	}

	if fieldCount == 9 {
		var idd o.IDBar
		idd.Index_code = array[0]
		idd.Date, _ = time.Parse("2006-01-02", array[1])
		idd.Open, _ = strconv.ParseFloat(array[2], 64)
		idd.Close, _ = strconv.ParseFloat(array[3], 64)
		idd.Low, _ = strconv.ParseFloat(array[4], 64)
		idd.High, _ = strconv.ParseFloat(array[5], 64)
		idd.Volume, _ = strconv.ParseFloat(array[6], 64)
		idd.Money, _ = strconv.ParseFloat(array[7], 64)
		idd.Change, _ = strconv.ParseFloat(array[8], 64)

		return &idd, nil
	}

	e := errors.NewError("100000", "This line is comment text, cannot be parsed to DBObject.")
	return nil, &e

}
