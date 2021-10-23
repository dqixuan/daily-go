package main

import (
	"strconv"
	"strings"
)

func main() {
	
}

/**
	ipv4   192.0.0.1
	数据库可以直接存储字符串类型，1.1.1.1  最少7位  192.192.193.192 15位
	字符串类型 varchar(15)
	可以将ip地址转换为uint(32)类型， 占用8位
	节省空间
 */
func transferIp2Int(ip string) int64 {
	partSlic := strings.Split(ip, ".")
	p0, _:= strconv.ParseInt(partSlic[0], 10, 64)
	p1, _:= strconv.ParseInt(partSlic[1], 10, 64)
	p2, _:= strconv.ParseInt(partSlic[2], 10, 64)
	p3, _:= strconv.ParseInt(partSlic[3], 10, 64)

	var sum int64
	sum += p0 << 24
	sum += p1 << 16
	sum += p2 << 8
	sum += p3

	return sum
}
