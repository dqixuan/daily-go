package main

/**
    name:丁其轩
    date:2021/6/21
    time:22:47
*/

// 读取配置文件中的信息
type config struct {
	Connection string //
	Database string
	Acronym map[string]string
	IncludeTables []string
	ExtTable []extTable
}

type extTable struct {
	Name string // 表名
	FieldName string // 字段名称
	Columns []extColumn
}

type extColumn struct {
	Name string
	FieldName string
	CustomType string
}