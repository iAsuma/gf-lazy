package qcfg

import (
	"context"
	"github.com/gogf/gf/v2/os/gcfg"
)

const (
	DefaultAppEnvConfigName = "server.appEnv"
	DefaultConfigDirName = "config" //默认的配置目录名
	DefaultBootConfigName = "boot" //默认的启动配置文件名
	DefaultFileType = "yaml" //默认的配置文件格式
	DefaultRuntimeDirName = "runtime" //默认的缓存目录
	DefaultLogDirName = "runtime/log" //默认的日志目录
)

type NacosTicket struct {
	IpAddr 				string
	Port 				string
	DataId 				string
	NamespaceId 		string
	Group 				string
	AccessKey           string
	SecretKey           string
	CacheDir			string
	LogDir				string
	LogLevel		    string
}

type Config struct {
	gcfg.Config
}

func Instance(name ...string) *Config {
	return &Config{
		*gcfg.Instance(name...),
	}
}

func (c *Config) GetString(ctx context.Context, pattern string, def ...interface{}) string {
	config, err := c.Get(ctx, pattern, def...)
	if err != nil {
		return ""
	}

	return config.String()
}

func (c *Config) GetToString(pattern string, def ...interface{}) string {
	return c.GetString(context.TODO(), pattern, def...)
}