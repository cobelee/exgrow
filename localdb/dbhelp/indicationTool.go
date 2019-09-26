package dbhelp

import (
	"exgrow/errors"
	o "exgrow/localdb/object"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SaveIndicSetToC 将 IndicSet 数据保存到mongo数据库
//
// @param  object DBObject
func SaveIndicSetToC(obj o.DBDoc) error {
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
