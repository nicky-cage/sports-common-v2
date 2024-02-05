package ws

import (
	"encoding/json"
	"fmt"
	"sports-common/config"
	"sports-common/log"
	"sports-common/mapping"
	"strings"

	"gopkg.in/ffmt.v1"
	"gopkg.in/olahol/melody.v1"
)

var handleMessage = func(sess *melody.Session, data []byte) {

	var reqData Request
	if err := json.Unmarshal(data, &reqData); err != nil { //提取数据格式
		if err := sess.Write(ResponseErrorBytes("协议解析错误")); err != nil {
			fmt.Println("发送消息[协议解析错误]失败: ", err)
		}
		return
	}

	// 处理心跳
	if reqData.Code == Codes.HeartBeat { // 如果是心跳
		if err := sess.Write(ResponseHeartBeatBytes()); err != nil {
			fmt.Println("发送消息[处理心跳]失败: ", err)
		}
		return
	}

	host := sess.Request.Host                 // 当前发送消息主机名称
	platform := config.GetPlatformByURL(host) // 获取平台识别号, 如果是其他节点发布过来的消息, 则需要自带平台识别号
	if platform == "" {                       // 如果来访无法获知platform, 则取得传送过来的platform
		platform = reqData.Platform
	}
	if platform == "" {
		if err := sess.Write(ResponseErrorBytes("host = " + host + ", 无法获取相应的平台识别号")); err != nil {
			fmt.Println("发送消息[平台识别号错误]失败: ", err)
		}
		return
	}
	if reqData.Platform == "" {
		reqData.Platform = platform
	}
	//fmt.Println("host = ", host, ", platform = ", platform, ", reqData = ", reqData)

	// 如果是登录 - 则将相应的链接写到内存当中
	if reqData.Code == Codes.Login {
		strID, err := getConnectionID(platform, reqData.Data)
		if err != nil {
			if err := sess.Write(ResponseErrorBytes(err.Error())); err != nil {
				fmt.Println("发送消息[getConnectionID]失败: ", err)
			} // 错误 - 关闭链接;
			sess.Close()
			return
		}
		sessionID := platform + ":" + endPoint + ":" + strID
		sessionLock.Lock()
		defer sessionLock.Unlock()
		ok := false
		for _, v := range sessionList {
			if v == sessionID {
				ok = true
			}
		}
		if !ok {
			sessionList[sess] = sessionID
		}

		if err := sess.Write(ResponseLoginBytes()); err != nil {
			fmt.Println("发送消息[登录]失败: ", err)
		} // 下发登录成功消息
		return
	}

	// 向后端api广播消息
	if reqData.Code == Codes.Broadcast.Backend { // 向后端广播
		if endPoint != EndPointTypes.Backend {
			return
		}
		sentNodeCount := 0
		totalNodeCount := 0
		r := Response{}
		if err := mapping.MapToStruct(reqData.Data.(map[string]interface{}), &r); err != nil {
			log.Logger.Error("数据格式错误, 无法下发: ", reqData.Data)
			ffmt.Puts([]interface{}{
				"数据格式有误",
				reqData.Data,
			})
		}

		for s, k := range sessionList {
			sArr := strings.Split(k, ":")
			sessionPlatform := sArr[0]
			sessionEndpoint := sArr[1]
			if platform == sessionPlatform && sessionEndpoint == EndPointTypes.Backend { // 平台对应, 并且端点类型对应
				if err := s.Write(ResponseDataBytes(r.Data, r.Code)); err != nil {
					fmt.Println("发送消息[", r.Code, "]失败: ", err)
				}
				sentNodeCount++
			}
			totalNodeCount++
		}
		log.Logger.Println("共向", sentNodeCount, "个节点发送消息, 共计", totalNodeCount, "个节点")
		return
	}
	// 向前端api广播消息
	if reqData.Code == Codes.Broadcast.Frontend { // 向前端广播
		if endPoint != EndPointTypes.Frontend {
			return
		}
		sentNodeCount := 0
		totalNodeCount := 0
		r := Response{}
		if err := mapping.MapToStruct(reqData.Data.(map[string]interface{}), &r); err != nil {
			log.Logger.Error("数据格式错误, 无法下发: ", reqData.Data)
			return
		}
		for s, k := range sessionList {
			sArr := strings.Split(k, ":")
			sessionPlatform := sArr[0]
			sessionEndpoint := sArr[1]
			if platform == sessionPlatform && sessionEndpoint == EndPointTypes.Frontend {
				if err := s.Write(ResponseDataBytes(r.Data, r.Code)); err != nil {
					fmt.Println("发送消息[", platform, ",", r.Code, "]失败: ", err)
				}
				sentNodeCount++
			}
			totalNodeCount++
		}
		log.Logger.Println("共向", sentNodeCount, "个节点发送消息, 共计", totalNodeCount, "个节点")
		return
	}
	// 向后端指定用户推送消息
	if reqData.Code == Codes.Push.Backend {
		if endPoint != EndPointTypes.Backend {
			return
		}
		if checkBackendID == nil {
			log.Logger.Error("缺少获取指定用户的 backend - xxConnectID 函数")
			return
		}
		r := RequestDataLoginID{}
		if err := mapping.MapToStruct(reqData.Data.(map[string]interface{}), &r); err != nil {
			log.Logger.Error("数据格式错误, 无法下发: ", reqData.Data)
			return
		}
		r2 := Response{}
		if err := mapping.MapToStruct(r.Data.(map[string]interface{}), &r2); err != nil {
			log.Logger.Error("数据格式错误, 无法下发: ", r.Data)
			return
		}
		for s, k := range sessionList {
			sArr := strings.Split(k, ":")
			sessionPlatform := sArr[0]
			sessionEndpoint := sArr[1]
			if platform == sessionPlatform && sessionEndpoint == EndPointTypes.Backend && checkBackendID(platform, r.LoginID, k) {
				if err := s.Write(ResponseDataBytes(r2.Data, r2.Code)); err != nil {
					fmt.Println("发送消息[后端][", platform, ",", r2.Code, "]失败: ", err)
				}
			}
		}
		return
	}
	// 向前端指定用户推送消息
	if reqData.Code == Codes.Push.Frontend {
		if endPoint != EndPointTypes.Frontend {
			return
		}
		if checkFrontendID == nil {
			log.Logger.Error("缺少获取指定用户的 frontend - xxConnectID 函数")
			return
		}
		r := RequestDataLoginID{}
		if err := mapping.MapToStruct(reqData.Data.(map[string]interface{}), &r); err != nil {
			log.Logger.Error("数据格式错误, 无法下发: ", reqData.Data)
			return
		}
		r2 := Response{}
		if err := mapping.MapToStruct(r.Data.(map[string]interface{}), &r2); err != nil {
			log.Logger.Error("数据格式错误, 无法下发: ", r.Data)
			return
		}
		for s, k := range sessionList {
			sArr := strings.Split(k, ":")
			sessionPlatform := sArr[0]
			sessionEndpoint := sArr[1]
			if platform == sessionPlatform && sessionEndpoint == EndPointTypes.Backend && checkFrontendID(platform, r.LoginID, k) {
				if err := s.Write(ResponseDataBytes(r2.Data, r2.Code)); err != nil {
					fmt.Println("发送消息[后端][", platform, ",", r2.Code, "]失败: ", err)
				}
			}
		}
		return
	}

	// 检测处理函数是否绑定过了
	callback, exists := webSocketHandleFuncs[reqData.Code]
	if !exists {
		if err := sess.Write(ResponseErrorBytes("协议指令错误")); err != nil {
			fmt.Println("发送消息[缺少协议指令]失败: ", err)
		}
		return
	}

	result := callback(sess, reqData.Data)
	if reqData.Code == Codes.LiveChat { //直播返回下发的code
		reqData.Code = Codes.LiveChatResponse
	}
	if err := sess.Write(ResponseDataBytes(result, reqData.Code)); err != nil {
		fmt.Println("发送消息[内部函数调用]失败: ", err)
	}
}
