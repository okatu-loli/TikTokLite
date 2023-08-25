package model

import (
	"douyinFavoriteList_4/api/internal/logic/mysqlop"
	"fmt"
	"testing"
)

func TestUser_GetUserById_grom(t *testing.T) {
	var fs User
	db, _ := mysqlop.GetDBGrom()
	db.Where(1).Find(&fs)
	fmt.Println(fs)
}
