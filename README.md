# integrated-common

综合 - 公用

## 一、配置文件
请将本目录下 setting.ini 拷贝至你的具体项目生成的bin同目录下面

## 二、Log 使用
在 config.LoadConfigs() 之后调用: 
```go
config.LoadConfigs()
log.Start() //启动日志程序
// log.Logger 即为 *logrus.Logger 类型, 可以直接调用
log.Logger.WithFields(map[string]interface{}{ 
    "key": "KEY",
}).Info("---------------")
// 或者直接调用
log.Logger.Info("hello ................")
```

## 三、配置使用
在 config.LoadConfigs() 之后调用:
```go
config.LoadConfigs()
// 假如配置节点如下: 
// [database]
// user = "root"
// ...
// 则获取方式如下: 
dbUser := config.Get("database.user") //返回为string类型
```

### 四、各游戏场馆端口和服务

integrated-game-api 虚拟host: integrated-game-api.prod port 8802


integrated-member-api 虚拟host: integrated-member-api.prod port 8801


integrated-admin 虚拟host: integrated-admin.prod port 8818