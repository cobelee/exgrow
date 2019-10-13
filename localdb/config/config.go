package config

import (
	"github.com/BurntSushi/toml"
)

//订制配置文件解析载体
type config struct {
	DBName *DBName `toml:"dbName"`
}

//订制Database块
type DBName struct {
	Analysis         string
	StockMarketRawD1 string
	StockD1          string
	StockW1          string
	StockM1          string
	StockQ1          string
	StockY1          string
	IndexMarketRawD1 string
	IndexD1          string
	IndexW1          string
	IndexM1          string
	IndexQ1          string
	IndexY1          string
}

// 实例化后的数据库配置信息对象
var DBConfig config

func init() {
	DBConfig = GetConfig()
}

//订制SQL语句结构
/*
type SQL struct{
    Sql1 string `toml:"sql_1"`
    Sql2 string `toml:"sql_2"`
    Sql3 string `toml:"sql_3"`
    Sql4 string `toml:"sql_4"`
}
*/
func GetConfig() config {
	var config config
	filePath := "localdb/config/config.toml"
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}
	return config
}
