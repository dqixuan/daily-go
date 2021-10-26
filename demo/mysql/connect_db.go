package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

/**
	1、输入的MySQL配置格式
		//mysql dsn格式
		//涉及参数:
		//username   数据库账号
		//password   数据库密码
		//host       数据库连接地址，可以是Ip或者域名
		//port       数据库端口
		//Dbname     数据库名
		username:password@tcp(host:port)/Dbname?charset=utf8&parseTime=True&loc=Local
	2、import mysql驱动
		_ "github.com/go-sql-driver/mysql"
 */

func InitDb() *sqlx.DB {
	var err error
	DB, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/wenda?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return DB
}

type User struct {
	ID int `db:"id"`
	Name string `db:"name"`
	Password string `db:"password"`
	Salt string `db:"salt"`
	HeadUrl string `db:"headUrl"`
}

func getUsrs(db *sqlx.DB) (result []*User, err error) {
	q := fmt.Sprintf("select * from %s", "user")
	err = db.Select(&result, q)
	return
}
