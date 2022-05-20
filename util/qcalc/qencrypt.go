package qcalc

import (
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/iasuma/gf-lazy/os/qconst"
)

// Md5 md5加密
func Md5(data string, salt ...string) string {
	var saltStr string
	if len(salt) > 0 {
		saltStr = salt[0]
	}

	dataStr := fmt.Sprintf("%s%s", data, saltStr)
	encryptString, err := gmd5.EncryptString(dataStr)
	if err != nil {
		return ""
	}

	return encryptString
}

// SafeMd5 用SceneSecretKey作为salt进行Md5加密
func SafeMd5(data string) string {
	return Md5(data, qconst.SceneSecretKey)
}
