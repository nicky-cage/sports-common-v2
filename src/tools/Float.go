package tools

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Round 四舍五入
func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// ToRound
func Fixed(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

// ToFixed 转换为2位小数1
func ToFixed(f float64, precision int) float64 {
	if precision == 0 {
		precision = 2
	}

	output, err := GetFloatStringPrecise(f, precision)
	if err != nil {
		fmt.Println(err.Error())
	}
	return output
}

// Decimal 转换为小数
func Decimal(value float64, fixNum string) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+fixNum+"f", value), 64)
	return value
}

// GetFloatString 返回浮点型的字符串,有精度问题 比如12.456 截取2位小数点是12.46
func GetFloatString(value float64, fixNum string) string {
	if fixNum == "0" {
		return fmt.Sprintf("%f", value)
	} else {
		return fmt.Sprintf("%."+fixNum+"f", value)
	}
}

// GetFloatStringPrecise 通过位数对比，并用字符串"."来切割来保证精度
func GetFloatStringPrecise(f float64, m int) (float64, error) {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return 0.00, errors.New("f转n出错,n是空值")
	}
	if strings.Index(n, ".") > 0 {
		newn := strings.Split(n, ".")
		if len(newn) < 2 || m >= len(newn[1]) {
			if m == len(newn[1]) {
				return f, nil
			}

			f0 := strings.Repeat("0", m-len(newn[1]))

			n = n + f0
		} else {
			f0 := ""
			if m >= len(newn[1]) {
				f0 = strings.Repeat("0", m-len(newn[1]))
			}
			n = newn[0] + "." + newn[1][:m] + f0
		}
	} else {
		return f, nil
	}
	v, err := strconv.ParseFloat(n, 64)
	return v, err
}
