package response

import (
	"github.com/flosch/pongo2"
)

// 模板函数
var templateFuncs = map[string]func(...*pongo2.Value) *pongo2.Value{}

// LoadTemplateFuncs 加载模板函数
func LoadTemplateFuncs(fArr map[string]func(...*pongo2.Value) *pongo2.Value) {
	for k, f := range fArr {
		templateFuncs[k] = f
	}
}
