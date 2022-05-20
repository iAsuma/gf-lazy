# GoFrame 懒人包
gf-lazy 是专为GoFrame v2的懒人包，简化部分方法在业务代码中的`nil`判断，忽略部分非必要`error`参数接收，当然她最大的特色是接入`nacos`的配置

> 特色：加载`nacos`配置

##### 使用
```
go get -u github.com/iasuma/gf-lazy@{tag}
```

##### Why 'qxxxx'?
Just My Name Contains 'q'!

## 配置包 `qcfg`

### 加载`Nacos`配置

使用方法详见[Nacos配置说明](os/qcfg/README_nacos.md)

### 应用环境判断
``` go
qcfg.EnvIsProd(env ...interface{}) //判断是否生产环境
qcfg.EnvIsXXXX(env ...interface{})
```

## 环境包 qenv

读取系统环境（会从命令行参数中查找），`genv.GetWithCmd()`的实现

```
envVar := qenv.GetString(key string, def ...interface{}) //读取系统环境
```

`qcfg`环境判断的底层实现

``` go
qenv.IsProd(envName interface{})  //是否生成环境
qenv.IsXXXXX(envName interface{})
```

