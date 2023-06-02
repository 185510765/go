package validator

//gin > 1.4.0

//将验证器错误翻译成中文
// 1、同时返回多个字段的错误，可通过遍历错误信息，只返回一条
// 2、返回的错误信息为英文的，要转译成中文
// 3、对应错误信息的字段名称为英文，需要转为中文

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Init() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	// 获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	// 注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

	// 将验证法字段名 映射为中文名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
}

// Translate 翻译错误信息 返回单条
func Translate(err error) string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}

	return fmt.Sprint(errors[0].Translate(trans)) // 多条验证 返回第一条
}

// // Translate 翻译错误信息 返回多条
// func Translate(err error) map[string][]string {
// 	var result = make(map[string][]string)
// 	errors := err.(validator.ValidationErrors)

// 	for _, err := range errors {
// 		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
// 	}

// 	return result
// }
