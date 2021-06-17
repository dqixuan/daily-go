package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
  name:丁其轩
  date:2021/6/17
  time:22:26
*/

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	if err != nil {
		panic("db can not connect")
	}
}

func main()  {
	m := tableInfo("person")
	fmt.Println(m)
}

func tableInfo(dbName string) map[string]string {
	q := `select column_name fname, column_comment fdesc, data_type type from information_schema.columns  where table_name = ?`
	var result = make(map[string]string)

	rows, err := db.Query(q, dbName)
	if err != nil {
		panic(fmt.Errorf("query: %w", err))
	}
	for rows.Next() {
		var fname, fdesc, typ string
		err = rows.Scan(&fname, &fdesc, &typ)
		if err != nil {
			panic("scan")
		}
		fmt.Printf(" %s, %s, %s\n", fname, fdesc, typ)
		if len(fdesc) == 0 {
			fdesc = fname
		}
		result[fname]  = fdesc
	}
	return result
}

/**
	打印结果
 	id, , int
	name, , varchar
	age, , int
	sex, , tinyint
	word, , text
 */
