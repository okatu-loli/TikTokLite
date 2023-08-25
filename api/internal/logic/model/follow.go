package model

import (
	"database/sql"
	"douyinFavoriteList_4/api/internal/logic/mysqlop"
	"douyinFavoriteList_4/api/utool"

	"fmt"
	"time"
)

type Follow struct {
	Id         int64        `db:"id"`          // 关注表的主键id
	UserId     int64        `db:"user_id"`     // 关注人的id
	ToUserId   int64        `db:"to_user_id"`  // 被关注人id
	Status     int64        `db:"status"`      // 是否为互相关注
	CreateTime time.Time    `db:"create_time"` // 创建时间
	UpdateTime time.Time    `db:"update_time"` // 更新时间
	DeletedAt  sql.NullTime `db:"deleted_at"`  // 逻辑删除
}

func (f Follow) GetFollowsByUids_grom(uid int64, touid int64) Follow {
	var fs Follow
	db, _ := mysqlop.GetDBGrom()
	db.Where("user_id=  ? and to_user_id=?", uid, touid).Find(&fs)
	return fs
}

func (f Follow) GetFollowsByUids(uid int64, touid int64) Follow {
	db, _ := mysqlop.GetDB()
	//1.sql语句
	sqlStr := "select * from Follow where user_id=? and to_user_id=?;"
	//2.执行
	rows, err := db.Query(sqlStr, uid, touid)
	if err != nil {
		fmt.Printf("%s query failed,err:%v\n", sqlStr, err)
		return Follow{}
	}
	//3一定要关闭rows
	defer rows.Close()
	//4.循环取值
	var u Follow
	u.Id = -1

	for rows.Next() {
		var CreateTimeCache []uint8
		var UpdateTimeCavhe []uint8
		var DeletedAtCavhe []uint8
		err := rows.Scan(&u.Id, &u.UserId, &u.ToUserId, &u.Status, &CreateTimeCache, &UpdateTimeCavhe, &DeletedAtCavhe)
		if err != nil {
			return u
		}
		u.CreateTime = utool.Uint82Time(CreateTimeCache)
		u.UpdateTime = utool.Uint82Time(UpdateTimeCavhe)
		u.DeletedAt = utool.Uint82NullTime(DeletedAtCavhe)

		//fmt.Printf("u1:%#v\n", u)

	}
	return u
}

func (f Follow) Print(text string) {
	println(text)
	println("vvvvvvvvvvvvvvvvvvvv\n")
	fmt.Printf("Id        %v\nUserId    %v\nToUserId  %v\nStatus    %v\nCreateTime%v\nUpdateTime%v\nDeletedAt %v",
		f.Id,
		f.UserId,
		f.ToUserId,
		f.Status,
		f.CreateTime,
		f.UpdateTime,
		f.DeletedAt)
	println("\n^^^^^^^^^^^^^^^^^^")
}
