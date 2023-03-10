package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 检查传入的用户信息是否合法
func UserInfoCheck(username string, password string) error {

	return nil
}

// Md5Encode MD5加密字符串
func Md5Encode(str string) string {
	secret := "secret"
	h := md5.New()
	h.Write([]byte(str + secret))
	return hex.EncodeToString(h.Sum(nil))
}

// CheckVideo 检查视频格式
func CheckVideo(str string) bool {
	videoExt := []string{"mp4"}
	for _, v := range videoExt {
		if str == v {
			return true
		}
	}
	return false
}
