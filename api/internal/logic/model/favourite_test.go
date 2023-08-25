package model

import (
	"douyinFavoriteList_4/api/internal/logic/mysqlop"
	"testing"
)

func TestFavourite_GetFavouritesByUid_grom(t *testing.T) {
	var fs []Favourite
	//var f Favourite
	db, _ := mysqlop.GetDBGrom()
	db.Where("user_id=  ?", 1).Find(&fs)
	for _, i := range fs {
		i.Print(">>>")
	}
}
