package utils

import "encoding/base64"

// Base64Encode 加密 []byte -> string
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64Decode 解密 string -> []byte
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
