package validator

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"goWebDemo/utils/errmsg"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	unt := ut.New(zh_Hans_CN.New())
	trans, _ := unt.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err.Error(), errmsg.Error
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.Error
		}
	}
	return "", errmsg.Success
}