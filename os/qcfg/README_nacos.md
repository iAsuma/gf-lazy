# GoFrame 加载`Nacos`配置

## 加载`Nacos`配置
``` go
var ticket qcfg.NacosTicket
ticket.IpAddr = "your nacos Ip Addr"
ticket.Port = "your nacos Port"
ticket.DataId = "your nacos Data Id"
ticket.Group = "your nacos Group"
ticket.NamespaceId = "your nacos Namespace Id"

qcfg.LoadGoFrameNacosConfig(ticket, true) // 加载goframe框架nacos配置
```

## 快捷方法
``` go
qcfg.ConnectConfigCenter(ctx context.Context, isListen bool)
```

## 示例

 连接配置中心（Nacos）示例：

1. 在你的交付清单文件夹`manifest/config/`下应该有`boot.yaml`文件，这是连接nacos的配置文件
```
nacos:
   ip_addr: "{REPALCE_IP_ADDR}"
   port: "{REPALCE_PORT}"
   data_id: "{REPALCE_DATA_ID}"
   namespace_id: "{REPALCE_NAMESPACE_ID}"
   group: "{REPALCE_GROUP}"
```
2. 复制`boot.yaml`文件到项目根目录`config`下

```bash
 cp manifest/config/boot.yaml config/
```
2. 修改boot.yaml

```bash
 vi boot.yaml
 nacos:
   ip_addr: "127.0.0.1"
   port: 8848
   data_id: "your_data_id"
   namespace_id: "your_namespace_id"
   group: "your group, likes DEFAULT_GROUP"
```

3. 入口添加启动

```go
 package main
 import(
   "context"
   "github.com/iasuma/gf-lazy/os/qcfg"
 ) 
 
 func main (){
     // 加载配置中心配置
     qcfg.ConnectConfigCenter(context.TODO(), true)
     
     // 项目启动run
 }
 ```

 4. 启动
```bash
go run main.go
```
--- 系统自动加载`boot.yaml`中连接的配置
启动后会在`config/`下生成`config.yaml`，同时会删除根目录下可能存在的config

 5. 如果你的启动配置不是`boot.yaml`请自行实现获取启动配置的方法
示例：
``` go
bootName := "start" // start.yaml
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
```

 使用其他启动配置：

```bash
 cp manifest/config/boot.yaml config/boot.your.yml
 # 1. 配置系统环境变量 BOOT_ENV=your
 #	 go run main.go 会自动启动`boot.your.yml`中连接的配置
 # 2. 或者直接执行 go run main.go --boot.env=your
 # 3. 或者用IDE 配置当前项目的启动环境变量 BOOT_ENV=your
```

 --  `qcfg.ConnectConfigCenter`读取环境变量配置的优先级顺序取决于goframe的`genv.GetWithCmd`的读取顺序

**注：`BOOT_ENV` 不是应用运行的环境配置，只是用来决定启动读取的连接配置**
