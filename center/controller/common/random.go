package common

import (
	"math/rand"
	"time"
)

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
