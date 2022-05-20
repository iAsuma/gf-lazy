package qcfg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/iasuma/gf-lazy/util/qutil"
	"os"
)

// GetString 获取配置 to string
func GetString(ctx context.Context, pattern string, def ...interface{}) string {
	return Instance().GetString(ctx, pattern, def...)
}

// GetToString 获取配置 to string
func GetToString(pattern string, def ...interface{}) string {
	return Instance().GetString(context.TODO(), pattern, def...)
}

// GetConfigDir 获取运行配置完整目录
func GetConfigDir() string {
	return qutil.GetCurrentPath() + string(os.PathSeparator)  + DefaultConfigDirName
}

// GetConfigFileName 获取完整的配置文件名
func GetConfigFileName(filename ...string) string {
	key := gcfg.DefaultConfigFileName
	if len(filename) > 0 && filename[0] != "" {
		key = filename[0]
	}

	return fmt.Sprintf(`%s.%s`, key, DefaultFileType)
}

// GetRuntimeDir 获取框架缓存目录
func GetRuntimeDir() string {
	return qutil.GetCurrentPath() + string(os.PathSeparator) + DefaultRuntimeDirName
}

// GetLogDir 获取框架日志目录
func GetLogDir() string {
	return  qutil.GetCurrentPath() + string(os.PathSeparator) + DefaultLogDirName
}

// SetGoFrameConfig 设置GoFrame框架配置
func SetGoFrameConfig(configFile string, content string)  {
	workDir := qutil.GetCurrentPath() + string(os.PathSeparator)

	// 禁止使用工作空间根目录配置文件 -- 根目录下的config.yaml会覆盖config/config.yaml的配置
	c1 := workDir + GetConfigFileName()
	if qutil.IsFile(c1) {
		_ = gfile.Remove(c1)
	}

	// 获取配置目录
	configDir := GetConfigDir()

	configFileName := configDir + string(os.PathSeparator) + configFile
	qutil.WriteFile(configFileName, content)
	gcfg.Instance().GetAdapter().(*gcfg.AdapterFile).SetFileName(configFile)
}

