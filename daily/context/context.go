package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1* time.Millisecond)
	defer cancel()

	select {
	case <- time.After(1 * time.Second):
		fmt.Println("time After work")
	case <- ctx.Done():
		fmt.Println("ctx Done work")
	}

	avg := makeAverager()
	fmt.Println(avg(10))
	fmt.Println(avg(30))
}

// 工程中 goroutine内结合for+select 针对context的事件进行处理，达到跨goroutine控制的目的
func text(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

/**
	context使用的正确姿势
	1、对第三方调用要传入context
	2、将context参数放在函数参数第一位， ctx context.Context，尽量不要将context封装在结构体中
	3、函数调用链必须传入context
	4、context的派生
		1）context.WithTimeOut、WithDeadline()、WithCancel
	5、context不能为空， 不知道传什么的时候用context.TODO()或 context.Background()
	6、context中仅传必要的值
 */

func httpRequest(ctx context.Context) error {
	req, err := http.NewRequest(http.MethodGet, "https://eddycjy.com/", nil)
	if err != nil {
		fmt.Println("http.NewRequest err=%v", err)
		return err
	}
	ctx, cancel := context.WithTimeout(req.Context(), 50 * time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http.DefaultClient.Do failed, err=%v", err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

// 闭包

func makeAverager() func(va float32) float32 {
	series := make([]float32, 0) // 内存逃逸到
	return func(val float32) float32 {
		series = append(series, val)
		total := float32(0)
		for _, v := range series {
			total += v
		}
		return total / float32(len(series))
	}
}
