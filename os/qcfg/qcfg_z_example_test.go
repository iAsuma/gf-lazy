package qcfg

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iasuma/gf-lazy/os/qenv"
)

func Example_loadConfig() {
	bootName := DefaultBootConfigName
	bootEnv := qenv.GetString(qenv.DefaultBootEnvName)
	if bootEnv != "" {
		bootName = bootName + "." + bootEnv
	}

	boot := GetString(context.TODO(), "nacos")
	bootMap := gconv.Map(boot)

	var ticket NacosTicket
	ticket.IpAddr = bootMap["ip_addr"].(string)
	ticket.Port = gconv.String(bootMap["port"])
	ticket.DataId = bootMap["data_id"].(string)
	ticket.Group = bootMap["group"].(string)
	ticket.NamespaceId = bootMap["namespace_id"].(string)
	LoadGoFrameNacosConfig(ticket, true)
}

func Example_simpleLoadConfig() {
	// qcfg.ConnectConfigCenter
	ConnectConfigCenter(context.TODO(), true)
}
