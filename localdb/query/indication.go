package query

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/* Query price data of stock.
    para: sc  -- stockCode  For sample: sh600000
		pt  --periodType   It could be D W M Q Y
		t   --time 			date of period
		p   --priceType    It could be "open", "close", "high", "low"


/* Query indication data of stock.
    para: sc   -- stockCode      For sample: sh600000
		  pt  --periodType       It could be D W M Q Y
	      t   --time 			date of period
*/
func GetIndicBar(sc string, pt string, d time.Time) o.IndicBar {
	var session *mgo.Session
	session = h.GetSession()

	var dbName string
	switch pt {
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
	c := session.DB(dbName).C(sc)
	var bar o.IndicBar

	c.Find(
		bson.M{
			"date": bson.M{
				"$eq": d,
			},
		}).One(&bar)
	return bar
}
