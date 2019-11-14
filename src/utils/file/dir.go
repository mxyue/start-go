package file

import (
	"os"
)

//Mkdir 创建目录
func Mkdir(path string) {
	if ok, err := PathExists(path); !ok || err != nil {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			panic(err)
		}
	}
}
