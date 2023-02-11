package utils

import (
	"github.com/duke-git/lancet/v2/validator"
	"regexp"
)

// 校验是否不符合正则格式
func mismatch(str string, regex string) bool {
	if validator.IsEmptyString(str) {
		return true
	}
	matched, _ := regexp.MatchString(regex, str)

	return !matched
}

// PasswordInvalid 是否是无效密码格式
func PasswordInvalid(password string) bool {
	return mismatch(password, PasswordRegex)
}
