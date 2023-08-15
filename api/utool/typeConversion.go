package utool

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func Uint82Time(uint8Data []uint8) time.Time {
	// 假设你有一个 []uint8 数据

	// 将 []uint8 转化为字符串
	strData := string(uint8Data)

	// 使用 time.Parse 解析字符串为 time.Time 类型
	t, err := time.Parse("2006-01-02 15:04:05", strData)
	if err != nil {
		fmt.Println(">>>解析时间失败:", err)
		fmt.Println()
		return t
	}

	// 输出解析得到的时间
	//fmt.Println("解析得到的时间:", t)
	return t
}
func Uint82NullTime(uint8Data []uint8) sql.NullTime {
	t := Uint82Time(uint8Data)

	// 将 time.Time 转换为 sql.NullTime
	nullTime := sql.NullTime{
		Time:  t,
		Valid: true,
	}

	// 输出 sql.NullTime
	//fmt.Println("转换得到的 sql.NullTime:", nullTime)
	return nullTime
}
func S2I64(s string) int64 {
	// string到int64
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
func S2I(s string) int {
	// string到int64
	i, _ := strconv.Atoi(s)
	return i
}
func I2S(i int) string {
	// string到int64
	s := strconv.Itoa(i)
	return s
}
func I642S(i int64) string {
	// string到int64
	s := strconv.FormatInt(i, 10)
	return s
}
