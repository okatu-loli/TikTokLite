package model

import (
	"douyinFavoriteList_4/api/internal/logic/mysqlop"
	"fmt"
	"testing"
)

func TestVideo_GetVideosByAid_grom(t *testing.T) {
	var fs []Video
	db, _ := mysqlop.GetDBGrom()
	db.Where("author_id=?", 15).Find(&fs)
	fmt.Println(fs)
}
