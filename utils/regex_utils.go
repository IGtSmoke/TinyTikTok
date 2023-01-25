package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// 校验是否不符合正则格式
func mismatch(str string, regex string) bool {
	str = strings.NewReplacer("\n", "", "\r", "", " ", "", "\t", "").Replace(str)
	if str == "" {
		return true
	}
	fmt.Print(str)
	matched, _ := regexp.MatchString(regex, str)

	return !matched
}

// PhoneInvalid 是否是无效手机格式
func PhoneInvalid(phone string) bool {
	return mismatch(phone, PhoneRegex)
}

// EmailInvalid 是否是无效邮箱格式
func EmailInvalid(email string) bool {
	return mismatch(email, EmailRegex)
}

// PasswordInvalid 是否是无效密码格式
func PasswordInvalid(password string) bool {
	return mismatch(password, PasswordRegex)
}
