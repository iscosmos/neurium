package validator

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/iscosmos/neurium/pkg/code"
	"reflect"
	"strings"
)

func Validate(data interface{}) error {
	zhCN := zh.New()

	// 注册一个提取json的tag作为翻译名字的函数
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// 注册翻译名字的中文翻译器
	uni := ut.New(zhCN)
	trans, _ := uni.GetTranslator("zh")
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)

	// 校验数据
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return code.NewCode(40010, err.Translate(trans))
		}
	}

	return nil
}
