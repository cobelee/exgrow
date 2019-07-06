package localdb

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session
var database *mgo.Database

// 初始化 MongoHelper
func init() {

	var err error

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"127.0.0.1"},
		Direct:    false,
		Timeout:   time.Second * 2,
		PoolLimit: 1024, // Session.SetPoolLimit
	}

	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Println(err.Error())
	}
	session.SetMode(mgo.Monotonic, true)

}

// 获取 mongo Session
func GetSession() *mgo.Session {
	return session
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}

// Upsert 将Obj数据更新插入到mongo数据库
func Upsert(dbName string, cName string, selector bson.M, obj interface{}) {
	var session *mgo.Session
	session = GetSession()

	c := session.DB(dbName).C(cName)

	c.Upsert(selector, obj)
}

// 获取指定数据库下的集合名列表
func GetCollectionNames(dbName string) (names []string, err error) {
	var session *mgo.Session
	session = GetSession()

	return session.DB(dbName).CollectionNames()
}
