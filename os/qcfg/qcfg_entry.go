package qcfg

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iasuma/gf-lazy/os/qenv"
)

// GetNacosTicket 获取配置中nacos的连接信息
func GetNacosTicket(ctx context.Context) NacosTicket {
	bootName := DefaultBootConfigName
	bootEnv := qenv.GetString(qenv.DefaultBootEnvName)
	if bootEnv != "" {
		bootName = bootName + "." + bootEnv
	}

	boot := Instance(bootName).GetString(ctx, "nacos")
	bootMap := gconv.Map(boot)

	var ticket NacosTicket
	ticket.IpAddr = gconv.String(bootMap["ip_addr"])
	ticket.Port = gconv.String(bootMap["port"])
	ticket.DataId = gconv.String(bootMap["data_id"])
	ticket.Group = gconv.String(bootMap["group"])
	ticket.NamespaceId = gconv.String(bootMap["namespace_id"])
	ticket.LogLevel = gconv.String(bootMap["log_level"])

	return ticket
}

// ConnectConfigCenter 连接配置中心
func ConnectConfigCenter(ctx context.Context, isListen bool)  {
	ticket := GetNacosTicket(ctx)
	LoadGoFrameNacosConfig(ticket, isListen)
}