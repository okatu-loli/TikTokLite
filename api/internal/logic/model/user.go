package model

import (
	"database/sql"
	"douyinFavoriteList_4/api/internal/logic/mysqlop"
	"douyinFavoriteList_4/api/utool"

	"fmt"
	"strconv"
	"time"
)

type User struct {
	Id            int64        `db:"id"`             // 用户的主键id
	UserName      string       `db:"user_name"`      // 用户名设置二级索引
	Password      string       `db:"password"`       // 用户的登录密码
	FollowCount   int64        `db:"follow_count"`   // 用户关注数
	FollowerCount int64        `db:"follower_count"` // 用户粉丝数
	CreateTime    time.Time    `db:"create_time"`    // 创建时间
	UpdateTime    time.Time    `db:"update_time"`    // 更新时间
	DeletedAt     sql.NullTime `db:"deleted_at"`     // 逻辑删除
}

func (u User) GetUserById(uid int64) User {
	db, _ := mysqlop.GetDB()
	//1.sql语句
	sqlStr := "select * from user where id =1 or id = ?;"
	//2.执行
	rows, err := db.Query(sqlStr, uid)
	if err != nil {
		fmt.Printf("%s query failed,err:%v\n", sqlStr, err)
		return User{}
	}
	//3一定要关闭rows

	var user User

	//4.循环取值
	for rows.Next() {
		var CreateTimeCache []uint8
		var UpdateTimeCache []uint8
		var DeletedAtCache []uint8
		_ = rows.Scan(&user.Id, &user.UserName, &user.Password, &user.FollowCount, &user.FollowerCount, &CreateTimeCache, &UpdateTimeCache, &DeletedAtCache)
		user.CreateTime = utool.Uint82Time(CreateTimeCache)
		user.UpdateTime = utool.Uint82Time(UpdateTimeCache)
		user.DeletedAt = utool.Uint82NullTime(DeletedAtCache)

		//fmt.Printf("u1:%#v\n", user)
	}
	defer rows.Close()
	return user
}
func (u User) GetAvatar() string {
	return ""
}
func (u User) GetBackgroundImage() string {
	return ""
}
func (u User) GetUserFavoriteCount() int64 {
	var f Favourite
	return int64(len(f.GetFavouritesByUid(u.Id)))

}

func (u User) GetUserTotalFavorited() string {
	var f Favourite
	return strconv.Itoa(len(f.GetFavouritesByUid(u.Id)))

}
func (u User) GetSignature() string {
	return ""
}
func (u User) GetWorkCount() int64 {
	var video Video
	videos := video.GetVideosByAid(u.Id)
	var count int64 = 0
	for _, v := range videos {
		count += v.FavoriteCount
	}
	return count
}
func (u User) Print(text string) {
	println(text)
	println("vvvvvvvvvvvvvvvvvvvv\n")
	fmt.Printf("Id            %v\nUserName      %v\nPassword      %v\nFollowCount   %v\nFollowerCount %v\nCreateTime    %v\nUpdateTime    %v\nDeletedAt     %v",
		u.Id,
		u.UserName,
		u.Password,
		u.FollowCount,
		u.FollowerCount,
		u.CreateTime,
		u.UpdateTime,
		u.DeletedAt)
	println("\n^^^^^^^^^^^^^^^^^^")
}
