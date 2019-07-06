package testcode

import (
	"fmt"
	"io/ioutil"
)

// ReadFiles 样例：读取指定文件夹中的所有文件列表
func ReadFiles() {
	dir_list, e := ioutil.ReadDir("data/")
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	for i, v := range dir_list {
		if v.IsDir() == false {
			fmt.Println(i, " = ", v.Name())
		}
	}
}
