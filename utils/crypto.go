package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(text string) (string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
    panic("加密失败" + err.Error())
	}
	return string(hash)
}

// 比对数据库存储的 hash 值与用户输入的密码是否匹配
func CompareHashAndPassword(hash, password string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil 
}
