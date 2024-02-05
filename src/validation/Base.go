package validation

import (
	"errors"
	"fmt"
	"regexp"
	"sports-common/log"
	"strconv"
	"strings"
)

// Validator 校验器
type Validator struct {
	Data      map[string]interface{} //要校验的字段
	Errors    []string               //错误信息
	FieldName string                 //当前字段
	Required  bool                   //当前字段是否允许为空
}

// New 生成校验器
func New(postedData map[string]interface{}) *Validator {
	return &Validator{
		Data:     postedData,
		Required: true,
	}
}

// Null 当前字段是否允许为空值
func (ths *Validator) Null(args ...bool) *Validator {
	argCount := len(args)
	if argCount >= 1 {
		ths.Required = args[0]
	} else {
		ths.Required = false //默认不允许为空值
	}
	return ths
}

// Field 设置当前的字段名称
func (ths *Validator) Field(name string) *Validator {
	ths.FieldName = name
	return ths
}

// AppendError 追加错误信息
func (ths *Validator) AppendError(message string) *Validator {
	ths.Errors = append(ths.Errors, message)
	return ths
}

// Check 检测各个字段
func (ths *Validator) Check(message string, callback func(interface{}) error) *Validator {
	val, exists := ths.Data[ths.FieldName]
	if !exists {
		if !ths.Required {
			return ths
		}
		return ths.AppendError(message)
	}
	if err := callback(val); err != nil {
		log.Err("\n格式化时出错: %v\n字段: %v\n提交数据: %v\n\n", err, ths.FieldName, ths.Data)
		ths.AppendError(message)
	}
	return ths
}

// CheckReg 检测正则表达式
func (ths *Validator) CheckReg(regex string, message string) *Validator {
	val, exists := ths.Data[ths.FieldName]
	if !exists {
		if !ths.Required {
			return ths
		}
		return ths.AppendError(message)
	}
	matched, err := regexp.MatchString(regex, val.(string))
	if err != nil {
		ths.AppendError(err.Error())
		return ths
	}
	if !matched {
		ths.AppendError(message)
	}
	return ths
}

// Length 字段长度必须介于二者之间
func (ths *Validator) Length(min int, max int, message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		length := len([]rune(val.(string)))
		if !ths.Required && length == 0 { //如果不是必须,并且长度为0
			return nil
		}
		if length < min || length > max {
			return errors.New("字符串长度不在区间范围之内")
		}
		return nil
	})
}

// Numeric 是数字, 可以为整数, 负数, 小数
func (ths *Validator) Numeric(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		_, err := strconv.ParseFloat(fmt.Sprintf("%v", val), 10)
		return err
	})
}

// Numeric 是数字, 可以为整数,小数,大于0
func (ths *Validator) NumericGt0(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		if num, err := strconv.ParseFloat(fmt.Sprintf("%v", val), 10); err == nil && num > 0 {
			return nil
		}
		return errors.New("不是大于0的数")
	})
}

// Numeric 是数字, 可以为整数,小数,大于等于0
func (ths *Validator) NumericEq0(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		if num, err := strconv.ParseFloat(fmt.Sprintf("%v", val), 10); err == nil && num >= 0 {
			return nil
		}
		return errors.New("不是大于0的数")
	})
}

// DateTime 必须是日期时间格式
func (ths *Validator) DateTime(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		matched, err := regexp.MatchString(`\d{4}\-\d{2}\-\d{2} \d{2}:\d{2}:\d{2}`, val.(string))
		if !matched {
			return errors.New("日期时间格式错误")
		}
		return err
	})
}

// Date 必须是日期格式
func (ths *Validator) Date(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		matched, err := regexp.MatchString(`\d{4}\-\d{2}\-\d{2}`, val.(string))
		if !matched {
			return errors.New("日期时间格式错误")
		}
		return err
	})
}

// Time 必须是时间格式
func (ths *Validator) Time(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		matched, err := regexp.MatchString(`\d{2}:\d{2}:\d{2}`, val.(string))
		if !matched {
			return errors.New("日期时间格式错误")
		}
		return err
	})
}

// Int 是整数
func (ths *Validator) Int(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		_, err := strconv.Atoi(fmt.Sprintf("%v", val))
		return err
	})
}

// Uint0 是正整数-包括零
func (ths *Validator) Uint0(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		_, err := strconv.ParseUint(fmt.Sprintf("%v", val), 10, 64)
		return err
	})
}

// Uint 正整数非零
func (ths *Validator) Uint(message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		if num, err := strconv.ParseUint(fmt.Sprintf("%v", val), 10, 64); err == nil && num > 0 {
			return nil
		}
		return errors.New("不是非零的正整数")
	})
}

// Mobile 是手机号码
func (ths *Validator) Mobile(args ...string) *Validator {
	message := "手机号码格式不正确"
	if len(args) >= 1 {
		message = args[0]
	}
	return ths.CheckReg(`1[356789]{1}\d{9}`, message)
}

// Mail 判断是否是电子邮件
func (ths *Validator) Mail(args ...string) *Validator {
	message := "电子邮件格式不正确"
	if len(args) >= 1 {
		message = args[0]
	}
	return ths.CheckReg(`[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+`, message)
}

// Password 必须是密码
func (ths *Validator) Password(args ...string) *Validator {
	message := "密码格式不正确"
	if len(args) >= 1 {
		message = args[0]
	}
	return ths.Check(message, func(val interface{}) error {
		strLen := len(val.(string))
		if strLen > 20 || strLen < 6 {
			return errors.New("密码需要6-20位")
		}
		return nil
	})
}

// Equal 两个字段必须相等
func (ths *Validator) Equal(field string, args ...string) *Validator {
	message := "确认内容与校验内容不一致"
	if len(args) >= 1 {
		message = args[0]
	}
	return ths.Check(message, func(val interface{}) error {
		v := val.(string)
		vField, exists := ths.Data[field]
		if !exists {
			return errors.New("两次输入的值不一致")
		}
		if v != vField {
			return errors.New("两次内容不一致")
		}
		return nil
	})
}

// Gender 性别
func (ths *Validator) Gender(args ...string) *Validator {
	message := "性别选择不正确"
	if len(args) >= 1 {
		message = args[0]
	}
	return ths.Check(message, func(val interface{}) error {
		v, err := strconv.Atoi(fmt.Sprintf("%v", val))
		if err != nil {
			return err
		}
		if v < 0 || v > 2 {
			return errors.New("性别不在合法区间之内")
		}
		return nil
	})
}

// State 必须是状态值
func (ths *Validator) State(args ...string) *Validator {
	message := "状态值无效"
	if len(args) >= 1 {
		message = args[0]
	}
	return ths.Check(message, func(val interface{}) error {
		sta := fmt.Sprintf("%v", val)
		if sta == "1" || sta == "0" {
			return nil
		}
		return errors.New("状态值无效")
	})
}

// InValues 值在区间范围之内
func (ths *Validator) InValues(values []string, message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		value := val.(string)
		for _, v := range values {
			if value == v {
				return nil
			}
		}
		return errors.New("值不在区间范围之内")
	})
}

// InIntValues 值在区间范围之内(整形)
func (ths *Validator) InIntValues(values []int, message string) *Validator {
	return ths.Check(message, func(val interface{}) error {
		value := val.(int)
		for _, v := range values {
			if value == v {
				return nil
			}
		}
		return errors.New("值不在区间范围之内")
	})
}

// BankCard 银行卡号
func (ths *Validator) BankCard(args ...string) *Validator {
	message := "银行卡号格式不正确"
	if len(args) >= 1 {
		message = args[0]
	}
	return ths.CheckReg(`([1-9]{1})(\d{14}|\d{18})`, message)
}

// UserName 用户名称
func (ths *Validator) UserName(args ...string) *Validator {
	message := "用户名称格式不正确"
	if len(args) >= 1 {
		message = args[0]
	}
	return ths.CheckReg(`[A-Za-z]{1}[A-Za-z0-9]{4,17}`, message)
}

// Validate 生成校验结果
func (ths *Validator) Validate() error {
	if len(ths.Errors) == 0 {
		return nil
	}
	return errors.New(strings.Join(ths.Errors, ","))
}
