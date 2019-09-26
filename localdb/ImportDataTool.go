/*
	9此包提供了用于将“预测者”网站csv格式行情数据导入mongo数据库的工具。
*/

// ImportDataTool.go 将来自于“预测者”网站（www.yucezhe.com）的证券历史日线数据导入mongo数据库。
package localdb

import (
	"bufio"
	"exgrow/errors"
	c "exgrow/localdb/config"
	"exgrow/localdb/dbhelp"
	o "exgrow/localdb/object"
	t "exgrow/localdb/tools"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	// "runtime"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	将来自于“预测者”网站（www.yucezhe.com）的证券历史日线数据导入mongo数据库。
	将历史日线数据压缩文件包解压到exgrow/data目录下，如 sh600000.csv等文件。
	执行 exgrow db i 命令，可将 exgrow/data目录下的所有 csv 文件中的数据导入到mongo数据库中。
*/
func ImportSHD_FromDir() {
	var dir string = "data/"
	dir_list, e := ioutil.ReadDir(dir)
	if e != nil {
		fmt.Println("read dir error!\n")
		return
	}

	// runtime.GOMAXPROCS(runtime.NumCPU())

	filesTotal := len(dir_list) // 获取总文件数
	// var c chan int
	// c = make(chan int, 8)
	for i, v := range dir_list {
		if v.IsDir() == false {
			path := dir + v.Name()
			// go func() {
			// 	runtime.Gosched()
			// 	e := ImportFromCSV(path, o.ParseToDBObject)
			// 	if e == nil {
			// 		c <- 1
			// 	}
			// }()
			ImportFromCSV(path, o.ParseToDBObject)
			fmt.Fprintf(os.Stdout, "Importing Stock History Data from %v. (%d / %d)\r", v.Name(), i+1, filesTotal)
		}
	}

	// var fileNum int = 0
	// for v := range c {
	// 	fileNum += v
	// 	fmt.Fprintf(os.Stdout, "Importing Stock History Data. (%d / %d)\r", fileNum, filesTotal)
	// 	if fileNum == filesTotal {
	// 		fmt.Println("Import completed.")
	// 		close(c)
	// 		break
	// 	}
	// }

}

/*
	ImportFromCSV 数据导入函数

	从股票历史行情数据文件，如：sz002623.csv，读取行数据，并转换为“股票日行情数据”结构体，
	存入mongo数据库。
	数据库名称为：StockMarketRawD1，
*/
func ImportFromCSV(filePath string, hookfn func([]byte) (o.DBDoc, error)) error {
	f, e := os.Open(filePath)
	if e != nil {
		e1 := errors.NewError("100002", "Open file error:"+e.Error())
		return &e1 // 直接返回e1是不对的，没有实现 error() 接口，&e1实现了error()接口。
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			break
		}

		obj, e := o.ParseToDBObject(line)
		if e == nil {
			SaveObjToC(obj) // obj 实现了DBObject接口
		}
	}
	return nil
}

// SaveObjToC 将Obj数据保存到mongo数据库
//
// @param  object DBObject
func SaveObjToC(obj o.DBDoc) error {
	var session *mgo.Session
	session = dbhelp.GetSession()
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

/* TypifyHeaderLine 将csv文件的首行加入数据类型说明
 * 此类型化方法，既可以类型化 Stock 类型csv文件，也可以类型化 Index 类型csv文件。
 * @Param csvFile string csv文件的路径
 */
func TypifyHeaderLine(csvFile string) error {
	f, e := os.Open(csvFile)
	if e != nil {
		e1 := errors.NewError("100002", "Open file error:"+e.Error())
		return &e1 // 直接返回e1是不对的，没有实现 error() 接口，&e1实现了error()接口。
	}
	defer f.Close()

	var strArray []string
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()

		// 针对未曾类型化的字段标题，进行类型化处理。
		if i == 1 && strings.Contains(line, "code") && !strings.Contains(line, "code.string()") {
			newLine := o.TypifyFields(line)
			strArray = append(strArray, newLine)
		} else {
			strArray = append(strArray, line)
		}

	}

	f1, e := os.Create(csvFile)
	defer f1.Close()
	w := bufio.NewWriter(f1)
	w.Write([]byte(strings.Join(strArray, "\n")))
	w.Flush()
	return nil
}

/*
	TypifyHeaderLine_FromDir 预处理“预测者”网站（www.yucezhe.com）的证券历史日线数据。
	将历史日线数据压缩文件包解压到exgrow/data目录下，如 sh600000.csv等文件。
	执行 exgrow db i 命令，可将 exgrow/data目录下的所有 csv 文件中的数据导入到mongo数据库中。
*/
func TypifyHeaderLine_FromDir() {

	type csvData struct {
		DataType string
		StoreDir string
	}

	var idata = []csvData{
		{"stock history data", "data/stock_h_data/"},
		{"index history data", "data/index_h_data/"},
	}

	for _, dir := range idata {

		dir_list, e := ioutil.ReadDir(dir.StoreDir)
		if e != nil {
			fmt.Println("read dir error!\n")
			return
		}

		filesTotal := len(dir_list) // 获取总文件数
		for i, csv := range dir_list {
			if csv.IsDir() == false {
				path := dir.StoreDir + csv.Name()
				TypifyHeaderLine(path)
				fmt.Fprintf(os.Stdout, "Typify header Line in %s csv file  %v. (%d / %d)\r", dir.DataType, csv.Name(), i+1, filesTotal)
			}
		}
		fmt.Printf("Typify %s completed.\n", dir.DataType)
	}

}

/*
*	MongoImportSHD_FromDir 利用mongoimport工具，将来自于“预测者”网站（www.yucezhe.com）的证券历史日线数据导入mongo数据库。
*	将历史日线数据压缩文件包解压到exgrow/data目录下，如 sh600000.csv等文件。
*	执行 exgrow db mishd 命令，可将 exgrow/data目录下的所有 csv 文件中的数据导入到mongo数据库中。
 */
func MongoImportSHD_FromDir() {
	var (
		dir    string = "data/stock_h_data/"
		dbName string = c.DBConfig.DBName.StockMarketRawD1
	)
	dir_list, e := ioutil.ReadDir(dir)
	if e != nil {
		fmt.Println("read dir error!\n")
		return
	}

	filesTotal := len(dir_list) // 获取总文件数
	for i, v := range dir_list {
		if v.IsDir() == false {
			path := dir + v.Name() // 相对路径
			MongoImport_FromCSV(path, dbName)
			fmt.Fprintf(os.Stdout, "MongoImporting Stock History Data from %v. (%d / %d)\r", v.Name(), i+1, filesTotal)
		}
	}
}

/*
*	MongoImportIHD_FromDir 利用mongoimport工具，将来自于“预测者”网站（www.yucezhe.com）的指数历史日线数据导入mongo数据库。
*	将历史日线数据压缩文件包解压到exgrow/data/index_h_data目录下，如 sh000001.csv等文件。
*	执行 exgrow db mii 命令，可将 exgrow/data/index_h_data 目录下的所有 csv 文件中的数据导入到mongo数据库中。
 */
func MongoImportIHD_FromDir() {
	var (
		dir    string = "data/index_h_data/"
		dbName string = c.DBConfig.DBName.IndexMarketRawD1
	)

	if dbName == "" {
		fmt.Println("未指定 dbName 的值。")
		return
	}

	dir_list, e := ioutil.ReadDir(dir)
	if e != nil {
		fmt.Println("read dir error!\n")
		return
	}

	filesTotal := len(dir_list) // 获取总文件数
	for i, v := range dir_list {
		if v.IsDir() == false {
			path := dir + v.Name() // 相对路径
			MongoImport_FromCSV(path, dbName)
			fmt.Fprintf(os.Stdout, "MongoImporting Index History Data from %v. (%d / %d)\r", v.Name(), i+1, filesTotal)
		}
	}
}

// 将csv中的数据导入数据库中
func MongoImport_FromCSV(csvFileWithPath string, dbName string) error {
	var (
		result []byte
		err    error
		cmd    *exec.Cmd

		cName string = t.GetFileNameOnly(csvFileWithPath) // 表名称

		arg1 string = "--db=" + dbName                                     // 指定数据库名称
		arg2 string = "--collection=" + cName                              // 指定集合名称
		arg3 string = "--type=csv"                                         // 指定要导入的文件类型
		arg4 string = "--headerline"                                       // 标识文件中是否包含标题行
		arg5 string = "--columnsHaveTypes"                                 // 标识标题行中是否包含了数据类型
		arg6 string = "--ignoreBlanks"                                     // 标识遇到空白行是否忽略导入
		arg7 string = "--file=/home/cobe/go/src/exgrow/" + csvFileWithPath // 要导入的文件名称，包括路径
		arg8 string = "--drop"                                             // 要导入的集合若在数据中存在，先销毁原集合
	)

	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("mongoimport", arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	if result, err = cmd.Output(); err != nil {
		fmt.Println(string(result))
		e := errors.NewError("100006", "Shell Command exec error."+err.Error())
		return &e
	}
	return nil
}
