package utils

/*
正则常量
*/
const (
	// PhoneRegex 手机号正则
	PhoneRegex string = "^1([38][0-9]|4[579]|5[0-3,5-9]|6[6]|7[0135678]|9[89])\\d{8}$"
	// EmailRegex 邮箱正则
	EmailRegex string = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	// PasswordRegex 密码正则,6~32位的字母、数字、下划线
	PasswordRegex string = "^\\w{6,32}$"
)
