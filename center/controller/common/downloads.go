package common

import (
	"fmt"
	"io/ioutil"
	"lhc.go.game.center/logs"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Downloads(urls string,dir string) string  {
	parse, err := url.Parse(urls)
	if !parse.IsAbs() || err!=nil {
		return ""
	}
	resp, err := http.Get(urls)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	lastIndex := strings.LastIndex(urls,"/")
	reDir := "/static/download/"+dir
	logDirPath := "."+reDir
	filename := logDirPath+urls[lastIndex:]

	//判断目录是否存在
	if !FileExist(logDirPath) {
		err := os.MkdirAll(logDirPath, os.ModePerm)
		if err != nil {
			logs.Corn.Info(err)
			fmt.Printf("Create img dir %s error: %s\n", logDirPath, err)
			return ""
		}
	}
	//判断文件是否存在
	if !FileExist(filename) {
		if err := ioutil.WriteFile(filename, body, 0644);err != nil {
			logs.Corn.Error("图片写入错误：",err)
			fmt.Printf("write img  %s error: %s\n", filename, err)
			return ""
		}
	}

	return reDir+urls[lastIndex:]
}
