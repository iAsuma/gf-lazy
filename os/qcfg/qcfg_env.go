package qcfg

import (
	"github.com/iasuma/gf-lazy/os/qenv"
)

// EnvIsProd 判断是否生产环境
func EnvIsProd(env ...interface{}) bool  {
	envName := interface{}(nil)
	if len(env) > 0 {
		envName = env[0]
	} else {
		envName = GetToString(DefaultAppEnvConfigName)
	}

	return qenv.IsProd(envName)
}

// EnvIsPre 判断是否准生产环境
func EnvIsPre(env ...interface{}) bool {
	envName := interface{}(nil)
	if len(env) > 0 {
		envName = env[0]
	} else {
		envName = GetToString(DefaultAppEnvConfigName)
	}

	return qenv.IsPre(envName)
}

// EnvIsOnLine 判断是否线上环境数据（prd、uat，生产和准生产均为线上数据的情况）
func EnvIsOnLine(env ...interface{}) bool {
	envName := interface{}(nil)
	if len(env) > 0 {
		envName = env[0]
	} else {
		envName = GetToString(DefaultAppEnvConfigName)
	}

	return qenv.IsOnLine(envName)
}

// EnvIsUat 判断是否测试环境（uat、test）
func EnvIsUat(env ...interface{}) bool {
	envName := interface{}(nil)
	if len(env) > 0 {
		envName = env[0]
	} else {
		envName = GetToString(DefaultAppEnvConfigName)
	}

	return qenv.IsUat(envName)
}

// EnvIsTest 判断是否测试环境（uat、test、dev，通常是服务器部署环境）
func EnvIsTest(env ...interface{}) bool {
	envName := interface{}(nil)
	if len(env) > 0 {
		envName = env[0]
	} else {
		envName = GetToString(DefaultAppEnvConfigName)
	}

	return qenv.IsTest(envName)
}

// EnvIsDev 判断是否开发环境
func EnvIsDev(env ...interface{}) bool {
	envName := interface{}(nil)
	if len(env) > 0 {
		envName = env[0]
	} else {
		envName = GetToString(DefaultAppEnvConfigName)
	}

	return qenv.IsDev(envName)
}

// EnvIsLocal 判断是否本地开发环境
func EnvIsLocal(env ...interface{}) bool {
	envName := interface{}(nil)
	if len(env) > 0 {
		envName = env[0]
	} else {
		envName = GetToString(DefaultAppEnvConfigName)
	}

	return qenv.IsLocal(envName)
}