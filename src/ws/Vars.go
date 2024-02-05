package ws

import (
	"sync"

	"gopkg.in/olahol/melody.v1"
)

// heartBeatTimeout 心跳间隔时间
var heartBeatTimeout = 5

// Codes 常用WebScoket指令
var Codes = struct {
	HeartBeat         int // 心跳
	HeartBeatResponse int // 心跳下发
	Login             int // 登录
	LoginResponse     int // 登录下发
	LiveChat          int
	LiveChatResponse  int //直播下发
	Broadcast         struct {
		Backend  int
		Frontend int
	} // 广播 当前节点 -> 其他节点
	Push struct {
		Backend  int
		Frontend int
	} // 推送 当前节点 -> 其他节点
	Frontend struct {
		UserInfo int
	}
	Backend struct {
		Finance int // 充值下发
	}
}{
	HeartBeat:         100000, // 心跳
	HeartBeatResponse: 100001, // 心跳下发
	Login:             100010, // 登录 上传 >> {code: 100001, data: { token: ""}} 下发 >> { code: 200, message: "" }
	LoginResponse:     100011, // 登录下发
	LiveChat:          200201,
	LiveChatResponse:  200202,
	Broadcast: struct {
		Backend  int // 后端广播
		Frontend int // 前播广播
	}{
		Backend:  800010,
		Frontend: 600010,
	}, // 广播
	Push: struct {
		Backend  int // 后端推送
		Frontend int // 前端推送
	}{
		Backend:  800020,
		Frontend: 600020,
	}, // 推送
	Frontend: struct {
		UserInfo int
	}{
		UserInfo: 600100,
	},
	Backend: struct {
		Finance int
	}{
		Finance: 800110,
	},
}

var sessionList map[*melody.Session]string                    // webSocketSessions 所有链接
var sessionLock = new(sync.Mutex)                             // 锁
var getConnectionID func(string, interface{}) (string, error) // 获取链接id的函数

// 函数处理代码 => 处理函数
// { 100101: controllers.Index.Test, }
var webSocketHandleFuncs map[int]func(*melody.Session, interface{}) interface{}

// websocket 对象
var currentNode *melody.Melody

// inboxFrontend 信箱
var inboxList map[string]*chan Request = map[string]*chan Request{}

// endPoint 后端/前端 -> front/back
var endPoint string = "front"

// EndPointTypes 类型
var EndPointTypes = struct {
	Backend  string
	Frontend string
}{
	Backend:  "backend",
	Frontend: "frontend",
}

// 检测前端session-id
var checkFrontendID func(string, string, string) bool = nil

// 检测后端session-id
var checkBackendID func(string, string, string) bool = nil
