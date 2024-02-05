package common

import (
	"sports-common/caches"
	"sports-common/log"
	"sports-common/request"
	"sports-common/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"xorm.io/builder"
)

// Backend 后台控制器基类
type Backend struct {
	Model        IModel                                            //模型
	Headers      []map[string]string                               //字段映射
	OptionAll    bool                                              //是否有all方法
	OptionList   bool                                              //列表
	OptionCreate bool                                              //添加
	OptionUpdate bool                                              //修改
	OptionDelete bool                                              //删除
	OptionDetail bool                                              //查看
	OptionOther  bool                                              //没有编辑和删除时,有其他操作按钮,需要显示操作栏
	OptionUtils  []string                                          //其他可用操作选项
	Platform     string                                            //平台识别号
	QueryCond    map[string]interface{}                            //查询条件
	ProcessRow   func(interface{})                                 //默认处理函数
	Rows         func() interface{}                                //多条记录
	RowsProcess  bool                                              //是否处理
	Validator    func(map[string]interface{}) error                //校验器
	Row          func() interface{}                                //单条记录
	RowMg        func() interface{}                                //单条记录-for-mg
	OrderBy      func(*gin.Context) string                         //获取排序
	CreateBefore func(*gin.Context, *map[string]interface{}) error //添加数据之前处理, 可以中断
	CreateAfter  func(*gin.Context, *map[string]interface{})       //添加数据之后处理, 不可中断
	UpdateBefore func(*gin.Context, *map[string]interface{}) error //修改数据之前处理, 可以中断
	UpdateAfter  func(*gin.Context, *map[string]interface{})       //修改数据之后处理, 不可中断
	SaveBefore   func(*gin.Context, *map[string]interface{}) error //保存数据之前处理, 可以中断
	SaveAfter    func(*gin.Context, *map[string]interface{})       //保存数据之后处理, 不可中断
	DeleteBefore func(*gin.Context, interface{}) error             //删除之前处理, 可以中断
	DeleteAfter  func(*gin.Context, interface{})                   //删除之后处理, 不可中断
}

// List 列表
func (ths *Backend) List(c *gin.Context) {
	//if !ths.OptionList {
	//	response.Err(c, "缺少权限")
	//	return
	//}
	rows := ths.Rows()
	limit, offset := request.GetOffsets(c)
	platform := ths.getPlatform(c)
	request.GetJwtToken(c)

	var total uint64
	var err error
	if ths.OrderBy != nil {
		total, err = ths.Model.FindAll(platform, rows, request.GetQueryCond(c, ths.QueryCond), limit, offset, ths.OrderBy(c))
	} else {
		total, err = ths.Model.FindAll(platform, rows, request.GetQueryCond(c, ths.QueryCond), limit, offset)
	}
	if ths.ProcessRow != nil {
		ths.ProcessRow(rows)
	}
	if err != nil {
		log.Err("获取列表信息出错: %v\n", err)
		response.Err(c, "获取列表错误")
		return
	}
	response.Result(c, struct {
		Headers []map[string]string `json:"headers"`
		Options map[string]bool     `json:"options"`
		List    interface{}         `json:"list"`
		Total   uint64              `json:"total"`
	}{
		Headers: ths.Headers,
		Options: map[string]bool{
			"create": ths.OptionCreate,
			"update": ths.OptionUpdate,
			"delete": ths.OptionDelete,
			"detail": ths.OptionDetail,
			"other":  ths.OptionOther,
		},
		List:  rows,
		Total: total,
	})
}

// Detail 详情
func (ths *Backend) Detail(c *gin.Context) {
	//if !ths.OptionDetail {
	//	response.Err(c, "缺少权限")
	//	return
	//}
	idStr, exists := c.GetQuery("id")
	if !exists || idStr == "" {
		log.Err("无法获取id信息!\n")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		log.Err("ID转换失败: %v\n", err)
		return
	}

	row := ths.Row()
	cond := builder.NewCond().And(builder.Eq{"id": id})
	platform := ths.getPlatform(c)
	exists, err = ths.Model.Find(platform, row, cond)
	if !exists || err != nil {
		log.Err("获取详情失败: %v\n", err)
		return
	}
	response.Result(c, row)
}

// Create 添加
func (ths *Backend) Create(c *gin.Context) {
	if !ths.OptionCreate {
		response.Err(c, "缺少权限")
		return
	}
	postedData := request.GetPostedData(c)
	if ths.CreateBefore != nil {
		if err := ths.CreateBefore(c, &postedData); err != nil { //可以取消
			response.Err(c, err.Error())
			return
		}
	}
	if ths.SaveBefore != nil {
		if err := ths.SaveBefore(c, &postedData); err != nil {
			response.Err(c, err.Error())
			return
		}
	}
	if ths.Validator != nil {
		if err := ths.Validator(postedData); err != nil {
			response.Err(c, err.Error())
			return
		}
	}
	platform := ths.getPlatform(c)
	id, err := ths.Model.Create(platform, postedData)
	if err != nil {
		response.Err(c, err.Error())
		return
	}
	if ths.CreateAfter != nil {
		ths.CreateAfter(c, &postedData) //不可取消
	}
	if ths.SaveAfter != nil {
		ths.SaveAfter(c, &postedData)
	}
	response.Result(c, struct {
		ID uint64 `json:"id"`
	}{
		ID: id,
	})
}

// Update 修改
func (ths *Backend) Update(c *gin.Context) {
	if !ths.OptionUpdate {
		response.Err(c, "缺少权限")
		return
	}
	postedData := request.GetPostedData(c)
	if ths.UpdateBefore != nil {
		if err := ths.UpdateBefore(c, &postedData); err != nil {
			response.Err(c, err.Error())
			return
		}
	}
	if ths.SaveBefore != nil {
		if err := ths.SaveBefore(c, &postedData); err != nil {
			response.Err(c, err.Error())
			return
		}
	}
	if ths.Validator != nil {
		if err := ths.Validator(postedData); err != nil {
			response.Err(c, err.Error())
			return
		}
	}
	platform := ths.getPlatform(c)
	if err := ths.Model.Update(platform, postedData); err != nil {
		response.Err(c, err.Error())
		return
	}
	if ths.UpdateAfter != nil {
		ths.UpdateAfter(c, &postedData)
	}
	if ths.SaveAfter != nil {
		ths.SaveAfter(c, &postedData)
	}
	response.Ok(c)
}

// Delete 删除
func (ths *Backend) Delete(c *gin.Context) {
	if !ths.OptionDelete {
		response.Err(c, "缺少权限")
		return
	}
	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Err(c, err.Error())
		return
	}
	platform := ths.getPlatform(c)
	// 先查找相应的记录
	cond := builder.And(builder.Eq{"id": id})
	row := ths.Row()
	exists, err := ths.Model.Find(platform, row, cond)
	if err != nil {
		response.Err(c, err.Error())
		return
	}
	if !exists {
		response.Err(c, "删除失败: 此记录不存在")
		return
	}
	if ths.DeleteBefore != nil {
		if err := ths.DeleteBefore(c, row); err != nil {
			response.Err(c, err.Error())
			return
		}
	}
	err = ths.Model.Delete(platform, idStr)
	if err != nil {
		response.Err(c, err.Error())
		return
	}
	if ths.DeleteAfter != nil {
		ths.DeleteAfter(c, row)
	}
	response.Ok(c)
}

// All 获取所有id-name样式
func (ths *Backend) All(c *gin.Context) {
	if !ths.OptionAll {
		response.Err(c, "缺少权限")
		return
	}
	platform := ths.getPlatform(c)
	rd := Redis(platform)
	RedisRestore(platform, rd)
	if values := caches.Global.Get(platform, rd, ths.Model.GetTableName(), func() interface{} {
		rows := &[]IdName{}
		_, _ = ths.Model.FindAll(platform, rows, nil, 100)
		return rows
	}); values != nil {
		response.Result(c, values)
		return
	}
	response.Err(c, "获取失败")
}

// 得到平台识别号
func (ths *Backend) getPlatform(c *gin.Context) string {
	if ths.Platform != "" {
		return ths.Platform
	}
	return request.GetPlatform(c)
}
