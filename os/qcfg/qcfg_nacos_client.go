package qcfg

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"strconv"
)

type nacosClient struct {
	config NacosTicket
	client config_client.IConfigClient
}

func NacosClient(config NacosTicket) *nacosClient {
	var c = new(nacosClient)
	c.instance(config)
	return c
}

func (n *nacosClient) instance(config NacosTicket) config_client.IConfigClient {
	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	ipAddr := config.IpAddr
	port := gconv.Uint64(config.Port)
	endpoint := fmt.Sprintf("%s:%s", ipAddr, strconv.FormatUint(port, 10))
	namespaceId := config.NamespaceId
	cacheDir := config.CacheDir
	logDir := config.LogDir
	logLevel := config.LogLevel

	// 推荐使用 RAM 用户的 accessKey、secretKey
	//var accessKey = config.AccessKey
	//var secretKey = config.SecretKey

	sc := []constant.ServerConfig{{
		IpAddr: ipAddr,
		Port:   port,
	}}

	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint,
		NamespaceId:    namespaceId,
		//AccessKey:      accessKey,
		//SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		//ListenInterval: 30 * 1000,
		CacheDir:	cacheDir,
		LogDir: logDir,
		LogLevel: logLevel,
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig": clientConfig,
	})

	if err != nil {
		panic(err)
	}

	n.config = config
	n.client = client
	return client
}

func (n *nacosClient) GetConfig() string {
	client := n.client
	config := n.config

	var dataId = config.DataId
	var group = config.Group

	// 获取配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	if err != nil {
		panic(err)
	}

	return content
}

func (n *nacosClient) ListenConfig(callback func(string)) (err error) {
	client := n.client
	config := n.config

	err = client.ListenConfig(vo.ConfigParam{
		DataId: config.DataId,
		Group:  config.Group,
		OnChange: func(namespace, group, dataId, data string) {
			callback(data)
		},
	})

	return err
}

