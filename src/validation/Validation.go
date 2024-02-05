package validation

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	_ = zh.RegisterDefaultTranslations(validate, trans)
}

// Verify 执行校验
func Verify(obj interface{}) error {
	err := validate.Struct(obj)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		var errInfo string
		for _, e := range errs {
			errInfo += e.Translate(trans) + ";"
		}
		if errInfo != "" {
			return errors.New(errInfo)
		}
	}
	return err
}
