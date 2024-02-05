package consts

// 菜单等级
var MenuLevels = map[uint8]string{
	1: "主菜单",
	2: "子菜单",
	3: "右侧标签",
	4: "动作/路由",
	5: "弹出窗口/路由",
	6: "弹窗二级菜单",
}

var UserLogModules = map[int]string{
	0: "常规访问",
	1: "系统操作",
	2: "用户中心",
	3: "财务管理",
	4: "支付",
	5: "取款",
	6: "邀请好友",
	7: "游戏",
}

// 前台模块
var UserModules = struct {
	System   int // 系统操作
	User     int // 用户中心
	Finance  int // 财务管理
	Payment  int // 支付
	Withdraw int // 取款
	Invite   int // 邀请好友
	Game     int
}{
	System:   1,
	User:     2,
	Finance:  3,
	Payment:  4,
	Withdraw: 5,
	Invite:   6,
	Game:     7,
}

var UserLogTypes = map[int]string{
	0: "访问",
	1: "添加",
	2: "修改",
	3: "删除",
	4: "读取",
}

// 操作类型
var UserOperationTypes = struct {
	Create int
	Update int
	Delete int
	Read   int
}{
	Create: 1,
	Update: 2,
	Delete: 3,
	Read:   4,
}

// 操作风险级别
var UserOperationLevels = struct {
	High   int
	Middle int
	Low    int
}{
	High:   1,
	Middle: 2,
	Low:    3,
}
