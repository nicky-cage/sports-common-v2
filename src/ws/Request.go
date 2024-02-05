package ws

import "encoding/json"

// Request 请求数据类型 - 默认
type Request struct {
	Code     int         `json:"code"`     //请求代码
	Platform string      `json:"platform"` // 平台识别号码
	Data     interface{} `json:"data"`     //请求信息
}

// RequestLoginByToken 依据token登录
type RequestLoginByToken struct {
	Token string `json:"token"` // Token
}

// RequestDataLoginID 依据login-id
type RequestDataLoginID struct {
	LoginID string      `json:"login_id"` // Token
	Data    interface{} `json:"data"`
}

// RequestLoginByAdmin 依据管理员编号登录
type RequestLoginByAdmin struct {
	AdminID string `json:"admin_id"` // Token
}

// RequestLoginByUser 依据用户编号登录
type RequestLoginByUser struct {
	UserID string `json:"user_id"` // Token
}

// HeartBeatJSONBytes 心跳包字节码
// var HeartBeatJSONBytes []byte

// HeartBeatJSONBytes
// HeartBeatJSON 生成心跳包
// var HeartBeatJSON = func() []byte {
var HeartBeatJSONBytes = func() []byte {
	//if HeartBeatJSONBytes != nil {
	//	return HeartBeatJSONBytes
	//}
	req := Request{
		Code: Codes.HeartBeat,
		Data: "",
	}
	//HeartBeatJSONBytes, _ = json.Marshal(&req)
	//return HeartBeatJSONBytes

	bytes, _ := json.Marshal(&req)
	return bytes
}()

// RequestHeartBeat 心跳包信息
var RequestHeartBeat = Request{
	Code: Codes.HeartBeat,
	Data: map[string]interface{}{},
}
