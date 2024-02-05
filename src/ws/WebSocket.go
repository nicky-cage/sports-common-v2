package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sports-common/config"
	"sports-common/log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/z9905080/gowebsocket"
	"gopkg.in/olahol/melody.v1"
)

// Websocket
var WebSocket = struct {
	// 参数一: 登录时提交的数据
	// 返回一: 保存在本节点, 用做唯一识别的key
	SetConnectID func(func(string, interface{}) (string, error))              // 获取链接ID
	Start        func()                                                       // 初始化
	Handle       func(*gin.Context)                                           // 处理函数
	AddRouters   func(map[int]func(*melody.Session, interface{}) interface{}) // 增加相关路由
	InitNodes    func()                                                       // *** 此方法必须要在其他方法之后执行 ***
	Backend      struct {                                                     // 向后端发送消息
		// 参数一: 要广播的消息
		Broadcast func(string, Response) // 广播给到所有用户
		// 参数一: 要发送消息给的用户编号
		// 参数二: 要发送给用户的消息
		Push func(string, string, Response) // 推送给单个用户
		// 参数一: 要发送消息给的用户编号
		// 参数二: 存储在websocket-node上的用户id(token或者用户编号, 由用户自定义)
		SetCheckID func(func(string, string, string) bool) // 用于检测是否是当前用户(要推送消息的用户)
	}
	Frontend struct { // 向前端发送消息
		// 参数一: 要广播的消息
		Broadcast func(string, Response) // 广播给到所有用户
		// 参数一: 要发送消息给的用户编号
		// 参数二: 要发送给用户的消息
		Push func(string, string, Response) // 推送给单个用户
		// 参数一: 要发送消息给的用户编号
		// 参数二: 存储在websocket-node上的用户id(token或者用户编号, 由用户自定义)
		SetCheckID func(func(string, string, string) bool) // 用于检测是否是当前用户(要推送消息的用户)
	}
	SetEndPoint func(string)
}{
	Start: func() {
		currentNode = melody.New()
		currentNode.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		currentNode.Config.MaxMessageSize = 1024
		currentNode.Config.WriteWait = time.Second * 10
		currentNode.Config.PongWait = 60 * time.Second
		currentNode.Config.PingPeriod = 50 * time.Second
		currentNode.Config.MessageBufferSize = 256
		sessionList = map[*melody.Session]string{}
		currentNode.HandleDisconnect(func(s *melody.Session) { // 处理断开链接
			sessionLock.Lock()         // 共享 - 加锁
			defer sessionLock.Unlock() // 共享 - 释放
			fmt.Println("---------------------------- 连接中断: ", s.Request.RemoteAddr, " --------------------------")
			delete(sessionList, s) // 删除关闭的链接
		})
		currentNode.HandleConnect(func(sess *melody.Session) {
			addr := sess.Request.RemoteAddr
			fmt.Println("*** 建立连接... FROM: ", addr, " ***")
		})
		currentNode.HandleMessage(handleMessage) // 处理信息响应
	},
	SetConnectID: func(callback func(string, interface{}) (string, error)) {
		getConnectionID = callback
	},
	Handle: func(c *gin.Context) { // 必须先执行Start())方法
		if currentNode == nil {
			log.Logger.Error("缺少 melody websocket 服务器")
			return
		}
		err := currentNode.HandleRequest(c.Writer, c.Request)
		if err != nil {
			log.Logger.Error("WS: 连接错误: ", err)
		}
	},
	AddRouters: func(routers map[int]func(sess *melody.Session, data interface{}) interface{}) {
		webSocketHandleFuncs = routers //设置websocket路由
	},
	Backend: struct {
		Broadcast  func(string, Response)
		Push       func(string, string, Response)
		SetCheckID func(func(string, string, string) bool)
	}{
		Broadcast: func(platform string, data Response) {
			if endPoint == EndPointTypes.Backend { // 如果当前节点就是后端节点, 则先往本节点下所有连接发消息
				for s := range sessionList {
					s.Write(ResponseDataBytes(data.Data, data.Code))
				}
			}
			req := Request{
				Code:     Codes.Broadcast.Backend,
				Platform: platform,
				Data:     data,
			}
			for nodeName, v := range inboxList { // 给每个信箱发消息
				fmt.Println("广播各个节点, 当前节点 = ", nodeName)
				*v <- req
			}
		},
		Push: func(platform, loginID string, data Response) {
			if endPoint == EndPointTypes.Backend { // 先向当前节点满足条件的地方发送消息
				for s, k := range sessionList {
					if k == loginID {
						s.Write(ResponseDataBytes(data.Data, data.Code))
					}
				}
			}
			req := Request{
				Code:     Codes.Broadcast.Backend,
				Platform: platform,
				Data: RequestDataLoginID{
					LoginID: loginID,
					Data:    data,
				},
			}
			for _, v := range inboxList { // 给每个信箱发消息
				*v <- req
			}
		},
		SetCheckID: func(callback func(string, string, string) bool) {
			checkBackendID = callback
		},
	},
	Frontend: struct {
		Broadcast  func(string, Response)
		Push       func(string, string, Response)
		SetCheckID func(func(string, string, string) bool)
	}{
		Broadcast: func(platform string, data Response) {
			//if endPoint == EndPointTypes.Frontend { // 如果当前节点是前端节点, 先往本节点发送消息
			//	for s := range sessionList {
			//		s.Write(ResponseDataBytes(data.Data, data.Code))
			//	}
			//}
			req := Request{
				Code:     Codes.Broadcast.Frontend,
				Platform: platform,
				Data:     data,
			}
			for _, v := range inboxList {
				*v <- req
			}
		},
		Push: func(platform, loginID string, data Response) {
			if endPoint == EndPointTypes.Frontend { // 先向当前节点发送消息
				for s, k := range sessionList {
					platformIndex := strings.Index(k, ":")
					realKey := k[platformIndex+1:]
					if realKey == loginID {
						s.Write(ResponseDataBytes(data.Data, data.Code))
						return
					}
				}
			}
			req := Request{
				Code:     Codes.Broadcast.Frontend,
				Platform: platform,
				Data: RequestDataLoginID{
					LoginID: loginID,
					Data:    data,
				},
			}
			for _, v := range inboxList { // 向每一个节点信息发送消息
				*v <- req
			}
		},
		SetCheckID: func(callback func(string, string, string) bool) {
			checkFrontendID = callback
		},
	},
	InitNodes: func() {
		nodeListStr := config.Get("websocket.node_list") // 以,号隔开
		nodeListArray := strings.Split(nodeListStr, ",")
		for _, v := range nodeListArray {
			nodeStr := strings.TrimSpace(v)
			if nodeStr == "" { // 忽略空节点内容
				continue
			}
			inbox := make(chan Request, 100) // 收件箱列表
			inboxList[nodeStr] = &inbox
			go func() {
				var socket gowebsocket.Socket
				reConnect := func(args ...[]byte) {
					for {
						socket = gowebsocket.New(nodeStr)
						socket.Connect()
						// fmt.Println("正在建立连接至服务器: ", nodeStr, " ===> ", socket.IsConnected)
						if !socket.IsConnected {
							// fmt.Println("建立到服务器 ", nodeStr, " 的连接失败, 5s 之后重试 ...")
							time.Sleep(time.Second * 5)
							continue
						}
						if len(args) >= 1 {
							bytes := args[0]
							err := socket.Conn.WriteMessage(websocket.TextMessage, bytes) // 发送消息
							if err != nil {
								fmt.Println("发送消息错误: ", err, ", 当前 status = ")
								continue
							}
						}
						break
					}
				}
				sendMessage := func(data interface{}) {
					bytes, err := json.Marshal(&data)
					if err != nil {
						fmt.Println("将数据转换为 bytes 出现错误:", err, ", data = %v", data)
						return
					}
					if !socket.IsConnected {
						fmt.Println("还没有建立连接, 或者连接已断开, 重新建立连接至服务器: ", nodeStr)
						reConnect(bytes)
						return
					}
					err = socket.Conn.WriteMessage(websocket.TextMessage, bytes) // 发送消息
					if err != nil {
						fmt.Println("发送消息出现错误 -> writeMessage: ", err)
						socket.Close()
						reConnect(bytes)
					}
				}
				reConnect() // 如果链接关闭, 则自动重新链接
				counter := 1
				for {
					select {
					case message := <-inbox: // 信箱消息 - 本节点
						sendMessage(message)
					default: // 处理完信箱消息之后, 再处理心跳包
						if counter%heartBeatTimeout == 0 {
							sendMessage(RequestHeartBeat)
						}
						counter++
						time.Sleep(time.Second)
					}
				}
			}()
		}
	},
	SetEndPoint: func(name string) { // back/front
		endPoint = name
	},
}
