package file

import (
	"io/ioutil"
	"os"
)

//Overwrite 覆盖写入字符串到文件
func Overwrite(path, content string) {
	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		panic(err)
	}
}

//PathExists 路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
