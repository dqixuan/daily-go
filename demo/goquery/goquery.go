package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	/**
	url:="http://www.baidu.com"
	   dom,err:=goquery.NewDocument(url)
	   if err != nil {
	      log.Fatalln(err)
	   }
	   //使用find()函数查找元素"div"、类".class"或ID"#id", 并使用Each遍历所有匹配结果赋值给selection,
	   dom.Find("div").Each(func(i int, selection *goquery.Selection){
	      fmt.Println(selection.Text())
	   })
	 */

	dom, err := goquery.NewDocument("http://album.zhenai.com/u/1790536836")
	if err != nil {
		panic("NewDocumentFromReader")
	}
	dom.Find("div[class=purple-btns").Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
