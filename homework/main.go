package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func main() {
	MustInitDb()
	fmt.Println("db init success")
	u, err := GetUserByID(Db, 10)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("对应用户不存在")
		return
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user=%+v", u)
}
