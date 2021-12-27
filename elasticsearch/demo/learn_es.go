package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"time"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetBasicAuth("user", "secret"),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetSniff(false),
		)
	if err != nil {
		fmt.Printf("connect failed: %v\n", err)
	} else {
		fmt.Printf("connect succeed\n")
	}

	ctx := context.TODO()
	// 写入
	//blog := Article{
	//	Title:     "golang es 教程",
	//	Content:   "go如何操作es",
	//	Author:    "tizi",
	//	CreatedAt: time.Now(),
	//}
	//put1, err := client.Index().Index("blogs").Id("1").BodyJson(blog).Do(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)

	// 查询
	get1, err := client.Get().Index("blogs").Id("1").Do(ctx)
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("result=%+v", get1)
	}

}




type Article struct {
	Title string
	Content string
	Author string
	CreatedAt time.Time
}
