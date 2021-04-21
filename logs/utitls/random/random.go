package random

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.Mutex
//生成随机盐
func GetRandSalt() string {

	var salt string

	randStr:="0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNPQRSTUVWXYZ"

	rand.Seed(time.Now().Unix())
	for i:=0;i<10 ;i++  {
		randIndex := rand.Intn(len(randStr))
		salt += randStr[randIndex:randIndex+1]
	}

	return salt
}

// 生成请求id
func GetRequstId() string  {
	defer wg.Unlock()
	wg.Lock()
	fmt.Println(time.Now().Nanosecond())
	format := time.Now().Format("20060102")
	return format + "_" + strconv.Itoa(int(time.Now().UnixNano()))
}
