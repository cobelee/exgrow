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


*/
func price(sc string, pt string, t time.Time, p string) float64 {
	return 0
}

/* Query stock market information.
    para: sc   -- stockCode      For sample: sh600000
		  pt  --periodType       It could be D W M Q Y
	      t   --time 			date of period
*/
func GetSDDBar(sc string, pt string, t time.Time) o.SDDBar {
	var session *mgo.Session
	session = h.GetSession()

	var dbName string
	switch pt {
	case "D":
		dbName = c.DBConfig.DBName.StockMarketRawD1
	case "W":
		dbName = c.DBConfig.DBName.StockMarketW1
	case "M":
		dbName = c.DBConfig.DBName.StockMarketM1
	case "Q":
		dbName = c.DBConfig.DBName.StockMarketQ1
	case "Y":
		dbName = c.DBConfig.DBName.StockMarketY1
	default:
		dbName = c.DBConfig.DBName.StockMarketRawD1
	}
	c := session.DB(dbName).C(sc)
	var bar o.SDDBar

	c.Find(
		bson.M{
			"date": bson.M{
				"$eq": t,
			},
		}).One(&bar)
	return bar
}
