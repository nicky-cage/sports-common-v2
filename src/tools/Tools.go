package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"runtime"
	"sports-common/consts"
	"sports-common/log"
	"strconv"
	"strings"
	"time"

	"gopkg.in/ffmt.v1"
)

// 加密字符-KEY
var mcryptLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GoID 协程编号
// 注意: 以获取调用栈的方式获取id, 性能堪忧, 请尽量不要使用此方法
func GoID() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// StructureToMapNew 结构转map,因为struct不能循环,map可以.
func StructureToMapNew(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		log.Logger.Info(tag)
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = StructureToMapNew(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

// StructureToMap 结构转map,因为struct不能循环,map可以.
func StructureToMap(request interface{}, tagInfo string) map[string]interface{} {
	result := make(map[string]interface{})
	refType := reflect.TypeOf(request)
	refValue := reflect.ValueOf(request)
	for i := 0; i < refType.NumField(); i++ {
		tag := refType.Field(i).Tag.Get(tagInfo)
		if tag == "" || tag == "-" {
			continue
		}
		val := ""
		switch refValue.Field(i).Interface().(type) {
		case string:
			_tmp := refValue.Field(i).Interface().(string)
			val = fmt.Sprint(_tmp)
		case int:
			_tmp := refValue.Field(i).Interface().(int)
			val = fmt.Sprint(_tmp)
		case uint:
			_tmp := refValue.Field(i).Interface().(int)
			val = fmt.Sprint(_tmp)
		case int32:
			_tmp := refValue.Field(i).Interface().(int32)
			val = fmt.Sprint(_tmp)
		case uint32:
			_tmp := refValue.Field(i).Interface().(uint32)
			val = fmt.Sprint(_tmp)
		case int64:
			_tmp := refValue.Field(i).Interface().(int64)
			val = fmt.Sprint(_tmp)
		case uint64:
			_tmp := refValue.Field(i).Interface().(uint64)
			val = fmt.Sprint(_tmp)
		case float32:
			_tmp := refValue.Field(i).Interface().(float32)
			val = fmt.Sprint(_tmp)
		case float64:
			_tmp := refValue.Field(i).Interface().(float64)
			val = fmt.Sprint(_tmp)
		}

		result[tag] = val
	}
	return result
}

// StructureToSlice 结构转map,因为struct不能循环,map可以.
func StructureToSlice(request interface{}, tagInfo string) []string {
	var result []string
	refType := reflect.TypeOf(request)
	refValue := reflect.ValueOf(request)
	for i := 0; i < refType.NumField(); i++ {
		tag := refType.Field(i).Tag.Get(tagInfo)
		if tag == "" {
			continue
		}
		result = append(result, tag)
		val := ""
		switch refValue.Field(i).Interface().(type) {
		case string:
			_tmp := refValue.Field(i).Interface().(string)
			val = fmt.Sprint(_tmp)
		case int:
			_tmp := refValue.Field(i).Interface().(int)
			val = fmt.Sprint(_tmp)
		case uint:
			_tmp := refValue.Field(i).Interface().(int)
			val = fmt.Sprint(_tmp)
		case int32:
			_tmp := refValue.Field(i).Interface().(int32)
			val = fmt.Sprint(_tmp)
		case uint32:
			_tmp := refValue.Field(i).Interface().(uint32)
			val = fmt.Sprint(_tmp)
		case int64:
			_tmp := refValue.Field(i).Interface().(int64)
			val = fmt.Sprint(_tmp)
		case uint64:
			_tmp := refValue.Field(i).Interface().(uint64)
			val = fmt.Sprint(_tmp)
		case float32:
			_tmp := refValue.Field(i).Interface().(float32)
			val = fmt.Sprint(_tmp)
		case float64:
			_tmp := refValue.Field(i).Interface().(float64)
			val = fmt.Sprint(_tmp)
		}

		result = append(result, val)
	}
	return result
}

// GenVCode 验证码
func GenVCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

// SendSMS 发送验证码短信
func SendSMS(phone string, code string, tmp int, channel string) error {
	if channel == "1" { // 首选短信通道
		return MessageChannel(phone, code, tmp)
	} else if channel == "2" { // 备用短信通道  未对接
		return MessageChannelSpare(phone, code, tmp)
	} else if channel == "3" { // 次用短信通道  未对接
		return MessageChannelSecond(phone, code, tmp)
	}
	return errors.New("暂无短信通道")
}

// 首选短信通道
func MessageChannel(phone string, code string, tmp int) error {
	/**
	253云通讯】 必须是备案的
	账号：YZM5612013
	状态： 正常
	密码：CkeXzI6DUJf9d8
	接口：http://smssh1.253.com/msg/send/json
	{
	"code":"0",
	"msgId":"17041010383624511",
	"time":"20170410103836",
	"errorMsg":""
	}
	*/
	//templateList := []string{
	//	"【天际】 您的验证码为：" + code + "，非本人操作请忽略本信息。",
	//	"【天际】" + code + " 号，密码输错3次，账户被锁定！",
	//	"【天际】" + code + " 号，输入应急码，账户被锁定！",
	//}
	templateList := []string{
		"【威尼斯】 您的验证码为：" + code + "，非本人操作请忽略本信息。",
		"【威尼斯】" + code + " 号，密码输错3次，账户被锁定！",
		"【威尼斯】" + code + " 号，输入应急码，账户被锁定！",
	}
	if len(templateList) <= tmp {
		return errors.New("模板不存在")
	}
	msg := templateList[tmp]

	message := url.QueryEscape(msg)
	var params string

	params = "&UserId=tj88888&UserPwd=qaztj1238!@A&SendPhone=" + phone + "&SendMessage=" + message + ""
	msgURL := "http://api.eait.cn/Sms/SmsHttp_U.aspx?Action=SendMessage" + params //"http://smssh1.253.com/msg/send/json"

	resp, err := HttpGet(msgURL)
	if err != nil {
		ffmt.Println(err.Error())
		return errors.New("网络错误")
	}

	if resp == "" {
		return errors.New("发送失败")
	}
	if resp[0:1] == "1" {
		return nil
	}
	log.Logger.Errorf("发送信息失败 %s", resp)
	return errors.New("发送失败")
}

// 备用短信通道
func MessageChannelSpare(phone string, code string, tmp int) error {
	return errors.New("暂无对接备用短信通道")
}

// 次用短信通道
func MessageChannelSecond(phone string, code string, tmp int) error {
	return errors.New("暂无对接次用短信通道")
}

// SendEmail 发送邮件
func SendEmail(email, title, content string) error {
	accesskey := "5ca830e287b65f6374f3e1f3"
	secretkey := "eddf6f346eeb4ece92f7b7cf3ab99583"
	randInt := RandInt64(100000, 999999)
	randStr := strconv.Itoa(int(randInt))
	url := "https://live.kewail.com/directmail/v1/singleSendMail?accesskey=" +
		accesskey + "&random=" + randStr
	fromEmail := "mail@service.mnbkop.com"
	timeStamp := time.Now().Unix()
	timeStampStr := strconv.Itoa(int(timeStamp))

	data := "secretkey=" + secretkey + "&random=" + randStr + "&time=" +
		timeStampStr + "&fromEmail=" + fromEmail

	postMap := map[string]interface{}{
		"ext":         "",
		"replyEmail":  fromEmail,
		"fromAlias":   "验证码服务",
		"htmlBody":    content,
		"needToReply": true,
		"subject":     title,
		"clickTrace":  "0",
		"time":        timeStamp,
		"type":        0,
		"toEmail":     email,
		"fromEmail":   fromEmail,
		"sig":         SHA256(data),
	}
	postStr, _ := json.Marshal(postMap)
	resp, err := HttpPost(url, postStr, 0)
	if err != nil {
		return err
	}

	var result = struct {
		Result int    `json:"result"`
		ErrMsg string `json:"errmsg"`
	}{}
	_ = json.Unmarshal([]byte(resp), &result)
	if result.Result == 0 && result.ErrMsg == "OK" {
		return nil
	}

	return errors.New("发送失败")
}

// GetFieldByStruct 获取结构体字段
func GetFieldByStruct(v interface{}, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

// GetBillNo 生成订单号 字符串型
func GetBillNo(prefix string, n int) string {
	if n == 0 {
		n = 5
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = mcryptLetters[rand.Intn(len(mcryptLetters))]
	}
	randStr := string(b)
	return prefix + (time.Now().Format(consts.TimeBillLayoutYmd)) + randStr
}

// GetBillNoInt 生成订单号数字
func GetBillNoInt(prefix string, n int) string {
	if n == 0 {
		n = 5
	}
	min, max := 0, 0
	if n == 5 {
		min = 0
		max = 99999
	}

	if min >= max {
		return "0"
	}
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return prefix + (time.Now().Format(consts.TimeBillLayoutYmd)) + strconv.Itoa(randNum)
}

// EncryptedPassword 生成用户密码
func EncryptedPassword(pwd string) string {
	pwd = strings.ToLower(pwd)
	pwd = MD5(pwd)
	return strings.ToLower(pwd)
}

// VerifyPassword 校验密码
func VerifyPassword(pwd, enPwd string) bool {
	pwd = strings.ToLower(pwd)
	pwd = MD5(pwd)
	pwd = strings.ToLower(pwd)
	return pwd == enPwd
}

// SetStructFieldByJsonName 将hgetall出来的map[string]string转换成对应的结构体
func SetStructFieldByJsonName(ptr interface{}, fields map[string]string) {

	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("json")

		//去掉逗号后面内容 如 `json:"voucher_usage,omitempty"`
		name = strings.Split(name, ",")[0]

		if value, ok := fields[name]; ok {

			//给结构体赋值
			//保证赋值时数据类型一致
			//fmt.Println("类型1：", reflect.ValueOf(value).Type(), "类型2：", v.FieldByName(fieldInfo.Name).Type())
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "uint64" {
				iv, _ := strconv.Atoi(value)
				nv := uint64(iv)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(nv))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "float64" {
				fv, _ := strconv.ParseFloat(value, 64)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(fv))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "int32" {
				iv, _ := strconv.Atoi(value)
				nv := int32(iv)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(nv))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "uint32" {
				iv, _ := strconv.Atoi(value)
				nv := uint32(iv)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(nv))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "int" {
				iv, _ := strconv.Atoi(value)
				nv := iv
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(nv))
			}

		}
	}

}

// EncodeUserId 编码用户编号
func EncodeUserId(uid uint64) string {
	uid = uid + consts.IncrUserId
	acode := Dec2Hex(int(uid))
	return acode
}

// DecodeUserId 解码用户编号
func DecodeUserId(icode string) (int, error) {
	code, err := Hex2Dec(icode)
	if err != nil {
		return 0, errors.New("邀请编码有误")
	}

	code = code & 0x0FFFFFFF
	code = code - consts.IncrUserId
	return code, nil
}
