package mysqlop

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var f = 0

var db *sql.DB //连接池对象

var (
	user     = "root"
	pwd      = "abc123"
	ipport   = "127.0.0.1:3306"
	database = "douyin"
	dsn      = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, ipport, database)
)

func GetDB() (*sql.DB, error) {

	if f == 1 {
		return db, nil
	}
	f = 1
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

	return db, nil
}

var g *gorm.DB //连接池对象
var gf bool = false

func GetDBGrom() (*gorm.DB, error) {
	if gf {
		return g, nil
	}
	g, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println("连接失败:", err, "\n", dsn)
		return nil, err
	}
	fmt.Println("连接成功:", dsn)
	return g, nil
}
