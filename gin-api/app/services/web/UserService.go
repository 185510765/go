package service_web

import (
	"gin-api/app/common/validate"
)

// 校验email
func ValidateEmail(email string) (int, string) {
	if email == "" {
		return 0, "邮箱不能为空"
	}

	isEmail := validate.IsEmail(email)
	if !isEmail {
		return 0, "邮箱格式不正确"
	}

	return 1, ""
}
