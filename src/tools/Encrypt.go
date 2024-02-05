package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// 后台加密用到的KEY, 切勿改动
const SecretKey = "0jas3^128*@#*Ls]/<M';sdf3Dnbm?@$"

// 十六进制转十进制
func Hex2Dec(val string) (int, error) {
	n, err := strconv.ParseUint(val, 16, 32)
	return int(n), err
}

// 十进制转16进制
func Dec2Hex(val int) string {
	return fmt.Sprintf("%X", val)
}

// MD5 md5加密
func MD5(data string) string {
	h := md5.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Secret 得到一个Salt
func Secret() string {
	return RandString(32)
}

// GetPassword 依据password/salt生成密码
func GetPassword(password, salt string) string {
	key := fmt.Sprintf("%s:%s:%s", SecretKey, MD5(password), salt)
	return MD5(key)
}

/*
	CBC加密 按照golang标准库的例子代码
 	不过里面没有填充的部分,所以补上
*/

// PKCS7Padding 使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 补齐
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func SHA256(v string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}
