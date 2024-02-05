package ws

import "encoding/json"

// websocket 返回数据类型
type Response struct {
	Code    int         `json:"code"`    //错误代码
	Message string      `json:"message"` //错误信息
	Data    interface{} `json:"data"`    //返回相关结果
}

// ResponseFailure  默认失败信息
var ResponseFailure = Response{
	Code:    500,
	Message: "",
}

// ResponseErrorBytes 默认错误信息
var ResponseErrorBytes = func(message string) []byte {
	result := Response{
		Code:    500,
		Message: message,
	}
	bytes, _ := json.Marshal(&result)
	return bytes
}

// ResponseMessageBytes 默认成功信息
var ResponseMessageBytes = func(message string) []byte {
	result := Response{
		Code:    200,
		Message: message,
	}
	bytes, _ := json.Marshal(&result)
	return bytes
}

// ResponseHeartBeatBytes 心跳
var ResponseHeartBeatBytes = func() []byte {
	result := Response{
		Code:    Codes.HeartBeatResponse,
		Message: "",
	}
	bytes, _ := json.Marshal(&result)
	return bytes
}

// ResponseLoginBytes 登录
var ResponseLoginBytes = func() []byte {
	result := Response{
		Code:    Codes.LoginResponse,
		Message: "",
	}
	bytes, _ := json.Marshal(&result)
	return bytes
}

// ResponseDataBytes 默认成功信息
var ResponseDataBytes = func(data interface{}, args ...int) []byte {
	code := 200
	if len(args) >= 1 {
		code = args[0]
	}
	result := Response{
		Code: code,
		Data: data,
	}
	bytes, _ := json.Marshal(&result)
	return bytes
}
