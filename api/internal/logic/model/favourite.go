package model

import (
	"douyinFavoriteList_4/api/internal/logic/mysqlop"
	"douyinFavoriteList_4/api/utool"
	"fmt"
	"time"
)

type Favourite struct {
	Id         int64     `db:"id"`          // 点赞表的主键
	UserId     int64     `db:"user_id"`     // 点赞人的id
	VideoId    int64     `db:"video_id"`    // 视频id
	Status     int64     `db:"status"`      // 点赞状态1为点赞2为取消
	CreateTime time.Time `db:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time"` // 更新时间
}

func (f Favourite) GetFavouritesByUid(uid int64) []Favourite {
	db, _ := mysqlop.GetDB()
	//1.sql语句
	sqlStr := "select * from favourite where user_id=  ?;"
	//2.执行
	rows, err := db.Query(sqlStr, uid)
	if err != nil {
		fmt.Printf("%s query failed,err:%v\n", sqlStr, err)
		return make([]Favourite, 0)
	}
	//3一定要关闭rows
	defer rows.Close()
	//4.循环取值
	var u Favourite
	var F []Favourite
	for rows.Next() {
		var CreateTimeCache []uint8
		var UpdateTimeVavhe []uint8
		rows.Scan(&u.Id, &u.UserId, &u.VideoId, &u.Status, &CreateTimeCache, &UpdateTimeVavhe)
		u.CreateTime = utool.Uint82Time(CreateTimeCache)
		u.UpdateTime = utool.Uint82Time(UpdateTimeVavhe)

		//fmt.Printf("u1:%#v\n", u)
		F = append(F, u)
	}
	return F
}
func (f Favourite) GetFavouritesByUid_grom(uid int64) []Favourite {
	var fs []Favourite
	db, _ := mysqlop.GetDBGrom()
	db.Where("user_id=  ?", uid).Find(&fs)
	return fs
}
func (f Favourite) Print(text string) {
	println(text)
	println("vvvvvvvvvvvvvvvvvvvv\n")
	fmt.Printf("Id        %v\nUserId    %v\nVideoId   %v\nStatus    %v\nCreateTime%v\nUpdateTime%v",
		f.Id,
		f.UserId,
		f.VideoId,
		f.Status,
		f.CreateTime,
		f.UpdateTime)
	println("\n^^^^^^^^^^^^^^^^^^")
}
