package model

import (
	"database/sql"
	"douyinFavoriteList_4/api/internal/logic/mysqlop"
	"douyinFavoriteList_4/api/utool"

	"fmt"
	"time"
)

type Video struct {
	Id            int64        `db:"id"`             // 作品的id
	Title         string       `db:"title"`          // 作品标题
	AuthorId      int64        `db:"author_id"`      // 作者的id
	PlayUrl       string       `db:"play_url"`       // 视频资源的url
	CoverUrl      string       `db:"cover_url"`      // 封面的url
	FavoriteCount int64        `db:"favorite_count"` // 点赞数量
	CommentCount  int64        `db:"comment_count"`  // 评论数量
	CreateTime    time.Time    `db:"create_time"`    // 创建时间
	UpdateTime    time.Time    `db:"update_time"`    // 更新时间
	DeletedAt     sql.NullTime `db:"deleted_at"`     // 逻辑删除
}

func (f Video) GetVideosByVid_grom(vid int64) []Video {
	var fs []Video
	db, _ := mysqlop.GetDBGrom()
	db.Where("id=?", vid).Find(&fs)
	return fs
}

func (f Video) GetVideosByVid(vid int) []Video {
	db, _ := mysqlop.GetDB()
	//1.sql语句
	sqlStr := "select * from Video where id=?;"
	//2.执行
	rows, err := db.Query(sqlStr, vid)
	if err != nil {
		fmt.Printf("%s query failed,err:%v\n", sqlStr, err)
		return make([]Video, 0)
	}
	//3一定要关闭rows
	defer rows.Close()
	//4.循环取值
	var u Video
	var F []Video
	for rows.Next() {
		var CreateTimeCache []uint8
		var UpdateTimeCavhe []uint8
		var DeletedAtCavhe []uint8
		rows.Scan(&u.Id, &u.Title, &u.AuthorId, &u.PlayUrl, &u.CoverUrl, &u.FavoriteCount, &u.CommentCount, &CreateTimeCache, &UpdateTimeCavhe, &DeletedAtCavhe)
		u.CreateTime = utool.Uint82Time(CreateTimeCache)
		u.UpdateTime = utool.Uint82Time(UpdateTimeCavhe)
		u.DeletedAt = utool.Uint82NullTime(DeletedAtCavhe)

		fmt.Printf("u1:%#v\n", u)
		F = append(F, u)
	}
	return F
}

func (f Video) GetVideosByAid_grom(aid int64) []Video {
	var fs []Video
	db, _ := mysqlop.GetDBGrom()
	db.Where("author_id=?", aid).Find(&fs)
	return fs
}

func (f Video) GetVideosByAid(aid int64) []Video {
	db, _ := mysqlop.GetDB()
	//1.sql语句
	sqlStr := "select * from Video where author_id=?;"
	//2.执行
	rows, err := db.Query(sqlStr, aid)
	if err != nil {
		fmt.Printf("%s query failed,err:%v\n", sqlStr, err)
		return make([]Video, 0)
	}
	//3一定要关闭rows
	defer rows.Close()
	//4.循环取值
	var u Video
	var F []Video
	for rows.Next() {
		var CreateTimeCache []uint8
		var UpdateTimeCavhe []uint8
		var DeletedAtCavhe []uint8
		rows.Scan(&u.Id, &u.Title, &u.AuthorId, &u.PlayUrl, &u.CoverUrl, &u.FavoriteCount, &u.CommentCount, &CreateTimeCache, &UpdateTimeCavhe, &DeletedAtCavhe)
		u.CreateTime = utool.Uint82Time(CreateTimeCache)
		u.UpdateTime = utool.Uint82Time(UpdateTimeCavhe)
		u.DeletedAt = utool.Uint82NullTime(DeletedAtCavhe)

		//fmt.Printf("u1:%#v\n", u)
		F = append(F, u)
	}
	return F
}

func (v Video) Print(text string) {
	println(text)
	println("vvvvvvvvvvvvvvvvvvvv\n")
	fmt.Printf("Id            %v\nTitle         %v\nAuthorId      %v\nPlayUrl       %v\nCoverUrl      %v\nFavoriteCount %v\nCommentCount  %v\nCreateTime    %v\nUpdateTime    %v\nDeletedAt     %v",
		v.Id,
		v.Title,
		v.AuthorId,
		v.PlayUrl,
		v.CoverUrl,
		v.FavoriteCount,
		v.CommentCount,
		v.CreateTime,
		v.UpdateTime,
		v.DeletedAt)
	println("\n^^^^^^^^^^^^^^^^^^")
}
