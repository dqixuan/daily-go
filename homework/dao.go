package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)


type User struct {
	ID int64
	Name string
	Age int32
}


var Db *sqlx.DB

func MustInitDb() {
	var err error
	Db, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mytest")
	if err != nil {
		panic(err)
	}
}

const UserTable = "user"

// 根据ID获取用户信息
func GetUserByID(db *sqlx.DB, userID int64) (user User, err error) {
	q := fmt.Sprintf("select I_ID, CH_NAME, I_AGE from %s where I_ID = ?", UserTable)
	err = db.Get(&user, q, userID)
	if err != nil {
		err = fmt.Errorf("sql=%s userID=%d, err=%w", q, userID, err)
		return
	}
	return
}
