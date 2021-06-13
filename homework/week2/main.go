package week2

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)


/**
  name:丁其轩
  date:2021/6/13
  time:22:49
*/


func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return http.ListenAndServe(":8888", )
	})


	err := g.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}
