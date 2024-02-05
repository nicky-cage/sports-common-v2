package common

// IdName 通用IDName
type IdName struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// AdminToken 后台登录用户信息数据
type AdminToken struct {
	Id   uint64 `json:"id"`   //用户编号
	Name string `json:"name"` //用户名称
}

// IdNames 转换为序列
func IdNames(data map[uint8]string) []IdName {
	rows := []IdName{}
	for k, v := range data {
		rows = append(rows, IdName{
			Id:   uint64(k),
			Name: v,
		})
	}
	return rows
}
