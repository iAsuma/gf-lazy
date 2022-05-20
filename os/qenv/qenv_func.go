package qenv

import (
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetString 获取环境变量中的值，支持从命令行参数中获取
func GetString(key string, def ...interface{}) string {
	value := interface{}(nil)
	if len(def) > 0 {
		value = def[0]
	}

	env := genv.GetWithCmd(key, value)
	return gconv.String(env)
}