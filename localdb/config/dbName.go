package config

/*  根据时间周期，获取不同周期Stock数据库的名称。

    由于该功能使用频率高，故单独写成函数。
*/
func GetStockDBNameByPeriodType(pt PeriodType) string {

	var dbName string

	switch pt {
	case Period_D:
		dbName = DBConfig.DBName.StockD1
	case Period_W:
		dbName = DBConfig.DBName.StockW1
	case Period_M:
		dbName = DBConfig.DBName.StockM1
	case Period_Q:
		dbName = DBConfig.DBName.StockQ1
	case Period_Y:
		dbName = DBConfig.DBName.StockY1
	default:
		dbName = ""
	}

	return dbName

}
