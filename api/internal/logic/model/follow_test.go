package model

import (
	"douyinFavoriteList_4/api/internal/logic/mysqlop"
	"fmt"
	"testing"
)

func TestFollow_GetFollowsByUids_grom(t *testing.T) {
	var fs Follow
	db, _ := mysqlop.GetDBGrom()
	db.Where("user_id=  ? and to_user_id=?", 5, 1).Find(&fs)
	fmt.Println(fs)
}
