package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

/**
  name:丁其轩
  date:2021/6/17
  time:22:26
*/

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/campfire?charset=utf8")
	if err != nil {
		panic("db can not connect")
	}
}

func main() {
	m := tableInfo("activity")
	fmt.Printf("result=%+v", m)
	generateModel(m)
}

func toCamelName(s string) string {
	parts := strings.Split(s, "_")
	for i, v := range parts {
		parts[i] = strings.Title(strings.ToLower(v))
	}
	return strings.Join(parts, "")
}

func tableInfo(tableName string) (t table) {
	q := `select column_name fname, column_comment fdesc, data_type type from information_schema.columns  where table_name = ? and table_schema = ?`
	t.TableName = tableName
	rows, err := db.Query(q, tableName, "campfire")
	if err != nil {
		panic(fmt.Errorf("query: %w", err))
	}
	cols := []column{}
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
		cols = append(cols, column{
			RawName:    fname,
			ColumnName: toCamelColumnName(fname),
			ColumnType: transferType[typ],
			Comment:    fdesc,
		})
	}
	t.Columns = cols
	t.LowFirstTableName = toCamelName(tableName)
	return
}

func generateModel(t table) {
	pwd, _ := os.Getwd()
	output := pwd + "/model" + "/" + t.LowFirstTableName + ".autogen.model.go"
	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	err = modelTemplate.Execute(f, t)
	if err != nil {
		panic(fmt.Errorf("err2=%w", err))
	}
}

func toCamelColumnName(name string) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		parts[i] = strings.Title(strings.ToLower(part))
	}
	return strings.Join(parts[1:], "")
}

/**
	打印结果
 	id, , int
	name, , varchar
	age, , int
	sex, , tinyint
	word, , text
*/
