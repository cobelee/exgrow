package dbhelp

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 指定数据库名和指定集合名，获取集合数据。
func RemoveAll(dbName string, longCode string, sinceDate time.Time) {
	var session *mgo.Session
	session = GetSession().Copy()
	defer session.Close()

	c := session.DB(dbName).C(longCode)
	c.RemoveAll(bson.M{"date": bson.M{"$gte": sinceDate}})
}
