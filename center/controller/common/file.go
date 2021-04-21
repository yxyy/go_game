package common

import (
	"fmt"
	"os"
)

//判断文件是否存在
func FileExist(name string) bool {
	_, err := os.Stat(name)
	if err!=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CheckFileExists(path string) bool {
	fmt.Println(path)
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return false
		}
	}
	return true
}
