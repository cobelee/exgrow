// ImportData_OneDay.go 将来自于“预测者”网站（www.yucezhe.com）的指定某天的日线数据导入mongo数据库。
package localdb

import (
	"bufio"
	o "exgrow/localdb/object"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
	ImportDD_FromDir() 导入 证券 以及 指数 单日日线数据

	将来自于“预测者”网站（www.yucezhe.com）的单日日线数据导入mongo数据库。
	将单日日线数据压缩文件包解压到exgrow/csv目录下，如 2015-04-14 data.csv等文件。
	执行 exgrow db idd 命令，可将 exgrow/csv 目录下的 csv 数据文件中的数据导入到mongo数据库中。
	该命令能够自动识别是否为目标数据文件的类型，数据包括证券、指数的日线数据。
*/
func ImportDD_FromDir() {
	var dir string = "data/daily_data/"
	dir_list, e := ioutil.ReadDir(dir)
	if e != nil {
		fmt.Println("read dir error!")
		return
	}

	filesCount := len(dir_list) // 获取 csv/ 目录下的总文件数
	for i, v := range dir_list {
		if v.IsDir() == false {
			filePath := dir + v.Name()

			fmt.Printf("Importing file %s...", filePath)

			if IsSDD_File(filePath) {
				err := ImportFromCSV(filePath, o.ParseToDBObject)

				if err == nil {
					// fmt.Fprintf(os.Stdout, "Handled %v csv file(s). Stock Daily Data is imported.\n", sddCount)
				}
			}

			if IsIDD_File(filePath) {
				err := ImportFromCSV(filePath, o.ParseToDBObject)

				if err == nil {
					// fmt.Fprintf(os.Stdout, "Handled %v csv file(s). Index Daily Data is imported.\n", sddCount)
				}
			}

			fmt.Printf("    Completed.  (%d/%d)\r\n", i+1, filesCount)

		}
	}

}

/* IsSDD_File (IsStockDailyDataFile) 判断指定文件是否是 证券 单日日线数据文件。

   返回 true | false, true表示指定文件是证券单日日线数据文件。
*/
func IsSDD_File(filePath string) bool {
	var isSDD bool = false

	f, e := os.Open(filePath)
	if e != nil {
		return false
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	var indexLine int = 0
	var includeCodeInFirstLine bool = false  // 第一行包含code
	var includeIndexInFirstLine bool = false // 第一行包含Index
	var isLen20InFirstLine bool = false      // 第一行有19个字段
	var isDiffInLine23 bool = false          // 第二行和第三行的第一个字段不同
	var fwInLine2 string                     // 第二行的第一个字段
	var fwInLine3 string                     // 第三行的第一个字段

	for {
		indexLine++

		line, err := reader.ReadBytes('\n')
		if err != nil {
			isSDD = false
			break
		}

		if indexLine == 1 {
			if strings.Index(string(line), "code") >= 0 {
				includeCodeInFirstLine = true
			}
			if strings.Index(string(line), "index") >= 0 {
				includeIndexInFirstLine = true
			}
			var arrayStr []string
			arrayStr = strings.Split(string(line), ",")
			if len(arrayStr) == 20 {
				isLen20InFirstLine = true
			}

			if includeCodeInFirstLine == false && includeIndexInFirstLine == false {
				isSDD = false
				return isSDD
			}
		}

		if indexLine == 2 {
			var arrayStr []string
			arrayStr = strings.Split(string(line), ",")
			fwInLine2 = arrayStr[0]
		}

		if indexLine == 3 {
			var arrayStr []string
			arrayStr = strings.Split(string(line), ",")
			fwInLine3 = arrayStr[0]
			if fwInLine2 != fwInLine3 {
				isDiffInLine23 = true
			}
			break
		}

	}

	if includeCodeInFirstLine == true && includeIndexInFirstLine == false &&
		isLen20InFirstLine == true && isDiffInLine23 == true {
		isSDD = true
	}

	return isSDD
}

/*
	IsIDD_File  (IsIndexDailyDataFile) 判断指定文件是否是 指数 单日日线数据文件。

	返回 true | false, true表示指定文件是 指数 单日日线数据文件。
*/
func IsIDD_File(filePath string) bool {
	var isIndexDD bool = false

	f, e := os.Open(filePath)
	if e != nil {
		return false
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	var indexLine int = 0
	var includeCodeInFirstLine bool = false  // 第一行包含Code
	var includeIndexInFirstLine bool = false // 第一行包含Index
	var isLen9InFirstLine bool = false       // 第一行有9个字段
	var isDiffInLine23 bool = false          // 第二行和第三行的第一个字段不同
	var fwInLine2 string                     // 第二行的第一个字段
	var fwInLine3 string                     // 第三行的第一个字段

	for {
		indexLine++

		line, err := reader.ReadBytes('\n')
		if err != nil {
			isIndexDD = false
			break
		}

		if indexLine == 1 {
			if strings.Index(string(line), "code") >= 0 {
				includeCodeInFirstLine = true
			}
			if strings.Index(string(line), "index") >= 0 {
				includeIndexInFirstLine = true
			}
			var arrayStr []string
			arrayStr = strings.Split(string(line), ",")
			if len(arrayStr) == 9 {
				isLen9InFirstLine = true
			}

			if includeCodeInFirstLine == false && includeIndexInFirstLine == false {
				isIndexDD = false
				return isIndexDD
			}
		}

		if indexLine == 2 {
			var arrayStr []string
			arrayStr = strings.Split(string(line), ",")
			fwInLine2 = arrayStr[0]
		}

		if indexLine == 3 {
			var arrayStr []string
			arrayStr = strings.Split(string(line), ",")
			fwInLine3 = arrayStr[0]
			if fwInLine2 != fwInLine3 {
				isDiffInLine23 = true
			}
			break
		}

	}

	if includeCodeInFirstLine == true && includeIndexInFirstLine == true &&
		isLen9InFirstLine == true && isDiffInLine23 == true {
		isIndexDD = true
	}

	return isIndexDD
}
