package tool

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptByString 将字符串加密为bcrypt散列
func BcryptByString(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// BcryptByBytes 将字节切片加密为bcrypt散列
func BcryptByBytes(b []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash 用于比较密码与其散列版本
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
