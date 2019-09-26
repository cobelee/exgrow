package tools

import (
	"path"
	"strings"
)

/* 获取删除扩展名后的文件名
*
 */
func GetFileNameOnly(fullFilename string) string {
	filenameWithSuffix := path.Base(fullFilename)               //获取文件名带后缀
	ext := path.Ext(filenameWithSuffix)                         //获取文件后缀
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, ext) //获取文件名
	return filenameOnly
}
