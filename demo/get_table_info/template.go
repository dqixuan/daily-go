package main

import "html/template"

func openBrace() string {
	return "`"
}

var funsMap = template.FuncMap{
	"openBrace":openBrace,
}

var transferType = map[string]string {
	"bigint": "int64",
	"tinyint":"int32",
	"varchar":"string",
	"datetime": "time.Time",
	"text": "string",
}

type column struct {
	ColumnName string
	RawName string
	LowFirstColName string
	ColumnType string
	Comment string

}

type table struct {
	TableName string
	LowFirstTableName string
	Columns []column

}

var modelTemplate = template.Must(template.New("model").Funcs(funsMap).Parse(`
package model
type {{.TableName}} struct {
{{range $k, $v :=.Columns}}{{$v.ColumnName}} {{$v.ColumnType}} {{openBrace}}db:"{{$v.RawName}}"{{openBrace}} // {{$v.Comment}}
{{end}}}
`))


