//

package keltner

import (
	c "exgrow/localdb/config"
)

type KCRecord struct {
	Code   string  `bson:"code"`  // The stock code. such as 'sh600000', 上证股票以sh开头，深证股票以sz开头
	CloseT float64 `bson:"closet` // The true close price that not adjusted.
	KCD    float64 `bson:"kcd"`   // Keltner Channels --- Daily indication.
	KCW    float64 `bson:"kcw"`   // Keltner Channels --- Weekly indication.
	KCM    float64 `bson:"kcm"`   // Keltner Channels --- Monthly indication.
	KCQ    float64 `bson:"kcq"`   // Keltner Channels --- Quarterly indication.
	KCY    float64 `bson:"kcy"`   // Keltner Channels --- Yearly indication.
}

// 此对象存储到的数据库名称。
func (this *KCRecord) GetDBName() string {
	// 从项目配置文件中获取
	dbName := c.DBConfig.DBName.Analysis
	return dbName
}

// 此对象存储到数据库中的集合名称。
func (this *KCRecord) GetCollectionName() string {
	return "keltner"
}

// 获取此对象在数据库集合中的键值。
func (this *KCRecord) MajorKey() string {
	return this.Code
}
