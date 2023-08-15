package mysqlop

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var f = 0

var db *sql.DB //连接池对象

func GetDB() (*sql.DB, error) {

	if f == 1 {
		return db, nil
	}
	f = 1
	//数据库
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := "root:abc123@tcp(127.0.0.1:3306)/douyin"
	//连接数据集
	var err error
	db, err = sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		return nil, err
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		return nil, err
	}
	fmt.Println("连接数据库成功~")
	//设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10)
	//defer func() {
	//	err := db.Close()
	//	if err != nil {
	//		return
	//	}
	//	f = 0
	//}()

	return db, nil
}
