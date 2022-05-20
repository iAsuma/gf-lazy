package qcfg

// LoadGoFrameNacosConfig nacos引入goframe支持
func LoadGoFrameNacosConfig(ticket NacosTicket, isListen bool, filename ...string)  {
	if ticket.CacheDir == "" {
		ticket.CacheDir = GetRuntimeDir()
	}

	if ticket.LogDir == "" {
		ticket.LogDir = GetLogDir()
	}

	if ticket.LogLevel == "" {
		ticket.LogLevel = "error"
	}

	client := NacosClient(ticket)
	content := client.GetConfig()

	configFile := GetConfigFileName(filename...)
	SetGoFrameConfig(configFile, content)

	// 是否监听配置文件变化
	if isListen {
		_ = client.ListenConfig(func(data string) {
			SetGoFrameConfig(configFile, data)
		})
	}
}

// LoadBeegoNacosConfig nacos引入beego支持
func LoadBeegoNacosConfig()  {
	//TODO
}