package qenv

import (
	"github.com/iasuma/gf-lazy/util/qutil"
)

const DefaultBootEnvName = "boot.env" // 默认的启动配置环境变量名
const DefaultEnvInputVar = "app.env" //APP环境变量名 APP_ENV

var envProd = []interface{}{"prd", "prod", "production", "1", 1, "true", true} // 生产环境
var envPre = []interface{}{"pre", "staging"} //准生产环境
var envTest = []interface{}{"uat", "test"} // 测试环境
var envDev = []interface{}{"dev", "development"} //部署开发环境
var envLocal = []interface{}{"local"} //本地开发环境

// IsProd 判断是否生产环境
func IsProd(envName interface{}) bool  {
	if envName == "" {
		return true
	}

	if qutil.InSlice(envProd, envName) {
		return true
	}

	return false
}

// IsPre 判断是否准生产环境
func IsPre(envName interface{}) bool {
	if qutil.InSlice(envPre, envName) {
		return true
	}

	return false
}

// IsOnLine 判断是否线上环境数据（prd、uat，生产和准生产均为线上数据的情况）
func IsOnLine(envName interface{}) bool {
	allPrd := append(envProd, envPre...)

	if qutil.InSlice(allPrd, envName) {
		return true
	}

	return false
}

// IsUat 判断是否测试环境（uat、test）
func IsUat(envName interface{}) bool {
	if qutil.InSlice(envTest, envName) {
		return true
	}

	return false
}

// IsTest 判断是否测试环境（uat、test、dev，通常是服务器部署环境）
func IsTest(envName interface{}) bool {
	allTest := append(envTest, envDev...)
	if qutil.InSlice(allTest, envName) {
		return true
	}

	return false
}

// IsDev 判断是否开发环境
func IsDev(envName interface{}) bool {
	if qutil.InSlice(envDev, envName) {
		return true
	}

	return false
}

// IsLocal 判断是否本地开发环境
func IsLocal(envName interface{}) bool {
	if qutil.InSlice(envLocal, envName) {
		return true
	}

	return false
}